package main

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// reconcileBatchSize is the retained-document page size reconciliation reads.
const reconcileBatchSize int64 = 1000

// reconciliationReport records the state observed after a migration run. The
// rehearsal treats any mismatch as a failure before cutover.
type reconciliationReport struct {
	Units              int
	Logs               int
	Memberships        int
	SourceLogs         int
	SourceMemberships  int
	Quarantined        int
	QuarantineByReason map[string]int
	Mismatches         []string
}

func (report reconciliationReport) String() string {
	var builder strings.Builder
	builder.WriteString("reconciliation:\n")
	fmt.Fprintf(&builder, "  units=%d daily_user_logs=%d memberships=%d source_logs=%d source_memberships=%d\n",
		report.Units, report.Logs, report.Memberships, report.SourceLogs, report.SourceMemberships)
	fmt.Fprintf(&builder, "  quarantined=%d\n", report.Quarantined)
	reasons := make([]string, 0, len(report.QuarantineByReason))
	for reason := range report.QuarantineByReason {
		reasons = append(reasons, reason)
	}
	sort.Strings(reasons)
	for _, reason := range reasons {
		fmt.Fprintf(&builder, "    %s=%d\n", reason, report.QuarantineByReason[reason])
	}
	if len(report.Mismatches) == 0 {
		builder.WriteString("  mismatches: none\n")
		return builder.String()
	}
	builder.WriteString("  mismatches:\n")
	for _, mismatch := range report.Mismatches {
		fmt.Fprintf(&builder, "    - %s\n", mismatch)
	}
	return builder.String()
}

// reconcile compares the retained MongoDB source with the migrated PostgreSQL
// logs through the ledger. It verifies unit, daily-log, and membership counts,
// date bucketing, and first-seen membership order. It never repairs; a mismatch
// fails the run.
func (m *migrator) reconcile(ctx context.Context) (reconciliationReport, error) {
	report := reconciliationReport{QuarantineByReason: map[string]int{}}

	ledger, err := m.loadLedgerTargets(ctx)
	if err != nil {
		return report, err
	}
	quarantined, err := m.loadQuarantinedLegacyIDs(ctx)
	if err != nil {
		return report, err
	}
	report.Units = len(ledger)

	expected, sourceLogs, unaccounted, err := m.loadExpectedDailyLogs(ctx, ledger, quarantined)
	if err != nil {
		return report, err
	}
	report.SourceLogs = sourceLogs
	report.SourceMemberships = countMembers(expected)

	target, targetIDs, err := m.loadTargetDailyLogs(ctx)
	if err != nil {
		return report, err
	}
	report.Logs = len(targetIDs)
	report.Memberships = countMembers(target)

	if report.SourceLogs != report.Units {
		report.Mismatches = append(report.Mismatches,
			fmt.Sprintf("%d source logs back %d ledger units", report.SourceLogs, report.Units))
	}
	if len(unaccounted) > 0 {
		report.Mismatches = append(report.Mismatches,
			fmt.Sprintf("%d retained daily logs are neither migrated nor quarantined: %s",
				len(unaccounted), strings.Join(firstN(unaccounted, 5), ", ")))
	}
	if report.Logs != len(expected) {
		report.Mismatches = append(report.Mismatches,
			fmt.Sprintf("%d target logs cover %d expected dates", report.Logs, len(expected)))
	}
	compareDailyLogBuckets(&report, expected, target)

	if err := m.loadQuarantineSummary(ctx, &report); err != nil {
		return report, err
	}
	return report, nil
}

// loadLedgerTargets returns the completed legacy identity to fresh PostgreSQL
// log identity mapping.
func (m *migrator) loadLedgerTargets(ctx context.Context) (map[string]string, error) {
	rows, err := m.pool.Query(ctx, `SELECT legacy_id, target_id FROM migration_ledger WHERE kind = $1`, ledgerKindDailyUserLog)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	targets := map[string]string{}
	for rows.Next() {
		var legacyID, targetID string
		if err := rows.Scan(&legacyID, &targetID); err != nil {
			return nil, err
		}
		targets[legacyID] = targetID
	}
	return targets, rows.Err()
}

// loadQuarantinedLegacyIDs returns every retained log recorded in the
// append-only quarantine ledger for this kind.
func (m *migrator) loadQuarantinedLegacyIDs(ctx context.Context) (map[string]bool, error) {
	rows, err := m.pool.Query(ctx, `SELECT legacy_id FROM migration_quarantine WHERE kind = $1`, ledgerKindDailyUserLog)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quarantined := map[string]bool{}
	for rows.Next() {
		var legacyID string
		if err := rows.Scan(&legacyID); err != nil {
			return nil, err
		}
		quarantined[legacyID] = true
	}
	return quarantined, rows.Err()
}

