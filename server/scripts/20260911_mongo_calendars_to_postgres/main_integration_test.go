package main

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"timeful/server/models"
	pgstore "timeful/server/postgres"
	"timeful/server/utils"
)

// newCalendarRehearsalContext connects to the isolated Mongo and PostgreSQL
// test databases and drops a per-test Mongo source database.
func newCalendarRehearsalContext(t *testing.T, suffix string) (context.Context, *mongo.Database, *pgxpool.Pool) {
	t.Helper()
	mongoURI := os.Getenv("MONGODB_URI")
	baseDatabase := os.Getenv("MONGODB_DATABASE")
	postgresURI := os.Getenv("POSTGRES_APPLICATION_URI")
	if mongoURI == "" || postgresURI == "" || baseDatabase == "" {
		t.Skip("MONGODB_URI, MONGODB_DATABASE, and POSTGRES_APPLICATION_URI are required")
	}
	if baseDatabase != "timeful-test" && !strings.HasPrefix(baseDatabase, "timeful-test-") {
		t.Fatalf("requires an isolated test database, got %q", baseDatabase)
	}
	if !strings.Contains(postgresURI, "timeful-test") {
		t.Fatalf("requires an isolated PostgreSQL test database")
	}

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = mongoClient.Disconnect(ctx) })

	pool, err := pgxpool.New(ctx, postgresURI)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)

	database := mongoClient.Database(baseDatabase + "_calendar_backfill_" + suffix)
	if err := database.Drop(ctx); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = database.Drop(ctx) })
	return ctx, database, pool
}

// insertCalendarUser upserts a retained MongoDB user document.
func insertCalendarUser(t *testing.T, ctx context.Context, database *mongo.Database, user models.User) {
	t.Helper()
	raw, err := bson.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}
	var fields bson.M
	if err := bson.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	if _, err := database.Collection("users").ReplaceOne(ctx, bson.M{"_id": fields["_id"]}, user, options.Replace().SetUpsert(true)); err != nil {
		t.Fatal(err)
	}
}

// seedCalendarOwner creates the authoritative PostgreSQL account and platform
// identity a retained calendar document resolves through.
func seedCalendarOwner(t *testing.T, ctx context.Context, pool *pgxpool.Pool, externalID string) {
	t.Helper()
	if _, err := pgstore.NewRepository(pool).FindOrCreateAccount(ctx, externalID, pgstore.Account{Email: externalID + "@example.com", FirstName: "Owner"}); err != nil {
		t.Fatal(err)
	}
}

func cleanupCalendarBatch(t *testing.T, ctx context.Context, pool *pgxpool.Pool, batch string, externalIDs []string) {
	t.Helper()
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		if _, err := pool.Exec(cleanupCtx, `DELETE FROM migration_quarantine WHERE batch = $1`, batch); err != nil {
			t.Errorf("cleanup quarantine: %v", err)
		}
		if _, err := pool.Exec(cleanupCtx, `DELETE FROM migration_ledger WHERE batch = $1`, batch); err != nil {
			t.Errorf("cleanup ledger: %v", err)
		}
		if len(externalIDs) > 0 {
			// Accounts reference the platform identity without a cascade, so
			// remove them first; calendar rows cascade from the identity.
			if _, err := pool.Exec(cleanupCtx, `DELETE FROM accounts WHERE platform_identity_id IN (SELECT id FROM platform_identities WHERE external_user_id = ANY($1))`, externalIDs); err != nil {
				t.Errorf("cleanup accounts: %v", err)
			}
			if _, err := pool.Exec(cleanupCtx, `DELETE FROM platform_identities WHERE external_user_id = ANY($1)`, externalIDs); err != nil {
				t.Errorf("cleanup identities: %v", err)
			}
		}
	})
}

