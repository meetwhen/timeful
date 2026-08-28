---
id: TASK-0081
title: Renumber canonical quality requirements
status: Done
assignee: []
created_date: '2026-08-26 15:50'
updated_date: '2026-08-26 15:51'
labels:
  - requirements
  - quality
  - migration
dependencies: []
references:
  - docs/requirements/README.md
  - docs/requirements/migration/README.md
documentation:
  - docs/requirements/quality/README.md
  - docs/requirements/migration/README.md
modified_files:
  - docs/requirements/README.md
  - docs/requirements/quality-requirements.md
  - docs/requirements/quality/qr/
  - docs/requirements/migration/backlog-fr-inventory.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/
  - docs/design/architecture/README.md
  - docs/design/architecture/adr/ADR-008.md
  - docs/design/architecture/adr/ADR-009.md
  - docs/design/architecture/adr/ADR-010.md
  - prettier/markdown/sentences-per-line.test.mjs
  - scripts/markdown.mjs
priority: medium
type: docs
ordinal: 88000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Make canonical quality requirement identifiers contiguous from QR-001, preserve the retired pre-canonical quality notes as untriaged migration candidates, and update active traceability references without rewriting historical Backlog records.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Canonical quality requirements are numbered QR-001 through QR-011 with matching filenames and metadata
- [x] #2 The three retired quality notes are retained as untriaged CAND records in the migration inventory
- [x] #3 Active documentation and tests reference the renumbered canonical QR records
- [x] #4 The retired quality-requirements source is deleted and Markdown tooling tolerates its working-tree deletion
- [x] #5 Markdown lint format checks and Markdown tests pass
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Preserve the three retired quality notes as separate untriaged CAND records and add them to the migration inventory.
2. Delete the retired source and remove its canonical-guide reference.
3. Renumber canonical QR files, IDs, headings, and active traceability references from QR-004 through QR-014 to QR-001 through QR-011.
4. Update the Markdown formatter test and ensure the Markdown file collector ignores tracked files deleted from the working tree.
5. Verify contiguous canonical filenames, absence of stale active references, Markdown checks, tests, and diff whitespace.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
The implementation was completed before this task was created at the user's request. The retired source was preserved as CAND-212 through CAND-214 with needs-decision triage metadata. Historical Backlog records remain unchanged because they are managed historical evidence rather than active traceability. Verification confirmed the QR-001 through QR-011 file set, no stale active QR-012 or higher references, no active quality-requirements.md references, and a clean diff check.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Renumbered canonical quality requirements from QR-004 through QR-014 to contiguous QR-001 through QR-011, including filenames, IDs, headings, canonical index entries, active ADR traceability, migration-candidate references, and the formatter test description. Preserved the retired pre-canonical quality notes as untriaged CAND-212 through CAND-214 records and removed the retired source file. Updated Markdown tooling to skip tracked files deleted from the working tree. Verified with npm run lint:markdown, npm run format:markdown:check, npm run format:markdown, npm run test:markdown-rules, npm run test:markdown-format, stale-reference searches, and git diff --check. Unit and e2e tests were exempt for this documentation-only change.
<!-- SECTION:FINAL_SUMMARY:END -->
