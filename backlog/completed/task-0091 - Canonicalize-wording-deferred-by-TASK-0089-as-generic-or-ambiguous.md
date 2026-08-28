---
id: TASK-0091
title: Canonicalize wording deferred by TASK-0089 as generic or ambiguous
status: Done
assignee:
  - opencode
created_date: '2026-08-27 16:02'
updated_date: '2026-08-28 17:37'
labels:
  - documentation
  - terminology
dependencies: []
references:
  - >-
    backlog/tasks/task-0089 -
    Align-functional-requirements-wording-with-the-glossary.md
documentation:
  - docs/terminology/glossary.md
  - docs/terminology/README.md
  - docs/requirements/README.md
  - docs/requirements/AGENTS.md
priority: medium
type: docs
ordinal: 98000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
TASK-0089 (Align functional requirements wording with the glossary) finished with a list of wording it "Deliberately left unchanged as generic or ambiguous wording": FR-030 "Active Timed Slots" (title hybrid), FR-051 verb "Picked" (verb participle), FR-058 title "event domain" (flagged as pending review), FR-093 "active grid cells", FR-094 "disabled timeslot", FR-086 title "specific-times grids" (not in the glossary's rejected alias list), and FR-099 body backticked `Disabled` status phrasing. The implementation notes also deferred the linked [Timed Event] page label pattern.

Follow up by deciding, per item, whether the wording should be canonicalized against existing glossary terms, the glossary should be extended with a new canonical term or rejected alias, or the generic wording should be kept and documented as intentional. Then apply the resulting wording decisions in docs/requirements/functional/ per the rules in docs/terminology/README.md. Read docs/terminology/README.md and docs/requirements/AGENTS.md before changing glossary or requirement records. Keep the change wording-only: do not alter requirement intent, scope, components, status, or IDs.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Every term on TASK-0089's deliberately-unchanged list (FR-030 "Active Timed Slots", FR-051 verb "Picked", FR-058 title "event domain", FR-086 title "specific-times grids", FR-093 "active grid cells", FR-094 "disabled timeslot", FR-099 backticked `Disabled` status phrasing) has an explicit recorded decision in this task: canonicalize to an existing glossary term, introduce or extend a glossary term or rejected alias, or keep as intentionally generic wording.
- [x] #2 Any new canonical terms, inflections, or rejected aliases are added to docs/terminology/glossary.md first, following docs/terminology/README.md canonicalization rules, before FR wording changes.
- [x] #3 FR wording and title changes follow docs/terminology/README.md: first-occurrence glossary links, bold same-scope repeats, possessives inside link labels, one sentence per physical source line, and no change to requirement intent, scope, components, status, or IDs.
- [x] #4 Any changed FR title is mirrored in docs/requirements/README.md and related FR cross-references.
- [x] #5 Grep sweeps show no remaining instances of decided rejected aliases in docs/requirements/functional/, and all glossary anchors in changed files resolve to real glossary headings.
- [x] #6 All changed Markdown files pass npm run format:markdown:check.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: TASK-0092 execution
created: 2026-08-28 17:34
---
TASK-0092 decided all seven deferred items on 2026-08-28 (owner-confirmed decision record in TASK-0092's description) and applied them in docs/requirements/functional/:

- FR-030 "Active Timed Slots": canonicalized to Active Time Slots — title "Persist only Active Time Slots in browser storage"; body stores [Active Slots] and derives the [Enabled Domain].
- FR-051 verb "Picked": canonicalized to Event Picked Dates — title "Keep Active Slots off newly added Event Picked Dates in Custom Timed Domain Mode"; body reworded around Event Picked Dates plus Custom Timed Domain Mode, intent unchanged (new dates contribute no Active Slots).
- FR-058 title "event domain": canonicalized to the Enabled Domain — title "Reject availability outside the Enabled Domain".
- FR-086 title "specific-times grids": canonicalized to Timed Grids — title "Render the Timed Event Occurrence Span on Timed Grids"; body names the Timed Event Scheduling Page and the Timed Grid, with the setting sentence reworded to "without painting availability".
- FR-093 "active grid cells": canonicalized to Active Slots — title "Clear the slot selection when interacting outside Active Slots".
- FR-094 "disabled timeslot": verified against the timed-grid interaction code (taps on non-selectable cells clear the selection; the `Disabled` wording belongs to cells outside the event dates, per SpecificTimesInstructions.vue and ColorLegend.vue) and canonicalized to Disabled Status — title "Clear the selection immediately when tapping a disabled cell on mobile"; body: tapping a [Timed Grid] cell that shows [Disabled Status] clears the selection, [Padding Cells] included.
- FR-099 backticked `Disabled` status phrasing: canonicalized to the Disabled Status controlled term — body now links [Disabled Status].

Also recorded: the deferred "[Timed Event] page" label pattern stays as intentional generic wording for "a page of a Timed Event"; requirements that mean a specific surface name the Event Pages family entry (for example FR-086 uses Timed Event Scheduling Page).

Glossary terms were added first in the TASK-0092 glossary rewrite (Slot, Time Slot, Date Slot, Event Picked Dates, Enabled Domain, Inactive Slots, Padding Cell, Disabled Status, Event Occurrence Span family), then FR wording was applied with first-occurrence links, bold repeats, possessives inside markup, and one sentence per line; titles are mirrored in docs/requirements/README.md; requirement intent, scope, components, status, and IDs are unchanged.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Executed inside TASK-0092 as its decision record required. All seven deferred items were decided (recorded in the task comment) and applied in docs/requirements/functional/: FR-030 → Active Time Slots; FR-051 → newly added Event Picked Dates; FR-058 → Enabled Domain; FR-086 → Timed Grids; FR-093 → Active Slots; FR-094 → disabled cell showing Disabled Status with Padding Cells included (verified against the timed-grid interaction code and the `Disabled` labels in SpecificTimesInstructions.vue and ColorLegend.vue); FR-099 → Disabled Status controlled term. The "[Timed Event] page" label pattern was recorded as intentionally generic wording. New glossary terms landed first in the TASK-0092 glossary rewrite, FR titles are mirrored in docs/requirements/README.md, intent/scope/components/status/IDs are unchanged, grep sweeps for the decided aliases are clean in docs/requirements/functional/, all glossary anchors resolve, and npm run format:markdown:check passes at the repository root.
<!-- SECTION:FINAL_SUMMARY:END -->
