---
id: TASK-0068
title: Create event descriptions before event publication
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-25 10:41'
updated_date: '2026-08-25 11:11'
labels:
  - events
  - event-description
dependencies: []
references:
  - docs/requirements/functional/fr/FR-019.md
  - docs/requirements/functional/fr/FR-035.md
priority: medium
type: feature
ordinal: 70000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Event creators need to provide an optional description while creating an event. Published event pages should present a saved description as read-only content and avoid reserving a description area when none was provided.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 A creator can enter an optional multiline description while creating an event
- [ ] #2 A created description is persisted and returned with the event
- [ ] #3 The event page renders a non-empty description as read-only content
- [ ] #4 The event page has no description card or add/edit control when the description is empty
- [ ] #5 Description behavior is covered by frontend and server regression tests
- [ ] #6 Functional requirements and generated API documentation reflect the delivered behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Update FR-035 and FR-077 to use the repository's canonical functional-requirement language, including linked controlled Event Owner terminology.
2. Treat descriptions containing only whitespace as empty in EventDescription, so no description card is rendered. Make the Event.vue description slot conditional too, so mobile does not retain an empty layout item.
3. Extend component coverage for whitespace-only omission and replace renderer-internal empty assertions; extend the event-page E2E regression to exercise desktop and mobile empty-description layout.
4. Run targeted tests, the required frontend checks, Firefox E2E, and graphify incremental update; then record objective evidence in the task.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Replaced obsolete event-page description editor coverage with read-only multiline and empty-state unit coverage. Creation payload coverage now verifies multiline description POST inclusion, while edit payload coverage verifies omission. Updated event-page component binding, the description fixture/inspection scenarios, the targeted layout E2E spec, and Mongo creation route regression assertions.

Normalized event-page descriptions with `trim()` and made the Event.vue description wrapper conditional, removing whitespace-only cards and their mobile layout item. The desktop Schedule event action now uses `sm:tw-ml-auto` so it remains right-aligned when no description column exists. Added component and Chromium E2E coverage for whitespace omission, mobile layout omission, and desktop Schedule alignment. `npm run test:e2e -- --project=firefox-desktop` had 21 passing, 1 skipped, and 1 unrelated failure: `timed-event-timezone-menu-firefox.spec.ts` measured 504px against its 514px lower bound.
<!-- SECTION:NOTES:END -->
