---
id: TASK-0065
title: Classify requirement migration candidates
status: Done
assignee:
  - OpenCode
created_date: '2026-08-25 08:51'
updated_date: '2026-08-25 09:13'
labels:
  - requirements
  - migration
  - triage
dependencies: []
references:
  - docs/requirements/migration/README.md
documentation:
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/requirements/quality/README.md
modified_files:
  - docs/requirements/migration/README.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/
priority: medium
type: docs
ordinal: 67000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Normalize the 211 durable CAND migration records with machine-readable triage verdicts that identify prospective FR/QR requirements, already-covered behavior, excluded source material, and product decisions. Preserve source provenance and make mappings to canonical requirements auditable.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Every CAND record has valid YAML front matter with a normalized verdict and confidence
- [x] #2 Prospective requirements state whether they are functional or quality requirements
- [x] #3 Covered candidates identify related canonical FR or QR records where known
- [x] #4 The migration inventory schema documents the front matter values and their relationship to the existing review fields
- [x] #5 A repository validation confirms all candidate files conform to the approved metadata schema
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Review the candidate schema and canonical FR/QR records, then validate candidate verdicts in parallel batches.
2. Normalize each candidate into YAML metadata: verdict, requirement_type when proposed, related_requirements, and confidence.
3. Document the metadata schema and its mapping from existing review fields in the migration guide.
4. Run a repository-wide YAML, canonical-reference, and prose-to-metadata consistency validation.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Normalized CAND-161 through CAND-180 with YAML triage metadata. Preserved Source, Candidate behavior, and Applicability; human-readable classifications and dispositions already agreed with the validated verdicts. Verified all 20 metadata blocks by scoped field inspection and `git diff --check`.

Applied normalized YAML triage metadata and synchronized review fields for CAND-141 through CAND-160 only. Verified all 20 front-matter blocks and classifications with targeted searches and `git diff --check`. No repository-wide schema validation was run because this request scoped changes to the selected candidates.

Reviewed all 211 candidates in eleven parallel batches against the canonical FR/QR inventory. Normalized 73 proposed requirements, 31 covered entries, 73 exclusions, and 34 product decisions. Corrected legacy prose where it contradicted a revised prospective-requirement verdict. The repository-wide Ruby validation parsed every metadata block, checked controlled values, verified related canonical IDs, and verified correspondence with the human-readable Classification and Confidence fields. `git diff --check` passed. This was documentation-only work, so frontend unit and e2e suites were exempt under project policy.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added normalized YAML triage metadata to all 211 CAND migration records and synchronized their human-readable review fields with the validated verdicts. Documented the schema and its mapping to the legacy Classification field in the migration guide. The inventory now contains 73 prospective requirements, 31 covered records linked to canonical FR/QR entries, 73 exclusions, and 34 items needing product decisions. Verified all front matter, controlled values, canonical references, and prose-to-metadata consistency with a repository-wide Ruby check; `git diff --check` passed. Frontend unit and e2e tests were not run because this is documentation-only work.
<!-- SECTION:FINAL_SUMMARY:END -->
