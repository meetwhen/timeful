---
id: TASK-0132
title: >-
  Fix event-page e2e specs that still expect the removed inline Add description
  button
status: Done
assignee: []
created_date: '2026-09-02 09:39'
updated_date: '2026-09-02 10:25'
labels:
  - e2e
  - frontend
  - test-maintenance
dependencies: []
references:
  - TASK-0134
priority: high
type: bug
ordinal: 145300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The e2e suite is red: event-page specs still wait for the removed inline "+ Add description" trigger, hitting the default 30s per-test timeout and cascading serial skips.

Contract (user-confirmed): the event page renders a saved description read-only; there is no Add description button on mobile or desktop; descriptions are edited only in the create and edit forms. The app already implements this (Event.vue renders EventDescription only when event.description?.trim(); EventDescription.vue is read-only; the inline trigger and editor were removed in f6f65e7d). No Add description button code remains in frontend/src - the stale behavior lives only in the specs.

Stale specs found by running npm run test:e2e (5 failed, 17 skipped, 11 did not run, 50 passed in 5.1m):
- frontend/e2e/event-page-no-responses-layout.spec.ts:11 - desktop branch waits for getByRole("button", { name: /^\+\s*add description$/i }) (30s timeout); its unconditional mobile branch also expects #desktop-primary-availability-btn, which the current mobile layout (bottom action bar, #mobile-primary-availability-btn) does not render.
- frontend/e2e/event-page-days-only-layout.spec.ts:55 - clicks the removed button and expects a [role="textbox"] description editor (30s timeout).
- frontend/e2e/event-description-header-layout.spec.ts:39 - fails on chromium-mobile because it unconditionally reads #desktop-schedule-event-btn, which is gated to !isPhone in Event.vue.
Also observed once: schedule-overlap-mobile-touch-firefox.spec.ts:146 tooltip anchoring predicate failure (5s) - assess flakiness by rerunning in isolation; do not fix in this task unless reproducible.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 No tracked e2e spec waits for, clicks, or asserts an Add description button on the event page; the asserted contract is that a saved description renders read-only and descriptions are edited only in the create and edit forms.
- [x] #2 event-page-no-responses-layout.spec.ts passes on chromium-desktop without the Add description button, keeps the shared right action-column alignment assertions, and skips chromium-mobile as a desktop-only layout like the other tests in the file.
- [x] #3 event-page-days-only-layout.spec.ts keeps the inline Start on Monday alignment contract on chromium-desktop without the removed inline description-editor flow.
- [x] #4 event-description-header-layout.spec.ts no-description test gates the desktop-only Schedule event alignment check so it passes on both chromium projects.
- [x] #5 The affected specs pass on their configured projects and the required frontend checks (lint, typecheck, build, test:unit) pass.
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
1. event-page-no-responses-layout.spec.ts: delete the addDescriptionButton locator, addDescriptionBox destructure, null-check entry, and the [addDescriptionBox, scheduleEventBox] row pairing; keep the scheduleEventBox width/x alignment with the shared right action column; add test.skip(chromium-mobile, "Desktop-only header layout") like the sibling tests because its unconditional branch asserts desktop-only selectors (#desktop-primary-availability-btn, #event-header-actions .desktop-primary-availability-anchor) that the mobile bottom-bar layout does not render.
2. event-page-days-only-layout.spec.ts: in the chromium-desktop branch drop the + Add description click, the [role=textbox] expectation, and the description-vs-Schedule-event overlap block; keep the Add availability vs Start on Monday center alignment assertions.
3. event-description-header-layout.spec.ts: in the no-description test, accept testInfo and run the #desktop-schedule-event-btn alignment evaluate() only on chromium-desktop; keep the mobile-valid negative assertions (no shell, no add-description button, no header description) running everywhere.
4. Verify: run the three specs via npm run test:e2e -- --project=chromium-desktop (and chromium-mobile for the description spec), rerun schedule-overlap-mobile-touch-firefox once to assess the tooltip flake; then npm run lint / typecheck / build / test:unit.
5. Report the legacy-code finding: no Add description button code remains in frontend/src (removed in f6f65e7d); only the specs referenced it.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Confirmed the :449 root cause in code: Event.vue mounts ScheduleOverlap via defineAsyncComponent behind v-if="scheduleOverlapReady", and useEventEditing.addAvailability() silently returns when the scheduleOverlap template ref is still null (frontend/src/composables/event/useEventEditing.ts:63-65). Specs that click #desktop-primary-availability-btn right after waitForEventShell can land in that window; the identical flow passes in isolation, matching the race hypothesis from the prior session.

Fix (e2e-only, per the confirmed contract that the app is already correct): added waitForScheduleOverlapMounted(page) to frontend/e2e/helpers/timed-event-helpers.ts, polling for .schedule-overlap-layout, a class rendered only by ScheduleOverlap.vue; its DOM presence proves the async component mounted and the template ref is assigned. Applied it before all three click-after-open sites: event-page-days-only-layout.spec.ts (edit-availability click and add-availability click) and event-page-no-responses-layout.spec.ts:259, which shared the same latent race.

Deleted the temporary frontend/e2e/__debug-days-only.spec.ts; its diagnosis is recorded here and in the prior handoff.

Verification after the mount-wait fix: npm run test:e2e -- --project=chromium-desktop --project=chromium-mobile over event-page-no-responses-layout.spec.ts, event-page-days-only-layout.spec.ts, and event-description-header-layout.spec.ts passed 20 tests with 10 intentional chromium-mobile skips in 1.1m, including the previously failing dates-only add-availability test. Required checks pass: npm run lint, npm run typecheck, npm run test:unit (138 files, 985 tests), npm run build.

User decision on the reproducible schedule-overlap-mobile-touch-firefox.spec.ts:146 tooltip failure (failed 3 of 3 runs including in isolation on the final worktree): keep it out of TASK-0132 and file a follow-up. Created TASK-0134 (bug, Medium) with the failure evidence and diagnosis direction pointing at shared Tooltip repositioning and TASK-0033 as a possible regression origin.

format:markdown run for DoD #4: it only surfaced pre-existing table-column drift in docs/requirements/README.md, unrelated to this task, and was reverted; no Markdown changed by this task lies outside Backlog management (the formatter excludes backlog/**).
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Fixed the event-page e2e specs that still expected the removed inline Add description button, under the user-confirmed contract: a saved description renders read-only, there is no Add description button on mobile or desktop, and descriptions are edited only in the create and edit forms. The app needed no changes; no Add description button code exists in frontend/src (removed in f6f65e7d) — only the specs still referenced it.

Spec changes:

- event-page-no-responses-layout.spec.ts: removed the Add description button locator, its box destructure, null-check, and row pairing; kept the shared right action-column alignment assertions; added the chromium-mobile skip ("Desktop-only header layout") matching the sibling tests because the unconditional branch asserts desktop-only selectors the mobile bottom-bar layout does not render.
- event-page-days-only-layout.spec.ts: removed the inline description-editor flow (button click plus [role="textbox"] editor assertions) while keeping the inline Start on Monday vs Add availability alignment contract.
- event-description-header-layout.spec.ts: gated the desktop-only Schedule event alignment check to chromium-desktop via testInfo; the mobile-valid negative assertions (no description shell, no Add description button, no header description) still run on both chromium projects.

Unmasked failure fixed (e2e-only): with the stale timeout gone, a pre-existing race surfaced at event-page-days-only-layout.spec.ts:449 — clicking #desktop-primary-availability-btn right after waitForEventShell could land while Event.vue's async ScheduleOverlap (defineAsyncComponent behind v-if="scheduleOverlapReady") had not mounted, and useEventEditing.addAvailability() silently no-ops when its scheduleOverlap template ref is null (frontend/src/composables/event/useEventEditing.ts:63). Added waitForScheduleOverlapMounted(page) to frontend/e2e/helpers/timed-event-helpers.ts, which deterministically polls for .schedule-overlap-layout (a class rendered only by ScheduleOverlap.vue; its DOM presence proves the template ref is assigned), and applied it before all three click-after-open sites across the two specs.

Verification:

- npm run test:e2e -- --project=chromium-desktop --project=chromium-mobile over the three affected specs: 20 passed, 10 skipped (intentional chromium-mobile skips) in 1.1m, including the previously failing dates-only add-availability test.
- npm run lint, npm run typecheck, npm run test:unit (138 files, 985 tests), and npm run build all pass.

Out of scope per user decision: the reproducible schedule-overlap-mobile-touch-firefox.spec.ts:146 tooltip re-anchoring failure on firefox-touch (pre-existing; failed 3 of 3 runs including isolation) is tracked in follow-up TASK-0134 with the recorded evidence and diagnosis direction.
<!-- SECTION:FINAL_SUMMARY:END -->
