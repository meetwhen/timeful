---
id: TASK-0083
title: Promote decided requirement candidates to functional requirements
status: Done
assignee: []
created_date: '2026-08-26 21:32'
updated_date: '2026-08-26 22:10'
labels: []
dependencies: []
type: docs
ordinal: 90000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Promote triaged migration-inventory candidates into proposed functional requirements per decisions recorded in session triage (user-approved product decisions).

Scope:
- Draft 29 atomic FR files (FR-085 through FR-113) under docs/requirements/functional/fr/ from promoted CAND records; conventions per docs/requirements/README.md and docs/requirements/functional/README.md.
- Add a "Grid Pointer" controlled term to docs/terminology/glossary.md and link it in new FR prose per terminology linking rules.
- Register two new candidate records (CAND-215 More-options aggregation, CAND-216 corrected desktop show-all-hours placement) with provenance, then promote them like other candidates.
- Append matching index rows to docs/requirements/README.md FR table.
- Close each promoted candidate record: Disposition records the permanent FR ID; related_requirements front matter updated together.

Constraints:
- New FRs use status: proposed.
- Applicability comes from each candidate record plus explicit user-mandated extensions; do not generalize beyond recorded decisions.
- One Markdown sentence per source line; table rows stay on one line; glossary-link first controlled-term use.
- Do not touch held candidates (needs-product-decision items, CAND-167 QR).
- Documentation-only change; no runtime code modifications; no commit unless separately requested.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 docs/requirements/functional/fr contains FR-085.md through FR-113.md conforming to the requirements README format with status proposed
- [x] #2 Each new FR statement preserves actor, location, event kind, and exclusions decided in triage without generalizing beyond recorded decisions
- [x] #3 Grid Pointer exists as a glossary entry and new FR prose links controlled terms on first use per paragraph
- [x] #4 CAND-215 and CAND-216 exist as schema-conformant candidate records with provenance referencing this decision round
- [x] #5 Promoted candidate records' Disposition sections name their permanent FR IDs and related_requirements lists them
- [x] #6 docs/requirements/README.md FR index contains rows for FR-085 through FR-113 with resolvable relative links
- [x] #7 Repository Markdown lint/format checks pass on all modified files
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
"CONTINUATION PLAN (execution ordered; each step self-contained). 1) Read docs/terminology/glossary.md + docs/terminology/README.md fully. Insert the new controlled term 'Grid Pointer' maintaining alphabetical position; definition direction: 'Grid Pointer - the interactive highlight marking the grid cell currently under the pointer during availability editing or scheduling; never rendered on collapsed-hours strips'; authoritative context = FR-095 (create FR-095 first or backlink after); follow existing entry link style (../../../terminology/glossary.md anchors are built from heading slugs). 2) Create CAND-215 (More-options aggregation rule) and CAND-216 (desktop show-all-hours placement correction) under backlog-fr-inventory-candidates/, copying the exact section structure/front-matter style of the newest neighbors (the CAND-200+ block); verdict: proposed-requirement, requirement_type: FR, confidence: confirmed (owner decisions this session); Source section quotes this session's ruling text verbatim since these derive from conversation, not backlog; Applicability filled from DECISION LOG notes. Index them in backlog-fr-inventory.md under a NEW trailing group heading '## Session Triage Additions' listed after Temporary Quality Requirements (keeps per-group numeric contiguity intact). 3) Author FR files in ascending order FR-085..FR-113 applying DECISION LOG 1/3, 2/3, 3/3 mappings. File template: YAML front matter with keys id/title/type/components/status (status: proposed everywhere; components from mapping; FR-113 is [frontend, backend]); then '# FR-0NN: Title Case Title'; body = short requirement sentences, ONE SENTENCE PER PHYSICAL LINE, imperative 'system shall' style matching FR-084; state actor/location/event-kind/exclusions decided; link first controlled-term occurrence per paragraph/list-item to ../../../terminology/glossary.md#anchor; never link inside headings or backticked UI labels ('Legend', 'Responses', 'More options' button labels stay literal where quoting UI text). Before writing, grep glossary.md headings to collect exact anchors for every term you will use (Timed Grid, Timed Slot, Availability Response, Availability Response names, Availability Editing, Responses? check existence, Collapsed hours? check, Active Slots, Scheduled Event Time, Event Kind, Dates-Only Event, Display Timezone, Slot Increment, Custom Domain Editing, Grid Pointer). Do not invent glossary anchors. 4) For FR-103 resolve the Slot Increment persistence question by reading the event-settings update path (frontend submit + server route/service); set components accordingly and record findings in task comments. Same check for FR-113 backend rejection (does a route validate range length? requirement mandates it, keep backend regardless since decision 20a said enforce server-side). 5) docs/requirements/README.md: append 29 rows to the Functional Requirements table after FR-084 row; columns ID|Title|Components exactly matching existing row format ('[FR-085](functional/fr/FR-085.md)' etc.), one row per line. 6) Close promoted source candidates: for each promoted CAND record edit ONLY its Disposition (prefix 'Promoted to FR-0NN.' retaining original sentence) and its front matter related_requirements array adding the new FR ID (replace [] where empty); special cases: CAND-147 disposition 'Superseded during triage: contextual summary wording moved to FR-101/FR-102' related_requirements [FR-101, FR-102]; CAND-198 disposition 'Corrected by owner triage: desktop placement split out as CAND-216 and promoted to FR-111' related_requirements [FR-111]; CAND-162 disposition 'Folded into FR-104 with CAND-163'; CAND-163 also references FR-104. 28 promoted sources total + 2 new = 29 FRs. 7) DO NOT touch anything outside this scope; held/unpromoted candidates remain byte-identical (list in IMPLEMENTATION NOTES continuation entry). 8) Run npm run format:markdown at repo root over changed files; search package.json for additional markdown lint scripts and run those too; fix reported issues without reformatting unrelated files. 9) Finalize: acceptance criteria checkboxes, finalSummary summarizing outputs + deferred QR-005 contrast alignment follow-up recommendation, set status Done, leave in Done folder (no archival). Never commit unless separately requested."
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
DECISION LOG 1/3 - Candidate-to-FR mapping with owner rulings (triage session 2026-08-26). Source dir: docs/requirements/migration/backlog-fr-inventory-candidates/. --- FR-085 <- CAND-066: during Availability Editing for a timed event having a scheduled event time, scheduled-event cells are not rendered/editable-drawn; actor availability editor; frontend. FR-086 <- CAND-068: Scheduled Event Time renders on the specific-times timed grid in specific-times mode (extends FR-012 which only defines the field); frontend. FR-087 <- CAND-071: on mobile, an open Responses offcanvas prevents interactions with the grid behind it; frontend. FR-088 <- CAND-105: on mobile, opening the Responses offcanvas preserves the selected Timed Slot and its visible tooltip (extends FR-020 scroll case); frontend. FR-089 <- CAND-204: in the new-event form, selecting any form button does not scroll the form to top; exclusion: deliberate navigation/submission; frontend. FR-090 <- CAND-210: every whole-hour horizontal grid line is hour-labeled on the left, including the line at the top of a collapsed-hours strip; exclusions half-hour lines and dates-only grids; frontend. FR-091 <- CAND-211: time-range picker end-of-day labels: 24h mode zero-padded 00:00-23:00 followed by 24:00; 12h mode end option is 12 AM; selecting end-of-day renders as 00:00 in the next date column; relates FR-002/FR-024; frontend.

