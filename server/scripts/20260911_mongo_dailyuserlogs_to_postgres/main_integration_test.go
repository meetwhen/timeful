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

	pgstore "timeful/server/postgres"
)

// newDailyUserLogRehearsalContext connects to the isolated Mongo and PostgreSQL
// test databases and drops a per-test Mongo source database.
func newDailyUserLogRehearsalContext(t *testing.T, suffix string) (context.Context, *mongo.Database, *pgxpool.Pool) {
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

	database := mongoClient.Database(baseDatabase + "_dailyuserlogs_backfill_" + suffix)
	if err := database.Drop(ctx); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = database.Drop(ctx) })
	return ctx, database, pool
}

// seedDailyLogAccount creates the authoritative PostgreSQL account a retained
// daily-log membership resolves through.
func seedDailyLogAccount(t *testing.T, ctx context.Context, pool *pgxpool.Pool, externalID string) {
	t.Helper()
	if _, err := pgstore.NewRepository(pool).FindOrCreateAccount(ctx, externalID, pgstore.Account{Email: externalID + "@example.com", FirstName: "Member"}); err != nil {
		t.Fatal(err)
	}
}

// seedDailyLogIdentityOnly creates a platform identity without an account row
// so the rehearsal covers an account that cannot be resolved.
func seedDailyLogIdentityOnly(t *testing.T, ctx context.Context, pool *pgxpool.Pool, externalID string) {
	t.Helper()
	if _, err := pool.Exec(ctx, `INSERT INTO platform_identities (external_user_id) VALUES ($1)
ON CONFLICT (external_user_id) DO NOTHING`, externalID); err != nil {
		t.Fatal(err)
	}
}

func insertLegacyDailyUserLog(t *testing.T, ctx context.Context, database *mongo.Database, log legacyDailyUserLog) {
	t.Helper()
	if _, err := database.Collection("dailyuserlogs").InsertOne(ctx, log); err != nil {
		t.Fatal(err)
	}
}

func cleanupDailyUserLogBatch(t *testing.T, ctx context.Context, pool *pgxpool.Pool, batch string, from, to time.Time, externalIDs []string) {
	t.Helper()
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		if _, err := pool.Exec(cleanupCtx, `DELETE FROM daily_user_logs WHERE log_date BETWEEN $1::date AND $2::date`,
			from.Format("2006-01-02"), to.Format("2006-01-02")); err != nil {
			t.Errorf("cleanup daily logs: %v", err)
		}
		if _, err := pool.Exec(cleanupCtx, `DELETE FROM migration_quarantine WHERE batch = $1`, batch); err != nil {
			t.Errorf("cleanup quarantine: %v", err)
		}
		if _, err := pool.Exec(cleanupCtx, `DELETE FROM migration_ledger WHERE batch = $1`, batch); err != nil {
			t.Errorf("cleanup ledger: %v", err)
		}
		if len(externalIDs) > 0 {
			if _, err := pool.Exec(cleanupCtx, `DELETE FROM accounts WHERE platform_identity_id IN (SELECT id FROM platform_identities WHERE external_user_id = ANY($1))`, externalIDs); err != nil {
				t.Errorf("cleanup accounts: %v", err)
			}
			if _, err := pool.Exec(cleanupCtx, `DELETE FROM platform_identities WHERE external_user_id = ANY($1)`, externalIDs); err != nil {
				t.Errorf("cleanup identities: %v", err)
			}
		}
	})
}

