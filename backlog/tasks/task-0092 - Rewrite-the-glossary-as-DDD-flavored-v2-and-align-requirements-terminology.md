---
id: TASK-0092
title: Rewrite the glossary as DDD-flavored v2 and align requirements terminology
status: Done
assignee:
  - opencode
created_date: '2026-08-28 16:51'
updated_date: '2026-08-28 17:37'
labels:
  - documentation
  - terminology
dependencies: []
references:
  - >-
    backlog/tasks/task-0091 -
    Canonicalize-wording-deferred-by-TASK-0089-as-generic-or-ambiguous.md
  - >-
    backlog/tasks/task-0087 -
    Align-code-naming-with-canonical-glossary-terminology.md
  - frontend/src/components/schedule_overlap/ToolRow.vue
  - server/db/events.go
  - server/postgres/repository.go
documentation:
  - docs/terminology/README.md
  - docs/terminology/glossary.md
  - docs/requirements/README.md
  - docs/requirements/AGENTS.md
  - docs/requirements/functional/README.md
  - docs/design/architecture/adr/ADR-010.md
  - BACKLOG_WORKFLOW.md
priority: medium
type: docs
ordinal: 97500
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Replace the staged annotated draft of docs/terminology/glossary.md with a DDD-flavored v2 and sweep requirements terminology to match. All decisions below were made with the product owner on 2026-08-28; do not relitigate them. Implement wording exactly; where a detail is marked VERIFY, confirm against the linked requirement before applying.

# Outcome
A single controlled vocabulary grouped by bounded context (Scheduling Domain, Event Pages, Identity & Access, Responses, Presentation, Product, Text appendix), plus requirements (FR/QR), ADR-010, the requirements README index, and PLUGIN_API_README.md prose updated to the new canonical terms. Documentation-only change: no runtime code, tests, payload shapes, or code identifiers change.

