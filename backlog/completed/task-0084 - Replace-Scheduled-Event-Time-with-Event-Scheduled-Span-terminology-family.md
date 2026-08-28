---
id: TASK-0084
title: Replace Scheduled Event Time with Event Scheduled Span terminology family
status: Done
assignee:
  - opencode
created_date: '2026-08-27 10:16'
updated_date: '2026-08-27 10:26'
labels:
  - terminology
  - documentation
dependencies: []
references:
  - docs/terminology/glossary.md
  - docs/terminology/README.md
  - docs/requirements/README.md
  - BACKLOG_WORKFLOW.md
modified_files:
  - docs/terminology/glossary.md
  - docs/requirements/README.md
  - docs/requirements/functional/fr/FR-012.md
  - docs/requirements/functional/fr/FR-085.md
  - docs/requirements/functional/fr/FR-086.md
  - docs/requirements/functional/fr/FR-113.md
type: docs
ordinal: 92000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Replace the single controlled term "Scheduled Event Time" in docs/terminology/glossary.md with the three-term taxonomy family approved by the owner:

- Event Scheduled Span (umbrella: optional saved occurrence recorded via an event's scheduling page; saveable, replaceable, clearable; shape kind-specific)
- Timed Event Scheduled Span (time range; contexts FR-012, FR-086 rendering, FR-113 non-empty range)
- Dates-Only Event Scheduled Span (a single date today; context FR-012)

Record "Scheduled Event Time" as a rejected alias on the umbrella entry. Canonical capitalization is Dates-Only (capital O). This is vocabulary-only: it does not enable multi-day dates-only spans and does not rename code identifiers.

Then apply a full corpus pass so no live occurrence references the retired term: glossary internal usages (Timed Event entry line ~27, Dates-Only Event entry line ~60, Dates-Only/Timed Event Scheduling Page entries), FR-012 (title triple: YAML title == H1 == requirements README index cell, body, both acceptance-criteria bullets), FR-085, FR-086 (title triple + body), FR-113, and the requirements README index rows for FR-012 and FR-086.

Out of scope: backlog/**, docs/requirements/migration/** (frozen; stale #scheduled-event-time anchors there accepted), runtime code identifiers, and the pre-existing rejected-alias wording "specific-times" in FR-085/FR-086/readmes (record as follow-up candidate). Documentation-only change: unit/e2e tests exempt per project DoD; required checks are the markdown format gates plus anchor/casing sweeps.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 #1 glossary.md records Event Scheduled Span, Timed Event Scheduled Span, and Dates-Only Event Scheduled Span with concise definitions linking FR-012 (+FR-086/FR-113 where relevant), and names Scheduled Event Time as a rejected alias
- [x] #2 #2 No Scheduled Event Time or generic lowercase scheduled-event-time/date prose remains in docs/terminology/glossary.md or docs/requirements/functional/fr/{FR-012,FR-085,FR-086,FR-113}.md outside frozen zones
- [x] #3 #3 FR-012 keeps YAML title == H1 == README index cell and its acceptance criteria use the kind-specific span terms; FR-086 title triple updates consistently
- [x] #4 #4 requirements README index rows for FR-012 and FR-086 reference the new titles and resolve to real glossary anchors
- [x] #5 #5 Out-of-scope material unchanged: backlog/**, docs/requirements/migration/**, all runtime code, PLUGIN_API_README.md
- [x] #6 #6 Zero scheduled-event-time matches across scoped dirs excluding backtick-code exceptions; every referenced #*-span anchor resolves against actual glossary H3 headings
- [x] #7 #7 npm run format:markdown passes from frontend/
- [x] #8 #8 graphify update . runs after edits
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
Execute in this order:

1. Glossary pass (docs/terminology/glossary.md): replace the Scheduled Event Time entry with three entries (Event Scheduled Span umbrella naming Scheduled Event Time as a rejected alias; Timed Event Scheduled Span; Dates-Only Event Scheduled Span), keeping them adjacent at the old entry's position in Grid Rendering Terms. Also rewrite the generic lowercase usages: Timed Event entry sentence ("may have a scheduled event time range"), Dates-Only Event entry sentence ("scheduled event date"), and the two Scheduling Page entries' set-or-clear sentences to their kind-specific terms.
2. FR-012.md: retitle FM + H1 to "Maintain an optional Event Scheduled Span", swap body link to #event-scheduled-span, and rewrite both acceptance-criteria bullets using Timed Event Scheduled Span / Dates-Only Event Scheduled Span.
3. FR-085.md, FR-113.md: swap occurrences to [Timed Event Scheduled Span](../../../terminology/glossary.md#timed-event-scheduled-span).
4. FR-086.md: same swap in body; retitle FM + H1 so the title triple stays synced.
5. docs/requirements/README.md rows 124 and 198: copy new FM title strings verbatim into index cells and fix anchors.
6. Verify loop until silent: grep zero scheduled-event-time occurrences in scoped dirs; resolve every glossary anchor referenced from changed files against actual H3 headings; one-sentence-per-line preserved; cd frontend && npm run format:markdown; graphify update .
7. Finalize: tick ACs, finalSummary, status Done (leave unarchived).

Constraints: never touch backlog/**, docs/requirements/migration/**, runtime code, PLUGIN_API_README.md. No anchor links inside headings or backticked UI labels. Do not purge pre-existing "specific-times" alias wording (defer).
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Replaced the "Scheduled Event Time" controlled term with the approved three-tier Span family.

Glossary (docs/terminology/glossary.md): new entries Event Scheduled Span (umbrella; names "Scheduled Event Time" as a rejected alias), Timed Event Scheduled Span (time range; non-empty save), and Dates-Only Event Scheduled Span (a single date) adjacent at the old entry's position in Grid Rendering Terms. Four internal glossary usages updated: Timed Event entry ("may have a **Timed Event Scheduled Span**"), Dates-Only Event entry, and both Scheduling Page entries' set-or-clear sentences.

Requirements corpus full pass:
- FR-012 retitled to "Maintain an optional Event Scheduled Span" (FM == H1 variant == index cell); body links #event-scheduled-span; both AC bullets use kind-specific terms.
- FR-085 harmonized beyond original plan scope with owner approval: FM/H1/index row now "Exclude Timed Event Scheduled Span cells from Availability Editing"; body phrases no longer use retired vocabulary.
- FR-086 retitled FM/H1/index to "Render the Timed Event Scheduled Span on Specific-Times Grids"; body occurrences swapped.
- FR-113 two occurrences swapped to Timed Event Scheduled Span.
- docs/requirements/README.md index rows FR-012, FR-085, FR-086 updated with matching anchors.

Verification evidence: case-insensitive sweep leaves only the intentional rejected-alias sentence in the glossary (frozen zones excluded per scope); every referenced #*-span anchor resolves against actual glossary H3 headings; prettier format pass realigned README.md; test:markdown-format 4/4 passed, test:markdown-rules 15/15 passed; graphify update . rebuilt (4633 nodes).

Deferred follow-up candidate: purging pre-existing rejected-alias wording "specific-times page/mode/editing" from FR-085/FR-086/index prose. Multi-day Dates-Only spans remain un-normative until a future requirement changes behavior.
<!-- SECTION:FINAL_SUMMARY:END -->
