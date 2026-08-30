---
id: TASK-0105
title: Move mobile Calendar options into the Available/If needed editing panel row
status: Done
assignee: []
created_date: '2026-08-29 12:11'
updated_date: '2026-08-29 12:31'
labels:
  - mobile
  - schedule-overlap
  - ui
milestone: Mobile event response editing
dependencies: []
references:
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlap.vue
priority: medium
type: enhancement
ordinal: 111000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On mobile (phone viewport), on the Timed Event Response Editing Page, the Availability editing stack currently renders the Available/If needed toggle as a full-width elevated panel above the Delete/Cancel/Save bar, and a collapsible "Options" section (ExpandableSection) above the Legend Section in the sidebar containing only a "Calendar options..." button.

Outcome:
- A "Calendar options..." button sits on the left of one row inside the elevated editing panel above the Delete/Cancel/Save bar and opens the existing calendar options dialog (buffer time / working hours), which must keep working from the new location.
- The Available/If needed toggle is narrower and sits on the right of that same row.
- The collapsible "Options" section above the Legend Section is removed on mobile. The desktop sidebar "Calendar options..." button and the dialog are unchanged.

Note: the row order was corrected by user request after initial delivery (the first implementation placed the toggle left and the button right); the outcome above reflects the corrected order.

Scope is frontend-only, phone viewports, EDIT_AVAILABILITY state. Keep the layout-based styling approach and the existing shared elevation class. The options-section element exposure used for scroll-visibility tracking must keep working.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On a phone viewport while editing availability, the elevated panel above the Delete/Cancel/Save bar shows a Calendar options... button on the left and the Available/If needed toggle on the right in the same row, with the toggle narrower than the previous full-width layout
- [x] #2 Clicking the Calendar options... button in the mobile editing panel opens the existing calendar options dialog with buffer time and working hours controls
- [x] #3 On a phone viewport, the collapsible Options section above the Legend Section is gone and no Calendar options button remains in the sidebar while editing
- [x] #4 Desktop sidebar Calendar options button, dialog behavior, and the options-section element exposure used for scroll-visibility tracking are unchanged
- [x] #5 Unit regression coverage covers the new mobile row, the dialog-open emit, and the removal of the mobile Options section
- [x] #6 Required frontend checks (lint, typecheck, build, unit) pass
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
## Research findings

- The mobile editing panel is `ScheduleOverlapMobileOverlay.vue` (elevated panel with `AvailabilityTypeToggle` at `tw-w-full`), fixed above the Delete/Cancel/Save bar.
- The collapsible "Options" section above the legend is the phone-only `ExpandableSection` in `ScheduleOverlapSidebar.vue` (lines 174-201) holding the "Calendar options..." button; the dialog with BufferTimeSwitch/WorkingHoursToggle lives in the same wrapper, gated by `!sidebar.event.daysOnly && sidebar.showCalendarOptions`.
- `showCalendarOptions` is computed in `ScheduleOverlap.vue` (not addingAvailabilityAsGuest + calendarPermissionGranted + (isGroup || !userHasResponded)) and passed to the sidebar viewmodel; the overlay viewmodel lacks it.
- `calendarOptionsDialog` state is owned by `useScheduleOverlapUI` and updated via `updateCalendarOptionsDialog` in ScheduleOverlap.vue; `mobileOverlayListeners` does not currently relay it.
- `showEditOptions`/`toggleShowEditOptions` exist only to persist the phone Options section expanded state and become dead code after removal (sidebar, contracts, viewmodels, ScheduleOverlap destructure/listeners, useScheduleOverlapUI localStorage state).
- `optionsSectionEl` exposure is consumed by `checkElementsVisible` for scroll-visibility tracking; keeping the wrapper div with the dialog preserves it on both desktop and phone.
- Desktop-only e2e (`event-page-no-responses-layout.spec.ts`) targets `.calendar-options-button`; desktop button stays, so it is unaffected.