# Decision record
1. Event Response family: Availability Response → Event Response; Protected Response → Protected Event Response; Open Response → Open Event Response; Availability Response Overlay → Event Response Overlay; Availability Response Edit Credential → Event Response Edit Credential, scoped explicitly to legacy MongoDB events. Rationale: owner chose the shorter event-aligned family name; churn accepted.
2. Availability State → Availability Status (values Available, If needed, Unavailable unchanged; canonical spellings preserved). Rationale: status fits the displayed enum; "Slot Availability Status" was rejected because states attach to dates too.
3. Slot model: Slot = the selectable unit of an event's Enabled Domain. Timed Slot → Time Slot; new Date Slot for Dates-Only Events. Active Slots generalizes to both kinds. Timed Event, Timed Grid, Timed Event Page names unchanged. Rationale: parallel attributive-noun pair chosen for grammar; owner confirmed.
4. Domain folding: Enabled Slots folds into Enabled Domain (Timed units are Time Slots). Active Domain, Inactive Domain, Disabled Domain fold away: Active Slots serves both kinds; Inactive Slots = Enabled minus Active (absorbs Enabled Inactive Slot); everything outside the Enabled Domain is presentation-level Disabled (Padding Cells on Timed Grids, non-picked dates with Disabled Status on Dates-Only Grids). Rationale: owner flagged this sprawl; folding reduces four overlapping domain terms to two.
5. Slot Increment → Slot Duration: the length of each Time Slot; consecutive Time Slots of that duration tile each Event Picked Date from 00:00 inclusive through the next 00:00 exclusive in the Event Timezone. FR-074/FR-075/FR-050/FR-010/FR-103 phrasing changes from "at the ... Slot Increment" to duration phrasing. Rationale: owner correction — it is the slot length (FR-050's 09:00–17:00 @ 15 min example proves it), not an interval between slots.
6. Ranged Domain Mode → Range Timed Domain Mode; Custom Domain Mode → Custom Timed Domain Mode; Custom Domain Editing → Custom Timed Domain Editing (still an editing surface, not a Page). Rationale: owner wants explicit Timed scoping and "Range" spelling.
7. Picked Dates → Event Picked Dates (owner's candidate term; keeps the picker connotation).
8. Enabled Inactive Slot → Inactive Slot; Disabled Padding Cell → Padding Cell (definition must explain padding-only and that it is not a Slot). Rationale: Slot = domain unit, Cell = grid rendering; resolves TASK-0091 items FR-093/FR-094.
9. Saved Active-Range Band: term dropped. FR-014 states the axis rule directly: on Custom Timed Domain Mode event pages the read-only grid's time axis covers the saved Active Slots' extent, falling back to the full Enabled Domain when no Active Slots exist; the full civil-day axis appears only in Custom Timed Domain Editing or with the Show all hours Option. Rationale: owner could not attach meaning to the old name; it is a one-off rendering rule.
10. Event Scheduled Span → Event Occurrence Span (subtypes: Timed Event Occurrence Span, Dates-Only Event Occurrence Span). "Scheduled Event Time" stays a rejected alias.
11. Cross-Device Access Transfer → Access Transfer. Definition states browser-to-browser, PostgreSQL events, delegating (never transferring) authority. Rationale: owner pick; "device" was inaccurate.
12. Blind Availability → Blind Availability Mode; Freemium → Freemium Mode; Legend → Legend Section; heading "Show all hours" Option (Option outside the quotes); Shared Link → Event Link. Owner confirmations.
13. Platform Identity → Platform Visitor Identity; new entries Authenticated Platform Visitor, Anonymous Platform Visitor, Platform Sign-In. Rationale: parallel with Event Visitor Identity; owner chose the rename.
14. Magic Link → Sign-In Link (FR-031/FR-032). Owner confirmation.
15. Availability Editor: term dropped. FR-005, FR-066, FR-102 reworded to name the page and grid precisely. Owner leaned this way and confirmed.
16. Kept terms: Civil Date (contradictory definition fixed: a bare calendar date with no time or timezone of its own, interpreted only when paired with a timezone such as Event Timezone or Display Timezone); Availability Editing (name kept; definition reworded to "creating, viewing, and editing Event Responses", dropping "selecting"); Schedule Overlap (name kept; non-circular definition: the event-page view aggregating the Availability Statuses of included Event Responses per Slot); EVCC acronyms; Event Owner Edit Token; Wipe Rule (kept as a named invariant; FR-015 and FR-057 gain the name); FR-024 (kept; the glossary must cite FR-024 instead of frontend code files as authoritative context).
17. End-of-Day Boundary: always labeled 24:00, including the 12-hour picker; FR-091's 12-hour `12 AM` boundary rule is removed. Owner decision for consistency.
18. FR-012 actor fix: "Anyone with an event link shall be able to save, replace, or clear" becomes Event Owner authority only, matching the owner-only Scheduling Pages and FR-016/FR-085/FR-086. Deliberate change to an accepted requirement; VERIFY during execution whether the app already enforces owner-only span saving and note a finding if it does not (no code change in this task).
19. Untouched by design: FR-062/FR-081/FR-082 MongoDB-preservation clauses (MongoDB is live: server/db/* uses mongo-driver, PostgreSQL lives in server/postgres/*); FR-046/FR-047 control-label wording (mismatch with the UI's `Shown in` text at frontend/src/components/schedule_overlap/ToolRow.vue:147 is a recorded follow-up, not changed here); code identifiers and API field names (do not rename AvailabilityResponse types, eventVisitorId, etc.).
20. Scope: docs/terminology/glossary.md rewrite; docs sweep across docs/requirements/** (bodies and titles), docs/design/architecture/adr/ADR-010.md, docs/requirements/README.md index rows for retitled FRs, and PLUGIN_API_README.md prose (terminology only, no payload changes). The staged comment-annotated glossary draft is fully replaced; review comments must not survive into the final text.

# Relationship to other tasks
- TASK-0091 (deferred wording): this task decides all seven of its items (FR-030 → Active Time Slots; FR-051 → Event Picked Dates participle; FR-058 → Enabled Domain; FR-086 → Timed Grid; FR-093 → Active Slots; FR-094 → Disabled Status or Padding Cell, VERIFY cell semantics; FR-099 → Disabled Status). Record the decisions as a comment on TASK-0091 during execution, apply its items here, then finalize TASK-0091 per its own acceptance criteria.
- TASK-0087 (code naming): must execute AFTER this task and align code identifiers to the new canonical terms; its description may need updating to the new names.

# Detailed edit plan
The per-entry glossary inventory, per-file edit list, FR title table, and verification commands are recorded in this task's implementation plan. Read docs/terminology/README.md, docs/requirements/AGENTS.md, docs/requirements/functional/README.md, and BACKLOG_WORKFLOW.md before editing. Authoring rules: one sentence per physical source line; link first occurrence per paragraph/list item/table cell; bold repeats; possessives inside markup; regenerate the glossary TOC after any heading change.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Every item in the task's decision record is applied: the glossary uses exactly the new canonical terms and section structure, dropped terms (Enabled Slots, Active/Inactive/Disabled Domain, Saved Active-Range Band, Availability Editor) are gone, and every entry keeps a concise definition plus an authoritative FR/ADR/spec context (no code-file contexts).
- [x] #2 All docs files (FRs, QR-001/002/003/011, ADR-010, docs/requirements/README.md index, PLUGIN_API_README.md prose) use the new canonical terms with first-occurrence glossary links per docs/terminology/README.md; every changed FR title is mirrored in the requirements README index.
- [x] #3 Semantic corrections are applied exactly: FR-012 restricted to Event Owner authority, FR-091 boundary always labeled 24:00, Slot Duration phrasing in FR-074/075/050/010/103, FR-014 axis rule replacing Saved Active-Range Band, FR-018 gains the EVCC-converse sentence, Wipe Rule named in FR-015 and FR-057.
- [x] #4 FR-062/FR-081/FR-082 MongoDB-preservation clauses and every other requirement's intent, scope, components, status, and ID are unchanged (except the deliberate changes in the semantic-corrections criterion); no runtime code, tests, build config, or payload shapes are modified.
- [x] #5 TASK-0091's deferred-wording items (FR-030, FR-051, FR-058, FR-086, FR-093, FR-094, FR-099) are applied per this task's decisions and recorded as a comment on TASK-0091, which is then finalized per its own acceptance criteria.
- [x] #6 Every glossary anchor referenced anywhere under docs/ resolves to a real glossary heading, and a ripgrep sweep for retired term spellings (full list in the implementation plan) returns zero matches outside the glossary's own alias and rejected-variant notes.
- [x] #7 npm run format:markdown:check passes at the repository root, and git status shows changes only under docs/, PLUGIN_API_README.md, and Backlog-managed records.
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
# Implementation Plan (research completed 2026-08-28; decisions locked with the owner — see description)

## Read first
- docs/terminology/README.md, docs/requirements/AGENTS.md, docs/requirements/functional/README.md, BACKLOG_WORKFLOW.md.
- The working-tree glossary draft (staged) contains the owner's inline review comments; they are inputs only and must not survive into the final text.

## Target glossary structure (bounded-context sections, in order)
1. **Text Terms** (appendix): Unicode Normalization Form C (NFC) — unchanged.
2. **Scheduling Domain**: Event Kind; Timed Event; Dates-Only Event; Event Picked Dates (renamed from Picked Dates); Slot (new); Time Slot (renamed from Timed Slot); Date Slot (new); Instant; Civil Date; Slot Duration (renamed from Slot Increment); Enabled Domain (absorbs Enabled Slots); Active Slots (generalized to both kinds); Inactive Slots (new; absorbs Enabled Inactive Slot and Inactive Domain); Range Timed Domain Mode (renamed); Custom Timed Domain Mode (renamed); Custom Timed Domain Editing (renamed); Active Slot Settings; Active Slot Range; Event Timezone; Display Timezone; Wipe Rule; End-of-Day Boundary; Event Occurrence Span (renamed) with Timed Event Occurrence Span and Dates-Only Event Occurrence Span subtypes; Event Settings (moved here from Grid Rendering Terms).
3. **Event Pages**: Timed Event Page and Dates-Only Event Page as umbrellas; each carries the "one physical route renders every event-page surface in this family" sentence once. Variants (owner, scheduling, response creation, response editing — 8 entries) deduplicated and shortened; scheduling-page definitions simplified; "specific-times page/mode/editing" stay listed as rejected aliases.
4. **Identity & Access**: Platform Visitor; Platform Visitor Identity (renamed from Platform Identity); Authenticated Platform Visitor (new); Anonymous Platform Visitor (new); Event Visitor; Event Visitor Identity; Authenticated Event Visitor (reworded: "signed in with a Platform Visitor Identity"); Anonymous Event Visitor (consistent wording); Event Guest; Event Owner; Event Visitor Control Credential (EVCC); Granted Event Visitor Control Credential (Granted EVCC); Event Owner Edit Token; Event Response Edit Credential (renamed; legacy-MongoDB-scoped); Access Transfer (renamed from Cross-Device Access Transfer); Event Sign-In; Platform Sign-In (new); Sign-In Link (renamed from Magic Link).
5. **Responses**: Event Response (renamed from Availability Response); Availability Status (renamed from Availability State) with Available / If needed / Unavailable children (canonical spellings unchanged); Protected Event Response (renamed); Open Event Response (renamed); Blind Availability Mode (renamed); Event Response Overlay (renamed); Availability Editing (kept; reworded); Event Link (renamed from Shared Link).
6. **Presentation**: Timed Grid; Dates-Only Grid; Projected Date Column; Grid Pointer; Padding Cell (renamed from Disabled Padding Cell); Disabled Status (kept); Schedule Overlap (kept; redefined); Legend Section (renamed from Legend); "Show all hours" Option (renamed heading; Option outside the quotes).
7. **Product**: Freemium Mode (renamed from Freemium).

Dropped entirely: Enabled Slots, Active Domain, Inactive Domain, Disabled Domain, Saved Active-Range Band, Availability Editor. Regenerate the TOC after headings are final.

## Definition fixes (apply this intent verbatim; respect one-sentence-per-line)
- Slot: "The selectable unit of an event's Enabled Domain. A Timed Event's Slots are Time Slots; a Dates-Only Event's Slots are Date Slots."
- Time Slot: "A discrete availability unit of a Timed Event. A Time Slot has an Instant; consecutive Time Slots of the Slot Duration tile each Event Picked Date from 00:00 inclusive through the next 00:00 exclusive in the Event Timezone."
- Date Slot: "The Dates-Only selectable unit: an Event Picked Date offered for a response."
- Slot Duration: "The length of each Time Slot of a Timed Event."
- Enabled Domain: "Everything generated from Event Picked Dates. For a Timed Event, full-civil-day Time Slots of the Slot Duration through the next 00:00 exclusive in the Event Timezone; for a Dates-Only Event, the Event Picked Dates as Date Slots." State that everything outside it is not answerable and renders as the Disabled treatment (Padding Cells or disabled dates).
- Active Slots: "The canonical subset of the Enabled Domain that respondents can answer on. Every Active Slot belongs to the Enabled Domain." Timed Events maintain them via the domain modes; VERIFY FR-051/FR-052 to state the Dates-Only Active Slot behavior precisely before writing that sentence.
- Inactive Slots: "The Enabled Domain minus Active Slots. Editable in Custom Timed Domain Editing but not respondent-selectable; visually distinct from both Active Slots and Padding Cells."
- Civil Date: "A bare calendar date with no time-of-day or timezone of its own; it gains interpretation only when paired with a timezone, such as the Event Timezone or Display Timezone."
- Instant: replace "clock label" with "rendered clock time".
- Display Timezone: remove the `Shown in` selector mention; reference the event-page Display Timezone control (FR-047).
- Active Slot Settings: "contains the Timed Domain Mode and, depending on the mode, an Active Slot Range (Range Timed Domain Mode) or the full Active Slots selection (Custom Timed Domain Mode)."
- Custom Timed Domain Mode: "its Active Slots are an explicitly chosen subset of the Enabled Domain, maintained through Custom Timed Domain Editing." Drop the mode-switching sentence; FR-054 governs regeneration.
- Wipe Rule: open with "On save, each Active Instant outside the Enabled Domain is dropped." Keep the next-day 00:00/00:30 example and the Event Timezone note.
- End-of-Day Boundary: "labeled `24:00` in both 12-hour and 24-hour pickers"; keep the Projected Date Column rendering sentence.
- Event Guest: drop the Event Owner sentence (atomic definitions; the Owner entry already references Guest).
- EVCC: replace "PostgreSQL Event Visitor" with "an Event Visitor of a PostgreSQL event"; add "It authorizes response management only; it never authorizes Event Settings edits."
- Granted EVCC: clarify delegation — the source Event Visitor Identity remains the responses' owner; the Granted EVCC delegates management authority until the source revokes it or the target browser's data is cleared.
- Event Response Overlay: "The elevated rendering of the Event Response being edited, shown above other Event Responses so its Availability Statuses stay visible and editable where other responses would otherwise paint over them."
- Schedule Overlap: "The event-page view that aggregates the Availability Statuses of the included Event Responses per Slot." Keep the Available/If needed equality and Unavailable exclusion.
- Availability Editing: "The activity of creating, viewing, and editing Event Responses on the response creation and editing pages."
- Event Link: "An event-scoped link whose validity gates disclosure of event metadata, respondent names, and availability data."
- Legend Section / "Show all hours" Option / Blind Availability Mode / Freemium Mode: definitions unchanged apart from names.
- Event Time Format / Display Time Format: cite FR-024 (and FR-046) as authoritative context instead of frontend code files.
- Authenticated/Anonymous Platform Visitor (new): "A Platform Visitor signed in with a Platform Visitor Identity" / "A Platform Visitor not signed in with a Platform Visitor Identity."
- Platform Sign-In (new): platform-wide authentication, distinct from Event Sign-In; gated application-wide (FR-007).

## Execution order
1. Rewrite docs/terminology/glossary.md per the structure and fixes above (replaces the staged annotated draft).
2. Update docs/requirements/README.md index rows for every retitled FR (title table in the task notes).
3. Apply semantic corrections: FR-012, FR-091, FR-074, FR-075, FR-050, FR-010, FR-103, FR-014, FR-015, FR-057, FR-018 (details in notes).
4. Apply the mechanical rename sweep across docs/requirements/functional/fr/*.md, docs/requirements/quality/qr/*.md, docs/design/architecture/adr/ADR-010.md, and PLUGIN_API_README.md prose: term swap plus first-occurrence link per paragraph/list item/table cell, bold repeats, possessives inside markup.
5. Apply TASK-0091 items (FR-030, FR-051, FR-058, FR-086, FR-093, FR-094, FR-099); record the seven decisions as a comment on TASK-0091 and finalize it per its own acceptance criteria.
6. Run verification; fix findings; re-run until clean.

## Verification
- Anchor check: extract every `terminology/glossary.md#<slug>` reference under docs/ and assert each slug matches a heading anchor in the new glossary (small node or rg + sort -u script).
- Retired-spelling sweep (list in notes) must return zero matches outside the glossary's own alias/rejected notes; review hits manually since some new terms contain old substrings.
- `npm run format:markdown:check` at the repository root; run `npm run format:markdown` only on changed files if it flags them.
- `git status` must show changes only under docs/, PLUGIN_API_README.md, and Backlog-managed records; commit is a separate, owner-initiated step.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
# Working Notes

## FR title changes (IDs and filenames permanent; mirror every change in docs/requirements/README.md)
- FR-001 → Preserve Event Guest Event Responses
- FR-002 → Project Time Slots into their Projected Date Columns
- FR-003 → Delete Event Responses
- FR-004 → Require a non-empty Event Response
- FR-005 → Render the edited Event Response elevated above other Event Responses
- FR-006 → Keep Availability Statuses mutually exclusive
- FR-007 → Hide Platform Sign-In entry points when Platform Sign-In is disabled (VERIFY the app-wide flag gates platform sign-in only)
- FR-008 → Initialize the Range Timed Domain Mode Active Slot Range
- FR-010 → Support Custom Timed Domain Editing
- FR-012 → Maintain an optional Event Occurrence Span (plus owner-only reword)
- FR-013 → Preserve Time Slot Instants across Display Timezone changes
- FR-017 → Preserve Time Slots across daylight saving transitions
- FR-020 → Keep the mobile selected Time Slot tooltip available after scrolling
- FR-030 → Persist only Active Time Slots in browser storage
- FR-031 → Send registration Sign-In Links by email
- FR-032 → Complete registration after a Sign-In Link sign-in
- FR-043 → Place Time Slot tooltips at the selected slot
- FR-044 → Explain Protected Event Responses
- FR-050 → Generate Active Slots for newly added Event Picked Dates
- FR-051 → Keep Active Slots off newly added Event Picked Dates in Custom Timed Domain Mode
- FR-052 → Remove Active Slots for removed Event Picked Dates
- FR-053 → Preserve Active Slots when converting to Custom Timed Domain Mode
- FR-054 → Regenerate Active Slots when converting to Range Timed Domain Mode
- FR-058 → Reject availability outside the Enabled Domain
- FR-059 → Validate and normalize Event Response names
- FR-060 → Create Protected Event Responses by default
- FR-061 → Allow Event Guests to open their Event Responses
- FR-064 → Keep Protected Event Responses in Schedule Overlap
- FR-066 → Filter the response-editing grid by Event Response selection
- FR-067 → Create Timed Events in Custom Timed Domain Mode
- FR-069 → Calculate Schedule Overlap from Availability Statuses
- FR-070/071/072 → … when Freemium Mode is disabled (all three)
- FR-075 → Generate initial Range Timed Domain Mode Active Slots
- FR-082 → Transfer Platform Sign-In from a PostgreSQL event to another browser
- FR-084 → Preserve Blind Availability Mode privacy
- FR-085 → Exclude Timed Event Occurrence Span cells from Availability Editing
- FR-086 → Render the Timed Event Occurrence Span on Timed Grids
- FR-092 → Keep the Legend Section visible without Event Responses
- FR-097 → Include the If needed item in both Legend Sections
- FR-102 → Use fixed wording for personal availability on response editing pages
- FR-103 → Expose the editable Slot Duration in new-event Advanced options
- FR-104 → Align the Legend Section label with the Responses label
- FR-113 → Require a non-empty Timed Event Occurrence Span before saving

Titles not listed stay unchanged (e.g., FR-011, FR-016, FR-024, FR-027, FR-046, FR-047, FR-074, FR-081, FR-083, FR-091, FR-095, FR-099, FR-111; FR-111/FR-048 only swap the linked heading).

## Retired spellings to sweep (zero expected outside glossary alias/rejected notes)
Availability Response(s), Availability State(s), Availability Response Overlay, Availability Response Edit Credential, Protected Response(s), Open Response(s), Timed Slot(s), Slot Increment, Enabled Slots, Enabled Inactive Slot(s), Active Domain, Inactive Domain, Disabled Domain, Disabled Padding Cell(s), Saved Active-Range Band, Ranged Domain Mode, Custom Domain Mode, Custom Domain Editing, Picked Dates (as controlled term; review generic-prose hits individually), Availability Editor, Magic Link(s), Platform Identity, Cross-Device Access Transfer(s), Event Scheduled Span, Scheduled Span, Blind Availability (bare), Shared Link, Freemium (bare), Legend (bare controlled-term refs). Do NOT sweep: Availability Editing, Civil Date, Event Kind, Active Slots (kept). Use word-boundary patterns and review hits; "Event Response" contains "Response", so naive substring matching over-matches.

## Per-file specifics
- FR-012: "The Event Owner shall be able to save, replace, or clear an optional Event Occurrence Span. Other Event Visitors shall not be able to change it." Keep the Timed/Dates-Only ACs with renamed subtypes.
- FR-091: remove the 12-hour `12 AM` boundary sentence; hours 1–12 keep 12-hour labels; the End-of-Day Boundary is always `24:00`.
- FR-074: "For each Event Picked Date of a Timed Event, the system shall derive Enabled Domain Time Slots of the configured Slot Duration, tiling the full Civil Date from 00:00 inclusive through the next 00:00 exclusive in the Event Timezone."
- FR-075/FR-050: "generate Active Slots of the configured Slot Duration for every Event Picked Date from its Active Slot Range" (start inclusive, end exclusive — the last Active Slot starts at end − Slot Duration).
- FR-010: "every Time Slot of the configured Slot Duration from 00:00 inclusive through the next 00:00 exclusive".
- FR-103: Slot Duration wording; control behavior unchanged.
- FR-014: rewrite with Inactive Slot / Padding Cell; replace the Saved Active-Range Band AC with the axis rule (read-only Custom Timed Domain Mode pages: axis covers saved Active Slots' extent, fallback full Enabled Domain; full civil day only in Custom Timed Domain Editing or with the "Show all hours" Option).
- FR-015: "…shall discard each Active Slot outside that event's Enabled Domain (the Wipe Rule)."
- FR-057: "…keep its Event Picked Dates stable, rebuild its Enabled Domain, and apply the Wipe Rule to its Active Slots." Keep the response-slot discard sentence.
- FR-018: append "The Event Visitor Control Credential shall not authorize Event Settings edits."
- FR-017/FR-013: optionally align "displayed clock label" with the Instant definition's "rendered clock time".
- FR-086: body reword — "On the Timed Event Scheduling Page, selecting a range on the Timed Grid shall set the Timed Event Occurrence Span directly, without painting availability."
- FR-093: "active grid cells" → Active Slots. FR-094: VERIFY which grid element the mobile tap targets (Padding Cell vs Disabled Status area) and use that term.
- FR-051: body reworded around Event Picked Dates + Custom Timed Domain Mode; keep intent (new dates contribute no Active Slots).
- QR-001/QR-002: "shared link" → Event Link. QR-003: VERIFY body for legacy credential naming; terminology only. QR-011: Access Transfer.
- ADR-010: Access Transfer, Blind Availability Mode, Event Response family, Platform Visitor Identity; decisions unchanged.
- PLUGIN_API_README.md: prose term swap only; payload field names and shapes unchanged (AGENTS.md rule about window.postMessage shapes).
- docs/requirements/functional/README.md: check for term usage and update if present.

## Follow-up candidates (do NOT execute inside this task)
- MongoDB event retirement (proposed task): afterwards remove FR-062/081/082 preservation clauses and the legacy Event Response Edit Credential entry.
- FR-046/FR-047 label reconciliation: FR-047 mandates `Display time zone`, FR-046 `Display time format`, but the event-page toolbar shows `Shown in` (frontend/src/components/schedule_overlap/ToolRow.vue:147). Needs a product decision; the glossary stays label-free either way.
- TASK-0087 (code naming): re-scope to the new canonical terms after this task lands.
- FR-012 owner-only enforcement: if the app lets non-owners save spans, file a bug task; this change states intended behavior only.

## Session handoff context
- Decisions were made interactively with the owner on 2026-08-28; the description is the authoritative decision record.
- The staged glossary diff (+109 comment lines) is the owner's review annotations; the rewrite supersedes them. Re-staging is the owner's step at commit time.
- Nothing has been edited in docs/ yet; this session only researched and decided.

## VERIFY findings (recorded during execution)
- FR-012 owner-only enforcement: NOT enforced today. The MongoDB handlers saveTimefulSchedule/clearTimefulSchedule carry no authorization and a source comment states scheduling is intentionally available to anyone with its link (server/routes/events.go:109-164); the PostgreSQL handler postgresUpdateSchedule also performs no owner authorization (server/routes/postgres_event_routes.go:336-377). Per the plan, no code change here; file a bug task to enforce owner-only span saving (follow-up candidate).
- FR-094 tap target: verified in useTimedGridInteractions tests and cell-state code that the mobile tap targets non-selectable grid cells; the `Disabled` wording belongs to cells outside the event dates (SpecificTimesInstructions.vue:27, ColorLegend.vue). Reworded to Disabled Status with Padding Cells explicitly included.
- Dates-Only Active Slot behavior (Active Slots glossary entry): a Dates-Only Event has no domain modes and no inactive subset; every Date Slot of its Enabled Domain is answerable, so its Active Slots are exactly its Date Slots (FR-058 rejects values outside the Event Picked Dates).
- FR-007 flag scope: VITE_ENABLE_SIGN_IN gates sign-in/sign-up availability application-wide (frontend/src/utils/signInAvailability.ts, docs/environments.md); there is no separate event-sign-in flag, so the Platform Sign-In rename is consistent.
- QR-003 legacy credential naming: the MongoDB clause's unnamed "unpredictable edit token" now cites the canonical Event Response Edit Credential.

## Scope execution notes
- PLUGIN_API_README.md prose contains no controlled-term usages, so it needed no changes (payload shapes untouched).
- docs/requirements/migration/ CAND files: only broken glossary links in review-authored prose were repointed to the renamed anchors; Raw Source quotes and candidate titles keep their verbatim historical wording per the migration README's provenance rules.
- ADR-010 was retitled to "Source-Confirmed Access Transfers" and its index row in docs/design/README.md updated to match.
- Known remaining retired spellings outside this task's sweep scope (follow-up candidates, unchanged by design): ADR-003's "Freemium" title/body and its docs/design/README.md row, ADR-008/ADR-009 "Platform Identity" prose, docs/environments.md VITE_ENABLE_FREEMIUM flag name, and migration provenance text.
- docs/terminology/README.md examples were updated to the new canonical terms so its teaching examples resolve to real glossary anchors.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Rewrote docs/terminology/glossary.md as the DDD-flavored v2 controlled vocabulary grouped by bounded context (Text Terms appendix; Scheduling Domain; Event Pages; Identity & Access; Responses; Presentation with the time-format entries; Product), replacing the staged annotated draft and dropping Enabled Slots, Active/Inactive/Disabled Domain, Saved Active-Range Band, and Availability Editor. Applied the full decision record: the Event Response family, Availability Status, the Slot/Time Slot/Date Slot model, Slot Duration, Range/Custom Timed Domain Mode, Event Picked Dates, Inactive Slots, Padding Cell, Event Occurrence Span family, Access Transfer, Platform Visitor Identity family with new Authenticated/Anonymous Platform Visitor and Platform Sign-In entries, Sign-In Link, Event Link, Blind Availability Mode, Freemium Mode, Legend Section, the "Show all hours" Option, fixed Civil Date/Wipe Rule/Schedule Overlap/EVCC/Granted EVCC definitions, and FR-024-based authoritative contexts for the time formats. Swept every FR (titles, YAML, bodies), QR-001/002/003/011, ADR-010 (retitled), the requirements README index (113 FR rows plus QR-011), and the terminology README examples onto the new canonical terms with first-occurrence links, bold repeats, possessives inside markup, and one sentence per line. Applied semantic corrections exactly: FR-012 owner-only Event Occurrence Span authority, FR-091 always-24:00 boundary, Slot Duration phrasing in FR-074/075/050/010/103, FR-014 axis rule, FR-015/FR-057 Wipe Rule naming, FR-018 EVCC-converse sentence, and FR-005/FR-066/FR-102 Availability Editor removal. FR-062/081/082 MongoDB preservation clauses keep their intent; no runtime code, tests, build config, or payload shapes changed. TASK-0091's seven items were applied and recorded as a task comment, and TASK-0091 is finalized and Done. Verification: glossary anchor check over all of docs/ resolves (87 anchors, 0 broken, GitHub slug rules), word-boundary ripgrep sweeps over the in-scope files return zero retired spellings outside the glossary's own alias/rejected notes, npm run format:markdown:check and npm run lint:markdown pass at the repository root, and git status shows changes only under docs/ plus Backlog-managed records. PLUGIN_API_README.md needed no edits (no controlled-term usage). Known out-of-scope remainders are recorded in the task notes as follow-up candidates (ADR-003/008/009 spellings, FR-012 enforcement bug task, FR-046/FR-047 label reconciliation, TASK-0087 re-scoping). Commit and re-staging of the glossary remain owner-initiated steps.
<!-- SECTION:FINAL_SUMMARY:END -->
