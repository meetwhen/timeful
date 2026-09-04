---
id: TASK-0141
title: >-
  Show a single bold red invalid state with message and alert icon for the event
  name field
status: Done
assignee:
  - '@opencode'
created_date: '2026-09-02 13:50'
updated_date: '2026-09-02 14:48'
labels:
  - frontend
dependencies: []
references:
  - frontend/src/components/NewEvent.vue
  - frontend/src/components/NewEvent.test.ts
modified_files:
  - frontend/src/components/NewEvent.vue
  - frontend/src/components/NewEvent.test.ts
priority: medium
type: enhancement
ordinal: 154300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Follow-up to TASK-0140: in the new-event/edit-event form (frontend/src/components/NewEvent.vue), an invalid event name currently shows the default grey Vuetify border plus a separate 1px dark-red CSS outline at the same time, which reads as two competing outlines.

Outcome: the invalid name field shows a single, brighter, bolder red outline (no grey), a red validation message ("Event name is required") below the input, and a red alert icon at the end of the input, matching the reference treatment the user provided (Tailwind docs email error: red border, red message below, filled red exclamation icon inside the field).

Context:
- TASK-0140 removed the shared required rule so no message renders; this task intentionally restores a validation message with proper error styling.
- The invalid-name red cue is currently the `new-event-name-field--invalid` class plus a `.v-field` outline rule; the grey border is the Vuetify outlined variant border.
- The FR-119 100-character rule and its message must keep working (edit mode with legacy >100-char names).
- Scope is the event name field in the new-event/edit-event form only; NewSignUp.vue and other forms keep their current behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 When the new-event/edit-event name field is invalid, only a red outline is shown on the field: the default grey Vuetify border is not visible alongside it.
- [x] #2 The invalid-name red outline is brighter and bolder than the current 1px dark-red outline so it clearly attracts attention.
- [x] #3 When the name field is invalid, a red validation message ("Event name is required") is shown below the input, styled like the reference treatment (red text under the field).
- [x] #4 When the name field is invalid, a red alert icon is shown at the end (inner-right) of the input.
- [x] #5 The invalid styling continues to cover the FR-119 100-character rule for the name field in edit mode, with its message, red outline, and icon.
- [x] #6 The sign-up form name field (NewSignUp.vue) is unaffected.
- [x] #7 npm run lint, fmt:check, typecheck, build, and test:unit pass, and affected e2e specs pass.
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
Research findings (verified against installed Vuetify 3 source and app theme):
- VTextField passes `error: isValid.value === false` to VField (VTextField.js:178), so any rule failure applies both `.v-input--error` (root, from validationClasses) and `.v-field--error` (field element).
- With `.v-field--error`, Vuetify natively renders the outlined border red at full opacity: `.v-field--error:not(.v-field--disabled) .v-field__outline { color: rgb(var(--v-theme-error)) }` and `.v-input.v-input--error .v-field__outline { --v-field-border-opacity: 1 }` (VField.css:372-376). The app theme error color is #DB1616 (src/plugins/vuetify.ts), which is brighter than the current 1px #dc2626 outline (from --timeful-error-foreground). The grey border and the red outline are the same border element recolored, so a rule-driven red border is a single outline.
- Vuetify natively colors append-inner icons red in error state: `.v-field--error:not(.v-field--disabled) .v-field__append-inner > .v-icon { color: rgb(var(--v-theme-error)) }` (VField.css:245-249). There is no built-in error icon prop, so the icon must be provided and shown conditionally.
- Outlined border thickness is driven by `--v-field-border-width` (VField.css:427+ consumes it; focused already gets 2px natively at VField.css:378).
- The grey+red double outline seen today exists because TASK-0140 removed the required rule, so the field never enters Vuetify error state; the grey Vuetify border stays and the custom class-based outline draws red next to it.
- The current class-based cue (`new-event-name-field--invalid` + showNameFieldError/hasBlurredNameField/isNameFieldFocused/handleNameFieldFocus/handleNameFieldBlur) would duplicate the rules-driven state; the length-rule error must get the same treatment, so the rules-driven state should be the single source of truth.
- The app already uses `mdi-alert-circle` icon names (e.g. AlertText.vue:5); `mdi-alert-circle` is a filled red circle with knocked-out exclamation, matching the reference screenshot.

