package postgres

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// calendarTestEncryptionKey is exactly 32 raw bytes, matching the ENCRYPTION_KEY
// contract the credential codec validates.
const calendarTestEncryptionKey = "0123456789abcdef0123456789abcdef"

// newCalendarTestRepository applies the migrations that define calendar
// integrations into a transaction-scoped set of temporary tables. Temp tables
// shadow the real schema so the isolated tests never mutate test-stack records.
func newCalendarTestRepository(t *testing.T) (context.Context, *Repository, pgx.Tx) {
	t.Helper()
	uri := os.Getenv("POSTGRES_APPLICATION_URI")
	if uri == "" {
		t.Skip("POSTGRES_APPLICATION_URI is required")
	}
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		t.Fatal(err)
	}
	if config.ConnConfig.Database != "timeful-test" && !strings.HasPrefix(config.ConnConfig.Database, "timeful-test-") {
		t.Fatal("requires an isolated test database")
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = tx.Rollback(ctx) })
	apply := func(name string) {
		t.Helper()
		data, err := os.ReadFile("../migrations/" + name)
		if err != nil {
			t.Fatal(err)
		}
		up := strings.Split(string(data), "-- +goose Down")[0]
		up = strings.ReplaceAll(up, "CREATE TABLE ", "CREATE TEMP TABLE ")
		if _, err := tx.Exec(ctx, up); err != nil {
			t.Fatalf("apply %s: %v", name, err)
		}
	}
	apply("20260814170000_postgres_anonymous_event_compatibility.sql")
	apply("20260815100000_postgres_event_short_id_only.sql")
	apply("20260908160000_visitor_identities.sql")
	apply("20260909090000_event_owner_authority.sql")
	apply("20260909110000_access_transfers.sql")
	apply("20260910120000_accounts.sql")
	apply("20260910130000_account_deletion.sql")
	apply("20260910140000_folders.sql")
	apply("20260911120000_drop_folder_legacy_event_id.sql")
	apply("20260910150000_signup_forms.sql")
	apply("20260910160000_availability_groups.sql")
	apply("20260910170000_migration_ledger.sql")
	apply("20260910180000_calendar_integrations.sql")
	return ctx, &Repository{db: tx}, tx
}

func seedCalendarOwner(t *testing.T, ctx context.Context, repo *Repository, externalUserID string) *PlatformIdentity {
	t.Helper()
	identity, err := repo.FindOrCreatePlatformIdentity(ctx, externalUserID)
	if err != nil {
		t.Fatal(err)
	}
	return identity
}

// TestCalendarCredentialCodecRoundTripAndTamperDetection proves the AES-256-GCM
// envelope round-trips a secret, carries the version prefix, rejects tampering,
// rejects an unsupported version, and refuses a wrong-length key.
func TestCalendarCredentialCodecRoundTripAndTamperDetection(t *testing.T) {
	codec, err := newCredentialCodec([]byte(calendarTestEncryptionKey))
	if err != nil {
		t.Fatal(err)
	}
	secret := "ya29.super-secret-refresh-token"
	envelope, err := codec.encrypt(secret)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(envelope, "v1:") {
		t.Fatalf("envelope is missing its version prefix: %q", envelope)
	}
	if strings.Contains(envelope, secret) {
		t.Fatal("envelope leaked the plaintext secret")
	}
	decrypted, err := codec.decrypt(envelope)
	if err != nil || decrypted != secret {
		t.Fatalf("round-trip = %q, %v; want %q", decrypted, err, secret)
	}

	// A second encryption must use a fresh nonce so equal secrets differ.
	again, err := codec.encrypt(secret)
	if err != nil {
		t.Fatal(err)
	}
	if again == envelope {
		t.Fatal("two encryptions reused the nonce")
	}

	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(envelope, "v1:"))
	if err != nil {
		t.Fatal(err)
	}
	raw[len(raw)-1] ^= 0xff
	tampered := "v1:" + base64.StdEncoding.EncodeToString(raw)
	if _, err := codec.decrypt(tampered); err == nil {
		t.Fatal("tampered ciphertext decrypted successfully")
	}
	if _, err := codec.decrypt("v2:" + strings.TrimPrefix(envelope, "v1:")); err == nil {
		t.Fatal("unsupported envelope version decrypted successfully")
	}
	if _, err := codec.decrypt("not-an-envelope"); err == nil {
		t.Fatal("malformed envelope decrypted successfully")
	}
	if _, err := newCredentialCodec([]byte("short")); !errors.Is(err, ErrEncryptionKeyUnavailable) {
		t.Fatalf("short key error = %v, want ErrEncryptionKeyUnavailable", err)
	}
}

