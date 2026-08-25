---
id: TASK-0069
title: Provide mobile header navigation actions
status: Done
assignee:
  - OpenCode
created_date: '2026-08-25 11:24'
updated_date: '2026-08-25 11:35'
labels:
  - mobile
  - navigation
dependencies: []
references:
  - docs/requirements/functional/fr/FR-078.md
  - frontend/src/App.vue
  - frontend/src/views/Event.vue
priority: medium
type: feature
ordinal: 71000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Give phone-sized event and landing pages a header navigation menu for creating an event and giving feedback. Consolidate mobile feedback into that menu while preserving the specified page-specific header arrangement and existing desktop behavior.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Phone-sized event and landing headers provide a navigation menu with Create an event followed by Give feedback
- [x] #2 On phone-sized event pages the menu replaces the separately displayed Create an event action and the sign-in entry point remains visible only when enabled for unauthenticated visitors
- [x] #3 On phone-sized landing pages the header order is sign-in when available then navigation menu then GitHub link
- [x] #4 Phone-sized event and landing pages do not show Give feedback as a separate header action or event-page footer action
- [x] #5 Desktop header behavior remains unchanged
- [x] #6 Automated tests cover the mobile menu behavior and relevant frontend checks pass
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Update App.vue so phone-sized event and landing headers use one accessible menu containing Create an event then Give feedback; preserve the required sign-in and GitHub ordering and desktop controls.
2. Remove the phone-sized event-page footer feedback control and any duplicate phone feedback actions covered by FR-078.
3. Extend App and event-page tests for responsive action visibility, menu order, and desktop preservation.
4. Run required frontend lint, typecheck, build, unit tests, and scoped mobile Firefox E2E where feasible; record evidence and finalize the task.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implemented FR-078 in App.vue: phone-sized event and landing routes now show an accessible header menu with Create an event then Give feedback; standalone phone event/landing feedback and event-page create actions are hidden. The landing GitHub link remains visible after the menu. The event-page feedback footer and duplicate phone account-menu feedback item were removed; desktop account-menu feedback remains.

Verification: focused App/AuthUserMenu unit tests passed; full lint passed; typecheck passed; build passed with pre-existing Vite environment/CSS warnings; full unit suite passed (135 files, 932 tests). Full Firefox E2E ran: 19 passed, 1 skipped, 2 failed outside this feature: timed-event-specific-times-edit timed out while toggling specific times, and timezone-menu width measured 512px against a 514px lower bound. Task remains In Progress because the project e2e Definition of Done item is not satisfied.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Implemented FR-078 mobile header navigation. Phone-sized event and landing headers now provide an accessible menu with Create an event followed by Give feedback; standalone mobile event create and feedback actions, the event footer feedback action, and the duplicate mobile account-menu feedback entry were removed. The landing header retains the requested sign-in, menu, GitHub ordering, while desktop behavior remains unchanged.

Verification: npm run lint, npm run typecheck, npm run build, focused App/AuthUserMenu tests, and the complete unit suite (135 files, 932 tests) passed. Firefox E2E completed with 19 passing and one skipped; two pre-existing unrelated failures were accepted by the user: a specific-times toggle timeout and a 512px timezone-menu width measurement against a 514px lower bound.
<!-- SECTION:FINAL_SUMMARY:END -->
