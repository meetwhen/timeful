---
id: TASK-0025
title: >-
  F-SCHEDULE-OVERLAP-007: drags landing on inactive cells and blocks spanning
  collapsed rows
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-007.md
priority: medium
type: bug
ordinal: 25000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/composables/schedule_overlap/useDragPaint.ts, src/components/schedule_overlap/ScheduleOverlap.vue

Problem: Schedule-event drags could end in enabled but inactive specific-time cells, and timed blocks retained their full height when their range crossed a collapsed-hours row.

Why it matters: Drags could land on inactive cells and timed blocks rendered over collapsed rows, confusing which intervals are actually selectable.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Schedule-event drags must not end in enabled but inactive specific-time cells
- [ ] #2 Timed blocks whose range crosses a collapsed-hours row must render visible contiguous fragments
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
**Root cause:**
- Drag movement and pointer-up did not validate endpoints against active slots. Timed blocks only translated their top position to visible rows, leaving their base-grid duration height unchanged.

**Resolution:**
- Schedule drags retain their last active endpoint. Timed calendar and scheduled-event blocks now render visible contiguous fragments, skipping collapsed rows.

**Regression coverage:**
- `useDragPaint.test.ts` covers inactive endpoint rejection; `scheduleOverlapRendering.test.ts` covers block fragmentation around collapsed rows.
---
<!-- COMMENTS:END -->
