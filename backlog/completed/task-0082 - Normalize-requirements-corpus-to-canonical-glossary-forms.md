---
id: TASK-0082
title: Normalize requirements corpus to canonical glossary forms
status: Done
assignee:
  - opencode
created_date: '2026-08-26 17:05'
updated_date: '2026-08-26 17:31'
labels:
  - terminology
  - documentation
dependencies: []
references:
  - docs/terminology/glossary.md
  - docs/terminology/README.md
  - docs/requirements/README.md
  - BACKLOG_WORKFLOW.md
  - tmp/handoff-casing-1.md
  - tmp/handoff-casing-2.md
  - tmp/handoff-link.md
modified_files:
  - docs/requirements/README.md
  - docs/requirements/AGENTS.md
  - docs/requirements/functional/README.md
  - docs/requirements/quality/README.md
  - docs/requirements/functional/fr/
  - docs/requirements/quality/qr/
priority: medium
type: docs
ordinal: 89000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Retroactive normalization that TASK-0078 explicitly deferred: make every requirement record, guide prose, and index title use controlled terms exactly as recorded in docs/terminology/glossary.md, and satisfy the first-occurrence linking rule from docs/terminology/README.md.

Audits are complete; findings are recorded in task notes (authoritative replacement table) and mirrored in tmp/handoff-casing-1.md, tmp/handoff-casing-2.md, tmp/handoff-link.md (ephemeral working copies — regenerate via the audit greps in the plan if missing).

