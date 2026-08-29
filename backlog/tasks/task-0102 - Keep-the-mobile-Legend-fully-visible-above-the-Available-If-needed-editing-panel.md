---
id: TASK-0102
title: >-
  Keep the mobile Legend fully visible above the Available/If needed editing
  panel
status: In Progress
assignee:
  - '@OpenCode'
created_date: '2026-08-29 10:57'
updated_date: '2026-08-29 11:30'
labels:
  - mobile
  - schedule-overlap
  - ui
dependencies: []
references:
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - frontend/src/views/Event.vue
modified_files:
  - frontend/src/components/schedule_overlap/ScheduleOverlap.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlap.mobileLegendClearance.test.ts
priority: medium
type: bug
ordinal: 108000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On mobile, while editing availability, the fixed Available/If needed panel (ScheduleOverlapMobileOverlay elevated panel above the Event bottom action bar) covers the bottom of the event-page Legend Section (ColorLegend at the end of ScheduleOverlapSidebar). The last legend rows, including the Scheduled event row, cannot be brought above the fixed panel stack. The page must reserve in-flow space for the fixed mobile bottom panel stack while editing so the Legend Section is fully visible at maximum scroll.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On a phone viewport, while editing availability, the event-page Legend Section can be scrolled so that no legend row is obscured by the fixed Available/If needed panel or the bottom action bar, including the last Scheduled event row
- [x] #2 The fix reserves in-flow layout space for the fixed mobile bottom panel stack instead of covering, z-index, or scroll-hack overrides
- [x] #3 Unit regression coverage verifies the mobile legend clearance contract
- [x] #4 Required frontend checks (lint, typecheck, build, unit) pass
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
1. Measure the ScheduleOverlapMobileOverlay rendered height with a ResizeObserver (guarded like existing usages) and emit it to the parent; the overlay root already sits above the bottom action bar offset.
2. In ScheduleOverlap, while on phone and editing (the states that render the elevated toggle/week panels), apply the measured overlay height plus the overlay bottom offset as in-flow padding-bottom on the schedule-overlap root so the Legend Section clears the fixed stack at maximum scroll.
3. Add unit regression coverage asserting the clearance contract (padding applied while editing on phone, absent otherwise), following scheduleOverlapTestUtils stubs.
4. Run required frontend checks and update the graph.
<!-- SECTION:PLAN:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-29 11:30
---
Root cause of the earlier failed attempt: the project sets `important: true` in tailwind.config.cjs, so `tw-py-2` emits `padding-bottom: .5rem !important` and silently overrode the inline `paddingBottom` binding (computed padding stayed 8px). Fix: replaced the inline padding with a conditional in-flow spacer element `.schedule-overlap__mobile-editing-clearance` (inline height is not fought by any utility), sized `calc(<measured overlay height>px + <bottomOffset> - 64px)` where 64px is the static tail already below the legend at max scroll (sidebar/root bottom padding + page footer gap). Verified in a real mobile Chromium viewport (390x844) against /e/6cc4d: last Scheduled-event legend row clears the fixed panel stack by 26px at max scroll; desktop and non-editing states reserve nothing.
---
<!-- COMMENTS:END -->
