---
id: TASK-0140
title: >-
  Drop redundant "Event name is required" text under the new-event/edit-event
  name field labeled "Event name (required)"
status: Done
assignee:
  - '@opencode'
created_date: '2026-09-02 12:48'
updated_date: '2026-09-02 13:29'
labels:
  - frontend
dependencies: []
references:
  - frontend/src/components/NewEvent.vue
  - frontend/src/composables/event/useEventEditorState.ts
modified_files:
  - frontend/src/components/NewEvent.vue
  - frontend/src/components/NewEvent.test.ts
priority: medium
type: enhancement
ordinal: 153300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Follow-up to TASK-0136: the new-event/edit-event form's name field (frontend/src/components/NewEvent.vue) now has a persistent label "Event name (required)", so the validation error text "Event name is required" that appears under the field when it is empty is redundant.

Outcome: submitting or blurring an empty event name in the new-event/edit-event form must still be blocked and visually flagged (the existing invalid-name red outline cue), but without displaying the duplicate "Event name is required" message under the field.

Constraints:
- The shared nameRules in frontend/src/composables/event/useEventEditorState.ts are also consumed by NewSignUp.vue, whose name field has no "(required)" label and still needs its message; do not regress the sign-up form.
- Keep the FR-119 100-character cap and its rule behavior from TASK-0136 intact.
- Scope is the new-event/edit-event form only.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 When the new-event/edit-event name field fails required validation, no redundant "Event name is required" text error is rendered below the field, while the existing invalid-name red outline cue still appears.
- [x] #2 Non-redundant name-field validation feedback in the new-event/edit-event form (for example the 100-character rule from FR-119) still surfaces its message where applicable.
- [x] #3 The sign-up form name field in NewSignUp.vue is unaffected and still shows its required-name validation message, since its field has no "(required)" label.
- [x] #4 Unit tests that assert the old "Event name is required" rule behavior in NewEvent.test.ts are updated to the new behavior.
- [x] #5 npm run lint, fmt:check, typecheck, build, and test:unit pass, and affected e2e specs pass.
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
Research findings (verified against installed Vuetify 3 source):
- The invalid-name red outline cue is independent of rules: it is the `new-event-name-field--invalid` class driven by the `showNameFieldError` computed (blur/focus state) plus CSS outline on `.v-field` (NewEvent.vue:1145). Rules only produce the message text.
- Vuetify rule validation: a rule returning `""` still counts as an error (`internalErrorMessages` keeps it, `isValid` becomes false) and VInput renders the details row whenever `messages.length > 0` (vuetify/lib/components/VInput/VInput.js:94-104), so an empty-string rule message would render an empty `.v-messages__message` and a blank details gap. CSS-hiding framework internals would be a hack.
- `submit()` (NewEvent.vue:878) gates on `formRef.validate()`; its only caller is `submitIfAllowed` (NewEvent.vue:1070) which already checks `submitBlocked` (includes `!hasName`). `showSubmitError` ("Please fix form errors before continuing") never shows for empty names because `submitAttempted` is only set when `submitBlocked` is false.
- No e2e helper or spec references the "Event name is required" text; only NewEvent.test.ts:1375 asserts the rule message. The shared `nameRules` (useEventEditorState.ts:229) is also consumed by NewSignUp.vue and must remain untouched.

Plan:
1) NewEvent.vue: stop spreading the shared required rule into `eventNameRules`; make it a plain const containing only the FR-119 length rule. Remove the now-unused `nameRules` from the editorState destructure. NewSignUp.vue keeps consuming the shared rules unchanged.
2) NewEvent.vue: add an explicit empty-name guard at the top of `submit()` (`if (!hasName.value) return`) so the removal of the required rule cannot open a path to submitting an empty name (TASK-0137 API-side validation is still open).
3) NewEvent.test.ts: update the "styles the event name field and caps it at 100 characters" test to the single-rule behavior (empty string valid at rule level, 100 chars valid, 101 chars rejected with the length message) and assert the required message is gone from the field rules.
4) Checks: npm run lint, fmt:check, typecheck, build, test:unit; then affected e2e specs (create + an edit spec) via npm run test:e2e -- --project=firefox-desktop.

Expected behavior after the change: blurring an empty name shows only the red outline cue (no text, no details row, no layout gap); a legacy >100-char name in edit mode still surfaces the length message; empty-name submission remains blocked.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Executed as planned with no scope changes. NewEvent.vue: eventNameRules is now a plain const with only the FR-119 length rule (the shared required rule is no longer spread in), the unused nameRules destructure was removed, and submit() gained an empty-name early-return guard so the removed rule cannot open a path to submitting an empty name. NewSignUp.vue and useEventEditorState.ts are untouched.

First e2e run had both specs fail on page.goto timeouts (Vite cold start, 30s per-test timeout, 10 tests never ran); identical code passed all 12 tests on retry, confirming infra flake rather than a regression.

Checks: lint 0 errors (2 pre-existing one-component-per-file warnings in NewEvent.test.ts), fmt:check pass, typecheck pass, build pass, test:unit 138 files / 986 tests pass, e2e firefox-desktop create + specific-times-edit specs 12/12 pass.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Removed the redundant "Event name is required" text under the new-event/edit-event name field (label already says "Event name (required)") while keeping the empty-name red outline cue and the FR-119 100-character cap.

Changes:
- frontend/src/components/NewEvent.vue: `eventNameRules` no longer spreads the shared `nameRules`, so the required rule (the source of the redundant message) no longer applies to this field; it is now a plain const holding only the FR-119 length rule. Removed the now-unused `nameRules` from the editorState destructure. Added an empty-name early-return guard at the top of `submit()` so removing the rule cannot open a path to submitting an empty name (TASK-0137 API-side length validation is still open). The invalid cue is unchanged: it is the `new-event-name-field--invalid` class (driven by blur/focus state) plus the `.v-field` outline rule, so the red outline still appears without any rules.
- frontend/src/components/NewEvent.test.ts: updated the name-field test to the single-rule behavior: `rules` has length 1, empty string and "Planning sync" and 100 chars are valid, 101 chars returns "Event name must be 100 characters or fewer".
- NewSignUp.vue and the shared `nameRules` in useEventEditorState.ts are untouched, so the sign-up form keeps its required-name message (verified by the untouched git diff and passing sign-up unit tests).

Design note: an empty-string rule message was rejected because Vuetify still counts it as an error and renders the details row (empty message div, blank gap); CSS-hiding framework internals would have been a hack. Dropping the rule is the clean fix since the cue does not depend on rules.

Validation:
- npm run lint: 0 errors (2 pre-existing one-component-per-file warnings in NewEvent.test.ts).
- npm run fmt:check, typecheck, build: pass.
- npm run test:unit: 138 files / 986 tests pass, including the updated name-field test and the semantic-tokens style test that still asserts the `.new-event-name-field--invalid` outline cue.
- e2e firefox-desktop: timed-event-create-firefox.spec.ts (6 tests) and timed-event-specific-times-edit-firefox.spec.ts (6 tests) all pass, 12/12 in 2.0m on the isolated test stack. A first run failed on page.goto timeouts during Vite cold start (infra flake, 10 tests never ran); the retry with identical code passed everything.
- No Markdown files were changed, so the format:markdown DoD item is satisfied vacuously.
<!-- SECTION:FINAL_SUMMARY:END -->
