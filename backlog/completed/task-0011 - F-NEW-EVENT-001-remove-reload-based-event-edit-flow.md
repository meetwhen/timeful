---
id: TASK-0011
title: 'F-NEW-EVENT-001: remove reload-based event-edit flow'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-NEW-EVENT-001.md
priority: high
type: task
ordinal: 11000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/NewEvent.vue

Problem: The editor still uses window.location.reload() and storage-backed reload flags as part of the core event-editing flow.

Why it matters: Reload-driven control flow obscures ownership, disrupts testability, and makes regressions more likely when the editor changes.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Replace reload-based flow with explicit state transitions or navigation
- [ ] #2 Move persistence behind clear boundaries
- [ ] #3 Preserve event-editing behavior
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed in `frontend/`.
- Added edit-save coverage in `src/components/NewEvent.test.ts` and `src/components/NewDialog.test.ts`.

**Implementation notes:**
- Successful event edits now emit an explicit refresh event instead of writing `from-edit-event-*` storage flags and calling `window.location.reload()`. `Event.vue` refreshes the loaded event in-process and remounts `ScheduleOverlap` when specific-times state needs to be reseeded.
---
<!-- COMMENTS:END -->