// calendarFixtures is the representative fixture set from the retained-data
// contract, including the missing-owner quarantine path.
type calendarFixtures struct {
	userIDs     []string
	externalIDs []string

	google    primitive.ObjectID
	outlook   primitive.ObjectID
	apple     primitive.ObjectID
	ics       primitive.ObjectID
	two       primitive.ObjectID
	prefsOnly primitive.ObjectID
	missing   primitive.ObjectID

	appleCiphertext string
}

func seedCalendarFixtures(t *testing.T, ctx context.Context, database *mongo.Database, pool *pgxpool.Pool) calendarFixtures {
	t.Helper()
	fixtures := calendarFixtures{
		google:    primitive.NewObjectID(),
		outlook:   primitive.NewObjectID(),
		apple:     primitive.NewObjectID(),
		ics:       primitive.NewObjectID(),
		two:       primitive.NewObjectID(),
		prefsOnly: primitive.NewObjectID(),
		missing:   primitive.NewObjectID(),
	}

	expiresAt := primitive.NewDateTimeFromTime(time.Now().Add(time.Hour).UTC())
	googleSubs := map[string]models.SubCalendar{
		"work":     {Name: "Work", Enabled: boolPointer(true)},
		"personal": {Name: "Personal", Enabled: boolPointer(false)},
		"hidden":   {Name: "Hidden"},
	}
	twoSubs := map[string]models.SubCalendar{
		"team": {Name: "Team"},
	}
	calendarOptions := &models.CalendarOptions{
		BufferTime:   models.BufferTimeOptions{Enabled: true, Time: 10},
		WorkingHours: models.WorkingHoursOptions{Enabled: true, StartTime: 9, EndTime: 17},
	}

	applePassword, err := utils.Encrypt("apple-secret")
	if err != nil {
		t.Fatal(err)
	}
	fixtures.appleCiphertext = applePassword

	googleKey := "google.owner@example.com_google"
	twoGoogleKey := "Two.User@Example.com_google"
	twoICSKey := "second@example.com_ics"

	users := []models.User{
		{
			Id:    fixtures.google,
			Email: "google.owner@example.com",
			CalendarAccounts: map[string]models.CalendarAccount{
				googleKey: {
					CalendarType: models.GoogleCalendarType,
					Email:        "google.owner@example.com",
					Picture:      "https://example.com/google.png",
					Enabled:      boolPointer(true),
					OAuth2CalendarAuth: &models.OAuth2CalendarAuth{
						AccessToken:           "google-access",
						RefreshToken:          "google-refresh",
						Scope:                 "calendar",
						AccessTokenExpireDate: expiresAt,
					},
					SubCalendars: &googleSubs,
				},
			},
			PrimaryAccountKey: stringPointer(googleKey),
			TokenOrigin:       models.WEB,
			CalendarOptions:   calendarOptions,
		},
		{
			Id:    fixtures.outlook,
			Email: "outlook.owner@example.com",
			CalendarAccounts: map[string]models.CalendarAccount{
				"outlook.owner@example.com_outlook": {
					CalendarType: models.OutlookCalendarType,
					Email:        "outlook.owner@example.com",
					OAuth2CalendarAuth: &models.OAuth2CalendarAuth{
						AccessToken:           "outlook-access",
						RefreshToken:          "outlook-refresh",
						Scope:                 "offline_access",
						AccessTokenExpireDate: expiresAt,
					},
				},
			},
		},
		{
			Id:    fixtures.apple,
			Email: "apple.owner@example.com",
			CalendarAccounts: map[string]models.CalendarAccount{
				"apple.owner@example.com_apple": {
					CalendarType: models.AppleCalendarType,
					Email:        "apple.owner@example.com",
					Enabled:      boolPointer(true),
					AppleCalendarAuth: &models.AppleCalendarAuth{
						Email:    "apple.owner@example.com",
						Password: applePassword,
					},
				},
			},
		},
		{
			Id:    fixtures.ics,
			Email: "ics.owner@example.com",
			CalendarAccounts: map[string]models.CalendarAccount{
				"My Feed_ics": {
					CalendarType: models.ICSCalendarType,
					Email:        "My Feed",
					Enabled:      boolPointer(true),
					ICSCalendarAuth: &models.ICSCalendarAuth{
						FeedURL: "https://example.com/feed.ics",
						Label:   "My Feed",
					},
				},
			},
		},
		{
			Id:    fixtures.two,
			Email: "owner.two@example.com",
			CalendarAccounts: map[string]models.CalendarAccount{
				twoGoogleKey: {
					CalendarType: models.GoogleCalendarType,
					Email:        "two.user@example.com",
					OAuth2CalendarAuth: &models.OAuth2CalendarAuth{
						AccessToken:           "two-access",
						RefreshToken:          "two-refresh",
						AccessTokenExpireDate: expiresAt,
					},
					SubCalendars: &twoSubs,
				},
				twoICSKey: {
					CalendarType: models.ICSCalendarType,
					Email:        "second@example.com",
					ICSCalendarAuth: &models.ICSCalendarAuth{
						FeedURL: "https://example.com/second.ics",
						Label:   "second@example.com",
					},
				},
			},
			PrimaryAccountKey: stringPointer(twoGoogleKey),
		},
		{
			Id:                fixtures.prefsOnly,
			Email:             "prefs.only@example.com",
			PrimaryAccountKey: stringPointer("prefs.only@example.com_google"),
			TokenOrigin:       models.IOS,
			CalendarOptions:   calendarOptions,
		},
		{
			Id:    fixtures.missing,
			Email: "ghost@example.com",
			CalendarAccounts: map[string]models.CalendarAccount{
				"ghost@example.com_google": {
					CalendarType: models.GoogleCalendarType,
					Email:        "ghost@example.com",
					OAuth2CalendarAuth: &models.OAuth2CalendarAuth{
						AccessToken:  "ghost-access",
						RefreshToken: "ghost-refresh",
					},
				},
			},
		},
	}
	for _, user := range users {
		insertCalendarUser(t, ctx, database, user)
		fixtures.userIDs = append(fixtures.userIDs, user.Id.Hex())
	}
	for _, owner := range []primitive.ObjectID{fixtures.google, fixtures.outlook, fixtures.apple, fixtures.ics, fixtures.two, fixtures.prefsOnly} {
		seedCalendarOwner(t, ctx, pool, owner.Hex())
		fixtures.externalIDs = append(fixtures.externalIDs, owner.Hex())
	}
	return fixtures
}

