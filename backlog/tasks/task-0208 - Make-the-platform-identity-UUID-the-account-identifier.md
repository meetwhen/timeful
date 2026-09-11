---
id: TASK-0208
title: Make the platform identity UUID the account identifier
status: To Do
assignee: []
created_date: '2026-09-11 21:02'
labels: []
dependencies: []
references:
  - TASK-0177.01
documentation:
  - docs/design/architecture/adr/ADR-019.md
  - docs/requirements/functional/fr/FR-121.md
  - server/docs/postgres-data-boundaries.md
  - server/migrations/20260908160000_visitor_identities.sql
priority: medium
type: chore
ordinal: 247000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The account reference is currently a separate ObjectID-shaped 24-hex string held in platform_identities.external_user_id and denormalized into TEXT columns (account_user_id, postgres_events.owner_external_id, access_transfers.external_user_id, account_deletion_tombstones.external_user_id). Everything else in PostgreSQL already keys on native UUIDv7 (platform_identities.id, accounts.id, event_visitor_identities.platform_identity_id, postgres_events.owner_platform_identity_id). Consolidate account identity onto platform_identities.id so there is exactly one account identifier: UUIDv7, database-validated, and consistent with the rest of the schema. Outcome: sign-in sessions and API payloads carry the platform identity UUID; account-referencing tables use uuid foreign keys; the separate external identifier is removed with no compatibility lookup. The change is a wire-format break and requires a one-time data consolidation and forced re-sign-in for existing sessions. A new ADR supersedes ADR-019 and records the decision.

Constraints:
- Hard cutover: no request path, column, or payload accepts or emits the legacy 24-hex account identifier after the change.
- Native uuid storage for consolidated references; platform_identities.id keeps its uuidv7() default.
- account_deletion_tombstones must survive platform identity deletion, so it stores the uuid without a foreign key.
- Guest and unowned zero-identity semantics must be preserved.
- TASK-0177.01 (bson-objectid in Dashboard sorting and signup block local identities) overlaps this initiative and is reconciled rather than duplicated.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Every account reference in the runtime resolves through platform_identities.id, and no request path, stored column, or API payload accepts or emits the legacy 24-hex account identifier
- [ ] #2 The superseding ADR, backend boundary, schema consolidation, frontend adoption, and end-to-end verification subtasks are all complete
- [ ] #3 TASK-0177.01 has a recorded reconciliation (dependency or update) so its bson-objectid removal does not conflict with this identifier change
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
