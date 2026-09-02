---
id: TASK-0131
title: Add top overflow fade to event editor form scroll area
status: Done
assignee:
  - opencode
created_date: '2026-09-02 08:59'
updated_date: '2026-09-02 09:17'
labels:
  - frontend
  - ui
dependencies: []
modified_files:
  - frontend/src/components/NewEvent.vue
  - frontend/src/components/NewEvent.test.ts
priority: medium
type: enhancement
ordinal: 144300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The event editor form (frontend/src/components/NewEvent.vue, shared by new-event and edit-event, and identical for mobile and desktop since both render the same component inside NewDialog) has a scrollable settings list in its v-card-text with a bottom overflow fade (OverflowGradient). When the user scrolls down, the top of the list is cut off with no visual hint of more content above. Add a matching top fade so the scroll affordance is symmetric.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The event editor form's scrollable settings area shows a top fade when the list is scrolled down and hides it when the list is scrolled back to the top
- [x] #2 The top fade visually matches the existing bottom fade and works in both mobile and desktop variants without variant-specific code
- [x] #3 Unit test coverage documents the top fade wiring in the event editor form
- [x] #4 npm run lint, npm run typecheck, npm run build, and npm run test:unit pass
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. NewEvent.vue: wrap the scrollable `v-card-text` in a `tw-relative tw-flex tw-min-h-0 tw-flex-1 tw-flex-col` div so the gradient can anchor to the scroll area itself instead of doing header-height math against the card (header height varies with the dialog subtitle).
2. Inside the wrapper, add a second `<OverflowGradient position="top" :scroll-container="cardTextElement" :show-arrow="false">` guarded by the same `hasMounted && cardTextElement` condition; reuse OverflowGradient's existing top-position support (same pattern as RespondentsList.vue lines 259-271).
3. Keep the existing bottom gradient untouched (`tw-bottom-[90px]` relative to the card).
4. NewEvent.test.ts: add a unit test asserting the top gradient wiring using the established source-assertion pattern.
5. Verify: npm run lint, typecheck, build, test:unit in frontend/.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
OverflowGradient already had position="top" support, so no shared-component change was needed.

Anchored the gradient to a new wrapper div around v-card-text instead of offsetting from the card top: EditorDialogHeader height varies (subtitle only when dialog && showHelp), so a fixed tw-top-[...] offset would be fragile. Wrapper is tw-min-h-0 so the inner overflow-auto region keeps scrolling; v-card-text keeps tw-flex-1 inside it.

Passed :show-arrow="false" on the top gradient because OverflowGradient's arrow hardcodes chevron-down/scrollToBottom, which is meaningless for a top fade. Matches RespondentsList.vue.

Findings: both NewEvent.vue and NewEvent.test.ts fail prettier --check in their committed state, so whole-file prettier --write was reverted; changes were applied surgically (scripted 2-space re-indent of the card-text block) to keep the diff minimal. Repo does not enforce prettier on Vue sources.

Graphify graph updated after the code change.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Add a top overflow fade to the event editor form's scrollable settings list, mirroring the existing bottom fade.

Changes (frontend only):
- NewEvent.vue: wrapped the scrollable `v-card-text` in a `tw-relative tw-flex tw-min-h-0 tw-flex-1 tw-flex-col` div so overflow gradients anchor to the scroll area instead of the card (avoids header-height math; header height varies with the dialog subtitle). Added a second `<OverflowGradient position="top" :scroll-container="cardTextElement" :show-arrow="false">` inside the wrapper with the same `hasMounted && cardTextElement` guard. The bottom gradient is untouched. No mobile/desktop branching: both variants render the same component, so the fade works in the fullscreen phone dialog and the desktop dialog identically.
- NewEvent.test.ts: added "wires top and bottom overflow gradients to the scrollable form area", asserting both gradients share the `cardTextElement` scroll container, the top one uses `position="top"` with `:show-arrow="false"`, and the bottom one keeps its default arrow.

Rationale: `OverflowGradient` already supported `position="top"` (shows when `scrollTop > 1`, hides at top), so this reuses the proven component, same pattern as RespondentsList.vue.

Checks: npm run lint, npm run typecheck, npm run build, npm run test:unit (138 files, 984 tests) all pass. E2E/browser verification intentionally skipped per user instruction; revisit if the fade needs visual confirmation on device.

Note: the event editor's top fade renders without a scroll-to-top arrow because OverflowGradient's arrow only supports scroll-to-bottom; extending it would be a small follow-up if wanted.
<!-- SECTION:FINAL_SUMMARY:END -->
