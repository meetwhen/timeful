---
id: TASK-0128
title: >-
  Fix NewEvent create-dialog e2e flake: import ExpandableSection and make the
  specific-times switch toggle deterministic
status: Done
assignee: []
created_date: '2026-09-01 18:38'
updated_date: '2026-09-01 18:38'
labels:
  - frontend
  - e2e
dependencies: []
references:
  - frontend/src/components/NewEvent.vue
  - frontend/src/components/NewSignUp.vue
  - frontend/e2e/helpers/timed-event-helpers.ts
  - commit 9f1e0613 (introduced the setSpecificTimesEnabled force-click loop)
documentation:
  - frontend/src/components/NewEvent.test.ts
  - frontend/e2e/helpers/timed-event-helpers.ts
  - frontend/src/components/NewSignUp.vue
modified_files:
  - frontend/src/components/NewEvent.vue
  - frontend/src/components/NewEvent.test.ts
  - frontend/e2e/helpers/timed-event-helpers.ts
type: bug
ordinal: 136300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The firefox-touch e2e suite intermittently failed 1-5 dialog-dependent tests per run: createSpecificTimesEventFromDialog hung for the whole test timeout, the specific-times switch ignored clicks, or the create dialog closed mid-flow. Root causes (both verified during TASK-0127 finalization; the user approved splitting this work into its own task):

1. App bug: NewEvent.vue rendered <ExpandableSection> (the Email reminders section, lines 258-299) without importing it. Every dialog open emitted '[Vue warn]: Failed to resolve component: ExpandableSection', and when signed in the section rendered as an unknown element. NewSignUp.vue imports the same component correctly and was the reference fix.

2. Test helper bug: timed-event-helpers.ts setSpecificTimesEnabled built selectionControl as toggle.locator('xpath=ancestor::*[contains(@class,"v-selection-control")]'), but the v-switch renders INSIDE the [data-testid="specific-times-toggle"] div, so no ancestor carries v-selection-control and that locator never resolved; any click on it burned the entire test timeout (passing runs only succeeded because an earlier action toggled first). The force-click fallbacks also raced the fullscreen dialog entrance animation: force clicks skip the stability wait, so click points computed mid-animation landed on stale geometry (e.g. the .compact-switch__label span at the toggle row's center, which has no click handler), producing 'Clicking the checkbox did not change its state' and swallowed clicks. Debug instrumentation captured elementFromPoint hitting the label span and the fullscreen (375x900, z-index 2400) v-dialog overlay.

Confirmed decisions for this work:
- The fix is app-side only for the missing import (one import line); all other changes are test-side. No transport, timezone, or Temporal boundary changes.
- Force clicks remain only as last-resort fallbacks after a stability-waiting non-force click on the selection control.
- A source-level unit regression test pins that every PascalCase component rendered in NewEvent.vue's template has a .vue import.
- Out of scope: other force-click call sites (create-event button, date cells), unrelated helpers, and the mobile sticky Responses panel work that stays in TASK-0127.

Split out of TASK-0127 finalization on 2026-09-01 with user approval; this task records already-implemented and already-verified work.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 NewEvent.vue imports the ExpandableSection component it renders and no '[Vue warn]: Failed to resolve component' is emitted when the create dialog opens; a unit regression test asserts every PascalCase component rendered in the template has a .vue import.
- [x] #2 The e2e helper setSpecificTimesEnabled targets the switch's .v-selection-control with a deterministic descendant selector and leads with a non-force, stability-waiting click; force clicks are last-resort fallbacks only.
- [x] #3 A repeated determinism loop (dialog open, then helper toggle on-off-on immediately after open) passes multiple consecutive iterations without hangs or swallowed clicks.
- [x] #4 The full firefox-touch e2e suite passes twice consecutively, including the four previously-flaky dialog tests.
- [x] #5 Required frontend checks pass: npm run lint, npm run typecheck, npm run build, npm run test:unit.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Fixed the NewEvent create-dialog e2e flake (split from TASK-0127 finalization with user approval). App bug: NewEvent.vue rendered <ExpandableSection> without importing it; added the missing import (matching NewSignUp.vue), which removes the per-open '[Vue warn]: Failed to resolve component: ExpandableSection' and the broken signed-in rendering, and NewEvent.test.ts gained 'imports every component it renders in its template', which scans the SFC template block for PascalCase tags and asserts each has a .vue import (regex handles default-plus-named imports; the test failed before the import fix). Test helper: setSpecificTimesEnabled in frontend/e2e/helpers/timed-event-helpers.ts previously built selectionControl via a never-resolving ancestor-axis XPath (the v-switch is a descendant of the testid div, so no ancestor carries v-selection-control) and led with force clicks that raced the fullscreen dialog entrance animation; it now targets .v-selection-control as a descendant and leads with a non-force stability-waiting click, keeping force check/uncheck and toggle.click only as fallbacks. Validation: a temporary debug loop spec (6 iterations of dialog open + helper toggle on-off-on immediately after open) passed 6/6 in 43.4s and was then deleted; npm run lint, typecheck, build, test:unit (978/978) all green; npm run test:e2e -- --project=firefox-touch passed 6/6 twice consecutively (58.7s and 1.0m), including all four previously-flaky dialog tests.
<!-- SECTION:FINAL_SUMMARY:END -->
