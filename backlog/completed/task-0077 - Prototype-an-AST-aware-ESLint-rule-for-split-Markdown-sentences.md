---
id: TASK-0077
title: Prototype an AST-aware ESLint rule for split Markdown sentences
status: Done
assignee:
  - OpenCode
created_date: '2026-08-26 11:20'
updated_date: '2026-08-26 12:19'
labels:
  - documentation
  - tooling
  - markdown
  - eslint
dependencies: []
references:
  - TASK-0075
  - TASK-0076
  - TASK-0076.01
  - eslint.config.ts
  - package.json
  - scripts/markdown.mjs
  - 'https://github.com/eslint/markdown'
  - 'https://github.com/JoshuaKGoldberg/sentences-per-line'
modified_files:
  - eslint.config.ts
  - package.json
  - scripts/markdown.mjs
  - eslint/markdown/no-split-sentence.js
  - eslint/markdown/no-split-sentence.test.mjs
priority: medium
type: enhancement
ordinal: 82000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Create a conservative local ESLint prototype that detects prose sentences continued across physical Markdown source lines and can repair only unambiguous soft-wrap boundaries. The repository currently uses `@eslint/markdown` and `eslint-plugin-sentences-per-line` to prevent multiple sentences on one line, but it does not detect the inverse case. TASK-0076 and TASK-0076.01 found that `textlint-rule-idiomatic-lines` is unsafe because list-item filtering does not cover all list-like input. The prototype must use the existing Markdown AST integration so structural content is not treated as ordinary prose. It is intentionally not a repository-wide enforcement or corpus-normalization change: TASK-0075 retains ownership of reviewing and normalizing the existing Markdown corpus before any future CI rollout.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 A local JavaScript ESLint rule reports a prose sentence that continues across a physical Markdown line boundary
- [x] #2 The rule does not report headings front matter fenced or indented code tables list items nested list items hard line breaks or ambiguous inline-structure boundaries
- [x] #3 The rule provides an autofix only for an unambiguous soft line break in eligible plain paragraph text and the fix preserves Markdown meaning
- [x] #4 Automated fixtures verify diagnostics and fixed output for ordinary prose abbreviations punctuation and the excluded structural cases
- [x] #5 The prototype has a reproducible repository command for its fixture tests
- [x] #6 Existing lint:markdown and Markdown CI enforcement remain unchanged and the task makes no repository-wide Markdown corpus edits
- [x] #7 The local rule and its safe-fix boundary are documented for the future enforcement task after TASK-0075 is complete
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
1. Claim the task and inspect installed tooling (@eslint/markdown 8.0.3, ESLint 10.9.1, flat-config Linter API verified working with language markdown/gfm).
2. Add eslint/markdown/no-split-sentence.js: visit only paragraphs whose direct parent is root; scan top-level text children for soft-wrap newlines; skip hard breaks (trailing backslash or 2+ spaces become break nodes); report boundaries whose previous line lacks confident sentence-terminal punctuation unless a digit period; treat whitelisted abbreviations and initialisms as continuations; autofix joins the boundary with a single space only when both edges match a plain-character allowlist.
3. Add eslint/markdown/no-split-sentence.test.mjs: Vitest suite driving new Linter(configType flat) with an inline array config (files **/*.md, language markdown/gfm, languageOptions.frontmatter yaml, local plugin rule error); fixtures cover ordinary prose diagnostics plus fixed output, abbreviations, terminator punctuation, multi-boundary wraps, unsafe-edge no-fix, and one structural exclusion fixture covering front matter, headings, fenced and indented code, tables, nested/lazy/task-list items, blockquotes, setext traps, hard breaks, and inline code/link/emphasis wrapping.
4. Root package.json: add vitest devDependency aligned with frontend and script test:markdown-rules = vitest run eslint/markdown/.
5. eslint.config.ts: register the local plugin with the rule explicitly off; lint:markdown behavior must stay identical.
6. Add eslint/markdown/README.md documenting purpose, heuristics, safe-fix boundary, and future enablement after TASK-0075; write it one sentence per line so it passes the scoped checks.
7. Validate: npm run test:markdown-rules; npm run lint:markdown unchanged; npm run format:markdown:check; git diff --check; no corpus edits beyond new files; graphify update .; finalize with evidence.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research confirmed @eslint/markdown 8.0.3 + ESLint 10.9.1 flat Linter works with language markdown/gfm; mdast represents soft wraps as newlines inside text node values while hard breaks become break nodes, so boundary scanning is exact.

Eligibility gate is paragraph with direct parent root; this excludes list items even for lazy unindented continuations, which was the textlint failure mode from TASK-0076.01.

Initial test run surfaced three issues, all fixed: the fixture config initially omitted the markdown plugin registration, the fixable tail allowlist needed . for abbreviation continuations, and the setext trap fixture used a non-underline second line so it was genuinely split prose.

Vitest 4 installed at root per user decision; lint:markdown verified unchanged with local/no-split-sentence registered but explicitly off.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Implemented the conservative AST-aware `no-split-sentence` ESLint prototype for Markdown prose sentences that continue across physical lines. The rule visits only paragraphs whose direct mdast parent is the document root, which structurally excludes headings, front matter, fenced and indented code, tables, blockquotes, and all list items including lazy unindented continuations, closing the TASK-0076.01 gap. It scans top-level text children for soft-wrap newlines, skips hard breaks (trailing backslash or two spaces) and inline code/link/emphasis boundaries, reports missing or abbreviation/initialism sentence terminals, treats digit periods as ends, and autofixes only when both edges are plain characters by joining with a single space.

Added `eslint/markdown/no-split-sentence.js`, a Vitest fixture suite in `eslint/markdown/no-split-sentence.test.mjs` driving a flat-config Linter over markdown/gfm with YAML front matter enabled, documentation in `eslint/markdown/README.md`, the root `test:markdown-rules` script, vitest 4.x as a root devDependency, and plugin registration in `eslint.config.ts` with the rule explicitly off.

Validation: `npm run test:markdown-rules` passed 8 of 8 fixtures covering diagnostics, fixed output, abbreviations, terminators, multi-boundary wraps, unsafe-edge no-fix, and the structural exclusion fixture; `npm run lint:markdown` passed unchanged with the rule disabled; `npm run format:markdown:check` and a scoped Prettier check on the new README passed; `git diff --check` clean; no repository-wide Markdown corpus edits.
<!-- SECTION:FINAL_SUMMARY:END -->
