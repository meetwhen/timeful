---
id: TASK-0019
title: 'F-SCHEDULE-OVERLAP-001: move storage-backed scheduling state behind boundaries'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-001.md
priority: high
type: task
ordinal: 19000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/schedule_overlap/ScheduleOverlap.vue

Problem: Core scheduling state is still seeded from localStorage, and the component exposes a large instance-style API for parent coordination.

Why it matters: Storage-coupled core state and imperative coordination make the main scheduling surface hard to test and risky to refactor.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move storage-backed state behind explicit domain or composable boundaries
- [ ] #2 Reduce imperative exposure
- [ ] #3 Preserve scheduling behavior and parity
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed in `frontend/`.
- Added `src/composables/schedule_overlap/scheduleOverlapStorage.test.ts`.
- Firefox inspection of `event-respondents-panel` passed against the frontend on `http://127.0.0.1:4173`.

**Implementation notes:**
- Remaining schedule-overlap storage ownership now goes through explicit helpers in `src/composables/schedule_overlap/scheduleOverlapStorage.ts` and `src/components/schedule_overlap/useScheduleOverlapPreferences.ts`, with shared guest-name reads and writes reused by `ScheduleOverlap`, `useAvailabilityData`, `useEventLoader`, and the event plugin bridge.
- The exposed `ScheduleOverlap` instance contract was reduced by dropping the `states` export and moving the state comparison to `Event.vue`.
---
<!-- COMMENTS:END -->
