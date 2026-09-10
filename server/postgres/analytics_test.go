package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

// insertAnalyticsEvent inserts one PostgreSQL event row with only the fields the
// creator aggregations read, so tests exercise the query window and creator
// filter directly against authoritative event storage.
func insertAnalyticsEvent(t *testing.T, ctx context.Context, tx pgx.Tx, creatorPosthogID *string, createdAt time.Time, isDeleted bool) {
	t.Helper()
	shortID, err := GenerateShortID()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO postgres_events
 (short_id, name, type, is_deleted, creator_posthog_id, created_at, updated_at, payload)
VALUES ($1, 'analytics', 'specific_dates', $2, $3, $4, $4, '{}'::jsonb)`,
		shortID, isDeleted, creatorPosthogID, createdAt); err != nil {
		t.Fatalf("insert analytics event: %v", err)
	}
}

func analyticsCreator(value string) *string { return &value }

// TestCountDistinctMonthlyActiveEventCreatorsReadsPostgresWindow proves the
// distinct-creator report counts each creator once over the half-open
// [date-30d, date) PostgreSQL window, ignores absent or empty attribution, and
// matches the legacy behavior of counting soft-deleted events.
func TestCountDistinctMonthlyActiveEventCreatorsReadsPostgresWindow(t *testing.T) {
	ctx, repo, tx := newAccountsTestRepository(t)
	date := time.Date(2026, 9, 10, 23, 59, 59, 0, time.UTC)
	lowerBound := date.AddDate(0, 0, -monthlyActiveCreatorLookback)

	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-a"), date.AddDate(0, 0, -1), false)
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-a"), date.AddDate(0, 0, -10), false)
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-b"), date.AddDate(0, 0, -5), false)
	// Exactly the lower bound is included.
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-c"), lowerBound, false)
	// Just before the lower bound is excluded.
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-d"), lowerBound.Add(-time.Second), false)
	// Exactly the reporting instant is excluded by the ObjectID-equivalent
	// half-open upper bound.
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-e"), date, false)
	// Absent and empty attribution are excluded.
	insertAnalyticsEvent(t, ctx, tx, nil, date.AddDate(0, 0, -2), false)
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator(""), date.AddDate(0, 0, -2), false)
	// Soft-deleted events were counted by the legacy aggregation, so they stay
	// counted here without consulting two stores.
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-h"), date.AddDate(0, 0, -3), true)

	got, err := repo.CountDistinctMonthlyActiveEventCreators(ctx, date)
	if err != nil {
		t.Fatal(err)
	}
	if got != 4 {
		t.Fatalf("distinct monthly active creators = %d, want 4 (creator-a, creator-b, creator-c, creator-h)", got)
	}
}

// TestCountDistinctMonthlyActiveEventCreatorsWithMoreThanXEvents proves the
// grouped report applies the same counting boundary and uses an inclusive
// threshold, so a creator with exactly x events qualifies and an event outside
// the window is never counted.
func TestCountDistinctMonthlyActiveEventCreatorsWithMoreThanXEvents(t *testing.T) {
	ctx, repo, tx := newAccountsTestRepository(t)
	date := time.Date(2026, 9, 10, 23, 59, 59, 0, time.UTC)

	// creator-a has three events in the window.
	for i := 0; i < 3; i++ {
		insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-a"), date.AddDate(0, 0, -(i+1)), false)
	}
	// creator-b has exactly two events in the window.
	for i := 0; i < 2; i++ {
		insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-b"), date.AddDate(0, 0, -(i+1)), false)
	}
	// creator-c has two events, but one falls outside the window.
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-c"), date.AddDate(0, 0, -1), false)
	insertAnalyticsEvent(t, ctx, tx, analyticsCreator("creator-c"), date.AddDate(0, 0, -31), false)

	cases := []struct {
		x    int
		want int64
	}{
		{x: 4, want: 0},
		{x: 3, want: 1},
		{x: 2, want: 2},
		{x: 1, want: 3},
	}
	for _, tc := range cases {
		got, err := repo.CountDistinctMonthlyActiveEventCreatorsWithMoreThanXEvents(ctx, date, tc.x)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Fatalf("creators with >= %d events = %d, want %d", tc.x, got, tc.want)
		}
	}
}

// TestCreatorAnalyticsCountEachPostgresEventOnce proves that a creator with
// multiple events is counted once by the distinct report and that the grouped
// aggregation counts rows rather than distinct creators, so migrated and newly
// created PostgreSQL events each contribute exactly once.
func TestCreatorAnalyticsCountEachPostgresEventOnce(t *testing.T) {
	ctx, repo, tx := newAccountsTestRepository(t)
	date := time.Date(2026, 9, 10, 23, 59, 59, 0, time.UTC)

	migrated := analyticsCreator("creator-migrated")
	insertAnalyticsEvent(t, ctx, tx, migrated, date.AddDate(0, 0, -20), false)
	insertAnalyticsEvent(t, ctx, tx, migrated, date.AddDate(0, 0, -10), false)
	newEvent := analyticsCreator("creator-new")
	insertAnalyticsEvent(t, ctx, tx, newEvent, date.AddDate(0, 0, -1), false)

	distinct, err := repo.CountDistinctMonthlyActiveEventCreators(ctx, date)
	if err != nil {
		t.Fatal(err)
	}
	if distinct != 2 {
		t.Fatalf("distinct creators = %d, want 2", distinct)
	}

	atLeastTwo, err := repo.CountDistinctMonthlyActiveEventCreatorsWithMoreThanXEvents(ctx, date, 2)
	if err != nil {
		t.Fatal(err)
	}
	if atLeastTwo != 1 {
		t.Fatalf("creators with >= 2 events = %d, want 1 (the migrated creator)", atLeastTwo)
	}
}
