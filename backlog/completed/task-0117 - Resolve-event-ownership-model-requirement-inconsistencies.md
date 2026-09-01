---
id: TASK-0117
title: Resolve event ownership model requirement inconsistencies
status: Done
assignee:
  - opencode
created_date: '2026-08-30 20:05'
updated_date: '2026-08-30 20:26'
labels:
  - requirements
  - terminology
dependencies: []
type: docs
ordinal: 124300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Requirements review found four contradictions and five gaps in the event ownership model (Event Owner, Event Guest, Event Visitor Identity, Platform Visitor Identity, EVCC/Granted EVCC, Event Owner Edit Token, Access Transfer). Resolve all nine with these product decisions already confirmed by the owner; do not reopen them.

Files involved: docs/requirements/functional/fr/FR-018.md, FR-059.md, FR-061.md, FR-062.md, FR-063.md, FR-079.md, FR-083.md, FR-084.md; docs/requirements/quality/qr/QR-003.md; docs/terminology/glossary.md; docs/requirements/README.md index; new FR files in docs/requirements/functional/fr/.

Confirmed decisions to encode:
1. Owner-issued Granted EVCC authorizes Event Settings edits (FR-083). FR-018 must list it as a third authorizer; base EVCC still never authorizes settings edits. Fix the glossary EVCC entry accordingly.
2. Sign-in association: FR-079's auto-association applies to the current browser's Event Visitor Identity; associating the source Event Visitor Identity when a Granted EVCC is active follows FR-062's explicit confirmation. Current browser's EVI is still associated per FR-079; the Granted EVCC survives until revoked or the target browser's data is cleared.
3. Blind Availability Mode overrides Open Event Response status: a non-owner guest cannot edit another guest's response even when it is Open.
4. QR-003's 'raw Platform Visitor Identity value' prohibition means an unauthenticated identifier supplied as a credential; authenticated PVI sessions authorize Protected Event Response edits per FR-062 and ownership per FR-063. Clarify wording, not semantics.
5. Event Owner Edit Token is PostgreSQL-only; MongoDB event-settings authorization remains legacy and unchanged (mirror the 'MongoDB behavior unchanged' pattern used by FR-062/FR-081/FR-083).
6. Define missing owner powers as new FR records: event archive authority (owner-only), event deletion authority (owner-only), and pending Access Transfer cancellation by the source browser (QR-011 already requires rejecting 'cancelled' transfers with no defining FR). Reword FR-083's power list to reference defined requirements only.
7. FR-059 duplicate-name rule scope: Event Response display names are unique within an event across guests (two guests named 'John Smith' cannot both use that name). Address the FR-080 profile-name default collision case.
8. Ownership binds to exactly one Platform Visitor Identity; proving the Event Owner Edit Token with a different PVI moves ownership and the previous PVI loses authority.
9. Anonymous-owner lockout is accepted current behavior: clearing browser-local data forfeits owner access with no recovery. Document it as accepted; the recovery idea is tracked separately in the Draft task and must stay out of scope here.

