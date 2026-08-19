---
id: TASK-0021
title: 'F-SCHEDULE-OVERLAP-003: base timed-grid rows on displayed local minutes'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-003.md
priority: medium
type: task
ordinal: 21000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/composables/schedule_overlap/useCalendarGrid.ts

Problem: The timed-grid row model mixed displayed-clock ordering with day-index shifting, filler rows, and split-specific rendering rules, so wrapped timezone cases could produce fake discontinuities and brittle row ownership.

Why it matters: For timezone-shifted timed events, the migrated grid could render a structural split even when the displayed local-day timeline was continuous, and the old split mapping made border rendering, disabled rows, and per-column date ownership harder to reason about safely.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Base timed-grid rows on displayed local clock minutes plus per-local-day interval ownership
- [ ] #2 Collapse wrapped ranges whenever the displayed local-day segments overlap or touch
- [ ] #3 Keep a split only for a real hidden gap
- [ ] #4 Preserve per-column local-date ownership
- [ ] #5 Avoid duplicate displayed times or inconsistent row separators
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed in `frontend/`.
- Added targeted regressions in `src/composables/schedule_overlap/useCalendarGrid.test.ts`, `src/components/schedule_overlap/ScheduleOverlap.test.ts`, and `src/components/schedule_overlap/scheduleOverlapRendering.test.ts` for the real-gap `UTC+3:30`, touching `UTC+4:00`, and overlapped `Asia/Kathmandu` cases.

**Implementation notes:**
- `useCalendarGrid` now carries `absoluteMinutes` and `displayedMinutes` on timed rows, resolves displayed rows through local-day interval ownership instead of the old first-split or second-split day-index adjustment, and treats wrapped ranges as one continuous displayed sequence when `displayStartMinutes <= wrappedEndMinutes`.
- `ScheduleOverlap.vue` no longer synthesizes filler rows between wrapped segments and instead inserts an explicit `split-gap` row only when a real second split exists.
- The rendering helpers now derive separator styling from displayed local minutes so enabled and disabled rows share the same dashed or solid boundary semantics in odd-offset zones such as Kathmandu. The remaining true-split case is the real-gap `UTC+3:30` shape where one local date owns `23:00 -> 24:00` and `00:00 -> 03:00` with a hidden middle interval.
- During this session we also confirmed the product-level rule: if the UI shows a continuous full-day axis with grey cells for non-owned times, timezone switching alone does not require a structural split. The manual event repro at `http://127.0.0.1:4173/e/ACbfC` in `UTC+4:00` behaved consistently with that rule and did not need a split block.
---
<!-- COMMENTS:END -->
