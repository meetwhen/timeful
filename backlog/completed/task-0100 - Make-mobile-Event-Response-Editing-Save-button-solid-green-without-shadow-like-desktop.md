---
id: TASK-0100
title: >-
  Make mobile Event Response Editing Save button solid green without shadow like
  desktop
status: Done
assignee: []
created_date: '2026-08-29 10:25'
updated_date: '2026-08-29 10:25'
labels: []
dependencies: []
priority: medium
type: enhancement
ordinal: 106000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On mobile, on the Event Response Editing Page, the Save button in the bottom event action bar was rendered as a white elevated button with green text (timeful-elevated-button tw-bg-white tw-text-green), which is visually inconsistent with the desktop editing Save button that is a solid green button with white text and no shadow.

Required outcome:

- The mobile editing Save button in the bottom event action bar uses the same coloring as the desktop editing Save button: solid green background with white text.
- The mobile editing Save button has no shadow, matching the desktop Save button.
- Do not change the mobile editing Cancel button, the desktop Save/Cancel buttons, or the mobile primary availability button styling.

Scope is limited to frontend/src/views/Event.vue (mobile editing Save button classes and its shadow rule) with unit test coverage in frontend/src/views/Event.test.ts.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On mobile, when editing availability responses on the Event page, the Save button in the bottom event action bar uses a solid green background with white text (tw-bg-green tw-text-white), matching the desktop editing Save button coloring.
- [x] #2 The mobile editing Save button has no shadow and no elevated-button border; it no longer uses the timeful-elevated-button class.
- [x] #3 The desktop editing Save button and the mobile primary availability button styling are unchanged.
- [x] #4 Unit test coverage asserts the mobile editing Save button renders with tw-bg-green and tw-text-white and without timeful-elevated-button.
- [x] #5 npm run lint, npm run typecheck, npm run build, and npm run test:unit pass in frontend/.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: agent
created: 2026-08-29 10:25
---
Objective evidence for finalization: AC1-AC3 verified in frontend/src/views/Event.vue (mobile Save button now uses tw-bg-green tw-text-white, timeful-elevated-button removed, and .mobile-editing-save-button shares the desktop shadow-none rule at Event.vue:2581; desktop Save/Cancel and the mobile primary availability button were untouched). AC4 verified by the extended unit test in frontend/src/views/Event.test.ts. AC5 verified: npm run lint, npm run typecheck, npm run build, and npm run test:unit all pass (135 test files, 935 tests). e2e was not run: per frontend AGENTS.md the required frontend checks are lint/typecheck/build/test:unit, and the relevant e2e spec (event-mobile-editing-options.spec.ts) only asserts .mobile-editing-save-button visibility, which the class change preserves.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
On the mobile Event Response Editing Page, the Save button in the bottom event action bar now matches the desktop Save button coloring:

- frontend/src/views/Event.vue: replaced the mobile editing Save button classes "mobile-editing-save-button timeful-elevated-button tw-bg-white tw-text-green" with "mobile-editing-save-button tw-bg-green tw-text-white", so the button is solid green with white text and no longer carries the elevated-button shadow or border.
- frontend/src/views/Event.vue: grouped ".mobile-editing-save-button" with ".desktop-editing-save-button" under the shared box-shadow: none !important rule, mirroring the desktop Save button shadow removal.
- frontend/src/views/Event.test.ts: extended the "renders mobile editing actions with outlined cancel and flat save" test to assert the Save button renders with tw-bg-green and tw-text-white and without timeful-elevated-button.

Verification (all in frontend/): npm run lint, npm run typecheck, npm run build, and npm run test:unit all pass (135 test files, 935 tests). The desktop Save/Cancel buttons and the mobile primary availability button styling were left unchanged. No Markdown files were modified, so format:markdown was not needed.
<!-- SECTION:FINAL_SUMMARY:END -->
