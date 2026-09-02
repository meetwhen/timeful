---
id: TASK-0146
title: >-
  Flush mobile sticky Responses panel with the action bar and add the white
  overflow gradient
status: Done
assignee:
  - opencode
created_date: '2026-09-02 17:21'
updated_date: '2026-09-02 22:29'
labels:
  - frontend
  - mobile
dependencies: []
priority: medium
type: bug
ordinal: 159300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On phones, the sticky Responses panel (ScheduleOverlapMobileOverlay) renders with bottom padding that leaves a white strip between the respondents list and the fixed mobile action bar. The action bar's upward box-shadow paints over that strip, darkening it. Also, the respondents list in the sticky panel has no white overflow fade, unlike the desktop respondents list.

Outcome:
1. The sticky Responses panel no longer has bottom padding; the respondents list ends flush with the top edge of the mobile action bar.
2. The action bar's shadow must not paint over the sticky Responses panel (or the editing panel above it). Preferred mechanism: stack the fixed action bar wrapper below the z-60 mobile bottom overlay layer (e.g. z-50) so the overlay panels' white surfaces occlude the shadow, while the bar must still occlude the shared z-50 slot tooltip (TASK-0143 contract).
3. The respondents list inside the sticky panel uses the shared white OverflowGradient fades (top and bottom) like the desktop sidebar respondents list.

Constraints:
- Do not weaken the TASK-0143 stacking contract: the mobile action bar must still hit-test above the shared tooltip (e2e schedule-overlap-mobile-touch-firefox.spec.ts "mobile action bar stacks above an overlapping mobile tooltip").
- Keep layout-based fixes; reuse the shared OverflowGradient component and existing timeful-* tokens rather than adding one-off styles.
- Do not change desktop respondents list behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The sticky mobile Responses panel has no bottom padding, so the respondents list ends flush with the mobile action bar
- [x] #2 The mobile action bar's shadow is not visible over the sticky Responses panel or the editing panel surfaces above it
- [x] #3 The action bar still occludes the shared slot tooltip when they overlap (TASK-0143 contract holds)
- [x] #4 The sticky respondents list renders the shared white OverflowGradient fades when its scroll view overflows, and no gradient when it does not
- [x] #5 Desktop respondents list gradient behavior is unchanged
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
1. ScheduleOverlapMobileOverlay.vue: change the sticky respondents panel wrapper from `timeful-mobile-elevated-panel tw-p-4` to `tw-px-4 tw-pt-4` (drop bottom padding only; keep the editing panel untouched).
2. Event.vue: replace `timeful-bottom-overlay-layer` (z-60) on the fixed action bar wrapper with `tw-z-50`, so the z-60 mobile bottom overlay panels (white surfaces) paint above the bar and occlude its upward shadow. The bar still paints above the shared z-50 tooltip because it comes later in the DOM at the same z-level, preserving the TASK-0143 hit-test contract (verified by the firefox-touch e2e test).
3. RespondentsList.vue: relax both OverflowGradient v-ifs from `hasMounted && !isPhone && respondentsScrollView && !maxHeight` to `hasMounted && respondentsScrollView && scrollViewMaxHeight`, so the phone sticky list (maxHeight 240) gets the same white fades as the desktop 300px-capped list, while the uncapped phone list keeps none.
4. Tests: update Event.test.ts mobile footer stacking assertions (tw-z-50, not timeful-bottom-overlay-layer); add sticky-panel padding assertions in ScheduleOverlapMobileOverlay.test.ts; add a RespondentsList.test.ts case rendering the gradients for the capped phone list.
5. Checks: npm run lint / fmt:check / typecheck / build / test:unit; run the firefox-touch e2e stacking spec for the tooltip contract.

Revision (user-directed): keep .timeful-action-bar-layer { z-index: 50 } on the fixed action bar wrapper (Event.vue) instead of the tw-z-50 utility, and lower the shared tooltip below the bar with a new semantic layer .timeful-tooltip-layer { z-index: 40 } (index.css) applied in Tooltip.vue. This restores the TASK-0143 bar-over-tooltip contract by explicit z-index rather than DOM order, while keeping the z-60 bottom overlay panels above the bar so their white surfaces occlude its upward shadow.

Update tooltip selectors in ScheduleOverlap.mobileTooltip.test.ts and e2e/schedule-overlap-mobile-touch-firefox.spec.ts from .tw-fixed.tw-z-50 to .tw-fixed.timeful-tooltip-layer (the old selector became ambiguous once the bar stopped being z-60).
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Design note: the bar uses the new semantic `.timeful-action-bar-layer` (z-50) instead of the suggested raw `tw-z-50` utility because unit and e2e tooltip tests select the shared tooltip via `.tw-fixed.tw-z-50`; a raw utility on the bar wrapper made that selector ambiguous and broke the e2e strict-mode locators.

Pre-existing failure noted outside this scope: e2e firefox-touch "touching a timeslot keeps its mobile tooltip anchored while scrolling" (spec line 224) fails on clean main as well; baseline confirmed by stashing this task's changes and rerunning the spec.

