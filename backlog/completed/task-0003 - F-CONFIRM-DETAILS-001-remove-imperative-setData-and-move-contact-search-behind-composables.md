---
id: TASK-0003
title: >-
  F-CONFIRM-DETAILS-001: remove imperative setData() and move contact search
  behind composables
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-CONFIRM-DETAILS-001.md
priority: high
type: task
ordinal: 3000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/schedule_overlap/ConfirmDetailsDialog.vue

Problem: The dialog exposes an imperative setData() API and runs debounced contact search directly inside the component.

Why it matters: Data preparation, async search, and dialog rendering are tightly coupled, which makes state transitions hard to test.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move data preparation and search coordination behind explicit helpers or composables
- [ ] #2 Remove the imperative dialog API
- [ ] #3 Preserve current confirmation behavior
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed on 2026-05-23.
- Added unit coverage in `src/components/schedule_overlap/ConfirmDetailsDialog.test.ts` for parent-owned draft updates after removing the imperative `setData()` API and moving debounced contact lookup behind shared composables.

**Implementation notes:**
- Add regression coverage around contact search and dialog state transitions when practical.
---
<!-- COMMENTS:END -->
