---
id: TASK-0136
title: >-
  Style the event name field like the description field with a 100-character cap
  (FR-119)
status: Done
assignee:
  - opencode
created_date: '2026-09-02 11:18'
updated_date: '2026-09-02 11:43'
labels:
  - frontend
dependencies: []
references:
  - docs/requirements/functional/fr/FR-119.md
modified_files:
  - frontend/src/components/NewEvent.vue
  - frontend/src/components/NewEvent.test.ts
  - frontend/e2e/helpers/timed-event-helpers.ts
  - frontend/e2e/helpers/firefox-timed-event-harness.ts
  - frontend/e2e/inspect/src/dom-resolvers.ts
  - frontend/e2e/inspect/src/scenarios/helpers.ts
  - frontend/e2e/repro/repro-date-added-clears-grid-selections.ts
  - frontend/e2e/repro/repro-date-added-visual-gap-and-duplicate.ts
priority: medium
type: enhancement
ordinal: 149300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
In the new-event and edit-event form (frontend/src/components/NewEvent.vue), the event name input currently uses a solo variant with placeholder "Name your event..." and no label, while the Description input uses an outlined variant with a persistent label ("Description (optional)").

Restyle the name input to match the Description input and cap its length per FR-119 (docs/requirements/functional/fr/FR-119.md):

- Label: "Event name (required)". Placeholder: "Name your event ..." (note the space before the ellipsis).
- Cap the name at 100 characters so longer input cannot be entered or submitted.
- Keep the required-name validation and the invalid-name error cue working with the new outlined treatment: the current `.new-event-name-field--invalid` outline rule and the `.new-event-name-field .v-field__outline { display: none }` override assume the solo variant and need revisiting.
- Keep autofocus and enter-to-blur behavior.
- Scope is the new-event/edit-event form only: NewSignUp.vue and the EventItem.vue duplicate dialog intentionally keep their current styling and "Name your event..." placeholder.

The placeholder change affects selectors that target the editor name input: frontend/e2e/helpers/timed-event-helpers.ts, frontend/e2e/helpers/firefox-timed-event-harness.ts, frontend/e2e/inspect/src/dom-resolvers.ts, frontend/e2e/inspect/src/scenarios/helpers.ts, and frontend/e2e/repro/repro-date-added-*.ts. Update those to match; verify which ones target the NewEvent form versus other fields before editing.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The new-event/edit-event name input renders with the same outlined treatment and visible label as the Description input, labeled "Event name (required)" with placeholder "Name your event ...".
- [x] #2 The name input prevents entering more than 100 characters in both create and edit modes, per FR-119.
- [x] #3 The existing required-name validation and invalid-name error cue still work after the restyle.
- [x] #4 E2e and unit selectors that target the new-event/edit-event name input use the new placeholder; selectors for the duplicate dialog (EventItem.vue) and sign-up form (NewSignUp.vue) are unchanged.
- [x] #5 npm run lint, typecheck, build, and test:unit pass, and affected e2e specs pass.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research findings: nameRules in useEventEditorState.ts is shared with NewSignUp.vue (sign-up name field), so the 100-char cap must NOT go into the shared nameRules; it will be added as a NewEvent-local rule. NewEvent.test.ts stubs v-text-field as nullStub, so behavioral assertions need a capturing stub. e2e selector occurrences confirmed to all target the editor form: timed-event-helpers.ts:146 (keep the 'Name your group...' alternative), firefox-timed-event-harness.ts:157,364, inspect/dom-resolvers.ts:200, inspect/scenarios/helpers.ts:73, repro-date-added-clears-grid-selections.ts:28, repro-date-added-visual-gap-and-duplicate.ts:49.

Plan: 1) NewEvent.vue name field: variant solo->outlined, add label 'Event name (required)', placeholder 'Name your event ...', maxlength=100, drop timeful-solo-field class; keep autofocus, enter-to-blur, focus/blur handlers, and --invalid class binding. 2) Add NewEvent-local eventNameRules = shared nameRules + max-100 rule so edit mode cannot save a legacy >100 name (FR-119 submission block). 3) Style block: remove the .v-field__outline display:none override for the name field; keep the --invalid semantic-token outline cue, dropping its border-radius so it follows outlined rounding. 4) Update NewEvent.test.ts (capturing-stub props test for label/placeholder/variant/maxlength/rules; keep and extend invalid-cue style assertions; negative assertion that v-field__outline is no longer hidden). 5) Update the 7 e2e selector occurrences to the new placeholder. 6) Run lint/typecheck/build/test:unit, then affected e2e specs (create + an edit spec) via the isolated test stack.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Restyled the new-event/edit-event name field in NewEvent.vue to match the Description field and enforced the FR-119 100-character cap.

Changes:
- NewEvent.vue template: name field switched from variant="solo" to variant="outlined" with label "Event name (required)", placeholder "Name your event ...", and maxlength="100"; dropped the timeful-solo-field class. Autofocus, enter-to-blur, focus/blur handlers, and the new-event-name-field--invalid class binding are unchanged.
- NewEvent.vue script: added NewEvent-local eventNameRules (shared nameRules + a 100-character rule). The cap is deliberately not in useEventEditorState's shared nameRules because NewSignUp.vue also consumes it, and FR-119 scope is event names only. The rule blocks saving a legacy >100-char name in edit mode; maxlength prevents entry in both modes (Vuetify forwards non-prop attrs like maxlength to the input element via filterInputAttrs, verified in vuetify/lib/util/helpers.js).
- NewEvent.vue style block: removed the solo-specific `.new-event-name-field .v-field__outline { display: none }` override so the outlined variant renders its outline; kept the `.new-event-name-field--invalid` outline cue using the --timeful-error-foreground semantic token, dropping its border-radius so it follows the outlined field's native rounding.
- NewEvent.test.ts: added a capturing VTextField stub and a behavioral test asserting label/placeholder/variant/maxlength props, the rendered maxlength attribute, and rule behavior (required rule intact, 100 chars valid, 101 chars rejected). Extended the semantic-tokens style test with a negative assertion that the v-field__outline hiding rule is gone.
- E2e/inspect selectors updated from "Name your event..." to "Name your event ..." in all six editor-targeting files: timed-event-helpers.ts (kept the "Name your group..." alternative), firefox-timed-event-harness.ts, inspect/src/dom-resolvers.ts, inspect/src/scenarios/helpers.ts, and both repro-date-added-*.ts repro scripts. NewSignUp.vue, EventItem.vue, and EventItem.test.ts verified untouched, keeping their intentional old placeholder.

Validation:
- npm run lint: 0 errors (2 pre-existing one-component-per-file warnings in the test file).
- npm run typecheck: pass. npm run build: pass.
- npm run test:unit: 138 files / 986 tests pass, including the new name-field tests.
- Affected e2e specs pass on the isolated test stack (firefox-desktop): timed-event-create-firefox.spec.ts (6 tests) and timed-event-weekly-firefox.spec.ts (2 tests), 8 passed in 1.8m.
- No Markdown files changed, so format:markdown was not needed.

FR-119 remains "proposed" because its backend half (TASK-0137, API rejection) is still open; the frontend acceptance criteria of FR-119 are now implemented.
<!-- SECTION:FINAL_SUMMARY:END -->
