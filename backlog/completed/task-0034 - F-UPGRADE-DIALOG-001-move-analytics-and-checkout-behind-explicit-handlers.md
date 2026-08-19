---
id: TASK-0034
title: 'F-UPGRADE-DIALOG-001: move analytics and checkout behind explicit handlers'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:01'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-UPGRADE-DIALOG-001.md
priority: high
type: task
ordinal: 34000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/pricing/UpgradeDialog.vue

Problem: Watchers still trigger analytics and initialization, and checkout flow redirects through window.location.href.

Why it matters: Pricing flow ownership is spread across view state and browser-side effects, which makes upgrade behavior fragile.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move analytics and checkout coordination behind explicit actions or composables
- [ ] #2 Replace direct redirect wiring where appropriate
- [ ] #3 Preserve current upgrade behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 14:01
---
**Verification evidence:**
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed in `frontend/`.
- Existing and updated coverage passed in `src/components/pricing/UpgradeDialog.test.ts`.

**Implementation notes:**
- Upgrade analytics, student-toggle tracking, dialog-view tracking, and checkout/sign-up redirects now run through explicit handlers instead of watcher-owned side effects while keeping the existing pricing and navigation behavior unchanged.
---
<!-- COMMENTS:END -->
