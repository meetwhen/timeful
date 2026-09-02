---
id: TASK-0147
title: Show mobile selected-slot cursor at z-index 40 like the tooltip
status: Done
assignee:
  - opencode
created_date: '2026-09-02 22:39'
updated_date: '2026-09-02 22:56'
labels: []
dependencies: []
priority: medium
type: enhancement
ordinal: 160300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On mobile, the currently clicked slot is highlighted with a cursor outline (the `__selected-timeslot` cursor). It currently stacks at z-index 10 in the timed grid and with no explicit z-index in the days-only grid, so it can end up below other stacked layers. The mobile slot tooltip renders at z-40. Raise the selection cursor to that same z-40 layer so it stays visible above overlapping elements, consistently across the grids that draw it.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On mobile, the cursor highlighting the currently clicked slot renders at z-index 40, the same layer as the mobile slot tooltip, so overlapping elements no longer cover it
- [x] #2 The raised cursor layering applies wherever the selection cursor is drawn (timed grid and days-only grid), keeping both grids consistent
- [x] #3 The selection cursor remains pointer-events-transparent so slot interaction is unchanged
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
1. Timed grid: in `ScheduleOverlapTimeGrid.vue`, change `.schedule-overlap-time-grid__selected-timeslot::after` from `z-index: 10` to `z-index: 40` (same numeric layer as `.timeful-tooltip-layer` in `frontend/src/index.css`).
2. Days-only grid: in `ScheduleOverlapDaysOnlyGrid.vue`, add `z-index: 40` to `.schedule-overlap-days-only-grid__selected-timeslot::after` so both grids that draw the cursor share the same layer (currently implicit/auto).
3. Keep `pointer-events: none` on both rules so slot interaction is unchanged.
4. The rendering layer (`scheduleOverlapRendering.ts`) only applies the `tw-relative __selected-timeslot` classes; no changes needed there and no existing tests assert the old z-index.
5. Verify with the required frontend checks: lint, fmt:check, typecheck, build, test:unit.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Layer ladder confirmed in frontend/src/index.css: bottom-overlay 60, action-bar 50, tooltip-layer 40. Cursor now sits at the tooltip's numeric layer; tooltip stays readable since the fixed tooltip element mounts later in the DOM (tie-break).

First e2e run failed at page setup on a cold stack (Go compile within the 30s test timeout); rerun on the warm stack passed. Scratch spec removed after verification.

Timed-grid cursor previously z-index 10; days-only cursor had no explicit z-index (auto).
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Raised the mobile selected-slot cursor to z-index 40, the same numeric layer as the mobile tooltip (`.timeful-tooltip-layer` in `frontend/src/index.css`).

## Changes

- `frontend/src/components/schedule_overlap/ScheduleOverlapTimeGrid.vue`: `.schedule-overlap-time-grid__selected-timeslot::after` z-index 10 → 40.
- `frontend/src/components/schedule_overlap/ScheduleOverlapDaysOnlyGrid.vue`: `.schedule-overlap-days-only-grid__selected-timeslot::after` gained an explicit `z-index: 40` (was implicit/auto), so both grids that draw the cursor share the same layer.
- `pointer-events: none` retained on both rules, so the cursor never intercepts slot interaction.

## Verification

- Required frontend checks: `npm run lint` (clean; two pre-existing unrelated warnings), `npm run fmt:check`, `npm run typecheck`, `npm run build`, `npm run test:unit` (138 files / 990 tests passed).
- Browser evidence on the isolated e2e stack (`chromium-mobile`, iPhone 13 viewport) via a temporary scratch spec (removed after the run): after tapping/clicking a slot, `getComputedStyle(slot, '::after')` reported `zIndex: "40"` and `pointerEvents: "none"` in both the timed grid and the days-only grid, and `document.elementFromPoint` at the slot center still hit the slot element, confirming interaction is unchanged.
- The scratch spec was verification-only and not added to the tracked suite; a tracked regression spec can be added as follow-up work if desired.

## Risks / notes

- At layer 40 the cursor now paints above the sticky weekday header (z-10) and in-grid overlay/event blocks; this is the intended outcome of matching the tooltip layer.
<!-- SECTION:FINAL_SUMMARY:END -->
