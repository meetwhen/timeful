---
id: TASK-0018
title: 'F-RESPONDENTS-LIST-001: decompose respondents panel responsibilities'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-RESPONDENTS-LIST-001.md
priority: medium
type: task
ordinal: 18000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/schedule_overlap/RespondentsList.vue

Problem: Export behavior, DOM measurement, resize handling, and list-state coordination are still concentrated in one large component.

Why it matters: The respondents panel is difficult to evolve or test because several unrelated responsibilities change together.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Decompose the panel into clearer rendering and behavior boundaries
- [ ] #2 Preserve export and layout behavior
- [ ] #3 Keep parity intact
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:59
---
**Verification evidence:**
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed in `frontend/`.
- Existing `src/components/schedule_overlap/RespondentsList.test.ts` coverage stayed green across the extraction, and Firefox `inspect -- --target event-respondents-panel` passed against the migrated app.

**Implementation notes:**
- `RespondentsList.vue` now delegates export behavior, desktop sizing, and respondent ordering or selection state to `useRespondentsCsvExport.ts`, `useRespondentsListSizing.ts`, and `useRespondentsListState.ts`, while keeping the emitted parent contract and sidebar DOM surface intact.
---
<!-- COMMENTS:END -->
