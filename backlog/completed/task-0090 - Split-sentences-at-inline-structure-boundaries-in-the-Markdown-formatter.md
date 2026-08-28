---
id: TASK-0090
title: Split sentences at inline-structure boundaries in the Markdown formatter
status: Done
assignee:
  - '@OpenCode'
created_date: '2026-08-27 15:16'
updated_date: '2026-08-27 15:50'
labels:
  - tooling
  - markdown
dependencies: []
references:
  - prettier/markdown/sentences-per-line.js
  - prettier/markdown/sentences-per-line.test.mjs
  - scripts/markdown.mjs
  - docs/terminology/glossary.md
  - eslint/markdown/no-split-sentence.js
priority: medium
type: chore
ordinal: 98000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
`npm run format:markdown` cannot repair prose lines where a sentence boundary sits at an inline structure. The joined Timed Domain Mode entry in `docs/terminology/glossary.md` (two sentences on one physical line, the second starting with a `[Ranged Domain Mode](#ranged-domain-mode)` link) is byte-stable under both `npm run format:markdown` and `npm run lint:markdown`, so no tool in the pipeline can restore one-sentence-per-line.

Root cause (verified 2026-08-27):

1. Prettier's built-in Markdown preprocess converts every text node into a `sentence` node whose children are `word`/`whitespace` tokens; `prettier-plugin-sentences-per-line` then inserts `sentenceBreak` nodes only inside those `sentence` nodes. A sentence boundary that falls between paragraph children (text → link → text) is never inside a single `sentence` node, so no break is inserted. Formatting through the JS API pipeline does split plain prose (the CLI shadows plugins and splits nothing); it is only blind at inline-structure gaps.
2. `eslint-plugin-sentences-per-line`'s `one` rule only inspects direct text children of paragraphs, so each fragment around a link holds at most one sentence and it never reports the joined line.
3. `Intl.Segmenter('en', { granularity: 'sentence' })` (verified) segments the real glossary line correctly, does not break at semicolons or colons, and does not break before lowercase continuations. It does break after digit periods and unknown abbreviations before capitals, and Markdown syntax can suppress boundaries, so guards are required.

Desired outcome: the repo-local wrapper `prettier/markdown/sentences-per-line.js` gains gap-aware splitting in its existing `preprocess` hook, using native `Intl.Segmenter` for boundary detection (repo pins Node >= 26.5.0 in engines and CI, so no polyfill dependency is needed), so that `npm run format:markdown` reliably restores one-sentence-per-line for the glossary and the whole corpus while tables stay fully intact.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 npm run format:markdown splits prose onto separate lines when a sentence boundary sits at an inline-structure gap (text ending a sentence followed by a link emphasis or similar, or a structure ending a sentence followed by new prose), altering neither wording nor rendered Markdown meaning
- [x] #2 Markdown table rows remain fully intact: no new splitting inside tables and the existing QR-011-style round-trip coverage still passes
- [x] #3 Formatting is idempotent for the new boundary classes: formatting already-formatted output twice produces no further changes, including boundaries adjacent to links
- [x] #4 Guard rails hold: no splits at digit-period tails, at known or custom abbreviations (including sentencesPerLineAdditionalAbbreviations), at semicolon or colon boundaries, before lowercase continuations, at boundaries that already span a physical line break, or at hard-break siblings; inline code and raw HTML content never triggers a boundary
- [x] #5 Automated coverage in prettier/markdown/sentences-per-line.test.mjs mirrors the new boundary classes and npm run test:markdown-format passes
- [x] #6 The joined Timed Domain Mode entry in docs/terminology/glossary.md is fixed by npm run format:markdown and the full Markdown check set passes: format:markdown:check, lint:markdown, test:markdown-rules, and git diff --check
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
Insert gap splits in prettier/markdown/sentences-per-line.js preprocess, after the base preprocess (which produces sentence nodes) and before restoreTableWhitespace: walk all nodes, skip table subtrees entirely, and for paragraph and heading nodes compute sibling-gap sentence boundaries.

Boundary detection: slice the paragraph raw source from options.originalText using node.position, blank inlineCode and raw HTML leaf ranges with same-length spaces so their content can never originate a boundary, then segment with a lazily created singleton Intl.Segmenter('en', { granularity: 'sentence' }).

Insertion rule per segment boundary (e = end of segment k, s = start of segment k+1): find the last child whose start offset is at or before e; skip if it is the last child, if either sibling is a hard break node, if e lies past the next child's start (boundary is inside a child - upstream already handles those), if the raw slice from e to s contains a newline (boundary already spans lines), or if the segment tail suppresses the boundary (digit-period tail, or tail word ending in a period that doesEndWithIgnoredWord from sentences-per-line accepts, or that word is in the ESLint rule's knownAbbreviations set, or in sentencesPerLineAdditionalAbbreviations).

