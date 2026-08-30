---
id: TASK-0097
title: Align desktop editing Delete availability button under right header column
status: Done
assignee:
  - opencode
created_date: '2026-08-28 21:18'
updated_date: '2026-08-28 21:22'
labels: []
milestone: Frontend polish
dependencies: []
priority: medium
type: bug
ordinal: 103000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On the desktop timed event response editing page, the Delete availability button currently renders in a 22rem header actions box at the left edge of its own header row, under the "Edit event" / "Copy link" buttons. While editing, the row's left description slot is never rendered (it requires !isEditing), so nothing pushes the Delete box into the right-hand column where Cancel/Save and Overlay availability / More options live.

Required outcome: when the Delete availability action is shown in the desktop editing header, its container must be aligned to the right edge of the header row so the Delete button lines up under the other right-column header buttons, matching their column width.

Constraints:
- Use a layout-based fix (e.g. flex auto margin) consistent with how the schedule event buttons achieve right alignment in the same row; no absolute positioning or pixel hacks.
- Do not alter non-editing header rows, the mobile editing bottom bar (which intentionally places Delete before Cancel/Save), or the delete availability confirmation dialog.
- Update the existing Event view unit test that asserts the exact delete actions class list, and add/extend regression coverage for the right alignment.
- Frontend required checks must pass.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On the desktop timed event response editing page, the Delete availability button is right-aligned in the event header so it sits in the right-hand column directly under the Cancel/Save row and the Overlay availability / More options row, matching the 22rem header actions column width.
- [x] #2 The alignment fix is layout-based (no visual hacks) and does not change the non-editing header rows, the schedule event buttons, or the mobile editing bottom bar.
- [x] #3 Regression coverage at the unit test layer asserts the delete actions container is right-aligned in the desktop editing header, and existing Event view tests are updated where they assert the old class list.
- [x] #4 npm run lint, npm run typecheck, npm run build, and npm run test:unit pass in frontend/.
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
1. In frontend/src/views/Event.vue (~line 665), add `sm:tw-ml-auto` to the desktop delete actions container (`desktop-editing-delete-actions desktop-event-header-actions tw-flex tw-justify-end`) so the fixed-width 22rem box is pushed to the right edge of its `sm:tw-flex-row` header row, matching the schedule event button pattern (`sm:tw-ml-auto` at line ~685). The delete container only renders while editing (`isEditing && showDeleteAvailabilityAction && !isPhone && !isGroup`), where the description slot is never rendered, so the auto margin has no effect on other rows.
2. Update frontend/src/views/Event.test.ts (~line 2669) source assertion for the new class list and extend it to assert right alignment (`sm:tw-ml-auto`) on the delete actions container.
3. Run frontend required checks: lint, typecheck, build, test:unit.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Fix applied at Event.vue:665: `sm:tw-ml-auto` on the delete actions container mirrors the schedule-event-button right-alignment pattern in the same row. Delete container renders only while editing, where the description slot never renders, so the auto margin cannot affect other rows.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Aligned the desktop editing Delete availability button with the right-hand header column on the timed event response editing page.

## Change

- `frontend/src/views/Event.vue`: added `sm:tw-ml-auto` to the `.desktop-editing-delete-actions` container (Event.vue:665). The container has the shared 22rem `.desktop-event-header-actions` width but is the sole child of its `sm:tw-flex-row` header row while editing (the row's description slot requires `!isEditing`), so it previously sat at the left edge under "Edit event" / "Copy link". The auto margin pushes it to the right edge, lining the Delete button up under Cancel/Save and Overlay availability / More options, matching the existing schedule-event-button alignment pattern in the same row. The container only renders when `!isPhone && !isGroup && isEditing && showDeleteAvailabilityAction`, so no other header row, the mobile bottom bar, or the confirmation dialog is affected.

## Verification

- `frontend/src/views/Event.test.ts`: updated the desktop editing header test to assert the rendered delete actions container has `sm:tw-ml-auto` and updated its source-level class-list assertion.
- Required checks all pass in `frontend/`: `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit` (135 files / 932 tests).

## Notes

- No e2e tests exist for this layout state and none were required by the task scope; regression coverage was added at the unit test layer per the work order.
<!-- SECTION:FINAL_SUMMARY:END -->
