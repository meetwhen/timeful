---
id: TASK-0042
title: Migrate canonical functional requirements
status: Done
assignee:
  - OpenCode
created_date: '2026-08-21 15:03'
updated_date: '2026-08-21 15:21'
labels: []
dependencies: []
references:
  - docs/requirements/README.md
  - docs/requirements/functional-requirements.md
  - backlog/backlog.md
modified_files:
  - docs/requirements/README.md
  - docs/requirements/functional/FR-001.md
  - docs/requirements/functional/FR-002.md
  - docs/requirements/functional/FR-003.md
  - docs/requirements/functional/FR-008.md
  - docs/requirements/functional/FR-011.md
  - docs/requirements/functional/FR-013.md
  - docs/requirements/functional/FR-015.md
  - docs/requirements/functional/FR-017.md
  - docs/requirements/functional/FR-018.md
  - docs/requirements/functional/FR-021.md
  - docs/requirements/functional/FR-023.md
  - docs/requirements/functional/FR-024.md
  - docs/requirements/functional/FR-025.md
  - docs/requirements/functional/FR-026.md
  - docs/requirements/functional/FR-028.md
  - docs/requirements/functional/FR-029.md
  - docs/requirements/functional/FR-030.md
  - docs/requirements/functional/FR-031.md
  - docs/requirements/functional/FR-032.md
  - docs/requirements/functional/FR-033.md
  - docs/requirements/functional/FR-034.md
  - docs/requirements/functional/FR-035.md
  - docs/requirements/functional/FR-036.md
priority: medium
type: docs
ordinal: 30000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move clearly stated functional requirements from the temporary legacy source and selected open backlog items into canonical per-file FR records. Wording must reflect completed backlog behavior, preserve verifiable details, and avoid statements contradicted by completed work.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Canonical FR files use the documented YAML front matter and ID-only filenames
- [x] #2 Every migrated requirement is atomic verifiable and has responsible components
- [x] #3 Requirement wording is reconciled with relevant completed backlog behavior and does not contradict it
- [x] #4 Open but valid requirements are marked proposed and confirmed behavior is marked accepted
- [x] #5 The Functional Requirements index links every canonical FR file
- [x] #6 Temporary migration-source files remain intact
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Reconcile legacy FR candidates and selected open backlog requirements against completed backlog behavior.
2. Create atomic canonical files under docs/requirements/functional with accepted or proposed status and responsible components.
3. Preserve detailed verifiable constraints as acceptance criteria where they jointly define one behavior; omit styling and unsettled authorization decisions.
4. Update docs/requirements/README.md with the linked Functional Requirements index.
5. Verify front matter, ID/filename consistency, index coverage, source preservation, and documentation formatting.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Reconciled legacy wording with completed backlog behavior. Replaced the legacy no-delete rule with the completed delete action; retained grid-state and scheduling details as acceptance criteria; marked unresolved requirements proposed. Validation: `npm exec prettier -- --check ../docs/requirements/README.md ../docs/requirements/functional/*.md` from `frontend/` passed. Unit and E2E suites were not run because this change only adds and formats Markdown documentation.

Canonical record set: `docs/requirements/functional/FR-001.md` through `FR-023.md`. The initial migration created these 23 records from `docs/requirements/functional-requirements.md` and selected entries in `backlog/backlog.md`; the follow-up task `TASK-0042.01` renumbered the initial non-contiguous identifiers into this sequence. The unified index is `docs/requirements/README.md`.

Source coverage from `functional-requirements.md`: migrated legacy entries formerly identified as FR-001, FR-002, FR-003, FR-008, FR-011, FR-013, FR-015, FR-017, FR-018, FR-021, FR-023, FR-024, FR-025, FR-026, and FR-027. The migration split some broader source statements into more atomic canonical requirements, including the saved-slot rule that was part of the former FR-021.

Selected backlog-source coverage: guest response records, default new-event hours, creator-only event metadata edits, multiline descriptions, persistent mobile selected-slot tooltips, scheduling only within active slots, browser-derived event timezones, adjacent-month date-picker days, and email registration status during sign-in. These entries were selected because they expressed durable product behavior rather than implementation work or investigations.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Created 23 canonical functional-requirement files and populated the requirements index. The migration preserves detailed timed-grid, scheduling, timezone, availability, and sign-in behavior while using completed backlog entries to avoid superseded rules. Confirmed behavior is marked accepted; unresolved product or implementation work is marked proposed. Prettier validation passed for all migrated Markdown; unit and E2E suites were not run because no application code changed.
<!-- SECTION:FINAL_SUMMARY:END -->
