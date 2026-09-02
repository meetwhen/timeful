---
id: TASK-0139
title: Migrate Markdown formatting to oxfmt once it supports custom Prettier plugins
status: To Do
assignee: []
created_date: '2026-09-02 12:39'
labels:
  - tooling
  - formatting
  - upstream-blocked
dependencies: []
references:
  - 'https://oxc.rs/docs/guide/usage/formatter/unsupported-features.html'
  - 'https://github.com/oxc-project/oxc/milestone/19'
priority: low
type: chore
ordinal: 152300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Why

Root Markdown is currently formatted by a standalone Prettier pipeline (`scripts/markdown.mjs`) that loads the custom wrapper plugin `prettier/markdown/sentences-per-line.js` (extends `prettier-plugin-sentences-per-line` with sentence-break insertion at paragraph gaps, table whitespace restoration, and abbreviation suppression). oxfmt handles only JS under `scripts/` and `prettier/` and ignores all Markdown via `ignorePatterns: ["**/*.md"]` in `.oxfmtrc.json`.

oxfmt (checked at 0.66.0) cannot load this plugin: its bundled Prettier only registers hard-coded, sentinel-gated plugins (`_useSveltePlugin`, `_useTailwindPlugin`, `_oxfmtPluginOptionsJson` in oxfmt's `src-js/libs/apis.ts`), and there is no `plugins` config option.

## Unblock condition

Start this task when oxfmt ships one of:

- a config option to load custom Prettier plugins (track the "Prettier plugins" entry on https://oxc.rs/docs/guide/usage/formatter/unsupported-features.html and the oxc formatter milestone), or
- a native sentences-per-line / sentence-aware Markdown feature that is equivalent for our use.

Re-evaluate at that time: if a native built-in covers our semantics, prefer it over plugin loading.

## Outcome

A single oxfmt-based formatting flow that keeps the current sentences-per-line semantics:

- one sentence per physical prose line, with the current abbreviation suppression (upstream ignored words, `knownAbbreviations`, and the `sentencesPerLineAdditionalAbbreviations` option) still respected
- Markdown table cells keep their single-space separation
- no reliance on oxfmt's native Markdown formatter while it lacks sentence-awareness (do not simply drop the `**/*.md` ignore)

## Constraints

- Formatting parity: `format:markdown:check` (or its successor) must pass with no unintended reformatting of existing tracked Markdown; any intentional output differences are reviewed and committed deliberately.
- Documentation with Markdown sentences-per-line rules (root AGENTS.md, docs/, backlog/) must stay stable after migration.
- Do not change sentence semantics as part of this migration; behavior parity is the goal.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Repo Markdown is formatted by oxfmt with sentences-per-line semantics equivalent to the current scripts/markdown.mjs pipeline, including abbreviation suppression and table whitespace behavior.
- [ ] #2 Existing tracked Markdown files are not reformatted as part of the migration (formatting check passes; intentional differences are reviewed and committed separately).
- [ ] #3 The standalone Prettier Markdown driver (scripts/markdown.mjs and prettier/markdown/sentences-per-line.js) is retired or reduced to whatever the new oxfmt flow requires, and prettier-plugin-sentences-per-line is dropped from devDependencies if no longer used.
- [ ] #4 Formatting entry points are updated: package.json scripts (fmt, fmt:check, lint:markdown, format:markdown, format:markdown:check), .oxfmtrc.json ignore patterns, and root AGENTS.md notes about the Prettier sentences-per-line pipeline.
- [ ] #5 Markdown formatting tests covering the sentences-per-line rules (currently under prettier/markdown/) pass in the new setup or are ported to an equivalent suite.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