## Implementation plan

1. `scheduleOverlapViewModelContracts.ts` + `useScheduleOverlapViewModels.ts` + `scheduleOverlapTestUtils.ts`: add `showCalendarOptions` to the mobile overlay viewmodel; drop `showEditOptions` from the sidebar viewmodel contract/builder.
2. `ScheduleOverlapMobileOverlay.vue`: restructure the editing panel into a row — AvailabilityTypeToggle on the left (flex-1, no longer full-width) and an outlined "Calendar options..." v-btn on the right, gated by `!overlay.event.daysOnly && overlay.showCalendarOptions`; add the `update:calendarOptionsDialog` emit.
3. `ScheduleOverlap.vue`: relay `"update:calendarOptionsDialog"` in `mobileOverlayListeners`; remove `showEditOptions`/`toggleShowEditOptions` destructures and the sidebar listener.
4. `ScheduleOverlapSidebar.vue`: remove the phone `ExpandableSection` branch and its import/emit, keep the desktop button and move the dialog out of the removed section (keep wrapper div + `optionsSectionRef`).
5. `useScheduleOverlapUI.ts`: remove the now-dead `showEditOptions` state, localStorage handling, and `toggleShowEditOptions`.
6. Tests: `ScheduleOverlapSidebar.test.ts` replace the mobile Options-section test with one asserting phone editing renders no Options section/Calendar options button; `ScheduleOverlapMobileOverlay.test.ts` add coverage for the toggle+button row and the `update:calendarOptionsDialog` emit (and gating).
7. Run lint, typecheck, build, unit in `frontend/`; update the graph with `graphify update .`.

## Correction (user request): swap the mobile editing row order

The initial delivery placed the Available/If needed toggle left and the Calendar options button right. User request: Calendar options on the left, Available/If needed on the right.

1. `ScheduleOverlapMobileOverlay.vue`: move the outlined `Calendar options...` v-btn before `AvailabilityTypeToggle` inside the row flex container; keep the toggle's `tw-flex-1 tw-min-w-0` and the button's `tw-shrink-0` and gating (`!overlay.event.daysOnly && overlay.showCalendarOptions`) unchanged.
2. `ScheduleOverlapMobileOverlay.test.ts`: update the row test to assert DOM order (button before toggle) and keep the emit + gating coverage.
3. Re-run required checks in `frontend/` (lint, typecheck, build, unit); update graph with `graphify update .`.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Correction executed: swapped the mobile editing row so the Calendar options button renders before AvailabilityTypeToggle (button left, toggle right). Gating, classes, emit path, and dialog state handling unchanged. Updated the overlay row test to assert DOM order and re-ran lint, typecheck, build, and the full unit suite — all green (136 files / 942 tests). Graph refreshed via `graphify update .`.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: user
created: 2026-08-29 12:28
---
Correction follow-up to the delivered row layout: put **Calendar options... on the left** and the **Available/If needed toggle on the right** (initial delivery had toggle left / button right). Track and verify in this task.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## What changed

- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue`: the elevated editing panel above the Delete/Cancel/Save bar is now one row — Available/If needed toggle on the left (`tw-flex-1 tw-min-w-0`, replacing full width) and an outlined "Calendar options..." button on the right, gated by `!overlay.event.daysOnly && overlay.showCalendarOptions`. Clicking emits `update:calendarOptionsDialog`.
- `ScheduleOverlap.vue`: `mobileOverlayListeners` relays `update:calendarOptionsDialog` to the existing `updateCalendarOptionsDialog` handler, so the dialog state path is identical to the desktop button.
- `ScheduleOverlapSidebar.vue`: removed the phone-only collapsible "Options" `ExpandableSection` (and its import/emit). The desktop "Calendar options..." button and the dialog (buffer time / working hours) are unchanged; the wrapper div with `optionsSectionRef` remains so the `optionsSectionEl` exposure used for scroll-visibility tracking still resolves on both phone and desktop.
- Dead plumbing removed: `showEditOptions` state + localStorage read + `toggleShowEditOptions` in `useScheduleOverlapUI.ts`, and `showEditOptions` from the sidebar viewmodel contract, viewmodels composable, ScheduleOverlap destructures/listeners, and test builders. Added `showCalendarOptions` to the mobile overlay viewmodel.

## Verification (AC evidence)

- AC #1/#3: unit tests — overlay test asserts the toggle carries `tw-flex-1` (not `tw-w-full`) with the `.calendar-options-button` beside it; sidebar test asserts phone editing renders no `.calendar-options-button` and no Options section.
- AC #2: unit test asserts the overlay button click emits `update:calendarOptionsDialog` [true]; ScheduleOverlap relays it to the same dialog-state handler the desktop button uses (desktop button click emit still asserted in `ScheduleOverlapSidebar.test.ts`).
- AC #4: desktop `.calendar-options-button` legend-adjacency e2e (`event-page-no-responses-layout.spec.ts` "timed add availability controls stay close to the Legend") passes; `optionsSectionEl` exposure unit test passes.
- AC #5: new overlay tests (row + emit + days-only/showCalendarOptions gating) and replacement sidebar test.
- AC #6: `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit` (136 files / 942 tests) all pass in `frontend/`.

## e2e notes

- Isolated test stack: `chromium-mobile event-mobile-editing-options.spec.ts` — 2 passed (mobile editing flow with the new panel).
- Pre-existing, unrelated failures on a clean tree (verified by stashing this change): `event-page-no-responses-layout.spec.ts` "pairs each header row..." (both projects) and `event-page-days-only-layout.spec.ts` "inline Start on Monday switch" fail identically without this change; all other tests in those files that ran match the clean-tree results.

## Risks / follow-ups

- The two pre-existing e2e failures above are untouched by this task and may warrant their own investigation.
- Graph updated with `graphify update .`; changed Markdown formatted via `npm run format:markdown`.

## Correction (row order swap)

### What changed

- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue`: the elevated editing row now renders the outlined `Calendar options...` button on the left and the `AvailabilityTypeToggle` on the right. The toggle keeps `tw-flex-1 tw-min-w-0` (narrower than the original full-width layout) and the button keeps `tw-shrink-0`, `variant="outlined"`, and the days-only / `showCalendarOptions` gating. The `update:calendarOptionsDialog` emit path is unchanged.
- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts`: the row test now asserts DOM order — `.calendar-options-button` precedes the `AvailabilityTypeToggle` element among the row's children — and still covers the toggle's `tw-flex-1` (not `tw-w-full`) classes, the `update:calendarOptionsDialog` [true] emit on click, and the days-only / `showCalendarOptions` gating.

### Verification (AC evidence)

- AC #1: unit test asserts the button is left of the toggle in DOM order and the toggle keeps `tw-flex-1` / not `tw-w-full`.
- AC #2: the same unit test asserts clicking the button emits `update:calendarOptionsDialog` [true]; the ScheduleOverlap relay to `updateCalendarOptionsDialog` is unchanged from the original delivery.
- AC #3/#4: untouched by this correction — full unit suite (136 files / 942 tests) passes, including the sidebar tests asserting no mobile Options section / sidebar button and the desktop button behavior, plus the `optionsSectionEl` exposure test.
- AC #5: ordering assertion added to the overlay row test.
- AC #6: `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit` all pass in `frontend/`.

### Notes

- Graph updated with `graphify update .`; no repo Markdown files changed in this correction (task record is Backlog-managed).
- e2e impact: none — the mobile e2e spec exercises the dialog open flow, not side-specific ordering; the desktop `.calendar-options-button` e2e is unaffected. The two pre-existing e2e failures recorded in the original final summary remain out of scope.
<!-- SECTION:FINAL_SUMMARY:END -->
