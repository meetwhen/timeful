---
id: TASK-0063
title: Consolidate backlog candidate inventory index
status: Done
assignee:
  - OpenCode
created_date: '2026-08-25 07:58'
updated_date: '2026-08-25 08:00'
labels: []
dependencies: []
references:
  - docs/requirements/migration/backlog-fr-inventory.md
  - docs/requirements/migration/README.md
priority: low
type: docs
ordinal: 65000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Use one consolidated index for the completed-backlog candidate records. The existing individual CAND records remain durable review artifacts; remove the redundant numbered batch-index layer and align the migration guide with the single-index layout.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The completed-backlog candidate inventory has one consolidated index linking every CAND record
- [x] #2 The numbered batch-index documents are removed
- [x] #3 The migration guide describes the single-index layout without references to numbered batches
- [x] #4 Documentation formatting and local links validate
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Retain `backlog-fr-inventory.md` as the only navigation document for CAND-001 through CAND-199.
2. Remove the four redundant numbered batch indexes under `backlog-fr-inventory-batches/`.
3. Update `migration/README.md` to replace batch-based scope and hierarchy text with the single-index arrangement while preserving candidate schema and review guidance.
4. Validate the candidate-link set, check remaining references, run scoped Prettier, and run `git diff --check`.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Removed the four duplicate numbered batch indexes and updated the migration guide to make `backlog-fr-inventory.md` the sole candidate navigation document. Staged the resulting migration documentation changes.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Consolidated completed-backlog candidate navigation on `docs/requirements/migration/backlog-fr-inventory.md`, which links all 199 durable CAND records. Removed the four numbered batch-index documents and updated the migration guide to describe the single-index structure. Validation passed: scoped Prettier, candidate-link/file-count equality (199), staged migration `git diff --check`, and no remaining batch-index references under the migration documentation. Unit and E2E tests were exempt because this is documentation-only work. A repository-wide `git diff --check` remains blocked by pre-existing trailing whitespace in `docs/requirements/functional-requirements.md`, outside this task's scope.
<!-- SECTION:FINAL_SUMMARY:END -->