func TestMigrateCalendarsRehearsal(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testEncryptionKey)
	ctx, database, pool := newCalendarRehearsalContext(t, "rehearsal")
	fixtures := seedCalendarFixtures(t, ctx, database, pool)
	batch := "calendar-rehearsal"
	cleanupCalendarBatch(t, ctx, pool, batch, fixtures.externalIDs)

	migrate := &migrator{database: database, pool: pool, batch: batch, apply: true}
	summary, complete, err := migrate.migrateCalendars(ctx, configuration{batchSize: 10})
	if err != nil {
		t.Fatal(err)
	}
	if !complete {
		t.Fatal("full run must complete")
	}
	if summary.Scanned != 7 || summary.Migrated != 6 || summary.Skipped != 0 || summary.Quarantined != 1 {
		t.Fatalf("first run summary = %#v", summary)
	}

	assertCalendarTargetCounts(t, ctx, pool, batch)
	assertCalendarValues(t, ctx, pool, fixtures)
	assertCalendarQuarantineBreakdown(t, ctx, pool, batch)
	assertCalendarSourceUntouched(t, ctx, database, fixtures)

	// Replay: every completed unit is skipped and no duplicate rows are created.
	summary, complete, err = migrate.migrateCalendars(ctx, configuration{batchSize: 10})
	if err != nil {
		t.Fatal(err)
	}
	if !complete || summary.Migrated != 0 || summary.Skipped != 6 || summary.Quarantined != 1 {
		t.Fatalf("replay summary = %#v", summary)
	}
	assertCalendarTargetCounts(t, ctx, pool, batch)

	report, err := migrate.reconcile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Mismatches) != 0 {
		t.Fatalf("reconciliation mismatches: %#v\n%s", report.Mismatches, report.String())
	}
	if report.Units != 6 || report.Accounts != 6 || report.SubCalendars != 4 || report.CredentialRows != 6 || report.Preferences != 3 {
		t.Fatalf("reconciliation counts = %#v", report)
	}
	if report.Quarantined != 1 || report.QuarantineByReason[reasonMissingOwnerAccount] != 1 {
		t.Fatalf("reconciliation quarantine = %#v", report)
	}
}

