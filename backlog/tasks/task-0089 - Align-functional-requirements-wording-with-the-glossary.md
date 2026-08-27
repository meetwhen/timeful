---
id: TASK-0089
title: Align functional requirements wording with the glossary
status: Done
assignee: []
created_date: '2026-08-27 15:10'
updated_date: '2026-08-27 15:50'
labels:
  - documentation
  - terminology
dependencies: []
documentation:
  - docs/terminology/glossary.md
  - docs/terminology/README.md
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
priority: medium
type: docs
ordinal: 97000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Functional requirement records under docs/requirements/functional/fr/ must use controlled terminology exactly as docs/terminology/glossary.md records it, following the canonicalization and linking rules in docs/terminology/README.md. Audit every FR-*.md file for non-canonical term spellings, missing or incorrect glossary links, incorrect bold styling, and rejected aliases, then fix the wording in place. Keep each sentence on one physical source line. Do not change requirement intent, scope, components, status, or IDs; wording-only alignment. If an FR title changes, update the matching row in docs/requirements/README.md.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Every controlled term used in functional requirement prose matches its glossary.md canonical spelling and capitalization, allowing only inflections that preserve the capitalization of the corresponding glossary words.
- [x] #2 The first occurrence of each controlled term in every functional requirement paragraph, list item, and table cell links to its glossary.md anchor, and further occurrences in the same scope are bold, except in headings, code blocks or spans, backticked UI labels, existing link labels, and text that is already linked.
- [x] #3 Rejected aliases recorded in glossary.md (for example Scheduled Event Time and specific-times page/mode/editing) do not appear in functional requirement prose except when the requirement itself documents the alias as rejected.
- [x] #4 The docs/requirements/README.md functional index stays consistent with any functional requirement titles changed by this task.
- [x] #5 All changed Markdown files pass the project Markdown format check.
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
1. Apply per-file wording fixes from the triaged audit findings, reading each file first to respect link-scope rules (paragraph / list item / table cell). 2. Canonicalize flagged front-matter titles and their H1 headings, then update the matching rows in docs/requirements/README.md. 3. Run npm run format:markdown on changed files. 4. Grep-verify: no rejected aliases, no unstyled canonical-term repeats, no ]'s possessives outside labels. 5. Record final summary, mark acceptance criteria and DoD, set task Done.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Audited all 113 FR files with four parallel sub-agents against docs/terminology/glossary.md and docs/terminology/README.md. Findings triaged.

Apply firm violations: missing first-occurrence state links (FR-004, FR-006, FR-069), missing bold on same-scope repeats (FR-005, FR-018, FR-020, FR-030, FR-049, FR-057, FR-062, FR-065, FR-066, FR-072, FR-073, FR-075, FR-079, FR-081, FR-083, FR-084, FR-085), possessives split outside link labels (FR-013, FR-017, FR-066, FR-073, FR-086, FR-113), lowercase concept references (FR-009 Legend, FR-091 End-of-Day Boundary, FR-092 Legend, FR-100 Dates-Only Grid, FR-105/106/108/110 event-page names, FR-103 Event Settings, FR-076 Availability States, FR-063/FR-079/FR-081/FR-083 missing links), verbatim rejected alias in FR-086 (specific-times mode), FR-090 Dates-Only Grids.

Apply well-founded judgment items: titles FR-001, FR-015, FR-028, FR-062, FR-063, FR-074, FR-097, FR-099, FR-113 canonicalized (index rows updated accordingly); FR-011/FR-028 body availability-editing references; FR-041/FR-045 Unavailable state references after file review.

Leave unchanged as generic or ambiguous: FR-030 title hybrid, FR-051 verb participle, FR-058 title event domain pending review, FR-093 active grid cells, FR-094 disabled timeslot, linked [Timed Event] page label pattern, FR-099 backticked `Disabled` status body phrasing (mirrors glossary's own definition).

Verification evidence: grep sweeps found zero rejected-alias matches and zero `)'s` possessive artifacts in docs/requirements/functional/; an anchor-validation Node script confirmed 0 bad glossary anchors across all 114 functional-directory Markdown files; npm run format:markdown:check exits 0 on three consecutive runs at completion.

Repository note: during execution the working tree had concurrent uncommitted edits outside this task (prettier/markdown/sentences-per-line.js plugin work, docs/terminology/glossary.md, PLUGIN_API_README.md). A concurrent process split FR-049 sentence lines and reformatted the glossary at 18:48 local time; all TASK-0089 edits were verified intact afterward and format:check passes.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Aligned all 113 functional requirement records with the controlled terminology.

Canonicalization and linking fixes in 47 FR files: linked first-occurrence availability states (Available, If needed, Unavailable) in FR-004, FR-006, FR-069; bolded same-scope repeats (FR-005, FR-018, FR-020, FR-030, FR-049, FR-057, FR-062, FR-065, FR-066, FR-072, FR-073, FR-075, FR-079, FR-081, FR-083, FR-084, FR-085); moved possessives inside link labels per the inflection rule (FR-013, FR-017, FR-066, FR-073, FR-086, FR-113); canonicalized lowercase concept references with links (Legend in FR-009/FR-092, End-of-Day Boundary in FR-091, Dates-Only Grid in FR-090/FR-100, Timed/Dates-Only Event Page in FR-105/FR-106/FR-108/FR-110, Event Settings in FR-103, Availability States in FR-076, Enabled Domain in FR-017, Unavailable state in FR-041/FR-045, Picked Dates in FR-058); replaced the rejected alias "specific-times mode" with the Timed Event Scheduling Page in FR-086; converted linked repeats to bold (FR-079); unwrapped sentences edited for terminology onto one physical line (FR-013, FR-073).

Titles canonicalized in FR files and mirrored in the requirements index: FR-001, FR-015, FR-028, FR-062, FR-063, FR-074, FR-091, FR-092, FR-097, FR-099, FR-113; index repeat-styling fixed for FR-005, FR-045, FR-072.

Deliberately left unchanged as generic or ambiguous wording: FR-030 "Active Timed Slots", FR-051 verb "Picked", FR-058 title "event domain", FR-093 "active grid cells", FR-094 "disabled timeslot", FR-086 title "specific-times grids" (not in the glossary's rejected alias list), FR-099 body backticked `Disabled` status phrasing (mirrors the glossary's own definition).

Verification: four parallel audit agents produced the findings; grep sweeps confirm no rejected aliases and no possessives split outside link labels remain in functional/; an anchor-validation script confirms every glossary.md# anchor in docs/requirements resolves to a real glossary heading; npm run format:markdown:check passes; npm run format:markdown was run and reformatted docs/requirements/README.md.
<!-- SECTION:FINAL_SUMMARY:END -->
