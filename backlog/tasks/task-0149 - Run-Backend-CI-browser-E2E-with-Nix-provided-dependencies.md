---
id: TASK-0149
title: Run Backend CI browser E2E with Nix-provided dependencies
status: To Do
assignee: []
created_date: '2026-09-03 14:46'
updated_date: '2026-09-03 14:46'
labels:
  - ci
  - nix
  - playwright
dependencies: []
references:
  - .github/workflows/backend-ci.yml
  - 'https://github.com/whensync/timeful/actions/runs/33316150651/job/99269864485'
documentation:
  - frontend/e2e/isolated-test-stack.ts
  - frontend/playwright.config.ts
  - docs/environments.md
modified_files:
  - flake.nix
  - .github/workflows/backend-ci.yml
priority: high
type: task
ordinal: 162300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Follow-up to TASK-0145 (closed once the run-tests step went green). Backend CI still fails at the "Run PostgreSQL browser lifecycle test" step because Playwright browsers are not installed on the runner.

Goal: provide the browser E2E dependencies from the repository flake in CI without pulling the whole devshell (the devshell carries go, python, backlog-md, and graphify that this step never needs), and keep the Nix store cached across runs via the nix-community/cache-nix-action step already present in .github/workflows/backend-ci.yml.

Verified facts (2026-09-03, pinned nixpkgs c8f90650):

- The pinned nixpkgs playwright-driver is 1.61.1 with browser revision firefox-1532, exactly matching frontend playwright-core 1.61.0 (node_modules/playwright-core/browsers.json), so PLAYWRIGHT_BROWSERS_PATH pointing at pkgs.playwright-driver.browsers is registry-compatible with the npm Playwright version.
- Nix playwright browsers are patched builds: the firefox binary's program interpreter and all shared libraries resolve inside /nix/store (readelf/ldd verified, 0 missing), so no apt browser dependencies are needed on ubuntu-latest.
- The pkgs.playwright-driver.browsers closure is 2.15 GB (firefox alone 1.07 GB), which fits the workflow's gc-max-store-size-linux: 5G cache budget, though the budget may need raising as the store grows.

Non-goals: the Nix devshell browsers (TASK-0067) stay as-is; speeding up npm ci via setup-node caching in Backend CI is out of scope.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 flake.nix exposes a Nix-run entry (package and app, e.g. frontend-e2e) whose runtime closure includes nodejs_26 and pkgs.playwright-driver.browsers but not the full devshell (no go, python, backlog-md, graphify)
- [ ] #2 The entry script sets PLAYWRIGHT_BROWSERS_PATH to the Nix browsers store path and runs npm ci followed by npm run test:e2e with pass-through arguments from frontend/
- [ ] #3 backend-ci.yml runs the browser lifecycle spec via nix run .#frontend-e2e with E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true, and the stray trailing-whitespace lines under the cache step are removed
- [ ] #4 cache-nix-action retains the browsers closure between runs without exceeding its gc-max-store-size budget (raise the 5G budget if the 2.15 GB browsers closure plus store growth proves too tight)
- [ ] #5 Local verification runs the exact CI command (E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true nix run .#frontend-e2e -- --project=firefox-desktop timed-event-postgres-plugin-firefox.spec.ts) green against the isolated test stack
- [ ] #6 Backend CI completes end to end green on the next run
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
Round 1 plan (carried from TASK-0145 round 4):

1. flake.nix: add `frontend-e2e` via pkgs.writeShellScriptBin; expose it as packages.<system>.frontend-e2e and apps.<system>.frontend-e2e so `nix run .#frontend-e2e` works; its runtime closure must contain only nodejs_26 and pkgs.playwright-driver.browsers (no devshell, no go/python/backlog-md/graphify).
2. Script body: set -euo pipefail; export PLAYWRIGHT_BROWSERS_PATH="${pkgs.playwright-driver.browsers}"; REPO_ROOT="$(git rev-parse --show-toplevel)"; cd "$REPO_ROOT/frontend"; npm ci; exec npm run test:e2e -- "$@".
3. .github/workflows/backend-ci.yml: replace the inline `E2E_...=true npm ci && E2E_...=true npm run test:e2e ...` browser step with `nix run .#frontend-e2e -- --project=firefox-desktop timed-event-postgres-plugin-firefox.spec.ts`, keeping E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true as step env; remove the stray trailing-whitespace lines left under the cache step.
4. Cache budget check: browsers closure is 2.15 GB against gc-max-store-size-linux: 5G; confirm the restored store plus run outputs stays under budget, otherwise raise the value.
5. Verify: actionlint on backend-ci.yml (actionlint is in the devshell); run the exact CI command locally against the isolated E2E stack (the harness owns mongo-test/postgres-test/server-test on 3003 and Vite on 4174, and creates the external timeful-test-go-build-cache volume itself); leave AC #6 unchecked until the next CI run completes end to end, keeping the task In Progress until then.
<!-- SECTION:PLAN:END -->