func TestMigrateCalendarsResumesAfterInterruption(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testEncryptionKey)
	ctx, database, pool := newCalendarRehearsalContext(t, "resume")
	owners := []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID()}
	externalIDs := make([]string, 0, len(owners))
	for index, owner := range owners {
		insertCalendarUser(t, ctx, database, models.User{
			Id:    owner,
			Email: "resume-" + string(rune('a'+index)) + "@example.com",
			CalendarAccounts: map[string]models.CalendarAccount{
				"resume@example.com_ics": {
					CalendarType:    models.ICSCalendarType,
					Email:           "resume@example.com",
					ICSCalendarAuth: &models.ICSCalendarAuth{FeedURL: "https://example.com/resume.ics", Label: "resume@example.com"},
				},
			},
		})
		seedCalendarOwner(t, ctx, pool, owner.Hex())
		externalIDs = append(externalIDs, owner.Hex())
	}
	batch := "calendar-resume"
	cleanupCalendarBatch(t, ctx, pool, batch, externalIDs)

	migrate := &migrator{database: database, pool: pool, batch: batch, apply: true}
	summary, complete, err := migrate.migrateCalendars(ctx, configuration{batchSize: 1, limit: 1})
	if err != nil {
		t.Fatal(err)
	}
	if complete || summary.Scanned != 1 || summary.Migrated != 1 {
		t.Fatalf("interrupted run = complete %v summary %#v", complete, summary)
	}

	summary, complete, err = migrate.migrateCalendars(ctx, configuration{batchSize: 10})
	if err != nil {
		t.Fatal(err)
	}
	if !complete || summary.Scanned != 3 || summary.Migrated != 2 || summary.Skipped != 1 {
		t.Fatalf("resumed run = complete %v summary %#v", complete, summary)
	}

	report, err := migrate.reconcile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Mismatches) != 0 || report.Units != 3 || report.Accounts != 3 {
		t.Fatalf("resume reconciliation = %#v\n%s", report, report.String())
	}
}

func assertCalendarTargetCounts(t *testing.T, ctx context.Context, pool *pgxpool.Pool, batch string) {
	t.Helper()
	counts := map[string]int{
		"units":                        queryCount(t, ctx, pool, `SELECT count(*) FROM migration_ledger WHERE kind = 'calendar-account' AND batch = $1`, batch),
		"calendar_accounts":            queryCount(t, ctx, pool, `SELECT count(*) FROM calendar_accounts a JOIN platform_identities p ON p.id = a.platform_identity_id JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id WHERE l.batch = $1`, batch),
		"calendar_sub_calendars":       queryCount(t, ctx, pool, `SELECT count(*) FROM calendar_sub_calendars s JOIN calendar_accounts a ON a.id = s.calendar_account_id JOIN platform_identities p ON p.id = a.platform_identity_id JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id WHERE l.batch = $1`, batch),
		"calendar_account_credentials": queryCount(t, ctx, pool, `SELECT count(*) FROM calendar_account_credentials c JOIN calendar_accounts a ON a.id = c.calendar_account_id JOIN platform_identities p ON p.id = a.platform_identity_id JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id WHERE l.batch = $1`, batch),
		"calendar_preferences":         queryCount(t, ctx, pool, `SELECT count(*) FROM calendar_preferences pref JOIN platform_identities p ON p.id = pref.platform_identity_id JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id WHERE l.batch = $1`, batch),
	}
	want := map[string]int{
		"units": 6, "calendar_accounts": 6, "calendar_sub_calendars": 4,
		"calendar_account_credentials": 6, "calendar_preferences": 3,
	}
	for name, value := range want {
		if counts[name] != value {
			t.Fatalf("%s = %d, want %d (all: %#v)", name, counts[name], value, counts)
		}
	}
}

