---
id: TASK-0101
title: >-
  Make Edit availability button outline #00994c like its solid fill on mobile
  and desktop
status: Done
assignee: []
created_date: '2026-08-29 10:37'
updated_date: '2026-08-29 10:44'
labels: []
dependencies: []
priority: medium
type: bug
ordinal: 107000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On the event page, the primary availability button in its "Edit availability" state renders with a solid #00994c fill but a lighter green outline (currently 1px solid #29bc68) in both viewport layouts. The outline color should be #00994c so it matches the solid fill, on both mobile and desktop.

Scope:
- frontend/src/views/Event.vue non-scoped style block: .desktop-primary-availability-button and .mobile-primary-availability-button--edit border declarations.
- Use the existing --timeful-primary-action-bg semantic token (#00994c) instead of adding a component-local raw hex, consistent with RespondentsList.vue border usage of the same token.
- Do not touch other #29bc68 usages: .timeful-elevated-button variants, .timeful-switch, and --timeful-compact-switch-track-active-border in index.css are separate components and out of scope.
- Desktop base class border also covers the "Add availability" state, which shares the same solid green fill; changing the color (not removing the border) is the minimal faithful change.

Regression coverage: jsdom does not apply SFC style blocks, so follow the existing source-level assertion pattern used by NewEvent.test.ts (read component source, extract style block, assert rules).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The desktop primary availability button (Edit availability state) outline is #00994c, matching its solid fill, expressed via the existing --timeful-primary-action-bg token rather than a new raw palette value
- [x] #2 The mobile primary availability button (Edit availability state) outline is #00994c, matching its solid fill, expressed via the existing --timeful-primary-action-bg token
- [x] #3 Unrelated #29bc68 usages elsewhere (timeful-elevated-button, timeful-switch, --timeful-compact-switch-track-active-border) are unchanged
- [x] #4 Unit regression coverage asserts both the desktop and mobile outline rules so a future change back to a mismatched green fails
- [x] #5 npm run lint, typecheck, build, and test:unit pass in frontend/
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
1. In frontend/src/views/Event.vue non-scoped style block, change .desktop-primary-availability-button border from 1px solid #29bc68 to 1px solid var(--timeful-primary-action-bg).
2. Same for .mobile-primary-availability-button--edit.
3. Add source-level unit test asserting both border rules reference var(--timeful-primary-action-bg) and neither contains #29bc68, following the NewEvent.test.ts source-assertion pattern; place it near existing Event view style coverage or a small focused spec under src/views.
4. Verify unrelated #29bc68 usages in index.css remain untouched.
5. Run lint, typecheck, build, test:unit.
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Changed the primary availability button outline to match its solid #00994c fill in both layouts:

- frontend/src/views/Event.vue: .desktop-primary-availability-button and .mobile-primary-availability-button--edit borders changed from 1px solid #29bc68 to 1px solid var(--timeful-primary-action-bg) (defined as #00994c in src/index.css), following the existing semantic-token pattern used by RespondentsList.vue.
- Regression coverage in frontend/src/views/Event.test.ts ("Event primary availability button outline"): extracts both style rules from the raw component source, asserts each border uses var(--timeful-primary-action-bg), asserts neither contains #29bc68, and asserts the token still maps to #00994c in src/index.css.

Verification: targeted vitest run passed; npm run lint, npm run typecheck, npm run build, and npm run test:unit (135 files, 937 tests) all passed in frontend/. Unrelated #29bc68 usages (timeful-elevated-button variants, timeful-switch, --timeful-compact-switch-track-active-border) remain unchanged. No e2e spec exists for this style rule; browser E2E was not run since the change is covered by unit-level source assertions.
<!-- SECTION:FINAL_SUMMARY:END -->
