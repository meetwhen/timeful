---
id: TASK-0134
title: >-
  Fix mobile touch tooltip not re-anchoring to the touched slot after scroll on
  Firefox
status: To Do
assignee: []
created_date: '2026-09-02 10:23'
updated_date: '2026-09-02 10:23'
labels:
  - e2e
  - frontend
dependencies: []
priority: medium
type: bug
ordinal: 147300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The e2e spec frontend/e2e/schedule-overlap-mobile-touch-firefox.spec.ts:146 ("touching a timeslot keeps its mobile tooltip anchored while scrolling") fails reproducibly on the firefox-touch project: after touching a timeslot and running window.scrollBy({ top: 50 }), the expect.poll predicate that compares the tooltip's inline left/top style against the touched slot's bounding rect times out after 5s, so the tooltip no longer re-anchors to the slot while scrolling.

Evidence (2026-09-02): failed in a full npm run test:e2e run, failed again in a targeted multi-spec run, and failed in isolation on firefox-touch (5 passed, 1 failed) against the TASK-0132 worktree. It is pre-existing and independent of the Add-description spec cleanup in TASK-0132, which does not touch the tooltip path.

Diagnosis direction: the shared Tooltip component's repositioning appears not to track page scroll on touch devices. TASK-0033 ("F-TOOLTIP-001: centralize tooltip placement and visibility logic", Done) centralized tooltip placement and may be the origin of a regression; review its placement logic for scroll listeners and touch-specific position overrides. Decide after diagnosis whether the app must keep the tooltip anchored during scroll or the spec's expectation is wrong, and record that decision.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The root cause of the tooltip failing to re-anchor to the touched slot after scroll on firefox-touch is identified (app bug vs stale spec expectation) and the decision is recorded in the task.
- [ ] #2 On the firefox-touch project, either the app keeps the touched slot's tooltip anchored while the page scrolls, or the spec's anchoring expectation is corrected with recorded justification.
- [ ] #3 frontend/e2e/schedule-overlap-mobile-touch-firefox.spec.ts passes fully on the firefox-touch project.
- [ ] #4 The required frontend checks (lint, typecheck, build, test:unit) pass.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