Scope: docs/requirements/README.md (index tables), docs/requirements/AGENTS.md, docs/requirements/functional/{README.md,fr/*.md}, docs/requirements/quality/{README.md,qr/*.md}. OUT of scope: docs/requirements/migration/**, ADRs/specs, code, frontend AGENTS.md.

Policy rulings (full detail in implementation notes):
- Attributive hyphen compounds de-compound into spaced canonical form with canonical caps (timed-event → Timed Event, display-timezone changes → Display Timezone changes, protected-response recovery → Protected Response recovery, event-owner authority → Event Owner authority, days-only-event workload → Dates-Only Event workload).
- freemium → Freemium everywhere including feature-flag phrasing (user decision).
- The days-only family (QR-007, QR-008, README index) renames to Dates-Only Event (user decision).
- "Show all hours" keeps its non-title-case form even inside Title Case H1s and index cells (glossary beats heuristics; see FR-048).
- Titles stay synced as triples: YAML title: ≡ H1 (strip "N:", Title Case every word incl. articles/To/With, FR-004 "a" anomaly left as-is) ≡ README index cell (verbatim FM text); FR-035's stale index row gets replaced with its real title.
- Untouched by design: fenced example in requirements/README.md:37–53, status enum line 84, code spans/backticked UI labels (incl. `Display time zone` vs Display Timezone spelling — separate defect), bare "edit token(s)/edit credential" stems awaiting a semantic pass, generic bare "timezones"/"editing availability", FR-058:12 "dates-only value", quality/README ISO vocabulary.

Documentation-only change: per BACKLOG_WORKFLOW.md DoD, unit/e2e tests are not required. Mark In Progress only when executing.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 All controlled-term occurrences across scoped files use canonical glossary capitalization and spacing; the audit greps rerun in verification return zero unexplained matches.
- [x] #2 Every in-scope controlled-term mention that appears first in its paragraph, list item, or table cell links to ../../../terminology/glossary.md#<slug> (correct depth per location) and every anchor resolves against actual glossary H3 headings; no links added in headings, code spans, or foreign link labels.
- [x] #3 Each record keeps YAML title:, H1, and README index cell consistent (case-insensitive equality for FM vs index; H1 follows the all-words Title Case exemplar pattern), including fixing FR-035's stale index row.
- [x] #4 No "days-only" variant remains in scoped files; renamed usages read Dates-Only Event.
- [x] #5 Freemium is capitalized in all scoped occurrences, including feature-flag phrasing.
- [x] #6 "Show all hours" retains glossary casing everywhere it occurs in scope.
- [x] #7 Documentation formatting holds: one sentence per physical source line, table rows on one physical line, wrapped cross-line span in FR-065:19-20 merged onto a single line; prettier/eslint Markdown gates pass.
- [x] #8 Out-of-scope material unchanged: docs/requirements/migration/**, README:37-53 fenced example, status enum line, code-span UI labels, ISO vocabulary, all runtime code.
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
Execute in this order; it is scripted-codemod-friendly but every step has a human-checkable assertion.

0. Hygiene: read BACKLOG_WORKFLOW.md; mark task In Progress; confirm scope globs unchanged since audits (git status clean-ish, no new FR/QR files; if IDs > FR-084/QR-011 exist, re-audit them with the Part-D exclusions before touching).
1. CASING PASS: apply Part A (FM titles) + Part-B/H1/index span swaps mechanically. Preferred tooling: throwaway python3/node script in /tmp/opencode consuming a literal (file,line,old,new) table generated FROM the notes below; assert old-string occurs exactly once per target line (or expected count); any mismatch = stop and fix by hand, never fuzzy-replace whole lines. GNARLY LINES TO EDIT BY HAND: FR-065:19–20 merge; FR-073:21 stray-token deletion; FR-067 title decision (read body first); README.md:147 row swap; FR-024 FM expansion ripple into H1/index.
2. TITLE TRIPLE SYNC PASS: regenerate H1s from new FM via span-swaps + Title-Case rule; copy FM strings verbatim into matching README index cells (except FR-035 special). Assert: lowercase(FM)==lowercase(indexCell) for ALL 95 records; H1 pattern scan finds Title Case everywhere except documented exceptions (Show all hours x-row, FR-004 'a').
3. LINK PASS: walk Parts C1/C2/C3 adding markdown links around FIRST-in-block occurrences (block defn: paragraph/list-item/table CELL; reset per cell in tables; never in headings/code/fences/foreign labels). Reuse anchor slugs from glossary; after edits run checker: extract every glossary.md#anchor in scope → assert slug matches some glossary H3-derived heading; assert relative depth resolves (fr/qr ../../../, guides ../, index none-prefix).
4. VERIFY LOOP (rerun until silent):
   a. casing grep — the two inventory greps from handoff files rerun over scope minus migration/; expect zero except whitelist: fenced 37–53, enum line 84, backtick labels, bare credential stems, ISO tables, 'dates-only value', 'editing availability', generic timezones, H1 exceptions (Show all hours / FR-004 a).
   b. days-only and freemium[lowercase] greps over scope: zero.
   c. orphan check: controlled-term mention with NO matching-anchor link anywhere in its file whose first-block occurrence lacks link → must be empty.
   d. cd frontend && npm run format:markdown (covers prettier/eslint sentences-per-line Markdown gates; project DoD item #4).
   e. git diff review sampled at minimum: FR-013, FR-020, FR-024, FR-030, FR-035, FR-048, FR-051, FR-064, FR-065, FR-067, FR-072, FR-073, QR-007, QR-008, QR-011, README index block.
5. FINALIZE: record finalSummary (counts applied/excluded, FR-067 wording chosen, any reconciliations), tick ACs, status Done (leave unarchived; workflow forbids archival at finalization). Documentation-only ⇒ unit/e2e exempt.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
AUTHORITATIVE REPLACEMENT TABLE (merged from tmp/handoff-casing-1.md, casing-2, handoff-link.md; conflict-resolved).
Slug map (glossary anchors): event-guest event-owner event-owner-edit-token event-visitor platform-visitor platform-identity event-visitor-identity authenticated-event-visitor anonymous-event-visitor event-visitor-control-credential-evcc granted-event-visitor-control-credential-granted-evcc availability-response availability-response-overlay availability-state availability-editing availability-editor protected-response open-response blind-availability timed-event timed-slot instant civil-date slot-incr… (full set: derive from glossary H3 headings; verified specials listed here)
Link path prefix: fr/qr files ../../../terminology/glossary.md#<slug>; guides docs/requirements/*.md → ../terminology/glossary.md#…; index README cells → terminology/glossary.md#…. Link labels keep the surrounding local text verbatim (inflected plurals allowed).

── PART A — YAML front matter title: full new values ──
FR-001 Preserve Event Guest availability responses
FR-002 Project Timed Slots into their Projected Date Columns
FR-003 Delete Availability Responses
FR-004 Require a non-empty Availability Response
FR-005 Render the edited Availability Response as an overlay above other Availability Responses
FR-006 Keep Availability States mutually exclusive
FR-008 Initialize the Ranged Domain Mode Active Slot Range
FR-009 Present Timed Grid states consistently
FR-010 Support Custom Domain Editing
FR-011 Collapse inactive Timed Grid runs
FR-012 Maintain an optional Scheduled Event Time
FR-013 Preserve Timed Slot Instants across Display Timezone changes ⚠de-compound
FR-014 Derive Timed Grid cell states from event domains ("event domains" stays generic-lowercase)
FR-015 Save Active Slots only within the enabled domain
FR-016 Limit scheduling to Active Slots
FR-017 Preserve Timed Slots across daylight saving transitions
FR-018 Restrict Event Settings editing to the Event Owner
FR-021 Initialize a new Event Timezone from the browser
FR-024 Keep the Event Time Format and Display Time Format independent ⚠shared head expanded
FR-026 Define Event Kind values
FR-027 Restrict Dates-Only Event Timezone settings to the Event Owner
FR-029 Initialize the Display Timezone from the browser
FR-030 Persist only Active Timed Slots in browser storage ⚠minimal caps chosen; dropping "Timed" optional follow-up
FR-031 Send registration Magic Links by email
FR-035 UNCHANGED (Create events with optional descriptions) — only its stale INDEX ROW is replaced (Part C4)
FR-039 Show Event Settings to non-owners
FR-041 Cue Availability Editing from Unavailable Timed Grid areas
FR-042 Keep Timed Grid highlights within their cells
FR-045 span-based only: `timed grid`→`Timed Grid` inside FM/H1/index (rest of line mirrors current README caps)
FR-048 span-based only: `Show All Hours`→`Show all hours` in FM/H1/index ⚠glossary case beats Title Case
FR-050 Generate Active Slots for newly Picked Dates
FR-051 Keep newly Picked Custom Domain Mode dates inactive ⚠caps-only minimal diff
FR-052 Remove Active Slots for removed Picked Dates
FR-053 Preserve Active Slots when converting to Custom Domain Mode
FR-054 Regenerate Active Slots when converting to Ranged Domain Mode
FR-055 Persist changed Timed Event timezones ⚠de-compound
FR-057 Apply Timed Event timezone transition rules ⚠generic "timezone transition rules" stays lowercase
FR-064 Keep Protected Availability Responses in Schedule Overlap ⚠dual reading accepted
FR-065 Show Availability Responses for Availability Editing
FR-066 Filter the Availability Editor grid by Availability Response selection
FR-067 Create Custom-Domain Timed Events PROVISIONAL — read file body; prefer "Create Custom Domain Mode Timed Events" if Mode is meant ⚠R7
FR-070 Omit advertising when Freemium is disabled
FR-071 Remove advertising layout reservations when Freemium is disabled
FR-072 Bypass Freemium restrictions when Freemium is disabled
FR-074 Derive the Timed Event enabled domain ⚠de-compound
FR-075 Generate initial ranged Timed Event Active Slots ⚠de-compound
FR-083 Transfer anonymous PostgreSQL Event Owner authority to another browser ⚠de-compound
FR-084 Preserve Blind Availability privacy
QR-006 Respond promptly for Timed Events
QR-007 Respond promptly for Dates-Only Events ⚠rename
QR-011 Authenticate Cross-Device Access Transfers

H1 rule: apply same span swaps inside H1; H1 stays all-words Title Case incl. articles/To/With EXCEPT verbatim "Show all hours"; FR-004's lone lowercase "a" left as-is. Index cells := FM value verbatim (see Part C4).

── PART B — BODY span replacements (exact old→new) ──
FR-018:14 `event-settings editing form`→`Event Settings editing form`
FR-018:17 `event-settings edits`→`Event Settings edits`
FR-020:15 `when no timed slot is selected.`→`when no Timed Slot is selected.`
FR-030:12 `canonical timed-event state`→`canonical Timed Event state`
FR-049:14 `those timed-event controls`→`those Timed Event controls`
FR-059:19 `Availability-response names`→`Availability Response names`
FR-061:14 `edit an open response.`→`edit an Open Response.`
FR-062:17 `authorize a protected response created by`→`authorize a Protected Response created by`
FR-062:18 `MongoDB protected-response recovery`→`MongoDB Protected Response recovery`
FR-063:15 `permit event-settings management`→`permit Event Settings management`
FR-065:13 `the event visitor cannot edit`→`the Event Visitor cannot edit`
FR-065:14 `availability responses the Event Visitor cannot edit` → cap `Availability Responses`; also `event visitor`(if another lowercase instance remains on :14)→`Event Visitor`
FR-065:19–20 sentence currently wraps after `orders availability` / continuing `responses owned by …`: merge onto ONE physical line ending at its period; inside it `availability-response subset`→`Availability Response subset`, `orders availability responses`→`orders Availability Responses` ⚠also fixes sentences-per-line violation
FR-066:19 `unselected availability responses`→`unselected Availability Responses`
FR-066:21 `every availability response's`→`every Availability Response's`
FR-072:12 `freemium paywall`→`Freemium paywall`; `freemium-gated action`→`Freemium-gated action`
FR-073:18 `that event visitor's`→`that Event Visitor's`
FR-073:21/22 DELETE stray duplicate lowercase token so chain reads `the future [Event Owner](…)'s [Event Visitor Identity](…)` ⚠copy artifact, editorial removal approved
FR-076:12 `mobile timed event page`→`mobile Timed Event page`
FR-076:13 `availability states in the grid`→`Availability States in the grid`
FR-081:23 `anonymous event-owner authority`→`anonymous Event Owner authority`
FR-083:15 `source event-owner powers`→`source Event Owner powers`; `including event settings`→`including Event Settings`
FR-084:15 `only the availability responses the guest`→`only the Availability Responses the guest`
FR-084:16 `availability editing`→`Availability Editing`; `schedule-overlap views`→`Schedule Overlap views`
FR-084:17 `view every availability response.`→`view every Availability Response.`
QR-002:21 `or an availability response without`→`or an Availability Response without`
QR-003:34 `anonymous protected-response mutation`→`anonymous Protected Response mutation`
QR-006:17 `view a timed event`→`view a Timed Event`
QR-006:25 `representative timed-event workload`→`representative Timed Event workload`; `distributed across timed events`→`distributed across Timed Events`
QR-006:29 `Timed-event view and availability-save operations.`→`Timed Event view and availability-save operations.` (availability-save hyphen is generic — keep)
QR-007:17 `view a days-only event`→`view a Dates-Only Event`
QR-007:25 `days-only-event workload`→`Dates-Only Event workload`; `across days-only events`→`across Dates-Only Events`
QR-007:29 `Days-only-event view`→`Dates-Only Event view`
QR-008:25 `Timed and days-only event workloads`→`Timed Event and Dates-Only Event workloads`
QR-011:17 `An event guest or signed-in platform visitor`→`An Event Guest or signed-in Platform Visitor`
QR-011:29 `Cross-device session issuance`→`Cross-Device session issuance` ⚠partial-term; do not invent Access Transfer linkage
Guides: requirements/AGENTS.md:16 `, event kind, mode`→`, Event Kind, mode`
functional/README.md:15 `Name the event kind`→`Name the Event Kind`
quality/README.md:42 `deployment, event kind, permission state`→capped variant
quality/README.md:53 `Name the event kind, operation`→capped variant

── PART C — LINK ADDITIONS (first occurrence per paragraph/list-item/cell ONLY) ──
C1 records (slug):
FR-016:14 enabled-inactive-slot · FR-020:15 timed-slot · FR-059:19 availability-response · FR-061:14 open-response · FR-062:14 event-visitor-identity · FR-062:17 event-visitor-identity AND protected-response · FR-062:18 protected-response · FR-063:15 event-settings · FR-065:19 availability-response · FR-076:12 timed-event · FR-079:14 platform-identity AND event-visitor-identity (plural inflection link acceptable) · FR-081:19 event-visitor-identity · QR-002:21 availability-response · QR-003:34 protected-response · QR-006:17/:25/:29 timed-event · QR-007:17/:25/:29 dates-only-event (post-rename spans) · QR-008:25 dates-only-event AND timed-event
C2 guides (../terminology path): AGENTS.md:16 · functional/README.md:15 · quality/README.md:42 · quality/README.md:53 → all glossary#event-kind, link first occurrence in each list item block.
C3 index cells (terminology/ path): add first-occurrence links inside title cells for rows: 113 availability-response,event-guest · 114 timed-slot,projected-date-column · 115–118 availability-response/-state per audit · 120 ranged-domain-mode,active-slot-range · 121,123,126,154 timed-grid · 122 custom-domain-editing · 124 scheduled-event-time · 125 timed-slot,instant,display-timezone · 127,128 active-slots · 129 timed-slot · 130 event-settings,event-owner · 132 timed-slot · 133 event-timezone · 136 event-time-format,display-time-format · 137 event-kind · 139 dates-only-event,event-owner(+event-timezone if present) · 141 display-timezone · 142 timed-slot · 143,144 magic-link · 148 display-timezone,event-timezone · 149 event-timezone · 151 event-settings · 153 availability-editing,timed-grid · 155 timed-slot · 156 protected-response · 157 event-guest,timed-grid · 158 display-time-format · 159 display-timezone · 160 show-all-hours · 162 active-slots,picked-dates · 163 custom-domain-mode · 164 active-slots,picked-dates · 165 active-slots,custom-domain-mode · 166 active-slots,ranged-domain-mode · 167 timed-event · 168 active-slot-settings · 169 timed-event · 170(née FR-059 row) availability-response · 171? dedupe · 172 protected-response · 173 event-guest,availability-response · 174 event-guest · 175 event-owner · 176 availability-response,schedule-overlap · 177 availability-response,availability-editing · 178 availability-editor,availability-response · 179 timed-event (+custom-domain-mode pending FR-067 ruling) · 180 active-slots,active-slot-range · 181 schedule-overlap,availability-state · 182–184 freemium (first freemium token per cell only; 184 has two) · 185 event-visitor-identity · 186 timed-event · 187 timed-event,active-slots · 188 timed-grid · 191 event-visitor-identity,platform-identity · 196 blind-availability · 209 timed-event · 210 dates-only-event (post-rename) · 214 cross-device-access-transfer
C4 README.md:147 FR-035 row: replace stale title text with `Create events with optional descriptions` (no links — no controlled terms).

── PART D — explicit DO-NOT-TOUCH list ──
migration/**; requirements/README.md:37–53 fenced example; line 84 enum literal `superseded by FR-005`; backticked labels (`Dates and times`, `Dates only`, `Display time zone`, `Time range`) incl. FR-026 body; bare stems `edit token(s)`/`edit credential`/`anonymous edit credentials`; quality/README ISO characteristic tables; FR-011 verb phrase `editing availability`; FR-058:12 `a dates-only value` (semantic check separate); generic `timezone(s)` outside Part A/B spans; FR-004 H1 article `a`; all ADR/SPEC/code files.

Execution integrity rules: apply every replacement as an exact-string substitution verified against expected file:line BEFORE writing; abort-and-reconcile manually on any count mismatch; never rewrite inside code spans, fences, link URLs, or foreign link labels; keep each edited sentence on one physical line; keep each README table row physically single-line.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: opencode
created: 2026-08-26 17:31
---
Executed 2026-08-26. Applied via assertion-checked codemods (/tmp/opencode/t82): every replacement required an exact single-count match before write; 70 files changed. Reconciliations vs the notes: (1) FR-067 resolved to "Create Custom Domain Mode Timed Events" per provisional ruling R7 (body selects Custom Domain Mode). (2) FR-068/FR-069 titles were missing from Part A despite matching audit findings; normalized to satisfy AC#1/#3. (3) Index/guide glossary links use ../terminology/glossary.md (notes said none-prefix); repo evidence README.md:105 and handoff-link depth audit confirm ../ is correct for docs/requirements/*. (4) quality/qr/TEMPLATE.md falls under the declared qr/*.md scope glob but was outside the audits; its "event kind" phrase normalized+linked accordingly (first-occurrence list-item link). Verification: FM≡H1≡index case-insensitively for all 95 records with the two documented H1 exceptions; all 463 glossary links resolve against real glossary H3 anchors at correct depths; days-only=0; freemium drift=0; prose hyphen compounds=0; format:markdown + format:check + markdown lint exit clean; migration/**, fenced README example, status enum line, code-span UI labels untouched.
---

author: opencode
created: 2026-08-26 17:31
---
Tooling note for future agents: /tmp/opencode/t82/inspect.py accidentally shadowed Python's stdlib inspect module when sibling scripts imported it lazily, producing garbage output; removed. Keep stdlib-shadowing filenames out of script dirs.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Normalized the requirements corpus to canonical glossary forms across 70 files (all in docs/requirements/). Casing: 51 YAML titles canonicalized (50 from Part A plus reconciled FR-068/FR-069), 10 H1s rewritten, FR-048 H1 swap to glossary-cased "Show all hours", ~35 Part-B body span fixes; FR-067 titled "Create Custom Domain Mode Timed Events" per R7. Title triples: FM ≡ H1 ≡ README index cell verified for all 95 records (FR-035's stale row replaced with its real title; the two documented H1 exceptions retained: FR-004 lowercase "a", FR-048 "Show all hours"). Links: all enumerated first-occurrence additions applied — records (C1), guides Event Kind (C2), 70 index title cells (C3), plus TEMPLATE.md under the declared qr/*.md glob; checker validates 463 links against 54 real glossary H3 anchors with correct relative depths (../../../ for fr|qr, ../ elsewhere — notes' "none-prefix" corrected per repo evidence at README.md:105). Grep gates: days-only=0, lowercase freemium=0, prose hyphen compounds=0, unexplained casing findings=0. Format gates: npm run format:markdown applied (table realignment), format:markdown:check clean, markdown ESLint exit 0; sentences-per-line preserved incl. FR-065 wrapped bullet merged to one physical line. Out-of-scope confirmed untouched: migration/**, fenced example README:37–53, status enum line 84, code-span UI labels, ISO tables, runtime code. Reconciliations logged in task comments. Documentation-only change: unit/e2e exempt per workflow DoD.
<!-- SECTION:FINAL_SUMMARY:END -->
