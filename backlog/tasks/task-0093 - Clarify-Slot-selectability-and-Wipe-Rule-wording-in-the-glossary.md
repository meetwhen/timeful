---
id: TASK-0093
title: Clarify Slot selectability and Wipe Rule wording in the glossary
status: Done
assignee:
  - opencode
created_date: '2026-08-28 20:33'
updated_date: '2026-08-28 20:34'
labels:
  - documentation
  - terminology
dependencies: []
references:
  - >-
    backlog/tasks/task-0092 -
    Rewrite-the-glossary-as-DDD-flavored-v2-and-align-requirements-terminology.md
documentation:
  - docs/terminology/README.md
  - docs/terminology/glossary.md
  - docs/requirements/functional/fr/FR-014.md
  - docs/requirements/functional/fr/FR-015.md
  - docs/requirements/functional/fr/FR-057.md
priority: medium
type: docs
ordinal: 99000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Post-completion review of TASK-0092 found three glossary wording tensions between the new definitions and the authoritative requirements. The owner approved the exact rewordings below on 2026-08-28; implement them exactly, do not relitigate.

# Approved rewordings (before → after)
1. Slot entry, first sentence — before: "The selectable unit of an event's Enabled Domain." after: "The unit of an event's Enabled Domain; respondents can select only Active Slots."
   Why: FR-014 (authoritative) and the Inactive Slots entry state Inactive Slots are not respondent-selectable, so "the selectable unit" contradicts the entry's own subsection. The term Slot is unchanged; only the modifier is corrected per docs/terminology/README.md's authority rule.
2. Wipe Rule entry, first sentence — before: "On save, each active Instant outside the Enabled Domain is dropped." after: "On save, each Active Slot outside the Enabled Domain is dropped."
   Why: FR-015 (authoritative, updated by TASK-0092) names the rule over Active Slots; "active Instant" is a v1-model artifact.
3. Wipe Rule entry, third sentence — before: "Changing the Event Timezone does not change Event Picked Dates membership to preserve an otherwise out-of-domain Slot." after: "Changing the Event Timezone rebuilds the Enabled Domain and applies this rule; it does not alter Event Picked Dates to preserve an otherwise out-of-domain Slot."
   Why: the carried-over v1 sentence misparses; FR-057 (authoritative) states the rebuild-and-wipe semantics this rewording mirrors.

# Constraints
- No term renames, no heading changes, no FR or QR body changes, no rejected-alias changes.
- Apply docs/terminology/README.md linking rules in the touched sentences: first-occurrence links for Active Slots, Enabled Domain, Event Timezone, Event Picked Dates, and Slot where they first occur in the entry body; bold for repeats within the same entry.
- One sentence per physical source line.
- The Date Slot entry's "selectable unit" phrasing stays (accurate: all Date Slots are Active Slots).

# Outcome
The glossary's Slot and Wipe Rule entries agree with FR-014/FR-015/FR-057 and parse unambiguously. Documentation-only change.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The Slot entry opens with the owner-approved wording 'The unit of an event's Enabled Domain; respondents can select only Active Slots.' with first-occurrence glossary links, keeping its second sentence and FR-074/FR-058 authoritative context unchanged.
- [x] #2 The Wipe Rule entry opens with 'On save, each Active Slot outside the Enabled Domain is dropped.' with first-occurrence glossary links, keeps the next-day 00:00/00:30 example, and its timezone sentence states that changing the Event Timezone rebuilds the Enabled Domain, applies the rule, and does not alter Event Picked Dates to preserve an otherwise out-of-domain Slot.
- [x] #3 No other glossary entries, requirement bodies, or rejected-alias notes change; every glossary anchor referenced under docs/ still resolves; npm run format:markdown:check passes at the repository root.
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
# Implementation Plan (research completed during the TASK-0092 review session, 2026-08-28)

