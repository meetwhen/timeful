---
id: TASK-0035
title: 'F-UPGRADE-DIALOG-002: separate upgrade dialog presentation from provider logic'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:01'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-UPGRADE-DIALOG-002.md
priority: medium
type: task
ordinal: 35000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/pricing/UpgradeDialog.vue

Problem: Pricing presentation, freemium gating, analytics, and checkout coordination still live in a single large dialog component.

Why it matters: UI work and behavior changes are coupled more tightly than they need to be, which raises regression risk.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Separate dialog presentation from provider-flow and gating logic
- [ ] #2 Preserve pricing parity and upgrade behavior
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
- Expanded coverage passed in `src/components/pricing/UpgradeDialog.test.ts`.

**Implementation notes:**
- Pricing fetches and checkout-session creation now live behind typed helpers, while `useUpgradeDialogState` owns dialog analytics, price derivation, student-mode state, and upgrade routing so `UpgradeDialog.vue` can stay presentation-focused without changing its public contract.
---
<!-- COMMENTS:END -->
