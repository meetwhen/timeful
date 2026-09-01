---
id: TASK-0119
title: Restyle response editing page delete and cancel buttons
status: Done
assignee: []
created_date: '2026-08-31 12:08'
updated_date: '2026-08-31 12:12'
labels:
  - ui-styling
dependencies: []
priority: low
type: enhancement
ordinal: 126300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On the event response editing page (`frontend/src/views/Event.vue`), restyle two action buttons.

Desktop: the Delete availability button in the desktop editing header actions (currently `variant="outlined"` `color="error"` at Event.vue ~line 649) should become a solid red filled button. Follow existing Vuetify 3 conventions (explicit variant props, `color="error"` or the project's red token).

Mobile: the Cancel button in the mobile editing action bar (currently `variant="outlined"` with `tw-border-green tw-text-green` at Event.vue ~line 932) should be red instead of green (outlined red border and red text). The project Tailwind config defines `red: #DB1616` (tailwind.config.cjs), and `tw-text-red`/`tw-border-red` are already used elsewhere.

Keep the mobile Save button green and all click behaviors unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On the desktop response editing page, the Delete availability button (currently `desktop-delete-availability-btn` in Event.vue, outlined error) renders as a solid red filled button with no outline variant.
- [x] #2 On the mobile response editing page, the Cancel button renders with a red border and red text instead of green.
- [x] #3 Delete, Save, and Cancel behaviors and other editing buttons are unchanged.
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
1. Desktop delete button (Event.vue ~649): use `variant="flat"` with `color="error"` — solid red fill with no outline and no elevation shadow, per Vuetify 3 explicit-variant conventions used elsewhere in the codebase.
2. Mobile Cancel button (Event.vue ~932): swap `tw-border-green tw-text-green` for `tw-border-red tw-text-red` (red token `#DB1616` already in tailwind.config.cjs).
3. Run lint, typecheck, build, unit tests.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
User refinement: the desktop Delete button must have no shadow. Used explicit `variant="flat"` with `color="error"` (solid fill, no elevation) instead of relying on the default elevated variant; this matches existing `variant="flat"` usage in the codebase (Event.vue:968, NewEvent.vue, NewGroup.vue, NewSignUp.vue).
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Restyled the response editing page action buttons in `frontend/src/views/Event.vue`:

- Desktop Delete availability button (`desktop-delete-availability-btn`) now uses `variant="flat"` with `color="error"` — a solid red filled button with no outline and no elevation shadow.
- Mobile editing Cancel button now uses `tw-border-red tw-text-red` (project red token `#DB1616`) instead of `tw-border-green tw-text-green`; it remains `variant="outlined"`.

No behavior changes: click handlers, disabled states, and the mobile Save button (green) are unchanged.

Checks: `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` all pass (137 files, 959 tests). E2E was not run; the styling change does not affect any e2e assertions (the mobile editing spec asserts button visibility and click behavior only, not colors). Visual verification is available at the local debug entry points.
<!-- SECTION:FINAL_SUMMARY:END -->
