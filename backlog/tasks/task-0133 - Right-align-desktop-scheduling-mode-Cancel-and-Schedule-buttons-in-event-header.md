---
id: TASK-0133
title: >-
  Right-align desktop scheduling-mode Cancel and Schedule buttons in event
  header
status: Done
assignee: []
created_date: '2026-09-02 09:42'
updated_date: '2026-09-02 09:48'
labels: []
dependencies: []
priority: medium
type: bug
ordinal: 146300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On the desktop event page scheduling mode (Cancel / Clear / Schedule in the event header), the button group container at frontend/src/views/Event.vue line ~692 uses class "desktop-event-header-actions tw-flex tw-gap-2" with no right alignment. The sibling containers for the Delete button (line ~647) and the "Schedule event" button (line ~667) both use "tw-flex tw-justify-end sm:tw-ml-auto". When the event description is empty, no flex-1 description spacer renders, so the Cancel/Schedule group sits on the left under Edit event / Copy link instead of on the right under the Show best times / More options options row. Fix with a layout-based change consistent with the sibling containers; no override hacks. Add a regression test asserting the scheduling button group container is right-aligned (sm:tw-ml-auto), following the existing parentElement className assertion pattern in Event.test.ts (see "moves Clear next to Cancel while rescheduling on desktop" and the sm:tw-ml-auto assertion near line 1252). Scope is desktop header layout only; do not change mobile action bar behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 When the event description is empty, the desktop scheduling-mode Cancel/Clear/Schedule button group aligns to the right edge of the header row, directly under the Show best times / More options options row
- [x] #2 When the event description is present, the desktop scheduling-mode button group remains right-aligned
- [x] #3 A regression test in frontend/src/views/Event.test.ts asserts the scheduling button group container is right-aligned (sm:tw-ml-auto)
- [x] #4 npm run lint, npm run typecheck, npm run build, and npm run test:unit pass from frontend/
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Event.vue ~line 692: change the scheduling button group container class from "desktop-event-header-actions tw-flex tw-gap-2" to "desktop-event-header-actions tw-flex tw-gap-2 tw-justify-end sm:tw-ml-auto", matching the right-alignment pattern of the Delete and Schedule event sibling containers (lines ~647 and ~667).

2. Event.test.ts: add a desktop scheduling-mode regression test asserting the Cancel/Schedule button group container (parent of the Cancel button) includes sm:tw-ml-auto, using the existing parentElement className assertion pattern.

3. Run npm run lint, typecheck, build, test:unit from frontend/.
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Right-aligned the desktop scheduling-mode Cancel/Clear/Schedule button group in the event header. The container in frontend/src/views/Event.vue now uses "desktop-event-header-actions tw-flex tw-gap-2 tw-justify-end sm:tw-ml-auto", matching the right-alignment pattern already used by the Delete and "Schedule event" sibling containers, so the group sits on the right under the Show best times / More options options row whether or not the event description renders. Added a regression test in frontend/src/views/Event.test.ts ("right-aligns desktop scheduling buttons under the header options when the description is empty") that mounts with the scheduling stub, confirms the description stub is absent (empty-description scenario), and asserts the Cancel button's container includes desktop-event-header-actions and sm:tw-ml-auto. Required checks from frontend/: npm run lint, npm run typecheck, npm run build, and npm run test:unit all pass (138 files, 985 tests). No Markdown files changed.
<!-- SECTION:FINAL_SUMMARY:END -->
