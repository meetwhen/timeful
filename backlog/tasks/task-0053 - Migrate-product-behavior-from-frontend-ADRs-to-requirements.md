---
id: TASK-0053
title: Migrate product behavior from frontend ADRs to requirements
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-23 14:35'
updated_date: '2026-08-23 18:30'
labels:
  - architecture
  - requirements
  - frontend
dependencies: []
references:
  - frontend/adr/003-frontend-freemium-operational-gating.md
  - frontend/adr/007-frontend-ui-primitive-semantics.md
  - frontend/adr/010-frontend-event-ownership-semantics.md
  - frontend/adr/011-frontend-guest-response-ownership-semantics.md
  - frontend/adr/012-frontend-timed-event-instant-slot-model.md
documentation:
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/terminology/README.md
modified_files:
  - docs/requirements/README.md
  - docs/requirements/functional/fr
  - docs/terminology/glossary.md
  - frontend/adr
  - frontend/AGENTS.md
priority: medium
type: docs
ordinal: 47000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Separate functional requirements from frontend decision records so the ADR set retains durable architectural choices and the requirements catalog owns independently verifiable product behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Functional requirements capture the agreed freemium behaviors, browser-local guest credential persistence, specific-date enabled-domain invariant, and initial ranged-slot generation as atomic verifiable records
- [x] #2 FR-015 and FR-057 define wipe-only specific-date timed-event behavior with stable picked dates and no re-anchor exception
- [x] #3 ADR-003, ADR-010, and ADR-011 retain only architectural decisions and cross-reference applicable functional requirements
- [x] #4 ADR-012 remains as a concise superseded historical record, while ADR-007 is removed and ADR-005/ADR-009 retain only architectural decisions
- [x] #5 Glossary and frontend guidance links resolve to current ADRs or functional requirements, including corrected ADR-001/ADR-002 filenames
- [x] #6 Requirements index lists every newly created functional requirement and Markdown links and index consistency are validated
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Validate Markdown links and requirement index consistency
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add FR-070 through FR-075 for the three disabled-freemium outcomes, browser-local guest credential persistence, the specific-date enabled-domain invariant, and initial ranged-slot generation; add index rows.
2. Broaden FR-015 from custom-domain saves to specific-date timed-event saves, and revise FR-057 to keep picked dates stable and wipe out-of-domain active slots after event-timezone changes without re-anchoring.
3. Narrow ADR-003, ADR-010, and ADR-011 to architectural decisions with links to the applicable FRs; reduce ADR-005 and ADR-009 to decisions and rules while moving procedural detail to frontend/AGENTS.md.
4. Replace ADR-012 with a concise superseded historical record that links to the requirements catalog; delete ADR-007 and remove its references.
5. Repoint glossary authorities from ADR-011/012 to the relevant FRs and update frontend/AGENTS.md with corrected ADR filenames, requirement-based timed-slot authority, retained framework guidance, and relocated conventions.
6. Validate requirement IDs/front matter/index rows, Markdown links and anchors, and git diff whitespace. If research during the docs work shows current code still re-anchors picked dates, create a follow-up implementation task rather than changing runtime behavior in this documentation-only task.

Review fixes: narrow the Enabled Slots glossary definition to timed specific-date events until weekly/group requirements exist; remove FR-074's persistence clause so FR-074 derives the domain and FR-015 enforces save-time wiping; clarify FR-057 to remove Availability Response selections whose corresponding Active Slots were discarded after a timezone change; retain TASK-0053.02 as a tracked child task; remove the accidental empty constraints.md artifact.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research found frontend re-anchoring in `specificTimesEditDraft.ts`, its unit test, and the Firefox reprojection test. Per the approved documentation-only scope, created dependent subtask TASK-0053.01 to remove that behavior and update regression coverage; no runtime files were changed here.

Validation: `git diff --check` passed. `markdown-link-check` passed for docs/requirements/README.md (95 links), docs/terminology/glossary.md (33 links), all changed ADRs, and all changed FRs. Requirement file count, front-matter ID count, and README index-row count are each 75. Created TASK-0053.02 for 54 unchanged FR files with invalid legacy glossary-link paths.

Review decisions approved: weekly/group enabled-slot semantics are deferred; re-anchoring remains documentation-only follow-up work in TASK-0053.01; TASK-0053 remains In Progress while its child tasks remain open.

Applied approved review fixes: narrowed Enabled Slots terminology to timed specific-date events, assigned derivation to FR-074 and save-time enforcement to FR-015, clarified Availability Response cleanup in FR-057, retained TASK-0053.02, and removed the empty constraints artifact. Validation passed: staged and working-tree `git diff --check`; `markdown-link-check` for the changed requirements, glossary, ADRs, and frontend guidance; staged requirement front matter and index count checks.
<!-- SECTION:NOTES:END -->