// TestCalendarMigrationSchemaConstraints proves the migration enforces provider
// types, non-empty keys, one connection per owner-and-key, one credential row
// per connection, one sub-calendar per connection-and-id, and a valid token
// origin.
func TestCalendarMigrationSchemaConstraints(t *testing.T) {
	ctx, repo, tx := newCalendarTestRepository(t)
	identity := seedCalendarOwner(t, ctx, repo, "aaaaaaaaaaaaaaaaaaaaaaaa")

	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO calendar_accounts (platform_identity_id, calendar_key, calendar_type) VALUES ($1, '', 'google')`, identity.ID)
		return err
	})
	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO calendar_accounts (platform_identity_id, calendar_key, calendar_type) VALUES ($1, 'a_google', 'bogus')`, identity.ID)
		return err
	})
	var accountID string
	if err := tx.QueryRow(ctx, `INSERT INTO calendar_accounts (platform_identity_id, calendar_key, calendar_type) VALUES ($1, 'a_google', 'google') RETURNING id`, identity.ID).Scan(&accountID); err != nil {
		t.Fatal(err)
	}
	duplicate := expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO calendar_accounts (platform_identity_id, calendar_key, calendar_type) VALUES ($1, 'a_google', 'google')`, identity.ID)
		return err
	})
	if !IsUniqueViolation(duplicate) {
		t.Fatalf("duplicate calendar key error = %v, want a unique violation", duplicate)
	}

	if _, err := tx.Exec(ctx, `INSERT INTO calendar_account_credentials (calendar_account_id, oauth_access_token_ciphertext) VALUES ($1, 'v1:abc')`, accountID); err != nil {
		t.Fatal(err)
	}
	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO calendar_account_credentials (calendar_account_id, oauth_scope) VALUES ($1, 'x')`, accountID)
		return err
	})

	if _, err := tx.Exec(ctx, `INSERT INTO calendar_sub_calendars (calendar_account_id, sub_calendar_id, name) VALUES ($1, 'primary', 'Primary')`, accountID); err != nil {
		t.Fatal(err)
	}
	subDuplicate := expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO calendar_sub_calendars (calendar_account_id, sub_calendar_id, name) VALUES ($1, 'primary', 'Again')`, accountID)
		return err
	})
	if !IsUniqueViolation(subDuplicate) {
		t.Fatalf("duplicate sub-calendar error = %v, want a unique violation", subDuplicate)
	}
	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO calendar_sub_calendars (calendar_account_id, sub_calendar_id) VALUES (gen_random_uuid(), 'orphan')`)
		return err
	})

	expectSavepointError(t, ctx, tx, func() error {
		_, err := tx.Exec(ctx, `INSERT INTO calendar_preferences (platform_identity_id, token_origin) VALUES ($1, 'desktop')`, identity.ID)
		return err
	})
}

