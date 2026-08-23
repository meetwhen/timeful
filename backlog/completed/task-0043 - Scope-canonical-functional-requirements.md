---
id: TASK-0043
title: Scope canonical functional requirements
status: Done
assignee:
  - OpenCode
created_date: '2026-08-21 16:53'
updated_date: '2026-08-21 16:55'
labels: []
dependencies: []
references:
  - docs/requirements/README.md
  - docs/requirements/functional/
  - backlog/backlog.md
  - >-
    backlog/tasks/task-0042.04 -
    Migrate-next-canonical-functional-requirements.md
priority: medium
type: docs
ordinal: 35000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Make the canonical functional requirement records independently understandable by adding the page, event-kind, interaction-state, and entry-point applicability context that was lost during migration from the backlog and prior frontend behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 FR-011 and FR-016 distinguish their applicable timed-grid and scheduling modes
- [x] #2 FR-033 and FR-034 identify the affected frontend entry points or banner context
- [x] #3 FR-042 FR-044 and FR-048 identify their relevant event-page state and control scope
- [x] #4 FR-049 distinguishes timed and dates-only mobile control layouts including response state
- [x] #5 Requirement Markdown formatting and the changed-file diff are valid
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Update FR-011 and FR-016 with the event-page timed-grid and scheduling-mode scope established by the backlog and schedule presentation change.
2. Update FR-033 through FR-034 and FR-042 through FR-049 with their originating route, event kind, response state, permission state, and entry-point scope.
3. Run Prettier on the eight requirement files, inspect the diff, and record verification.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Scoped the eight migrated records using their originating backlog and frontend behavior. No application code changed; `npm exec prettier -- --check` passed for all eight records, and `git diff --check` passed. Unit and E2E suites were not run because the change is documentation-only.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Scoped eight canonical functional requirements so each identifies its applicable page, event kind, interaction state, permission state, or UI entry-point coverage. FR-011/016 now exclude unrelated editor modes; FR-033/034 define feedback and banner scope; FR-042/044/048 define their event-page controls; and FR-049 distinguishes response, timed-event, and dates-only layouts. Verified with scoped Prettier and `git diff --check`; application unit and E2E suites were not run because no application code changed.
<!-- SECTION:FINAL_SUMMARY:END -->
