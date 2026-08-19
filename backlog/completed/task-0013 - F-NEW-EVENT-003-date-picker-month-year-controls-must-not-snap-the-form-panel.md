---
id: TASK-0013
title: 'F-NEW-EVENT-003: date-picker month/year controls must not snap the form panel'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-NEW-EVENT-003.md
priority: medium
type: bug
ordinal: 13000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/DatePicker.vue

Problem: Clicking the migrated new-event date picker's month and year controls could move focus in a way that snapped the internal form panel back to the top.

Why it matters: The date picker sits inside a scrollable event editor. Losing the current panel position interrupts date selection and makes month or year navigation feel broken.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Clicking the date-picker month and year controls must keep the NewEvent form panel scroll position stable
- [ ] #2 Vuetify month and year navigation must still work
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:59
---
**Verification evidence:**
- Added focused control-button coverage in `src/components/DatePicker.test.ts` so the press event prevents the focus-scroll path without blocking the follow-up click.
- Added a Firefox regression in `e2e/timed-event-create-firefox.spec.ts` that asserts the new-event form panel keeps its scroll position when the month and year controls are clicked.
- Passed `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` in `frontend/`.

**Implementation notes:**
- The shared date picker now treats Vuetify control-button press events separately from date-cell interactions. Control presses call `preventDefault()` to avoid the focus-driven scroll jump, but they do not stop propagation, so Vuetify still receives the click and updates the calendar view.
---
<!-- COMMENTS:END -->
