---
id: TASK-0062
title: Create a completed-backlog FR extraction inventory
status: Done
assignee:
  - OpenCode
created_date: '2026-08-24 20:20'
updated_date: '2026-08-24 21:13'
labels:
  - requirements
  - traceability
  - provenance
  - backlog-migration
dependencies: []
references:
  - backlog/backlog.md
  - TASK-0053.04
documentation:
  - docs/requirements/AGENTS.md
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/terminology/README.md
modified_files:
  - docs/requirements/migration/README.md
  - docs/requirements/migration/backlog-fr-inventory.md
  - >-
    docs/requirements/migration/backlog-fr-inventory-batches/01-layout-and-timed-grid.md
  - >-
    docs/requirements/migration/backlog-fr-inventory-batches/02-event-page-interaction.md
  - >-
    docs/requirements/migration/backlog-fr-inventory-batches/03-dates-only-and-platform.md
  - >-
    docs/requirements/migration/backlog-fr-inventory-batches/04-infrastructure-and-final-ui.md
priority: medium
type: docs
ordinal: 64000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create a non-normative, versioned inventory that turns completed items in `backlog/backlog.md` into reviewable candidates for later migration to atomic functional and quality requirements. The inventory must preserve the original wording and provenance while distinguishing durable expected behavior from implementation details, decisions, investigations, duplicates, and existing requirements. Store it at `docs/requirements/migration/backlog-fr-inventory.md`; add `docs/requirements/migration/README.md` that states the inventory is not normative and that accepted `FR-*` and `QR-*` records remain canonical. Use sequential temporary `CAND-NNN` section identifiers solely for review and traceability; they must never become permanent requirement IDs. Each candidate section must include Source, Candidate behavior, Applicability, Classification, Disposition, and Open Questions when needed. Candidate behavior must be rewritten as observable, independently testable behavior, with role, location, event kind, interaction mode, viewport, state, and exclusions recorded when applicable. Preserve original task wording verbatim in a quoted Source section. Use controlled glossary terminology and links when drafting candidate behavior, but do not invent glossary terms or make uncertain terminology normative. Record overlap with existing FRs/QRs and a confidence classification of confirmed, inferred, or needs product decision. Retain migrated inventory entries with their resulting permanent requirement ID for provenance. This task complements, but does not replace, TASK-0053.04, which defines source-neutral provenance metadata for canonical functional requirements.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 A non-normative inventory exists at docs/requirements/migration/backlog-fr-inventory.md and its README states that canonical requirements remain the accepted FR-* and QR-* records
- [x] #2 The inventory uses one sequential CAND-NNN section per reviewed source item or coherent source-item group and states that CAND identifiers are temporary and never become requirement identifiers
- [x] #3 Every candidate section preserves original backlog wording in a quoted Source section and records a stable source reference
- [x] #4 Every candidate section includes Candidate behavior Applicability Classification and Disposition sections plus Open Questions when scope or intent is unresolved
- [x] #5 Candidate behavior describes one observable independently testable outcome and records applicable actor location event kind mode viewport state and exclusions without generalizing beyond the source evidence
- [x] #6 Every reviewed item is classified as a candidate FR candidate QR existing requirement ADR or decision implementation detail bug or investigation duplicate or refinement or needs product decision
- [x] #7 Each candidate records any matching existing FR or QR and confidence as confirmed inferred or needs product decision
- [x] #8 Migrated or resolved candidates remain in the inventory with their permanent FR or QR identifier or final disposition so provenance remains auditable
- [x] #9 Candidate behavior uses approved glossary terminology and first-use links where terms are established while uncertain terms are recorded for glossary or product review rather than treated as normative
- [x] #10 The inventory format is documented with at least one resolved-existing-requirement example and one unresolved candidate example
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Build a source manifest for every checked top-level source group in the completed MUST - Done, SHOULD - Done, and COULD - Done sections; preserve child bullets with their parent and reserve one sequential CAND-NNN ID per group.
2. Create four durable, non-normative batch inventories under docs/requirements/migration/backlog-fr-inventory-batches/, with fixed source ranges and reserved IDs: 001-080 (267-380), 081-139 (381-481), 140-167 (482-535), and 168-199 (536-585). Each batch uses the canonical candidate schema and quotes source wording verbatim.
3. Add docs/requirements/migration/README.md documenting the authority boundary, temporary identifiers, controlled classifications/confidence values, schema, terminology rule, batch hierarchy, and resolved/unresolved examples.
4. Merge the batch candidate records verbatim and in numeric order into docs/requirements/migration/backlog-fr-inventory.md; retain the batch files as durable review artifacts with backlinks to the consolidated inventory.
5. Validate exact source coverage, sequential IDs, required sections and applicability fields, controlled values, Open Questions conditions, permanent-ID dispositions, source anchors, local links, and glossary anchors. This is documentation-only work, so unit and E2E suites are exempt.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Created the requested non-normative migration README and consolidated inventory. Mechanically concatenated all four durable batches in CAND numeric order, normalizing consolidated source links to ../../../backlog/backlog.md and copied FR/QR links to their consolidated relative locations. Validation confirmed CAND-001 through CAND-199 occur exactly once and in order; all required candidate fields, controlled classification/confidence values, applicable Open Questions, source links, and local Markdown links resolve. Documentation-only change: unit and E2E suites exempt. `git diff --check` reports pre-existing trailing whitespace in docs/requirements/functional-requirements.md; the consolidated raw source quote also retains its batch-provenance trailing whitespace.

Reopened after review confirmed the prior consolidated inventory and migration README retained obsolete Batch 03/04 candidate ranges. The corrected durable batches are authoritative: CAND-140 through CAND-167 cover source lines 482-535 and CAND-168 through CAND-199 cover lines 536-585. Regenerate the consolidated inventory mechanically from the corrected batches, then re-run source-coverage and documentation validation before finalization.

Regenerated the consolidated inventory mechanically from the corrected durable batches. Corrected README Batch 03/04 ranges and replaced the obsolete combined source-group explanation. Exact-source validation found eight Markdown blockquotes that Prettier would reindent or collapse; targeted prettier-ignore comments preserve their quoted backlog wording without trailing whitespace. Final validation passed for 199 ordered batch candidates and 199 ordered consolidated candidates against 199 checked top-level backlog groups, including exact source quotes, schema, applicability fields, controlled values, required questions, source links, and local Markdown links. Scoped Prettier and git diff --check passed. This remains documentation-only work, so unit and E2E suites are exempt.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Regenerated the non-normative completed-backlog candidate inventory from the corrected durable batch artifacts and corrected the migration guide's Batch 03/04 ranges to CAND-140 through CAND-167 and CAND-168 through CAND-199. The consolidated inventory now preserves one candidate for each checked top-level completed-backlog group, with normalized links for its location. Targeted Prettier ignore directives preserve the few raw Source blockquotes whose significant indentation or line structure would otherwise be rewritten. Validation passed: 199 ordered batch candidates and 199 ordered consolidated candidates match 199 exact checked source groups; schema, applicability, controlled classifications/confidence, required Open Questions, source links, and local links pass. Scoped Prettier and git diff --check pass. Unit and E2E suites are exempt because this task changes documentation only.
<!-- SECTION:FINAL_SUMMARY:END -->
