---
id: TASK-0040
title: Hide mobile Responses panel while editing availability
status: Done
assignee: []
created_date: '2026-08-21 12:06'
updated_date: '2026-08-21 12:09'
labels:
  - mobile
  - availability
dependencies: []
references:
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts
modified_files:
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts
priority: medium
type: bug
ordinal: 28000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On a phone, selecting a timeslot shows the Responses panel. If the respondent then enters Edit availability and selects another grid timeslot, the panel must show only the availability controls rather than stale response details.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The mobile overlay hides the Responses panel while editing availability
- [x] #2 The Available and If needed controls remain visible while editing availability
- [x] #3 Regression coverage verifies an active selected slot cannot render Responses during availability editing
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All unit tests pass
- [x] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Guard the mobile sticky respondents section with the existing editing view-model flag.
2. Add a component regression test covering editing with a selected slot.
3. Run the focused test and required frontend checks.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Made the sticky mobile respondents section conditional on not editing. The existing broad rendering fixture now represents the editing state without sticky respondents; a dedicated regression test verifies sticky respondents are suppressed while the availability-type control remains rendered.

Verification passed: focused ScheduleOverlapMobileOverlay unit test, npm run lint, npm run typecheck, npm run build, npm run test:unit (135 files, 934 tests), and Firefox scoped E2E (6 tests).
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Hide the mobile Responses panel whenever availability editing is active, preventing stale response details from appearing alongside the Available and If needed selector after a grid selection. Added focused component coverage for the conflicting editing-plus-selected-slot state and clarified the existing rendering fixture. Verified with lint, typecheck, production build, the full unit suite, and the scoped Firefox mobile event E2E suite.
<!-- SECTION:FINAL_SUMMARY:END -->
