---
id: TASK-0106
title: >-
  Move mobile Editing availability as indicator into the editing panel above
  Calendar options
status: Done
assignee:
  - '@OpenCode'
created_date: '2026-08-29 13:00'
updated_date: '2026-08-29 13:41'
labels:
  - mobile
  - schedule-overlap
  - ui
milestone: Mobile event response editing
dependencies: []
references:
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlap.vue
priority: medium
type: enhancement
ordinal: 112000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On the Timed Event Response Editing Page at phone viewports while editing availability (EDIT_AVAILABILITY state), the italic "Editing/Adding availability as <name>" indicator currently renders in the sidebar between the bottom of the availability grid and the Legend section. The elevated editing panel above the Delete/Cancel/Save bar (reshaped by TASK-0105) currently contains one row with the "Calendar options..." button on the left and the Available/If needed toggle on the right.

Outcome:
- While editing availability on a phone viewport, the "Editing availability as <name>" (or "Adding availability as ...") indicator renders inside the elevated editing panel, above the row containing the Calendar options button and the Available/If needed toggle.
- The indicator keeps its dynamic actor text and name resolution (authenticated user, guest with editable name plus the pencil affordance that opens the Edit guest name dialog, or the anonymous "a guest" fallback), all working from the new location.
- The indicator no longer renders between the grid and the Legend on phone viewports, so the Legend no longer sits below that text block.
- Desktop editing layout (sidebar indicator placement and the Edit guest name dialog) is unchanged.

Scope is frontend-only, phone viewports, EDIT_AVAILABILITY state. Keep the layout-based styling approach and the existing shared elevation class on the panel. Related in-flight work: TASK-0102 (mobile Legend visibility above this same panel) touches the same region; confirm the current tree state before implementing and coordinate any overlapping layout math.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On a phone viewport while editing availability, the elevated panel above the Delete/Cancel/Save bar shows the Editing/Adding availability as indicator above the row containing the Calendar options button and the Available/If needed toggle
- [x] #2 The indicator keeps its dynamic actor text and name: authenticated user name, guest name with a working pencil affordance that opens the Edit guest name dialog and saves correctly, or the anonymous a-guest fallback
- [x] #3 On a phone viewport while editing, the indicator no longer renders between the grid and the Legend
- [x] #4 Desktop editing layout keeps the existing sidebar indicator placement with unchanged behavior
- [x] #5 Unit regression coverage covers the new indicator placement in the editing panel, its removal from the mobile grid-and-Legend location, and the guest-name edit affordance from the new location
- [x] #6 Required frontend checks (lint, typecheck, build, unit) pass
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
## Research findings

- The indicator lives in `ScheduleOverlapSidebar.vue` (EDIT_AVAILABILITY branch, first child of the `tw-gap-5` container, lines ~99-149) and includes the Edit guest name dialog. Display strings are computed locally in the sidebar component (`showEditingAsText`, `availabilityActorActionText`, `availabilityActorName`, `currentGuestName`).
- The elevated panel is `ScheduleOverlapMobileOverlay.vue`, gated `overlay.editing && !isGroup && !isSignUp`; it currently holds one row: Calendar options button (left, gated `!daysOnly && showCalendarOptions`) + AvailabilityTypeToggle (right, `tw-flex-1`). The mobile overlay only mounts on phone (`v-if="isPhone && !calendarOnly"` in ScheduleOverlap.vue).
- `editing` = EDIT_AVAILABILITY || EDIT_SIGN_UP_BLOCKS, so the panel renders exactly while the sidebar indicator block renders, except for groups (panel is group-gated off; sidebar indicator must stay for phone+group).
- Sidebar visibility logic `showEditingAsText = !(calendarPermissionGranted && !daysOnly && !addingAvailabilityAsGuest)`; actor text Editing/Adding from `userHasResponded`/`addingAvailabilityAsGuest`/`curGuestId`; actor name precedence: signed-in non-guest-user full name, then guest name, then "a guest"; pencil shown when `curGuestId && canEditGuestName` (opens the guest-name dialog).
- ScheduleOverlap.vue already owns `openEditGuestNameDialog`, `saveGuestName`, `updateNewGuestName`, `updateEditGuestNameDialog` (used by `sidebarListeners`); `mobileOverlayListeners` lacks them.
- Legend clearance spacer height derives from the overlay's ResizeObserver-measured height, so a taller panel self-adjusts (TASK-0102 mechanism). No legend test asserts the indicator position.

## Implementation plan

