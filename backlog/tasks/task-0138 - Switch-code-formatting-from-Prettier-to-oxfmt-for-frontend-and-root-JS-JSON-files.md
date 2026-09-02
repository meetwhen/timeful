---
id: TASK-0138
title: >-
  Switch code formatting from Prettier to oxfmt for frontend and root JS/JSON
  files
status: Done
assignee:
  - opencode
created_date: '2026-09-02 11:57'
updated_date: '2026-09-02 12:27'
labels:
  - tooling
dependencies: []
references:
  - frontend/.prettierrc
  - frontend/.prettierignore
  - frontend/tailwind.config.cjs
  - .github/workflows/frontend-ci.yml
  - .github/workflows/markdown-ci.yml
documentation:
  - 'https://oxc.rs/docs/guide/usage/formatter.html'
  - 'https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html'
  - 'https://oxc.rs/docs/guide/usage/formatter/sorting.html'
  - 'https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html'
modified_files:
  - frontend/package.json
  - frontend/.prettierrc
  - frontend/.prettierignore
  - package.json
  - .github/workflows/frontend-ci.yml
  - .github/workflows/markdown-ci.yml
  - AGENTS.md
  - frontend/AGENTS.md
priority: medium
type: chore
ordinal: 151300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Why

Frontend code formatting is currently ad-hoc: Prettier is configured in `frontend/.prettierrc` (semi: false, 2-space, prettier-plugin-tailwindcss) but there is no format script, no CI gate, and 340 files currently fail `prettier --check`. The repo already uses the Oxc stack for linting (oxlint + oxlint-tsgolint + eslint). Switching formatting to oxfmt aligns formatting with the lint toolchain, replaces prettier-plugin-tailwindcss with built-in Tailwind class sorting, gives roughly 30x faster formatting, and for the first time establishes an enforced canonical format.

## Confirmed decisions

- Tool: oxfmt standalone, not vite-plus (vite-plus is at 0.3.0 and pins older oxlint/oxfmt versions; it can be adopted later as the unified toolchain).
- Scope: frontend (src, e2e, playwright.config.ts — the same scope as lint) plus root JS files under `scripts/` and `prettier/`; `.vue` files are included.
- Enforcement: add fmt:check to CI and apply a one-time reformat sweep now.
- Style parity with the Prettier era: semi false, tabWidth 2, useTabs false, printWidth 80, double quotes (oxfmt defaults are double quotes but printWidth 100 — override printWidth to 80).
- Tailwind class sorting: enable via the oxfmt Tailwind sorting option and point it explicitly at `frontend/tailwind.config.cjs`; oxfmt auto-detection only looks for `tailwind.config.js`. Do not silently drop class sorting if the .cjs config is not read — stop and ask.
- Import sorting stays disabled (candidate follow-up); package.json key sorting may remain at oxfmt defaults.
- The root Markdown pipeline stays entirely on Prettier: root `.prettierrc`, prettier-plugin-sentences-per-line, the `prettier/markdown/` wrapper plugin and its tests, `scripts/markdown.mjs`, the format:markdown / format:markdown:check / lint:markdown scripts, and markdown-ci.yml must remain unchanged and green.
- oxfmt must not format any Markdown; scope oxfmt invocations to code paths so the docs Markdown corpus stays governed by the sentences-per-line pipeline.
- Generated `frontend/src/types/api.ts` (output of `npm run gen:api`) is excluded from formatting checks so regeneration never breaks fmt:check.
- Keep `eslint-config-prettier` in frontend devDependencies (it still disables conflicting ESLint stylistic rules).

## Current-state pointers

