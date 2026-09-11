# PostgreSQL Operations Runbook

PostgreSQL is the only supported database.
This runbook documents the ongoing operational procedures: backup and restore, and retention cleanup.

## Backup And Restore

Backups use the PostgreSQL 18 client tools with a custom-format `pg_dump` and the least-privilege backup role.
The backup role needs read access to every table, granted at creation and re-applied to existing databases with:

```sql
GRANT pg_read_all_data TO <backup role> WITH INHERIT TRUE;
```

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
The destructive restore over the live database adds `--clean --if-exists` and requires a change ticket.
A backup is verified only after its restore reconciliation reports matching counts and digests.
Off-host replication, automated scheduling, and recovery objectives remain later operational work.

## Retention Cleanup

The schema is additive and is never dropped to roll back a release.
`migration_ledger` and `migration_quarantine` are migration-tooling tables that no request path reads and that hold only the last executed run's records.
After the retention window and cutover validation complete, drop both tables to finish the cleanup.
Until the cleanup runs, keep the verified backup and the retention window intact.
