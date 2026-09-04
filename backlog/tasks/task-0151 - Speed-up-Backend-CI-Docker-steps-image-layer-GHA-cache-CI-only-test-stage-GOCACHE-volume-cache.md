---
id: TASK-0151
title: >-
  Speed up Backend CI Docker steps (image/layer GHA cache, CI-only test stage,
  GOCACHE volume cache)
status: In Progress
assignee: []
created_date: '2026-09-04 07:23'
updated_date: '2026-09-04 08:06'
labels: []
milestone: CI reliability and speed
dependencies: []
priority: medium
type: chore
ordinal: 164300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The backend-quality job (~4m9s, run 33804085825 job 100810256124) re-pulls base images and recompiles Go from zero on every run because hosted runners are ephemeral and no cacheable builder is configured. Breakdown: start DB services 80s (pulls + migrator build), backend tests 96s (builder target builds a full server binary whose output is discarded, then go test recompiles with a cold GOCACHE volume), browser E2E 52s, nix save 17s.

Plan (ranked):
1. Add a CI-only Dockerfile stage that stops after `go mod download` (no `RUN go build`); point compose.test.yaml server-route-test and server-test build targets at it.
2. Set up buildx and pre-build the server test images with docker/build-push-action using cache-from/cache-to type=gha mode=max, loading them tagged with the Compose-computed image names (timeful-test-server-route-test, timeful-test-postgres-test-migrate) so compose up/run skip builds; drop `--build` from the up step.
3. Persist the GOCACHE volume with actions/cache: restore a host directory, copy it into the external timeful-test-go-build-cache volume before tests, copy back after tests for saving.
4. Minor: keep mongo:7 pull as-is or digest-pin it; pre-pulling optional.

Constraints: must not change test semantics (`-count=1` stays); E2E stack ownership (frontend/e2e isolated-test-stack.ts) must keep working, including its mirror of ensureGoBuildCacheVolume; keep the workflow's existing comments updated where behavior changes.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Backend CI no longer builds the discarded full-server `go build` layer: test Compose services use a Dockerfile stage that stops after module download, and runtime behavior of server-route-test/server-test is unchanged.
- [ ] #2 Server images used by the test stack are built with a buildx builder using the GitHub Actions cache backend (cache-from/cache-to type=gha, mode=max) and tagged with the Compose-computed names so up/run skip redundant builds.
- [ ] #3 The Go build cache (GOCACHE volume) is persisted across CI runs via actions/cache, restored into the external volume before tests and saved after.
- [ ] #4 The modified workflow is validated (e.g., actionlint or equivalent syntax check) and the full workflow passes on a pushed branch.
- [ ] #5 The measured job duration on a run with a warm cache is recorded in the task and is meaningfully lower than the ~4m9s baseline.
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
1. server/Dockerfile: extract `testdeps` stage (golang base + go.mod/go.sum COPY + go mod download); retarget `builder` stage FROM testdeps to dedupe module-download layers.\n2. compose.test.yaml: point server-route-test and server-test `build.target` at `testdeps` (test containers mount ./server and compile live, so the full binary build was discarded work).\n3. .github/workflows/backend-ci.yml:\n   a. docker/setup-buildx-action (docker-container driver) after compose config validation.\n   b. Two docker/build-push-action steps (testdeps -> timeful-test-server-route-test, migrator -> timeful-test-postgres-test-migrate) with cache-from/cache-to type=gha mode=max and load: true; exact-commit pinned like existing actions.\n   c. actions/cache for .go-build-cache keyed on server/go.sum; seed step creates the external volume and copies the restored cache in; save step copies it back out with if: always() before down -v.\n   d. Drop `--build` from the up step so Compose uses the pre-built tags.\n4. Leave frontend/e2e/isolated-test-stack.ts untouched: its `up -d --build server-test` keeps local semantics; on CI it now builds the cheaper testdeps target instead of the discarded binary.\n5. Validate with actionlint + compose config, verify the full route-test suite locally against the test stack, then push a branch and measure cold and warm CI runs (rerun reuses GHA cache).
<!-- SECTION:PLAN:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-09-04 08:06
---
Verified locally with Docker 28.2.2 before pushing:
- `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml config` confirms project name `timeful-test`, so image tags are `timeful-test-server-route-test` and `timeful-test-postgres-test-migrate`.
- `docker compose build server-route-test postgres-test-migrate` builds cleanly with the new `testdeps` stage (no full server binary build).
- `up -d` without `--build` starts the stack from the pre-built images; bootstrap and migrate one-shots complete.
- `compose run --rm server-route-test go test ./... -count=1` passes all 14 test packages (repeat runs are fast thanks to the GOCACHE volume).
- actionlint 1.7.12 passes on backend-ci.yml.
Note for local devs: after changing a Compose build target, run `docker compose build <service>` once or remove the stale image, since compose reuses an existing image without detecting target changes. CI is unaffected (fresh runners).
---
<!-- COMMENTS:END -->