// TestCalendarAccountRepositoryEncryptsCredentialsAtRest proves every provider
// secret round-trips through the repository, is stored as a versioned GCM
// envelope rather than plaintext, and that non-secret fields stay plaintext.
func TestCalendarAccountRepositoryEncryptsCredentialsAtRest(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", calendarTestEncryptionKey)
	ctx, repo, tx := newCalendarTestRepository(t)
	seedCalendarOwner(t, ctx, repo, "bbbbbbbbbbbbbbbbbbbbbbbb")

	expiresAt := time.UnixMilli(1700000000000).UTC()
	oauth := &CalendarAccount{
		CalendarKey:  "ada@example.com_google",
		CalendarType: CalendarTypeGoogle,
		Email:        "ada@example.com",
		OAuth2: &CalendarOAuth2Credentials{
			AccessToken:          "access-token-value",
			RefreshToken:         "refresh-token-value",
			AccessTokenExpiresAt: &expiresAt,
			Scope:                "calendar.readonly",
		},
	}
	if err := repo.CreateCalendarAccount(ctx, "bbbbbbbbbbbbbbbbbbbbbbbb", oauth); err != nil {
		t.Fatal(err)
	}
	apple := &CalendarAccount{
		CalendarKey:  "ada@example.com_apple",
		CalendarType: CalendarTypeApple,
		Email:        "ada@example.com",
		Apple:        &CalendarAppleCredentials{Password: "app-specific-password"},
	}
	if err := repo.CreateCalendarAccount(ctx, "bbbbbbbbbbbbbbbbbbbbbbbb", apple); err != nil {
		t.Fatal(err)
	}
	ics := &CalendarAccount{
		CalendarKey:  "Team feed_ics",
		CalendarType: CalendarTypeICS,
		Email:        "Team feed",
		ICS:          &CalendarICSCredentials{FeedURL: "https://example.com/private/feed.ics?token=secret"},
	}
	if err := repo.CreateCalendarAccount(ctx, "bbbbbbbbbbbbbbbbbbbbbbbb", ics); err != nil {
		t.Fatal(err)
	}

	stored, err := repo.GetCalendarAccountByKey(ctx, "bbbbbbbbbbbbbbbbbbbbbbbb", oauth.CalendarKey)
	if err != nil {
		t.Fatal(err)
	}
	if stored.OAuth2 == nil || stored.OAuth2.AccessToken != "access-token-value" || stored.OAuth2.RefreshToken != "refresh-token-value" {
		t.Fatalf("oauth round-trip lost credentials: %#v", stored.OAuth2)
	}
	if stored.OAuth2.Scope != "calendar.readonly" || stored.OAuth2.AccessTokenExpiresAt == nil || stored.OAuth2.AccessTokenExpiresAt.UnixMilli() != expiresAt.UnixMilli() {
		t.Fatalf("non-secret oauth fields are unexpected: %#v", stored.OAuth2)
	}
	if stored.CalendarKey != oauth.CalendarKey || stored.CalendarType != CalendarTypeGoogle || stored.Email != "ada@example.com" {
		t.Fatalf("connection identity is unexpected: %#v", stored)
	}

	storedApple, err := repo.GetCalendarAccountByKey(ctx, "bbbbbbbbbbbbbbbbbbbbbbbb", apple.CalendarKey)
	if err != nil || storedApple.Apple == nil || storedApple.Apple.Password != "app-specific-password" {
		t.Fatalf("apple round-trip = %#v, %v", storedApple.Apple, err)
	}
	storedICS, err := repo.GetCalendarAccountByKey(ctx, "bbbbbbbbbbbbbbbbbbbbbbbb", ics.CalendarKey)
	if err != nil || storedICS.ICS == nil || storedICS.ICS.FeedURL != "https://example.com/private/feed.ics?token=secret" {
		t.Fatalf("ics round-trip = %#v, %v", storedICS.ICS, err)
	}

	var accessToken, refreshToken string
	if err := tx.QueryRow(ctx, `SELECT oauth_access_token_ciphertext, oauth_refresh_token_ciphertext FROM calendar_account_credentials WHERE calendar_account_id = $1`, oauth.ID).Scan(&accessToken, &refreshToken); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(accessToken, "v1:") || !strings.HasPrefix(refreshToken, "v1:") {
		t.Fatalf("oauth tokens are not versioned envelopes: %q, %q", accessToken, refreshToken)
	}
	if strings.Contains(accessToken, "access-token-value") || strings.Contains(refreshToken, "refresh-token-value") {
		t.Fatal("oauth token ciphertext leaked plaintext")
	}
	// The oauth connection has no Apple or ICS secrets, so those stay NULL.
	var appleEnvelope, feedEnvelope *string
	if err := tx.QueryRow(ctx, `SELECT apple_password_ciphertext, ics_feed_url_ciphertext FROM calendar_account_credentials WHERE calendar_account_id = $1`, oauth.ID).Scan(&appleEnvelope, &feedEnvelope); err != nil {
		t.Fatal(err)
	}
	if appleEnvelope != nil || feedEnvelope != nil {
		t.Fatalf("unrelated credential columns were not left absent: %v, %v", appleEnvelope, feedEnvelope)
	}
	var appleStored, feedStored string
	if err := tx.QueryRow(ctx, `SELECT apple_password_ciphertext FROM calendar_account_credentials WHERE calendar_account_id = $1`, apple.ID).Scan(&appleStored); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(appleStored, "v1:") || strings.Contains(appleStored, "app-specific-password") {
		t.Fatalf("apple password is not encrypted at rest: %q", appleStored)
	}
	if err := tx.QueryRow(ctx, `SELECT ics_feed_url_ciphertext FROM calendar_account_credentials WHERE calendar_account_id = $1`, ics.ID).Scan(&feedStored); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(feedStored, "v1:") || strings.Contains(feedStored, "token=secret") {
		t.Fatalf("ics feed url is not encrypted at rest: %q", feedStored)
	}
	var scope string
	if err := tx.QueryRow(ctx, `SELECT oauth_scope FROM calendar_account_credentials WHERE calendar_account_id = $1`, oauth.ID).Scan(&scope); err != nil || scope != "calendar.readonly" {
		t.Fatalf("scope should stay plaintext: %q, %v", scope, err)
	}
}

