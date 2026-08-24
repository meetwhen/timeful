---
id: TASK-0059
title: Clarify controlled terminology for event visitors and scheduling
status: Done
assignee:
  - '@OpenCode'
created_date: '2026-08-24 18:37'
updated_date: '2026-08-24 19:25'
labels:
  - requirements
  - terminology
dependencies: []
references:
  - docs/terminology/glossary.md
  - docs/requirements/functional/fr/FR-073.md
  - docs/design/architecture/adr/ADR-009.md
documentation:
  - docs/terminology/README.md
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
modified_files:
  - docs/terminology/glossary.md
  - docs/design/architecture/adr/ADR-009.md
  - docs/requirements/functional/fr/FR-010.md
  - docs/requirements/functional/fr/FR-018.md
  - docs/requirements/functional/fr/FR-062.md
  - docs/requirements/functional/fr/FR-063.md
  - docs/requirements/functional/fr/FR-072.md
  - docs/requirements/functional/fr/FR-073.md
priority: medium
type: docs
ordinal: 61000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Align the controlled glossary and canonical functional requirements with the agreed event-visitor identity model and consistent scheduling, availability-editing, product-mode, and UI terminology. An Event Visitor receives a browser-local, event-scoped, non-authorizing Event Visitor Identity when opening an event; an Event Owner receives one when event creation begins. Event Owner and Event Guest are separate Event Visitor roles: owners may edit Event Settings and create responses, while guests may create responses but may not edit Event Settings. An Event Owner Edit Token authorizes Event Settings edits and is distinct from guest-response credentials. Timed Event and Dates-Only Event remain canonical domain names, with Dates and times and Dates only as UI labels.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The glossary defines the agreed visitor identity role scheduling availability editing product and UI terms with authoritative context
- [x] #2 Canonical functional requirements use the controlled terminology consistently and link first applicable occurrences
- [x] #3 Timed Event and Dates-Only Event are distinguished from their Dates and times and Dates only UI labels
- [x] #4 Event Visitor Identity persistence and the separate Event Owner and Event Guest roles are documented consistently in relevant requirements and ADRs
- [x] #5 Glossary anchors terminology links and requirement title front matter and index consistency are verified
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Correct the glossary so Event Owner and Event Guest are distinct Event Visitor roles, add Platform Visitor for platform-wide actors, and define the distinct Event Owner Edit Token.
2. Update FR-073 to define the proposed event-scoped, non-authorizing Event Visitor Identity lifecycle, including earlier creation for Event Owners; update FR-018 and FR-063 with the Event Owner Edit Token; update FR-072 to use Platform Visitor.
3. Correct ADR-009 so its existing guestId describes guest-response ownership rather than the future Event Visitor Identity.
4. Validate glossary anchors, local terminology links, and requirement title/front-matter/index consistency. This documentation-only task defines the planned owner-token behavior but does not implement it.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Reopened after review identified a mismatch between the staged role model and the agreed terminology. The correction distinguishes Event Owner from Event Guest, introduces Platform Visitor, preserves `guestId` as guest-response ownership data, and documents the future Event Visitor Identity separately.

Review-driven clarification finalized the role model: Event Owner and Event Guest are separate Event Visitor roles. Event Visitor Identity is a browser-local, event-scoped, non-authorizing `eventVisitorId`; owners receive it when creation begins. Event Owner Edit Token is event-scoped and distinct from guest-response credentials. ADR-009 now scopes `guestId` to guest-response ownership.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Corrected the staged terminology model after review. The glossary now distinguishes Platform Visitor, Event Visitor, Event Owner, and Event Guest; documents the proposed non-authorizing, event-scoped `eventVisitorId`; and defines the separate Event Owner Edit Token. FR-018, FR-062, FR-063, FR-072, and FR-073 now reflect the approved role, authentication, and identity semantics. ADR-009 no longer conflates `guestId` with Event Visitor Identity, and FR-010 now correctly distinguishes Timed Slots from Slot Increment.

Verification: `git diff --check`, scoped Prettier, and Markdown link validation for all changed documentation passed. Unit and E2E tests were not run because this is documentation-only work. The full-worktree link check reports unrelated unavailable local-development URLs in `backlog/backlog.md`.
<!-- SECTION:FINAL_SUMMARY:END -->
