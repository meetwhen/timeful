# PostgreSQL Staging And Production Cutover Runbook

## Scope

This runbook is the operator procedure for making PostgreSQL the authoritative store for accounts, events, responses, groups, signup forms, folders, calendar integrations, OTP challenges, and daily user logs.
It implements the tooling described in the [PostgreSQL Core Migration Runbook](../server/docs/postgres-core-migration-runbook.md) and the store boundaries in the [PostgreSQL Retained-Data Migration Contracts](../server/docs/postgres-retained-data-contracts.md).
It records prerequisites, source-write coordination, validation gates, monitoring, rollback limits, and the handling of failed or quarantined records.
The one-off backfill commands no longer exist in the repository, so the source-write and backfill sections record the procedure the rehearsals executed rather than a repeatable command sequence.
A live cutover is a separately scheduled operational action with a named operator, an approved change window, and the evidence gates below.
The isolated rehearsal that validated this procedure is recorded in the [Final Isolated Cutover Rehearsal](#final-isolated-cutover-rehearsal) section.
TASK-0199.09 removed the MongoDB runtime, driver, migration tooling, Compose services, environment variables, and browser E2E support on 2026-09-11; the [Final MongoDB Collection Retirement](#final-mongodb-collection-retirement) section records the executed backfills and the irreversible rollback boundary.

## Prerequisites

- The preceding migration stages are complete and verified: signed-in and anonymous creation route to PostgreSQL, the account, event, calendar, and daily-log backfills reconcile in rehearsal, and the retained-data cutover of calendar integrations, OTP challenges, daily logs, and analytics has passed.
- The release under cutover passed backend CI, frontend checks, and the browser rehearsal suites in the [Final Isolated Cutover Rehearsal](#final-isolated-cutover-rehearsal) section.
- Every required PostgreSQL role credential and URI is populated in the selected environment file, including `POSTGRES_BACKUP_USERNAME` and `POSTGRES_BACKUP_PASSWORD`.
- The backup role can read every table: run the one-time bootstrap grant `GRANT pg_read_all_data TO <backup role> WITH INHERIT TRUE;` for databases created before this runbook, or confirm the container bootstrap already applied it.
- A fresh custom-format backup exists and is verified by the restore reconciliation described in [Backup And Restore](#backup-and-restore).
- A named operator, a change ticket, an open retention window, and access to logs and readiness endpoints are available.
- Historical `dailyuserlogs` backfill (TASK-0199.10) and MongoDB runtime removal (TASK-0199.09) are complete; see [Final MongoDB Collection Retirement](#final-mongodb-collection-retirement).

## Deploy

```sh
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml config --quiet
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml up -d --build
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml ps
curl -fsS https://staging.timeful.fun/api/health/live
curl -fsS https://staging.timeful.fun/api/health
```

Confirm `postgres-migrate` completed successfully and both readiness endpoints return successfully.

## Source-Write Coordination

The executed backfill ran against the live MongoDB source before the freeze:

1. The initial pass copied migration units in batches while the application kept serving.
2. A bounded catch-up pass re-read records changed during the initial pass, because legacy event documents carry only an insertion timestamp.
3. After catch-up, the affected record kind entered a short write freeze in which new mutations were rejected or queued and never written to both stores.
4. During the freeze the migration performed one final idempotent pass and reconciliation.
5. Only after reconciliation passed did reads and writes for that record kind switch to PostgreSQL, ending the freeze.

No permanent dual write existed, and no record is authored in two stores after its cutover.
The batch label and reconciliation evidence were recorded so the operator could evidence which pass produced the cutover state.

## Validation Gates

Each record kind had to pass every gate before the write freeze ended:

- Preflight: run the migration command without `--apply` and confirm no validation errors.
- Complete pass: run the event and folder pass with the intended `--batch-label`; folders wait for a run that completes all events.
- Idempotent replay: rerun the command and confirm every unit reports as skipped and no rows are written.
- Reconciliation: the printed report must have no mismatches, including per-kind counts, `num_responses` equality, referential integrity, and quarantine totals.
- Quarantine review: every unresolved record is either accepted with a recorded reason or separately resolved.
- Backup: a fresh custom-format backup exists and its restore reconciliation matches the source.
- Application smoke: anonymous and signed-in event creation, guest and account responses, folder membership, calendar reads, and OTP sign-in all work from the deployed application.

The isolated rehearsal exercised every gate and recorded the results in the [Final Isolated Cutover Rehearsal](#final-isolated-cutover-rehearsal) section.

## Failed And Quarantined Records

Ambiguous or corrupt records are reported in `migration_quarantine` and are never repaired by inference.
The core runbook lists each reason code and its handling.
At cutover, group the quarantine totals by reason code and record one of two decisions per group:

- Accept the quarantine outcome, such as an ownerless event or an unimported response with no identity, and record why the outcome is safe.
- Resolve the record separately before cutover, then rerun the idempotent pass so the ledger and quarantine report reflect the resolution.

The quarantine ledger is append-only, so replayed runs preserve the same decisions.
A failed migration step stops the cutover; do not proceed to the freeze while a reconciliation mismatch or an unreviewed quarantine group remains.

## Backup And Restore

Backups use the PostgreSQL 18 client tools and the least-privilege backup role.
The commands run inside the `postgres` container so the role names and database name come from the container environment and local socket authentication applies; no host environment file is sourced and no password is interpolated.
Run the commands on the deployment host from a checkout of the release:

```sh
# Backup the selected environment.
docker compose --project-name timeful-production --env-file .env.production -f compose.yaml -f compose.production.yaml exec -T postgres \
  sh -ec 'pg_dump --format=custom --no-owner --username "$POSTGRES_BACKUP_USERNAME" --dbname "$POSTGRES_DB"' \
  > "timeful-$(date -u +%Y%m%dT%H%M%SZ).dump"

# Restore into a scratch maintenance database for verification.
docker compose --project-name timeful-production --env-file .env.production -f compose.yaml -f compose.production.yaml exec -T postgres \
  sh -ec 'createdb --username "$POSTGRES_USER" timeful-restore-check'
docker compose --project-name timeful-production --env-file .env.production -f compose.yaml -f compose.production.yaml exec -T postgres \
  sh -ec 'pg_restore --no-owner --exit-on-error --username "$POSTGRES_USER" --dbname timeful-restore-check' < "timeful-<timestamp>.dump"

# Reconcile per-table row counts and key content digests before declaring the backup verified.
```

Both restore commands run as the container's bootstrap superuser, because `pg_restore` creates and drops objects that the read-only backup role cannot.
The destructive restore over the live database adds `--clean --if-exists` and requires the change ticket.
The retired backup and restore rehearsal, removed with the migration tooling by TASK-0199.09.03, exercised the same dump and restore cycle against representative migrated records and reconciled eleven relations by row count and full-row digest on 2026-09-11.

## Monitoring

Watch the following during the backfill, freeze, and post-cutover window:

- Readiness: `/api/health/live` and `/api/health`, including PostgreSQL health.
- Errors: route error rates, PostgreSQL connection errors, transaction failures, and uniqueness conflicts.
- Capacity: PostgreSQL storage growth, active connections, and migration batch duration.
- Progress: `migration_ledger` completed units and `migration_quarantine` totals for the active batch label.
- Backups: backup completion, exit status, and age; a failed or stale backup blocks cutover.
- Retained integrations: calendar read errors, OTP send or verify failures, and daily-log writes after their cutover.

## Rollback Limits

- Before a record kind's write freeze, rollback was trivial: MongoDB remained unmodified and authoritative, and disabling the PostgreSQL route for that kind returned to it.
- The write freeze was the practical point of no return, because mutations accepted by PostgreSQL are not dual-written and a MongoDB rollback would have lost post-cutover writes unless they were exported back.
- Schema changes are additive and destructive downgrades are refused during the migration.
- Supported event creation has no flag-based rollback and no MongoDB route; the migration rollback boundaries only ever applied before each record kind's write freeze.
- MongoDB source documents were retained unmodified through the retention window and final validation, and repository code can no longer read them.
- The MongoDB runtime, driver, migration tooling, and configuration removal on 2026-09-11 is irreversible: the repository has no MongoDB read path, credential, or recovery command, and a verified PostgreSQL backup is the only repository-supported recovery artifact.

## Final Isolated Cutover Rehearsal

Run the rehearsal in the isolated test stack; it must never target a development or production database.

```sh
# Backend route, store, model, and analytics suites.
docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test

# Browser rehearsals.
cd e2e
E2E_FRONTEND=bundled npm run test:e2e -- --project=firefox-desktop
npm run test:e2e -- --project=chromium-desktop --project=chromium-mobile --project=firefox-touch
npm run test:e2e -- --project=chromium-production-desktop --project=chromium-production-mobile
```

The rehearsal covers the cutover surface as follows:

| Rehearsal area               | Evidence                                                                                                                                                                                      |
| ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Account sign-in              | `account-deletion.spec.ts`, `e2e/helpers/account-auth.ts`, `server/routes/auth_otp_test.go`, `server/postgres/otp_test.go`                                                                    |
| Every event kind             | `timed-event-create-firefox.spec.ts`, `timed-event-postgres-dashboard-folders-firefox.spec.ts`, `events_timed_payloads_test.go`, `postgres_signed_in_event_test.go`                           |
| Availability group           | `timed-event-group-postgres-firefox.spec.ts`, `server/routes/postgres_group_test.go`                                                                                                          |
| Signup form                  | `timed-event-signup-postgres-firefox.spec.ts`, `server/routes/postgres_signup_test.go`                                                                                                        |
| Folders                      | `timed-event-postgres-dashboard-folders-firefox.spec.ts`, `server/routes/postgres_folders_test.go`                                                                                            |
| Existing links               | The event backfill rehearsal rewrites folder membership and documents that pre-cutover public URLs may change; browser specs open `/e/<shortId>` links                                        |
| Identity and owner authority | `timed-event-owner-authority-firefox.spec.ts`, `timed-event-access-transfer-firefox.spec.ts`, `server/routes/postgres_owner_test.go`, `server/routes/postgres_transfers_test.go`              |
| Calendar integration         | `timed-event-calendar-integration-firefox.spec.ts`, the retired calendar backfill rehearsal (TASK-0199.09.03), `server/postgres/calendar_test.go`                                             |
| OTP challenge                | `server/routes/auth_otp_test.go`, `server/postgres/otp_test.go`, OTP sign-in through `e2e/helpers/account-auth.ts`                                                                            |
| Friend-request retirement    | TASK-0199.08 removed the collection, model, accessors, and runtime references; no rehearsal re-enables them                                                                                   |
| Historical daily logs        | `server/postgres/dailylogs_test.go` covers PostgreSQL-owned writes and reads; the retired daily-log backfill rehearsal (TASK-0199.09.03) reconciled retained history under isolated rehearsal |
| Event analytics              | `server/postgres/analytics_test.go` proves migrated and new PostgreSQL events are each counted once; TASK-0199.06 deleted the MongoDB `server/db/analytics.go` path                           |
| Backup and restore           | The retired backup and restore rehearsal (TASK-0199.09.03) reconciled the restored schema and migrated relations                                                                              |

The 2026-09-11 rehearsal recorded all backend packages passing, including backup and restore reconciliation of 11 representative relations with matching full-row digests; Firefox desktop with 56 passed, 1 skipped, and 0 failed; Chromium desktop, Chromium mobile, and Firefox touch with 55 passed, 20 skipped, and 0 failed; and the production desktop and mobile projects with 3 passed and 0 failed.
Record the run identifiers, pass and skip counts, and any accepted skip in the change ticket.

## Final MongoDB Collection Retirement

TASK-0199.09 executed the repository-side MongoDB retirement on 2026-09-11, after the core-record cutover rehearsal in TASK-0190.08, the retained-data cutover, and the historical daily-log backfill in TASK-0199.10 passed.

The executed backfills were the account, anonymous-event, event and folder, retained calendar-integration, and historical daily-log backfills.
Their isolated rehearsals recorded clean reconciliation, idempotent replay, resume after interruption, source immutability, and re-encrypted provider credentials.

The repository removal stages were:

1. TASK-0199.09.01 removed the legacy MongoDB event runtime and the `server/db` package, including the Mongo collection variables, the Mongo health-check branch, and the historical one-off scripts that imported `server/db`.
2. TASK-0199.09.02 replaced the remaining mongo-driver types in the PostgreSQL runtime with canonical `models.ID` and `models.DateTime` types while preserving the JSON and stored-payload wire format.
3. TASK-0199.09.03 deleted the five mongo-to-postgres backfill commands, their representative fixtures, and the script-only `internal/legacybson` package, then `go mod tidy` removed `go.mongodb.org/mongo-driver` from `server/go.mod` and `server/go.sum`.
4. TASK-0199.09.04 removed the `mongo` and `mongo-test` Compose services, volumes, health checks, `MONGODB_*` environment variables, `mongo/` and `scripts/mongo/` bootstrap tooling, the Mongo CI path filters, and the E2E MongoDB inspection helper.

The three 2026 one-off maintenance commands `20260527_canonical_respondent_names`, `20260724_canonical_timed_events`, and `20260810_shortid_unique_index` were deleted with `server/db` before their production run or not-run status was confirmed; no repository procedure can rerun them.
The only non-Mongo one-off command that remains is `server/scripts/20240721_apple_calendar_test`.

The final collection retirement is irreversible:

- No repository code reads or writes MongoDB; there is no fallback read path, credential, or collection variable.
- Retained documents in a deployed MongoDB are no longer reachable by repository code and are not a recovery source; a verified PostgreSQL backup is the only repository-supported recovery artifact.
- MongoDB rollback is no longer possible: the pre-freeze routing rollback applied only before each record kind's write freeze and cannot be reconstructed from the repository.
- An operator can decommission a deployed MongoDB service and its volume only after confirming the PostgreSQL backup and retention requirements; this repository no longer performs or documents that removal as a runnable step.

## Smoke Test

Supported new events are always created in PostgreSQL; no creation flag is required.
Create one anonymous [Timed Event](terminology/glossary.md#timed-event) poll and one [Dates-Only Event](terminology/glossary.md#dates-only-event) poll.
Confirm each has one bare eight-character Crockford ID and exercise guest response mutation, selected schedule save and clear, plugin `set-slots` and `get-slots`, and a signed-in response.
Confirm PostgreSQL account responses remain absent from the dashboard and `/api/user/events`.
Open a [Calendar Connection](terminology/glossary.md#calendar-connection) page and confirm reads still resolve, and request an OTP code for a test account to confirm sign-in delivery.

## Rollback

Supported creation has no flag-based rollback and no MongoDB route.
The window for reverting a record kind to MongoDB closed at its write freeze, and the [Final MongoDB Collection Retirement](#final-mongodb-collection-retirement) removed the repository's MongoDB read path and credentials.
Do not roll back the additive SQL schema.
Routine PostgreSQL backup and restore are exercised by the isolated rehearsal and documented in [Backup And Restore](#backup-and-restore); off-host replication, recovery objectives, and automated scheduling remain later operational work.
