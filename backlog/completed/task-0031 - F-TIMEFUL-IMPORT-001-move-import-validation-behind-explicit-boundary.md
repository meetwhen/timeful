---
id: TASK-0031
title: 'F-TIMEFUL-IMPORT-001: move import validation behind explicit boundary'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-TIMEFUL-IMPORT-001.md
priority: high
type: task
ordinal: 31000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/TimefulImportDialog.vue

Problem: Import validation depends directly on window.location.hostname, and dialog close clears form state through a writable computed shim.

Why it matters: Environment checks and form-lifecycle rules are buried inside the dialog, which makes the flow harder to reuse and test.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move environment-sensitive validation behind an explicit boundary
- [ ] #2 Separate dialog visibility from form reset rules
- [ ] #3 Preserve current import behavior
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
- Added component coverage for the extracted import-URL validation boundary and explicit close/reset behavior in `src/components/TimefulImportDialog.test.ts`.

**Implementation notes:**
- Extracted hostname blocking into `src/utils/timefulImport.ts` and replaced the writable computed dialog shim with explicit close and visibility handlers.
---
<!-- COMMENTS:END -->