Plan:
1) NewEvent.vue template: on the name field add `append-inner-icon="mdi-alert-circle"`; drop the `new-event-name-field--invalid` class binding and the `@focus`/`@blur` cue handlers.
2) NewEvent.vue script: restore spreading the shared `nameRules` (required rule, message "Event name is required") into `eventNameRules` alongside the FR-119 length rule; re-add `nameRules` to the editorState destructure; remove the now-unused cue machinery (hasBlurredNameField, isNameFieldFocused, showNameFieldError, handleNameFieldFocus, handleNameFieldBlur, and their reset() lines). Keep the TASK-0140 empty-name guard in submit() and enter-to-blur.
3) NewEvent.vue style block: replace the `.new-event-name-field--invalid .v-field` outline rule with:
   - `.new-event-name-field .v-field__append-inner { visibility: hidden }` (reserve space, no layout shift),
   - `.new-event-name-field.v-input--error .v-field__append-inner { visibility: visible }`,
   - `.new-event-name-field.v-input--error .v-field__outline { --v-field-border-width: 2px }` (bolder).
   Red color comes from the Vuetify theme error (#DB1616) natively for border, icon, and message; no new color overrides.
4) NewEvent.test.ts: revert the rules assertions to the two-rule behavior (empty -> "Event name is required", 100 chars valid, 101 chars -> length message); assert `appendInnerIcon` prop is "mdi-alert-circle"; update the semantic-token style test to assert the new visibility/border-width rules and that the old `new-event-name-field--invalid` outline rule is gone.
5) Checks: npm run lint, fmt:check, typecheck, build, test:unit; e2e firefox-desktop create + specific-times-edit specs.

Expected behavior: on invalid name (empty, or >100 chars in edit mode) the field shows a single 2px bright-red border (#DB1616), a red message below ("Event name is required" or the length message), and a red alert icon at the inner end; focused-typing clears the error as soon as the value is valid; NewSignUp.vue untouched.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Execution deviations from the recorded plan: (1) the extra outline was traced to the global App.vue rule .v-input--error .v-field, .v-field--error { outline: red solid; border-radius: 3px; } (a Vuetify 2-era leftover), so a scoped outline: none override was added for the name field instead of relying on the plan assumption that removing the class-based cue suffices. (2) Per user decision during execution, the invalid color is the brand theme error color (#DB1616, no overrides) instead of raw #ff0000; the --v-theme-error override was dropped and no palette literals are added. (3) App-wide removal of the global rule plus a shared invalid-field class (index.css, timeful-solo-field pattern) was deferred to a follow-up task covering NewEvent name, NewSignUp, NewGroup, GuestDialog, SignUpForSlotDialog, EmailInput, and ConfirmDetailsDialog.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
The invalid event name field now shows a single bold red treatment using the brand error color: restored the required rule ("Event name is required") alongside the FR-119 100-character rule so the field enters Vuetify error state natively; added append-inner-icon mdi-alert-circle with visibility reserved when idle and shown on error; neutralized the App.vue global error outline for this field via a scoped outline: none rule so only the native 2px theme-error border shows (label notch intact, no outline crossing the floating label); kept the 2px border-width override for boldness. Border, label, message, and icon all come from the Vuetify theme error color (#DB1616) with no color overrides, per the user decision to use the brand color instead of #ff0000. NewSignUp.vue and other forms are untouched. Checks: lint, fmt:check, typecheck, build, test:unit (986 passed) and firefox-desktop e2e timed-event-create + timed-event-specific-times-edit (12 passed). App-wide unification (remove App.vue global rule, shared invalid-field class) recorded as a follow-up task.
<!-- SECTION:FINAL_SUMMARY:END -->
