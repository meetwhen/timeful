package postgres

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"timeful/server/models"
)

const (
	baselineMigrationName       = "20260912000000_baseline_schema.sql"
	consolidationMigrationName  = "20260912120000_account_identity_platform_uuid.sql"
	payloadRewriteMigrationName = "20260912130000_rewrite_legacy_account_payloads.sql"
)

func newMigrationFileTransaction(t *testing.T) (context.Context, pgx.Tx) {
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
	return ctx, tx
}

// applyMigrationUpFile applies one goose migration's up section into the
// transaction. CREATE TABLE becomes CREATE TEMP TABLE so the real schema is
// shadowed, matching newMigrationTestRepository.
func applyMigrationUpFile(t *testing.T, tx pgx.Tx, name string) error {
	t.Helper()
	data, err := os.ReadFile("../migrations/" + name)
	if err != nil {
		t.Fatal(err)
	}
	up := strings.Split(string(data), "-- +goose Down")[0]
	up = strings.ReplaceAll(up, "CREATE TABLE ", "CREATE TEMP TABLE ")
	_, err = tx.Exec(context.Background(), up)
	return err
}

func hasColumn(t *testing.T, tx pgx.Tx, table, column string) bool {
	t.Helper()
	var exists bool
	if err := tx.QueryRow(context.Background(), `SELECT EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = to_regclass($1) AND attname = $2 AND attnum > 0 AND NOT attisdropped
    )`, table, column).Scan(&exists); err != nil {
		t.Fatal(err)
	}
	return exists
}

