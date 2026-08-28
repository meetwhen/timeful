---
id: TASK-0064
title: Add temporary functional-requirement candidates to migration inventory
status: Done
assignee:
  - '@OpenCode'
created_date: '2026-08-25 08:15'
updated_date: '2026-08-25 08:21'
labels:
  - requirements
  - migration
  - traceability
dependencies: []
references:
  - docs/requirements/functional-requirements.md
  - docs/requirements/migration/backlog-fr-inventory.md
documentation:
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/requirements/migration/README.md
  - docs/terminology/README.md
modified_files:
  - docs/requirements/migration/README.md
  - docs/requirements/migration/backlog-fr-inventory.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-200.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-201.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-202.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-203.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-204.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-205.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-206.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-207.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-208.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-209.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-210.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-211.md
priority: medium
type: docs
ordinal: 66000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Capture every entry in docs/requirements/functional-requirements.md as a durable, non-normative CAND record. Preserve the source text and the user's clarified unstaged wording, distinguish confirmed existing requirements from new candidates and bugs, and keep the migration inventory's source rules accurate.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Each temporary functional-requirements source section has one new sequential candidate record
- [x] #2 Each record preserves source wording and provides the required candidate review fields
- [x] #3 Candidate behavior uses the clarified unstaged wording where it is complete and is independently reviewable
- [x] #4 Confirmed overlap with existing requirements and non-requirement bug material are classified and dispositioned accurately
- [x] #5 The migration guide and consolidated index support and link the new source-backed candidates
- [x] #6 Documentation validation confirms required fields numbering and local links
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Update the migration guide so candidate Source references may cite either the completed Backlog or a temporary requirements migration source.
2. Add CAND-200 through CAND-211 for the twelve sections in functional-requirements.md, preserving each section verbatim and applying clarified wording as review-authored candidate behavior when complete.
3. Classify existing overlaps, the FR-010 defect report, and unresolved product choices without creating canonical requirements.
4. Add the candidate links to the consolidated inventory and validate sequential numbering, required fields, controlled values, Markdown links, formatting, and whitespace.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Added CAND-200 through CAND-211 from every section of the temporary functional-requirements source. FR-010 remains bug provenance; CAND-206 records the unresolved hide-versus-disable product choice; CAND-202 and CAND-207 map to FR-013 and proposed FR-074 respectively. Corrected CAND-208 after the source removed its incomplete draft fragment, retaining the pointer-based scheduling-drag tooltip behavior.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added twelve temporary functional-requirement candidate records, CAND-200 through CAND-211, and indexed them under a dedicated inventory section. Updated the migration guide so candidate provenance may cite a temporary requirements migration source as well as completed Backlog material. Mapped confirmed existing behavior to FR-013 and proposed FR-074, retained FR-010 as bug provenance, and recorded the Edit availability presentation choice as unresolved. Validation passed: candidate numbering and required fields, git diff --check, Prettier, and Markdown link checks. Unit and E2E suites were exempt because the change is documentation-only.
<!-- SECTION:FINAL_SUMMARY:END -->
