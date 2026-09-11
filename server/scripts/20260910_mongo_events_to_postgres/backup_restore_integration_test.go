package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// rehearsalRelation names one batch-scoped relation whose rows must survive a
// backup and restore byte-for-byte. The query selects each row as text so the
// digest covers every column of every migrated record, not just row counts.
type rehearsalRelation struct {
	name  string
	query string
	args  []any
}

// backupRestoreRelations returns the representative migrated relations from the
// rehearsal fixture set, scoped to this test's batch and owner identities so
// unrelated concurrent packages cannot change the reconciliation result.
func backupRestoreRelations(fixtures rehearsalFixtures, batch string) []rehearsalRelation {
	return []rehearsalRelation{
		{"postgres_events", `SELECT e::text AS row_text FROM postgres_events e JOIN migration_ledger l ON l.kind = 'event' AND l.target_id::uuid = e.id WHERE l.batch = $1`, []any{batch}},
		{"postgres_event_responses", `SELECT r::text AS row_text FROM postgres_event_responses r JOIN migration_ledger l ON l.kind = 'event' AND l.target_id::uuid = r.event_id WHERE l.batch = $1`, []any{batch}},
		{"event_signup_blocks", `SELECT b::text AS row_text FROM event_signup_blocks b JOIN migration_ledger l ON l.kind = 'event' AND l.target_id::uuid = b.event_id WHERE l.batch = $1`, []any{batch}},
		{"event_signup_responses", `SELECT s::text AS row_text FROM event_signup_responses s JOIN migration_ledger l ON l.kind = 'event' AND l.target_id::uuid = s.event_id WHERE l.batch = $1`, []any{batch}},
		{"event_attendees", `SELECT a::text AS row_text FROM event_attendees a JOIN migration_ledger l ON l.kind = 'event' AND l.target_id::uuid = a.event_id WHERE l.batch = $1`, []any{batch}},
		{"folders", `SELECT f::text AS row_text FROM folders f JOIN migration_ledger l ON l.kind = 'folder' AND l.target_id::uuid = f.id WHERE l.batch = $1`, []any{batch}},
		{"folder_events", `SELECT fe::text AS row_text FROM folder_events fe JOIN migration_ledger l ON l.kind = 'folder' AND l.target_id::uuid = fe.folder_id WHERE l.batch = $1`, []any{batch}},
		{"migration_ledger", `SELECT l::text AS row_text FROM migration_ledger l WHERE l.batch = $1`, []any{batch}},
		{"migration_quarantine", `SELECT q::text AS row_text FROM migration_quarantine q WHERE q.batch = $1`, []any{batch}},
		{"platform_identities", `SELECT p::text AS row_text FROM platform_identities p WHERE p.external_user_id = ANY($1)`, []any{fixtures.externalIDs}},
		{"accounts", `SELECT a::text AS row_text FROM accounts a JOIN platform_identities p ON p.id = a.platform_identity_id WHERE p.external_user_id = ANY($1)`, []any{fixtures.externalIDs}},
	}
}

// digestRelation returns the row count and an order-independent-of-insertion
// content digest for one batch-scoped relation.
func digestRelation(t *testing.T, ctx context.Context, conn *pgxpool.Conn, relation rehearsalRelation) (int, string) {
	t.Helper()
	query := `SELECT count(*), COALESCE(md5(string_agg(row_text, E'\n' ORDER BY row_text)), '') FROM (` + relation.query + `) rows`
	var count int
	var digest string
	if err := conn.QueryRow(ctx, query, relation.args...).Scan(&count, &digest); err != nil {
		t.Fatalf("digest %s: %v", relation.name, err)
	}
	return count, digest
}

// digestRehearsalRelations digests every relation on one connection with a
// fixed session time zone so row text renders identically in both databases.
func digestRehearsalRelations(t *testing.T, ctx context.Context, pool *pgxpool.Pool, relations []rehearsalRelation) map[string]string {
	t.Helper()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Release()
	if _, err := conn.Exec(ctx, `SET TIME ZONE 'UTC'`); err != nil {
		t.Fatal(err)
	}
	digests := make(map[string]string, len(relations))
	for _, relation := range relations {
		count, digest := digestRelation(t, ctx, conn, relation)
		digests[relation.name] = digest
		t.Logf("%s rows=%d md5=%s", relation.name, count, digest)
	}
	return digests
}

// replaceURIDatabase returns the connection URI pointed at another database.
func replaceURIDatabase(t *testing.T, rawURI string, database string) string {
	t.Helper()
	parsed, err := url.Parse(rawURI)
	if err != nil {
		t.Fatalf("parse connection URI: %v", err)
	}
	parsed.Path = "/" + database
	return parsed.String()
}

// grantBackupReadPrivileges makes the least-privilege backup role able to read
// every table, matching the grant the production bootstrap applies at creation.
func grantBackupReadPrivileges(t *testing.T, ctx context.Context, adminURI string, backupURI string) {
	t.Helper()
	parsed, err := url.Parse(backupURI)
	if err != nil {
		t.Fatalf("parse backup URI: %v", err)
	}
	username := parsed.User.Username()
	if username == "" {
		t.Fatal("backup URI has no username")
	}
	adminPool, err := pgxpool.New(ctx, adminURI)
	if err != nil {
		t.Fatal(err)
	}
	defer adminPool.Close()
	if _, err := adminPool.Exec(ctx, `GRANT pg_read_all_data TO `+pgx.Identifier{username}.Sanitize()+` WITH INHERIT TRUE`); err != nil {
		t.Fatalf("grant backup read privileges: %v", err)
	}
}

