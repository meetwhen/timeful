---
id: TASK-0096
title: Enforce local/no-split-sentence for Markdown prose
status: To Do
assignee: []
created_date: '2026-08-28 20:58'
labels:
  - tooling
  - markdown
dependencies: []
references:
  - eslint.config.ts
  - eslint/markdown/README.md
  - eslint/markdown/no-split-sentence.js
  - TASK-0075
  - TASK-0077
priority: medium
type: chore
ordinal: 102000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
`eslint/markdown/README.md` stages enforcement of the local `no-split-sentence` rule behind corpus normalization (TASK-0075), which is Done, and the required `frontmatter` language option is already paired in `eslint.config.ts`. The rule is currently registered as `off` in `eslint.config.ts`, so split-sentence violations at inline-structure boundaries are only repaired by the formatter and never reported by lint.

Enable the rule for Markdown prose and confirm the corpus is clean. If genuine violations surface, fix them in the same change or stop and propose a follow-up task before changing scope. Update the rule's README so it no longer describes enablement as pending.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 eslint.config.ts sets local/no-split-sentence to error for Markdown files.
- [ ] #2 npm run lint:markdown passes across the whole corpus with the rule enabled.
- [ ] #3 eslint/markdown/README.md no longer describes the rule as off or enablement as pending.
- [ ] #4 npm run test:markdown-rules fixture suite still passes.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
