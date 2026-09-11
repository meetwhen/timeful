# PostgreSQL Core Migration Runbook

## Scope And Status

This runbook covers the one-time backfill of existing events, responses, attendees, signup blocks and responses, folders, and folder membership from MongoDB to PostgreSQL.
It implements the [PostgreSQL Core Migration Contracts](postgres-core-migration-contracts.md) and complements the [PostgreSQL Anonymous Event Compatibility Contract](postgres-anonymous-event-compatibility.md).
The tooling and its isolated rehearsal are delivered by TASK-0190.06; no live deployment or production data migration is performed by that task.
TASK-0190.08 adds the backup and restore rehearsal and the final isolated cutover rehearsal recorded here, and the operator-facing staging and production procedure lives in the [PostgreSQL Staging And Production Cutover Runbook](../../docs/postgres-staging-rollout.md).
A live cutover is a separately scheduled operational action with a named operator and an approved change window.

This runbook covers the core-record backfill only.
[PostgreSQL Retained-Data Migration Contracts](postgres-retained-data-contracts.md) supersedes this runbook's statements that calendar, OTP, friend-request, and daily-log collections stay in MongoDB, and governs the retained-data backfill and the final MongoDB removal.

## Tooling

The migration lives at `server/scripts/20260910_mongo_events_to_postgres/` and runs as a dated one-off command.
It never writes to MongoDB.

```sh
go run ./scripts/20260910_mongo_events_to_postgres \
  --mongo-uri "$MONGODB_URI" \
  --mongo-database "$MONGODB_DATABASE" \
  --postgres-uri "$POSTGRES_APPLICATION_URI"
```

Flags:

| Flag                  | Purpose                                                                                                     |
| --------------------- | ----------------------------------------------------------------------------------------------------------- |
| `--apply`             | Writes PostgreSQL rows. Without it the command is a read-only preflight.                                    |
| `--batch-size N`      | Source documents read per MongoDB page.                                                                     |
| `--limit N`           | Stops after `N` committed event units; folders and reconciliation wait for a run that completes all events. |
| `--batch-label label` | Stamps every ledger and quarantine row with the run identity. Defaults to a UTC timestamp.                  |

The account backfill `server/scripts/20260910_mongo_accounts_to_postgres/` runs first so account ownership and response references resolve through `platform_identities.external_user_id`.

## Migration Units And Ordering

A migration unit is one aggregate copied atomically in a single PostgreSQL transaction.
Units are migrated in dependency order: accounts, then events with their responses, signup data, and attendees, then folders with their memberships.
The event tool owns the event, response, attendee, signup, folder, and membership units; the account tool owns account units; the retained calendar backfill owns calendar integration units.
OTP, friend-request, and daily-log collections are never copied by the event tool.

Each unit stamps its completion in `migration_ledger` inside the same transaction, including the fresh PostgreSQL target identity.
Folders rewrite each membership's legacy event reference to the migrated `postgres_events.id` recorded in the ledger, so migrated relationships never depend on a permanent legacy-identifier map.

## Resumability And Interruption Recovery

The run cursor advances only after a unit transaction commits.
A completed unit is skipped on the next run, and an incomplete unit is retried from scratch because its transaction rolled back.
Recovery is replay, not manual repair: rerun the same command with the same or a new `--batch-label`.
Replaying a completed run reports every unit as skipped and writes nothing.
The `--limit` flag simulates interruption and proves the resume point; the isolated rehearsal covers both interruption and replay.

## Source-Write Coordination

The backfill runs against the live MongoDB source without a maintenance window:

1. The initial pass copies units in batches while the application keeps serving MongoDB.
2. A bounded catch-up pass re-reads source records changed during the initial pass.
   Legacy event documents carry only an insertion timestamp, so catch-up re-runs the migration for the affected units rather than relying on a change cursor.
3. After catch-up, the affected record kind enters a short write freeze in which new mutations are rejected or queued and never dual-written.
4. During the freeze the migration performs one final idempotent pass and reconciliation.
5. Only after reconciliation passes are reads and writes switched for that record kind to PostgreSQL, ending the freeze.