// TestCalendarAccountRepositorySurfacesDecryptionFailure proves a corrupted or
// unsupported credential envelope returns an error instead of an empty secret.
func TestCalendarAccountRepositorySurfacesDecryptionFailure(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", calendarTestEncryptionKey)
	ctx, repo, tx := newCalendarTestRepository(t)
	seedCalendarOwner(t, ctx, repo, "cccccccccccccccccccccccc")
	account := &CalendarAccount{
		CalendarKey:  "ada@example.com_google",
		CalendarType: CalendarTypeGoogle,
		Email:        "ada@example.com",
		OAuth2:       &CalendarOAuth2Credentials{AccessToken: "access-token-value"},
	}
	if err := repo.CreateCalendarAccount(ctx, "cccccccccccccccccccccccc", account); err != nil {
		t.Fatal(err)
	}
	var envelope string
	if err := tx.QueryRow(ctx, `SELECT oauth_access_token_ciphertext FROM calendar_account_credentials WHERE calendar_account_id = $1`, account.ID).Scan(&envelope); err != nil {
		t.Fatal(err)
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(envelope, "v1:"))
	if err != nil {
		t.Fatal(err)
	}
	raw[len(raw)-1] ^= 0xff
	if _, err := tx.Exec(ctx, `UPDATE calendar_account_credentials SET oauth_access_token_ciphertext = $2 WHERE calendar_account_id = $1`, account.ID, "v1:"+base64.StdEncoding.EncodeToString(raw)); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.GetCalendarAccountByKey(ctx, "cccccccccccccccccccccccc", account.CalendarKey); err == nil {
		t.Fatal("corrupted credential envelope read without error")
	}
	if _, err := tx.Exec(ctx, `UPDATE calendar_account_credentials SET oauth_access_token_ciphertext = 'v2:still-an-envelope' WHERE calendar_account_id = $1`, account.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.GetCalendarAccountByKey(ctx, "cccccccccccccccccccccccc", account.CalendarKey); err == nil {
		t.Fatal("unsupported credential envelope read without error")
	}
}

// TestCalendarAccountEnabledAbsentVersusFalse proves nil enabled stays absent,
// an explicit false is stored, and a later nil upsert cannot clear it.
func TestCalendarAccountEnabledAbsentVersusFalse(t *testing.T) {
	ctx, repo, _ := newCalendarTestRepository(t)
	seedCalendarOwner(t, ctx, repo, "dddddddddddddddddddddddd")
	account := &CalendarAccount{CalendarKey: "ada@example.com_google", CalendarType: CalendarTypeGoogle, Email: "ada@example.com"}
	if err := repo.UpsertCalendarAccount(ctx, "dddddddddddddddddddddddd", account); err != nil {
		t.Fatal(err)
	}
	if account.Enabled != nil {
		t.Fatalf("absent enabled was not preserved: %#v", account.Enabled)
	}
	account.Enabled = boolPointer(false)
	if err := repo.UpsertCalendarAccount(ctx, "dddddddddddddddddddddddd", account); err != nil {
		t.Fatal(err)
	}
	account.Enabled = nil
	if err := repo.UpsertCalendarAccount(ctx, "dddddddddddddddddddddddd", account); err != nil {
		t.Fatal(err)
	}
	stored, err := repo.GetCalendarAccountByKey(ctx, "dddddddddddddddddddddddd", account.CalendarKey)
	if err != nil {
		t.Fatal(err)
	}
	if stored.Enabled == nil || *stored.Enabled {
		t.Fatalf("nil upsert cleared explicit false: %#v", stored.Enabled)
	}
	if err := repo.SetCalendarAccountEnabled(ctx, "dddddddddddddddddddddddd", account.CalendarKey, true); err != nil {
		t.Fatal(err)
	}
	stored, err = repo.GetCalendarAccountByKey(ctx, "dddddddddddddddddddddddd", account.CalendarKey)
	if err != nil || stored.Enabled == nil || !*stored.Enabled {
		t.Fatalf("explicit enable not stored: %#v, %v", stored.Enabled, err)
	}
	if err := repo.SetCalendarAccountEnabled(ctx, "dddddddddddddddddddddddd", "missing_key", true); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("missing connection toggle error = %v, want pgx.ErrNoRows", err)
	}
}

