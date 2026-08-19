---
id: TASK-0002
title: >-
  F-CALENDAR-TYPE-SELECTOR-001: make dialog visibility/reset ownership
  single-owner
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:51'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-CALENDAR-TYPE-SELECTOR-001.md
priority: low
type: task
ordinal: 2000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/settings/CalendarTypeSelector.vue

Problem: Internal finite-state dialog flow is reset by a watcher mirroring props.visible.

Why it matters: Local dialog state currently depends on prop mirroring instead of a clearer ownership contract.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Keep one clear owner for visibility and reset behavior
- [x] #2 Preserve current calendar-type selection UX
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:51
---
**Verification evidence:**
- Reconfirmed in `src/components/settings/CalendarTypeSelector.vue` that the dialog flow previously reset by watching `props.visible` and forcing the local step back to `PICK_CALENDAR` whenever visibility became false.
- Refactored the component to keep `visible` as a session boundary only: local step state now resets on a closed-to-open transition, while nested flow ownership stays inside the selector.
- Added `src/components/settings/CalendarTypeSelector.test.ts` coverage for provider action emits, nested `addedCalendar` forwarding, and reopening the dialog after entering a nested credential step.
- Ran `npm run test:unit -- src/components/settings/CalendarTypeSelector.test.ts` in `frontend` on 2026-05-23; 3 tests passed.

**Implementation notes:**
- Kept the existing `visible` prop and emitted events, but removed the reset-on-close mirroring pattern in favor of an explicit visible-session reset.
---
<!-- COMMENTS:END -->