func assertCalendarValues(t *testing.T, ctx context.Context, pool *pgxpool.Pool, fixtures calendarFixtures) {
	t.Helper()
	repository := pgstore.NewRepository(pool)

	googleAccounts := listOwnerAccounts(t, ctx, repository, fixtures.google.Hex())
	if len(googleAccounts) != 1 {
		t.Fatalf("google accounts = %#v", googleAccounts)
	}
	google := googleAccounts[0]
	if google.CalendarKey != "google.owner@example.com_google" || google.CalendarType != "google" {
		t.Fatalf("google identity = %q %q", google.CalendarKey, google.CalendarType)
	}
	if google.Enabled == nil || !*google.Enabled {
		t.Fatalf("google explicit enabled = %#v", google.Enabled)
	}
	if google.OAuth2 == nil || google.OAuth2.AccessToken != "google-access" || google.OAuth2.RefreshToken != "google-refresh" || google.OAuth2.Scope != "calendar" {
		t.Fatalf("google oauth round-trip = %#v", google.OAuth2)
	}
	if len(google.SubCalendars) != 3 {
		t.Fatalf("google sub-calendars = %#v", google.SubCalendars)
	}
	if google.SubCalendars["personal"].Enabled == nil || *google.SubCalendars["personal"].Enabled {
		t.Fatalf("explicit false sub-calendar lost: %#v", google.SubCalendars["personal"])
	}
	if google.SubCalendars["hidden"].Enabled != nil {
		t.Fatalf("absent sub-calendar enabled must stay nil")
	}
	preferences, err := repository.GetCalendarPreferences(ctx, fixtures.google.Hex())
	if err != nil {
		t.Fatal(err)
	}
	if preferences.PrimaryAccountKey == nil || *preferences.PrimaryAccountKey != "google.owner@example.com_google" {
		t.Fatalf("google primary account key = %#v", preferences.PrimaryAccountKey)
	}
	if preferences.TokenOrigin == nil || *preferences.TokenOrigin != "web" {
		t.Fatalf("google token origin = %#v", preferences.TokenOrigin)
	}
	if !strings.Contains(string(preferences.CalendarOptions), `"bufferTime"`) {
		t.Fatalf("google calendar options = %s", preferences.CalendarOptions)
	}

	outlookAccounts := listOwnerAccounts(t, ctx, repository, fixtures.outlook.Hex())
	if len(outlookAccounts) != 1 || outlookAccounts[0].Enabled != nil {
		t.Fatalf("outlook absent enabled must stay nil: %#v", outlookAccounts)
	}
	if outlookAccounts[0].OAuth2 == nil || outlookAccounts[0].OAuth2.RefreshToken != "outlook-refresh" {
		t.Fatalf("outlook oauth round-trip = %#v", outlookAccounts[0].OAuth2)
	}

	appleAccounts := listOwnerAccounts(t, ctx, repository, fixtures.apple.Hex())
	if len(appleAccounts) != 1 || appleAccounts[0].Apple == nil || appleAccounts[0].Apple.Password != "apple-secret" {
		t.Fatalf("apple password round-trip = %#v", appleAccounts)
	}

	icsAccounts := listOwnerAccounts(t, ctx, repository, fixtures.ics.Hex())
	if len(icsAccounts) != 1 || icsAccounts[0].ICS == nil || icsAccounts[0].ICS.FeedURL != "https://example.com/feed.ics" {
		t.Fatalf("ics feed round-trip = %#v", icsAccounts)
	}
	if icsAccounts[0].Email != "My Feed" {
		t.Fatalf("ics connection identity = %q", icsAccounts[0].Email)
	}

	twoAccounts := listOwnerAccounts(t, ctx, repository, fixtures.two.Hex())
	if len(twoAccounts) != 2 {
		t.Fatalf("two-connection account = %#v", twoAccounts)
	}
	keys := map[string]bool{}
	for _, account := range twoAccounts {
		keys[account.CalendarKey] = true
	}
	if !keys["Two.User@Example.com_google"] || !keys["second@example.com_ics"] {
		t.Fatalf("legacy keys not preserved: %#v", keys)
	}

	prefsOnly, err := repository.GetCalendarPreferences(ctx, fixtures.prefsOnly.Hex())
	if err != nil {
		t.Fatal(err)
	}
	if prefsOnly.TokenOrigin == nil || *prefsOnly.TokenOrigin != "ios" {
		t.Fatalf("prefs-only token origin = %#v", prefsOnly.TokenOrigin)
	}

	var ciphertext string
	if err := pool.QueryRow(ctx, `SELECT c.apple_password_ciphertext FROM calendar_account_credentials c
JOIN calendar_accounts a ON a.id = c.calendar_account_id
JOIN platform_identities p ON p.id = a.platform_identity_id
WHERE p.external_user_id = $1`, fixtures.apple.Hex()).Scan(&ciphertext); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(ciphertext, "v1:") || ciphertext == fixtures.appleCiphertext {
		t.Fatalf("apple credential was not re-encrypted with the GCM envelope: %q", ciphertext)
	}
}