// TestCalendarSubCalendarLifecycle proves add, update, remove, and enabled
// toggles preserve absent-versus-false and stay scoped to their connection.
func TestCalendarSubCalendarLifecycle(t *testing.T) {
	ctx, repo, _ := newCalendarTestRepository(t)
	seedCalendarOwner(t, ctx, repo, "eeeeeeeeeeeeeeeeeeeeeeee")
	account := &CalendarAccount{CalendarKey: "ada@example.com_google", CalendarType: CalendarTypeGoogle, Email: "ada@example.com"}
	if err := repo.CreateCalendarAccount(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account); err != nil {
		t.Fatal(err)
	}
	primary := &CalendarSubCalendar{SubCalendarID: "primary", Name: "Primary"}
	if err := repo.UpsertCalendarSubCalendar(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey, primary); err != nil {
		t.Fatal(err)
	}
	if primary.Enabled != nil {
		t.Fatalf("absent sub-calendar enabled was not preserved: %#v", primary.Enabled)
	}
	if err := repo.SetCalendarSubCalendarEnabled(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey, "primary", false); err != nil {
		t.Fatal(err)
	}
	// A provider refresh with naming but no enabled choice must not clear false.
	if err := repo.UpsertCalendarSubCalendar(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey, &CalendarSubCalendar{SubCalendarID: "primary", Name: "Renamed"}); err != nil {
		t.Fatal(err)
	}
	if err := repo.UpsertCalendarSubCalendar(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey, &CalendarSubCalendar{SubCalendarID: "work", Name: "Work"}); err != nil {
		t.Fatal(err)
	}

	account, err := repo.GetCalendarAccountByKey(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(account.SubCalendars) != 2 {
		t.Fatalf("unexpected sub-calendar set: %#v", account.SubCalendars)
	}
	renamed := account.SubCalendars["primary"]
	if renamed.Name != "Renamed" || renamed.Enabled == nil || *renamed.Enabled {
		t.Fatalf("sub-calendar update lost name or enabled: %#v", renamed)
	}
	if account.SubCalendars["work"].Enabled != nil {
		t.Fatalf("new sub-calendar should be absent enabled: %#v", account.SubCalendars["work"].Enabled)
	}
	if err := repo.RemoveCalendarSubCalendar(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey, "primary"); err != nil {
		t.Fatal(err)
	}
	account, err = repo.GetCalendarAccountByKey(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey)
	if err != nil || len(account.SubCalendars) != 1 {
		t.Fatalf("sub-calendar removal = %#v, %v", account.SubCalendars, err)
	}
	if _, ok := account.SubCalendars["primary"]; ok {
		t.Fatal("removed sub-calendar still present")
	}
	if err := repo.RemoveCalendarSubCalendar(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey, "primary"); err != nil {
		t.Fatalf("repeated removal should be a no-op: %v", err)
	}
	if err := repo.SetCalendarSubCalendarEnabled(ctx, "eeeeeeeeeeeeeeeeeeeeeeee", account.CalendarKey, "missing", true); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("missing sub-calendar toggle error = %v, want pgx.ErrNoRows", err)
	}
}

