---
id: TASK-0047
title: Migrate clarified functional requirements and ADR terminology
status: Done
assignee:
  - OpenCode
created_date: '2026-08-22 14:23'
updated_date: '2026-08-22 14:30'
labels:
  - requirements
  - adr
  - terminology
dependencies: []
references:
  - backlog/backlog.md
  - frontend/adr/010-frontend-event-ownership-semantics.md
  - frontend/adr/011-frontend-guest-response-ownership-semantics.md
  - frontend/adr/012-frontend-timed-event-instant-slot-model.md
documentation:
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/terminology/README.md
  - docs/terminology/glossary.md
priority: high
type: docs
ordinal: 39000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Document the agreed timed-event mode, response-access, cross-device event-sign-in, active-slot validation, and respondent-name behaviors as canonical requirements. Align the controlled glossary and affected frontend ADRs so the documents express one consistent product model.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 New functional requirements document the agreed range and specific-times picked-date behavior and conversion behavior
- [x] #2 New functional requirements document active-slot validation and respondent-name validation behavior
- [x] #3 New functional requirements document protected-by-default and explicitly-open response access including event sign-in transfer
- [x] #4 The glossary defines response author protected response and open response
- [x] #5 ADR-010 ADR-011 and ADR-012 align with the clarified requirement and glossary semantics
- [x] #6 The functional-requirements index links every new record and documentation validation passes
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add controlled response-access terms to the glossary, linked to ADR-011.
2. Add atomic FR-050+ records and index entries for the agreed timed-mode, validation, response-access, and event-sign-in behavior.
3. Revise ADR-012 to distinguish range and specific-times modes and their conversions; revise ADR-011 and ADR-010 for browser-local-to-authenticated event sign-in.
4. Validate YAML front matter, index coverage, anchors, and Markdown links. Run documentation-focused checks; full frontend suites are not proportionate to documentation-only changes.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Added FR-050 through FR-065, glossary response-access terms, and coordinated ADR-010/011/012 updates. Structural validation verified 65 contiguous FR files, required front matter, index rows, and glossary anchors. `git diff --check` passed. `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed (135 files, 934 tests). Firefox E2E was explicitly waived by the user for this documentation-only task; the attempted run did not start because another process was already listening on port 4174 (PID 394633), which was left untouched.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Migrated 16 agreed functional requirements (FR-050 through FR-065) covering mutually exclusive range/specific-times behavior, conversions, timed-event persistence and validation, respondent-name rules, response access, and cross-device event sign-in. Added controlled glossary definitions for response author, protected response, and open response. Updated ADR-010, ADR-011, and ADR-012 to align ownership, response access, and timed-mode semantics with the new requirements. Validation: requirement/index/glossary structural check, `git diff --check`, lint, typecheck, production build, and 934 unit tests passed. Firefox E2E was waived for the docs-only change after its startup found port 4174 in use; no existing process was stopped.
<!-- SECTION:FINAL_SUMMARY:END -->
