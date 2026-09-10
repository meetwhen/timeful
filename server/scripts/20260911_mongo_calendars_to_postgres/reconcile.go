package main

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

// reconciliationReport records the state observed after a migration run. The
// rehearsal treats any mismatch as a failure before cutover.
type reconciliationReport struct {
	Units              int
	Accounts           int
	SubCalendars       int
	CredentialRows     int
	Preferences        int
	Quarantined        int
	QuarantineByReason map[string]int
	Mismatches         []string
}

func (report reconciliationReport) String() string {
	var builder strings.Builder
	builder.WriteString("reconciliation:\n")
	fmt.Fprintf(&builder, "  units=%d calendar_accounts=%d sub_calendars=%d credential_rows=%d preferences=%d\n",
		report.Units, report.Accounts, report.SubCalendars, report.CredentialRows, report.Preferences)
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

// reconcile counts every migrated calendar record through the ledger and checks
// that credentials stay in the versioned AES-GCM envelope. It never repairs; a
// mismatch fails the run.
func (m *migrator) reconcile(ctx context.Context) (reconciliationReport, error) {
	report := reconciliationReport{QuarantineByReason: map[string]int{}}
	if err := m.pool.QueryRow(ctx, `SELECT count(*) FROM migration_ledger WHERE kind = $1`, ledgerKindCalendarAccount).Scan(&report.Units); err != nil {
		return report, err
	}

	counts := []struct {
		query string
		dest  *int
	}{
		{`SELECT count(*) FROM calendar_accounts a
JOIN platform_identities p ON p.id = a.platform_identity_id
JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id`, &report.Accounts},
		{`SELECT count(*) FROM calendar_sub_calendars s
JOIN calendar_accounts a ON a.id = s.calendar_account_id
JOIN platform_identities p ON p.id = a.platform_identity_id
JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id`, &report.SubCalendars},
		{`SELECT count(*) FROM calendar_account_credentials c
JOIN calendar_accounts a ON a.id = c.calendar_account_id
JOIN platform_identities p ON p.id = a.platform_identity_id
JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id`, &report.CredentialRows},
		{`SELECT count(*) FROM calendar_preferences pref
JOIN platform_identities p ON p.id = pref.platform_identity_id
JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id`, &report.Preferences},
	}
	for _, count := range counts {
		if err := m.pool.QueryRow(ctx, count.query).Scan(count.dest); err != nil {
			return report, err
		}
	}

	m.checkCredentialEnvelopes(ctx, &report)
	if err := m.loadQuarantineSummary(ctx, &report); err != nil {
		return report, err
	}
	return report, nil
}

// checkCredentialEnvelopes fails when a migrated provider secret is not stored
// in the versioned AES-256-GCM envelope, so a legacy CFB value or plaintext
// secret never survives the backfill.
func (m *migrator) checkCredentialEnvelopes(ctx context.Context, report *reconciliationReport) {
	var unversioned int
	err := m.pool.QueryRow(ctx, `SELECT count(*) FROM calendar_account_credentials c
JOIN calendar_accounts a ON a.id = c.calendar_account_id
JOIN platform_identities p ON p.id = a.platform_identity_id
JOIN migration_ledger l ON l.kind = 'calendar-account' AND l.legacy_id = p.external_user_id
WHERE (c.oauth_access_token_ciphertext IS NOT NULL AND c.oauth_access_token_ciphertext NOT LIKE 'v1:%')
   OR (c.oauth_refresh_token_ciphertext IS NOT NULL AND c.oauth_refresh_token_ciphertext NOT LIKE 'v1:%')
   OR (c.apple_password_ciphertext IS NOT NULL AND c.apple_password_ciphertext NOT LIKE 'v1:%')
   OR (c.ics_feed_url_ciphertext IS NOT NULL AND c.ics_feed_url_ciphertext NOT LIKE 'v1:%')`).Scan(&unversioned)
	if err != nil {
		report.Mismatches = append(report.Mismatches, "credential envelope check failed: "+err.Error())
		return
	}
	if unversioned != 0 {
		report.Mismatches = append(report.Mismatches, fmt.Sprintf("%d credential rows are not in the v1 GCM envelope", unversioned))
	}
}

func (m *migrator) loadQuarantineSummary(ctx context.Context, report *reconciliationReport) error {
	rows, err := m.pool.Query(ctx, `SELECT reason, count(*) FROM migration_quarantine GROUP BY reason ORDER BY reason`)
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
