---
id: TASK-0023
title: >-
  F-SCHEDULE-OVERLAP-005: shifted timed responses not bucketed onto rendered
  slots
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-005.md
priority: medium
type: bug
ordinal: 23000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/composables/schedule_overlap/useAvailabilityData.ts

Problem: Timed event heatmap aggregation rebuilt response buckets from day.dateObject.add(time.hoursOffset) instead of the rendered slot projection. For timezone-shifted timed grids that use absoluteMinutes or other display-domain normalization, fetched response slots were matched against the wrong instants, so the event page could show zero responded slots even when the response edit flow showed saved availability.

Why it matters: Responded availability disappeared from the main event page for shifted timed events, which broke best-times and heatmap trust for owners and guests.

Root cause: getResponsesFormatted() reconstructed timed slots from raw hoursOffset values, bypassing useCalendarGrid's rendered-slot normalization (including absoluteMinutes handling for timezone-shifted views), so response buckets landed on midnight-based instants instead of the visible 10:30-style slots.

Fix direction: Build timed response buckets by iterating visible row and column indexes and calling getDateFromRowCol(row, col) for each rendered slot instead of recomputing instants from hoursOffset.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 When a timed grid is rendered through a shifted local display domain, responsesFormatted must bucket fetched responses against the same slot instants produced by getDateFromRowCol(...)
- [ ] #2 Event-page heatmap cells and best-times counts must reflect the saved responses
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
- Added regression `"maps fetched timed responses onto rendered slots that use absolute local minutes"` in `src/utils/scheduleOverlap.regressions.test.ts`.
- Verified live against `http://127.0.0.1:4173/e/5B5af`: the page now shows four responded cells and no longer renders the `"There's no time when all 1 respondents are available."` note.
---
<!-- COMMENTS:END -->