No permanent dual write exists; a record is authored in exactly one store after its cutover.
Because the catch-up and freeze windows are operational, the tooling records the batch label and reconciliation evidence so the operator can evidence which pass produced the cutover state.

## Reconciliation Evidence

`--apply` runs reconciliation after a complete event and folder pass and prints a report.
It fails the run on any mismatch and never repairs.
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
The isolated rehearsal `TestBackupRestoreRehearsal` in `server/scripts/20260910_mongo_events_to_postgres/backup_restore_integration_test.go` seeds the representative fixture set, runs the event and folder migration, dumps the test database through the backup role, restores it into a fresh scratch database through the bootstrap role, and reconciles the restored table list plus every representative migrated relation by row count and full-row digest.
A reconciliation mismatch fails the rehearsal.

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

- A complete event and folder pass with the intended `--batch-label`.
- A replay that reports every unit skipped.
- A reconciliation report with no mismatches.
- A reviewed quarantine report with every unresolved record either accepted or separately resolved.

## Retention And Cleanup

The schema is additive and never dropped during migration.
MongoDB source documents are retained unmodified through the retention window and final validation.
Runtime reads and writes are PostgreSQL-only for accounts, events, responses, attendees, signup data, groups, folders, calendar integrations, OTP challenges, and new daily user logs; retained MongoDB `users` documents are recovery source that is never read or written at runtime.
After cutover validation and the retention window, drop `migration_ledger` and `migration_quarantine`; they are operational tooling and are never read by request paths.
MongoDB runtime removal is a separate stage that starts only after the core-record cutover in TASK-0190.08 and the retained-data cutover both pass.
The [Remaining MongoDB Collections And Deferred Scope](#remaining-mongodb-collections-and-deferred-scope) section records what still exists at that point and which task removes it.

## Isolated Rehearsal

The rehearsal lives in `server/scripts/20260910_mongo_events_to_postgres/main_integration_test.go` and runs inside the isolated Compose test stack.
It seeds the contract's representative fixture set, including one record per quarantine path:

- An anonymous timed event with protected and open guest responses.
- An anonymous dates-only event with responses.
- A day-of-week event.
- An authenticated timed event with an owner account and an account response.
- An availability group with active and declined attendees and a calendar-derived response.
- A signup form with multiple blocks, capacity limits, and account and guest signups.
- A folder with owned member events.
- A retained calendar-integration account with provider tokens and preferences.
- Quarantine records for every reason code.

The rehearsal asserts target counts, ownership and response access, nullable and schedule fields, the quarantine breakdown, that the MongoDB source is unchanged, a clean reconciliation, idempotent replay, and resume after an interrupted `--limit` run.

Run it with:

```sh
docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test \
  go test ./scripts/20260910_mongo_events_to_postgres/ -count=1
```

## Retained Calendar Backfill

The retained calendar backfill lives at `server/scripts/20260911_mongo_calendars_to_postgres/` and runs as a dated one-off command after the account backfill.
It never writes to MongoDB.

```sh
go run ./scripts/20260911_mongo_calendars_to_postgres \
  --mongo-uri "$MONGODB_URI" \
  --mongo-database "$MONGODB_DATABASE" \
  --postgres-uri "$POSTGRES_APPLICATION_URI"
```

Flags:

| Flag                  | Purpose                                                                                    |
| --------------------- | ------------------------------------------------------------------------------------------ |
| `--apply`             | Writes PostgreSQL calendar rows. Without it the command is a read-only preflight.          |
| `--batch-size N`      | Retained user documents read per MongoDB page.                                             |
| `--limit N`           | Stops after `N` committed account units to simulate interruption.                          |
| `--batch-label label` | Stamps every ledger and quarantine row with the run identity. Defaults to a UTC timestamp. |

A migration unit is one retained owner account and covers every calendar connection, encrypted credential, sub-calendar, and preference it owns.
The unit resolves its owner through `platform_identities.external_user_id` and requires an existing `accounts` row; a retained document whose owner or account is absent is quarantined as `missing-owner-account` and is never migrated.
Each unit is written in one PostgreSQL transaction that also records completion in `migration_ledger` under kind `calendar-account`, keyed by the legacy `users._id`, with the resolved platform identity as its target.
The legacy AES-CFB Apple password is decrypted and re-encrypted with the versioned AES-256-GCM envelope, while OAuth2 tokens and the ICS feed URL are plaintext in the retained document and are encrypted on write.

`--apply` runs reconciliation after a complete pass and prints a report.
It fails the run on any mismatch and never repairs.
Checks:

- Unit and per-kind target counts for connections, sub-calendars, credential rows, and preferences.
- Every migrated credential is stored in the `v1:` GCM envelope.
- Quarantine totals grouped by reason code.

Run the isolated rehearsal, which covers Google, Outlook, Apple, ICS, preferences, a two-connection legacy key suffix, missing-owner quarantine, replay, and resume after interruption, with:

```sh
docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test \
  go test ./scripts/20260911_mongo_calendars_to_postgres/ -count=1
```

## Final Isolated Cutover Rehearsal

TASK-0190.08 rehearses the whole cutover in the isolated test stack and records the evidence in the [PostgreSQL Staging And Production Cutover Runbook](../../docs/postgres-staging-rollout.md).

- Backend: `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test go test ./... -count=1`.
  This includes the account backfill, event backfill, calendar backfill, analytics, and backup and restore rehearsals.
- Browser: from `e2e/`, run `E2E_FRONTEND=bundled npm run test:e2e -- --project=firefox-desktop`, then `npm run test:e2e -- --project=chromium-desktop --project=chromium-mobile --project=firefox-touch`, then `npm run test:e2e -- --project=chromium-production-desktop --project=chromium-production-mobile`.
- The coverage map for account sign-in, every event kind, groups, signup forms, folders, event links, identity and owner authority, calendar integration, OTP challenges, friend-request retirement, historical daily logs, and backup and restore lives in the staging and production runbook.

A rehearsal failure blocks cutover.

## Remaining MongoDB Collections And Deferred Scope

PostgreSQL is authoritative after cutover for accounts, events, responses, attendees, signup data, groups, folders, calendar integrations, OTP challenges, and new daily user logs.
The MongoDB collections below remain only as recovery source or for legacy-only behavior; none is a second authority for a migrated record kind.

| Collection                                                         | Content at final cutover                                                                   | Runtime use                                                                                                                                     | Deferred scope                                                                 |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| `events`, `eventResponses`, `attendees`, `folders`, `folderEvents` | Legacy and noncanonical event, response, attendee, folder, and membership documents        | Served only by legacy routes for payload shapes the PostgreSQL creation classifier does not claim; supported records are authored in PostgreSQL | Removed with the MongoDB runtime by TASK-0199.09 after both cutovers pass.     |
| `users`                                                            | Retained account, calendar, and preference documents from before the retained-data cutover | Never read or written at runtime; read only by one-off backfill scripts                                                                         | Recovery source until the retention window ends, then removed by TASK-0199.09. |
| `dailyuserlogs`                                                    | Historical daily user logs written before the daily-log cutover                            | Never read or written at runtime; new logs go to PostgreSQL                                                                                     | Backfilled by TASK-0199.10, then recovery source until removal.                |
| `otpCodes`                                                         | Challenges from before the OTP cutover                                                     | Never read or written                                                                                                                           | They expire and are removed with MongoDB by TASK-0199.09.                      |
| `friendrequests`                                                   | Dormant request documents                                                                  | Never read or written                                                                                                                           | Retired by TASK-0199.08 and removed with MongoDB by TASK-0199.09.              |

Integration-only user fields have PostgreSQL destinations under the retained-data contracts: calendar connections, sub-calendars, and preferences move to `calendar_accounts`, `calendar_sub_calendars`, `calendar_account_credentials`, and `calendar_preferences`; OTP challenges move to `otp_challenges`; daily-log membership moves to `daily_user_logs` and `daily_user_log_members`; event creator attribution stays in `postgres_events.creator_posthog_id`; and active-user and signed-up-user reporting read `daily_user_logs` and `accounts`.
No integration-only user field remains authoritative in the retained `users` document, and copying a retained document back into runtime authority is never a recovery option.
