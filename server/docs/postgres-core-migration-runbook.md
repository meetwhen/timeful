# PostgreSQL Core Migration Runbook

## Scope And Status

This runbook covers the one-time backfill of existing events, responses, attendees, signup blocks and responses, folders, and folder membership from MongoDB to PostgreSQL.
It implements the [PostgreSQL Core Migration Contracts](postgres-core-migration-contracts.md) and complements the [PostgreSQL Anonymous Event Compatibility Contract](postgres-anonymous-event-compatibility.md).
The tooling and its isolated rehearsal were delivered by TASK-0190.06, and TASK-0199.10 added the retained daily-log backfill and its reconciliation.
TASK-0190.08 added the backup and restore rehearsal and the final isolated cutover rehearsal recorded here, and the operator-facing staging and production procedure lives in the [PostgreSQL Staging And Production Cutover Runbook](../../docs/postgres-staging-rollout.md).
TASK-0199.09.03 retired the one-off backfill commands and their representative fixtures after the rehearsals passed, so this runbook now records the executed behavior and evidence as history.
No procedure in this runbook can be run again from the repository.
A live cutover is a separately scheduled operational action with a named operator and an approved change window.

This runbook covers the core-record backfill only.
[PostgreSQL Retained-Data Migration Contracts](postgres-retained-data-contracts.md) supersedes this runbook's statements that calendar, OTP, friend-request, and daily-log collections stay in MongoDB, and governs the retained-data backfill and the final MongoDB removal.

## Backfill Tooling And Retirement

Five dated one-off commands under `server/scripts/` performed the backfill, each reading MongoDB and writing PostgreSQL without ever writing MongoDB:

- `20260815_mongo_anonymous_events_to_postgres` copied anonymous legacy events.
- `20260910_mongo_accounts_to_postgres` copied account units and ran first so ownership and response references resolved through `platform_identities.external_user_id`.
- `20260910_mongo_events_to_postgres` copied events with their responses, signup data, attendees, folders, and memberships, and carried the backup and restore rehearsal.
- `20260911_mongo_calendars_to_postgres` copied retained calendar integration units.
- `20260911_mongo_dailyuserlogs_to_postgres` copied retained historical daily user logs.

Each command supported `--apply` to write PostgreSQL, `--batch-size` to bound a MongoDB page, `--limit` to simulate interruption, and `--batch-label` to stamp ledger and quarantine rows; without `--apply` it was a read-only preflight.
TASK-0199.09.03 deleted these commands, their representative fixture sets, and the script-only `server/scripts/internal/legacybson` package they shared.

The executed evidence recorded before deletion:

- The full isolated backend suite passed every package on 2026-09-11, including the account, event, calendar, and daily-log backfill rehearsals, the analytics tests, and the backup and restore rehearsal.
- The backup and restore rehearsal restored the schema and reconciled eleven representative relations by row count and full-row digest.
- The daily-log rehearsal covered overlapping membership, an empty day, two retained documents sharing a date, missing-owner quarantine, replay, resume after interruption, source immutability, and clean reconciliation, and its replay reported every migrated unit as skipped.

The tooling retirement is irreversible: the backfill commands no longer exist, so re-copying a retained record is possible only from a verified PostgreSQL backup or from the retained MongoDB source while it still exists.
Once the MongoDB collections are retired, the PostgreSQL backup is the only recovery artifact.

## Migration Units And Ordering

A migration unit is one aggregate copied atomically in a single PostgreSQL transaction.
Units are migrated in dependency order: accounts, then events with their responses, signup data, and attendees, then folders with their memberships.
The event tool owned the event, response, attendee, signup, folder, and membership units; the account tool owned account units; the retained calendar backfill owned calendar integration units; and the retained daily-log backfill owned historical daily-log units.
OTP, friend-request, and daily-log collections were never copied by the event tool.

Each unit stamped its completion in `migration_ledger` inside the same transaction, including the fresh PostgreSQL target identity.
Folders rewrote each membership's legacy event reference to the migrated `postgres_events.id` recorded in the ledger, so migrated relationships never depended on a permanent legacy-identifier map.

## Resumability And Interruption Recovery

The run cursor advanced only after a unit transaction committed.
A completed unit was skipped on the next run, and an incomplete unit was retried from scratch because its transaction rolled back.
Recovery was replay, not manual repair: rerunning the same command with the same or a new batch label.
Replaying a completed run reported every unit as skipped and wrote nothing.
The `--limit` flag simulated interruption and proved the resume point, and the isolated rehearsal covered both interruption and replay.

## Source-Write Coordination

The backfill ran against the live MongoDB source without a maintenance window:

1. The initial pass copied units in batches while the application kept serving MongoDB.
2. A bounded catch-up pass re-read source records changed during the initial pass.
   Legacy event documents carry only an insertion timestamp, so catch-up reran the migration for the affected units rather than relying on a change cursor.
3. After catch-up, the affected record kind entered a short write freeze in which new mutations were rejected or queued and never dual-written.
4. During the freeze the migration performed one final idempotent pass and reconciliation.
5. Only after reconciliation passed were reads and writes switched for that record kind to PostgreSQL, ending the freeze.

No permanent dual write existed; a record is authored in exactly one store after its cutover.
Because the catch-up and freeze windows are operational, the tooling recorded the batch label and reconciliation evidence so the operator can evidence which pass produced the cutover state.

## Reconciliation Evidence

The event and folder run reconciled after a complete pass and printed a report.
It failed the run on any mismatch and never repaired.
Checks:

- Per-kind target counts for events, responses, signup blocks, signup responses, attendees, folders, and memberships.
- `num_responses` equality against the copied response rows for every non-signup migrated event.
- Referential integrity for folder membership and signup block claims.
- Quarantine totals grouped by reason code.

Reconciliation evidence is the input to the cutover gate; the operator records the printed report with the change ticket.

## Backup And Restore

Backups use the PostgreSQL 18 client tools with a custom-format dump and the least-privilege backup role.
The backup role needs read access to every table, granted at creation and re-applied to existing databases with:

```sql
GRANT pg_read_all_data TO <backup role> WITH INHERIT TRUE;
```

The operator commands for a dump, a scratch restore, and reconciliation live in the [PostgreSQL Staging And Production Cutover Runbook](../../docs/postgres-staging-rollout.md).
The retired backup and restore rehearsal seeded the representative fixture set, ran the event and folder migration, dumped the test database through the backup role, restored it into a fresh scratch database through the bootstrap role, and reconciled the restored table list plus every representative migrated relation by row count and full-row digest.
The 2026-09-11 run restored the schema and reconciled eleven relations with matching counts and digests.
A reconciliation mismatch failed the rehearsal.

## Quarantine Handling

Ambiguous or corrupt records are reported in `migration_quarantine` and are never repaired by inference.

| Reason code                  | Handling                                                                                        |
| ---------------------------- | ----------------------------------------------------------------------------------------------- |
| `missing-owner-account`      | The event migrates ownerless; no ownership is invented.                                         |
| `incomplete-token-ownership` | The response migrates under a fresh Event Visitor Identity; the missing credential is recorded. |
| `missing-response-identity`  | The response is not imported because it has no account or guest identity.                       |
| `invalid-guest-name`         | The response is not imported because the shared guest-name normalizer rejects it.               |
| `orphan-response`            | The response references an absent event and is not imported.                                    |
| `orphan-membership`          | The folder membership references an unmigrated event and is not imported.                       |
| `legacy-guest-credential`    | The legacy `guestEditToken` is recorded and never imported as PostgreSQL authority.             |

Migrated guest responses receive no base Event Visitor Control Credential and no Event Owner Edit Token.
A guest whose legacy token is quarantined regains PostgreSQL authority only by establishing new authority, such as creating a new response or Event Sign-In.
The quarantine ledger is append-only, so replayed runs preserve the same decisions.

## Cutover Gate And Rollback Boundaries

Before the write freeze, rollback is trivial: MongoDB remains unmodified and authoritative, and disabling the PostgreSQL route for the affected kind returns to MongoDB.
The write freeze is the practical point of no return because mutations accepted by PostgreSQL are not dual-written, so a MongoDB rollback after the freeze loses post-cutover writes unless they are first exported back.
The operational rollback boundary for each record kind is therefore the start of its freeze, not the end.

Cutover of a record kind requires:

- A complete event and folder pass with the intended batch label.
- A replay that reports every unit skipped.
- A reconciliation report with no mismatches.
- A reviewed quarantine report with every unresolved record either accepted or separately resolved.

## Retention And Cleanup

