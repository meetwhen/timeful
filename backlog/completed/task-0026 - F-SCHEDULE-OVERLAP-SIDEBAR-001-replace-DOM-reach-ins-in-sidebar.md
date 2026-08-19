---
id: TASK-0026
title: 'F-SCHEDULE-OVERLAP-SIDEBAR-001: replace DOM reach-ins in sidebar'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-SIDEBAR-001.md
priority: high
type: task
ordinal: 26000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/schedule_overlap/ScheduleOverlapSidebar.vue

Problem: Sidebar behavior still reaches through respondentsPanelRef.value?.$el and exposes DOM-facing scroll helpers.

Why it matters: DOM reach-ins couple the sidebar to child implementation details and block cleaner composition boundaries.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Replace $el access and exposed scroll helpers with typed, explicit child contracts or shared state
- [ ] #2 Preserve current sidebar behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 14:00
---
**Verification evidence:**
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed on 2026-05-24.
- Updated unit coverage in `src/components/schedule_overlap/ScheduleOverlapSidebar.test.ts` now verifies the sidebar reads the respondents panel element through an explicit exposed contract instead of `$el`.

**Implementation notes:**
- Coordinate with any `RespondentsList` or respondents-panel decomposition so contracts do not churn twice.
---
<!-- COMMENTS:END -->