User revised the stacking mechanism mid-review: keep the action bar in a semantic z-50 layer (timeful-action-bar-layer, already present in index.css) and lower the tooltip to an explicit z-40 instead of relying on DOM order at shared z-50. Added .timeful-tooltip-layer { z-index: 40 } to index.css and applied it in Tooltip.vue (replacing the tw-z-40 utility), giving the layer ladder tooltip 40 < action bar 50 < bottom overlay 60 = top navbar tw-z-[60].

Updated ScheduleOverlap.mobileTooltip.test.ts (7 selector swaps) and the firefox-touch e2e spec (8 selector swaps) from the ambiguous .tw-fixed.tw-z-50 to .tw-fixed.timeful-tooltip-layer; the old selector matched both tooltip and bar once the bar left z-60, causing a strict-mode failure in the bar/tooltip spec.

Definitive regression check via git stash: on the pre-change worktree the full firefox-touch spec gives the same result (6 pass; 'touching a timeslot keeps its mobile tooltip anchored while scrolling' fails identically), so that failure is pre-existing, not a TASK-0146 regression. Instrumented dump (temporary debug spec, since removed) identified two pre-existing root causes: (1) Tooltip.vue clamps horizontal position to the viewport margin (clamped left 119.917px vs slot center 115.75px at 375px viewport), so the spec's exact-equality predicate can never hold for this geometry; (2) after programmatic scroll Firefox synthesizes a mouseover under the last touch point and useTimedGridInteractions re-targets the tooltip to that other slot (content flipped 00:15 to 01:00 and the anchor followed the other slot).

Housekeeping: temporary diagnostic specs (session debug spec and the user's tmp-repro-highlight spec) were moved out of the linted tree to /tmp/opencode/task-0146-diagnostic-specs/ to keep lint/fmt/typecheck green; they are preserved there, not deleted.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Flush the mobile sticky Responses panel with the action bar and add the shared white overflow fades.

Changes:
- ScheduleOverlapMobileOverlay.vue: the sticky respondents panel wrapper now uses `tw-px-4 tw-pt-4` instead of `tw-p-4`, so the respondents list ends flush with the top edge of the mobile action bar (no white strip). The editing panel keeps `tw-p-4` (its bottom padding is white-on-white against the bar).
- Event.vue + index.css: the fixed mobile action bar wrapper moved from `timeful-bottom-overlay-layer` (z-60) to the semantic `timeful-action-bar-layer` (z-50), so the z-60 bottom overlay panels (sticky Responses panel, editing panel) paint above the bar and their white surfaces occlude the bar's upward box-shadow.
- Tooltip.vue + index.css: the shared tooltip moved from the raw `tw-z-50` utility (and an interim `tw-z-40` utility) to the semantic `timeful-tooltip-layer` (z-40). The final stacking ladder is explicit and DOM-order-independent: tooltip 40 < action bar 50 < bottom overlay panels 60 = top navbar `tw-z-[60]`. This strengthens the TASK-0143 contract: the bar hit-tests above the tooltip by explicit z-index, not by DOM order.
- RespondentsList.vue: both OverflowGradient fades now render whenever the scroll view is capped (`hasMounted && respondentsScrollView && scrollViewMaxHeight`) instead of desktop-only, so the phone sticky list (240px cap) gets the same white top/bottom fades as the desktop 300px-capped list, while the uncapped phone list still renders none.
- Selectors: ScheduleOverlap.mobileTooltip.test.ts (7 swaps) and e2e/schedule-overlap-mobile-touch-firefox.spec.ts (8 swaps) now target the tooltip via `.tw-fixed.timeful-tooltip-layer`; the old `.tw-fixed.tw-z-50` selector became ambiguous (matched tooltip and bar) once the bar left z-60.

Tests: Event.test.ts mobile footer stacking assertions updated to `.timeful-action-bar-layer` (not `timeful-bottom-overlay-layer`); ScheduleOverlapMobileOverlay.test.ts asserts the sticky panel has `tw-px-4`/`tw-pt-4` and no `tw-p-4`; RespondentsList.test.ts gained gradient cases for the capped phone list and the uncapped phone list.

Checks: lint (0 errors; 2 pre-existing vue/one-component-per-file warnings in NewEvent.test.ts), fmt:check, typecheck, build, test:unit (138 files, 990 tests), and format:markdown all pass. E2E: firefox-touch schedule-overlap-mobile-touch-firefox.spec.ts passes 6/7 in a real browser, including both stacking-contract tests ("Responses panel stacks above an overlapping mobile tooltip", "mobile action bar stacks above an overlapping mobile tooltip") and the sticky-list scroll test.

Known pre-existing failure (out of scope, not a regression): "touching a timeslot keeps its mobile tooltip anchored while scrolling" fails identically on clean main (verified via git stash baseline). Instrumented evidence shows two pre-existing root causes: (1) Tooltip.vue clamps horizontal position to the viewport margin (clamped 119.917px vs slot center 115.75px at 375px viewport), so the spec's exact-equality predicate can never hold for this geometry; (2) after programmatic scroll, Firefox synthesizes a mouseover under the last touch point and useTimedGridInteractions re-targets the tooltip to that other slot. Suggested follow-up task to fix the spec predicate and/or the touch mouseover re-targeting.
<!-- SECTION:FINAL_SUMMARY:END -->
