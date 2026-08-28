---
id: TASK-0076
title: Evaluate tooling for exact one-sentence Markdown lines
status: Done
assignee:
  - '@OpenCode'
created_date: '2026-08-26 10:59'
updated_date: '2026-08-26 11:02'
labels:
  - documentation
  - tooling
  - markdown
dependencies: []
references:
  - 'https://sembr.org/'
  - 'https://github.com/textlint/textlint'
  - 'https://github.com/textlint-ja/textlint-rule-idiomatic-lines'
  - TASK-0073
  - TASK-0075
priority: medium
type: spike
ordinal: 80000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Determine whether the repository can reliably detect and automatically repair non-empty Markdown prose lines that do not contain exactly one sentence, complementing the existing maximum-one-sentence-per-line tooling. Evaluate Sembr, textlint and relevant textlint rules, and the current ESLint and Prettier setup; produce a recommendation and, if low-risk, a minimally scoped experimental setup that does not replace the established formatter until its behavior is validated.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The current Markdown formatter and lint guard limitations for sentences split across lines are documented
- [x] #2 Sembr textlint and applicable textlint rules are evaluated against the repository's Markdown scope
- [x] #3 A recommendation identifies the appropriate enforcement and automatic-repair boundary with documented tradeoffs
- [x] #4 Any adopted experimental configuration is isolated from the established Markdown checks and has a reproducible validation command
- [x] #5 The relationship to TASK-0075 is documented so existing normalization work is not duplicated
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
1. Inspect the root Markdown runner, ESLint, Prettier, CI scope, and TASK-0075 to distinguish current one-way enforcement from the desired exact-line invariant.
2. Evaluate SemBr, the SemBr agent skill, textlint, and textlint-rule-idiomatic-lines from their published documentation and metadata.
3. Run an isolated fixture and a read-only full-corpus textlint trial with pinned textlint 15.8.0 and textlint-rule-idiomatic-lines 1.2.0; verify diagnostics and fixer output before considering integration.
4. Retain the existing ESLint and Prettier commands. Recommend a SemBr agent skill for reviewed reformatting and a future custom Markdown-AST formatter or rule only if exact CI enforcement remains required; do not duplicate TASK-0075's corpus cleanup.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research found that the current root runner passes tracked Markdown outside backlog/**, .opencode/**, and .agents/** to Prettier and ESLint. Prettier-plugin-sentences-per-line and eslint-plugin-sentences-per-line/one enforce or repair multiple sentences sharing a line, but do not detect a sentence continued on the next physical line.

SemBr is a writing convention rather than a linter. Its published skill reflows prose semantically, permits optional clause breaks, and therefore does not implement the repository's stricter exactly-one-sentence-per-non-empty-prose-line invariant. It is appropriate as an opt-in agent capability for reviewed rewrites, not CI enforcement.

The supplied idiomatic-lines repository URL is stale. The maintained npm package is textlint-rule-idiomatic-lines 1.2.0 from rahulsom/textlint-idiomatic-lines, with textlint >=12.2.0 peer dependency. textlint 15.8.0 supports Markdown and fixable rules on the repository Node 26.5.0 runtime.

Isolated fixture result: the rule correctly reported and dry-run fixed one split prose sentence and two sentences on one line. It also reported a wrapped list item despite its README saying list items are excluded; its --fix dry run joined that item with three spaces. This is unsafe for automatic repository formatting.

Read-only full-corpus trial with textlint 15.8.0 and idiomatic-lines 1.2.0 reported 637 diagnostics: 619 sentence-spanning-line reports and 18 multiple-sentence-line reports. The corpus result overlaps TASK-0075's intentional and unintentional wrapping inventory, so this task made no corpus edits and does not take ownership of normalization.

Validation: npm run lint:markdown passed; npm run format:markdown:check passed. Isolated textlint fixture reported the expected three fixable violations; the dry-run diff confirmed the list-item false positive and unsafe fixer behavior. No repository dependencies, configuration, or CI commands were changed.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Evaluated SemBr, textlint, and textlint-rule-idiomatic-lines without changing the established Markdown toolchain. The existing Prettier and ESLint setup remains the correct guard for no more than one sentence on a line, but cannot detect split sentences.

SemBr is recommended only as an opt-in, human-reviewed agent skill because its specification intentionally permits clause-level breaks. The candidate textlint rule detects both directions and fixes ordinary prose in an isolated fixture, but its 1.2.0 implementation also flags wrapped list items and fixes them by inserting unsafe spacing. A full read-only repository trial produced 637 diagnostics, including 619 split-line reports, so it is not safe to adopt in CI or as a formatter.

TASK-0075 retains ownership of reviewing and normalizing the existing corpus. If exact CI enforcement is still required after that cleanup, create a follow-up to implement and test a custom Markdown-AST-aware checker/fixer that explicitly excludes lists, tables, code, front matter, and other structural content.

Validation passed: npm run lint:markdown; npm run format:markdown:check; isolated textlint fixture and --fix --dry-run.
<!-- SECTION:FINAL_SUMMARY:END -->
