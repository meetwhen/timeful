---
id: TASK-0121
title: Reduce desktop gap between time format controls and Editing availability as
status: Done
assignee:
  - '@OpenCode'
created_date: '2026-08-31 12:48'
updated_date: '2026-08-31 12:48'
labels:
  - desktop
  - schedule-overlap
  - ui
milestone: Event page polish
dependencies: []
references:
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - frontend/src/components/schedule_overlap/ToolRow.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlap.mobileLegendSpacing.test.ts
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.test.ts
modified_files:
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlap.mobileLegendSpacing.test.ts
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.test.ts
priority: medium
type: enhancement
ordinal: 117000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On desktop, on the event Response editing page, there is an overly large vertical gap between the time format controls row (12h/24h toggle, timezone selector, reset) and the "Editing availability as" block. Reduce that gap so the editing panel sits closer to the controls, consistent with the spacing rhythm already used between the editing sections.

Constraints: the fix must be a clean layout-level spacing change at the section boundary, not overrides, hacks, or viewport magic. Do not regress the TASK-0110 mobile contract (phone editing drops the desktop editing top padding), the desktop days-only `tw-pt-16` offset, or the non-editing `tw-pt-2` spacing. The space above the time format controls inside the compact ToolRow is out of scope.

Context: the gap source is the sidebar body top padding for the EDIT_AVAILABILITY state in frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue; the compact ToolRow renders in flow directly above it.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On desktop, while editing availability on a timed event response page, the vertical space between the time format controls row and the Editing availability as block is visibly reduced from the previous 56px
- [x] #2 Mobile editing, desktop days-only, and non-editing sidebar layouts keep their current spacing
- [x] #3 The spacing change is implemented as clean section-boundary layout, not overrides or hacks
- [x] #4 Unit regression coverage pins the new desktop editing spacing contract
- [x] #5 Required frontend checks (lint, typecheck, build, unit) pass
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Executed approach: the desktop EDIT_AVAILABILITY body padding in ScheduleOverlapSidebar.vue was the entire gap (compact ToolRow is in flow above the body and has pb-0), so `tw-pt-14` (56px) became `tw-pt-5` (20px) to match the editing block's `tw-gap-5` internal rhythm. Spacing pins updated in ScheduleOverlap.mobileLegendSpacing.test.ts (phone drops `tw-pt-5`, desktop keeps it) and ScheduleOverlapSidebar.test.ts (body-scoped `tw-pt-5` assertion replacing the whole-HTML `tw-pt-14` check, which the unrelated ToolRow padding would otherwise satisfy vacuously). Space above the controls inside the compact ToolRow (`tw-pt-14` for pager clearance) left untouched as out of scope. Verification: lint, typecheck, build, test:unit (960 passed) green; no e2e spec references the changed classes (verified via rg over frontend/e2e excluding inspect), so no e2e run was required; no Markdown files changed. graphify update run after the change.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Reduced the desktop gap between the time format controls row and the "Editing availability as" block on the event Response editing page.

## Root cause
In the desktop EDIT_AVAILABILITY state, the sidebar body (`schedule-overlap-sidebar__body`) applied `tw-pt-14` (56px) top padding directly below the in-flow compact ToolRow (time format controls). That 56px padding was the entire controls-to-content gap.

## Changes
- `ScheduleOverlapSidebar.vue`: changed the desktop editing body padding from `tw-pt-14` (56px) to `tw-pt-5` (20px), matching the `tw-gap-5` rhythm already used between the editing sections (Editing availability as, Available/If needed, Calendar options).
- Updated the spacing pins in `ScheduleOverlap.mobileLegendSpacing.test.ts` (phone editing drops `tw-pt-5`; desktop editing keeps `tw-pt-5`) and `ScheduleOverlapSidebar.test.ts` (desktop editing asserts the body element carries `tw-pt-5`, scoped to the body instead of the whole HTML so the unrelated ToolRow `tw-pt-14` cannot satisfy it).

## Verification
- Checks: lint, typecheck, build, test:unit (960 passed) all green.
- Desktop editing gap: 56px → 20px. Mobile editing, days-only (`tw-pt-16`), and non-editing (`tw-pt-2`) spacing untouched.
<!-- SECTION:FINAL_SUMMARY:END -->
