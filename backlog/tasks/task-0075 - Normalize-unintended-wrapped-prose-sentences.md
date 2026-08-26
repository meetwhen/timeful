---
id: TASK-0075
title: Normalize unintended wrapped prose sentences
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-26 10:32'
updated_date: '2026-08-26 13:30'
labels: []
dependencies: []
references:
  - fe298f8ef5602d9abb15e127b0a6b4585b490762
  - TASK-0076
  - TASK-0077
  - eslint/markdown/README.md
  - eslint/markdown/no-split-sentence.js
priority: medium
type: chore
ordinal: 79000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Normalize prose sentences that continue across physical Markdown source lines across the repository's tracked Markdown corpus (outside backlog/**, .opencode/**, and .agents/**) so documentation is consistently readable and reviewable and the AST-aware no-split-sentence prototype delivered by TASK-0077 can be enabled in a future enforcement task. The wrapping predated formatting commit fe298f8ef5602d9abb15e127b0a6b4585b490762 which retained it; TASK-0076 ruled external sentence linters unsafe and confirmed this task owns corpus normalization before any CI enforcement. Scope expanded from the fe298f8 file set to the full corpus with user approval on 2026-08-26 because 82 files holding 349 reports were never touched by that commit.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 A report-only scan with the AST-aware local no-split-sentence prototype inventories every prose sentence continued across physical lines in tracked Markdown outside backlog/**, .opencode/**, and .agents/**
- [ ] #2 Every unintentionally wrapped sentence found by the scan is rewritten as exactly one sentence per physical line without altering wording
- [x] #3 Intentional structural line breaks such as headings front matter lists tables code blocks and blockquotes remain unchanged
- [ ] #4 A repeat of the report-only scan reports zero findings on the corpus
- [x] #5 Changed Markdown files pass npm run format:markdown:check npm run lint:markdown and git diff --check
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Build the candidate inventory with a report-only run of the local AST-aware rule over tracked Markdown outside backlog/**, .opencode/**, and .agents/**: npx eslint --config eslint.config.ts --rule '{"local/no-split-sentence": "error"}' <files>. No permanent config changes; the rule stays off in eslint.config.ts.
2. Apply the rule's safe autofix scoped to reported files (--fix): it joins only boundaries whose both edges are plain characters, preserving Markdown semantics. Commit nothing yet; review the resulting diff.
3. Manually review every remaining report in document context and join confirmed sentence continuations into one physical line without altering wording; leave genuinely intentional breaks (verse, labels, deliberate fragments) untouched and record any such exceptions in implementation notes.
4. Re-run the report-only scan expecting zero findings; then validate npm run test:markdown-rules (stays green), npm run format:markdown:check, npm run lint:markdown, and git diff --check. Unit and e2e suites are not required because the change is documentation-only.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Scope decision (user approved): expanded from commit-scoped (95 files / ~699 reports intersecting fe298f8) to the full tracked Markdown corpus after TASK-0076/TASK-0077 established this task owns corpus normalization before enforcement. Report-only scan totals: 177 files / 1048 reports.

Research verified the TASK-0077 prototype can drive this work without config changes: node_modules/.bin/eslint --config eslint.config.ts --rule '{"local/no-split-sentence": "error"}' works with flat config, reports boundaries, and marks most fixable. The rule structurally excludes headings, front matter, fenced/indented code, tables, blockquotes, HTML blocks, hard breaks, and list items including lazy continuations.

Normalization executed in three passes. Pass 1 (original rule) exposed that eslint.config.ts lacked languageOptions.frontmatter, so the scanner misread YAML front matter as prose and the autofix flattened metadata lines in several ADRs; the corpus was restored via git restore and languageOptions.frontmatter 'yaml' was added to the shared config (lint:markdown verified unchanged). Clean baseline rescan: 595 reports across the corpus.

Pass 2 (safe autofix) joined 502 of 595 boundaries across 177 files; 93 remained, all adjacent to or inside inline links/code/emphasis, which the TASK-0077 prototype deliberately skipped.

Subtask TASK-0075.01 reworked the rule to raw-source line scanning; its autofix then resolved 92 of those 93 automatically. Corpus result: exactly one remaining report, the intentionally unfixed pipe fragment in docs/requirements/quality/qr/QR-014.md.

QR-014 disposition per user decision ('ignore qr-014 for now'): its structural repair was attempted (restoring Response and Response measure rows whose leading pipes were lost to earlier formatting) but reverted to HEAD because prettier-plugin-sentences-per-line reformats multi-sentence table cells by splitting physical lines, which destroys valid rows and fails npm run format:markdown:check. The file also carries a pre-existing flattened front matter line. User floated writing an additional plugin/rule for sentences-per-line handling of tables; candidate follow-up task pending user confirmation.

Validation after final state: repeat report-only scan shows 1 finding (QR-014 guard) corpus-wide; npm run test:markdown-rules 15/15; npm run lint:markdown passed; npm run format:markdown:check passed; git diff --check clean; graphify updated.

Staging per user instruction: staged only implementation files (eslint.config.ts, eslint/markdown/no-split-sentence.js, eslint/markdown/no-split-sentence.test.mjs, eslint/markdown/README.md); the 95-file corpus normalization (+419/-1092) remains unstaged for separate review.

Progress commit: the 95-file corpus normalization (+491/-1092 including wording-identical joins across all documented boundary classes) is committed together with this task record. AC #2 and #4 stay open solely for the QR-014 pipe fragment, which awaits the table-aware sentences-per-line follow-up floated by the user; every other scanned sentence now sits on one physical line.
<!-- SECTION:NOTES:END -->
