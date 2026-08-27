---
id: TASK-0088
title: Canonicalize generic term references inside glossary definitions
status: Done
assignee:
  - OpenCode
created_date: '2026-08-27 13:40'
updated_date: '2026-08-27 14:36'
labels:
  - terminology
  - documentation
dependencies: []
documentation:
  - docs/terminology/README.md
modified_files:
  - docs/terminology/glossary.md
priority: medium
type: docs
ordinal: 96000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Finish the terminology cleanup in docs/terminology/glossary.md so that the glossary itself follows the canonicalization rules stated in docs/terminology/README.md (link first occurrence per paragraph/list item/table cell, bold repeated mentions covering the whole inflection, bold-no-link self-references).

Current state: many definitions refer to other glossary terms through unstyled lowercase generic phrasing instead of the canonical form, for example "a timed domain mode" inside Timed Event, "custom domain editing" and "ranged domain mode" inside Custom Domain Mode, "picked dates"/"event timezone" inside Enabled Slots, and "MongoDB availability response edit credential" inside Protected Response. Several compound informal phrases also act as shadow aliases: availability domain, picked-date domain, active-slot domain, enabled-slot domain, enabled full-day domain, schedule-overlap view, date selection.

Decisions already made by the user:
- Do not promote any phrase to a new canonical term in this task; add no new headings.
- Map the informal "availability domain" phrase onto [Enabled Domain] where meaning-preserving.
- FR-086's use of the rejected alias "specific-times mode/grids" and CAND-203's broken glossary.md#scheduled-event-time anchor are out of scope; only note them, do not fix them here.

Constraints:
- Only docs/terminology/glossary.md content changes.
- Keep every sentence on one physical source line.
- Rejected-alias sentences and all Authoritative context links stay unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Every cross-reference inside docs/terminology/glossary.md definitions names its concept using the canonical spelling from the corresponding glossary heading
- [x] #2 Cross-references are styled exactly per docs/terminology/README.md: first occurrence per paragraph, list item, or table cell links to the glossary anchor with an inflection-matching label; further occurrences in the same unit are bold covering the entire inflected form; an entry naming its own term is bold without linking to itself
- [x] #3 The informal phrasings availability domain, picked-date domain, active-slot domain, enabled-slot domain, enabled full-day domain, schedule-overlap view, and date selection no longer appear as standalone concepts in definitions; equivalent canonical-term prose replaces them where meaning-preserving
- [x] #4 No new controlled-term headings are added and existing authoritative-context links, anchors, and rejected-alias sentences (specific-times page/mode/editing and Scheduled Event Time) remain unchanged
- [x] #5 From the repository root, npm run lint:markdown and npm run format:markdown:check pass
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
Research summary: TASK-0085 established that a definition block (all consecutive sentences up to the Authoritative context line) is one styling unit; first occurrence of each controlled term links with a natural-inflection label, later occurrences bold covering the full inflection with possessives inside markers, and self-references stay bold without a link. TASK-0085 only converted existing Title Case mentions; lowercase generic phrasings were left behind, which is this task's scope.

Edits, all in docs/terminology/glossary.md (anchors resolve to existing headings; add none):

