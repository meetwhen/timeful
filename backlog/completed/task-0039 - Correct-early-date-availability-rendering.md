---
id: TASK-0039
title: Correct early-date availability rendering
status: Done
assignee:
  - OpenCode
created_date: '2026-08-20 14:59'
updated_date: '2026-08-20 15:16'
labels: []
dependencies: []
references:
  - 'http://127.0.0.1:4173/e/WQ03VRJZ'
priority: high
type: bug
ordinal: 27000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Availability submitted across every selected event date must display consistently when returning to the event page.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 An availability response spanning August 20 through August 30 renders as available on every submitted date
- [x] #2 The event page does not show active submitted slots as unavailable for the first week only
- [x] #3 Regression coverage verifies the affected multi-week date range
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All unit tests pass
- [x] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Pass the grid's unpaged day/time slot resolver into availability aggregation. 2. Build response buckets using unpaged day indexes so the complete event range is stable regardless of the currently visible page. 3. Add a composable regression test with an 11-day event while page 2 is active. 4. Run the focused test and required frontend checks.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
The saved event embeds response metadata without slots, and the event page correctly fetches filtered slots from `/responses`. Current live API and client state both contain the full multi-week response.

Investigated the live event on 2026-08-20. `/api/events/WQ03VRJZ/responses?timeMin=2026-08-19T21:00:00Z&timeMax=2026-08-30T21:00:00Z` returns all 88 slots for Danila. Live Firefox inspection confirms the event-page `parsedResponses` and `responsesFormatted` mark 18:00-20:00 available for every date, including August 20-26 and August 27-30. No code change made because the reported state is not currently reproducible.

Fixed the page-relative aggregation bug: `getResponsesFormatted` now resolves every event-day bucket through the grid's unpaged `getDateFromDayTimeIndex` resolver. Added a regression covering an 11-day response refresh while page 2 is selected. Verification: focused unit test; `npm run lint`; `npm run typecheck`; `npm run build`; `npm run test:unit` (135 files, 933 tests); `npm run test:e2e -- --project=firefox-desktop` (22 passed, 1 PostgreSQL-only skipped).
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Availability response aggregation no longer depends on the currently displayed page. The complete event date range is resolved through the grid's unpaged day/time mapping, so a response refresh while page 2 is open cannot omit the first week. Added a regression for an 11-day event with page 2 selected, and updated test harnesses for the new resolver.

Verified with lint, typecheck, production build, 933 unit tests, and Firefox E2E (22 passed; 1 PostgreSQL-only test skipped).
<!-- SECTION:FINAL_SUMMARY:END -->
