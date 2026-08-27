---
id: TASK-0085
title: >-
  Bold same-block term repeats and convert glossary cross-references to fragment
  links
status: Done
assignee: []
created_date: '2026-08-27 10:54'
updated_date: '2026-08-27 11:02'
labels:
  - terminology
  - documentation
dependencies: []
references:
  - docs/terminology/README.md
  - docs/terminology/glossary.md
modified_files:
  - docs/terminology/README.md
  - docs/terminology/glossary.md
ordinal: 93000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The controlled-terminology home from TASK-0044 defines first-occurrence linking but leaves repeated occurrences unstated, which shows up as inconsistent styling between the glossary (bold terms) and normative documents (links). Establish a complete presentation rule for controlled terms and apply it inside the canonical glossary itself: the first occurrence in a paragraph, list item, or table cell links to the canonical `glossary.md` anchor using the natural inflection of the sentence as the label; further occurrences within the same block render in bold spanning the entire inflected form, keeping a possessive inside the markers. Restrict the change to `docs/terminology/` per the agreed no-sweep decision; an automated checker is postponed and will be specified in a later task. Refines guidance introduced by TASK-0044.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 `docs/terminology/README.md` states one coherent convention: link each controlled term on its first occurrence in every paragraph, list item, and table cell to its `glossary.md` anchor using the natural inflection as the label; style further occurrences in the same block in bold covering the entire inflected form with a possessive inside the markers; never split markup around an inflection.
- [x] #2 `docs/terminology/README.md` preserves exemptions for headings, code blocks or spans, backticked UI labels, link labels, already-linked text, generic lowercase prose, and the dense-table omission; removes the glossary definitional-prose exemption so the glossary follows the same rule; drops the aspirational automated-check sentence pending a later task.
- [x] #3 `docs/terminology/glossary.md` contains no remaining bold-styled term cross-references except permitted same-block repeats and a self-reference naming the defined term; cross-references use fragment links whose targets resolve to existing heading slugs; obvious unstyled Title Case references such as the response-page lists are aligned; generic lowercase prose stays untouched.
- [x] #4 No files besides `docs/terminology/README.md` and `docs/terminology/glossary.md` change; repo-root `npm run lint:markdown`, `npm run test:markdown-rules`, and `npm run format:markdown:check` pass.
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
1. Rewrite the Linking section of `docs/terminology/README.md` into a single convention: first-in-block fragment links whose label may carry the natural inflection while the target stays canonical; same-block repeats bold across the whole inflected form with possessives inside markers; keep all exemptions, drop the definitional-prose exemption, remove the aspirational checker sentence.
2. Convert every bold cross-reference in `docs/terminology/glossary.md` into a same-file fragment link anchored at the canonical heading slug; let plural labels such as `Disabled Padding Cells` point at singular anchors.
3. Within the glossary, treat repeated occurrences inside one definition block as bold (no link), including the Enabled Domain entry naming itself; align obvious unstyled Title Case references such as the response-page lists in Availability Editing/Editor and stray mentions like Platform Identity inside the EVCC entry; leave generic lowercase prose unchanged.
4. Audit with grep for leftover term bolds and validate that every used fragment resolves to an existing heading slug via a throwaway script under /tmp/opencode.
5. Verify no other files changed (`git status --porcelain`) and run repo-root `npm run lint:markdown`, `npm run test:markdown-rules`, `npm run format:markdown:check`.
6. Record evidence per acceptance criterion, write the final summary, set Done.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation notes

- README rewrite: replaced the two-bullet linking guidance with position-based styling: first occurrence links per block, repeats render bold across the full inflected form, possessive stays inside markers; added self-reference allowance for glossary entries naming themselves.
- Glossary conversion performed entry-by-entry from the fresh grep inventory (53 bold-span lines), replacing bolds with fragment links using the heading slug map; plural labels such as [Disabled Padding Cells](#disabled-padding-cell) point at singular anchors.
- Two split-markup possessives found and fixed during conversion: **Dates-Only Event**'s and **Timed Grid**'s became [Dates-Only Event's] / [Timed Grid's] with canonical targets.
- Alignment set beyond existing bolds: EVCC entry (Event Visitor Identity, Platform Identity), Granted EVCC entry (Cross-Device Access Transfer), Availability Response Edit Credential entry (Event Owner Edit Token).
- Deliberately left unstyled: source EVCC, PostgreSQL EVCC authority, a Platform Identity session - shorthand/compound phrases, consistent with no-sweep scope agreed with user.
- Audit script (throwaway under /tmp/opencode pattern via node -e): slug extraction over 91 headings matched all 88 used fragments; leftover-bold grep returns exactly the permitted Enabled Domain self-reference.
- Checks passed at repo root: lint:markdown clean, vitest markdown rules 15/15, format:markdown:check clean. git status --porcelain confirms only docs/terminology/{README,glossary}.md modified plus Backlog task record.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-27 11:02
---
Audit detail for AC #3: the only remaining bold span in glossary.md is the Enabled Domain self-reference on the Together-the-Active-Domain line, which the new README rule explicitly permits. All 88 intra-file fragment links resolve against the 91 heading slugs. Judgment call recorded: `source EVCC`, `PostgreSQL EVCC authority`, and `a Platform Identity session` are shorthand or compound phrases rather than standalone canonical occurrences, so they stay unstyled like the generic lowercase prose; a later task could revisit if the project prefers them linked.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Codified the controlled-term presentation convention and applied it inside the canonical glossary.

Rule (`docs/terminology/README.md`, section renamed to Linking and Styling Terms): link each controlled term's first occurrence per paragraph/list item/table cell to its glossary anchor, using the natural inflection as the label while the target stays canonical (plural labels resolve to singular anchors); style further occurrences within the same block in bold covering the entire inflected form with possessives inside markers such as `**Event Timezone's**`; keep exemptions for headings, code spans, UI labels, link labels, already-linked text, generic lowercase prose, and dense tables; allow an entry to name itself in bold without linking; removed the definitional-prose exemption so the glossary follows its own rule, and dropped the aspirational automated-check sentence pending the postponed checker task.

Glossary (`docs/terminology/glossary.md`): converted all 53 previously bold cross-reference sites into same-file fragment links whose targets all resolve to existing heading slugs, kept the single permitted self-reference (Enabled Domain) bold, folded inflected possessives fully into markup (`[Dates-Only Event's](#dates-only-event)`, `[Timed Grid's](#timed-grid)`), and aligned four previously unstyled Title Case references (Event Visitor Identity and Platform Identity in the EVCC entry, Cross-Device Access Transfer in the Granted EVCC entry, Event Owner Edit Token in the Availability Response Edit Credential entry). Generic lowercase prose and shorthand compounds remain untouched.

Verification: root `npm run lint:markdown`, `npm run test:markdown-rules` (15/15), and `npm run format:markdown:check` pass; `git status --porcelain` shows no changes outside the two documentation files plus Backlog-generated records. Automation of the rule is deliberately postponed for a separately detailed follow-up task.
<!-- SECTION:FINAL_SUMMARY:END -->
