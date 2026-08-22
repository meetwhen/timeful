# ADR-011: Frontend Guest Response Ownership Semantics

Date: 2026-05-26

Status:

- Accepted

## Context

Within the frontend migration baseline described in ADR-008, guest response editing has a separate ownership model from signed-in user responses:

- signed-in users are identified by durable account IDs
- guest responses may be legacy name-keyed rows or token-backed rows with opaque guest credentials
- the frontend needs to decide which guest row is "yours" before and after event-specific sign-in
- the backend can authorize guest mutations only from data the client sends back on fetch and mutation paths

That creates a concrete product and architecture decision surface:

- the migrated frontend needs one canonical way to persist and reuse guest ownership state across event loads
- edit affordances must match backend mutation rules instead of treating all guest rows as interchangeable
- response authors need to retain protected-response access when moving to another device

This is not just a one-off respondents-list bug. It is a frontend ownership and UX contract that affects event loading, guest editing, deletion, rename flows, and how the UI explains editability.

## Decision

The frontend keeps one shared guest-response ownership model:

- guest ownership begins as browser-local state
- token-backed guest ownership uses opaque backend-issued credentials: `guestId` plus `guestEditToken`
- legacy guest ownership falls back to the stored guest display name when no token-backed identity exists
- guest-aware fetch and mutation paths reuse the stored guest identity through existing query semantics: `guestId` first, `guestName` second
- a response author can associate their proven browser-local identity with an authenticated account by signing in to manage the event
- frontend edit affordances must mirror protected and open response-access semantics

The sign-in entry point is neutral before authentication. After authentication,
the frontend derives the visitor's event-author and response-author capabilities
from the event's ownership records.

## Rules

### Canonical guest ownership state

- Shared guest ownership state belongs in the schedule-overlap guest-storage boundary, not in individual views or components.
- The canonical current-browser guest identity is the stored opaque `guestId` when present.
- If no stored `guestId` exists, the canonical fallback identity is the stored guest display name.
- Components and composables should consume shared guest ownership helpers instead of rebuilding guest lookup logic locally.
- A response author starts as the identity that created a response in the browser. After that identity is proven and associated through event sign-in, the authenticated account represents the same response author on another device.

### Fetch and mutation semantics

- Guest-aware event and response fetches should send `guestId` when token-backed ownership exists.
- Guest-aware event and response fetches should send `guestName` only as the legacy fallback when no token-backed identity exists.
- Guest mutation requests should preserve the distinction between the caller's guest identity and the target guest row being edited.
- Frontend editability helpers must match backend authorization:
  - legacy guest rows are editable only by the matching stored guest identity
  - protected rows are editable only by their response author through the local identity or authenticated event-sign-in association
  - open rows are editable by any event visitor

### UX semantics

- The UI may describe a row as editable only when the current visitor can actually satisfy the backend ownership rules for that row.
- New responses are protected by default. Their response author may explicitly make them open.
- An unauthenticated visitor may enter the neutral `Sign in to manage this event` flow; the UI then exposes event and response management capabilities that belong to the authenticated visitor.
- The UI may continue to show non-editable guest responses as readable respondents so their availability still participates in overlap and filtering flows.
- Loss of browser-local guest state is an accepted consequence of the current model, not a frontend inconsistency.

### Storage-loss and orphan semantics

- Clearing `localStorage` may remove the current browser's ability to prove authorship of a protected response.
- A response author can restore protected-response access on another device after their original browser proves its local identity and associates it through event sign-in.
- An open response remains editable by any event visitor after browser storage loss.
- The frontend may infer that a row is not editable for the current browser, but it must not claim that the row is globally orphaned.

### Boundary and test discipline

- Keep guest ownership transport fields at explicit boundaries, consistent with ADR-001.
- Do not widen raw guest ownership payload fields into ad hoc component-local contracts when shared helpers already define the semantics.
- Add focused regression coverage when guest editability rules change so respondents lists, selected-guest CTAs, and direct guest-edit flows remain aligned with backend authorization.

## Consequences

- Guest response editing starts without requiring accounts, while event sign-in provides cross-device restoration for proven response authors.
- The migrated frontend gets one explicit contract for browser-local identity, authenticated response authors, protected responses, and open responses.
- Guest responses can remain useful for scheduling even when they are no longer editable from the current browser.
- Product and support expectations should treat browser-storage loss as a loss of immediate editability, unless the response author previously associates the identity through event sign-in.
