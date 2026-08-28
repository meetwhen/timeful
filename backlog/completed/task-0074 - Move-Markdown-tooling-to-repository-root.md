---
id: TASK-0074
title: Move Markdown tooling to repository root
status: Done
assignee:
  - OpenCode
created_date: '2026-08-26 10:06'
updated_date: '2026-08-26 10:24'
labels:
  - documentation
  - tooling
dependencies: []
references:
  - eslint.config.ts
  - frontend/package.json
  - frontend/scripts/markdown.mjs
  - .github/workflows/frontend-ci.yml
  - TASK-0073
priority: medium
type: chore
ordinal: 78000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Make the repository-wide Markdown formatting and sentence-per-line lint tooling independently owned at the repository root rather than by the frontend package. This keeps tooling ownership aligned with the Markdown corpus it checks and allows CI to validate relevant non-frontend changes.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Repository-wide Markdown linting and formatting commands run from a root-owned tooling package
- [x] #2 Markdown tooling no longer depends on frontend package scripts configuration or frontend-local configuration files
- [x] #3 The tracked Markdown scope continues to exclude Backlog-managed and agent-configuration directories
- [x] #4 Continuous integration validates Markdown tooling when its configuration or checked documentation changes
- [x] #5 Frontend source linting remains limited to frontend source and test files
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Align the root Markdown tooling package engine declaration with the repository-wide pinned Node 26.5.0 version. 2. Regenerate the root lockfile and verify the Markdown formatter and linter.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Root tooling now owns the Markdown runner, ESLint and Prettier configuration, dependencies, lockfile, and CI workflow. The root configuration preserves the prior formatter plugin set so formatting remains compatible with the approved corpus. Root Markdown formatting and lint checks pass. Frontend lint, typecheck, build, and unit tests also pass; the build emitted the pre-existing undefined Vite environment and CSS :deep() warnings. `graphify update .` completed with its existing source-parser warnings.

Review follow-up accepted: harden root toolchain compatibility, reduce unused root formatting dependencies, and restore protected provenance formatting before revalidating the Markdown checks.

Follow-up validation: removed the root-only Tailwind Prettier plugin, declared Node >=22.14.0 for the sentence-per-line dependency, regenerated the root lockfile, and confirmed `npm run format:markdown:check`, `npm run lint:markdown`, and both staged and unstaged `git diff --check` pass. `graphify update .` completed with pre-existing parser warnings for JSON and SQL inputs.

Aligned the root toolchain baseline with the repository Node 26.5.0 CI version using `>=26.5.0`; this accepts the local 26.7.0 runtime without EBADENGINE warnings. Re-ran the root lockfile generation, Markdown format check, lint, and diff check successfully.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Moved repository-wide Markdown tooling out of `frontend/` into a root-owned npm package and `scripts/markdown.mjs`. The root package now owns the Markdown ESLint and Prettier dependencies, configuration, lockfile, and commands; frontend linting is again scoped only to frontend source and tests. Added a dedicated Markdown CI workflow that installs the root package and runs format and lint checks for Markdown and tooling changes. The existing tracked-file exclusions for `backlog/**`, `.opencode/**`, and `.agents/**` remain intact.

Validation passed: `npm ci`, `npm run format:markdown:check`, and `npm run lint:markdown` at the repository root; `npm ci`, `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` in `frontend/` (135 files, 932 tests). `graphify update .` completed. The frontend build retains its pre-existing Vite environment and CSS `:deep()` warnings.

Review follow-up: root Markdown tooling now declares Node >=22.14.0 and omits the unused Tailwind Prettier plugin, keeping frontend Tailwind formatting ownership local. Regenerated the root lockfile. Validation passed: `npm run format:markdown:check`, `npm run lint:markdown`, and `git diff --check`.

Aligned the root Node engine baseline to >=26.5.0, matching the repository's Node 26.5.0 CI baseline while permitting compatible newer Node 26 releases. Re-ran root lockfile generation, formatter, linter, and diff validation successfully.
<!-- SECTION:FINAL_SUMMARY:END -->
