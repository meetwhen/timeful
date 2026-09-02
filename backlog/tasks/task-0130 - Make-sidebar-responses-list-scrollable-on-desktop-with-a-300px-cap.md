---
id: TASK-0130
title: Make sidebar responses list scrollable on desktop with a 300px cap
status: Done
assignee:
  - opencode
created_date: '2026-09-02 08:26'
updated_date: '2026-09-02 08:49'
labels: []
dependencies: []
priority: medium
type: enhancement
ordinal: 143300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On desktop, the responses list in the schedule-overlap sidebar does not scroll at a useful point: it uses a viewport-derived max height with a 400px minimum (frontend/src/components/schedule_overlap/useRespondentsListSizing.ts), so the list can grow to nearly the full viewport and the desktop scroll behavior does not engage as intended. On mobile, the list is already capped (ScheduleOverlapMobileOverlay passes maxHeight=240).

Desired outcome: on desktop, cap the sidebar responses list at 300px; once content exceeds 300px, the list scrolls. Keep the mobile behavior unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On desktop (non-phone), the responses list in the schedule-overlap sidebar is capped at 300px regardless of viewport height
- [x] #2 When responses exceed the cap, the list scrolls vertically and overflow is contained inside the scroll region; the Responses heading, pending list, and legend remain outside the scroll region
- [x] #3 When responses fit within the cap, no scroll region appears and the surrounding layout is unchanged
- [x] #4 The mobile overlay respondents list keeps its existing behavior (240px max-height prop path)
- [x] #5 Unit tests cover the 300px desktop cap and that scrolling engages past the cap, added per the repo regression-test default
- [x] #6 Required frontend checks pass: lint, typecheck, build, and unit tests
- [x] #7 On desktop, when the respondents list is scrolled down, a white fade gradient renders below the Responses heading over the clipped top item, mirroring the existing bottom gradient; both gradients track scroll position (top fades only when scrolled away from the top, bottom fades only when content remains below)
- [x] #8 Existing OverflowGradient consumers (NewEvent, NewSignUp, SignUpBlocksList) keep their current bottom-gradient behavior
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
1. Root cause: the desktop cap from `useRespondentsListSizing` (viewport-derived, 400px floor) is applied to the outer `scrollableSection` in `RespondentsList.vue`, not to the `respondentsScrollView` scroll element, so desktop scrolling never engages. The mobile path (240px `maxHeight` prop) already caps the scroll element correctly.

2. Apply the proven mobile mechanism to desktop with a fixed 300px cap:
   - Delete `useRespondentsListSizing.ts` (its viewport-derived sizing is fully replaced by the fixed cap; only `RespondentsList.vue` uses it).
   - In `RespondentsList.vue`: add a `scrollViewMaxHeight` computed = `maxHeight` prop when set, else 300 on desktop, undefined on phone. Keep a local `hasMounted` mount flag for the OverflowGradient gate.
   - Template: move the inline `max-height: 300px !important` to the `respondentsScrollView` (same element that owns `overflow-y-auto`); remove the outer `scrollableSection` max-height binding and its now-unused ref. Heading, pending list, and legend stay outside the scroll region. Keep OverflowGradient desktop gating as-is.

3. Tests (RespondentsList.test.ts): replace the "desktop viewport-derived cap on scrollableSection" test with one asserting the desktop scroll view carries `max-height: 300px !important` plus `tw-overflow-y-auto`/`tw-overflow-x-hidden` while the outer section has no inline max-height; keep the existing mobile 240px-prop test as the unchanged-behavior guard.

4. Verify: `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit`; run `graphify update .` after code changes.

User-directed scope additions during review: (1) skip e2e; (2) add a white fade gradient at the top of the desktop respondents list (below the Responses heading) so the clipped item is softened, mirroring the existing bottom OverflowGradient.

Extend OverflowGradient.vue with a `position: top | bottom` prop (default bottom, so NewEvent/NewSignUp/SignUpBlocksList are unaffected): top variant anchors top-0 with an inverted white gradient and shows when scrollTop > 1; bottom variant unchanged. Render a second OverflowGradient (position=top) in RespondentsList inside the same overflow-hidden wrapper, gated like the bottom one (desktop only).

Tests: extend OverflowGradient.test.ts for the top-variant show/hide and anchoring; assert the RespondentsList template wires the top gradient via the existing raw-source assertion pattern (OverflowGradient is null-stubbed there).
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Executed per plan: cap moved onto the scroll element; composable deleted (single consumer, viewport logic fully replaced). `hasMounted` kept locally to preserve the OverflowGradient mount gate.

Mobile 240px prop path untouched; ScheduleOverlapMobileOverlay tests pass. Other OverflowGradient consumers unaffected via default position="bottom" (SignUpBlocksList tests pass).

Top gradient show condition mirrors the bottom one: scrollTop > 1 with a 1px epsilon, driven by the existing scroll/resize/mutation observers.

E2E skipped per user direction; aborted run's isolated test stack torn down. Required checks: lint, typecheck, build, unit (983 tests) all green.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-09-02 08:49
---
E2E was explicitly skipped at user direction mid-run; the isolated test stack it had started was torn down (`compose.test.yaml` stack down -v, dev stack untouched). DoD item #3 therefore remains unchecked by exception, not by failure.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Desktop sidebar responses list is now capped at 300px and scrolls past the cap, with fade gradients at both scroll edges.

What changed:

RespondentsList.vue: the desktop cap moved from the outer `scrollableSection` to the `respondentsScrollView` scroll element itself (the previous viewport-derived cap sat on a non-scrolling ancestor, so desktop scrolling never engaged). A `scrollViewMaxHeight` computed applies the `maxHeight` prop when present (mobile overlay, 240px unchanged) and otherwise caps desktop at 300px (undefined on phone). Outer section max-height binding and ref removed; heading, pending list, and legend stay outside the scroll region.

useRespondentsListSizing.ts deleted: its viewport-derived sizing (400px floor, resize listener) was fully replaced by the fixed cap; it had a single consumer.

OverflowGradient.vue: new optional `position: "top" | "bottom"` prop (default `bottom`); the top variant anchors `top-0`, inverts the white fade, and shows only when `scrollTop > 1`. RespondentsList renders a second gradient with `position="top"` so a clipped top item fades under the Responses heading.

Tests:

RespondentsList.test.ts: new desktop test asserts the 300px inline cap plus `overflow-y-auto`/`overflow-x-hidden` on the scroll view with no inline max-height on the outer section, and the heading outside the scroll region; the mobile 240px-prop test remains as the unchanged-behavior guard; new source test asserts both gradients are wired to the scroll view.

OverflowGradient.test.ts: new tests for bottom-anchored default and top-variant show/hide by scroll position and anchoring/gradient direction.

Checks: lint, typecheck, build, and unit tests (138 files / 983 tests) all pass. E2E skipped at user direction. Risk: none known; the mobile path and other OverflowGradient consumers (NewEvent, NewSignUp, SignUpBlocksList) keep their previous behavior.
<!-- SECTION:FINAL_SUMMARY:END -->