Apply actions in reverse child order: trim trailing whitespace tokens from the preceding sentence node and leading whitespace tokens from the following sentence node (mirroring upstream's whitespace removal when inserting breaks), then splice a { type: 'sentenceBreak' } node between the siblings. The printer already renders sentenceBreak anywhere in the tree.

Extend prettier/markdown/sentences-per-line.test.mjs with the corpus boundary classes: text ending before a link head (the glossary case), sentence ending inside a link label followed by prose, boundary before emphasis, idempotency round-trips, code-span and HTML inertness, digit and abbreviation guards, custom abbreviation option, semicolon and lowercase non-splits, table integrity, and a hard-break case.

Validate: npm run test:markdown-format, npm run format:markdown (must fix the glossary line), format:markdown:check, lint:markdown, test:markdown-rules, git diff --check, then graphify update .
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation went through four iterations driven by empirical UAX #29 behavior: (1) confirmed Prettier's own markdown preprocess creates sentence nodes and the upstream plugin only splits inside them; (2) space-blanking inline syntax lets UAX #29 absorb the blank into the preceding segment's trailing whitespace, shifting boundaries past child starts; (3) an uppercase-letter placeholder fires boundaries after '. ' but suppresses them when directly adjacent to the ATerm (initialism-style), so heads and tails need different treatment; (4) final design: heads stay raw, inline-code/HTML/image leaves stay raw (the next-segment-start condition makes code-internal boundaries inert), and only inline-structure tail spans (link URLs and closing syntax) are blanked with a non-letter placeholder.

All boundary decisions are position-mask based: a Uint8Array marks blanked source ranges, and nextRealStart/startsWithCapital/proseTailBefore consult the mask, so real X characters in prose never collide with the placeholder. The capital check mirrors upstream's /^[A-Z]/ word-token model, keeping the formatter consistent with the sentences-per-line ESLint rule.

Verification: 15 of 15 formatter fixtures (11 new covering the glossary round trip, link-label ends, emphasis on both sides, multi-boundary paragraphs, table integrity, digit/abbreviation/custom-abbreviation guards, semicolon and lowercase non-splits, code inertness and edges, hard breaks); corpus-wide npm run format:markdown repaired exactly three genuine violations (docs/terminology/glossary.md Timed Domain Mode entry, PLUGIN_API_README.md bold-boundary line, docs/requirements/functional/fr/FR-049.md link-head line) with no other churn; format:markdown:check, lint:markdown, test:markdown-rules (15), git diff --check all pass; graphify update completed.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Extended the repo-local Prettier wrapper `prettier/markdown/sentences-per-line.js` with gap-aware sentence splitting so `npm run format:markdown` now repairs prose lines whose sentence boundary sits at an inline structure — the class that was previously byte-stable under both the formatter and `lint:markdown`.

Mechanism: after the upstream preprocess (which only inserts breaks inside `sentence` nodes), the wrapper now detects sentence boundaries that fall between paragraph/heading children. It segments each paragraph's raw source with a lazily created singleton `Intl.Segmenter('en', { granularity: 'sentence' })` (native; repo pins Node >= 26.5.0, so no polyfill), with inline-structure tail spans (link URLs, closing syntax) blanked via a position mask so URLs can neither suppress nor originate boundaries. Insertions pass guards mirroring the established sentence model: digit-period tails, known abbreviations, `sentencesPerLineAdditionalAbbreviations`, semicolons/colons (Intl does not break), lowercase continuations via an /^[A-Z]/ check on the next real letter, boundaries already spanning a line break, hard-break siblings, and table subtrees (skipped entirely). Boundary whitespace is trimmed to match upstream's whitespace-removal behavior, and all decisions consult the mask so real X letters in prose never collide with the placeholder.

Corpus effect: `npm run format:markdown` repaired exactly three genuine violations — the Timed Domain Mode entry in `docs/terminology/glossary.md`, a bold-boundary line in `PLUGIN_API_README.md`, and a link-head line in `docs/requirements/functional/fr/FR-049.md` — with no other churn.

Verification: `npm run test:markdown-format` 15 of 15 fixtures (11 new: glossary round trip and idempotency, link-head and link-label-end splits, emphasis on both sides, multi-boundary paragraphs, table-row integrity, digit/known-abbreviation/custom-abbreviation guards, semicolon and lowercase non-splits, code-span inertness plus code-edge splitting, hard-break stability); `npm run format:markdown:check`, `npm run lint:markdown`, `npm run test:markdown-rules` (15), and `git diff --check` all pass; `graphify update .` completed.
<!-- SECTION:FINAL_SUMMARY:END -->
