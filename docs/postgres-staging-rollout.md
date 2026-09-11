# PostgreSQL Staging And Production Cutover Runbook

## Scope

This runbook is the operator procedure for making PostgreSQL the authoritative store for accounts, events, responses, groups, signup forms, folders, calendar integrations, OTP challenges, and daily user logs while validating the retained MongoDB integrations.
It implements the tooling described in the [PostgreSQL Core Migration Runbook](../server/docs/postgres-core-migration-runbook.md) and the store boundaries in the [PostgreSQL Retained-Data Migration Contracts](../server/docs/postgres-retained-data-contracts.md).
It records prerequisites, source-write coordination, validation gates, monitoring, rollback limits, and the handling of failed or quarantined records.
A live cutover is a separately scheduled operational action with a named operator, an approved change window, and the evidence gates below.
The isolated rehearsal that validated this procedure is recorded in the [Final Isolated Cutover Rehearsal](#final-isolated-cutover-rehearsal) section.

## Prerequisites

- The preceding migration stages are complete and verified: signed-in and anonymous creation route to PostgreSQL, the account, event, calendar, and daily-log backfills reconcile in rehearsal, and the retained-data cutover of calendar integrations, OTP challenges, daily logs, and analytics has passed.
- The release under cutover passed backend CI, frontend checks, and the browser rehearsal suites in the [Final Isolated Cutover Rehearsal](#final-isolated-cutover-rehearsal) section.
- Every required PostgreSQL role credential and URI is populated in the selected environment file, including `POSTGRES_BACKUP_USERNAME` and `POSTGRES_BACKUP_PASSWORD`.
- The backup role can read every table: run the one-time bootstrap grant `GRANT pg_read_all_data TO <backup role> WITH INHERIT TRUE;` for databases created before this runbook, or confirm the container bootstrap already applied it.
- A fresh custom-format backup exists and is verified by the restore reconciliation described in [Backup And Restore](#backup-and-restore).
- A named operator, a change ticket, an open retention window, and access to logs and readiness endpoints are available.
- Historical `dailyuserlogs` backfill (TASK-0199.10) and MongoDB runtime removal (TASK-0199.09) are scheduled separately after this cutover passes.

## Deploy

```sh
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml config --quiet
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml up -d --build
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml ps
curl -fsS https://staging.timeful.fun/api/health/live
curl -fsS https://staging.timeful.fun/api/health
```

Confirm `postgres-migrate` completed successfully, legacy MongoDB reads and writes for noncanonical records remain intact, and both readiness endpoints return successfully.

## Source-Write Coordination

The backfill runs against the live MongoDB source before the freeze:

1. The initial pass copies migration units in batches while the application keeps serving.
2. A bounded catch-up pass re-reads records changed during the initial pass, because legacy event documents carry only an insertion timestamp.
3. After catch-up, the affected record kind enters a short write freeze in which new mutations are rejected or queued and never written to both stores.
4. During the freeze the migration performs one final idempotent pass and reconciliation.
5. Only after reconciliation passes do reads and writes for that record kind switch to PostgreSQL, ending the freeze.

No permanent dual write exists, and no record is authored in two stores after its cutover.
Record the batch label and reconciliation evidence so the operator can evidence which pass produced the cutover state.

## Validation Gates

Each record kind must pass every gate before the write freeze ends:

- Preflight: run the migration command without `--apply` and confirm no validation errors.
- Complete pass: run the event and folder pass with the intended `--batch-label`; folders wait for a run that completes all events.
- Idempotent replay: rerun the command and confirm every unit reports as skipped and no rows are written.
- Reconciliation: the printed report must have no mismatches, including per-kind counts, `num_responses` equality, referential integrity, and quarantine totals.
- Quarantine review: every unresolved record is either accepted with a recorded reason or separately resolved.
- Backup: a fresh custom-format backup exists and its restore reconciliation matches the source.
- Application smoke: anonymous and signed-in event creation, guest and account responses, folder membership, calendar reads, and OTP sign-in all work from the deployed application.

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
The isolated rehearsal `TestBackupRestoreRehearsal` in `server/scripts/20260910_mongo_events_to_postgres/backup_restore_integration_test.go` exercises the same dump and restore cycle against representative migrated records and reconciles eleven relations by row count and full-row digest.

## Monitoring

Watch the following during the backfill, freeze, and post-cutover window:

- Readiness: `/api/health/live` and `/api/health`, including PostgreSQL and MongoDB health.
- Errors: route error rates, PostgreSQL connection errors, transaction failures, and uniqueness conflicts.
- Capacity: PostgreSQL storage growth, active connections, and migration batch duration.
- Progress: `migration_ledger` completed units and `migration_quarantine` totals for the active batch label.
- Backups: backup completion, exit status, and age; a failed or stale backup blocks cutover.
- Retained integrations: calendar read errors, OTP send or verify failures, and daily-log writes after their cutover.

## Rollback Limits

- Before a record kind's write freeze, rollback is trivial: MongoDB remains unmodified and authoritative, and disabling the PostgreSQL route for that kind returns to it.
- The write freeze is the practical point of no return, because mutations accepted by PostgreSQL are not dual-written and a MongoDB rollback would lose post-cutover writes unless they are exported back.
- Schema changes are additive and destructive downgrades are refused during the migration.
- Supported event creation no longer has a flag-based rollback; reverting to MongoDB creation requires the migration rollback boundaries, not an environment variable.
- MongoDB source documents are retained unmodified through the retention window and final validation.
- MongoDB runtime removal is irreversible and requires a completed core-record cutover, a completed retained-data cutover, an unexpired retention window, and verified backups.

## Final Isolated Cutover Rehearsal

Run the rehearsal in the isolated test stack; it must never target a development or production database.

```sh
# Backend, migration, analytics, and backup/restore rehearsals.
docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test

# Browser rehearsals.
cd e2e
E2E_FRONTEND=bundled npm run test:e2e -- --project=firefox-desktop
npm run test:e2e -- --project=chromium-desktop --project=chromium-mobile --project=firefox-touch
npm run test:e2e -- --project=chromium-production-desktop --project=chromium-production-mobile
```

The rehearsal covers the cutover surface as follows:

| Rehearsal area               | Evidence                                                                                                                                                                                                   |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Account sign-in              | `account-deletion.spec.ts`, `e2e/helpers/account-auth.ts`, `server/routes/auth_otp_test.go`, `server/postgres/otp_test.go`                                                                                 |
| Every event kind             | `timed-event-create-firefox.spec.ts`, `timed-event-postgres-dashboard-folders-firefox.spec.ts`, `events_timed_payloads_test.go`, `postgres_signed_in_event_test.go`                                        |
| Availability group           | `timed-event-group-postgres-firefox.spec.ts`, `server/routes/postgres_group_test.go`                                                                                                                       |
| Signup form                  | `timed-event-signup-postgres-firefox.spec.ts`, `server/routes/postgres_signup_test.go`                                                                                                                     |
| Folders                      | `timed-event-postgres-dashboard-folders-firefox.spec.ts`, `server/routes/postgres_folders_test.go`                                                                                                         |
| Existing links               | The event backfill rehearsal rewrites folder membership and documents that pre-cutover public URLs may change; browser specs open `/e/<shortId>` links                                                     |
| Identity and owner authority | `timed-event-owner-authority-firefox.spec.ts`, `timed-event-access-transfer-firefox.spec.ts`, `server/routes/postgres_owner_test.go`, `server/routes/postgres_transfers_test.go`                           |
| Calendar integration         | `timed-event-calendar-integration-firefox.spec.ts`, `server/scripts/20260911_mongo_calendars_to_postgres`, `server/postgres/calendar_test.go`                                                              |
| OTP challenge                | `server/routes/auth_otp_test.go`, `server/postgres/otp_test.go`, OTP sign-in through `e2e/helpers/account-auth.ts`                                                                                         |
| Friend-request retirement    | TASK-0199.08 removed the collection, model, accessors, and runtime references; no rehearsal re-enables them                                                                                                |
| Historical daily logs        | `server/postgres/dailylogs_test.go` covers PostgreSQL-owned writes and reads; `server/scripts/20260911_mongo_dailyuserlogs_to_postgres` backfills and reconciles retained history under isolated rehearsal |
| Event analytics              | `server/postgres/analytics_test.go` proves migrated and new PostgreSQL events are each counted once; TASK-0199.06 deleted the MongoDB `server/db/analytics.go` path                                        |
| Backup and restore           | `server/scripts/20260910_mongo_events_to_postgres/backup_restore_integration_test.go` reconciles the restored schema and migrated relations                                                                |

The 2026-09-11 rehearsal recorded all backend packages passing, including backup and restore reconciliation of 11 representative relations with matching full-row digests; Firefox desktop with 56 passed, 1 skipped, and 0 failed; Chromium desktop, Chromium mobile, and Firefox touch with 55 passed, 20 skipped, and 0 failed; and the production desktop and mobile projects with 3 passed and 0 failed.
Record the run identifiers, pass and skip counts, and any accepted skip in the change ticket.

## Smoke Test

Supported new events are always created in PostgreSQL; no creation flag is required.
Create one anonymous [Timed Event](terminology/glossary.md#timed-event) poll and one [Dates-Only Event](terminology/glossary.md#dates-only-event) poll.
Confirm each has one bare eight-character Crockford ID and exercise guest response mutation, selected schedule save and clear, plugin `set-slots` and `get-slots`, and a signed-in response.
Confirm PostgreSQL account responses remain absent from the dashboard and `/api/user/events`.
Open a [Calendar Connection](terminology/glossary.md#calendar-connection) page and confirm reads still resolve, and request an OTP code for a test account to confirm sign-in delivery.

## Rollback

Supported creation no longer has a flag-based rollback.
Reverting to MongoDB creation requires the migration rollback boundaries in the [PostgreSQL Core Migration Runbook](../server/docs/postgres-core-migration-runbook.md), not an environment variable.
Do not roll back the additive SQL schema.
Routine PostgreSQL backup and restore are exercised by the isolated rehearsal and documented in [Backup And Restore](#backup-and-restore); off-host replication, recovery objectives, and automated scheduling remain later operational work.
