package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// legacyDailyUserLog is one retained MongoDB dailyuserlogs document. It only
// decodes the fields the retained-data contract maps to PostgreSQL; the
// denormalized users array is never read.
type legacyDailyUserLog struct {
	ID      primitive.ObjectID   `bson:"_id"`
	Date    primitive.DateTime   `bson:"date"`
	UserIDs []primitive.ObjectID `bson:"userIds"`
}

// quarantineRecord is one unmigratable retained unit.
type quarantineRecord struct {
	legacyID string
	reason   string
	detail   string
}

// unitDailyUserLog is one retained daily log with its resolved membership. It
// is copied in a single PostgreSQL transaction that also records the ledger
// completion marker. logDate is the account-local calendar date at UTC
// midnight, which is also the PostgreSQL log identity key.
type unitDailyUserLog struct {
	legacyID   string
	logDate    time.Time
	members    []string
	quarantine []quarantineRecord
}

// dailyUserLogDate normalizes the retained log date to the account-local
// calendar date at UTC midnight. The legacy writer stored exactly that instant;
// normalizing keeps any time component from shifting the bucket.
func dailyUserLogDate(date primitive.DateTime) time.Time {
	stored := date.Time().UTC()
	return time.Date(stored.Year(), stored.Month(), stored.Day(), 0, 0, 0, 0, time.UTC)
}

// migrateDailyUserLogs pages the retained MongoDB dailyuserlogs collection and
// migrates one log per unit. It returns complete=false when --limit stopped the
// run before the source was exhausted so the caller can resume.
func (m *migrator) migrateDailyUserLogs(ctx context.Context, config configuration) (migrationSummary, bool, error) {
	collection := m.database.Collection("dailyuserlogs")
	summary := migrationSummary{}

	// hasCursor separates "no page read yet" from a legitimate zero ObjectID.
	var lastID primitive.ObjectID
	hasCursor := false
	for {
		cursor, err := collection.Find(ctx, pageFilter(lastID, hasCursor), options.Find().SetSort(bson.M{"_id": 1}).SetLimit(config.batchSize))
		if err != nil {
			return summary, false, err
		}
		logs := make([]legacyDailyUserLog, 0, config.batchSize)
		if err := cursor.All(ctx, &logs); err != nil {
			cursor.Close(ctx)
			return summary, false, err
		}
		cursor.Close(ctx)
		if len(logs) == 0 {
			break
		}

		for _, log := range logs {
			summary.Scanned++
			legacyID := log.ID.Hex()
			if _, completed, err := m.completedUnit(ctx, ledgerKindDailyUserLog, legacyID); err != nil {
				return summary, false, err
			} else if completed {
				summary.Skipped++
				lastID = log.ID
				hasCursor = true
				continue
			}

			unit, err := m.buildDailyUserLogUnit(ctx, log)
			if err != nil {
				return summary, false, err
			}
			if len(unit.quarantine) > 0 {
				summary.Quarantined++
			} else {
				summary.Migrated++
			}
			if m.apply {
				if err := m.persistDailyUserLogUnit(ctx, unit); err != nil {
					return summary, false, err
				}
			}

			// Advance the source cursor only after the unit committed, so an
			// interruption resumes from the first incomplete log.
			lastID = log.ID
			hasCursor = true
			if config.limit > 0 && summary.Migrated >= config.limit {
				return summary, false, nil
			}
		}
		if int64(len(logs)) < config.batchSize {
			break
		}
	}
	return summary, true, nil
}

// buildDailyUserLogUnit resolves every membership owner through the
// authoritative PostgreSQL account. A log whose membership includes an account
// that does not resolve is quarantined with missing-owner-account and migrates
// nothing, so no membership is invented or silently dropped.
func (m *migrator) buildDailyUserLogUnit(ctx context.Context, log legacyDailyUserLog) (unitDailyUserLog, error) {
	unit := unitDailyUserLog{legacyID: log.ID.Hex(), logDate: dailyUserLogDate(log.Date)}
	members := make([]string, 0, len(log.UserIDs))
	missing := []string{}
	for _, userID := range orderedDistinct(objectIDHexes(log.UserIDs)) {
		resolved, err := m.memberAccountExists(ctx, userID)
		if err != nil {
			return unit, err
		}
		if !resolved {
			missing = append(missing, userID)
			continue
		}
		members = append(members, userID)
	}
	if len(missing) > 0 {
		unit.quarantine = append(unit.quarantine, quarantineRecord{
			legacyID: unit.legacyID,
			reason:   reasonMissingOwnerAccount,
			detail:   "no PostgreSQL account resolves daily-log members " + strings.Join(missing, ", "),
		})
		return unit, nil
	}
	unit.members = members
	return unit, nil
}

// memberAccountExists reports whether an external user identifier resolves to a
// platform identity that owns a PostgreSQL account row.
func (m *migrator) memberAccountExists(ctx context.Context, externalUserID string) (bool, error) {
	var exists bool
	err := m.pool.QueryRow(ctx, `SELECT EXISTS (
SELECT 1 FROM platform_identities p
JOIN accounts a ON a.platform_identity_id = p.id
WHERE p.external_user_id = $1)`, externalUserID).Scan(&exists)
	return exists, err
}

// persistDailyUserLogUnit writes the log and its membership plus the ledger
// completion marker in one transaction. Two retained documents that share a
// date upsert the same fresh log identity and append membership in source
// order, so overlapping days merge without duplicates. A quarantined log only
// records the quarantine decision.
func (m *migrator) persistDailyUserLogUnit(ctx context.Context, unit unitDailyUserLog) error {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, record := range unit.quarantine {
		if err := m.recordQuarantine(ctx, tx, ledgerKindDailyUserLog, record.legacyID, record.reason, record.detail); err != nil {
			return err
		}
	}
	if len(unit.quarantine) > 0 {
		return tx.Commit(ctx)
	}

	var logID string
	if err := tx.QueryRow(ctx, `INSERT INTO daily_user_logs (log_date) VALUES ($1)
ON CONFLICT (log_date) DO UPDATE SET updated_at = daily_user_logs.updated_at
RETURNING id`, unit.logDate).Scan(&logID); err != nil {
		return fmt.Errorf("upsert daily log %s: %w", unit.logDate.Format("2006-01-02"), err)
	}
	for _, accountUserID := range unit.members {
		if _, err := tx.Exec(ctx, `INSERT INTO daily_user_log_members (daily_user_log_id, account_user_id, first_seen_position)
VALUES ($1, $2, COALESCE((SELECT MAX(first_seen_position) + 1 FROM daily_user_log_members WHERE daily_user_log_id = $1), 0))
ON CONFLICT (daily_user_log_id, account_user_id) DO NOTHING`, logID, accountUserID); err != nil {
			return fmt.Errorf("append daily-log member %s: %w", accountUserID, err)
		}
	}
	if err := m.recordLedger(ctx, tx, ledgerKindDailyUserLog, unit.legacyID, logID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// objectIDHexes converts legacy account identifiers to their external form.
func objectIDHexes(objectIDs []primitive.ObjectID) []string {
	hexes := make([]string, 0, len(objectIDs))
	for _, objectID := range objectIDs {
		hexes = append(hexes, objectID.Hex())
	}
	return hexes
}
