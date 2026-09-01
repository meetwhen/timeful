---
id: TASK-0129
title: >-
  Enforce Event Owner powers on PostgreSQL events (Event Owner Edit Token,
  FR-018/FR-115/FR-116)
status: To Do
assignee: []
created_date: '2026-09-01 19:03'
labels:
  - postgresql
  - identity
  - backend
  - authorization
dependencies:
  - TASK-0071
references:
  - docs/requirements/functional/fr/FR-018.md
  - docs/requirements/functional/fr/FR-063.md
  - docs/requirements/functional/fr/FR-083.md
  - docs/requirements/functional/fr/FR-115.md
  - docs/requirements/functional/fr/FR-116.md
  - docs/requirements/quality/qr/QR-003.md
documentation:
  - docs/design/architecture/adr/ADR-010.md
  - docs/terminology/glossary.md
  - server/docs/postgres-anonymous-event-compatibility.md
modified_files:
  - server/routes/postgres_event_routes.go
  - server/postgres/repository.go
  - server/errs/errors.go
priority: high
type: feature
ordinal: 142300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Follow-up to TASK-0071, deferred by owner decision on 2026-08-31. Enforce Event Owner powers on PostgreSQL events: issue and validate the Event Owner Edit Token, authorize Event Settings edits per FR-018 (token, associated Platform Visitor Identity, owner-issued Granted EVCC per FR-083; base EVCC never), and enforce owner-only event archive (FR-115) and deletion (FR-116). Ownership association and takeover semantics follow FR-063. Requires Granted EVCC issuance from TASK-0071.02 for the third authorizer; coordinate ordering accordingly. MongoDB owner authorization remains legacy and unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 PostgreSQL events issue an Event Owner Edit Token to the creating owner as a PostgreSQL-only credential per FR-018 and the glossary
- [ ] #2 Event Settings edits on PostgreSQL events authorize through the Event Owner Edit Token, the associated Platform Visitor Identity, or the owner-issued Granted Event Visitor Control Credential, while the base EVCC never authorizes settings edits (FR-018)
- [ ] #3 Event archive/unarchive (FR-115) and event deletion (FR-116) on PostgreSQL events are authorized through the same credential model FR-018 defines
- [ ] #4 Proving a valid Event Owner Edit Token with a different Platform Visitor Identity moves event ownership to the proving identity and the previously associated identity loses Event Settings authority (FR-063)
- [ ] #5 MongoDB event-settings authorization remains legacy and unchanged
- [ ] #6 Route and frontend regression coverage verifies the authorization model, including rejection cases for base-EVCC and anonymous attempts
- [ ] #7 Swagger annotations regenerated and npm run gen:api run from frontend/ where annotations changed
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