- `frontend/.prettierrc` and `frontend/.prettierignore` are the files to migrate/replace; `.prettierignore` currently lists dist, node_modules, .vscode, index.html.
- Frontend lint scope lives in the frontend/package.json scripts (oxlint + eslint over `src e2e playwright.config.ts`, ignoring `e2e/inspect/**`).
- CI: `.github/workflows/frontend-ci.yml` (add fmt:check after lint) and `.github/workflows/markdown-ci.yml` (add root fmt:check; it already triggers on scripts/ and prettier/** paths).
- Docs to update: root AGENTS.md Required Checks and frontend/AGENTS.md Required Checks.
- oxfmt reads `.prettierignore` for compatibility, but prefer ignorePatterns in the oxfmt config; nested configs are supported, so root and frontend each own their config.
- The one-time sweep is expected to touch roughly 340 frontend files; keep sweep changes separable from tooling config changes for review.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 frontend/package.json has oxfmt as a devDependency and no longer has prettier or prettier-plugin-tailwindcss; frontend/.prettierrc and frontend/.prettierignore are removed with their ignore entries migrated into the oxfmt config.
- [x] #2 Frontend npm scripts fmt and fmt:check exist and run oxfmt over src, e2e, and playwright.config.ts (same scope as lint); e2e/inspect/** is excluded from formatting.
- [x] #3 Frontend oxfmt config preserves current style: semi false, tabWidth 2, useTabs false, printWidth 80; .vue files are formatted as part of the scope.
- [x] #4 Tailwind class sorting is enabled in the frontend oxfmt config with an explicit reference to frontend/tailwind.config.cjs; spot-checked components produce the same class order prettier-plugin-tailwindcss produced. If oxfmt cannot read the .cjs config, stop and ask rather than dropping class sorting.
- [x] #5 src/types/api.ts is excluded from formatting checks, and running npm run gen:api followed by npm run fmt:check passes.
- [x] #6 Root package.json gains oxfmt as a devDependency plus fmt and fmt:check scripts covering the JS files under scripts/ and prettier/ using a root-level oxfmt config; root fmt:check passes and no Markdown file is formatted by oxfmt.
- [x] #7 Root Markdown pipeline is unchanged and green: npm run format:markdown:check, npm run lint:markdown, npm run test:markdown-format, and npm run test:markdown-rules all pass from the repo root.
- [x] #8 One-time reformat sweep is applied in both scopes; a second npm run fmt run is idempotent with zero changes; sweep changes remain separable from tooling config changes for review.
- [x] #9 CI enforces formatting: frontend-ci.yml runs the frontend fmt:check step and markdown-ci.yml runs the root fmt:check step.
- [x] #10 Root AGENTS.md and frontend/AGENTS.md Required Checks include fmt:check; root .prettierrc and root Prettier devDependencies remain, scoped to Markdown only.
- [x] #11 Frontend checks pass after migration: npm run lint, npm run typecheck, npm run build, npm run test:unit, and npm run fmt:check from frontend/; e2e is not required for this formatting-only change.
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
## Implementation Plan (oxfmt 0.66.0, standalone)

1. **Install oxfmt@latest** as a devDependency in `frontend/` and at repo root (`npm install -D oxfmt`). Prettier stays in root devDependencies for the Markdown pipeline; prettier + prettier-plugin-tailwindcss are removed from frontend only after Tailwind parity is verified.
2. **Frontend config** `frontend/.oxfmtrc.json`: `$schema`, `semi: false`, `tabWidth: 2`, `useTabs: false`, `printWidth: 80` (oxfmt default is 100), `sortTailwindcss: { config: "tailwind.config.cjs" }` (path resolves relative to the config file), `ignorePatterns` migrating `.prettierignore` entries (`dist`, `node_modules`, `.vscode`, `index.html`) plus `e2e/inspect/**` and `src/types/api.ts`.
3. **Tailwind parity gate (stop-and-ask point)**: while prettier-plugin-tailwindcss is still installed, run identical unsorted class inputs through `prettier --plugin prettier-plugin-tailwindcss --stdin-filepath` and `oxfmt --stdin-filepath` and diff the sorted output, including prefix-sensitive classes (`tw-*`, custom colors/screens from tailwind.config.cjs). If oxfmt cannot read the `.cjs` config or orders differ, stop and ask instead of dropping class sorting.
4. **Frontend scripts**: `fmt: oxfmt src e2e playwright.config.ts`, `fmt:check: oxfmt --check src e2e playwright.config.ts` (same scope as lint; e2e/inspect/** excluded via ignorePatterns). Remove `frontend/.prettierrc` and `frontend/.prettierignore`; drop `prettier` and `prettier-plugin-tailwindcss` from frontend devDependencies; keep `eslint-config-prettier` (still disables conflicting ESLint stylistic rules, and eslint.config.ts:123 references it).
5. **Root config** `.oxfmtrc.json`: same base style but `singleQuote: true` to match the existing root JS style in scripts/markdown.mjs and prettier/markdown/*.js (minimizes sweep churn; root .prettierrc does not govern these JS files today), `ignorePatterns: ["**/*.md"]` so oxfmt never touches the Markdown corpus (prettier/markdown/README.md lives inside the scope dirs). Root scripts: `fmt: oxfmt scripts prettier`, `fmt:check: oxfmt --check scripts prettier` (scope: only JS under scripts/ and prettier/).
6. **Nested config sanity**: nearest-config-wins means frontend files resolve to frontend/.oxfmtrc.json and root scope files to the root .oxfmtrc.json; verify with a check run in each scope.
7. **One-time sweep**: run `npm run fmt` in frontend (expected ~340 files) and root (2 files: scripts/markdown.mjs, prettier/markdown/sentences-per-line.js + its test). Re-run to confirm idempotency (zero changes on second run). Keep sweep edits separable from tooling config edits for review (no commits; report file lists).
8. **CI**: frontend-ci.yml adds a `fmt:check` step after lint; markdown-ci.yml adds a root `fmt:check` step and `.oxfmtrc.json` to trigger paths.
9. **Docs**: root AGENTS.md and frontend/AGENTS.md Required Checks gain `fmt:check` (one sentence per source line); run `npm run format:markdown` on changed Markdown.
10. **Verification against ACs**: frontend lint/typecheck/build/test:unit + fmt:check; `npm run gen:api` then `npm run fmt:check` passes (api.ts ignored); second `fmt` run idempotent; root fmt:check passes with zero Markdown files formatted; root markdown pipeline green (format:markdown:check, lint:markdown, test:markdown-format, test:markdown-rules).
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Session 2: Fixed the 5 sweep-broken unit tests (6 assertions total — ToolRow's 'Shown in' markup assertion also needed updating beyond the 5 listed in session 1). Updated ?raw substring assertions to post-sweep canonical snippets: ScheduleOverlap.breakpoints.test.ts (grid content class order, pager-button classes with tw-border-outline-neutral), NewEvent.test.ts (2x reflowed :class object), ToolRow.test.ts (compact class order + Shown-in label reflow). All test edits themselves pass oxfmt --check.

Session 2 verification: lint 0 errors (2 pre-existing warnings), fmt:check 377 files PASS, typecheck exit 0, build exit 0, test:unit 986/986 PASS. gen:api exit 0 then fmt:check PASS (AC #5). Root markdown pipeline all green (AC #7). format:markdown run on changed AGENTS.md files, already format-clean (DoD #4). Root fmt:check PASS (3 JS files only — no Markdown touched by oxfmt). Idempotency re-verified in both scopes: second fmt run produces zero changes (AC #8). graphify graph updated.

Review note: sweep changes and tooling config changes are all staged together but uncommitted; split into separate commits (tooling/CI/docs/test-assertions vs reformat sweep) for reviewability.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: opencode
created: 2026-09-02 12:17
---
## Progress (session 1)

**Done**
- oxfmt@0.66.0 installed as devDependency in frontend/ and at repo root.
- frontend/.oxfmtrc.json created: semi false, tabWidth 2, useTabs false, printWidth 80, sortTailwindcss.config = tailwind.config.cjs, ignorePatterns = dist, node_modules, .vscode, index.html, e2e/inspect/**, src/types/api.ts.
- Tailwind parity gate PASSED: oxfmt output byte-identical to prettier-plugin-tailwindcss on a synthetic prefix/custom-screen/custom-color input and on real components Footer.vue and RespondentsList.vue; the .cjs config is read correctly (tw- prefix, publift-* screens honored).
- frontend prettier + prettier-plugin-tailwindcss removed; frontend/.prettierrc and frontend/.prettierignore deleted; frontend fmt / fmt:check scripts added (oxfmt src e2e playwright.config.ts).
- Root .oxfmtrc.json created (semi false, singleQuote true to match existing root JS style, printWidth 80, ignorePatterns ["**/*.md"] so oxfmt never touches Markdown); root fmt / fmt:check scripts added (oxfmt scripts prettier).
- Root sweep: 3 files reformatted (scripts/markdown.mjs, prettier/markdown/sentences-per-line.js, prettier/markdown/sentences-per-line.test.mjs); second run idempotent.
- Frontend sweep: 377 files processed, 319 modified; second run idempotent; src/types/api.ts and e2e/inspect untouched.
- CI updated: frontend-ci.yml gains fmt:check step after lint; markdown-ci.yml gains root fmt:check step and .oxfmtrc.json in trigger paths (both pull_request and push).
- Root AGENTS.md and frontend/AGENTS.md Required Checks updated with fmt:check plus oxfmt/Prettier scope note.
- Frontend lint PASS (0 errors, 2 pre-existing vue/one-component-per-file warnings), typecheck PASS, build PASS, fmt:check PASS; root fmt:check PASS.

**Blocker found: 5 unit test failures caused by the sweep (clean tree passes)**
All 5 are `?raw` template-source substring assertions that encoded pre-sweep class order / line wrapping. Semantics intact; assertions must be updated to the new canonical formatting:
- ScheduleOverlap.breakpoints.test.ts (2): expects 'class="schedule-overlap-time-grid__content tw-grow tw-min-w-0"' but sweep sorted it to 'tw-min-w-0 tw-grow' (ScheduleOverlapTimeGrid.vue:54); also the pager-button substring tw-h-8 tw-w-8 tw-min-w-8 sm:tw-h-[36px] sm:tw-w-[36px] sm:tw-min-w-[36px] got reordered (tw-border-outline-neutral moved within the class list, ScheduleOverlapTimeGrid.vue:15 / ScheduleOverlapDaysOnlyGrid.vue:6,17).
- NewEvent.test.ts (2, lines ~913 and ~923): expects "'time-range-select-item--active': item.raw === selectedDateOption" on one source line but oxfmt reflowed the multiline :class object in NewEvent.vue:150.
- ToolRow.test.ts (1, lines ~99-104): expects 'tw-w-full tw-flex-col tw-items-start tw-justify-start tw-gap-0 tw-pt-14 tw-pb-0' but sweep reordered to '... tw-gap-0 tw-pb-0 tw-pt-14' (ToolRow.vue:17).

**Remaining work**
1. Update the 5 raw-source assertions above to the post-sweep canonical snippets (keep test intent).
2. Re-run npm run test:unit, then the full frontend check set (lint, typecheck, build, test:unit, fmt:check).
3. Verify AC #5: run npm run gen:api then npm run fmt:check in frontend (api.ts must be ignored).
4. Verify AC #7: from repo root run format:markdown:check, lint:markdown, test:markdown-format, test:markdown-rules.
5. Run npm run format:markdown on the changed AGENTS.md files (DoD #4), then re-run format:markdown:check.
6. Finalize per task-finalization guide: verify each AC with evidence, write finalSummary, mark Done. Sweep changes remain uncommitted/separable (no commits made in this session).
---
<!-- COMMENTS:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Switched code formatting from Prettier to oxfmt for the frontend and root JS files, establishing an enforced canonical format aligned with the existing Oxc lint stack.

## Tooling changes

- Added `oxfmt@^0.66.0` as devDependency in `frontend/` and at repo root; removed `prettier` and `prettier-plugin-tailwindcss` from frontend devDependencies (kept `eslint-config-prettier`, still disabling conflicting ESLint stylistic rules).
- Deleted `frontend/.prettierrc` and `frontend/.prettierignore`; their ignore entries migrated into `frontend/.oxfmtrc.json` (`dist`, `node_modules`, `.vscode`, `index.html`, plus `e2e/inspect/**` and `src/types/api.ts`).
- `frontend/.oxfmtrc.json`: semi false, tabWidth 2, useTabs false, printWidth 80 (oxfmt default 100), `sortTailwindcss.config = "tailwind.config.cjs"` so Tailwind class sorting reads the .cjs config explicitly.
- Root `.oxfmtrc.json`: same base style plus `singleQuote: true` (matches existing root JS style) and `ignorePatterns: ["**/*.md"]` so oxfmt never touches the Markdown corpus, which stays on the Prettier sentences-per-line pipeline.
- Scripts: frontend `fmt`/`fmt:check` run `oxfmt src e2e playwright.config.ts` (same scope as lint); root `fmt`/`fmt:check` run `oxfmt scripts prettier` (3 JS files).

## One-time reformat sweep (uncommitted; separable from tooling changes at commit time)

- Frontend: 377 files processed, 319 modified (class-list reordering, line rewrapping). Root: 3 files (scripts/markdown.mjs, prettier/markdown/sentences-per-line.js + its test).
- Second `fmt` run is idempotent in both scopes (verified this session: `git diff` empty, subsequent `fmt:check` exit 0).

## Test assertion updates (6 assertions across 5 failing tests)

The sweep reformatted `?raw` template-source substrings asserted by unit tests; assertions were updated to the new canonical snippets with test intent preserved: ScheduleOverlap.breakpoints.test.ts (content-class order, pager-button class order incl. `tw-border-outline-neutral` position), NewEvent.test.ts (2x reflowed `:class` object), ToolRow.test.ts (compact-classes order + reflowed "Shown in" label markup).

## CI and docs

- frontend-ci.yml: `fmt:check` step after lint. markdown-ci.yml: root `fmt:check` step and `.oxfmtrc.json` added to pull_request/push trigger paths.
- Root AGENTS.md and frontend/AGENTS.md Required Checks include `fmt:check`; root `.prettierrc` and root Prettier devDependencies remain, scoped to Markdown only. Changed Markdown formatted via `npm run format:markdown`.

## Verification

- Frontend: lint (0 errors, 2 pre-existing vue/one-component-per-file warnings), fmt:check (377 files PASS), typecheck (exit 0), build (exit 0), test:unit (138 files, 986/986 pass).
- AC #5: `npm run gen:api` exit 0 then `npm run fmt:check` PASS (api.ts excluded).
- Root markdown pipeline green: format:markdown:check, lint:markdown, test:markdown-format, test:markdown-rules all exit 0.
- Tailwind parity gate passed earlier: oxfmt output byte-identical to prettier-plugin-tailwindcss on synthetic prefix/custom-screen/custom-color input and real components (Footer.vue, RespondentsList.vue).

## Risks / follow-ups

- Commits not yet made: recommend splitting into (1) tooling config + CI + docs + test-assertion updates and (2) the reformat sweep, for reviewability.
- Import sorting remains disabled (candidate follow-up); vite-plus unified toolchain adoption deferred until it matures.
<!-- SECTION:FINAL_SUMMARY:END -->
