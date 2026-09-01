---
id: TASK-0125
title: Specify event-editor past-date selection rules (FR-118)
status: Done
assignee: []
created_date: '2026-09-01 10:33'
updated_date: '2026-09-01 10:33'
labels:
  - requirements
  - date-picker
dependencies: []
modified_files:
  - docs/requirements/functional/fr/FR-118.md
  - docs/requirements/README.md
type: docs
ordinal: 133300
---

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 docs/requirements/functional/fr/FR-118.md exists with README-conformant front matter (type: functional, components: [frontend], status: accepted) and one sentence per physical line
- [x] #2 The FR-118 index row is added to docs/requirements/README.md with a stable relative link
- [x] #3 Both changed Markdown files pass npm run format:markdown:check
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
- [x] #5 No runtime code, tests, build, or deployment configuration is modified (documentation-only change)
<!-- DOD:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Authored FR-118 (docs/requirements/functional/fr/FR-118.md) specifying event-editor date-picker past-date rules: no selection before the current date in the Event Timezone (create mode uses the form's selected timezone, which becomes the Event Timezone), existing Event Picked Dates before the current date remain selected while editing, no additional past dates selectable while editing, and the earliest selectable date re-derives when the selected Event Timezone changes. Added the FR-118 row to docs/requirements/README.md. npm run format:markdown:check passes. Decisions were confirmed with the product owner: keep current create-mode restriction, keep existing past dates on edit (non-destructive, matches CAND-032), disallow new past picks in edit mode, frontend-only enforcement, no ADR (CAND-032 retains the historical scope note).
<!-- SECTION:FINAL_SUMMARY:END -->