1. Timed Event and Dates-Only Event definitions: replace informal availability domain with [Enabled Domain]; convert leading An event kind to a linked [Event Kind]; make A timed event / A dates-only-event self-references bold (**Timed Event**, **Dates-Only Event**); link [Timed Domain Mode].
2. Timed Slot: link [Timed Event], bold **Timed Slot** self-reference, link [Instant], [Slot Increment], [Enabled Domain].
3. Instant: link [Timed Slot] and [Civil Date]; keep rendered timezone descriptive.
4. Civil Date: bold **Civil Date** self-reference.
5. Slot Increment: link [Timed Slots] and [Timed Event].
6. Timed Domain Mode: split the two-sentence source line into one sentence per line; convert active-slot domain to a set of [Active Slots] with [Timed Event] link; link [Active Slot Range] and [Custom Domain Editing].
7. Ranged Domain Mode: link [Timed Domain Mode], rewrite active-slot domain to set of [Active Slots], link [Active Slot Range], rewrite current picked-date domain to the event's current [Picked Dates].
8. Custom Domain Mode: link [Timed Domain Mode], [Enabled Slots], [Custom Domain Editing], [Ranged Domain Mode]; rewrite active-slot domain as above; the event's [Active Slot Range].
9. Custom Domain Editing: link [Custom Domain Mode], [Active Slots], [Picked Dates].
10. Picked Dates: link [Timed Events]; bold **Picked Dates** self-reference; rewrite enabled-slot domain to [Enabled Slots]; keep date picker descriptive.
11. Enabled Slots: link [Timed Event], [Picked Date] (singular label, plural anchor), [Event Timezone] inside the parenthetical; bold **Enabled Slots**, **Picked Dates**, **Event Timezone** repeats in the derivation sentence.
12. Active Slots: link [Enabled Slots]; invariant code span unchanged.
13. Event Timezone: link [Picked Dates], [Enabled Slots], [Enabled Domain], [Active Slots]; bold repeat **Picked Dates**.
14. Display Timezone: link [Timed Slots], [Picked Dates], [Enabled Slots], [Active Slots].
15. Active Slot Settings: link [Timed Event], [Active Slots], [Timed Domain Mode], [Active Slot Range]; bold repeat **Active Slots**.
16. Active Slot Range: link [Active Slot Settings], [Ranged Domain Mode], [Active Slots], [Picked Dates].
17. Wipe Rule: link [Instant], [Enabled Domain] replacing enabled full-day domain; bold possessive **Picked Dates'** membership; keep example instants generic per TASK-0085 precedent.
18. Enabled Domain: bold same-block repeats at the picked-dates-themselves and timed-event-definition sentences; keep full civil-day slots descriptive.
19. Inactive Domain: link [Event Kind] in for-any-event-kind.
20. Dates-Only Event Owner Page and Timed Event Owner/Page/Scheduling entries plus response creation/editing page pairs: convert remaining unstyled availability response mentions to [Availability Response] links; convert platform identity recovery to [Platform Identity] recovery.
21. Event Owner: link [Event Guest] and [Event Settings]; bold **Event Owner** self-reference; link [Platform Identity] and [Event Sign-In].
22. Identity family: link [Platform Visitor] in Event Visitor; bold **Event Visitor** self-reference; link [Event Visitor Identity] in Event Visitor and Authenticated/Anonymous variants; link [Platform Identity]; bold **Event Visitor Identity** self-reference; link [Event Owner]'s in the browser-establishes sentence.
23. EVCC: link [Event Visitor] and [Availability Response]; keep legacy MongoDB response credentials as a deliberate umbrella, noted.
24. Cross-Device Access Transfer: link [Platform Identity] session.
25. Availability Response: link [Availability States] and [Event Visitor Identity]; Overlay: link [Availability Responses] and [Availability States].
26. Protected Response: drop the hyphen in the default availability-response access mode and link [Availability Response]; link [Event Guest], MongoDB [Availability Response Edit Credential], [Platform Identity]. Open Response: link all three concepts.
27. Availability Response Edit Credential: link [Protected Response] and [Platform Identity].
28. Blind Availability: link [Availability Responses]/[Event Owner]/bold repeat; Schedule Overlap: bold **Schedule Overlap** view replacing schedule-overlap view; link [Availability Responses].
29. Timed Grid: link [Timed Slots]; Enabled Inactive Slot: link [Enabled Slot] lead, keep not-active bare, link [Custom Domain Editing], [Active Slots], [Disabled Padding Cells]; Saved Active-Range Band: link [Active Slots]/[Enabled Domain] with bold repeat, [Custom Domain Editing]; Event Settings: expand the enumeration into individually linked [Event Kind], [Picked Dates] (replacing date selection), [Event Timezone] and [Display Timezone], [Event Time Format] and [Display Time Format], [Active Slot Settings], [Active Slot Range]; link [Availability Response] in its lead sentence.
30. Legend: link [Event Kinds]; Freemium: capitalize and bold **Freemium** self-reference; Event Settings lead handled above.

Deliberately left as-is (record in notes): calendar dates, date picker, rendered/display-local/non-picked date descriptives, legacy MongoDB response credentials, example instants after the Wipe Rule link, unavailable treatment in Disabled Padding Cell, and the Scheduling pages' owner shorthand following an explicit Event Owner link.

Verification: npm run lint:markdown and npm run format:markdown:check from the repo root; manual anchor-target audit (every link target exists as a heading); visual diff review.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Followed the TASK-0085 conventions exactly: a definition block up to the Authoritative context line is one styling unit; first occurrence links with a natural-inflection label (singular labels may target plural anchors, possessives stay inside the markers); same-block repeats bold across the whole inflection; self-references bold without links.

Plan deviations, all within approved scope: split Timed Domain Mode's two-sentence source line to satisfy one sentence per physical source line; converted additional genuine cross-references found in a final plain-text sweep (Event Guest body, Availability Editing and Availability Editor leads, Event Time Format's Display Time Format reference, both owner pages' Event Settings mentions, Show all hours' Timed Grid and Saved Active-Range Band references, Disabled Padding Cell's Enabled Slot lead, response editing pages' Availability Response leads, Event Owner Edit Token's Event Owner/Event Settings references, Shared Link self-reference, Grid Pointer's Availability Editing reference, both scheduled-span entries' Event Scheduled Span self-references, Disabled Domain's Timed Grids plural).

Deliberately left as-is with rationale: calendar dates and date picker (descriptive UI phrasing, no promotion approved), rendered/display-local/non-picked descriptives in Instant, Projected Date Column, and Disabled Domain, legacy MongoDB response credentials in EVCC (an umbrella distinct from the single Availability Response Edit Credential), example instants and enabled day near Wipe Rule and End-of-Day Boundary per TASK-0085 precedent, unavailable treatment in Disabled Padding Cell, Scheduling pages' owner shorthand after an explicit Event Owner link, and platform-wide sign-in contrast in Event Sign-In.

