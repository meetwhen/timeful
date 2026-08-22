---
id: TASK-0044
title: Centralize controlled terminology documentation
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-22 06:08'
updated_date: '2026-08-22 06:16'
labels:
  - terminology
  - documentation
dependencies: []
references:
  - docs/requirements/README.md
  - frontend/AGENTS.md
  - frontend/glossary.md
modified_files:
  - docs/terminology/README.md
  - docs/terminology/glossary.md
  - docs/requirements/README.md
  - frontend/AGENTS.md
  - frontend/glossary.md
priority: medium
type: docs
ordinal: 36000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create a documentation-wide terminology home that defines controlled terms concisely, identifies their authoritative context, and establishes lightweight first-occurrence linking guidance. Migrate the existing frontend timed-event glossary without losing its source links, and point requirements and frontend guidance to the canonical documentation.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 A docs/terminology guide defines glossary authority and first-occurrence linking conventions
- [x] #2 The existing frontend glossary content is available from the canonical documentation location with authoritative-context links preserved
- [x] #3 Requirements and frontend authoring guidance point to the canonical terminology documentation
- [x] #4 The migration preserves existing unrelated worktree changes
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add docs/terminology/README.md as the documentation-wide authoring policy: concise glossary definitions, authoritative-context sources, and first-controlled-term linking per paragraph/list item/table cell.
2. Move the timed-event glossary content to docs/terminology/glossary.md, preserving definitions and repairing relative links to ADR-012, requirements, and frontend source files.
3. Remove the legacy frontend glossary and update frontend agent guidance plus requirements authoring guidance to the canonical terminology documentation.
4. Validate Markdown links and review the final diff without touching unrelated worktree changes.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Created `docs/terminology/README.md` with glossary authority and first-occurrence linking policy. Moved the timed-event glossary to `docs/terminology/glossary.md`, preserving authoritative-context references and correcting requirement links to canonical per-file records. Updated requirements and frontend guidance; removed the old frontend glossary.

Verification passed: `git diff --check`; local Markdown target and anchor validation; `npm run lint`; `npm run typecheck`; `npm run build`; and `npm run test:unit` (135 files, 934 tests). Firefox E2E did not pass: the full run had two browser-context timeout failures. Retrying the two failures at 60 seconds passed the create-flow case, but `timed-event-specific-times-edit-firefox.spec.ts` still timed out after the browser context closed. This is unrelated to the documentation-only changes, so it was not addressed in this task.
<!-- SECTION:NOTES:END -->
