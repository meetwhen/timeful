---
id: TASK-0143
title: Render mobile slot tooltip below the event-page bottom overlay panels
status: Done
assignee:
  - opencode
created_date: '2026-09-02 14:56'
updated_date: '2026-09-02 16:48'
labels: []
milestone: Mobile UX
dependencies: []
references:
  - TASK-0033
  - TASK-0111
  - TASK-0040
  - TASK-0127
priority: medium
type: bug
ordinal: 156300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On phones, tapping a slot on the event page opens the slot tooltip. When the user scrolls the page down, the fixed bottom overlay panels (sticky Responses panel, Available/If needed editing panel, mobile action bar) can overlap the tooltip. Currently the tooltip renders above these panels, so it visually sits on top of them and reads as floating over the panels instead of being part of the page content. The tooltip shall be stacked below the bottom overlay panels so the panels occlude it. Work should respect the centralized tooltip placement logic (TASK-0033) and the shared mobile elevated-panel treatment (TASK-0111).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On a phone-sized viewport, when a slot tooltip is visible on the event page and the user scrolls so a bottom overlay panel (sticky Responses panel, Available/If needed editing panel, or mobile action bar) overlaps the tooltip, the bottom overlay panels render above the tooltip and fully occlude it
- [x] #2 The stacking contract is encoded once in a shared bottom-overlay stacking layer (z-index 60) applied to every fixed bottom overlay container on the event page (mobile action bar wrapper and the overlay root hosting the sticky Responses panel and the Available/If needed editing panel); the shared tooltip stays at z-index 50 and no call site raises the tooltip above that layer
- [x] #3 No regression to the centralized tooltip placement and visibility behavior from TASK-0033, including tooltip re-anchoring on scroll
- [x] #4 A unit regression test asserts the tooltip stacking context sits below the shared mobile elevated bottom panels
- [x] #5 Tooltip behavior on desktop and on the event page grid above the bottom overlays is unchanged
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Decided stacking contract (user-approved 2026-09-02)

The event page's fixed bottom overlay containers (mobile action bar wrapper; ScheduleOverlapMobileOverlay root hosting the sticky Responses panel and the Available/If needed editing panel) share one bottom-overlay stacking layer (z-index 60) that always sits above the shared slot tooltip layer (z-index 50). The shared tooltip is not lowered; no per-call-site z-index patches.

## Steps

1. frontend/src/index.css: add `.timeful-bottom-overlay-layer { z-index: 60; }` beside `.timeful-mobile-elevated-panel`.
2. frontend/src/views/Event.vue mobile bottom bar wrapper (~line 839): replace `tw-z-20` with `timeful-bottom-overlay-layer`.
3. frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue root (line 4): replace `tw-z-[60]` with `timeful-bottom-overlay-layer`; keep `tw-isolate`.
4. Tooltip.vue unchanged (z-50; TASK-0033 placement logic untouched).
5. Unit regression tests: ScheduleOverlap.mobileTooltip.test.ts — phone viewport with tooltip visible, assert the tooltip keeps `tw-z-50` and `.schedule-overlap-mobile-overlay` carries `timeful-bottom-overlay-layer`. Event.test.ts — extend the existing mobile action bar test (~line 2491) to assert the fixed bottom wrapper carries `timeful-bottom-overlay-layer` instead of `tw-z-20`.
6. Verify: frontend npm run lint / fmt:check / typecheck / build / test:unit; targeted e2e stacking spec schedule-overlap-mobile-touch-firefox.spec.ts on firefox-desktop; graphify update .
7. Finalize: verify each acceptance criterion with evidence, record finalSummary, mark Done.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Decision (user-approved 2026-09-02): guarantee the slot tooltip always renders beneath the bottom overlay panels by elevating the fixed bottom overlay containers into one shared bottom-overlay stacking layer (z-index 60) above the shared tooltip layer (z-index 50). The tooltip is NOT lowered. Rationale: lowering the tooltip below z-20 would regress desktop (ScheduleOverlapSidebar pager is tw-z-20 and would occlude the tooltip), reverse the deliberate fix 3fea15d6 that raised the overlay root to z-60 above the tooltip, and force rewriting the .tw-fixed.tw-z-50 assertions in ScheduleOverlap.mobileTooltip.test.ts and the Firefox e2e locator. This revises the original AC #2 wording ('fix at the shared tooltip layer') into the shared bottom-overlay layer contract; the user explicitly asked to settle this decision and update the task.

