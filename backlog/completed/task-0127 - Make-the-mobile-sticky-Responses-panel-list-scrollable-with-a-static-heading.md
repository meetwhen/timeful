---
id: TASK-0127
title: Make the mobile sticky Responses panel list scrollable with a static heading
status: Done
assignee:
  - opencode
created_date: '2026-09-01 16:11'
updated_date: '2026-09-01 18:39'
labels:
  - frontend
  - mobile
  - layout
dependencies: []
references:
  - frontend/src/components/schedule_overlap/RespondentsList.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapRespondentsPanel.vue
  - >-
    commit e3b82e6d (introduced the overflow-hidden wrapper that broke the
    scroll chain)
documentation:
  - frontend/src/components/schedule_overlap/RespondentsList.test.ts
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts
  - frontend/e2e/schedule-overlap-mobile-touch-firefox.spec.ts
  - frontend/e2e/event-toolbar-mobile-layout.spec.ts
  - frontend/e2e/helpers/timed-event-helpers.ts
modified_files:
  - frontend/src/components/schedule_overlap/RespondentsList.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - frontend/src/components/schedule_overlap/RespondentsList.test.ts
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts
  - frontend/e2e/schedule-overlap-mobile-touch-firefox.spec.ts
priority: medium
type: bug
ordinal: 135300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On phones, the schedule-overlap sticky Responses panel (the fixed bottom panel shown when the page's Responses section is off-screen and a slot or respondent is selected) clips the respondent list at a fixed 100px cap, and the list cannot be scrolled, so responses beyond the first few are unreachable. The panel must let the user scroll through all responses while the "Responses" heading stays static.

Root cause (verified in this codebase): RespondentsList.vue applies the maxHeight cap to the outer scrollableSection flex container, but the element that actually scrolls (respondentsScrollView, which already gets tw-overflow-y-auto whenever maxHeight is set) sits nested inside the tw-relative tw-overflow-hidden wrapper introduced by commit e3b82e6d. That wrapper absorbs the flex shrink (overflow-hidden gives it an automatic min-height of 0) and clips the content, while the scroll element keeps height:auto, so its overflow-y-auto never engages and nothing scrolls.

Confirmed decisions for this work:
- The sticky panel list cap is 240px (replacing the current 100px) so the scroll window shows roughly 6-7 rows.
- The fix is layout-based: the cap must land on the scrolling element itself whenever the maxHeight prop is set. The desktop path (no maxHeight prop, non-phone, viewport-derived cap on scrollableSection) must remain behavior-identical.
- The "Responses" heading remains outside the capped scroll area so it stays static; no sticky/positioning hacks.
- Coverage is regression-first: add the unit and e2e assertions for the new layout behavior as part of this task, before or alongside the fix.
- Scope is the schedule-overlap respondents panel only. No sign-up list changes, no browser-plugin payload changes, and no transport, timezone, or Temporal boundary changes (ADR-001/002/004/005/006 are not implicated).

Follow-up: the NewEvent create-dialog e2e flake investigated and fixed while validating this task was moved to TASK-0128 with user approval; it is not part of this task's scope.

After code changes, run graphify update . to keep the knowledge graph current.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On a phone viewport, when a grid slot or respondent is selected while the page's Responses section is off-screen, the sticky Responses panel's respondent list scrolls vertically so every response is reachable.
- [x] #2 The "Responses" heading stays outside the capped scroll area and its rendered position does not move when the list is scrolled; no sticky/positioning hacks are added to keep it static.
- [x] #3 Unit test: whenever the maxHeight prop is set, the inline max-height cap and the overflow-y-auto overflow land on the respondents scroll element and the outer scrollableSection carries no inline max-height; the Responses heading is asserted to be outside the capped scroll element.
- [x] #4 Unit test: on a non-phone viewport without the maxHeight prop, the viewport-derived cap stays on scrollableSection and the scroll element gains no inline max-height, pinning the desktop path as unchanged.
- [x] #5 Unit test: the mobile overlay passes a 240px list cap to the sticky respondents panel, replacing the previous 100px cap.
- [x] #6 E2E (hasTouch-gated, following the existing mobile Responses touch specs): with multiple guest responses seeded via the API, tapping a slot shows the sticky panel whose list reports scrollHeight greater than clientHeight, scrolling it leaves the Responses heading bounding box unchanged, and the last seeded respondent becomes visible/reachable.
- [x] #7 Frontend required checks pass: npm run lint, npm run typecheck, npm run build, npm run test:unit; the touched Playwright spec passes via npm run test:e2e -- --project=firefox-touch (Playwright owns the isolated test stack).
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
1. RespondentsList.vue (layout fix, mobile maxHeight path only):
   - Move the inline max-height cap from the outer scrollableSection to the actual scroll element (respondentsScrollView) whenever the maxHeight prop is set, so overflow-y-auto engages and the tw-relative tw-overflow-hidden wrapper (from e3b82e6d) no longer absorbs the flex shrink and clips.
   - Keep the desktop path untouched: no maxHeight prop + non-phone still caps scrollableSection with the viewport-derived height and leaves the scroll element without an inline max-height.
   - Add data-testid="respondents-scrollable-section" and data-testid="respondents-scroll-view" for deterministic test/e2e selection (pattern already used in this component via data-testid="event-timezone").
   - The Responses heading row already sits outside scrollableSection, so it stays static; no sticky/positioning hacks.
2. ScheduleOverlapMobileOverlay.vue: change the sticky panel cap from :max-height="100" to :max-height="240".
3. Unit tests RespondentsList.test.ts:
   - AC3: with maxHeight set (phone), assert inline max-height and tw-overflow-y-auto land on the scroll element, scrollableSection has no inline max-height, and the Responses heading (tw-text-lg) is not contained in the capped scroll element.
   - AC4: non-phone without maxHeight, assert viewport-derived inline max-height stays on scrollableSection and the scroll element gains no inline max-height.
   - Add an optional maxHeight param to the mountRespondentsList helper.
4. Unit test ScheduleOverlapMobileOverlay.test.ts (AC5): with showStickyRespondents, assert ScheduleOverlapRespondentsPanel receives maxHeight 240.
5. E2E schedule-overlap-mobile-touch-firefox.spec.ts (AC6, hasTouch-gated by the file's beforeEach): seed a specific-dates event via seedCanonicalTimedEvent, POST 24 guest responses (Guest 01..Guest 24, zero-padded so alphabetical DOM order matches seeding order; slot taps leave curRespondents empty so ordering is pure firstName localeCompare), tap the 09:00 slot, then assert: scroll view scrollHeight > clientHeight and computed max-height 240px; Responses heading bounding box unchanged after scrolling the list to the bottom; the last seeded row (Guest 24) sits within the scroll viewport after scrolling (2-col grid on 375px keeps rows below sm:tw-block).
6. Run frontend required checks (lint, typecheck, build, test:unit) and the touched Playwright spec via npm run test:e2e -- --project=firefox-touch.
7. Run graphify update . at the end.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
SESSION LOG (fix implemented, checks done, e2e run blocked by an environment-level failure - details below). STATUS: code + unit tests + e2e spec are in the worktree, uncommitted; task stays In Progress.

IMPLEMENTATION DONE (uncommitted in worktree). (1) RespondentsList.vue: moved the maxHeight-prop cap off the outer scrollableSection onto the actual scroll element respondentsScrollView (new inline :style binding on the scroll element; scrollableSection style now only applies the viewport-derived cap when !maxHeight && !isPhone, so the desktop path is byte-for-byte behavior-identical). Added data-testid="respondents-scrollable-section" on scrollableSection and data-testid="respondents-scroll-view" on the respondentsScrollView div for deterministic test selection. The Responses heading row remains outside scrollableSection; no sticky/positioning hacks. (2) ScheduleOverlapMobileOverlay.vue: sticky panel cap changed from :max-height="100" to :max-height="240".

TESTS ADDED (uncommitted). RespondentsList.test.ts: mountRespondentsList helper gained an optional maxHeight param; two new tests - (a) with maxHeight 240 (phone): inline 'max-height: 240px !important;' + tw-overflow-y-auto/tw-overflow-x-hidden land on [data-testid=respondents-scroll-view], scrollableSection has no inline max-height, and the .tw-text-lg Responses heading is not contained in the scroll element or scrollableSection (AC3); (b) non-phone without maxHeight: scrollableSection style matches /max-height: \d+px !important;/, scroll element has no inline max-height, desktop overflow-y-auto pinned (AC4). ScheduleOverlapMobileOverlay.test.ts: new test asserting the sticky ScheduleOverlapRespondentsPanel receives props maxHeight === 240 when showStickyRespondents (AC5). frontend/e2e/schedule-overlap-mobile-touch-firefox.spec.ts: new hasTouch-gated test 'Responses panel list scrolls under a static Responses heading' (AC6) - seeds a specific-dates event via seedCanonicalTimedEvent, POSTs 24 guest responses (names 'Guest 01'..'Guest 24', zero-padded so alphabetical firstName order matches seeding order; slot taps leave curRespondents empty so DOM order is pure localeCompare), taps timeslot row 0 col 0, then asserts computed max-height 240px on [data-testid=respondents-scroll-view], polls scrollHeight - clientHeight > 0, records the Responses heading bounding box, sets scrollTop = scrollHeight, asserts heading x/y unchanged within 1px, and asserts the 'Guest 24' row bounding rect sits within the scroll view rect (reachable). Note: 24 guests in the 2-col phone grid (tw-grid-cols-2 below sm) give 12 rows ~ 344px > 240px cap. Imports in the spec now include buildSpecificDateSeed, openEventPage, seedCanonicalTimedEvent alongside the existing createSpecificTimesEventFromDialog (removing the old import broke lint/typecheck - keep all four).

CHECKS PASSED. npm run lint OK, npm run typecheck OK, npm run build OK, npm run test:unit OK (138 files / 977 tests, includes the 3 new unit tests; targeted run of the two touched files: 46/46).

E2E RUN RESULT (npm run test:e2e -- --project=firefox-touch): 1 passed / 5 failed in 3.2m. PASSED: 'disabled mobile Schedule keeps the scheduled-event blue at full opacity' (the only test that does not depend on the seeded/created event page grid persisting). FAILED: the four pre-existing touch tests AND the new AC6 test, all failing after createSpecificTimesEventFromDialog/openEventPage rendered the landing page instead of the event page. Full symptom detail for that run was environmental (cold first-run stack) and was superseded by the flake investigation recorded in TASK-0128; this task's own layout diff was never implicated (unit/build/typecheck layers were fully green and the passing test exercised the same event-page grid).

RESUME CHECKLIST FOR NEXT SESSION. Step 1: re-run `npm run test:e2e -- --project=firefox-touch` once to see whether the landing-page failures reproduce. Step 2: if they reproduce, with TEST_DB_PERSIST=true keep the stack up and inspect server-test logs around the run. Step 3: once the AC6 test passes, AC1/AC2 get behavior-level evidence from it. Step 4: finish required checks if anything changed, run `graphify update .`, then finalize: verify each AC, write finalSummary, check DoD items, mark task Done.

FINAL SESSION LOG (scope split; flake work moved to TASK-0128). Re-running the firefox-touch suite reproduced the failures and they were root-caused to the NewEvent create-dialog flake (missing ExpandableSection import in NewEvent.vue plus the never-resolving ancestor-axis XPath and force-click animation race in timed-event-helpers.ts setSpecificTimesEnabled). The user approved expanding scope to fix it; that investigation, fix, and validation evidence now live in TASK-0128 ('Fix NewEvent create-dialog e2e flake: import ExpandableSection and make the specific-times switch toggle deterministic'), which also carries the modified files frontend/src/components/NewEvent.vue, frontend/src/components/NewEvent.test.ts, and frontend/e2e/helpers/timed-event-helpers.ts. With the TASK-0128 fix in the worktree, this task's validation stands: npm run lint, typecheck, build OK; npm run test:unit 978/978 including the three unit tests added for AC3/AC4/AC5; npm run test:e2e -- --project=firefox-touch passed 6/6 twice consecutively (58.7s and 1.0m), including the AC6 test 'Responses panel list scrolls under a static Responses heading' (behavior evidence for AC1/AC2 and the 240px cap: scrollHeight > clientHeight, Responses heading bounding box unchanged within 1px after scrolling, Guest 24 row reachable). graphify update . run; temporary debug spec deleted; repo-root npm run format:markdown run.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Fixed the mobile sticky Responses panel: the respondent list now scrolls under a static "Responses" heading. In RespondentsList.vue the inline max-height cap moved from the outer scrollableSection to the actual scroll element (respondentsScrollView) whenever the maxHeight prop is set, so overflow-y-auto engages instead of being defeated by the e3b82e6d overflow-hidden wrapper; the desktop (no-prop, non-phone) path is unchanged, and the heading row stays outside the capped area with no sticky hacks. ScheduleOverlapMobileOverlay.vue raised the sticky panel cap from 100px to 240px. Regression coverage: RespondentsList.test.ts pins the phone cap on the scroll element (with overflow classes, no inline cap on scrollableSection, heading outside) and the unchanged desktop path; ScheduleOverlapMobileOverlay.test.ts asserts the 240px cap; the new hasTouch e2e test "Responses panel list scrolls under a static Responses heading" seeds 24 guest responses and verifies scrollHeight > clientHeight, the Responses heading bounding box unchanged within 1px after scrolling to the bottom, and the last row (Guest 24) reachable. The NewEvent create-dialog e2e flake encountered while validating this task was root-caused and fixed under the separate TASK-0128 (user-approved scope split); with that fix in place, npm run lint, typecheck, build, and test:unit (978/978) all pass, and npm run test:e2e -- --project=firefox-touch passed 6/6 twice consecutively, including the AC6 test and all four previously-flaky dialog tests. graphify update . run.
<!-- SECTION:FINAL_SUMMARY:END -->
