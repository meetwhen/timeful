---
id: TASK-0144
title: >-
  Clear Unsaved changes prompt state after Leave page so reopening Edit event
  starts clean
status: Done
assignee:
  - opencode
created_date: '2026-09-02 15:32'
updated_date: '2026-09-02 16:05'
labels: []
dependencies: []
modified_files:
  - frontend/src/components/NewDialog.vue
  - frontend/src/components/NewDialog.test.ts
priority: medium
type: bug
ordinal: 157300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
When a user edits an event, changes something, closes the modal, chooses Leave page on the Unsaved changes prompt, and then opens Edit event again, the Unsaved changes prompt shows immediately even though nothing is dirty anymore. Expected: the prompt must not appear after Leave page was chosen; leaving should fully clear the prompt state so a fresh reopen starts clean. The same modal is shared by the App-level create/edit dialog and the event-page edit dialog. Fix belongs in the frontend dialog flow; regressions must be covered at both unit and e2e layers.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Given an event is edited and a change is made, when the edit-event modal is closed and Leave page is chosen on the Unsaved changes prompt, then reopening the edit-event modal does not show the Unsaved changes prompt
- [x] #2 Leaving with unsaved changes still closes the modal and restores the form to the event's saved data
- [x] #3 A unit regression test covers leave-then-reopen clearing the Unsaved changes prompt state and fails against the unfixed component
- [x] #4 Required frontend checks pass: lint, fmt:check, typecheck, build, test:unit
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
Research (already done in session):

- The Unsaved changes prompt is `UnsavedChangesDialog.vue` rendered inside `NewDialog.vue` (shared by App create/edit dialog and Event.vue edit dialog).
- Close-with-edits sets `unsavedChangesDialog = true` (NewDialog.vue:186). Choosing "Leave page" emits `leave` only; `UnsavedChangesDialog.vue:22` does not clear its own model. `exitDialog()` (NewDialog.vue:189) closes the outer dialog and calls `resetToEventData()` but never resets `unsavedChangesDialog`, so the ref stays `true` and the remounted prompt shows open on the next open of the modal. The save path `handleRefreshEvent` (NewDialog.vue:197) already resets it explicitly; only the Leave path misses it.
- Dirty-form state itself is correctly cleared: `resetToEventData` → `applyEventData` restores all fields including extra captured state, so a subsequent close will not re-trigger the prompt.

Plan:

1. Unit regression test first (expected to fail pre-fix) in `frontend/src/components/NewDialog.test.ts`: extend `UnsavedChangesDialogStub` to render a leave button that emits `leave`; new test: `edit: true` + edited → close → `data-open="true"` → click leave → expect `update:modelValue [[false]]` and `resetToEventData` called → reopen via `setProps({ modelValue: true })` → assert `data-open="false"`.
2. E2E regression spec first (expected to fail pre-fix, no auth) in `frontend/e2e/event-edit-unsaved-changes.spec.ts`: spec name has no `-firefox` suffix so it runs on chromium-desktop and chromium-mobile; force desktop viewport 1440x1400 so `isPhone` is false and outside click works. Anonymous flow via `seedCanonicalTimedEvent` (plain POST /api/events) → `openEventPage` → `openEditDialog` (`#edit-event-btn`; anonymous events are editable by anyone) → dirty the form via `getEditorNameInput().fill(...)` → close via outside click on the scrim (`@click:outside` fires despite `persistent`) → assert "Leave page" visible → click it → editor card hidden → reopen edit dialog → assert "Leave page" hidden and the name input restored to the seeded value.
3. Fix: in `NewDialog.vue` `exitDialog()` add `unsavedChangesDialog.value = false`, consistent with `handleRefreshEvent`. One-line component-state change; no boundary/transport/Temporal changes, no ADR implications.
4. Verify: `cd frontend` — npm run lint, fmt:check, typecheck, build, test:unit; then `npm run test:e2e -- --project=chromium-desktop` and `--project=chromium-mobile` for the new spec (Playwright owns the isolated test stack on 3003/4174).
5. Finalize: verify acceptance criteria with test evidence, record final summary, mark Done; run `graphify update .`.

Executed approach deviation: step 2 (e2e spec) was dropped after user review; step 4's e2e runs are not applicable. Fix and unit coverage delivered as planned.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Scope change (user-approved 2026-09-02): the e2e regression spec was dropped; regression coverage is unit-only. Rationale: the bug is a deterministic component-state issue fully captured by the NewDialog unit test, while the e2e spec turned flaky for reasons unrelated to the fix.

E2E investigation findings (spec deleted, findings kept for future work): the bug was reproduced in a real browser pre-fix — after Leave page and reopening the edit dialog, the Unsaved changes prompt remounted open. Separately, an unrelated automation flake was observed: the first #edit-event-btn force-click after the leave flow was consistently swallowed (0 click events reached the button) even though no overlay was active, no element covered the button center (elementFromPoint returned the button's own label span), and nothing was inert; a second click ~800ms later always opened the dialog. Closed overlays remain in the DOM with display:flex and no active class. If a future e2e needs to reopen the dialog right after closing it, retry the click (toPass loop) instead of one force-click after waiting for .v-overlay--active to reach 0.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Problem

After editing an event, changing something, closing the modal, and choosing "Leave page" on the Unsaved changes prompt, reopening Edit event immediately showed the prompt again even though nothing was dirty.

## Root cause

In `NewDialog.vue`, the Leave path was the only exit that did not reset the prompt state: "Leave page" emits `leave` → `exitDialog()` closed the modal and reset the form but never cleared `unsavedChangesDialog`, and `UnsavedChangesDialog.vue`'s Leave button does not clear its own model. The ref stayed `true`, so when the dialog content remounted on the next open, the prompt rendered open. The save path (`handleRefreshEvent`) already reset it explicitly.

## Change

- `frontend/src/components/NewDialog.vue`: `exitDialog()` now sets `unsavedChangesDialog.value = false` before closing, consistent with `handleRefreshEvent`. One-line component-state fix; no boundary/transport/Temporal changes.
- `frontend/src/components/NewDialog.test.ts`: extended `UnsavedChangesDialogStub` to emit `leave` and added the regression test "clears the unsaved changes dialog after leave so reopening does not show it again" (edit mode + edited → close → prompt open → Leave page → modal closes with `resetToEventData` → reopen → prompt closed).

## Scope decision (user-approved)

The planned e2e regression spec was dropped; coverage is unit-only. The e2e had already reproduced the bug in a real browser pre-fix, but the spec itself was flaky for reasons unrelated to the fix: the first `#edit-event-btn` force-click after the leave flow was swallowed (0 click events reached the button) despite no active overlay, nothing covering the button center, and nothing inert; closed Vuetify overlays linger in the DOM with `display: flex`. Retry-based clicking would be needed; details recorded in implementation notes for future e2e work.

## Verification

- Regression test failed pre-fix (`data-open` was `"true"` after reopen) and passes post-fix.
- `npm run lint` — 0 errors (2 pre-existing warnings in `NewEvent.test.ts`), `npm run fmt:check` — pass, `npm run typecheck` — pass, `npm run build` — pass, `npm run test:unit` — 987 tests in 138 files all passing.
- Repo-root `npm run format:markdown` produced no changes.
<!-- SECTION:FINAL_SUMMARY:END -->
