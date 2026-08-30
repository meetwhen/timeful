---
id: TASK-0109
title: >-
  Align mobile event action bar buttons with the panel above by increasing
  activity bar padding
status: Done
assignee:
  - opencode
created_date: '2026-08-29 15:00'
updated_date: '2026-08-29 15:10'
labels:
  - mobile
  - styling
dependencies: []
priority: medium
type: bug
ordinal: 115000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On phones, buttons inside the bottom mobile event action bar (`.mobile-event-action-bar` in `frontend/src/views/Event.vue`, e.g. the Schedule/Reschedule button, Delete on the response editing page, and the Save button) sit closer to the screen edges than buttons in the elevated panel above the bar (e.g. Calendar options, Available/if needed). The bar uses smaller horizontal padding on phones than the panel above it, so button edges do not line up: the Delete button's left edge is left of the Calendar options button's left edge, and Save's right edge is right of Available/if needed's right edge.

Increase the activity bar's horizontal padding on phones so action buttons align with the buttons above them. Use a layout-based fix (adjust the bar's own padding) rather than per-button overrides, and keep shared visual-state classes consistent with the elevated panel styling conventions.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On phones, the left edge of the lowest-priority action in the mobile event action bar (e.g. Delete on the response editing page) aligns with the left edge of buttons directly above it (e.g. Calendar options).
- [x] #2 On phones, the right edge of the primary action in the mobile event action bar (e.g. Save) aligns with the right edge of buttons directly above it (e.g. Available/if needed).
- [x] #3 Alignment holds for every state that renders the mobile event action bar (not editing, editing, and scheduling states).
- [x] #4 Desktop and larger-than-phone layouts are visually unchanged.
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
1. Align the action bar's horizontal padding with the elevated panel above it (ScheduleOverlapMobileOverlay.vue:34 uses `timeful-mobile-elevated-panel tw-p-4` at all breakpoints; the bar in Event.vue:917 uses `tw-px-4 max-sm:tw-px-2`).
2. Edit `frontend/src/views/Event.vue:917`: remove `max-sm:tw-px-2` so the bar uses `tw-px-4` on phones too; desktop is unchanged.
3. Update the source-assertion test "compacts the single-row mobile footer actions below the sm breakpoint" in `frontend/src/views/Event.test.ts:579` to reflect the new padding (keep the other compaction assertions: gap and text-size compaction remain).
4. Verify alignment across bar states (not editing / editing / scheduling) — the bar is a single container, so one padding fix covers all states; desktop layout untouched.
5. Run required checks: lint, typecheck, build, unit tests. No e2e scenario exists for pixel alignment; unit source assertions cover the layout contract (regression coverage at unit layer since browser screenshot diffing is not part of this repo's tracked checks — document in final summary).
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Root cause: the bar dropped to 8px horizontal padding on phones (`max-sm:tw-px-2`) while the editing panel above keeps 16px (`tw-p-4`), so bar buttons sat 8px closer to the screen edges than panel buttons.

Fix chosen over per-button overrides per styling rules: single-container padding change.

E2E regression assertion added to the existing guest editing scenario in event-mobile-editing-options.spec.ts (chromium-mobile, 375px viewport) because alignment is a rendered-layout property unit tests cannot prove; temporarily restoring the old padding reproduced the failure.

During regression verification a `git stash` pathspec mistake from the frontend workdir briefly attempted to pop an older unrelated stash entry; the pop aborted cleanly on conflict and all pre-existing stash entries remain intact.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Aligned mobile event action bar buttons with the elevated panel above them.

Changes:
- `frontend/src/views/Event.vue`: removed the phone-only `max-sm:tw-px-2` override from `.mobile-event-action-bar`, so the bar now uses `tw-px-4` (16px) on phones, matching the elevated editing panel above it (`ScheduleOverlapMobileOverlay.vue` uses `tw-p-4` at all breakpoints). This is a single-container layout fix; it applies to every bar state (not editing, editing, scheduling) and leaves desktop (`tw-px-4`) unchanged.
- `frontend/src/views/Event.test.ts`: updated the source-assertion test to require the bar's `tw-px-4` padding and forbid `max-sm:tw-px-2`; kept the still-valid inner compaction assertions (gap/text-size below sm).
- `frontend/e2e/event-mobile-editing-options.spec.ts`: added browser-level alignment assertions to the guest-response editing scenario (the screenshot scenario): the Delete button's left edge must align with the Calendar options button's left edge, and Save's right edge must align with the right edge of the Available/If needed row, within 1px.

Verification:
- Regression proven: with the old `max-sm:tw-px-2` temporarily restored, the new e2e alignment assertion fails (Delete left edge mismatch); with the fix it passes on the chromium-mobile (375px) project.
- `npm run lint`, `npm run typecheck`, `npm run build` pass.
- `npm run test:unit`: 959 tests pass (138 files).
- `npm run test:e2e -- e2e/event-mobile-editing-options.spec.ts`: 2 passed (chromium-mobile), 2 skipped (desktop-only).

Risks/follow-ups: the bar is 8px more inset on phones than before; the Schedule/Reschedule row and scheduling Cancel/Clear/Schedule row gain the same inset, keeping all bar states consistent. No Markdown files changed (DoD format item n/a).
<!-- SECTION:FINAL_SUMMARY:END -->