// TestAccountIdentityMigrationBackfillsLegacyReferences rehearses the
// consolidation over representative legacy rows: every account reference is
// remapped to the platform identity uuid, the legacy columns are dropped, an
// unmappable tombstone is cleared with its unreachable legacy value, and a
// repeated run is a no-op.
func TestAccountIdentityMigrationBackfillsLegacyReferences(t *testing.T) {
	ctx, tx := newMigrationFileTransaction(t)
	if err := applyMigrationUpFile(t, tx, baselineMigrationName); err != nil {
		t.Fatalf("apply baseline: %v", err)
	}

	legacyAccountID := "507f1f77bcf86cd799439011"
	// A payload created before the cutover embeds the retired 24-hex identifier
	// and the legacy zero sentinel, which the strict canonical UUID decoder
	// rejects. The migration must rewrite these payloads to stay readable.
	legacyEventPayload := `{"_id":"000000000000000000000000","ownerId":"000000000000000000000000","name":"Legacy","type":"specific_dates",` +
		`"responses":{"507f1f77bcf86cd799439011":{"name":"Ada","userId":"507f1f77bcf86cd799439011"}},` +
		`"signUpResponses":{"507f1f77bcf86cd799439011":{"name":"Ada","userId":"507f1f77bcf86cd799439011"}}}`
	legacyResponsePayload := `{"name":"Legacy Account","userId":"507f1f77bcf86cd799439011",` +
		`"user":{"_id":"507f1f77bcf86cd799439011","firstName":"Legacy"}}`
	var legacyEvent models.Event
	if err := json.Unmarshal([]byte(legacyEventPayload), &legacyEvent); err == nil {
		t.Fatal("legacy event payload fixture unexpectedly decodes as canonical UUIDs")
	}
	var legacyResponse models.Response
	if err := json.Unmarshal([]byte(legacyResponsePayload), &legacyResponse); err == nil {
		t.Fatal("legacy response payload fixture unexpectedly decodes as canonical UUIDs")
	}
	var platformIdentityID string
	if err := tx.QueryRow(ctx, `INSERT INTO platform_identities (external_user_id) VALUES ($1) RETURNING id`, legacyAccountID).Scan(&platformIdentityID); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO accounts (platform_identity_id, email) VALUES ($1, 'legacy@example.com')`, platformIdentityID); err != nil {
		t.Fatal(err)
	}
	var eventID string
	if err := tx.QueryRow(ctx, `INSERT INTO postgres_events (short_id, name, type, owner_external_id, payload)
 VALUES ($1, 'Legacy', 'specific_dates', $2, $3::jsonb) RETURNING id`, signupTestShortID(t), legacyAccountID, legacyEventPayload).Scan(&eventID); err != nil {
		t.Fatal(err)
	}
	var ownerVisitorID, guestVisitorID string
	if err := tx.QueryRow(ctx, `INSERT INTO event_visitor_identities (event_id, platform_identity_id) VALUES ($1, $2) RETURNING id`, eventID, platformIdentityID).Scan(&ownerVisitorID); err != nil {
		t.Fatal(err)
	}
	if err := tx.QueryRow(ctx, `INSERT INTO event_visitor_identities (event_id) VALUES ($1) RETURNING id`, eventID).Scan(&guestVisitorID); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO postgres_event_responses (event_id, event_visitor_identity_id, respondent_kind, account_user_id, payload)
 VALUES ($1, $2, 'account', $3, $4::jsonb)`, eventID, ownerVisitorID, legacyAccountID, legacyResponsePayload); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO event_signup_responses (event_id, event_visitor_identity_id, respondent_kind, account_user_id, name)
 VALUES ($1, $2, 'account', $3, 'Legacy Signup')`, eventID, ownerVisitorID, legacyAccountID); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO event_attendees (event_id, email, account_user_id) VALUES ($1, 'legacy@example.com', $2)`, eventID, legacyAccountID); err != nil {
		t.Fatal(err)
	}
	var folderID string
	if err := tx.QueryRow(ctx, `INSERT INTO folders (account_user_id, name) VALUES ($1, 'Legacy Folder') RETURNING id`, legacyAccountID).Scan(&folderID); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO folder_events (account_user_id, folder_id, event_id) VALUES ($1, $2, $3)`, legacyAccountID, folderID, eventID); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO access_transfers (event_id, source_hash, external_user_id)
 VALUES ($1, decode(repeat('00', 32), 'hex'), $2)`, eventID, legacyAccountID); err != nil {
		t.Fatal(err)
	}
	var logID string
	if err := tx.QueryRow(ctx, `INSERT INTO daily_user_logs (log_date) VALUES ('2000-01-01') RETURNING id`).Scan(&logID); err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO daily_user_log_members (daily_user_log_id, account_user_id, first_seen_position)
 VALUES ($1, $2, 0)`, logID, legacyAccountID); err != nil {
		t.Fatal(err)
	}
	// The deletion transaction removes the identity before it writes a tombstone,
	// so an existing tombstone can never map to a live platform identity.
	if _, err := tx.Exec(ctx, `INSERT INTO account_deletion_tombstones (external_user_id) VALUES ('ffffffffffffffffffffffff')`); err != nil {
		t.Fatal(err)
	}

	if err := applyMigrationUpFile(t, tx, consolidationMigrationName); err != nil {
		t.Fatalf("apply consolidation: %v", err)
	}
	// A repeated application must be a no-op, not an error.
	if err := applyMigrationUpFile(t, tx, consolidationMigrationName); err != nil {
		t.Fatalf("repeat consolidation: %v", err)
	}
	// The payload rewrite ships as its own version because the consolidation
	// version had already been applied where the rewrite was added later.
	if err := applyMigrationUpFile(t, tx, payloadRewriteMigrationName); err != nil {
		t.Fatalf("apply payload rewrite: %v", err)
	}
	if err := applyMigrationUpFile(t, tx, payloadRewriteMigrationName); err != nil {
		t.Fatalf("repeat payload rewrite: %v", err)
	}

	var storedPlatformIdentityID string
	if err := tx.QueryRow(ctx, `SELECT owner_platform_identity_id FROM postgres_events WHERE id = $1`, eventID).Scan(&storedPlatformIdentityID); err != nil {
		t.Fatal(err)
	}
	if storedPlatformIdentityID != platformIdentityID {
		t.Fatalf("event owner = %q, want %q", storedPlatformIdentityID, platformIdentityID)
	}
	for _, check := range []struct {
		table string
		query string
		args  []any
	}{
		{"postgres_event_responses", `SELECT platform_identity_id FROM postgres_event_responses WHERE event_id = $1`, []any{eventID}},
		{"event_signup_responses", `SELECT platform_identity_id FROM event_signup_responses WHERE event_id = $1`, []any{eventID}},
		{"event_attendees", `SELECT platform_identity_id FROM event_attendees WHERE event_id = $1`, []any{eventID}},
		{"folders", `SELECT platform_identity_id FROM folders WHERE id = $1`, []any{folderID}},
		{"folder_events", `SELECT platform_identity_id FROM folder_events WHERE folder_id = $1`, []any{folderID}},
		{"access_transfers", `SELECT platform_identity_id FROM access_transfers WHERE event_id = $1`, []any{eventID}},
		{"daily_user_log_members", `SELECT platform_identity_id FROM daily_user_log_members WHERE daily_user_log_id = $1`, []any{logID}},
	} {
		var remapped string
		if err := tx.QueryRow(ctx, check.query, check.args...).Scan(&remapped); err != nil {
			t.Fatalf("%s backfill: %v", check.table, err)
		}
		if remapped != platformIdentityID {
			t.Fatalf("%s reference = %q, want %q", check.table, remapped, platformIdentityID)
		}
	}

	var tombstones int
	if err := tx.QueryRow(ctx, `SELECT count(*) FROM account_deletion_tombstones`).Scan(&tombstones); err != nil {
		t.Fatal(err)
	}
	if tombstones != 0 {
		t.Fatalf("unreachable legacy tombstones were not cleared: %d", tombstones)
	}

	// Stored payloads written before the cutover must decode under the strict
	// canonical UUID decoder after the payload rewrite removed the duplicated
	// account-identity keys.
	var eventPayload []byte
	if err := tx.QueryRow(ctx, `SELECT payload FROM postgres_events WHERE id = $1`, eventID).Scan(&eventPayload); err != nil {
		t.Fatal(err)
	}
	assertPayloadLacksKeys(t, eventPayload, "_id", "ownerId", "signUpResponses", "responses")
	var migratedEvent models.Event
	if err := json.Unmarshal(eventPayload, &migratedEvent); err != nil {
		t.Fatalf("migrated event payload did not decode: %v", err)
	}
	if !migratedEvent.Id.IsZero() || !migratedEvent.OwnerId.IsZero() {
		t.Fatal("migrated event payload still carries a stored identity")
	}

	var responsePayload []byte
	if err := tx.QueryRow(ctx, `SELECT payload FROM postgres_event_responses WHERE event_id = $1`, eventID).Scan(&responsePayload); err != nil {
		t.Fatal(err)
	}
	assertPayloadLacksKeys(t, responsePayload, "userId", "user")
	var migratedResponse models.Response
	if err := json.Unmarshal(responsePayload, &migratedResponse); err != nil {
		t.Fatalf("migrated response payload did not decode: %v", err)
	}
	if !migratedResponse.UserId.IsZero() {
		t.Fatal("migrated response payload still carries a stored identity")
	}

	for _, columnCheck := range []struct{ table, column string }{
		{"platform_identities", "external_user_id"},
		{"postgres_events", "owner_external_id"},
		{"postgres_event_responses", "account_user_id"},
		{"event_signup_responses", "account_user_id"},
		{"event_attendees", "account_user_id"},
		{"folders", "account_user_id"},
		{"folder_events", "account_user_id"},
		{"access_transfers", "external_user_id"},
		{"daily_user_log_members", "account_user_id"},
		{"account_deletion_tombstones", "external_user_id"},
	} {
		if hasColumn(t, tx, columnCheck.table, columnCheck.column) {
			t.Fatalf("%s.%s was not dropped", columnCheck.table, columnCheck.column)
		}
	}
}

// TestAccountIdentityMigrationRejectsUnmappedReferences proves the guarded
// backfill refuses to drop the legacy mapping when a row references an external
// account identifier that has no platform identity.
func TestAccountIdentityMigrationRejectsUnmappedReferences(t *testing.T) {
	ctx, tx := newMigrationFileTransaction(t)
	if err := applyMigrationUpFile(t, tx, baselineMigrationName); err != nil {
		t.Fatalf("apply baseline: %v", err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO postgres_events (short_id, name, type, owner_external_id)
 VALUES ($1, 'Orphan', 'specific_dates', 'ffffffffffffffffffffffff')`, signupTestShortID(t)); err != nil {
		t.Fatal(err)
	}
	err := applyMigrationUpFile(t, tx, consolidationMigrationName)
	if err == nil {
		t.Fatal("expected an unmapped owner external id to abort the migration")
	}
	if !strings.Contains(err.Error(), "unmapped") {
		t.Fatalf("migration error = %v, want an unmapped-reference failure", err)
	}
}