Progress log: applied the shared `timeful-bottom-overlay-layer` (z-60) to the Event.vue mobile action bar wrapper and the ScheduleOverlapMobileOverlay root; index.css hosts the class beside the TASK-0111 elevated-panel treatment. Added the unit stacking regression test (mutation-verified: it fails when the action-bar wrapper reverts to z-20) and extended the Event.test.ts sticky-footer test with layer assertions.

AC #1 browser evidence: added the e2e test `mobile action bar stacks above an overlapping mobile tooltip` to schedule-overlap-mobile-touch-firefox.spec.ts (firefox-touch project). The first attempt failed because the create-flow page is still in SET_SPECIFIC_TIMES where Event.vue hides the action bar; the test now touch-selects a slot, clicks specific-times-grid-next (real touch selection is required for numTempTimes), then asserts the bar occludes the tooltip. Both stacking e2e tests pass on the isolated test stack; the spec's known TASK-0134 re-anchoring failure was not run and is unrelated.

Checks: frontend lint (0 errors, 2 pre-existing warnings in untouched NewEvent.test.ts), fmt:check, typecheck, build, and the full unit suite (138 files / 988 tests) pass. graphify update . recompiled the knowledge graph; npm run format:markdown ran clean (backlog/backlog.md diff is the MCP's own task registration).
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

The slot tooltip on the event page always renders beneath the fixed bottom overlay panels. The stacking contract is encoded once in a shared bottom-overlay layer instead of per-call-site z-index patches or lowering the shared tooltip.

## Changes

- `frontend/src/index.css`: added `.timeful-bottom-overlay-layer { z-index: 60; }` beside `.timeful-mobile-elevated-panel` so the bottom overlay containers share one stacking layer above the tooltip's z-50.
- `frontend/src/views/Event.vue`: the mobile bottom action bar wrapper (Add availability / Edit availability) moved from `tw-z-20` to `timeful-bottom-overlay-layer`.
- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue`: root container moved from `tw-z-[60]` to `timeful-bottom-overlay-layer` (keeps `tw-isolate`).
- `Tooltip.vue` untouched: the shared tooltip keeps z-50 and the TASK-0033 placement logic is unchanged.
- Unit regression tests: `ScheduleOverlap.mobileTooltip.test.ts` "stacks the shared tooltip layer below the mobile bottom overlay panels" asserts the tooltip keeps `tw-z-50` (no shared-layer class) while the overlay root carries `timeful-bottom-overlay-layer`; `Event.test.ts` sticky-footer test asserts the fixed bottom wrapper carries `timeful-bottom-overlay-layer` (and not `tw-z-20`). Mutation-verified: reverting the action-bar class fails the test.
- E2e: added `mobile action bar stacks above an overlapping mobile tooltip` to `schedule-overlap-mobile-touch-firefox.spec.ts` (firefox-touch), hit-testing that the action bar occludes a tooltip moved over it; the existing "Responses panel stacks above an overlapping mobile tooltip" still passes.

## Evidence

- `npm run lint` (0 errors; 2 pre-existing warnings in untouched `NewEvent.test.ts`), `npm run fmt:check`, `npm run typecheck`, `npm run build`, `npm run test:unit` (138 files / 988 tests) all pass in `frontend/`.
- firefox-touch e2e: both stacking tests pass via the isolated test stack.
- AC #1/#2: real-browser occlusion covered by both e2e stacking tests; layer contract covered by unit tests. AC #3: placement logic untouched; all tooltip-related unit suites pass. AC #5: the action bar is phone-only (`isPhone` guard) and the tooltip z-index is unchanged, so desktop stacking is untouched.
- Note: the same spec contains the pre-existing TASK-0134 failure ("touching a timeslot keeps its mobile tooltip anchored while scrolling"), unrelated to this change and not run here.
- Decision record: user approved elevating the bottom panels into the shared layer over lowering the tooltip (see implementation notes); AC #2 was rewritten accordingly before implementation.
<!-- SECTION:FINAL_SUMMARY:END -->
