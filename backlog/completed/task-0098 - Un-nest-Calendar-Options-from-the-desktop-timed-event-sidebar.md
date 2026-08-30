---
id: TASK-0098
title: Un-nest Calendar Options from the desktop timed event sidebar
status: Done
assignee: []
created_date: '2026-08-29 09:42'
updated_date: '2026-08-29 09:42'
labels:
  - frontend
  - schedule-overlap
  - ui
dependencies: []
modified_files:
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.test.ts
  - frontend/e2e/event-page-no-responses-layout.spec.ts
type: enhancement
ordinal: 104000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
On the desktop timed event response editing/creation page, the sidebar renders an "Options" collapsible section whose only content is the "Calendar options..." button. The section header adds no information and hides the button behind a persisted collapse state.

Outcome:
- Desktop sidebar shows the "Calendar options..." button directly, with no "Options" section header.
- The calendar options dialog keeps its current contents and behavior.
- Mobile keeps the existing "Options" expandable section around the button.

Constraints:
- Change layout in the sidebar component; do not alter buffer time or working hours update behavior.
- Keep mobile behavior unchanged; scope the removal to the non-phone branch.
- Update any layout e2e assertions that anchor on the removed expandable section toggle.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On desktop timed event response editing, the sidebar shows no Options expandable section and the Calendar options... button renders directly in the sidebar and opens the calendar options dialog
- [x] #2 On mobile, the Options expandable section wrapping the Calendar options... button is preserved
- [x] #3 The e2e desktop sidebar layout spec anchors the last edit control above the Legend on the un-nested Calendar options... button
- [x] #4 lint, typecheck, build, and unit tests pass, and the updated e2e spec passes on the chromium-desktop project
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
Removed the "Options" expandable section from the desktop timed event response editing sidebar and un-nested the "Calendar options..." button so it renders directly in the sidebar.

Changes:
- frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue: the options block now branches on sidebar.isPhone. Desktop renders the "Calendar options..." button directly (new calendar-options-button class) and a shared controlled v-dialog; mobile keeps the existing ExpandableSection "Options" wrapper. Dialog contents (buffer time, working hours, group alert) and all update handlers are unchanged.
- frontend/e2e/event-page-no-responses-layout.spec.ts: "timed add availability controls stay close to the Legend" now anchors the last edit control on .calendar-options-button instead of .expandable-section-toggle.
- frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.test.ts: added coverage that desktop renders the button without the Options section and emits update:calendarOptionsDialog on click, and that mobile keeps the Options section wrapping the button.

Verification: npm run lint, typecheck, build, and test:unit (934 tests) pass; the updated e2e spec passes on chromium-desktop via the isolated test stack. Pre-existing unrelated failure: the "pairs each header row" test in the same spec expects a "+ Add description" button removed in f6f65e7d.
<!-- SECTION:FINAL_SUMMARY:END -->