DECISION LOG 2/3 --- FR-092 <- CAND-056: legend stays visible when a timed event has no Availability Responses and shows only enabled/active states; OWNER EXTENSION: also applies to Dates-Only Events ('dates-only events too'); scope: every viewport and grid mode (answer 1a); frontend. FR-093 <- CAND-108: clicking anywhere outside the active grid cells clears slot selection, including inter-grid gap and the collapsed strip; input-type-unified pointer+touch (answer 2a); frontend. FR-094 <- CAND-111: tapping a disabled timeslot clears the selection immediately with no replacement selection; padding cells included (answer 3 yes=a); frontend. FR-095 <- CAND-187: collapsed-hours strips never receive the interactive cell highlight (now named Grid Pointer); state-based so hover and touch behave alike (answer 4 yes=a); introduces NEW GLOSSARY TERM Grid Pointer usable in the Custom Domain editor; frontend. FR-096 <- CAND-197: on mobile event pages every tooltip remains fully within the visible screen (answer 5a ALL tooltips); frontend. FR-097 <- CAND-125: the If-needed legend item appears in BOTH timed and dates-only response presentations (answer 6a); frontend. FR-098 <- CAND-127: If-needed responses render yellow via the shared status-color token; contrast alignment with QR-005 deliberately DEFERRED (owner: 'current yellow is ok but better align later with the QR-005') - do not hardcode a threshold in the FR; frontend.

