---
id: TASK-0022
title: 'F-SCHEDULE-OVERLAP-004: stale specific-times slots leaked into edit drafts'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SCHEDULE-OVERLAP-004.md
priority: medium
type: bug
ordinal: 22000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/composables/event/specificTimesEditDraft.ts

Problem: buildSpecificTimesEditDraft(...) leaked stale canonical slot state in two edit paths. Removing a date in the edit dialog's date picker did not remove that date's slots from enabledSlots in the draft, and broad timed edits with specificTimesEnabled: false could preserve contradictory out-of-window slots from prior saved enabledSlots instead of rewriting the canonical window from the current schedule.

Why it matters: This broke the core edit flow - users could not remove dates from specific-times events.

Root cause: buildSpecificTimesEditDraft reused enabledSlotsSource too broadly. In the specific-times-enabled merge path, when sameSelectedDays was false but slotWindowMatches was true and timezoneChanged was false, the code merged old event enabledSlots (including slots for all dates) with new schedule enabledSlots (only the new date set), preserving slots for removed dates. In the !specificTimesEnabled branch, the draft filtered enabledSlotsSource rather than the current schedule.enabledSlots, so contradictory stale slots from prior state survived broad timed edits. Later, getEventDateSeeds -> buildTimedDateSeeds -> projectSlotsToLocalDays extracted unique local dates from the leaked enabledSlots, reviving removed or out-of-window state.

Fix direction: Keep the date-filtering guard, split the slot sources by mode. Specific-times-enabled branch keeps filtering enabledSlotsSource to the projected date set. !specificTimesEnabled branch derives the canonical slot set directly from schedule.enabledSlots, then filters that schedule-derived set by the normalized selected days as a safety guard. normalizeActiveSlots then rebuilds both enabledSlots and activeSlots from the canonical schedule window for broad timed edits.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 When a user removes a date from the date picker during edit, the specific times grid must not show a column for the removed date
- [ ] #2 When a user saves a timed edit with specific-times disabled, the canonical enabledSlots and activeSlots must be rebuilt from the current schedule window, not merged with historical sparse slots
- [ ] #3 Reopening specific-times after that save must seed only the rewritten canonical window
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 14:00
---
**Verification evidence:**
- Added regression tests `"filters out enabledSlots for removed dates when dates changed"` and `"rewrites non-specific timed edits to the schedule canonical slots instead of preserving stale out-of-window slots"` in `specificTimesEditDraft.test.ts`, plus Firefox browser regression `"broad timed edits rewrite contradictory canonical slots before reopen-specific-times seeding"` in `timed-event-specific-times-edit-firefox.spec.ts`.
- Frontend lint, typecheck, build, and unit tests pass.

**Implementation notes:**
- Exported `getLocalSlotDomainDay` from `timedEventSlots.ts` to reuse the same date-projection helper used by `projectSlotsToLocalDays`, ensuring consistent wrapped-window handling.
---
<!-- COMMENTS:END -->
