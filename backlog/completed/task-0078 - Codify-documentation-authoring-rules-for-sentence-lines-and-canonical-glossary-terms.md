---
id: TASK-0078
title: >-
  Codify documentation authoring rules for sentence lines and canonical glossary
  terms
status: Done
assignee: []
created_date: '2026-08-26 13:50'
updated_date: '2026-08-26 14:02'
labels: []
milestone: Documentation authoring rules
dependencies: []
modified_files:
  - AGENTS.md
  - docs/AGENTS.md
  - docs/terminology/README.md
priority: medium
type: docs
ordinal: 85000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Add written authoring rules so documentation authors follow the tooling's existing constraints and the requirements corpus's existing capitalization practice.

Rule 1: every sentence occupies exactly one physical Markdown line; sentences are never soft-wrapped across lines. Table rows are the exception: a row stays on exactly one physical line regardless of how many sentences its cells contain. This codifies what Prettier and eslint sentences-per-line already enforce at the repository root; the local no-split-sentence detector remains off pending TASK-0075 corpus normalization.

Rule 2: when prose refers to a concept defined in docs/terminology/glossary.md, the term is written in its canonical form exactly as recorded in the glossary entry, word for word including per-word capitalization. Most entries are title case; some are not (for example Show all hours), so the glossary wins over any general capitalization heuristic. Natural inflection (singular, plural, possessive) is allowed without changing word capitalization.

Placement decisions: concise pointer-style section in root AGENTS.md; full rule detail in docs/terminology/README.md next to the existing linking rules. No retroactive normalization of existing docs beyond what TASK-0075 already owns.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Root AGENTS.md contains a Documentation Authoring section that requires one sentence per physical Markdown line, states the table-row exception, and requires controlled terms to be written in their verbatim glossary form
- [x] #2 docs/terminology/README.md documents canonical term use in prose: verbatim glossary spelling and capitalization, permitted natural inflection, and explicitly out-of-scope uses such as generic lowercase prose, code blocks and spans, backticked UI labels, and the glossary's own definitional prose
- [x] #3 Changed Markdown files pass npm run format:markdown:check and the repository Markdown ESLint configuration
- [x] #4 docs/AGENTS.md applies the authoring rules within docs and directs authors to docs/terminology/README.md for controlled-term canonicalization and linking guidance
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
1. Add a concise Documentation Authoring section to root AGENTS.md.
2. Add canonical-form guidance beside the terminology linking rules.
3. Add docs/AGENTS.md as the local entry point that applies the rules under docs and directs authors to the terminology guide.
4. Run the Markdown format check and scoped Markdown ESLint.
5. Record verification evidence and complete the task.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Scope extended at the user's request on 2026-08-26 to add a docs-local authoring entry point.
Documentation-only change; unit and e2e tests are exempt under BACKLOG_WORKFLOW.md.

Verification passed on 2026-08-26:
- npm run format:markdown:check
- npx eslint AGENTS.md docs/AGENTS.md docs/terminology/README.md
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added repository-wide documentation authoring instructions for one sentence per physical Markdown line and single-line table rows. Added docs/AGENTS.md as the local entry point for documentation authors, including a pointer to terminology/README.md for canonicalization and linking guidance. Added glossary canonical-form guidance requiring the exact heading spelling and per-word capitalization, allowing natural inflection and defining exclusions for generic prose, code, UI labels, and glossary definitions. Verified with npm run format:markdown:check and npx eslint AGENTS.md docs/AGENTS.md docs/terminology/README.md.
<!-- SECTION:FINAL_SUMMARY:END -->