// runDumpRestoreCommand runs one PostgreSQL client tool and fails the test with
// its combined output on a nonzero exit.
func runDumpRestoreCommand(t *testing.T, ctx context.Context, name string, arguments ...string) {
	t.Helper()
	command := exec.CommandContext(ctx, name, arguments...)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("%s %s: %v\n%s", name, strings.Join(arguments, " "), err, output)
	}
}

// baseTableNames lists the public base tables of one database in name order.
func baseTableNames(t *testing.T, ctx context.Context, pool *pgxpool.Pool) []string {
	t.Helper()
	rows, err := pool.Query(ctx, `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE' ORDER BY table_name`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Fatal(err)
		}
		names = append(names, name)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	return names
}

// TestBackupRestoreRehearsal backs up the isolated PostgreSQL test database
// through the least-privilege backup role, restores it into a fresh scratch
// database through the bootstrap role, and reconciles every representative
// migrated relation by row count and content digest.
func TestBackupRestoreRehearsal(t *testing.T) {
	ctx, database, pool := newRehearsalContext(t, "backup")
	const batch = "backup-restore-itest"
	fixtures := seedRehearsalFixtures(t, ctx, database, pool)
	cleanupRehearsalBatch(t, ctx, pool, batch, fixtures.externalIDs)

	migrate := &migrator{database: database, pool: pool, batch: batch, apply: true}
	config := configuration{apply: true, batchSize: 10}
	if _, complete, err := migrate.migrateEvents(ctx, config); err != nil {
		t.Fatal(err)
	} else if !complete {
		t.Fatal("event migration did not complete")
	}
	if _, err := migrate.migrateFolders(ctx, config); err != nil {
		t.Fatal(err)
	}
	if _, err := migrate.scanOrphanResponses(ctx); err != nil {
		t.Fatal(err)
	}

	sourceURI := os.Getenv("POSTGRES_APPLICATION_URI")
	backupURI := os.Getenv("POSTGRES_BACKUP_URI")
	adminURI := os.Getenv("POSTGRES_BOOTSTRAP_URI")
	if sourceURI == "" || backupURI == "" || adminURI == "" {
		t.Skip("POSTGRES_APPLICATION_URI, POSTGRES_BACKUP_URI, and POSTGRES_BOOTSTRAP_URI are required")
	}
	for _, tool := range []string{"pg_dump", "pg_restore"} {
		if _, err := exec.LookPath(tool); err != nil {
			t.Skipf("%s is not installed in this test image", tool)
		}
	}

	grantBackupReadPrivileges(t, ctx, adminURI, backupURI)

	dumpFile := filepath.Join(t.TempDir(), "timeful-test.dump")
	runDumpRestoreCommand(t, ctx, "pg_dump", "--format=custom", "--no-owner", "--file", dumpFile, backupURI)
	dumpInfo, err := os.Stat(dumpFile)
	if err != nil {
		t.Fatal(err)
	}
	if dumpInfo.Size() == 0 {
		t.Fatal("pg_dump produced an empty archive")
	}

	scratchDatabase := fmt.Sprintf("timeful-test-restore-%d", time.Now().UnixNano()%1_000_000_000)
	adminPool, err := pgxpool.New(ctx, adminURI)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := adminPool.Exec(ctx, `CREATE DATABASE `+pgx.Identifier{scratchDatabase}.Sanitize()); err != nil {
		t.Fatalf("create scratch database: %v", err)
	}
	t.Cleanup(func() {
		if _, err := adminPool.Exec(context.Background(), `DROP DATABASE IF EXISTS `+pgx.Identifier{scratchDatabase}.Sanitize()+` WITH (FORCE)`); err != nil {
			t.Errorf("drop scratch database: %v", err)
		}
		adminPool.Close()
	})
	scratchURI := replaceURIDatabase(t, adminURI, scratchDatabase)
	runDumpRestoreCommand(t, ctx, "pg_restore", "--no-owner", "--exit-on-error", "--dbname", scratchURI, dumpFile)

	scratchPool, err := pgxpool.New(ctx, scratchURI)
	if err != nil {
		t.Fatal(err)
	}
	defer scratchPool.Close()

	sourceTables := baseTableNames(t, ctx, pool)
	restoredTables := baseTableNames(t, ctx, scratchPool)
	if strings.Join(sourceTables, ",") != strings.Join(restoredTables, ",") {
		t.Fatalf("restored schema tables differ from source: source=%v restored=%v", sourceTables, restoredTables)
	}
	t.Logf("restored schema tables=%d", len(restoredTables))

	relations := backupRestoreRelations(fixtures, batch)
	sourceDigests := digestRehearsalRelations(t, ctx, pool, relations)
	restoredDigests := digestRehearsalRelations(t, ctx, scratchPool, relations)

	mismatches := 0
	for _, relation := range relations {
		if sourceDigests[relation.name] != restoredDigests[relation.name] {
			mismatches++
			t.Errorf("%s digest mismatch: source=%s restored=%s", relation.name, sourceDigests[relation.name], restoredDigests[relation.name])
		}
	}
	if mismatches != 0 {
		t.Fatalf("backup/restore reconciliation failed: %d of %d relations differ", mismatches, len(relations))
	}
	t.Logf("backup/restore reconciliation passed: %d relations, dump %d bytes", len(relations), dumpInfo.Size())
}
