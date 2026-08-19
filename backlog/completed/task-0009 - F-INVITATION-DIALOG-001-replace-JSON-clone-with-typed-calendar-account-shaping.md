---
id: TASK-0009
title: >-
  F-INVITATION-DIALOG-001: replace JSON clone with typed calendar-account
  shaping
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-INVITATION-DIALOG-001.md
priority: low
type: task
ordinal: 9000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/groups/InvitationDialog.vue

Problem: Calendar account data is still deep-cloned with JSON.parse(JSON.stringify(...)) during mount.

Why it matters: Untyped cloning hides shape assumptions and conflicts with the repo's explicit boundary-model guidance.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Replace ad hoc cloning with explicit typed shaping
- [ ] #2 Preserve current dialog behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:58
---
**Verification evidence:**
- Reconfirmed that `src/components/groups/InvitationDialog.vue` cloned `authUser.calendarAccounts` with `JSON.parse(JSON.stringify(...))` during `onMounted` before this change.
- Replaced the ad hoc clone with the typed `cloneCalendarAccounts` boundary in `src/components/settings/useCalendarAccountsState.ts`, which preserves dialog-local editability while normalizing required `enabled` ownership for the local response payload.
- Added focused coverage in `src/components/groups/InvitationDialog.test.ts` for deep-enough clone isolation and response submission without mutating the auth-owned calendar state.
- Ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` in `frontend/` on 2026-05-24; all checks passed.

**Implementation notes:**
- Kept the dialog-local editable account shape explicit so the response payload path still receives required `enabled` and `subCalendars` values without widening the shared frontend user model.
---
<!-- COMMENTS:END -->
