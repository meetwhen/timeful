---
id: TASK-0001
title: >-
  F-CALENDAR-ACCOUNTS-001: move account source selection/persistence behind
  explicit boundaries
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:51'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-CALENDAR-ACCOUNTS-001.md
priority: high
type: task
ordinal: 1000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/settings/CalendarAccounts.vue

Problem: Account-source selection, lazy event fetching, and collapse persistence are split across mount logic, a watcher, and localStorage.

Why it matters: Data ownership is unclear, and storage or auth concerns leak directly into the view component.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Move source selection and persistence behind explicit boundaries
- [x] #2 Keep event loading behavior intact
- [x] #3 Preserve the current settings UX
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:51
---
**Verification evidence:**
- Extracted source selection, lazy event loading, and collapse persistence into `src/components/settings/useCalendarAccountsState.ts`.
- Added regression coverage in `src/components/settings/useCalendarAccountsState.test.ts` for auth-owned versus event-owned account selection, `showCalendars` storage persistence, and lazy `/user/calendars` loading only when the caller does not provide a map.
- Passed `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` in `frontend/` on 2026-05-23.

**Implementation notes:**
- Preserved the migrated app's empty-map fallback behavior so bare settings and invitation usages continue to read from `authUser.calendarAccounts`, while schedule-overlap ownership stays with the injected event/group map.
---
<!-- COMMENTS:END -->
