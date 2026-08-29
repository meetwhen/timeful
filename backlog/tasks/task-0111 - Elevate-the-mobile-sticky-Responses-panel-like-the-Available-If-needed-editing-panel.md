---
id: TASK-0111
title: >-
  Elevate the mobile sticky Responses panel like the Available/If needed editing
  panel
status: Done
assignee:
  - opencode
created_date: '2026-08-29 15:24'
updated_date: '2026-08-29 15:38'
labels: []
milestone: Mobile UX
dependencies: []
modified_files:
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts
priority: medium
type: enhancement
ordinal: 117000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On phones, the event page shows the Responses panel in the fixed bottom overlay. The Available/If needed editing panel and the mobile action bar already cast an upper-edge shadow via the shared mobile elevated-panel treatment, but the sticky Responses panel section in the mobile overlay renders as a flat white block, so the grid behind it reads as flush with the panel. Apply the same shared upper-edge elevation to the sticky Responses panel so all stacked mobile bottom panels share one consistent elevated look.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On a phone-sized viewport, when the sticky Responses panel is shown on the event page, its top edge casts the same upward shadow as the Available/If needed editing panel and the mobile action bar, using the shared mobile elevated-panel treatment
- [x] #2 The elevation is applied at the shared component level (ScheduleOverlapMobileOverlay sticky respondents section), not duplicated at call sites
- [x] #3 A unit regression test asserts the sticky respondents panel uses the shared mobile elevated-panel treatment when shown and not while editing
- [x] #4 No regression to the hidden-while-editing behavior of the sticky Responses panel
- [x] #5 The Available/If needed editing panel's upper-edge shadow is likewise visible while its expand transition slides (scope addition approved by the user in-session)
- [x] #6 Unit regression tests assert both mobile panels render the shared elevated treatment directly on their expand-transition target, with no intermediate wrapper that would clip the upward shadow during the slide
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. In `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue`, remove the intermediate wrapper divs for both elevated sections and put `timeful-mobile-elevated-panel tw-p-4` directly on the `v-if` element inside each `v-expand-transition` (sticky respondents section and Available/If needed editing panel). Rationale: the transition sets `overflow: hidden` on the target during the slide, clipping a nested panel's upward shadow; the target's own box-shadow is never clipped by its own overflow, so the shadow stays visible while sliding.
2. Update `ScheduleOverlapMobileOverlay.test.ts`: assert both panels use `.timeful-mobile-elevated-panel`, that the elevated element is the direct slot child of its expand-transition wrapper (no intermediate wrapper), and keep the hidden-while-editing coverage.
3. Run `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit` in `frontend/`.
4. Finalize the task with evidence and mark Done.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
User follow-up: the upper-edge shadow only appeared after the panel finished sliding. Root cause: v-expand-transition sets overflow:hidden on the transition target during enter/leave (vuetify expand-transition.js), which clips the nested elevated panel's upward box-shadow until overflow is restored on completion. Fix: render the shared elevated treatment directly on the transition target, since an element's own overflow:hidden does not clip its own box-shadow. User approved extending this fix to the Available/If needed editing panel as well.

Required checks all pass: lint, typecheck, build, test:unit (965 tests / 139 files). No Markdown files changed, so no format:markdown run needed. Graphify graph updated after the code change.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added the shared mobile upper-edge elevation to the sticky Responses panel and made the elevation visible during expand-transition slides for both mobile bottom panels.

Root cause of the slide-time shadow pop-in: `v-expand-transition` sets `overflow: hidden` on its transition target while animating (node_modules/vuetify/lib/components/transitions/expand-transition.js:19,44) and restores it on completion (line 69). With the elevated panel nested inside that target, the upward box-shadow was clipped during the slide and only rendered after the slide finished.

Changes:
- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue`: removed the intermediate wrapper divs so `timeful-mobile-elevated-panel tw-p-4` sits directly on the `v-if` element inside both `v-expand-transition` blocks (sticky Responses section and Available/If needed editing panel). An element's own `overflow: hidden` never clips its own box-shadow, so the shadow is now visible throughout the slide; final rendering is unchanged because the same shared treatment (surface token, `0 -2px 8px` shadow, bottom-safe clip-path) applies to the same box.
- `frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts`: new test asserts the sticky respondents panel uses the shared elevated treatment and is shown only when `showStickyRespondents` is true and not editing; both elevated-panel tests assert the panel sits directly on the expand-transition target with no intermediate wrapper (regression guard for the slide-time clipping); hidden-while-editing coverage retained.

Verification: frontend `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` all pass (965 unit tests across 139 files, including the 12 in ScheduleOverlapMobileOverlay.test.ts). No Markdown files were changed, so the format:markdown DoD item has nothing to format.
<!-- SECTION:FINAL_SUMMARY:END -->
