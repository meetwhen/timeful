---
id: TASK-0094
title: Enforce owner-only Event Occurrence Span saving (FR-012)
status: To Do
assignee: []
created_date: '2026-08-28 20:35'
labels:
  - backend
  - authorization
dependencies: []
references:
  - server/routes/events.go
  - server/routes/postgres_event_routes.go
  - >-
    backlog/tasks/task-0092 -
    Rewrite-the-glossary-as-DDD-flavored-v2-and-align-requirements-terminology.md
documentation:
  - docs/requirements/functional/fr/FR-012.md
  - docs/requirements/functional/fr/FR-018.md
priority: medium
type: bug
ordinal: 100000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
FR-012 (status: accepted; reworded in TASK-0092) restricts saving, replacing, or clearing the optional Event Occurrence Span to the Event Owner: "Other Event Visitors shall not be able to change it." The application does not enforce this today.

# Verified current behavior (2026-08-28, TASK-0092 VERIFY step)
- MongoDB handlers saveTimefulSchedule and clearTimefulSchedule perform no authorization at all; a source comment states scheduling is intentionally available to anyone with the event link (server/routes/events.go:109-164).
- The PostgreSQL handler postgresUpdateSchedule also performs no owner authorization (server/routes/postgres_event_routes.go:336-377).

# Scope
- Enforce owner-only Event Occurrence Span mutation for PostgreSQL events per FR-012, using the existing credential model (Event Owner Edit Token or the Platform Visitor Identity associated with the event ownership, as FR-018/FR-063 describe).
- Decide and record whether legacy MongoDB schedule endpoints stay behaviorally unchanged (MongoDB event retirement is a separate proposed task; see TASK-0092's follow-up candidates) or gain owner authority.
- Add route tests proving the owner-success and non-owner-rejection paths.

# Constraints
- Keep MongoDB access in server/db/ and PostgreSQL access in server/postgres/; route handlers must not query the stores directly (AGENTS.md backend conventions).
- Run Mongo-backed route tests via the isolated Compose overlay: docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml up -d mongo-test, then run --rm server-route-test; the test database must be timeful-test or carry a timeful-test- prefix (AGENTS.md Server Test Workflow).
- Add Swag annotations if route signatures or responses change, then regenerate swagger and the frontend API types per AGENTS.md.
- This bug was discovered during TASK-0092's VERIFY step; that documentation task deliberately made no code change. FR-012's accepted status means the intended behavior is already the requirement; this task closes the enforcement gap.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Saving, replacing, or clearing an Event Occurrence Span on a PostgreSQL event succeeds only for the Event Owner and is rejected for every other Event Visitor, matching FR-012's owner-only authority (owner authorization follows the FR-018 credential model: the Event Owner Edit Token or the associated Platform Visitor Identity).
- [ ] #2 Legacy MongoDB schedule endpoints are either left behaviorally unchanged or brought under owner authority, with the choice recorded in the task notes and kept consistent with the FR-062/FR-081/FR-082 MongoDB-preservation clauses.
- [ ] #3 Route tests cover both the owner-success and non-owner-rejection paths for the span endpoints and pass via the isolated Compose test stack (mongo-test, postgres-test, server-test per compose.test.yaml).
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
