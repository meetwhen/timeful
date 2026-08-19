---
id: TASK-0020
title: 'F-SCHEDULE-OVERLAP-002: split monolithic schedule-overlap view'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-002.md
priority: medium
type: task
ordinal: 20000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/schedule_overlap/ScheduleOverlap.vue

Problem: The schedule-overlap view is still monolithic, with a large computed surface and broad coordination responsibilities.

Why it matters: Even when behavior is correct, the component is expensive to change safely because unrelated concerns share one file.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Split the component along stable domain or UI boundaries without changing behavior
- [ ] #2 Keep extracted logic covered by targeted tests
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
- Firefox inspection of `event-respondents-panel` passed against the frontend on `http://127.0.0.1:4173`.

**Implementation notes:**
- `ScheduleOverlap.vue` now acts as the route-facing shell while extracted helpers own persistence preferences and view-model assembly.
- `src/components/schedule_overlap/useScheduleOverlapViewModels.ts` holds the sidebar, mobile overlay, tool-row, and grid view-model composition so the shell no longer owns every computed presentation boundary inline.
---
<!-- COMMENTS:END -->