DECISION LOG 3/3 --- FR-099 <- CAND-146: hovering a disabled Dates-Only Event date shows Disabled status mirroring the timed-grid Disabled wording and treatment (answer 8a); frontend. FR-100 <- CAND-148: dates-only enabled-cell hover draws a black inset square frame inside the cell (some px thickness, value not mandated) and the frame must stay contained in the cell without shifting date text; containment + no-shift + black-square-inset are the required behaviors (answer 10); frontend. FR-101 <- B-SPLIT part 1, SUPERSEDES the CAND-147 legacy legend label approach: event-page response summaries read unavailable -> 'Everyone is unavailable', green/some availability -> 'Someone is available'. FR-102 <- B-SPLIT part 2: personal availability view (= Availability Editor, confirmed) WITHOUT overlay: green -> 'You're available', red -> 'You're not available'; WITH the edited-response overlay (FR-005 behavior) active: an overlay-colored item 'You're available' PLUS a plain green item 'Someone is available' for responses beneath the overlay (mapping verified correct, answer B/b). FR-103 <- CAND-159: Slot Increment appears under Advanced options of the new-event form and is OWNER-EDITABLE (NOT read-only - owner correction), visible only for Timed Events (answer 11). COMPONENTS PENDING CODE CHECK: confirm whether the chosen increment persists through the existing event-settings update path needing backend acceptance/validation; if yes use components [frontend, backend], else [frontend].

FINAL DECISION LOG ENTRIES --- FR-104 <- CAND-163 with CAND-162 FOLDED IN: the sidebar label reads 'Legend' without colon (no non-event-page legends exist, answer 12a) and its font and font size match the 'Responses' label; frontend. FR-105 <- CAND-171: timed More-options lists 'Show all hours' above 'Hide if needed times', fixed across viewports and states (13a). FR-106 <- CAND-172: dates-only More-options lists 'Start on Monday' above 'Hide if needed days', fixed across viewports (14 bundle->a). FR-107 <- NEW RECORD CAND-215 registered this session: the 'More options' button is used only when more than one hideable toggle exists; otherwise those toggles render directly instead of the button (owner addition). SEPARATE atomic FR, though sibling orders FR-105/FR-106 sit under it. FR-108 <- CAND-192: the timezone control reserves a default fixed width (currently 112px) so its reset button never resizes the control; phrase as reserved-width outcome, implementation may keep the component configurable (15). FR-109 <- CAND-194: the timezone reset icon is a counter-clockwise arrow; accessible name covered by existing accessibility QRs, not this FR (16a). FR-110 <- CAND-196: time-format button sits LEFT of the timezone button regardless of whether the reset button is present (17 yes=a). FR-111 <- NEW RECORD CAND-216: 'Show all hours' sits to the RIGHT of the event description on DESKTOP where other options and action buttons (cancel, save, delete) live; this CORRECTS CAND-198 whose mobile premise was rejected by the owner (18). FR-112 <- CAND-201: between the two actions only 'Edit availability' is filled/primary; 'Add availability' is not filled (19). FR-113 <- CAND-203: a timed range must be non-empty before save; frontend gate AND backend rejection => components [frontend, backend] (20a).

CONTINUATION NOTES - held scope + guardrails + verification. HELD / DO-NOT-TOUCH (must remain byte-identical at task end): all proposed-requirement FR candidates NOT promoted: CAND-017, CAND-018, CAND-067, CAND-079, CAND-080, CAND-083, CAND-086, CAND-087, CAND-089, CAND-091, CAND-094, CAND-095, CAND-100, CAND-101, CAND-102, CAND-104, CAND-107, CAND-109, CAND-110, CAND-112, CAND-114, CAND-115, CAND-118, CAND-121, CAND-122, CAND-123, CAND-128, CAND-129, CAND-131, CAND-132, CAND-134, CAND-137, CAND-150, CAND-160, CAND-161, CAND-164, CAND-165, CAND-173, CAND-176, CAND-185, CAND-189 (41 records; several have open questions needing product decisions later), the QR candidate CAND-167, plus every covered/excluded/needs-decision record. PROMOTED SET (28): 056 066 068 071 105 108 111 125 127 146 147(superseded->FR-101/102) 148 159 162(folded) 163 171 172 187 192 194 196 197 198(corrected->CAND-216/FR-111) 201 203 204 210 211.

