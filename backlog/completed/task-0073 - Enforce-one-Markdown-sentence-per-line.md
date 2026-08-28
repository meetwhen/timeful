---
id: TASK-0073
title: Enforce one Markdown sentence per line
status: Done
assignee:
  - OpenCode
created_date: '2026-08-26 09:28'
updated_date: '2026-08-26 09:58'
labels:
  - documentation
  - tooling
dependencies: []
references:
  - frontend/package.json
  - frontend/eslint.config.ts
  - frontend/.prettierrc
  - eslint.config.ts
  - frontend/scripts/markdown.mjs
  - frontend/eslint/markdown.ts
priority: medium
type: chore
ordinal: 77000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Improve reviewability of repository documentation by enforcing and automatically formatting one sentence per source line for all Markdown outside Backlog-managed and agent-configuration directories.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Markdown outside backlog .opencode and .agents is checked for no more than one sentence per line
- [x] #2 A repository command automatically formats Markdown in the approved scope
- [x] #3 Repository linting includes the Markdown sentence-per-line check
- [x] #4 The initial approved Markdown corpus passes the new formatting and lint checks
- [x] #5 Existing frontend lint behavior remains intact
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add the required Markdown parser and sentence-per-line ESLint and Prettier plugins to the existing frontend tooling package.
2. Configure ESLint to lint Markdown and extend package scripts with lint and formatting commands covering all repository Markdown except backlog .opencode and .agents.
3. Run the formatter across the approved corpus then verify the dedicated Markdown checks and existing frontend lint behavior.
4. Record verification evidence and finalize the task.

5. Keep Markdown rules solely in the root configuration used by the scoped runner; do not register Markdown files in the frontend lint configuration, which would cause the existing frontend ESLint command to apply Vue rules to Markdown.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Validation on 2026-08-26: `npm run format:markdown`, `npm run format:markdown:check`, and `npm run lint:markdown` passed for the tracked 346-file scope. `npm run lint` initially exposed Markdown being included in frontend ESLint; removed the duplicate Markdown block from `frontend/eslint.config.ts`, retaining the root-only scoped configuration. Re-ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit`: all passed. Unit suite: 135 files and 932 tests. Build emitted existing Vite environment and CSS `:deep()` warnings, but completed successfully. `graphify update .` also completed.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Implemented repository-wide Markdown sentence-per-line tooling for tracked Markdown outside `backlog/**`, `.opencode/**`, and `.agents/**`. Added a Git-scoped runner, root ESLint Markdown configuration, Prettier sentence-per-line support, and package commands integrated into frontend linting; formatted the approved Markdown corpus. Kept Markdown lint isolated from the existing frontend configuration after validation found that the duplicate configuration caused Vue rules to run on Markdown files.

Validation passed: `npm run format:markdown`, `npm run format:markdown:check`, `npm run lint:markdown`, `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` (135 files, 932 tests). Build completed with pre-existing environment and CSS `:deep()` warnings. Graph updated with `graphify update .`.
<!-- SECTION:FINAL_SUMMARY:END -->