The schema is additive and never dropped during migration.
MongoDB source documents were retained unmodified through the retention window and final validation.
Runtime reads and writes are PostgreSQL-only for accounts, events, responses, attendees, signup data, groups, folders, calendar integrations, OTP challenges, and daily user logs; retained MongoDB `users` documents are recovery source that is never read or written at runtime.
After cutover validation and the retention window, drop `migration_ledger` and `migration_quarantine`; they are operational tooling and are never read by request paths.
The backfill tooling that wrote those tables was retired by TASK-0199.09.03, so they now hold only the last executed run's records.
MongoDB runtime and configuration removal is a separate stage owned by TASK-0199.09.04 and starts only after the core-record cutover in TASK-0190.08 and the retained-data cutover both pass.
The [Remaining MongoDB Collections And Deferred Scope](#remaining-mongodb-collections-and-deferred-scope) section records what still exists at that point and which task removes it.

## Executed Rehearsal Evidence

The isolated rehearsals ran inside the isolated Compose test stack and never targeted a development or production database.
The 2026-09-11 run recorded all backend packages passing.

- The account, event, calendar, and daily-log backfill rehearsals seeded one record per quarantine path, asserted target counts, ownership and response access, nullable and schedule fields, the quarantine breakdown, source immutability, clean reconciliation, idempotent replay, and resume after interruption.
- The daily-log rehearsal covered overlapping membership, an empty day, two retained documents sharing a date, missing-owner quarantine, replay, resume after interruption, and source immutability.
- The backup and restore rehearsal restored the schema and reconciled eleven representative relations by row count and full-row digest.
- The retained calendar rehearsal covered Google, Outlook, Apple, ICS, preferences, a two-connection legacy key suffix, missing-owner quarantine, replay, and resume after interruption.

The [PostgreSQL Staging And Production Cutover Runbook](../../docs/postgres-staging-rollout.md) records the matching browser rehearsal results.

## Final Isolated Cutover Rehearsal

TASK-0190.08 rehearsed the whole cutover in the isolated test stack and recorded the evidence in the [PostgreSQL Staging And Production Cutover Runbook](../../docs/postgres-staging-rollout.md).

- Backend: `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test go test ./... -count=1`.
  Before the tooling retirement this included the account backfill, event backfill, calendar backfill, daily-log backfill, analytics, and backup and restore rehearsals; it now covers the remaining PostgreSQL route, store, and model suites.
- Browser: from `e2e/`, run `E2E_FRONTEND=bundled npm run test:e2e -- --project=firefox-desktop`, then `npm run test:e2e -- --project=chromium-desktop --project=chromium-mobile --project=firefox-touch`, then `npm run test:e2e -- --project=chromium-production-desktop --project=chromium-production-mobile`.
- The coverage map for account sign-in, every event kind, groups, signup forms, folders, event links, identity and owner authority, calendar integration, OTP challenges, friend-request retirement, historical daily logs, and backup and restore lives in the staging and production runbook.

A rehearsal failure blocks cutover.

## Remaining MongoDB Collections And Deferred Scope

PostgreSQL is authoritative after cutover for accounts, events, responses, attendees, signup data, groups, folders, calendar integrations, OTP challenges, and daily user logs.
The MongoDB collections below remain only as recovery source or for legacy-only behavior; none is a second authority for a migrated record kind.

| Collection                                                         | Content at final cutover                                                                   | Runtime use                                                                                                                                     | Deferred scope                                                                 |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| `events`, `eventResponses`, `attendees`, `folders`, `folderEvents` | Legacy and noncanonical event, response, attendee, folder, and membership documents        | Served only by legacy routes for payload shapes the PostgreSQL creation classifier does not claim; supported records are authored in PostgreSQL | Removed with the MongoDB runtime by TASK-0199.09 after both cutovers pass.     |
| `users`                                                            | Retained account, calendar, and preference documents from before the retained-data cutover | Never read or written at runtime; read only by the retired one-off backfill tools                                                               | Recovery source until the retention window ends, then removed by TASK-0199.09. |
| `dailyuserlogs`                                                    | Historical daily user logs written before the daily-log cutover                            | Never read or written at runtime; new logs go to PostgreSQL                                                                                     | Backfilled by TASK-0199.10, then recovery source until removal.                |
| `otpCodes`                                                         | Challenges from before the OTP cutover                                                     | Never read or written                                                                                                                           | They expire and are removed with MongoDB by TASK-0199.09.                      |
| `friendrequests`                                                   | Dormant request documents                                                                  | Never read or written                                                                                                                           | Retired by TASK-0199.08 and removed with MongoDB by TASK-0199.09.              |

Integration-only user fields have PostgreSQL destinations under the retained-data contracts: calendar connections, sub-calendars, and preferences move to `calendar_accounts`, `calendar_sub_calendars`, `calendar_account_credentials`, and `calendar_preferences`; OTP challenges move to `otp_challenges`; daily-log membership moves to `daily_user_logs` and `daily_user_log_members`; event creator attribution stays in `postgres_events.creator_posthog_id`; and active-user and signed-up-user reporting read `daily_user_logs` and `accounts`.
No integration-only user field remains authoritative in the retained `users` document, and copying a retained document back into runtime authority is never a recovery option.
