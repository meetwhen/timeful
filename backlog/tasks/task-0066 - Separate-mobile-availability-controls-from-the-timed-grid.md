---
id: TASK-0066
title: Separate mobile event action bar from the timed grid
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-25 09:07'
updated_date: '2026-08-25 10:31'
labels:
  - mobile
  - schedule-overlap
  - ui
dependencies: []
references:
  - frontend/src/views/Event.vue
  - docs/requirements/functional/fr/FR-076.md
modified_files:
  - docs/requirements/README.md
  - docs/requirements/functional/fr/FR-076.md
  - frontend/e2e/schedule-overlap-mobile-touch-firefox.spec.ts
  - frontend/src/index.css
  - frontend/src/views/Event.vue
  - frontend/src/views/Event.test.ts
priority: medium
type: enhancement
ordinal: 68000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Make the mobile timed-event action bar visually distinct from available grid slots and record the resulting durable frontend behavior as a functional requirement. White is the confirmed bar treatment; green and blue remain semantic to availability and scheduling actions.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The mobile timed-event action bar appears on a white elevated surface that remains visually distinct from the timed grid
- [x] #2 Availability actions remain green and scheduling actions remain blue on the white action bar
- [x] #3 A functional requirement records the scoped mobile action-bar behavior and is listed in the requirements index
- [x] #4 Regression coverage verifies the mobile action-bar presentation contract
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Record FR-076 as an accepted frontend requirement for the mobile timed-event action bar and index it in docs/requirements/README.md.
2. Apply an owned white elevated surface to the mobile action bar in Event.vue for both normal and scheduling states; retain green for availability actions and blue for scheduling actions.
3. Extend Event-view unit coverage to assert the action-bar presentation contract and scheduling action border.
4. Run scoped Event tests and required frontend verification; record the existing Firefox host-library blocker if browser verification remains unavailable.

5. Refine the existing mobile action-bar controls: make normal Schedule blue outlined, make Edit availability green outlined without the elevated green shadow, and use a subdued blue fill for disabled Schedule when no slot is selected. Extend Event-view assertions for each state before rerunning frontend checks.

5. Refine the mobile disabled Schedule action after review: retain its inactive blue fill and border, render the label and icon white at full opacity despite Vuetify disabled defaults, and cover its disabled state in Event-view regression tests.

6. Align disabled mobile Schedule with the supplied desktop reference: use a solid medium-blue fill with matching blue border, retain white label and icon at full opacity, and revise the regression expectation.

7. Match the supplied Schedule reference exactly by using the existing scheduled-event blue (`#76AFF2`) for both the disabled mobile button fill and border.

7. Move the disabled Schedule full-opacity treatment into its inline state style because the component-level CSS override did not win in the rendered Vuetify button; keep its background and border identical to the desktop blue token.

8. Match the desktop disabled-control mechanism: render the mobile Schedule button with Vuetify's explicit flat variant, whose disabled state remains opaque, and express enabled/disabled colors through existing Tailwind semantic tokens. Update the Event regression test to assert variant and state classes rather than a fragile inline-style implementation.

9. Correct the Vuetify-disabled opacity regression on the mobile Schedule action: retain the scheduled-event semantic classes, add a named disabled state class that owns `--v-disabled-opacity: 1` and `opacity: 1`, assert that class in the Event regression test, then run required frontend checks.

10. Add a Firefox touch regression that enters scheduling without a picked slot and reads the disabled button's computed background, border, color, and opacity. Compare it with the desktop disabled Schedule control, then correct only the mobile Vuetify disabled-state selector necessary to make the computed values match.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implemented FR-076, the white elevated availability-controls panel, and a component regression assertion for its dedicated panel wrapper.

Verification passed: `npm run test:unit -- src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts` (4 tests), `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` (135 files, 934 tests). `git diff --check` also passed.

Browser verification is blocked: `npm run test:e2e -- --project=firefox-touch e2e/schedule-overlap-mobile-touch-firefox.spec.ts` cannot launch Firefox because the host lacks required system libraries (including libgtk-3.so.0 and libX11.so.6). The earlier firefox-desktop invocation found no tests because this spec is assigned to the `firefox-touch` project. Keep the task In Progress until the browser dependencies are available and this E2E suite can run.

The supplied screenshot identified the actual green surface as Event.vue's mobile event action bar, rather than ScheduleOverlapMobileOverlay. The implementation and FR were corrected to target that action bar; the unrelated schedule-overlap edit was removed. White is now the neutral surface in normal and scheduling modes, while availability actions remain green and Schedule actions remain blue.

Corrected-target verification passed: `npm run test:unit -- src/views/Event.test.ts` (55 tests), `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` (135 files, 934 tests). `git diff --check` passed. The graph was updated after the Event.vue change.

The Firefox touch E2E suite remains blocked by missing host libraries, as recorded previously; no browser result is claimed.

Refined disabled mobile Schedule to match the supplied desktop reference: solid `#006BE8` fill and border, white foreground, and a scoped full-opacity override for Vuetify's disabled state. Event regression coverage now asserts disabled semantics, colors, and the opacity override. Passed `npm run test:unit -- src/views/Event.test.ts` (55 tests), `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` (135 files, 934 tests). Build completed with pre-existing environment/CSS warnings. Ran `graphify update .`; it reported existing zero-node JSON and missing SQL parser warnings. Firefox E2E remains blocked by missing host system libraries.

Current staged code sets `--v-disabled-opacity: 1` for the mobile disabled Schedule button but does not set `opacity: 1`. Firefox inspection will verify whether Vuetify's `.v-btn--disabled` opacity rule still composites the scheduled-event token over the white action bar.

Firefox touch is now available. Added and ran the focused Playwright regression `disabled mobile Schedule keeps the scheduled-event blue at full opacity`: it creates a timed event, enters the mobile scheduling mode without selecting a meeting slot, and confirms computed background and border `rgb(118, 175, 242)`, white foreground, and `opacity: 1`. It passed. The pre-existing Firefox touch suite run now executes but has an unrelated existing failure in `touching a timeslot keeps its mobile tooltip anchored while scrolling` at line 153; the other three existing checks passed. Required checks passed: `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` (135 files, 934 tests). Build retained existing env and `:deep` CSS warnings. `graphify update .` completed with existing zero-node JSON and missing SQL parser warnings.
<!-- SECTION:NOTES:END -->
