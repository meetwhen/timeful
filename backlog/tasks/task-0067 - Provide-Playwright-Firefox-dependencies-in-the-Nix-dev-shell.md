---
id: TASK-0067
title: Provide Playwright Firefox dependencies in the Nix dev shell
status: In Progress
assignee:
  - OpenCode
created_date: '2026-08-25 10:08'
updated_date: '2026-08-25 10:08'
labels:
  - nix
  - e2e
  - playwright
dependencies: []
references:
  - flake.nix
  - frontend/playwright.config.ts
priority: medium
type: chore
ordinal: 69000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Enable contributors using the repository Nix development shell to launch the Playwright-managed Firefox browser for the frontend E2E suite. The current shell provides Node and tooling but lacks the desktop runtime libraries required by Firefox.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Entering the default Nix development shell provides the runtime libraries needed for Playwright Firefox to launch
- [ ] #2 The Firefox desktop Playwright E2E suite can start without missing shared-library errors
- [ ] #3 The flake evaluates successfully
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Inspect the default development shell and verify available Nixpkgs Playwright browser support. 2. Add the minimal Firefox browser package and environment needed for the Playwright test runner. 3. Evaluate the flake and run the Firefox desktop E2E suite from the Nix development shell, recording any remaining environmental blocker.
<!-- SECTION:PLAN:END -->
