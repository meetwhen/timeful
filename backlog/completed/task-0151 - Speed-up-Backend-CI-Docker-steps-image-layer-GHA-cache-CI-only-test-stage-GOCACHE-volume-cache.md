---
id: TASK-0151
title: >-
  Speed up Backend CI Docker steps (image/layer GHA cache, CI-only test stage,
  GOCACHE volume cache)
status: Done
assignee: []
created_date: '2026-09-04 07:23'
updated_date: '2026-09-04 12:29'
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
- [x] #1 Backend CI no longer builds the discarded full-server `go build` layer: test Compose services use the module-cache-free `testbase` Dockerfile stage (modules come from the cached GOMODCACHE volume at /go/pkg/mod), and runtime behavior of server-route-test/server-test is unchanged.
- [x] #2 The server test image is built with a buildx builder using the GitHub Actions cache backend (cache-from/cache-to type=gha, mode=max) and tagged with the Compose-computed name; the migrator image is cached as a docker save/load tarball via actions/cache keyed on server/Dockerfile and migrations (buildx gha never restored RUN-step results), so up/run skip redundant builds.
- [x] #3 The Go build cache (GOCACHE volume) is persisted across CI runs via actions/cache, restored into the external volume before tests and saved after.
- [x] #4 The modified workflow is validated (e.g., actionlint or equivalent syntax check) and the full workflow passes on a pushed branch.
- [x] #5 The measured job duration on a run with a warm cache is recorded in the task and is meaningfully lower than the ~4m9s baseline.
- [x] #6 The E2E step's `npm ci` is resilient to slow registry edges: `~/.npm` is cached via actions/cache keyed on frontend/package-lock.json, restoring the same cache pattern as the Go build cache.
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