// TestCalendarPreferencesAbsentVersusPresent proves an absent preference is
// distinct from a present empty one and that the explicit write path replaces
// the complete preference state.
func TestCalendarPreferencesAbsentVersusPresent(t *testing.T) {
	ctx, repo, _ := newCalendarTestRepository(t)
	seedCalendarOwner(t, ctx, repo, "ffffffffffffffffffffffff")
	externalUserID := "ffffffffffffffffffffffff"

	if _, err := repo.GetCalendarPreferences(ctx, externalUserID); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("missing preferences error = %v, want pgx.ErrNoRows", err)
	}
	if err := repo.UpsertCalendarPreferences(ctx, externalUserID, &CalendarPreferences{}); err != nil {
		t.Fatal(err)
	}
	stored, err := repo.GetCalendarPreferences(ctx, externalUserID)
	if err != nil {
		t.Fatal(err)
	}
	if stored.PrimaryAccountKey != nil || stored.TokenOrigin != nil || stored.CalendarOptions != nil {
		t.Fatalf("absent preferences were not preserved: %#v", stored)
	}

	primary := "ada@example.com_google"
	origin := "ios"
	options := []byte(`{"bufferTime":{"enabled":true,"time":15}}`)
	if err := repo.UpsertCalendarPreferences(ctx, externalUserID, &CalendarPreferences{
		PrimaryAccountKey: &primary,
		TokenOrigin:       &origin,
		CalendarOptions:   options,
	}); err != nil {
		t.Fatal(err)
	}
	stored, err = repo.GetCalendarPreferences(ctx, externalUserID)
	if err != nil {
		t.Fatal(err)
	}
	if stored.PrimaryAccountKey == nil || *stored.PrimaryAccountKey != primary {
		t.Fatalf("primary account key not stored: %#v", stored.PrimaryAccountKey)
	}
	if stored.TokenOrigin == nil || *stored.TokenOrigin != origin {
		t.Fatalf("token origin not stored: %#v", stored.TokenOrigin)
	}
	var storedOptions, expectedOptions any
	if err := json.Unmarshal(stored.CalendarOptions, &storedOptions); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(options, &expectedOptions); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(storedOptions, expectedOptions) {
		t.Fatalf("calendar options not stored: %s", stored.CalendarOptions)
	}

	// The explicit write path clears a field the caller omits.
	if err := repo.UpsertCalendarPreferences(ctx, externalUserID, &CalendarPreferences{}); err != nil {
		t.Fatal(err)
	}
	stored, err = repo.GetCalendarPreferences(ctx, externalUserID)
	if err != nil {
		t.Fatal(err)
	}
	if stored.PrimaryAccountKey != nil || stored.TokenOrigin != nil || stored.CalendarOptions != nil {
		t.Fatalf("explicit empty preferences did not clear state: %#v", stored)
	}
	if err := repo.DeleteCalendarPreferences(ctx, externalUserID); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.GetCalendarPreferences(ctx, externalUserID); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("deleted preferences error = %v, want pgx.ErrNoRows", err)
	}
}