// TestMigrateDailyUserLogsRehearsal covers the representative daily-log fixture
// set: overlapping membership across days, an empty day, two retained documents
// sharing a date, a missing-owner quarantine, replay idempotency, source
// immutability, and clean reconciliation.
func TestMigrateDailyUserLogsRehearsal(t *testing.T) {
	ctx, database, pool := newDailyUserLogRehearsalContext(t, "rehearsal")
	batch := "daily-user-logs-rehearsal"

	memberA := primitive.NewObjectID()
	memberB := primitive.NewObjectID()
	memberC := primitive.NewObjectID()
	memberD := primitive.NewObjectID()
	missingNoIdentity := primitive.NewObjectID()
	missingIdentityOnly := primitive.NewObjectID()
	externalIDs := []string{memberA.Hex(), memberB.Hex(), memberC.Hex(), memberD.Hex(), missingIdentityOnly.Hex()}

	dayOne := time.Date(2026, 8, 10, 0, 0, 0, 0, time.UTC)
	dayTwo := time.Date(2026, 8, 11, 0, 0, 0, 0, time.UTC)
	dayFour := time.Date(2026, 8, 13, 0, 0, 0, 0, time.UTC)
	dayFive := time.Date(2026, 8, 14, 0, 0, 0, 0, time.UTC)

	for _, member := range []primitive.ObjectID{memberA, memberB, memberC, memberD} {
		seedDailyLogAccount(t, ctx, pool, member.Hex())
	}
	seedDailyLogIdentityOnly(t, ctx, pool, missingIdentityOnly.Hex())

	logs := []legacyDailyUserLog{
		{ID: primitive.NewObjectID(), Date: primitive.NewDateTimeFromTime(dayOne), UserIDs: []primitive.ObjectID{memberA, memberB}},
		{ID: primitive.NewObjectID(), Date: primitive.NewDateTimeFromTime(dayTwo), UserIDs: []primitive.ObjectID{memberB, memberC}},
		{ID: primitive.NewObjectID(), Date: primitive.NewDateTimeFromTime(dayFour), UserIDs: []primitive.ObjectID{memberC, memberA}},
		{ID: primitive.NewObjectID(), Date: primitive.NewDateTimeFromTime(dayFour), UserIDs: []primitive.ObjectID{memberA, memberD}},
		{ID: primitive.NewObjectID(), Date: primitive.NewDateTimeFromTime(dayFive), UserIDs: []primitive.ObjectID{memberD, missingNoIdentity, missingIdentityOnly}},
	}
	for _, log := range logs {
		insertLegacyDailyUserLog(t, ctx, database, log)
	}
	cleanupDailyUserLogBatch(t, ctx, pool, batch, dayOne, dayFive, externalIDs)

	migrate := &migrator{database: database, pool: pool, batch: batch, apply: true}
	summary, complete, err := migrate.migrateDailyUserLogs(ctx, configuration{batchSize: 2})
	if err != nil {
		t.Fatal(err)
	}
	if !complete {
		t.Fatal("full run must complete")
	}
	if summary.Scanned != 5 || summary.Migrated != 4 || summary.Skipped != 0 || summary.Quarantined != 1 {
		t.Fatalf("first run summary = %#v", summary)
	}

	assertDailyLogMembers(t, ctx, pool, dayOne, []string{memberA.Hex(), memberB.Hex()})
	assertDailyLogMembers(t, ctx, pool, dayTwo, []string{memberB.Hex(), memberC.Hex()})
	assertDailyLogMembers(t, ctx, pool, dayFour, []string{memberC.Hex(), memberA.Hex(), memberD.Hex()})
	if got := countPostgres(t, ctx, pool, `SELECT count(*) FROM daily_user_logs WHERE log_date = $1::date`, dayTwo.AddDate(0, 0, 1).Format("2006-01-02")); got != 0 {
		t.Fatalf("empty day between logs has %d target rows, want 0", got)
	}
	if got := countPostgres(t, ctx, pool, `SELECT count(*) FROM daily_user_logs WHERE log_date = $1::date`, dayFive.Format("2006-01-02")); got != 0 {
		t.Fatalf("quarantined day has %d target rows, want 0", got)
	}
	if got := countPostgres(t, ctx, pool, `SELECT count(DISTINCT target_id) FROM migration_ledger WHERE kind = $1 AND batch = $2`, ledgerKindDailyUserLog, batch); got != 3 {
		t.Fatalf("distinct target logs = %d, want 3 (overlapping dates share one identity)", got)
	}
	if got := countPostgres(t, ctx, pool, `SELECT count(*) FROM migration_ledger WHERE kind = $1 AND batch = $2`, ledgerKindDailyUserLog, batch); got != 4 {
		t.Fatalf("ledger units = %d, want 4", got)
	}

	var reason string
	if err := pool.QueryRow(ctx, `SELECT reason FROM migration_quarantine WHERE kind = $1 AND legacy_id = $2`,
		ledgerKindDailyUserLog, logs[4].ID.Hex()).Scan(&reason); err != nil {
		t.Fatal(err)
	}
	if reason != reasonMissingOwnerAccount {
		t.Fatalf("quarantine reason = %q, want %q", reason, reasonMissingOwnerAccount)
	}

	// The retained MongoDB source is never modified.
	if got := countMongo(t, ctx, database, "dailyuserlogs"); got != 5 {
		t.Fatalf("source dailyuserlogs = %d, want 5", got)
	}
	var stored legacyDailyUserLog
	if err := database.Collection("dailyuserlogs").FindOne(ctx, bson.M{"_id": logs[4].ID}).Decode(&stored); err != nil {
		t.Fatal(err)
	}
	if !stored.Date.Time().Equal(dayFive) || len(stored.UserIDs) != 3 {
		t.Fatalf("retained source was modified: %#v", stored)
	}

	report, err := migrate.reconcile(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Mismatches) != 0 {
		t.Fatalf("reconciliation mismatches: %#v\n%s", report.Mismatches, report.String())
	}
	if report.Units != 4 || report.Logs != 3 || report.Memberships != 7 || report.SourceLogs != 4 || report.SourceMemberships != 7 {
		t.Fatalf("reconciliation counts = %#v", report)
	}
	if report.Quarantined != 1 || len(report.QuarantineByReason) != 1 || report.QuarantineByReason[reasonMissingOwnerAccount] != 1 {
		t.Fatalf("reconciliation quarantine = %#v", report)
	}

	// Replay: every migrated unit is skipped, the quarantine decision is
	// re-observed, and no duplicate rows are created.
	summary, complete, err = migrate.migrateDailyUserLogs(ctx, configuration{batchSize: 2})
	if err != nil {
		t.Fatal(err)
	}
	if !complete || summary.Migrated != 0 || summary.Skipped != 4 || summary.Quarantined != 1 {
		t.Fatalf("replay summary = %#v", summary)
	}
	assertDailyLogMembers(t, ctx, pool, dayFour, []string{memberC.Hex(), memberA.Hex(), memberD.Hex()})
	if got := countPostgres(t, ctx, pool, `SELECT count(*) FROM daily_user_log_members m
JOIN daily_user_logs l ON l.id = m.daily_user_log_id
WHERE l.log_date BETWEEN $1::date AND $2::date`, dayOne.Format("2006-01-02"), dayFive.Format("2006-01-02")); got != 7 {
		t.Fatalf("memberships after replay = %d, want 7", got)
	}
}

