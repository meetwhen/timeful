---
id: TASK-0068
title: Create event descriptions before event publication
status: Done
assignee:
  - OpenCode
created_date: '2026-08-25 10:41'
updated_date: '2026-09-01 20:19'
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
- [x] #1 A creator can enter an optional multiline description while creating an event
- [x] #2 A creator can enter an optional multiline description while editing an event, prefilled with the saved description
- [x] #3 A created or edited description is persisted and returned with the event
- [x] #4 The event page renders a non-empty description as read-only content
- [x] #5 The event page has no description card or add/edit control when the description is empty
- [x] #6 Description behavior is covered by frontend and server regression tests
- [x] #7 Functional requirements and generated API documentation reflect the delivered behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Update FR-035 and FR-077 to use the repository's canonical functional-requirement language, including linked controlled Event Owner terminology.
2. Treat descriptions containing only whitespace as empty in EventDescription, so no description card is rendered. Make the Event.vue description slot conditional too, so mobile does not retain an empty layout item.
3. Extend component coverage for whitespace-only omission and replace renderer-internal empty assertions; extend the event-page E2E regression to exercise desktop and mobile empty-description layout.
4. Run targeted tests, the required frontend checks, Firefox E2E, and graphify incremental update; then record objective evidence in the task.

5. Restore edit-form description support: render the description field in edit mode, hydrate it from the saved event, track it in editor dirty state, and include it in the create and edit payloads; update the edit-payload regression coverage.

6. Extend FR-035 to cover creating and editing events with optional descriptions and align the requirements index row.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Replaced obsolete event-page description editor coverage with read-only multiline and empty-state unit coverage. Creation payload coverage now verifies multiline description POST inclusion, while edit payload coverage verifies omission. Updated event-page component binding, the description fixture/inspection scenarios, the targeted layout E2E spec, and Mongo creation route regression assertions.

Normalized event-page descriptions with `trim()` and made the Event.vue description wrapper conditional, removing whitespace-only cards and their mobile layout item. The desktop Schedule event action now uses `sm:tw-ml-auto` so it remains right-aligned when no description column exists. Added component and Chromium E2E coverage for whitespace omission, mobile layout omission, and desktop Schedule alignment. `npm run test:e2e -- --project=firefox-desktop` had 21 passing, 1 skipped, and 1 unrelated failure: `timed-event-timezone-menu-firefox.spec.ts` measured 504px against its 514px lower bound.

Scope expanded by user request: the edit event form shall also provide the description field, prefilled with the saved description, and send it in the edit payload. Current state: create form has the field and POST includes it; edit form hides the field and omits it from the PUT payload; server edit routes on both Mongo and Postgres already persist an explicitly sent description and retain it when omitted; event page renders the description read-only via EventDescription.vue.

Edit-form support delivered: removed the `v-if="!edit"` gate from the NewEvent description textarea, hydrated `description` from the saved event in `onEventHydrate`, added description to the editor dirty-tracking snapshot (`captureExtraInitialState` / `isExtraEdited`), and made the create/edit payload always include `description`. Replaced the edit-payload omission regression with coverage asserting the saved description is prefilled and sent, plus an explicitly cleared description sending an empty string. Updated FR-035 to cover creating and editing (status accepted) and aligned its requirements-index row. All four required frontend checks pass; unit suite 138 files / 979 tests pass. Firefox E2E: 24 passed, 2 skipped, 1 pre-existing flake (`timed-event-timezone-menu-firefox.spec.ts` measured 512 vs 514 lower bound in-suite; previously failed at 504 before this change, and passes in isolation). Server code unchanged; existing Mongo and Postgres route description regressions cover persistence and omit-retain semantics.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Restored full event description functionality. The create and edit event forms now both provide the optional multiline description field: the edit form hydrates it from the saved event, includes it in dirty tracking, and sends it in the PUT payload (explicit empty string clears it; omitting it elsewhere retains the saved value per the server contract on both Mongo and Postgres). The event page continues to render a saved non-whitespace description as read-only content with no card or control when empty. Unit coverage asserts prefill, sent value on edit, and cleared-value on edit; server regressions already covered persistence and omit-retain semantics. FR-035 now covers creating and editing events with optional descriptions. All required frontend checks pass (lint, typecheck, build, 979 unit tests); Firefox E2E 24 passed / 2 skipped with one pre-existing flake unrelated to this change.
<!-- SECTION:FINAL_SUMMARY:END -->
