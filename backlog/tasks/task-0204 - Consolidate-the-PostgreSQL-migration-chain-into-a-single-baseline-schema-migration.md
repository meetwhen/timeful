---
id: TASK-0204
title: >-
  Consolidate the PostgreSQL migration chain into a single baseline schema
  migration
status: To Do
assignee: []
created_date: '2026-09-11 20:25'
labels: []
dependencies: []
references:
  - >-
    backlog/backlog.md open item: remove server/migrations from mongo,
    consolidate into a single init script for postgres
  - server/migrations/20260910170000_migration_ledger.sql
  - server/migrations/20260911120000_drop_folder_legacy_event_id.sql
documentation:
  - docs/postgres-operations.md
  - docs/environments.md
  - server/docs/postgres-data-boundaries.md
modified_files:
  - server/migrations
  - server/postgres
  - docs/postgres-operations.md
  - docs/environments.md
  - compose.yaml
  - compose.test.yaml
  - server/Dockerfile
priority: medium
type: chore
ordinal: 243000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
`server/migrations/` has accumulated 15 goose migrations (from anonymous-event compatibility through the folder legacy-column cleanup) that encode the MongoDB-to-PostgreSQL transition, Mongo backfill tooling, and intermediate schema revisions. MongoDB is fully retired, there is no Mongo data to migrate, and existing local, staging, and production PostgreSQL databases may be recreated, so no in-place upgrade path is required.

Replace the chain with one authoritative goose baseline migration that creates the complete schema the current server requires, so a fresh database initializes in one step and future schema changes continue as new incremental goose migrations on top.

The resulting schema must match the schema the current chain produces, except that the retired Mongo migration tooling tables (`migration_ledger`, `migration_quarantine`) are no longer created. Goose stays the migration tool, and the migrator image, Compose migration service, CI migration caches, and isolated test and E2E stacks must keep working. Historical ADRs and applied-migration records are not rewritten.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 server/migrations/ contains a single baseline goose migration that creates the complete current schema, and no other schema migration files remain.
- [ ] #2 Applying the baseline to an empty PostgreSQL 18 database succeeds and produces the same schema the current migration chain produces, with the only intended difference being that migration_ledger and migration_quarantine are not created (verified by a schema comparison, not by inspection alone).
- [ ] #3 Runtime compatibility columns and behavior still used by the current server are preserved in the baseline.
- [ ] #4 PostgreSQL-backed test harnesses initialize their schema from the baseline instead of individual migration filenames, and the isolated backend test suite passes.
- [ ] #5 A fresh development/test Compose stack reaches a healthy server using the baseline migration, and the Firefox browser E2E suite passes.
- [ ] #6 Documentation that describes the migration history, the forward-only compatibility statement, and the migration_ledger/migration_quarantine retention cleanup reflects the baseline model.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