// TestAccountIdentityMigrationDownRefuses proves the consolidation cannot be
// reversed destructively.
func TestAccountIdentityMigrationDownRefuses(t *testing.T) {
	ctx, tx := newMigrationFileTransaction(t)
	if err := applyMigrationUpFile(t, tx, baselineMigrationName); err != nil {
		t.Fatalf("apply baseline: %v", err)
	}
	if err := applyMigrationUpFile(t, tx, consolidationMigrationName); err != nil {
		t.Fatalf("apply consolidation: %v", err)
	}
	data, err := os.ReadFile("../migrations/" + consolidationMigrationName)
	if err != nil {
		t.Fatal(err)
	}
	down := strings.SplitN(string(data), "-- +goose Down", 2)
	if len(down) != 2 {
		t.Fatal("consolidation migration has no down section")
	}
	if _, err := tx.Exec(ctx, down[1]); err == nil {
		t.Fatal("expected the down migration to refuse")
	} else if !strings.Contains(err.Error(), "cannot be reversed") {
		t.Fatalf("down error = %v, want a refusal", err)
	}
}

// assertPayloadLacksKeys proves the consolidation removed the stored
// account-identity keys from a JSONB payload.
func assertPayloadLacksKeys(t *testing.T, payload []byte, keys ...string) {
	t.Helper()
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("decode payload: %v", err)
	}
	for _, key := range keys {
		if _, present := decoded[key]; present {
			t.Fatalf("payload still contains %q after consolidation", key)
		}
	}
}