## Current state (verified)
- docs/terminology/glossary.md Slot entry (lines ~132-137): opens "The selectable unit of an event's [Enabled Domain](#enabled-domain)." then the Timed/Dates-Only sentence, then Authoritative context: FR-074 and FR-058.
- Wipe Rule entry (lines ~247-253): "On save, each active [Instant](#instant) outside the [Enabled Domain](#enabled-domain) is dropped." / next-day 00:00/00:30 example / timezone sentence ending "...membership to preserve an otherwise out-of-domain [Slot](#slot)."
- FR-014/FR-015/FR-057 already carry the authoritative semantics; no FR edits needed.

## Edits (exact)
1. Slot entry sentence 1 → `The unit of an event's [Enabled Domain](#enabled-domain); respondents can select only [Active Slots](#active-slots).`
2. Wipe Rule sentence 1 → `On save, each [Active Slot](#active-slots) outside the [Enabled Domain](#enabled-domain) is dropped.`
3. Wipe Rule sentence 3 → `Changing the [Event Timezone](#event-timezone) rebuilds the **Enabled Domain** and applies this rule; it does not alter [Event Picked Dates](#event-picked-dates) to preserve an otherwise out-of-domain [Slot](#slot).`

## Linking decisions (docs/terminology/README.md)
- Slot entry: Enabled Domain and Active Slots are first occurrences in the touched sentence → link; the following sentence is unchanged.
- Wipe Rule: Active Slot and Enabled Domain first occurrences in sentence 1 → link; Event Timezone, Event Picked Dates, and Slot first occur in sentence 3 → link there; **Enabled Domain** is a repeat within the entry body → bold. The Instant link disappears from the entry (Instant no longer referenced there); the example sentence's lowercase "instants" stays generic prose.

## Verification
- Anchor check script over docs/ (external `terminology/glossary.md#` refs) plus glossary-internal `#` refs and TOC completeness.
- `npm run format:markdown:check` at the repository root.
- `git diff docs/terminology/glossary.md` shows exactly the three sentences changed.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Verification evidence: `git diff docs/terminology/glossary.md` shows exactly the three planned sentences changed (Slot s1, Wipe Rule s1, Wipe Rule s3). Anchor check re-run post-edit: 616 external glossary references under docs/ + PLUGIN_API_README.md resolve, 335 glossary-internal references resolve, 0 broken; TOC covers 86/86 headings. `npm run format:markdown:check` passes at the repository root. The plan's linking decisions were applied as recorded; the Instant link left the Wipe Rule entry with the subject change, and the example sentence's lowercase "instants" remains generic prose.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Applied the three owner-approved glossary rewordings to docs/terminology/glossary.md: the Slot entry now opens "The unit of an event's Enabled Domain; respondents can select only Active Slots." (resolving the contradiction with the not-respondent-selectable Inactive Slots entry and FR-014); the Wipe Rule entry now opens "On save, each Active Slot outside the Enabled Domain is dropped." (aligning the rule's subject with authoritative FR-015 instead of the v1 "active Instant" artifact); and the Wipe Rule timezone sentence now reads "Changing the Event Timezone rebuilds the Enabled Domain and applies this rule; it does not alter Event Picked Dates to preserve an otherwise out-of-domain Slot." (replacing the misparsable v1 carry-over with the FR-057 semantics). No other entries, FRs, QRs, or rejected-alias notes changed; no term renames or heading changes. First-occurrence links applied per docs/terminology/README.md (Active Slots, Enabled Domain, Event Timezone, Event Picked Dates, Slot), bold for the repeated Enabled Domain within the Wipe Rule entry, one sentence per line. Verification: git diff shows exactly the three sentences changed; glossary anchor checks pass (616 external references under docs/ plus PLUGIN_API_README.md and 335 internal references, 0 broken, GitHub slug rules); TOC complete (86/86 headings); npm run format:markdown:check passes at the repository root (which also satisfies the formatting DoD item). Documentation-only change; no runtime code, tests, or payload shapes touched. Changes remain uncommitted; re-staging and commit stay owner-initiated.
<!-- SECTION:FINAL_SUMMARY:END -->