// loadExpectedDailyLogs pages the retained source and builds the expected
// PostgreSQL state keyed by the account-local date. Overlapping retained
// documents on one date merge their membership in _id order, matching the
// migration's append order. A source log that is neither quarantined nor
// ledger-backed is reported as unaccounted.
func (m *migrator) loadExpectedDailyLogs(ctx context.Context, ledger map[string]string, quarantined map[string]bool) (map[string][]string, int, []string, error) {
	expected := map[string][]string{}
	sourceLogs := 0
	unaccounted := []string{}

	var lastID primitive.ObjectID
	hasCursor := false
	for {
		cursor, err := m.database.Collection("dailyuserlogs").Find(ctx, pageFilter(lastID, hasCursor), options.Find().SetSort(bson.M{"_id": 1}).SetLimit(reconcileBatchSize))
		if err != nil {
			return nil, 0, nil, err
		}
		logs := make([]legacyDailyUserLog, 0, reconcileBatchSize)
		if err := cursor.All(ctx, &logs); err != nil {
			cursor.Close(ctx)
			return nil, 0, nil, err
		}
		cursor.Close(ctx)
		if len(logs) == 0 {
			break
		}

		for _, log := range logs {
			lastID = log.ID
			hasCursor = true
			legacyID := log.ID.Hex()
			if quarantined[legacyID] {
				continue
			}
			if _, completed := ledger[legacyID]; !completed {
				unaccounted = append(unaccounted, legacyID)
				continue
			}
			date := dailyUserLogDate(log.Date).Format("2006-01-02")
			expected[date] = mergeOrdered(expected[date], objectIDHexes(log.UserIDs))
			sourceLogs++
		}
		if int64(len(logs)) < reconcileBatchSize {
			break
		}
	}
	return expected, sourceLogs, unaccounted, nil
}

// loadTargetDailyLogs returns the migrated logs keyed by date with membership in
// first-seen order, plus the set of fresh PostgreSQL log identities.
func (m *migrator) loadTargetDailyLogs(ctx context.Context) (map[string][]string, map[string]bool, error) {
	rows, err := m.pool.Query(ctx, `SELECT l.id, l.log_date FROM daily_user_logs l
WHERE EXISTS (SELECT 1 FROM migration_ledger ml
WHERE ml.kind = $1 AND ml.target_id = l.id::text)`, ledgerKindDailyUserLog)
	if err != nil {
		return nil, nil, err
	}
	targetIDs := map[string]bool{}
	dates := map[string]string{}
	for rows.Next() {
		var logID string
		var logDate time.Time
		if err := rows.Scan(&logID, &logDate); err != nil {
			rows.Close()
			return nil, nil, err
		}
		targetIDs[logID] = true
		dates[logID] = logDate.Format("2006-01-02")
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	memberRows, err := m.pool.Query(ctx, `SELECT m.daily_user_log_id, m.account_user_id FROM daily_user_log_members m
WHERE EXISTS (SELECT 1 FROM migration_ledger ml
WHERE ml.kind = $1 AND ml.target_id = m.daily_user_log_id::text)
ORDER BY m.daily_user_log_id, m.first_seen_position, m.id`, ledgerKindDailyUserLog)
	if err != nil {
		return nil, nil, err
	}
	defer memberRows.Close()

	members := map[string][]string{}
	for memberRows.Next() {
		var logID, accountUserID string
		if err := memberRows.Scan(&logID, &accountUserID); err != nil {
			return nil, nil, err
		}
		members[logID] = append(members[logID], accountUserID)
	}
	if err := memberRows.Err(); err != nil {
		return nil, nil, err
	}

	target := map[string][]string{}
	for logID := range targetIDs {
		target[dates[logID]] = members[logID]
	}
	return target, targetIDs, nil
}

// compareDailyLogBuckets fails on a missing or unexpected target date and on a
// membership order that differs from the retained first-seen order.
func compareDailyLogBuckets(report *reconciliationReport, expected, target map[string][]string) {
	expectedDates := sortedKeys(expected)
	for _, date := range expectedDates {
		got, ok := target[date]
		if !ok {
			report.Mismatches = append(report.Mismatches, "target is missing daily log "+date)
			continue
		}
		if !equalOrdered(got, expected[date]) {
			report.Mismatches = append(report.Mismatches,
				fmt.Sprintf("daily log %s membership order = %v, want %v", date, got, expected[date]))
		}
	}
	for _, date := range sortedKeys(target) {
		if _, ok := expected[date]; !ok {
			report.Mismatches = append(report.Mismatches, "target has unexpected daily log "+date)
		}
	}
}

func (m *migrator) loadQuarantineSummary(ctx context.Context, report *reconciliationReport) error {
	rows, err := m.pool.Query(ctx, `SELECT reason, count(*) FROM migration_quarantine WHERE kind = $1 GROUP BY reason ORDER BY reason`, ledgerKindDailyUserLog)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var reason string
		var count int
		if err := rows.Scan(&reason, &count); err != nil {
			return err
		}
		report.QuarantineByReason[reason] = count
		report.Quarantined += count
	}
	return rows.Err()
}

func countMembers(logs map[string][]string) int {
	total := 0
	for _, members := range logs {
		total += len(members)
	}
	return total
}

func equalOrdered(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func sortedKeys(logs map[string][]string) []string {
	keys := make([]string, 0, len(logs))
	for key := range logs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func firstN(values []string, n int) []string {
	if len(values) <= n {
		return values
	}
	return values[:n]
}