GUARDRAILS AND PITFALLS --- (a) Inventory candidate files use mixed heading levels (## in some blocks, ### in others); match the exact style of adjacent records when creating CAND-215/216 or editing dispositions, and do not renumber anything: CAND ids are permanent temporary review identifiers that never become FR ids. (b) The migration README §Candidate Schema shows #### headings but no record uses them - trust existing files over the README example. (c) FR file prose: use 'The system shall ...' statements; keep applicability sentences as separate lines when they are separate sentences; link terms like [Timed Grid](../../../terminology/glossary.md#timed-grid) mirroring FR-084 relative depth (functional/fr/ is three levels below docs/, so anchors are ../../../terminology/glossary.md#...). (d) UI literals such as `Legend`, `Responses`, `More options`, `Show all hours` are backticked labels, never glossary links (note glossary DOES contain 'Show all hours' as a controlled term - if referencing the concept rather than the visible label, link it per its canonical form). (e) README index titles must match each new FR's title field exactly except title-case differences already present in existing rows pattern. (f) Do not create FRs for deferred items (QR-005 contrast alignment is a recommended follow-up for a NEW future QR/refinement decision, not part of this batch).

VERIFICATION AND COMPLETENESS CHECKLIST for the resuming session --- (1) ls docs/requirements/functional/fr | confirm exactly FR-085.md..FR-113.md added (29 files). (2) Every new FR front matter parses: id matches filename, status proposed, components non-empty. (3) grep requirements README table rows count: rows FR-085..FR-113 present, links resolve from docs/requirements/README.md (functional/fr/<file>). (4) grep backlog-fr-inventory-candidates: every promoted record's Disposition contains its FR id; related_requirements updated in same commit-set. (5) CAND-215/CAND-216 exist + indexed under 'Session Triage Additions'. (6) Glossary contains '## Grid Pointer' entry linking authoritative context FR-095. (7) Run npm run format:markdown (repo root) and any markdown lint script discovered via package.json scripts; rerun until clean on changed files only. (8) No changes outside: 29 fr files, glossary.md, migration candidates dir (28 edited + 2 new), backlog-fr-inventory.md, requirements README.md, this task. If git status shows other files touched, revert them before finishing. (9) Record finalSummary including: created range FR-085..FR-113, supersessions (147,198 folds/corrections), folded pair (162+163), registered additions (215/216), deferred follow-up recommendation: consider QR refresh aligning If-needed yellow contrast with QR-005 rules.

SESSION STATUS as of note time: TASK-0083 plan + decision log complete; implementation not yet started beyond this record-keeping (no FR/glossary/inventory files written yet). Resume at CONTINUATION PLAN step 1.

IMPLEMENTATION COMPLETE. Wrote 29 proposed FR files FR-085.md..FR-113.md under docs/requirements/functional/fr/ following DECISION LOG 1/3-3/3 mappings; components frontend everywhere except FR-103 and FR-113 (frontend, backend).

FR-103 component question resolved by code check: backend normalizeTimedEventPayloadFields validates SlotGeneration.TimeIncrementMinutes in both create and event-settings update paths (server/routes/event_timed_slots.go, server/routes/events.go), so persisting the chosen increment does require backend acceptance/validation => [frontend, backend].

Added Grid Pointer entry to docs/terminology/glossary.md (Timed-Grid Rendering Terms section, between Enabled Inactive Slot and Projected Date Column) with authoritative context FR-095; new FR prose links controlled terms on first use per paragraph.

Registered CAND-215 and CAND-216 under migration/backlog-fr-inventory-candidates/ with session-triage provenance (ruling quoted verbatim, no retained backlog source) and indexed both under new trailing heading '## Session Triage Additions' after Temporary Quality Requirements.

Closed all 28 promoted candidate records: Disposition sections name their permanent FR IDs and related_requirements front matter updated together (CAND-147 superseded by FR-101+FR-102; CAND-162 folded into FR-104; CAND-198 corrected via CAND-216 -> FR-111). Held candidates untouched.

Verification: exactly 29 new FR files; front matter ids match filenames, all status proposed, components non-empty; README rows FR-085..FR-113 match each title field verbatim and links resolve; dispositions contain mapped FR ids for all 28 promoted records; glossary anchors referenced from every new file resolve; npm run format:markdown, format:markdown:check, and lint:markdown all clean; git status shows only the intended files changed.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: staged-review
created: 2026-08-26 22:10
---
POST-STAGE REVIEW FINDINGS (2026-08-27) - apply these doc-only fixes to the already-staged work; do not widen scope beyond them.

FINDING 1 (must fix): Per-paragraph controlled-term linking misses against docs/terminology/README.md "linking" rule (first occurrence per paragraph must be linked). Five sentences need links added:
- FR-085.md final line: `Rendering outside [Availability Editing](../../../terminology/glossary.md#availability-editing) is unchanged.`
- FR-086.md second line: `Selecting specific times shall set the [Scheduled Event Time](../../../terminology/glossary.md#scheduled-event-time) directly, without drawing it as availability.`
- FR-101.md second line: `An [Availability Response](../../../terminology/glossary.md#availability-response) summary reporting any availability shall read \`Someone is available\`.`
- FR-103.md second line: `The creator shall be able to edit the [Slot Increment](../../../terminology/glossary.md#slot-increment) value; the control shall not be read-only.`
- FR-103.md third line: `Saving shall persist the chosen [Slot Increment](../../../terminology/glossary.md#slot-increment), and the backend shall validate it when event settings are accepted.`
This is a real gap versus AC #3 ("links controlled terms on first use per paragraph"), which scripted anchor-resolution checks did not catch.

FINDING 2 (must fix): FR-113.md last paragraph uses generic lowercase for a defined concept: change `Clearing the scheduled event time entirely` to `Clearing the [Scheduled Event Time](../../../terminology/glossary.md#scheduled-event-time) entirely` (canonicalization plus per-paragraph link).

FINDING 3 (resolved by owner, NO CHANGE): FR-108.md keeps the `(currently 112 px)` parenthetical. The requirements README prefers observable outcomes over mechanisms, but the owner elected to retain the parenthesized current value for context.

Verified-clean during review (recorded so the resuming session does not redo it): all 29 FR front matter/H1/status/components conform; README rows match title fields verbatim with resolving links; all 30 candidate dispositions name their FR IDs with related_requirements updated together; CAND-215/216 schema-conformant and indexed under Session Triage Additions; glossary anchors referenced from new files resolve; format/lint/markdown-rule suites pass; staged file set is exactly the intended 63; FR-103 backend component confirmed via normalizeTimedEventPayloadFields validation on both write paths (server/routes/events.go:258, server/routes/events.go:481).

After applying Findings 1-2: rerun `npm run format:markdown:check`, `npm run lint:markdown`, `npm run test:markdown-rules`, `npm run test:markdown-format`, then restage only the touched FR files. No commit unless separately requested.
---

author: staged-review
created: 2026-08-26 22:10
---
POST-STAGE REVIEW FINDINGS (2026-08-27) - apply these doc-only fixes to the already-staged work; do not widen scope beyond them.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Promoted all 28 triaged requirement candidates plus 2 newly registered session-triage records into 29 proposed functional requirements FR-085..FR-113 under docs/requirements/functional/fr/.

What was created:
- FR files FR-085.md through FR-113.md, status proposed, one per promoted behavior; components frontend for all except FR-103 and FR-113, which are frontend+backend.
- New controlled term Grid Pointer in docs/terminology/glossary.md with authoritative context FR-095; it is used and glossary-linked from FR-095 prose.
- CAND-215 (More-options aggregation) and CAND-216 (desktop Show-all-hours placement) registered as schema-conformant candidate records with session-triage provenance, indexed under a new trailing group heading "Session Triage Additions" in docs/requirements/migration/backlog-fr-inventory.md.
- 29 index rows appended to the Functional Requirements table in docs/requirements/README.md with resolvable relative links matching each title field.

Candidate closures:
- All 28 promoted records now name their permanent FR ID in Disposition with related_requirements updated together. Special cases: CAND-147 superseded by FR-101 and FR-102; CAND-162 folded into FR-104 alongside CAND-163; CAND-198 corrected via new CAND-216 into FR-111 after the owner rejected its mobile premise.
- Held candidates (41 unpromoted proposed-requirement records, CAND-167 QR, and every covered/excluded/needs-decision record) are untouched.

Evidence:
- Scripted verification: exactly 29 new FR files with parsing front matter (id matches filename, status proposed, non-empty components); README rows FR-085..FR-113 match title fields verbatim; every promoted disposition names its FR id; all referenced glossary anchors resolve.
- npm run format:markdown, format:markdown:check, and lint:markdown pass at repo root; git status shows only the intended files changed.

FR-103 component decision: backend normalizeTimedEventPayloadFields validates SlotGeneration.TimeIncrementMinutes on both create and event-settings update paths, so the chosen increment does persist through backend validation => frontend+backend.

Deferred follow-up recommendation: consider a future QR refresh aligning the If-needed yellow contrast with QR-005 accessibility rules; deliberately not hardcoded in FR-098 per owner deferral.

No commit made, as instructed.
<!-- SECTION:FINAL_SUMMARY:END -->
