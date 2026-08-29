---
id: TASK-0099
title: >-
  Elevate mobile Event Response Editing bottom panel to match Event page bottom
  panel
status: Done
assignee:
  - opencode
created_date: '2026-08-29 09:49'
updated_date: '2026-08-29 10:11'
labels:
  - mobile
  - ui
  - design-tokens
dependencies: []
priority: medium
type: enhancement
ordinal: 105000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On mobile, the bottom panel of the Event Response Editing page (the panel with the Available and If needed actions) currently lacks the elevated surface treatment used by the bottom panel on the Event page. Visually the two mobile bottom panels should match: the response-editing panel should sit above the page content with the same elevation (shadow/surface treatment) as the Event page bottom panel.

Outcome: unify the elevation treatment of the two mobile bottom panels, preferably by extracting the elevation value into a shared constant or design token that both panels reference, so they stay consistent going forward.

Scope is frontend-only and mobile viewports. The worker should locate the current elevation styling of the Event page bottom panel and the Event Response Editing page bottom panel, then align the latter with the former.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On mobile viewports, the bottom action panel of the Event Response Editing page (Available / If needed actions) renders with the same elevation as the bottom panel on the Event page
- [x] #2 Both panels consume a single shared elevation token or constant so the two surfaces cannot drift apart in future changes
- [x] #3 The elevated panel keeps its current layout, spacing, and behavior on mobile, with no visual regressions to content underneath or safe-area handling
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

- Event page mobile bottom panel: `.mobile-event-action-bar` (frontend/src/views/Event.vue:917, scoped rule at Event.vue:2581) carries `background-color: var(--timeful-mobile-action-bar-surface)` and `box-shadow: 0 -2px 8px var(--timeful-mobile-action-bar-shadow)`; tokens live in `frontend/src/index.css:44-45`.
- Timed Event Response Editing Page bottom panel: `ScheduleOverlapMobileOverlay.vue` renders a fixed bottom overlay (`z-60`) above the editing action bar (Delete / Cancel / Save). The Available / If needed toggle (`AvailabilityTypeToggle`) sits in a `tw-bg-white tw-p-4` container with no elevation today.
- The toggle panel and the editing action bar read as one merged white sheet (confirmed by user screenshot). User decision: elevating just the top edge of the Available / If needed panel is sufficient; do not try to separate or individually elevate the stacked sections.

## Implementation plan

1. `frontend/src/index.css`: introduce a shared elevation class `timeful-mobile-elevated-panel` (surface + `0 -2px 8px` top-edge shadow) built on the existing `--timeful-mobile-action-bar-*` tokens, following the existing `.timeful-elevated-button` shared-class precedent. Keep tokens unchanged.
2. `frontend/src/views/Event.vue`: add `timeful-mobile-elevated-panel` to the mobile action bar element and remove the now-duplicated `.mobile-event-action-bar` background/shadow rule.
3. `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue`: apply `timeful-mobile-elevated-panel` to the editing panel container (the one holding `AvailabilityTypeToggle`) so its top edge gets the same elevation.
4. `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts`: add a unit assertion that the editing panel renders with the shared elevation class.
5. Run `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit` in `frontend/`.

## Validation

- AC #1/#2 evidence: both panels consume the single shared class referencing the same tokens; unit test asserts the editing panel carries it.
- AC #3 evidence: box-shadow-only change; layout, spacing, and safe-area handling untouched; full required-check suite passes.

Adjustment (user-directed): elevation must be top-edge-only. The box-shadow's natural bottom bleed painted over the action bar below the editing panel. Added `clip-path: inset(-16px -16px 0 -16px)` to the shared `.timeful-mobile-elevated-panel` class so the shadow renders only above (and beside) the panel and never over stacked sections below. Geometry (`0 -2px 8px` + token) is unchanged for both panels.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
User refinement after first implementation: the shadow on the Available / If needed panel bled downward over the Delete / Cancel / Save action bar, breaking the seamlessly-stacked single-sheet look. Fix: the shared `.timeful-mobile-elevated-panel` class now clips its own shadow at the element bottom edge (`clip-path: inset(-16px -16px 0 -16px)`), so the elevation is top-edge-only. Both panels keep the identical shared class, so the Event page bar look is unchanged while the editing panel no longer casts onto the bar below it.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## What changed

- `frontend/src/index.css`: introduced the shared elevation class `.timeful-mobile-elevated-panel` (surface `--timeful-mobile-action-bar-surface` + `box-shadow: 0 -2px 8px var(--timeful-mobile-action-bar-shadow)` + `clip-path: inset(-16px -16px 0 -16px)`). The clip makes the elevation top-edge-only so stacked bottom panels stay seamlessly merged instead of casting shadows onto each other.
- `frontend/src/views/Event.vue`: the mobile action bar now consumes the shared class; the duplicated `.mobile-event-action-bar` background/shadow rule was removed. Its rendered look is unchanged.
- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue`: the Timed Event Response Editing Page bottom panel (Available / If needed toggle container) now uses the shared class, replacing plain `tw-bg-white`, so its top edge carries the same elevation as the Event page action bar.
- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts`: added a unit test asserting the editing panel renders with the shared elevation class and contains the availability toggle.

## Verification

- Unit: `npm run test:unit` — 135 files / 935 tests passed, including the new overlay elevation test.
- Targeted e2e (isolated test stack): `chromium-mobile` `event-mobile-editing-options.spec.ts` — 2 passed.
- Browser computed-style check (throwaway spec, removed after run): while mobile editing, the editing panel and the action bar compute the identical box-shadow, and the editing panel's `clip-path` is applied (top-edge-only elevation, no seam shadow over the Delete / Cancel / Save bar).
- `npm run lint`, `npm run typecheck`, `npm run build` all pass.
- Full e2e suite was not run; verification focused on the affected mobile editing flow plus the computed-style check.

## Risks / follow-ups

- `clip-path` makes each elevated panel its own containing block for fixed descendants; none exist in the affected panels (Vuetify menus teleport to body), so no behavior change is expected.
- Future mobile bottom surfaces should reuse `.timeful-mobile-elevated-panel` rather than re-declaring shadows.
<!-- SECTION:FINAL_SUMMARY:END -->
