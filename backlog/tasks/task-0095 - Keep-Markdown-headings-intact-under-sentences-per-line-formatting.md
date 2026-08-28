---
id: TASK-0095
title: Keep Markdown headings intact under sentences-per-line formatting
status: To Do
assignee: []
created_date: '2026-08-28 20:58'
labels:
  - tooling
  - markdown
dependencies: []
references:
  - prettier/markdown/sentences-per-line.js
  - prettier/markdown/sentences-per-line.test.mjs
  - TASK-0090
  - TASK-0075.02
priority: medium
type: bug
ordinal: 101000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
`npm run format:markdown` splits headings that contain two sentences across physical source lines (`## Foo. Bar` becomes `## Foo.` followed by `Bar` on the next line). A Markdown heading cannot span physical lines, so on re-parse the second sentence leaves the heading and renders as a separate paragraph, changing rendered semantics. Re-formatting the plain case crashes Prettier's markdown printer with `TypeError: Cannot destructure property 'start' of 'e.position'`. The inline-structure case (`## Identity & Access — **Overview**. Second part here.`) is non-idempotent: the next run drifts the second sentence into a blank-line-separated paragraph.

The breaks originate from two mechanisms: the upstream `prettier-plugin-sentences-per-line` word-level splitter, which inserts breaks inside heading sentence nodes, and the repo wrapper's gap logic added by TASK-0090, whose `gapNodeTypes` includes `heading`. Both must leave headings structurally inert; how the wrapper neutralizes upstream-inserted breaks is an implementation decision.

Constraints: do not regress TASK-0090's paragraph gap splitting or TASK-0075.02's table-row integrity; escaped punctuation such as `&` and `\&` inside headings must survive byte-for-byte; the corpus (including `## Identity & Access` in `docs/terminology/glossary.md`) must remain unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 npm run format:markdown never splits a heading across physical lines; a two-sentence heading round-trips with heading membership and wording preserved, in both the plain case (## Foo. Bar) and the inline-structure case (## Identity & Access — **Overview**. Second part here.)
- [ ] #2 Formatting is idempotent for headings: repeated runs produce no blank-line drift and no printer crash (no TypeError: Cannot destructure property 'start' of 'e.position').
- [ ] #3 TASK-0090 paragraph gap splitting and TASK-0075.02 table-row integrity remain intact per existing fixtures; escaped punctuation such as & and \& inside headings survives byte-for-byte.
- [ ] #4 Regression fixtures for headings are added to prettier/markdown/sentences-per-line.test.mjs and all Markdown checks pass: npm run test:markdown-format, format:markdown:check, lint:markdown, test:markdown-rules, and git diff --check.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
