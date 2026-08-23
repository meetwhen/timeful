---
id: TASK-0053
title: Migrate product behavior from frontend ADRs to requirements
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-23 14:35'
updated_date: '2026-08-23 22:00'
labels:
  - architecture
  - requirements
  - frontend
dependencies: []
references:
  - docs/design/adr/003-frontend-freemium-operational-gating.md
  - docs/design/adr/008-frontend-event-ownership-semantics.md
  - docs/design/adr/009-frontend-guest-response-ownership-semantics.md
  - docs/design/README.md
documentation:
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/terminology/README.md
modified_files:
  - docs/design
  - docs/requirements/README.md
  - frontend/AGENTS.md
priority: medium
type: docs
ordinal: 47000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Migrate the retained frontend architecture decisions into the canonical design catalog, eliminate obsolete reference-only records, and make ADR-to-requirement traceability and quality-attribute rationale explicit.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Nine retained frontend ADRs reside in docs/design/adr with contiguous canonical IDs ADR-001 through ADR-009
- [x] #2 The obsolete framework-baseline ADR and superseded timed-event ADR are deleted and frontend/adr is removed
- [x] #3 ADR-005 contains only the civil-date and end-of-day modeling decision while relocated scheduling call-site conventions are documented in frontend guidance
- [x] #4 docs/design/README.md defines the ADR format and index and specifies that ADRs may link to FRs and QRs while requirements remain self-contained and do not cite ADRs as normative dependencies
- [x] #5 Each retained ADR identifies applicable ISO/IEC 25010:2023 quality attributes and its specific QR traceability or explicitly states that no specific QR currently applies
- [x] #6 ADR-008 links QR-005 and QR-006 while ADR-009 links QR-006 only
- [x] #7 Live documentation and frontend guidance use the canonical ADR paths and IDs and Markdown links plus documentation consistency checks pass
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
1. Move the nine retained frontend ADRs to docs/design/adr, delete old ADR-008 and ADR-012, and remove frontend/adr. Renumber the remaining records contiguously: retain 001-006, rename old 009 to 007, old 010 to 008, and old 011 to 009.
2. Narrow ADR-005 to civil-date and end-of-day modeling, relocating working-hours/overlap and rendered-week conventions to frontend/AGENTS.md.
3. Add docs/design/README.md as the sole design-record guide and ADR index. Add the approved directional traceability rule to it and docs/requirements/README.md.
4. Add ISO/IEC 25010:2023 Quality Attributes sections to retained ADRs. Link ADR-008 to QR-005 and QR-006, link ADR-009 to QR-006 only, and explicitly record no specific QR for ADR-001 through ADR-007.
5. Update live paths, titles, and references; do not repair archived Backlog or handoff links. Validate live Markdown links, ADR index coverage, absence of live frontend/adr links, requirement consistency, and whitespace.
6. Apply staged-review corrections: replace the ADR-based requirement-supersession example, distinguish self-contained requirement content from separately recorded provenance, map frontend ADR reading guidance by decision concern, correct ADR-009 QR and quality-attribute traceability, and repair design-guide wording.
7. Create sequential follow-up subtasks: establish canonical user stories, research and record source-neutral provenance for every FR, then add mdsh-generated reverse traceability with frozen validation. Defer the detailed provenance schema and FR-by-FR source research to TASK-0053.03 and TASK-0053.04.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research found frontend re-anchoring in `specificTimesEditDraft.ts`, its unit test, and the Firefox reprojection test. Per the approved documentation-only scope, created dependent subtask TASK-0053.01 to remove that behavior and update regression coverage; no runtime files were changed here.

Validation: `git diff --check` passed. `markdown-link-check` passed for docs/requirements/README.md (95 links), docs/terminology/glossary.md (33 links), all changed ADRs, and all changed FRs. Requirement file count, front-matter ID count, and README index-row count are each 75. Created TASK-0053.02 for 54 unchanged FR files with invalid legacy glossary-link paths.

Review decisions approved: weekly/group enabled-slot semantics are deferred; re-anchoring remains documentation-only follow-up work in TASK-0053.01; TASK-0053 remains In Progress while its child tasks remain open.

Applied approved review fixes: narrowed Enabled Slots terminology to timed specific-date events, assigned derivation to FR-074 and save-time enforcement to FR-015, clarified Availability Response cleanup in FR-057, retained TASK-0053.02, and removed the empty constraints artifact. Validation passed: staged and working-tree `git diff --check`; `markdown-link-check` for the changed requirements, glossary, ADRs, and frontend guidance; staged requirement front matter and index count checks.

Completed the approved ADR catalog migration. Nine frontend ADRs now reside in `docs/design/adr` with contiguous IDs ADR-001 through ADR-009; former ADR-008 and ADR-012 were deleted, and `frontend/adr` was removed. ADR-005 now retains only civil-date and end-of-day modeling; working-hours/overlap and weekly call-site conventions moved to `frontend/AGENTS.md`. Added the repository-wide design guide, directional ADR-to-FR/QR traceability rule, quality-attribute sections, QR-005/QR-006 links for ADR-008/009, and explicit no-QR statements for ADR-001 through ADR-007.

Validation: `git diff --check` passed; changed-document link check passed (116 links). The repository-wide link check retains pre-existing failures: ISO returns HTTP 403 for `docs/requirements/quality/README.md`, and `frontend/backlog/legacy.md` references 30 absent legacy finding files. Archived Backlog and handoff references were intentionally unchanged.

Staged-review fixes applied: requirements no longer use an ADR as a supersession target; design and requirement guides distinguish self-contained requirement content from separately recorded provenance; frontend ADR-reading guidance now identifies ADR-002, ADR-004, ADR-005, and ADR-006 by concern; ADR-009 retains only QR-006 and security authenticity. Created sequential follow-ups TASK-0053.03 (canonical user stories), TASK-0053.04 (FR provenance), and TASK-0053.05 (generated reverse traceability).

Final validation after staged-review fixes: `git diff --cached --check` passed. `npm exec markdown-link-check` passed for the design guide, all nine ADRs, requirements README, and frontend guidance (113 links total). The unrelated unstaged `backlog/backlog.md` change and untracked session handoff were left untouched. TASK-0053 remains In Progress because its dependent subtasks remain open.
<!-- SECTION:NOTES:END -->