func listOwnerAccounts(t *testing.T, ctx context.Context, repository *pgstore.Repository, externalID string) []pgstore.CalendarAccount {
	t.Helper()
	accounts, err := repository.ListCalendarAccountsForUser(ctx, externalID)
	if err != nil {
		t.Fatal(err)
	}
	return accounts
}

func assertCalendarQuarantineBreakdown(t *testing.T, ctx context.Context, pool *pgxpool.Pool, batch string) {
	t.Helper()
	counts := map[string]int{}
	rows, err := pool.Query(ctx, `SELECT reason, count(*) FROM migration_quarantine WHERE batch = $1 GROUP BY reason`, batch)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reason string
		var count int
		if err := rows.Scan(&reason, &count); err != nil {
			t.Fatal(err)
		}
		counts[reason] = count
	}
	if len(counts) != 1 || counts[reasonMissingOwnerAccount] != 1 {
		t.Fatalf("quarantine breakdown = %#v", counts)
	}
}

func assertCalendarSourceUntouched(t *testing.T, ctx context.Context, database *mongo.Database, fixtures calendarFixtures) {
	t.Helper()
	if count := queryMongoCount(t, ctx, database, "users"); count != 7 {
		t.Fatalf("source users = %d, want 7", count)
	}
	var user models.User
	if err := database.Collection("users").FindOne(ctx, bson.M{"_id": fixtures.apple}).Decode(&user); err != nil {
		t.Fatalf("source user not resolvable after migration: %v", err)
	}
	stored := user.CalendarAccounts["apple.owner@example.com_apple"].AppleCalendarAuth
	if stored == nil || stored.Password != fixtures.appleCiphertext {
		t.Fatalf("retained Apple credential was modified: %#v", stored)
	}
}

func queryCount(t *testing.T, ctx context.Context, pool *pgxpool.Pool, query string, args ...any) int {
	t.Helper()
	var count int
	if err := pool.QueryRow(ctx, query, args...).Scan(&count); err != nil {
		t.Fatal(err)
	}
	return count
}

func queryMongoCount(t *testing.T, ctx context.Context, database *mongo.Database, collection string) int {
	t.Helper()
	count, err := database.Collection(collection).CountDocuments(ctx, bson.M{})
	if err != nil {
		t.Fatal(err)
	}
	return int(count)
}