created: 2026-09-04 08:56
---
CI measurements (PR #17, meetwhen/timeful, run 33853348576):
- Attempt 1 (npm cache cold): 8m56s. Server test image 29s (GHA cache hit), migrator 120s (cache miss, cause unverified), backend tests 4s (GOCACHE seeding: 96s -> 4s), E2E 5m23s (npm ci still slow that run).
- Rerun (fully warm): ~6m20s. npm cache restore 3s, E2E 119s (npm fix works, was 5m23s), backend tests 3s, DB services 21s (baseline 80s). BUT migrator 120s again and server test image 88s.
- Log analysis of the warm rerun: migrator's ~75.6s was `#16 exporting to GitHub Actions Cache DONE 75.6s`. type=gha mode=max re-uploads every layer blob on every run; the goose-builder stage layer contains the full goose module dependency cache, and the testdeps stage layer contains the server module cache (~hundreds of MB). Cache I/O is now the bottleneck, not compilation.
- Next structural fix: slim the goose-builder layer (`RUN ... && go clean -cache -modcache`) so the migrator image cache is ~30MB; consider moving GOMODCACHE to a second cached volume mounted at /go/pkg/mod (requires compose.test.yaml + frontend/e2e/isolated-test-stack.ts mirror for the second external volume) instead of caching it inside the testdeps image.
- AC #5 (warm run meaningfully under 4m9s) is NOT yet satisfied; task stays In Progress.
---

author: agent
created: 2026-09-04 12:28
---
Round 3 (PR #17, commits e30a1482/a349b279/f4a36083) — structural fixes from the 08:56 diagnosis:
- server/Dockerfile: added `testbase` (module-cache-free base; CI test images stop there) and slimmed goose-builder with `go clean -cache -modcache` in the same RUN, so no cached image layer contains a module cache.
- Moved GOMODCACHE out of the image entirely: second external volume `timeful-test-go-mod-cache` mounted at /go/pkg/mod (GOMODCACHE=/go/pkg/mod) in both test services; seeded/saved via actions/cache keyed on server/go.sum (exact hits skip both copy-out and post-save); `frontend/e2e/isolated-test-stack.ts` creates it via the new `ensureExternalVolume` helper; AGENTS.md and docs/environments.md updated for the two-volume contract.
- Root-caused the 08:56 "transient" migrator cache miss: it was NOT transient. Across four measured runs, buildx gha mode=max restored FROM-stage records but never RUN-step results (the goose install re-ran at 40-55s every run even when the identical record was exported minutes earlier), and module-cache layers have unstable digests (fresh mtimes), so their blobs re-uploaded on every export (306MB layer: 34.7s upload + 17s docker load). Base-image blobs are stable and were deduped.
- Replaced buildx gha caching for the migrator with a docker save/load tarball cached by actions/cache keyed on server/Dockerfile + server/migrations/**; miss path: `compose build postgres-test-migrate` + `docker save` (~50s), hit path: `docker load` (2s measured).
- Gated the Save Go caches copy-out on primary-key != matched-key (exact hits would save nothing).
---

author: agent
created: 2026-09-04 12:28
---
CI measurements (PR #17, backend-quality job):
- Run 33870610187 (testbase + slimmed goose, buildx migrator, module cache cold): 3m42s wall. Server test image 15s (was 88s), migrator 56s (go install re-ran; export 6s vs 75.6s), backend tests 15s incl. 850MB module download into the empty volume, E2E 31s (was 119s), module cache saved 270MB compressed in 3s.
- Run 33871199548 (warm, gate commit): migrator goose-builder RUN re-executed again (54.9s) despite the identical record exported minutes earlier — confirmed RUN records never restore from gha (4/4 observations); step 68s. Job 3m56s wall.
- Run 33871887624 attempt 1 (tarball commit, migrator cache miss): migrator compose build 45s + docker save 2s; job 3m24s wall.
- Run 33871887624 attempt 2 (fully warm): migrator restore 1s + `docker load` (step 2s, was 120s at baseline and 120s/68s in rounds 1-2); backend tests 6s; seed volumes 21s; E2E 84s (run-to-run variance 26-84s, nix/compose noise inside the step); Save Go caches <1s (gate skips copy-out on exact hits); job wall 3m29s, active steps 3m18s.
- AC #5: warm run 3m18s active vs ~4m9s baseline (~20% lower), and the cold first run (3m42s) is also under baseline. npm cache validated across four runs (restore ~3s every time).

Local validation before pushing: compose config, actionlint 1.7.12, compose build of both targets, `go test ./... -count=1` all 14 packages twice (repeat run 9.1s with the module cache volume), docker save/load tarball round-trip plus an idempotent migrator run from the loaded image, the exact CI E2E spec (1 passed, 39.3s), frontend lint/typecheck/fmt:check/build/test:unit (991 tests), npm run format:markdown:check. Local test stack left running per convention with both external volumes populated.
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Backend CI (backend-quality) went from ~4m9s baseline (and ~6m20s warm after round 1) to 3m18s active / 3m29s wall on a fully warm run (PR #17, run 33871887624 attempt 2), with the cold first run at 3m42s.\n\nChanges:\n- server/Dockerfile: new `testbase` stage (golang base + WORKDIR only) as the build target for the CI test services, so no module download cache is baked into the cached image; `testdeps` (go mod download) remains solely as the builder-chain base; goose-builder now cleans module and build caches inside the same RUN so its cached layer holds only the ~30MB binary.\n- compose.test.yaml: server-test/server-route-test target `testbase`, set GOMODCACHE=/go/pkg/mod, and mount the new external `timeful-test-go-mod-cache` volume alongside the existing build-cache volume.\n- .github/workflows/backend-ci.yml: server test image built via buildx + gha cache (zero blob uploads now that no RUN layer exists); migrator image cached as a docker save/load tarball via actions/cache keyed on server/Dockerfile + server/migrations/** (miss: compose build ~45-55s + save; hit: docker load 2s) after four runs proved buildx gha never restores RUN-step results; GOMODCACHE seeded/saved via actions/cache keyed on server/go.sum; cache copy-out gated on key mismatch; comments updated.\n- frontend/e2e/isolated-test-stack.ts: ensureExternalVolume creates both external cache volumes.\n- AGENTS.md and docs/environments.md document the two-volume contract (create/reset both).\n\nValidation: actionlint; compose config; local compose build of both targets; go test ./... -count=1 (14 packages, twice); tarball save/load round-trip plus an idempotent migrator run from the loaded image; the exact CI E2E spec locally (1 passed, 39.3s); frontend lint/typecheck/fmt:check/build/test:unit (991 tests); npm run format:markdown:check. Four CI runs measured on PR #17; runtime behavior unchanged (-count=1 stays, E2E stack ownership preserved).
<!-- SECTION:FINAL_SUMMARY:END -->
