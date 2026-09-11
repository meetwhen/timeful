package main

import (
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestDailyUserLogDateNormalizesStoredDate proves the retained log date maps to
// the account-local calendar date at UTC midnight, so a stored time component
// can never shift the PostgreSQL log bucket.
func TestDailyUserLogDateNormalizesStoredDate(t *testing.T) {
	cases := []struct {
		name string
		at   time.Time
		want string
	}{
		{"stored midnight keeps its date", time.Date(2026, 9, 10, 0, 0, 0, 0, time.UTC), "2026-09-10"},
		{"late evening stays on the stored date", time.Date(2026, 9, 10, 23, 30, 0, 0, time.UTC), "2026-09-10"},
		{"previous-day tail stays on the stored date", time.Date(2026, 9, 9, 23, 59, 59, 0, time.UTC), "2026-09-09"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := dailyUserLogDate(primitive.NewDateTimeFromTime(tc.at))
			if got.Format("2006-01-02") != tc.want {
				t.Fatalf("dailyUserLogDate(%s) = %s, want %s", tc.at, got.Format("2006-01-02"), tc.want)
			}
			if !got.Equal(time.Date(got.Year(), got.Month(), got.Day(), 0, 0, 0, 0, time.UTC)) {
				t.Fatalf("dailyUserLogDate must be UTC midnight, got %s", got)
			}
		})
	}
}

// TestOrderedDistinctPreservesFirstSeenOrder proves duplicate membership
// collapses to its first occurrence because the PostgreSQL target allows one
// row per account per log.
func TestOrderedDistinctPreservesFirstSeenOrder(t *testing.T) {
	got := orderedDistinct([]string{"a", "b", "a", "c", "b"})
	want := []string{"a", "b", "c"}
	if !equalOrdered(got, want) {
		t.Fatalf("orderedDistinct = %v, want %v", got, want)
	}
}

// TestMergeOrderedAppendsOverlappingDates proves two retained documents that
// share a date merge membership in source order without repeating an account.
func TestMergeOrderedAppendsOverlappingDates(t *testing.T) {
	got := mergeOrdered([]string{"c", "a"}, []string{"a", "d"})
	want := []string{"c", "a", "d"}
	if !equalOrdered(got, want) {
		t.Fatalf("mergeOrdered = %v, want %v", got, want)
	}
}

func TestReconciliationReportFormatting(t *testing.T) {
	report := reconciliationReport{
		Units:              4,
		Logs:               3,
		Memberships:        7,
		SourceLogs:         4,
		SourceMemberships:  7,
		Quarantined:        1,
		QuarantineByReason: map[string]int{reasonMissingOwnerAccount: 1},
		Mismatches:         []string{"example mismatch"},
	}
	output := report.String()
	for _, want := range []string{
		"units=4 daily_user_logs=3 memberships=7 source_logs=4 source_memberships=7",
		"quarantined=1",
		"missing-owner-account=1",
		"example mismatch",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("report missing %q:\n%s", want, output)
		}
	}
}