Read docs/requirements/AGENTS.md, docs/requirements/README.md, docs/requirements/functional/README.md, and docs/terminology/README.md before editing. Authoring rules: one sentence per physical source line, controlled terms linked on first occurrence per paragraph/list item/table cell, new FRs use the next sequential IDs with status 'proposed', YAML front matter per README format, no ADR citations in normative FR/QR text. Research current implementation only as needed to state verifiable behavior for the new archive/deletion/cancellation FRs.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 FR-018 lists all three Event Settings authorizers: Event Owner Edit Token, associated Platform Visitor Identity, and the owner-issued Granted Event Visitor Control Credential per FR-083, while stating the base EVCC still never authorizes Event Settings edits
- [x] #2 Glossary EVCC entry scopes its 'never authorizes Event Settings edits' statement to the base EVCC, and the Granted EVCC entry states the owner-issued Granted EVCC authorizes Event Settings edits
- [x] #3 FR-079 states the sign-in association applies to the current browser's Event Visitor Identity, and that when a Granted EVCC is active, associating the source Event Visitor Identity follows FR-062's explicit confirmation instead
- [x] #4 FR-062 states the current browser's Event Visitor Identity is still associated per FR-079, and that the Granted EVCC survives the association until the source revokes it or the target browser's data is cleared
- [x] #5 FR-084 states Blind Availability Mode overrides Open Event Response status: a non-owner guest cannot edit another guest's response even when it is Open, and FR-061 carries a matching blind-mode caveat
- [x] #6 QR-003 clarifies that a 'raw Platform Visitor Identity value' means an unauthenticated identifier supplied as a credential, and that authenticated Platform Visitor Identity sessions authorize Protected Event Response edits per FR-062 and event ownership per FR-063
- [x] #7 FR-018 and the glossary Event Owner Edit Token entry declare the token PostgreSQL-only and state MongoDB event-settings authorization remains legacy and unchanged
- [x] #8 New FR records define event archive authority (owner-only), event deletion authority (owner-only), and pending Access Transfer cancellation by the source browser, so QR-011's 'cancelled' state has a defining requirement and FR-083's power list references defined requirements only
- [x] #9 New controlled terms introduced by the new FRs (for example Event Archive) have glossary entries following docs/terminology/README.md rules, and the new FRs appear as rows in the docs/requirements/README.md functional index with the next sequential IDs
- [x] #10 FR-059 states duplicate Event Response display names are not allowed within an event across guests, and its acceptance criteria address the FR-080 profile-name default collision case
- [x] #11 FR-063 and the glossary Event Owner entry state event ownership binds to exactly one Platform Visitor Identity and that proving the Event Owner Edit Token with a different Platform Visitor Identity moves ownership, with the previous identity losing authority
- [x] #12 The requirements record that clearing an anonymous owner's browser-local data forfeits owner access with no recovery, stated as accepted current behavior
- [x] #13 Every changed requirements file follows docs/requirements/AGENTS.md and docs/terminology/README.md conventions: one sentence per physical source line, controlled terms linked on first use per paragraph, status 'proposed' for new FRs, and no ADR citations inside FR/QR normative text
- [x] #14 All relative links in changed Markdown files resolve to existing files and anchors
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
## Implementation plan

New FR IDs continue the sequence: FR-115 (archive), FR-116 (deletion), FR-117 (transfer cancellation). All three are PostgreSQL-first, `status: proposed`, `components: [frontend, backend]`.

