---
id: TASK-0150
title: >-
  Fix timed-event-postgres-plugin-firefox spec blockers: grid readiness wait and
  cloneable get-slots payload
status: Done
assignee: []
created_date: '2026-09-03 15:09'
updated_date: '2026-09-03 15:23'
labels:
  - playwright
  - frontend
dependencies: []
references:
  - frontend/e2e/timed-event-postgres-plugin-firefox.spec.ts
  - frontend/src/views/Event.vue
  - frontend/src/views/event/pluginResponsesBoundary.ts
  - PLUGIN_API_README.md
modified_files:
  - frontend/e2e/timed-event-postgres-plugin-firefox.spec.ts
  - frontend/src/views/event/pluginResponsesBoundary.ts
  - frontend/src/utils/pluginBoundaries.regressions.test.ts
  - frontend/src/components/schedule_overlap/ScheduleOverlapTimeGrid.vue
priority: high
type: bug
ordinal: 163300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The first-ever run of frontend/e2e/timed-event-postgres-plugin-firefox.spec.ts (2026-09-03, via nix run .#frontend-e2e during TASK-0149) exposed two defects that make the spec unable to pass. TASK-0149 AC #5/#6 are blocked on this work.

Defect 1 - plugin message races grid mount: the spec posts set-slots immediately after the /api/events/{id} response resolves, but Event.vue registers the plugin listener at mount (Event.vue:2293) while the ScheduleOverlap grid mounts later via requestAnimationFrame+setTimeout (queueScheduleOverlapMount, Event.vue:1701-1718). timeSlotToRowCol comes from scheduleOverlap.getAllValidTimeRanges() and is an empty Map at that point, so set-slots returns "Time slot at index 0 ... falls outside the event's date/time range". Evidence: a throwaway diagnostic spec (since deleted) identical except awaiting .schedule-overlap-time-grid__scroller visibility made set-slots succeed.

Defect 2 - get-slots payload is not structured-cloneable: normalizePluginResponses (frontend/src/views/event/pluginResponsesBoundary.ts) returns PluginSlotEntry.availability/.ifNeeded as Temporal.ZonedDateTime[] and getSlots (Event.vue:2206-2216) passes them into sendPluginSuccess -> window.postMessage. Structured clone throws DataCloneError "Temporal.ZonedDateTime object could not be cloned", which the catch converts to "Failed to fetch responses: ...". PLUGIN_API_README.md documents availability as plain local ISO strings (e.g. ["2026-01-07T09:00:00", ...]), so the app violates its own documented plugin contract; browsers without native Temporal can never receive a successful get-slots response.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The spec performs an objective grid-readiness wait (e.g. the ScheduleOverlap time grid becoming visible) before posting any plugin message, with no fixed sleeps and no blind timeout raises
- [x] #2 get-slots success payload availability and ifNeeded values are plain local-ISO strings matching PLUGIN_API_README.md, so window.postMessage never receives Temporal objects
- [x] #3 Plugin response boundary serialization has unit regression coverage (Temporal.ZonedDateTime input yields the documented string payload)
- [x] #4 E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true nix run .#frontend-e2e -- --project=firefox-desktop timed-event-postgres-plugin-firefox.spec.ts runs green against the isolated test stack
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Defect 1: add a stable data-testid to the ScheduleOverlapTimeGrid scroller and make the spec await its visibility with an auto-retrying expect(...).toBeVisible() before posting set-slots (no sleeps, no timeout raises). 2. Defect 2: change the plugin response boundary (pluginResponsesBoundary.ts) so normalizePluginResponses emits the documented wire shape directly: availability/ifNeeded as plain local-ISO strings, keeping Temporal arithmetic (DOW shift) internal until serialization. 3. Update pluginBoundaries.regressions.test.ts expectations from Temporal accessors to documented strings and add a structured-cloneability regression test (AC #3). 4. Run required frontend checks and the AC #4 e2e command against the isolated test stack.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
- Defect 1 (spec-side fix): added an objective grid-readiness wait to frontend/e2e/timed-event-postgres-plugin-firefox.spec.ts. Before posting any plugin message the spec asserts page.getByTestId("schedule-overlap-time-grid-scroller") is visible. A data-testid="schedule-overlap-time-grid-scroller" was added to the ScheduleOverlapTimeGrid scroller div because no accessible role exists. The auto-retrying toBeVisible() replaces fixed sleeps; grid visibility implies the ScheduleOverlap grid mounted, so getAllValidTimeRanges() is populated for timed events (days/times are synchronous computeds from the event).
- Defect 2 (boundary fix): frontend/src/views/event/pluginResponsesBoundary.ts now serializes the get-slots payload at the plugin response boundary. PluginSlotEntry.availability/ifNeeded are string[] of plain local ISO values (toPlainDateTime().toString({ smallestUnit: "seconds" }) -> e.g. "2026-01-07T09:00:00"), matching PLUGIN_API_README.md. DOW one-hour-shift semantics still run on Temporal.ZonedDateTime before serialization. window.postMessage never receives Temporal objects, so structured clone cannot throw DataCloneError. Event.vue needed no change; it already passes the boundary output straight to sendPluginSuccess.
- Unit coverage: pluginBoundaries.regressions.test.ts updated from Temporal accessors (.hour/.timeZoneId) to exact documented string assertions for the fixed-offset, LA, and DOW round-trip cases, plus a new regression test that feeds Temporal.ZonedDateTime input and asserts the documented string payload survives structuredClone (reproduces the DataCloneError defect on Node 26 native Temporal).
- Verification: npm run lint (clean), npm run fmt:check (clean), npm run typecheck (clean), npm run build (success), npm run test:unit (138 files, 991 tests passed), and E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true nix run .#frontend-e2e -- --project=firefox-desktop timed-event-postgres-plugin-firefox.spec.ts -> 1 passed (26.9s) against the isolated test stack.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Fixed both TASK-0149 blockers for frontend/e2e/timed-event-postgres-plugin-firefox.spec.ts.

Grid readiness (AC #1): the spec now awaits page.getByTestId("schedule-overlap-time-grid-scroller") with an auto-retrying expect(...).toBeVisible() before posting any plugin message; the new test id was added to the ScheduleOverlapTimeGrid scroller div because no accessible role exists. Grid visibility guarantees the ScheduleOverlap grid mounted, so getAllValidTimeRanges() is populated when set-slots validates slots.

Cloneable get-slots payload (AC #2): frontend/src/views/event/pluginResponsesBoundary.ts now serializes the get-slots payload at the plugin response boundary. PluginSlotEntry.availability/ifNeeded are plain local-ISO strings (e.g. "2026-01-07T09:00:00") matching PLUGIN_API_README.md, produced via toPlainDateTime().toString({ smallestUnit: "seconds" }) after the DOW one-hour shift. window.postMessage never receives Temporal objects, so structured clone cannot throw DataCloneError. Event.vue required no change.

Regression coverage (AC #3): pluginBoundaries.regressions.test.ts now asserts the exact documented string payload for LA, fixed-offset, and DOW round-trip cases, and a new test feeds Temporal.ZonedDateTime input and asserts the payload survives structuredClone (reproduces the original DataCloneError on Node 26 native Temporal).

Verification (AC #4): npm run lint, fmt:check, typecheck, build all clean; npm run test:unit 138 files / 991 tests passed; E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true nix run .#frontend-e2e -- --project=firefox-desktop timed-event-postgres-plugin-firefox.spec.ts passed (1 passed, 26.9s) against the isolated test stack. No Markdown files changed, so DoD #4 is exempt.
<!-- SECTION:FINAL_SUMMARY:END -->
