---
id: TASK-0024
title: 'F-SCHEDULE-OVERLAP-006: active slots drop out of timezone-shifted timed grid'
status: To Do
assignee: []
created_date: '2026-08-19 13:40'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-006.md
priority: medium
type: bug
ordinal: 24000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/composables/schedule_overlap/timedGridRendering.ts (likely)

Problem: When a UTC+4 timed specific-dates event is opened in a default-UTC display timezone, the slot at Jun 24 03:00 UTC+4 projects to Jun 23 23:00 UTC - one calendar day before the earliest membership date column in the grid - so that active slot is invisible. Instead of 4 white grid cells (2 per day x 2 days), the grid shows only 3.

Why it matters: Active slots at the early edge of a timezone-shifted timed event drop out of the grid entirely. Users see fewer active cells than they saved, which undermines trust in the editing flow.

Root cause: The grid builds date columns from membership dates (picked dates in the event timezone), then projects slot instants into the display timezone to assign rows and columns. If a slot's projected display-timezone local day does not match any membership-date column, the slot is dropped. The projection logic does not account for the membership-date boundary crossing caused by a positive UTC offset.

Reproduction:
- Seed: UTC+4 (Asia/Dubai), Jun 24-25 membership, 03:00-05:00 window, 60-min increments
- All 4 slots are both enabled and active
- Open edit dialog, enter specific-times grid in default display timezone (UTC)
- Grid shows 3 instead of 4 white cells
- Repo-tracked e2e test: e2e/timed-event-utc4-edit-firefox.spec.ts
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 When the display timezone causes a slot to project into a local-day column that differs from its membership date, that slot must still be rendered in the appropriate membership-date column, not omitted
- [ ] #2 All 4 active slots must appear in the grid regardless of display timezone
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->