Verification evidence: npm run lint:markdown and npm run format:markdown:check pass from repo root; scripted audit confirms all 56 fragment anchors resolve against existing headings with zero misses; ripgrep confirms zero remaining occurrences of availability domain, picked-date domain, active-slot domain, enabled-slot domain, enabled full-day domain, schedule-overlap view, or date selection as standalone concepts; git diff shows no changed headings and no changed Authoritative context lines. Net diff: 85 insertions, 84 deletions (the +1 is the Timed Domain Mode sentence split).

Review correction: the original pass bolded three non-repeat occurrences and dropped two words. Fixes: Wipe Rule's first-occurrence **Picked Dates'** became the link [Picked Dates'](#picked-dates); the scheduled-span entries' **Event Scheduled Span** became [Event Scheduled Span](#event-scheduled-span) links (Event Scheduled Span is a distinct heading, not a self-reference); Availability Response's possessive now sits inside the label as [Event Guest's](#event-guest); Platform Identity's "an event or availability-response identifier" phrasing was restored to "an event or [Availability Response] identifier". lint:markdown, format:markdown:check, and the anchor audit re-pass.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: OpenCode
created: 2026-08-27 13:53
---
Follow-ups noted during execution, deliberately not fixed here (out of scope by user decision): (1) FR-086's title and body use the rejected aliases specific-times grids/mode even though the glossary records them as rejected for the scheduling page. (2) docs/requirements/migration/backlog-fr-inventory-candidates/CAND-203.md links glossary.md#scheduled-event-time, an anchor that does not exist because Scheduled Event Time is recorded as a rejected alias inside the Event Scheduled Span entry rather than as a heading.
---

author: OpenCode
created: 2026-08-27 14:36
---
Review corrections applied after Done (still documentation-only, same scope): (1) Wipe Rule: `**Picked Dates'**` was a first occurrence of Picked Dates in the block, so it is now the link [Picked Dates'](#picked-dates) per AC #2. (2) The Timed/Dates-Only Event Scheduled Span entries bolded **Event Scheduled Span**, but Event Scheduled Span is a distinct glossary heading, so both now link [Event Scheduled Span](#event-scheduled-span); the earlier notes mislabeled these as self-references. (3) Availability Response: possessive moved inside the link label as [Event Guest's](#event-guest) to match the stated possessives-inside-markers convention. (4) Platform Identity: restored the silently dropped "an event or" wording before the [Availability Response] link, which had changed meaning. Re-verified: npm run lint:markdown and npm run format:markdown:check pass; anchor audit passes with #event-scheduled-span now used.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Canonicalized all remaining generic term references inside docs/terminology/glossary.md so the glossary itself complies with docs/terminology/README.md (link first occurrence per block with an inflection-matching label; bold same-block repeats covering the whole inflection; bold-no-link self-references), building on TASK-0085's Title Case pass.

What changed: converted unstyled lowercase cross-references across roughly 45 definition lines into canonical fragment links or bold repeats, including the identity family (Platform Visitor, Event Visitor variants, Event Visitor Identity, EVCC, Cross-Device Access Transfer, Event Guest/Owner family), the slot-domain family (Timed/Dates-Only Event, Timed Slot, Instant, Civil Date, Slot Increment, domain modes and editing, Picked Dates, Enabled Slots, Active Slots, timezones, Active Slot Settings/Range, Wipe Rule), the response-access family (Blind Availability, Availability Response/State/Overlay, Protected/Open Response, Availability Response Edit Credential), grid-rendering terms (Timed Grid, Enabled Inactive Slot, Disabled Padding Cell, Projected Date Column, Saved Active-Range Band, Show all hours), and Event Settings' scope enumeration. Informal compound phrasings were rewritten meaning-preservingly onto existing terms: availability domain to Enabled Domain, active-slot domain to a set of Active Slots, picked-date domain to the event's current Picked Dates, enabled full-day domain to Enabled Domain, schedule-overlap view to a Schedule Overlap view, enabled-slot domain to Enabled Slots, and date selection to Picked Dates in the Event Settings enumeration. Also split one two-sentence source line in Timed Domain Mode to honor one sentence per physical source line. No new terms were added; no promotions were made (user decision).

Deliberately kept as descriptive prose: calendar dates, date picker, rendered/display-local/non-picked descriptives, legacy MongoDB response credentials, example instants/enabled day near Wipe Rule and End-of-Day Boundary, unavailable treatment, Scheduling pages' owner shorthand, platform-wide sign-in contrast.

Verification: npm run lint:markdown and npm run format:markdown:check pass at repo root; scripted audit shows all 56 fragment anchors resolve against existing headings; ripgrep confirms zero leftover informal-phrase occurrences; git diff confirms headings and Authoritative context lines untouched (85 insertions, 84 deletions). Documentation-only change, so unit/e2e tests are exempt per BACKLOG_WORKFLOW.md.

Follow-ups recorded as task comments (not fixed): FR-086 title/body use the rejected specific-times aliases; CAND-203 links a nonexistent glossary.md#scheduled-event-time anchor. Changes are left uncommitted for user review.
<!-- SECTION:FINAL_SUMMARY:END -->
