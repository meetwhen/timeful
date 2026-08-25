---
id: TASK-0071
title: Implement PostgreSQL platform and event visitor identities
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-25 16:15'
updated_date: '2026-08-25 21:29'
labels:
  - postgresql
  - identity
  - event-visitor
  - platform-identity
dependencies: []
references:
  - docs/requirements/functional/fr/FR-073.md
  - docs/requirements/functional/fr/FR-079.md
  - docs/requirements/functional/fr/FR-001.md
  - docs/requirements/functional/fr/FR-060.md
  - docs/requirements/functional/fr/FR-062.md
  - backlog/handoffs/handoff-2026-08-25T16-12-59Z.md
  - docs/requirements/functional/fr/FR-081.md
  - docs/requirements/functional/fr/FR-082.md
  - docs/requirements/quality/qr/QR-014.md
  - docs/requirements/functional/fr/FR-083.md
  - docs/requirements/functional/fr/FR-084.md
  - docs/requirements/quality/qr/QR-006.md
  - backlog/handoffs/handoff-2026-08-25T21-03-31Z.md
documentation:
  - docs/terminology/glossary.md
  - server/docs/postgres-anonymous-event-compatibility.md
  - docs/design/architecture/adr/ADR-009.md
  - docs/design/architecture/adr/ADR-010.md
priority: high
type: feature
ordinal: 73000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Implement the first PostgreSQL-only Platform Identity and Event Visitor Identity foundation. New PostgreSQL events establish opaque event-scoped visitor identities, private Event Visitor Control Credentials (EVCCs), and response ownership through Event Visitor Identities; authenticated visitors may associate those identities with a private Platform Identity. PostgreSQL responses use opaque response IDs and an explicit selected-response contract that supports multiple responses per visitor and blind availability. Preserve MongoDB persistence, credentials, request behavior, and data isolation without migration, mutation, or lookup.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 PostgreSQL stores Platform Identities and event-scoped Event Visitor Identities with internal UUID relations
- [ ] #2 A PostgreSQL Event Visitor Identity is opaque browser-local and retained independently from EVCCs and Platform Identity values
- [ ] #3 Authenticated PostgreSQL visitors associate Event Visitor Identities with private PostgreSQL Platform Identities without MongoDB queries or migration
- [ ] #4 PostgreSQL issues and validates private EVCC authority for every response owned by an Event Visitor Identity without exposing credential values to application JavaScript
- [ ] #5 PostgreSQL availability responses are owned by Event Visitor Identity relations and do not expose raw authentication subjects as response identifiers
- [ ] #6 One Event Visitor Identity can own multiple availability responses for an event and database integrity rejects an event and owner Event Visitor Identity mismatch
- [ ] #7 PostgreSQL response APIs use opaque response IDs and an explicit selection creation edit and deletion contract that supports multiple visitor-owned responses
- [ ] #8 PostgreSQL blind-availability reads expose a non-owner only authorized responses and no other-response counts while an owner can view all responses
- [ ] #9 MongoDB event and user persistence request behavior and guest credentials remain unchanged
- [ ] #10 Cross-device transfers matching-code approval and Granted EVCC issuance remain out of scope for this foundation
- [ ] #11 Repository route and frontend regression coverage verifies PostgreSQL identity behavior blind filtering and MongoDB non-regression
- [ ] #12 Relevant documentation reflects the delivered PostgreSQL identity boundary and deferred cross-device transfer behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Scope

Deliver the PostgreSQL-only Platform Identity and Event Visitor Identity foundation. PostgreSQL events receive opaque server-issued Event Visitor Identity UUIDs and private Event Visitor Control Credentials (EVCCs); authenticated visitors may associate those identities with private Platform Identity UUIDs. Every PostgreSQL availability response is owned by an Event Visitor Identity, and one visitor may own multiple responses for an event. Preserve MongoDB without migration, mutation, data lookup, or credential changes. Do not implement transfer links, matching-code approval, Granted EVCC issuance, source revocation, or event-owner transfer.

## Schema and repository

1. Add PostgreSQL Platform Identity and Event Visitor Identity relations with UUIDv7 internal identifiers and private session-subject lookup.
2. Make every PostgreSQL response refer to an Event Visitor Identity and enforce that the response event and owner identity belong to the same event through a composite constraint or equivalent invariant.
3. Migrate PostgreSQL account and guest responses to Event Visitor Identities without MongoDB access.
4. Issue, validate, and store EVCC authority separately from public eventVisitorId values. EVCCs authorize every response owned by their Event Visitor Identity; do not retain PostgreSQL per-response protected-response credentials.
5. Support multiple responses per Event Visitor Identity and opaque response-ID lookups.

## HTTP and frontend boundaries

1. Issue and retain public eventVisitorId context only for PostgreSQL events; it is non-authorizing.
2. Keep EVCC values private to browser credential transport and unavailable to application JavaScript.
3. Define PostgreSQL-only opaque response IDs plus explicit selected-response, create, edit, and delete behavior for visitors with multiple responses.
4. Filter blind-availability payloads so non-owners receive only responses they are authorized to manage and no other-response counts; owners receive all responses.
5. Associate validated PostgreSQL event/visitor pairs during sign-in without changing Mongo flows.
6. Keep Mongo request shapes and paths unchanged.

## Verification and documentation

1. Add repository, route, and frontend coverage for identities, EVCC authorization, event/visitor integrity, multiple responses, opaque response selection, blind filtering, and Mongo non-regression.
2. Update PostgreSQL compatibility documentation for the UUID, EVCC, response-selection, and blind-availability boundaries.
3. Run the isolated PostgreSQL route suite, server tests, required frontend checks, and graphify update after code changes.

## Review gate

The revised documentation and stored plan require explicit user review and approval before runtime implementation begins.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research completed before the review gate. PostgreSQL creation currently runs only for anonymous requests (`postgresCreationEnabled` rejects signed-in sessions), so creation must issue and return an unassociated Event Visitor Identity; the subsequent sign-in association flow attaches it to the private Platform Identity.

Current frontend response selection and analytics use `authUser._id` as a `responses` map key. PostgreSQL response UUID map keys require an explicit non-subject selected-response contract before implementation.

Confirmed seams: event-source dispatch is centralized in `server/routes/events.go`; all frontend requests pass through `frontend/src/utils/fetch_utils.ts`; event fetch/response decoding is `frontend/src/composables/event/eventTransportBoundary.ts`; event creation currently returns only `eventId`; OAuth sign-in payload is owned by `Auth.vue` and `UserService.ts`. `postgresResponseModel` currently calls the Mongo-dependent `populateResponsePayloadIdentity` and must be replaced for PostgreSQL. No runtime code changed.

The confirmed documentation model establishes source EVCC authority in this foundation. Matching-code transfers, Granted EVCC issuance, source revocation, and owner-role delegation remain deferred to TASK-0071.02.
<!-- SECTION:NOTES:END -->
