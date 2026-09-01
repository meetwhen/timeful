---
id: TASK-0114
title: Pin GitHub Actions to full-length SHAs and enable Dependabot updates
status: Done
assignee: []
created_date: '2026-08-30 13:28'
updated_date: '2026-08-30 13:31'
labels:
  - ci
  - security
dependencies: []
priority: medium
type: chore
ordinal: 120000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
All GitHub Actions workflows reference actions by mutable version tag (actions/checkout@v4, actions/setup-node@v4), which does not protect against tag-move supply-chain attacks. GitHub's security guidance is to pin third-party action references to full-length commit SHAs. Dependabot's github-actions ecosystem supports SHA-pinned refs and updates the same-line version comment on bump, so pinning and automated updates can coexist.

Outcome: every GitHub-hosted action reference in .github/workflows/*.yml is pinned to a verified full-length commit SHA with a same-line version comment, and a Dependabot config keeps both actions and npm dependencies updated via weekly PRs. markdown-ci and backend-ci also gain concurrency groups for consistency with frontend-ci, and checkout steps drop persisted git credentials since no step needs them.

Constraints: SHAs must be resolved from the official actions repositories (never guessed); if pinning to the latest stable release of an action breaks CI, fall back to the latest release of the major currently in use. Do not change the pinned Node.js version (26.5.0) or any job semantics beyond the agreed scope.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Every uses: reference to a GitHub-hosted action in .github/workflows/*.yml is pinned to a full-length commit SHA with a same-line version comment, and each SHA is verified against the official action repository rather than a fork
- [x] #2 .github/dependabot.yml enables weekly version updates for the github-actions ecosystem (directory /) plus npm ecosystems for the root and frontend directories
- [x] #3 markdown-ci and backend-ci workflows have concurrency groups with cancel-in-progress: true, consistent with the existing frontend-ci pattern
- [x] #4 actions/checkout steps in all workflows set persist-credentials: false
- [x] #5 All workflow files validate locally with actionlint or an equivalent YAML validator, and markdown-ci is verified against its own trigger paths
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
1. Resolve full-length SHAs for the latest stable releases of actions/checkout and actions/setup-node from the official repositories with git ls-remote; cross-check the tag-to-SHA mapping if any ambiguity exists.
2. Pin all uses: refs in markdown-ci.yml, frontend-ci.yml, and backend-ci.yml to those SHAs with same-line `# vX.Y.Z` comments; add persist-credentials: false to every checkout step.
3. Add concurrency groups (cancel-in-progress: true) to markdown-ci.yml and backend-ci.yml matching the existing frontend-ci.yml pattern.
4. Create .github/dependabot.yml with weekly github-actions updates (directory /) plus npm ecosystems for root and frontend.
5. Validate all workflow files with actionlint (docker image) and report results.
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Pinned every GitHub-hosted action reference across the three workflows to full-length commit SHAs resolved via git ls-remote from the official repositories (not forks), each with a same-line version comment so Dependabot can keep them updated: actions/checkout@3d3c42e5aac5ba805825da76410c181273ba90b1 # v7.0.1 and actions/setup-node@820762786026740c76f36085b0efc47a31fe5020 # v7.0.0 (majors v4 -> v7). Added persist-credentials: false to all checkout steps. Added concurrency groups (markdown-ci-*/backend-ci-* with cancel-in-progress: true) to markdown-ci.yml and backend-ci.yml, matching the existing frontend-ci pattern. Created .github/dependabot.yml with weekly version updates for github-actions (directory /) plus npm ecosystems for the root and frontend directories.

Verification: actionlint (rhysd/actionlint docker image) passed on all three workflow files with exit 0; .github/dependabot.yml parsed as valid YAML. The live markdown-ci run could not be executed locally because it requires GitHub-hosted runners; the workflow lists its own path in trigger paths, so the run is verified automatically on push. No unit or e2e tests were required since only CI configuration changed, and no Markdown files were modified.
<!-- SECTION:FINAL_SUMMARY:END -->
