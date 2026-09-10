package postgres

import (
	"context"
	"errors"
	"time"
)

// DailyUserLog is a PostgreSQL-owned historical daily user log. LogDate is the
// account-local month/day/year at UTC midnight, and Members holds one row per
// account that signed in that day in first-seen order.
type DailyUserLog struct {
	ID      string
	LogDate time.Time
	Members []DailyUserLogMember
}

// DailyUserLogMember is one account's membership in a daily log. The profile
// fields are rebuilt from the authoritative accounts table at read time and are
// never stored on the log.
type DailyUserLogMember struct {
	AccountUserID string
	FirstName     string
	LastName      string
	Email         string
	Position      int
}

// dailyLogDate returns the start of the account-local month/day/year as UTC
// midnight. It preserves the legacy bucketing: shift the server instant by the
// account's timezone offset and take that calendar date, so a sign-in late at
// night and one the next morning remain distinct in the account's own timezone.
func dailyLogDate(now time.Time, timezoneOffset int) time.Time {
	adjusted := now.Add(time.Duration(timezoneOffset) * time.Minute).UTC()
	return time.Date(adjusted.Year(), adjusted.Month(), adjusted.Day(), 0, 0, 0, 0, time.UTC)
}

// RecordDailyUserLogMembership records one account's sign-in for its
// account-local day. It is idempotent per account per day and appends new
// accounts after existing members so first-seen order is preserved. The log and
// its new membership are written in one transaction.
func (r *Repository) RecordDailyUserLogMembership(ctx context.Context, accountUserID string, timezoneOffset int) error {
	return r.recordDailyUserLogMembershipAt(ctx, accountUserID, timezoneOffset, time.Now())
}

// recordDailyUserLogMembershipAt is the deterministic core of
// RecordDailyUserLogMembership; tests call it with a fixed instant to exercise
// timezone bucketing without depending on the wall clock.
func (r *Repository) recordDailyUserLogMembershipAt(ctx context.Context, accountUserID string, timezoneOffset int, now time.Time) error {
	if accountUserID == "" {
		return errors.New("daily log account user ID is required")
	}
	logDate := dailyLogDate(now, timezoneOffset)
	return r.withTransaction(ctx, func(ctx context.Context, tx *Repository) error {
		var logID string
		if err := tx.db.QueryRow(ctx, `INSERT INTO daily_user_logs (log_date) VALUES ($1)
ON CONFLICT (log_date) DO UPDATE SET updated_at = daily_user_logs.updated_at
RETURNING id`, logDate).Scan(&logID); err != nil {
			return err
		}
		_, err := tx.db.Exec(ctx, `INSERT INTO daily_user_log_members (daily_user_log_id, account_user_id, first_seen_position)
VALUES ($1, $2, COALESCE((SELECT MAX(first_seen_position) + 1 FROM daily_user_log_members WHERE daily_user_log_id = $1), 0))
ON CONFLICT (daily_user_log_id, account_user_id) DO NOTHING`, logID, accountUserID)
		return err
	})
}

// ListDailyUserLogs returns every daily log on or after startDate, newest first,
// with each log's account members in first-seen order and their authoritative
// profile fields. A log with no members is returned with an empty member list.
func (r *Repository) ListDailyUserLogs(ctx context.Context, startDate time.Time) ([]DailyUserLog, error) {
	rows, err := r.db.Query(ctx, `SELECT l.id, l.log_date, m.account_user_id, COALESCE(a.first_name, ''), COALESCE(a.last_name, ''), COALESCE(a.email, ''), m.first_seen_position
FROM daily_user_logs l
LEFT JOIN daily_user_log_members m ON m.daily_user_log_id = l.id
LEFT JOIN platform_identities p ON p.external_user_id = m.account_user_id
LEFT JOIN accounts a ON a.platform_identity_id = p.id
WHERE l.log_date >= $1::date
ORDER BY l.log_date DESC, m.first_seen_position, m.id`, startDate.UTC().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := []DailyUserLog{}
	var current *DailyUserLog
	for rows.Next() {
		var logID string
		var logDate time.Time
		var accountUserID, firstName, lastName, email *string
		var position *int
		if err := rows.Scan(&logID, &logDate, &accountUserID, &firstName, &lastName, &email, &position); err != nil {
			return nil, err
		}
		if current == nil || current.ID != logID {
			logs = append(logs, DailyUserLog{ID: logID, LogDate: logDate, Members: []DailyUserLogMember{}})
			current = &logs[len(logs)-1]
		}
		if accountUserID != nil {
			current.Members = append(current.Members, DailyUserLogMember{
				AccountUserID: *accountUserID,
				FirstName:     derefString(firstName),
				LastName:      derefString(lastName),
				Email:         derefString(email),
				Position:      derefInt(position),
			})
		}
	}
	return logs, rows.Err()
}

// ListActiveUserDays returns active-user reporting days from startDate up to now,
// newest first, padding days without a log as empty so the reporting output
// lists every day in the range. Existing logs keep their first-seen member order.
func (r *Repository) ListActiveUserDays(ctx context.Context, startDate, now time.Time) ([]DailyUserLog, error) {
	logs, err := r.ListDailyUserLogs(ctx, startDate)
	if err != nil {
		return nil, err
	}
	return padDailyUserLogs(logs, startDate, now), nil
}

// CountAccounts returns the number of signed-up accounts for reporting.
func (r *Repository) CountAccounts(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.QueryRow(ctx, `SELECT count(*) FROM accounts`).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// padDailyUserLogs inserts empty days between startDate and now, keeping the
// slice newest first. It reproduces the legacy reporting padding so an empty
// day still appears in the list and chart with a zero count.
func padDailyUserLogs(logs []DailyUserLog, startDate, now time.Time) []DailyUserLog {
	curDate := startDate
	for i := len(logs) - 1; i >= 0; i-- {
		for !logs[i].LogDate.Equal(curDate) && curDate.Before(now) {
			logs = insertDailyUserLog(logs, i+1, DailyUserLog{LogDate: curDate, Members: []DailyUserLogMember{}})
			curDate = curDate.AddDate(0, 0, 1)
		}
		curDate = curDate.AddDate(0, 0, 1)
	}
	for curDate.Before(now) {
		logs = insertDailyUserLog(logs, 0, DailyUserLog{LogDate: curDate, Members: []DailyUserLogMember{}})
		curDate = curDate.AddDate(0, 0, 1)
	}
	return logs
}

func insertDailyUserLog(logs []DailyUserLog, index int, value DailyUserLog) []DailyUserLog {
	if index < 0 {
		index = 0
	}
	if index > len(logs) {
		index = len(logs)
	}
	logs = append(logs, DailyUserLog{})
	copy(logs[index+1:], logs[index:])
	logs[index] = value
	return logs
}

func derefString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func derefInt(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}
