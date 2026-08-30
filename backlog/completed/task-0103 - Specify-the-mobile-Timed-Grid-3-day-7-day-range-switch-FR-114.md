---
id: TASK-0103
title: Specify the mobile Timed Grid 3-day/7-day range switch (FR-114)
status: Done
assignee: []
created_date: '2026-08-29 11:05'
updated_date: '2026-08-29 11:09'
labels: []
dependencies: []
references:
  - docs/requirements/functional/fr/FR-049.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-018.md
  - frontend/src/components/schedule_overlap/ToolRow.vue
documentation:
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/terminology/README.md
modified_files:
  - docs/requirements/functional/fr/FR-114.md
  - docs/requirements/README.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-018.md
type: docs
ordinal: 109000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The mobile event page offers a `3 days`/`7 days` range switch for Timed Events, but no requirement specifies when it is shown or what it does. Decisions confirmed with the product owner: use one new requirement (FR-114) covering the switch effect, its visibility rule, and persistence; hide the switch whenever the Timed Grid can display 3 or fewer day columns (it is redundant exactly when both options render identical content); keep Dates-Only Event pages excluded as FR-049 already does; document the browser-local persistence with a 3 days default. Documentation-only change; no runtime code, tests, or build configuration are affected.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 FR-114.md exists in docs/requirements/functional/fr/ with front matter id FR-114, type functional, components [frontend], status proposed
- [x] #2 FR-114 states: on a mobile Timed Event Page the 3 days/7 days range switch sets how many day columns each Timed Grid page displays; the switch is shown only when the Timed Grid can display more than 3 day columns; Dates-Only Event pages do not show the switch; the selection persists in the visitor's browser and defaults to 3 days
- [x] #3 Each Markdown sentence in every edited file sits on one physical source line and controlled terms use canonical glossary forms with first-occurrence-per-paragraph links per docs/terminology/README.md
- [x] #4 docs/requirements/README.md Functional Requirements table gains an FR-114 row linking functional/fr/FR-114.md with components frontend
- [x] #5 CAND-018 records verdict covered, related requirement FR-114, confidence confirmed, a disposition naming FR-114, and resolved open questions
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
"1. Create `docs/requirements/functional/fr/FR-114.md`: YAML front matter (id FR-114, title, type functional, components [frontend], status proposed) followed by a title heading and four prose sentences, one sentence per physical source line: (a) on a mobile Timed Event Page the `3 days`/`7 days` range switch sets how many day columns each Timed Grid page displays; (b) show the switch only when the Timed Grid can display more than 3 day columns, otherwise hide it and display all day columns; (c) Dates-Only Event pages shall not show the switch; (d) the selection persists in the visitor's browser and defaults to `3 days`. Link controlled terms at first occurrence per paragraph to ../../../terminology/glossary.md anchors (#timed-event-page, #timed-grid, #dates-only-event); bold repeated occurrences in the same paragraph. Scope is mobile Timed Event Pages, so desktop needs no exclusion sentence.\n2. Append an FR-114 row to the Functional Requirements table in docs/requirements/README.md (ID, linked title with Timed Grid glossary link, Components: frontend).\n3. Update docs/requirements/migration/backlog-fr-inventory-candidates/CAND-018.md: verdict covered, related_requirements [FR-114], confidence confirmed; rewrite Existing Requirements, Disposition (specified by FR-114), and resolve Open Questions.\n4. Run npm run format:markdown on the changed files and verify it leaves them unchanged.\n5. Verify each acceptance criterion and finalize per the task-finalization guide.\nResearch notes: the switch is the segmented toggle in ToolRow.vue mobile row 1 (mobileNumDaysOptions, labels `3 days`/`7 days`), stored browser-locally with default 3 by useCalendarGrid.ts, and drives maxDaysPerPage (phone uses the setting, desktop always 7). Weekly Timed Events always render 7 columns; specific-dates Timed Events render one column per picked date plus cross-midnight neighbor columns. CAND-015 shows the covered-verdict record shape."
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Verified the switch implementation in ToolRow.vue (mobileNumDaysOptions labels `3 days`/`7 days`) and useCalendarGrid.ts (browser-local persistence, 3-day default, maxDaysPerPage: phone uses the setting, desktop always 7) before wording the requirement.

Visibility rule expressed observably as day-column capacity rather than picked-date counts so weekly Timed Events (always 7 columns) and specific-dates events (picked dates plus cross-midnight neighbor columns) both stay decidable.

npm run format:markdown covers only git-tracked .md files, so FR-114.md was verified separately with the repo Prettier config and the sentences-per-line plugin; it required no changes.

CAND-018 front matter follows the covered-record shape used by CAND-015 (no requirement_type field); classification stays `candidate FR` with the disposition naming FR-114.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added FR-114, "Switch mobile Timed Grid pages between 3 and 7 day columns", specifying the mobile event-page `3 days`/`7 days` range switch: it sets how many day columns each Timed Grid page displays on a mobile Timed Event Page; it is shown only when the Timed Grid can display more than 3 day columns (hidden otherwise, with every day column displayed); Dates-Only Event pages never show it; and the selection persists in the visitor's browser with a `3 days` default. Added the FR-114 row to the docs/requirements/README.md functional-requirements index and migrated CAND-018 to verdict covered with related requirement FR-114, resolving its open questions (segmented toggle in the mobile tool row, Timed Event kind, browser-local persistence). Documentation-only change: no runtime code, tests, or build configuration touched. Ran npm run format:markdown (clean) and verified FR-114.md is byte-identical under the repo Prettier config with the sentences-per-line plugin. Glossary anchors (#timed-event-page, #timed-grid, #dates-only-event) verified against docs/terminology/glossary.md headings.
<!-- SECTION:FINAL_SUMMARY:END -->