// TestMigrateDailyUserLogsResumesAfterInterruption proves a --limit stop leaves
// completed units committed and the next run resumes from the first incomplete
// unit without duplicating membership.
func TestMigrateDailyUserLogsResumesAfterInterruption(t *testing.T) {
	ctx, database, pool := newDailyUserLogRehearsalContext(t, "resume")
	batch := "daily-user-logs-resume"

	members := []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID(), primitive.NewObjectID()}
	days := []time.Time{
		time.Date(2026, 7, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2026, 7, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2026, 7, 7, 0, 0, 0, 0, time.UTC),
	}
	externalIDs := make([]string, 0, len(members))
	for index, member := range members {
		seedDailyLogAccount(t, ctx, pool, member.Hex())
		insertLegacyDailyUserLog(t, ctx, database, legacyDailyUserLog{
			ID:      primitive.NewObjectID(),
			Date:    primitive.NewDateTimeFromTime(days[index]),
			UserIDs: []primitive.ObjectID{member},
		})
		externalIDs = append(externalIDs, member.Hex())
	}
	cleanupDailyUserLogBatch(t, ctx, pool, batch, days[0], days[len(days)-1], externalIDs)

	migrate := &migrator{database: database, pool: pool, batch: batch, apply: true}
	summary, complete, err := migrate.migrateDailyUserLogs(ctx, configuration{batchSize: 1, limit: 1})
	if err != nil {
		t.Fatal(err)
	}
	if complete || summary.Scanned != 1 || summary.Migrated != 1 {
		t.Fatalf("interrupted run = complete %v summary %#v", complete, summary)
	}

	summary, complete, err = migrate.migrateDailyUserLogs(ctx, configuration{batchSize: 10})
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
	if len(report.Mismatches) != 0 || report.Units != 3 || report.Logs != 3 || report.Memberships != 3 {
		t.Fatalf("resume reconciliation = %#v\n%s", report, report.String())
	}
	if got := countPostgres(t, ctx, pool, `SELECT count(*) FROM daily_user_log_members m
JOIN daily_user_logs l ON l.id = m.daily_user_log_id
WHERE l.log_date BETWEEN $1::date AND $2::date`, days[0].Format("2006-01-02"), days[2].Format("2006-01-02")); got != 3 {
		t.Fatalf("memberships after resume = %d, want 3", got)
	}
}

func assertDailyLogMembers(t *testing.T, ctx context.Context, pool *pgxpool.Pool, date time.Time, want []string) {
	t.Helper()
	rows, err := pool.Query(ctx, `SELECT m.account_user_id FROM daily_user_log_members m
JOIN daily_user_logs l ON l.id = m.daily_user_log_id
WHERE l.log_date = $1::date
ORDER BY m.first_seen_position, m.id`, date.Format("2006-01-02"))
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	got := []string{}
	for rows.Next() {
		var accountUserID string
		if err := rows.Scan(&accountUserID); err != nil {
			t.Fatal(err)
		}
		got = append(got, accountUserID)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	if !equalOrdered(got, want) {
		t.Fatalf("daily log %s members = %v, want %v", date.Format("2006-01-02"), got, want)
	}
}

func countPostgres(t *testing.T, ctx context.Context, pool *pgxpool.Pool, query string, args ...any) int {
	t.Helper()
	var count int
	if err := pool.QueryRow(ctx, query, args...).Scan(&count); err != nil {
		t.Fatal(err)
	}
	return count
}

func countMongo(t *testing.T, ctx context.Context, database *mongo.Database, collection string) int {
	t.Helper()
	count, err := database.Collection(collection).CountDocuments(ctx, bson.M{})
	if err != nil {
		t.Fatal(err)
	}
	return int(count)
}