1. `scheduleOverlapViewModelContracts.ts`: add `ScheduleOverlapEditingAvailabilityAsViewModel` (`visible: boolean`, `actionText: "Editing" | "Adding"`, `actorName: string`, `editableGuestName: string | null`); add `editingAvailabilityAs` to the sidebar and overlay view models; add `editGuestNameDialog` + `newGuestName` to the overlay view model (dialog state travels with the indicator to the overlay on phone).
2. `useScheduleOverlapViewModels.ts`: compute `editingAvailabilityAs` once (exact port of the sidebar logic) and wire it into both view models.
3. New `EditingAvailabilityAs.vue` presentational component: italic "{actionText} availability as" line, bold guest name + pencil when `editableGuestName` is set (click emits `openEditGuestNameDialog`), and the Edit guest name dialog (props: `editingAs`, `editGuestNameDialog`, `newGuestName`; emits: `openEditGuestNameDialog`, `saveGuestName`, `update:newGuestName`, `update:editGuestNameDialog`). One canonical markup shared by both surfaces.
4. `ScheduleOverlapSidebar.vue`: replace the indicator+dialog block with the shared component, gated `visible && (!isPhone || isGroup)` so phone non-group defers to the overlay while desktop and phone groups keep the sidebar placement; delete the four local computeds.
5. `ScheduleOverlapMobileOverlay.vue`: render the shared component above the controls row (wrap panel content in a `tw-flex tw-flex-col tw-gap-3` column) and declare the four emits.
6. `ScheduleOverlap.vue`: relay the four guest-name events in `mobileOverlayListeners` to the existing handlers.
7. `scheduleOverlapTestUtils.ts`: add the new fields to both view-model builders (`visible: true` default).
8. Tests: overlay — indicator renders above the row when visible, absent when not, pencil emits `openEditGuestNameDialog`, dialog events re-emitted; sidebar — phone non-group editing renders no indicator, desktop still renders it.
9. Required checks (lint, typecheck, build, unit), the mobile editing e2e spec, then `graphify update .`.

Risks: duplicate dialog instances (sidebar + overlay) must never mount simultaneously — gating above guarantees it; mobile e2e may assert panel content ordering.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Executed the plan as written with no deviations. View-model gating guarantees the sidebar and overlay dialog instances never mount simultaneously (phone non-group: overlay only; desktop or phone group: sidebar only).

Verification evidence: lint, typecheck, build pass; full unit suite 137 files / 951 tests pass; chromium-mobile e2e event-mobile-editing-options.spec.ts 2/2 pass; format:markdown no-op (backlog/** excluded); graphify update . completed.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Moved the "Editing/Adding availability as <name>" indicator from the sidebar grid-and-Legend area into the elevated mobile editing panel (above the Calendar options + Available/If needed row) at phone viewports, keeping desktop behavior unchanged.

## Changes

- Added `EditingAvailabilityAs.vue`: shared presentational component rendering the italic "{actionText} availability as" line with dynamic actor name (authenticated user, guest with pencil affordance, or "a guest" fallback) and the Edit guest name dialog; one canonical markup shared by sidebar and mobile overlay.
- `scheduleOverlapViewModelContracts.ts`: new `ScheduleOverlapEditingAvailabilityAsViewModel` (`visible`, `actionText`, `actorName`, `editableGuestName`); wired `editingAvailabilityAs` into both sidebar and overlay view models, plus `editGuestNameDialog`/`newGuestName` on the overlay view model.
- `useScheduleOverlapViewModels.ts`: computes `editingAvailabilityAs` once (exact port of the former sidebar logic) so both surfaces share one source of truth.
- `ScheduleOverlapSidebar.vue`: replaced the local indicator+dialog block with the shared component, gated `visible && (!isPhone || isGroup)`; phone non-group defers to the overlay while desktop and phone-group keep sidebar placement; removed the four local computeds.
- `ScheduleOverlapMobileOverlay.vue`: renders the shared component above the controls row inside a `tw-flex-col tw-gap-3` column and declares the four guest-name emits.
- `ScheduleOverlap.vue`: relays the four guest-name events in `mobileOverlayListeners` to the existing handlers.

## Tests

- New `EditingAvailabilityAs.test.ts`: actor fallback, editable guest name + pencil emitting `openEditGuestNameDialog`, dialog cancel relay.
- Overlay tests: indicator renders above the Calendar options/toggle row while editing, hidden when the view model marks it invisible, guest-name events re-emitted.
- Sidebar tests: no indicator in phone sidebar while editing (non-group), indicator kept for phone group and desktop.
- `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit` (137 files, 951 tests) all pass; `event-mobile-editing-options.spec.ts` (chromium-mobile) passes 2/2; `graphify update .` run; `npm run format:markdown` no-op (only changed Markdown is the Backlog-managed task file, excluded by the formatter).

## Risks / follow-ups

- Sidebar and overlay dialog instances are mutually exclusive by the `!isPhone || isGroup` gating, so no duplicate dialogs can mount simultaneously.
<!-- SECTION:FINAL_SUMMARY:END -->