1. **FR-018** — add Granted EVCC (owner-issued, per FR-083) as third Event Settings authorizer; state the base EVCC never authorizes settings edits (keep existing sentence, sharpen wording); declare the Event Owner Edit Token and this authorization model PostgreSQL-only with one legacy-MongoDB sentence.
2. **FR-059** — reword duplicate-name rule: display names unique within an event across guests (compared after FR-059 normalization, case-sensitive); add acceptance criterion for the FR-080 profile-default collision (reject creation until a distinct name is chosen).
3. **FR-061** — add blind-mode caveat: Blind Availability Mode (FR-084) overrides Open status for non-owner guests.
4. **FR-062** — add: current browser's EVI is still associated per FR-079; Granted EVCC survives association until revoked or target data cleared.
5. **FR-063** — add: ownership binds to exactly one PVI; token proof with a different PVI moves ownership, previous PVI loses authority; add accepted-lockout paragraph (anonymous owner clearing browser data forfeits owner access, no recovery; signed-in owners recover via PVI association).
6. **FR-079** — clarify auto-association applies to the browser performing sign-in; source-EVI association under an active Granted EVCC follows FR-062's confirmation.
7. **FR-083** — reword power list to reference defined requirements (FR-018 settings, FR-084 visibility, FR-115 archive, FR-116 deletion).
8. **FR-115 (new)** — Event Owner archives/unarchives a PostgreSQL event; authorization mirrors FR-018 (token, associated PVI, owner-issued Granted EVCC; base EVCC never); archived events remain viewable via the Event Link but become read-only (no response or settings mutations); unarchive restores; archived events group under an Archived section of a signed-in owner's event list; authority only for anonymous owners (no UI specified); MongoDB archive behavior unchanged.
9. **FR-116 (new)** — Event Owner deletes a PostgreSQL event; same authorization model; observable outcome: the event link stops resolving for everyone and responses become inaccessible; explicitly does not specify retention, recovery, or export; MongoDB deletion behavior unchanged.
10. **FR-117 (new)** — source browser cancels a pending Access Transfer; cancelled transfers cannot be approved or consumed (satisfies QR-011's rejected 'cancelled' state); cancellation does not revoke an already-issued Granted EVCC.
11. **QR-003** — clarify 'raw Platform Visitor Identity value' = unauthenticated identifier supplied as a credential; authenticated sessions authorize via FR-062/FR-063 associations.
12. **Glossary** — EVCC entry: scope 'never authorizes Event Settings edits' to the base EVCC; Granted EVCC entry: add owner-variant authorizes Event Settings edits (cite FR-018); Event Owner Edit Token entry: PostgreSQL-only + legacy MongoDB sentence; Event Owner entry: single-PVI binding and move-on-token-proof; add new 'Archived Event' entry (authoritative context FR-115) in the Scheduling Domain section.
13. **requirements/README.md** — add index rows for FR-115/116/117 after FR-114.

Verification: sentence-per-line authoring, controlled-term linking rules, link resolution check across changed files, `npm run format:markdown` (frontend). Documentation-only change: unit/e2e tests exempt per DoD.
<!-- SECTION:PLAN:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-30 20:09
---
Owner clarification (session): MongoDB is legacy; write requirements with PostgreSQL in mind. New archive/deletion FRs are PostgreSQL-first requirements, with at most one sentence noting legacy MongoDB behavior stays unchanged. This refines decision 5 and the new-FR scoping in the task description.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Resolved all nine ownership-model findings across 11 changed and 3 new files, encoding the owner-confirmed decisions (PostgreSQL-first; MongoDB mentioned only as unchanged legacy).

Contradictions fixed:
- FR-018 now authorizes Event Settings edits through three credentials (Event Owner Edit Token, associated Platform Visitor Identity, owner-issued Granted EVCC per FR-083) and states the base EVCC never authorizes settings edits.
- FR-079/FR-062: sign-in auto-association is scoped to the browser performing sign-in; source-EVI association under an active Granted EVCC follows FR-062's explicit confirmation, and the Granted EVCC survives association until revocation or target data clearing.
- FR-084 + FR-061: Blind Availability Mode overrides Open Event Response status for non-owner guests.
- QR-003: "raw Platform Visitor Identity value" clarified as an unauthenticated supplied identifier; authenticated sessions authorize via the FR-062/FR-063 associations.

Gaps closed:
- Event Owner Edit Token declared PostgreSQL-only (FR-018 + glossary) with MongoDB owner authorization unchanged.
- New FR-115 (owner-only event archive/unarchive; archived events stay viewable but read-only; Archived grouping for signed-in owners; authority-only for anonymous owners), FR-116 (owner-only event deletion; link stops resolving; retention/recovery/export explicitly unspecified), FR-117 (source-browser cancellation of pending Access Transfers; cancellation does not revoke an earlier Granted EVCC). All mirror FR-018's authorization model; base EVCC excluded.
- FR-083's power list now references FR-018/FR-084/FR-115/FR-116 instead of undefined powers.
- FR-059: response display names unique per event across guests (post-normalization) with an FR-080 profile-default collision criterion.
- FR-063 + glossary Event Owner: ownership binds to exactly one PVI; token proof with a different PVI moves ownership.
- FR-063 records the anonymous-owner lockout (cleared browser data forfeits owner access, no recovery) as accepted current behavior.

Terminology: glossary gained an "Archived Event" entry (FR-115) in Scheduling Domain; EVCC, Granted EVCC, Event Owner Edit Token, and Event Owner entries updated to match. README functional index gained FR-115/116/117 rows.

Verification: scripted link and glossary-anchor check passed for all changed files; `npm run format:markdown` (repo root) run to idempotency. Documentation-only change, so unit/e2e tests exempt per DoD. Note: `format:markdown` lives in the root package.json, not frontend/.
<!-- SECTION:FINAL_SUMMARY:END -->
