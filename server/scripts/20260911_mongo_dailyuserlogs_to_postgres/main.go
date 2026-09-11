// mongo_dailyuserlogs_to_postgres backfills authoritative PostgreSQL historical
// daily user logs from retained MongoDB dailyuserlogs documents. It never
// modifies the MongoDB source. Every migration unit is one retained log with its
// membership list and is written in one PostgreSQL transaction that also records
// completion in the migration ledger, so a rerun resumes after the first
// incomplete unit and repeated runs never create duplicate logs or memberships.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const usage = "usage: go run ./scripts/20260911_mongo_dailyuserlogs_to_postgres [--apply] [--batch-size N] [--limit N] [--batch-label label]"

type configuration struct {
	apply       bool
	batchSize   int64
	limit       int
	batchLabel  string
	mongoURI    string
	mongoDB     string
	postgresURI string
}

// migrationSummary reports the work observed in one run. Migrated counts units
// whose writes committed in this run; Skipped counts units already completed by
// an earlier run; Quarantined counts retained logs whose membership could not
// resolve to PostgreSQL accounts.
type migrationSummary struct {
	Scanned     int
	Migrated    int
	Skipped     int
	Quarantined int
}

// migrator carries the two stores and the run identity for one backfill.
type migrator struct {
	database *mongo.Database
	pool     *pgxpool.Pool
	batch    string
	apply    bool
}

func main() {
	config, err := parseConfiguration(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.mongoURI))
	if err != nil {
		fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	postgresPool, err := pgxpool.New(ctx, config.postgresURI)
	if err != nil {
		fatal(err)
	}
	defer postgresPool.Close()
	if err := postgresPool.Ping(ctx); err != nil {
		fatal(err)
	}

	migrate := &migrator{
		database: mongoClient.Database(config.mongoDB),
		pool:     postgresPool,
		batch:    config.batchLabel,
		apply:    config.apply,
	}
	if migrate.batch == "" {
		migrate.batch = time.Now().UTC().Format("20060102T150405Z")
	}

	summary, complete, err := migrate.migrateDailyUserLogs(ctx, config)
	if err != nil {
		fatal(fmt.Errorf("migrate daily user logs: %w", err))
	}
	fmt.Printf("daily-user-logs: scanned=%d migrated=%d skipped=%d quarantined=%d\n", summary.Scanned, summary.Migrated, summary.Skipped, summary.Quarantined)
	if !complete {
		fmt.Println("stopped at --limit; rerun to resume from the first incomplete unit")
		return
	}

	if !config.apply {
		fmt.Println("preflight passed; rerun with --apply to write PostgreSQL daily-log rows")
		return
	}

	report, err := migrate.reconcile(ctx)
	if err != nil {
		fatal(fmt.Errorf("reconcile: %w", err))
	}
	fmt.Print(report.String())
	if len(report.Mismatches) > 0 {
		os.Exit(1)
	}
}

func parseConfiguration(arguments []string) (configuration, error) {
	flags := flag.NewFlagSet("mongo_dailyuserlogs_to_postgres", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	config := configuration{}
	flags.BoolVar(&config.apply, "apply", false, "write PostgreSQL daily-log rows")
	flags.Int64Var(&config.batchSize, "batch-size", 200, "source documents read per MongoDB page")
	flags.IntVar(&config.limit, "limit", 0, "stop after this many committed units (0 migrates all)")
	flags.StringVar(&config.batchLabel, "batch-label", "", "ledger batch label; defaults to a UTC timestamp")
	flags.StringVar(&config.mongoURI, "mongo-uri", os.Getenv("MONGODB_URI"), "MongoDB connection URI")
	flags.StringVar(&config.mongoDB, "mongo-database", os.Getenv("MONGODB_DATABASE"), "MongoDB database name")
	flags.StringVar(&config.postgresURI, "postgres-uri", os.Getenv("POSTGRES_APPLICATION_URI"), "PostgreSQL connection URI")
	if err := flags.Parse(arguments); err != nil {
		return config, err
	}
	if config.mongoURI == "" || config.mongoDB == "" || config.postgresURI == "" {
		return config, errors.New("MONGODB_URI, MONGODB_DATABASE, and POSTGRES_APPLICATION_URI must be set")
	}
	if config.batchSize < 1 {
		return config, errors.New("batch size must be positive")
	}
	if config.limit < 0 {
		return config, errors.New("limit must not be negative")
	}
	return config, nil
}

// recordQuarantine records an unmigratable retained log. The quarantine ledger
// is append-only, so a replay preserves the same decision, and it is a no-op
// during a read-only preflight.
func (m *migrator) recordQuarantine(ctx context.Context, tx querier, kind, legacyID, reason, detail string) error {
	if !m.apply {
		return nil
	}
	_, err := tx.Exec(ctx, `INSERT INTO migration_quarantine (kind, legacy_id, reason, detail, batch)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (kind, legacy_id, reason) DO NOTHING`,
		kind, legacyID, reason, detail, m.batch)
	return err
}

// recordLedger records a completed unit inside its own transaction so the
// target rows and the completion marker commit atomically.
func (m *migrator) recordLedger(ctx context.Context, tx querier, kind, legacyID, targetID string) error {
	if !m.apply {
		return nil
	}
	_, err := tx.Exec(ctx, `INSERT INTO migration_ledger (kind, legacy_id, target_id, batch)
VALUES ($1, $2, $3, $4)
ON CONFLICT (kind, legacy_id) DO NOTHING`,
		kind, legacyID, targetID, m.batch)
	return err
}

// completedUnit returns the recorded target identity for a completed unit.
func (m *migrator) completedUnit(ctx context.Context, kind, legacyID string) (string, bool, error) {
	var targetID string
	err := m.pool.QueryRow(ctx, `SELECT target_id FROM migration_ledger WHERE kind = $1 AND legacy_id = $2`, kind, legacyID).Scan(&targetID)
	if err != nil {
		if isNoRows(err) {
			return "", false, nil
		}
		return "", false, err
	}
	return targetID, true, nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
