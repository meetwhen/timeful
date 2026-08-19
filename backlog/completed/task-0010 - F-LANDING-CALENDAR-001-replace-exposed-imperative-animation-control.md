---
id: TASK-0010
title: 'F-LANDING-CALENDAR-001: replace exposed imperative animation control'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-LANDING-CALENDAR-001.md
priority: high
type: task
ordinal: 10000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/landing/LandingPageCalendar.vue

Problem: Animation control depends on mount-time setup, timers, and defineExpose({ playAnimation }).

Why it matters: Instance-style parent coordination makes the component harder to compose and test, and ties behavior to render timing.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Replace exposed imperative coordination with clearer state-driven control
- [ ] #2 Keep parity with the legacy landing animation
- [ ] #3 Preserve current timing behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:58
---
**Verification evidence:**
- Added `src/composables/useLandingPageCalendarAnimation.ts` to replace `defineExpose({ playAnimation })` with a replay-token driven animation boundary and explicit timer cleanup.
- Updated `src/components/landing/LandingPageCalendar.vue` to consume the composable and remove mount-coupled exposed instance control.
- Added `src/composables/useLandingPageCalendarAnimation.test.ts` covering restart sequencing and cancellation of overlapping timer runs.
- `npm run compare:landing-styles` completed against the active landing route; the reported landing diffs are on the existing rendered surface, while `LandingPageCalendar.vue` is currently not rendered by `Landing.vue`.
- Ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit`.

**Implementation notes:**
- Use Firefox comparator evidence for landing-page parity before and after the refactor.
---
<!-- COMMENTS:END -->
