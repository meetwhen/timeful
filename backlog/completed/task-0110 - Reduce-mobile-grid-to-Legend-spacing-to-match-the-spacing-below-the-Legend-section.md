---
id: TASK-0110
title: >-
  Reduce mobile grid-to-Legend spacing to match the spacing below the Legend
  section
status: Done
assignee:
  - OpenCode
created_date: '2026-08-29 15:17'
updated_date: '2026-08-29 15:46'
labels:
  - mobile
  - schedule-overlap
  - ui
dependencies: []
references:
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlap.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
modified_files:
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlap.mobileLegendSpacing.test.ts
priority: medium
type: enhancement
ordinal: 116000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On a phone viewport, while editing availability on the event response editing page, there is a large empty vertical gap between the bottom of the schedule grid and the Legend section heading, while the vertical space below the Legend section is much smaller. Reduce the grid-to-Legend gap so it visually matches the existing space below the Legend section.

Constraints: the fix must be a clean layout-level spacing change at the section boundary (margin, padding, or gap), not viewport-specific magic numbers, scroll hacks, or z-index/overlay work. Do not regress the TASK-0102 contract that the Legend Section fully clears the fixed Available/If needed panel stack at maximum scroll while editing on mobile. Desktop and non-editing layouts keep their current grid-to-Legend spacing.

Context: the Legend lives in the ColorLegend section at the end of frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue; the mobile editing experience spans frontend/src/components/schedule_overlap/ScheduleOverlap.vue and frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On a phone viewport, while editing availability, the vertical space between the schedule grid and the Legend section visually matches the vertical space below the Legend section
- [x] #2 Desktop and non-editing layouts keep their current spacing between the schedule grid and the Legend section
- [x] #3 At maximum scroll while editing on mobile, no legend row is obscured by the fixed Available/If needed panel or the bottom action bar (TASK-0102 contract still holds)
- [x] #4 The spacing change is implemented as clean section-boundary layout, not viewport-specific magic numbers or scroll hacks
- [x] #5 Unit regression coverage verifies the mobile grid-to-Legend spacing contract
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
1. In ScheduleOverlapSidebar.vue, guard the EDIT_AVAILABILITY `tw-pt-14` branch with `!sidebar.isPhone`. That 56px padding reserves clearance for the desktop-only absolute pager and the inline editing controls; on phone those controls live in ScheduleOverlapMobileOverlay (TASK-0105/0106), so the padding is dead space between grid and Legend. Phone editing then gets no extra top padding, giving grid→Legend = root wrapper py-2 (8px) + sidebar py-4 (16px) + editing block mb-2 (8px) = 32px, which visually matches the ~26-28px visible below the Legend with no magic numbers, and equals the grid→first-content spacing of every other mobile state.
2. Add a semantic class `schedule-overlap-sidebar__body` to that spacing div (matches existing `__pager` / `__tool-row` naming) as a stable test-selection hook.
3. New regression test ScheduleOverlap.mobileLegendSpacing.test.ts following the ScheduleOverlap.mobileLegendClearance.test.ts mount pattern: phone (viewportWidth 375) + editing → body div has no `tw-pt-14`; desktop (1024) + editing → body div keeps `tw-pt-14`; phone non-editing → body div keeps `tw-pt-2`.
4. Browser-verify at a phone viewport that the grid→Legend gap visually matches the space below the Legend and that the TASK-0102 contract still holds at max scroll (spacer math depends only on below-legend tail, which is unchanged).
5. Run required frontend checks: lint, typecheck, build, test:unit. Then graphify update and markdown format.

Execution deviation (recorded after research): browser decomposition showed a second dead-space source. The editing-options block rendered two always-empty wrapper divs on phone, and its tw-gap-5 added 20px between them. Added step: remove the single-child AvailabilityTypeToggle wrapper div (spacing-neutral on desktop) in addition to the pt-14 phone guard. Result: grid→Legend box gap 24px (16px sidebar py-4 + 8px editing block mb-2), measured 28.7px ink-to-ink, matching the ~26px below-Legend gap on the user's page.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Browser measurement found a second dead-space source beyond the pt-14 guard: the editing block's tw-gap-5 applied 20px between two always-empty wrapper divs on phone (AvailabilityTypeToggle wrapper and the Calendar options wrapper). Fixed by deleting the single-child toggle wrapper (spacing-neutral on desktop) and phone-guarding pt-14; measured gap grid→Legend heading: 44px → 28.7px on a 390px viewport, vs ~26px below the Legend on the user's page.

The below-Legend space at max scroll is page-dependent (TASK-0102's 64px tail assumption covers sidebar paddings + footer gap): the measurement page with a large footer showed 148px clearance, the user's page shows ~26px. The contract (no legend row obscured) holds on both; the below-legend flow is byte-identical before/after this change.

Temp measurement spec e2e/tmp-legend-spacing-measure.spec.ts was used for evidence and deleted afterwards; unit regression coverage in ScheduleOverlap.mobileLegendSpacing.test.ts is the permanent lock.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Reduced the mobile grid-to-Legend gap on the event response editing page so it visually matches the space below the Legend section.

## Root cause
Two sources of dead space in ScheduleOverlapSidebar, both phone-editing-only:
1. The sidebar body applied `tw-pt-14` (56px) for EDIT_AVAILABILITY regardless of viewport. That padding reserves clearance for the desktop-only absolute pager and the inline editing controls, which on phone moved into ScheduleOverlapMobileOverlay (TASK-0105/0106), leaving 56px of dead space between the grid and the Legend.
2. The editing-options block (`tw-flex tw-flex-col tw-gap-5`) rendered two always-empty wrapper divs on phone (the AvailabilityTypeToggle wrapper exists only for the desktop-only toggle), and `gap-5` still applied 20px between the empty wrappers.

## Changes
- `ScheduleOverlapSidebar.vue`: guarded the EDIT_AVAILABILITY `tw-pt-14` branch with `!sidebar.isPhone`; added a semantic `schedule-overlap-sidebar__body` class to the sidebar body (matches `__pager` / `__tool-row` naming); removed the single-child wrapper div around AvailabilityTypeToggle (provably spacing-neutral on desktop: the block's `tw-gap-5` now separates the toggle from its neighbors directly).
- New regression spec `ScheduleOverlap.mobileLegendSpacing.test.ts` (6 tests): phone editing drops `tw-pt-14` and renders no availability toggle; desktop editing keeps `tw-pt-14` and the toggle; phone/desktop non-editing keep `tw-pt-2`; desktop days-only keeps `tw-pt-16`.

## Verification
- Browser measurement (Playwright, chromium-mobile 390px viewport, seeded timed event, editing state): grid-pane bottom to Legend ink-top went 44px → 28.7px after the fix (was ~113px on the user's page before). The user's page shows ~26-28px below the Legend (TASK-0102 engineered 26px); below-legend flow is untouched by this change, so the two gaps now match within ~3px.
- TASK-0102 contract re-verified in-browser: at maximum scroll while editing, legend clearance above the fixed panel stack was 148px on a footer-heavy test page (contract requires ≥ 0; below-legend layout unchanged, so the user-page value stays ~26px).
- Checks: lint, typecheck, build, test:unit (966 passed) all green; chromium-mobile e2e for schedule-overlap-mobile-scroll and event-mobile-editing-options (3 tests) passed. Desktop spacing locked by unit tests.
<!-- SECTION:FINAL_SUMMARY:END -->