// TestCalendarRepositoryRequiresExistingOwner proves a calendar record never
// creates an account: an owner with no platform identity is an error.
func TestCalendarRepositoryRequiresExistingOwner(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", calendarTestEncryptionKey)
	ctx, repo, _ := newCalendarTestRepository(t)
	account := &CalendarAccount{CalendarKey: "ghost@example.com_google", CalendarType: CalendarTypeGoogle, Email: "ghost@example.com"}
	if err := repo.UpsertCalendarAccount(ctx, "999999999999999999999999", account); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("missing owner error = %v, want pgx.ErrNoRows", err)
	}
	if _, err := repo.ListCalendarAccountsForUser(ctx, "999999999999999999999999"); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("missing owner list error = %v, want pgx.ErrNoRows", err)
	}
	if err := repo.CreateCalendarAccount(ctx, "999999999999999999999999", &CalendarAccount{
		CalendarKey:  "ghost@example.com_google",
		CalendarType: CalendarTypeGoogle,
		OAuth2:       &CalendarOAuth2Credentials{AccessToken: "token"},
	}); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("missing owner create error = %v, want pgx.ErrNoRows", err)
	}
}

// TestCalendarAccountRepositoryListAndDelete proves connections are owned and
// listed independently and that deletion cascades credentials and sub-calendars
// while leaving another connection untouched.
func TestCalendarAccountRepositoryListAndDelete(t *testing.T) {
	ctx, repo, tx := newCalendarTestRepository(t)
	seedCalendarOwner(t, ctx, repo, "121212121212121212121212")
	first := &CalendarAccount{CalendarKey: "ada@example.com_google", CalendarType: CalendarTypeGoogle, Email: "ada@example.com"}
	second := &CalendarAccount{CalendarKey: "ada@example.com_outlook", CalendarType: CalendarTypeOutlook, Email: "ada@example.com"}
	if err := repo.CreateCalendarAccount(ctx, "121212121212121212121212", first); err != nil {
		t.Fatal(err)
	}
	if err := repo.CreateCalendarAccount(ctx, "121212121212121212121212", second); err != nil {
		t.Fatal(err)
	}
	if err := repo.UpsertCalendarSubCalendar(ctx, "121212121212121212121212", first.CalendarKey, &CalendarSubCalendar{SubCalendarID: "primary"}); err != nil {
		t.Fatal(err)
	}

	accounts, err := repo.ListCalendarAccountsForUser(ctx, "121212121212121212121212")
	if err != nil {
		t.Fatal(err)
	}
	if len(accounts) != 2 {
		t.Fatalf("unexpected connection count: %d", len(accounts))
	}
	if len(accounts[0].SubCalendars) != 1 {
		t.Fatalf("listed connection lost sub-calendars: %#v", accounts[0].SubCalendars)
	}

	if err := repo.DeleteCalendarAccount(ctx, "121212121212121212121212", first.CalendarKey); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.GetCalendarAccountByKey(ctx, "121212121212121212121212", first.CalendarKey); !errors.Is(err, pgx.ErrNoRows) {
		t.Fatalf("deleted connection lookup error = %v, want pgx.ErrNoRows", err)
	}
	var subs, credentials int
	if err := tx.QueryRow(ctx, `SELECT (SELECT count(*) FROM calendar_sub_calendars WHERE calendar_account_id = $1), (SELECT count(*) FROM calendar_account_credentials WHERE calendar_account_id = $1)`, first.ID).Scan(&subs, &credentials); err != nil {
		t.Fatal(err)
	}
	if subs != 0 || credentials != 0 {
		t.Fatalf("deletion did not cascade: sub-calendars=%d credentials=%d", subs, credentials)
	}
	remaining, err := repo.ListCalendarAccountsForUser(ctx, "121212121212121212121212")
	if err != nil || len(remaining) != 1 || remaining[0].CalendarKey != second.CalendarKey {
		t.Fatalf("unexpected remaining connections: %#v, %v", remaining, err)
	}
	// Deleting a missing connection is a no-op.
	if err := repo.DeleteCalendarAccount(ctx, "121212121212121212121212", "missing_key"); err != nil {
		t.Fatalf("repeated deletion should be a no-op: %v", err)
	}
}
