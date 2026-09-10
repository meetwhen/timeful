package postgres

import (
	"context"
	"time"
)

// monthlyActiveCreatorLookback is the window length used by the event-creator
// reporting queries. A creator is active when it appears on an event created in
// the 30 days before the reporting instant.
const monthlyActiveCreatorLookback = 30

// CountDistinctMonthlyActiveEventCreators returns the number of distinct
// creator_posthog_id values attributed to non-empty creators on events created
// in the half-open window [date-30d, date). Reading only PostgreSQL event
// storage counts every migrated or newly created event exactly once and never
// consults the MongoDB events collection.
//
// The legacy MongoDB aggregation filtered events by the ObjectID timestamp,
// which has second precision, so its upper bound was effectively exclusive of
// the reporting second. PostgreSQL created_at is backfilled from the legacy
// creation instant, so the half-open interval reproduces those numbers. Like
// the legacy query, it does not filter is_deleted.
func (r *Repository) CountDistinctMonthlyActiveEventCreators(ctx context.Context, date time.Time) (int64, error) {
	var count int64
	err := r.db.QueryRow(ctx, `SELECT count(DISTINCT creator_posthog_id)
FROM postgres_events
WHERE created_at >= $1
  AND created_at < $2
  AND creator_posthog_id IS NOT NULL
  AND creator_posthog_id <> ''`, date.AddDate(0, 0, -monthlyActiveCreatorLookback), date).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountDistinctMonthlyActiveEventCreatorsWithMoreThanXEvents returns the number
// of distinct non-empty creators with at least x events created in the same
// half-open window [date-30d, date). The aggregation runs entirely against
// PostgreSQL event storage, so a creator whose events span the legacy and new
// stores is never counted twice.
func (r *Repository) CountDistinctMonthlyActiveEventCreatorsWithMoreThanXEvents(ctx context.Context, date time.Time, x int) (int64, error) {
	var count int64
	err := r.db.QueryRow(ctx, `SELECT count(*)
FROM (
    SELECT creator_posthog_id
    FROM postgres_events
    WHERE created_at >= $1
      AND created_at < $2
      AND creator_posthog_id IS NOT NULL
      AND creator_posthog_id <> ''
    GROUP BY creator_posthog_id
    HAVING count(*) >= $3
) AS creators_with_at_least_x_events`, date.AddDate(0, 0, -monthlyActiveCreatorLookback), date, x).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
