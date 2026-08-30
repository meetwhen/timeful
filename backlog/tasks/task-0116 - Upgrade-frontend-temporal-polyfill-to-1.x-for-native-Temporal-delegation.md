---
id: TASK-0116
title: Upgrade frontend temporal-polyfill to 1.x for native Temporal delegation
status: To Do
assignee: []
created_date: '2026-08-30 14:48'
labels:
  - frontend
  - dependencies
  - temporal
dependencies: []
references:
  - >-
    https://github.com/fullcalendar/temporal-polyfill/releases/tag/temporal-polyfill%401.0.1
  - >-
    https://github.com/fullcalendar/temporal-polyfill/blob/main/polyfill/README.md
  - 'https://github.com/tc39/proposal-temporal'
  - 'https://www.npmjs.com/package/temporal-polyfill'
documentation:
  - docs/design/architecture/adr/ADR-004.md
  - docs/design/architecture/adr/ADR-002.md
  - docs/design/architecture/adr/ADR-006.md
  - frontend/AGENTS.md
  - AGENTS.md
modified_files:
  - frontend/package.json
  - frontend/package-lock.json
  - docs/design/architecture/adr/ADR-004.md
  - frontend/AGENTS.md
priority: medium
type: chore
ordinal: 123000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Outcome

Upgrade the frontend Temporal runtime from `temporal-polyfill@^0.3.2` to the latest 1.x release (1.0.4 as of Aug 2026; re-verify latest 1.x at execution time). This moves the app to current Stage 4 Temporal spec conformance and makes the runtime automatically use engine-native `Temporal` where it exists, while the package's JS polyfill keeps browsers without native Temporal (notably Safari/WebKit as of Aug 2026) working. User value: spec-accurate time behavior, bugfixes in exactly this app's risk areas (future-offset `ZonedDateTime.from`, `Duration.round`/`total`), and no extra cost on modern browsers.

## Researched background (Aug 2026; re-verify facts at execution)

- Current state: `temporal-polyfill@^0.3.2` in `frontend/package.json`, imported as a ponyfill (`import { Temporal } from "temporal-polyfill"` / `import type { Temporal }`) in roughly 100 files across `frontend/src/` and `frontend/e2e/`. Value-semantics helpers are centralized in `frontend/src/utils/temporalPrimitives.ts` (ADR-006).
- `temporal-polyfill@1.0.x` (repo: fullcalendar/temporal-polyfill) keeps the same ponyfill entrypoint and, new in 1.0, delegates to native Temporal when available; 0.3 never did. Min+gzip 19.5 kB (23.4 kB for `/full/`).
- Native Temporal is Stage 4 and shipped in Firefox 139 (2025-05-27), Chrome 144 (2026-01-13), and Node 26 (2026-05-05). Safari/JavaScriptCore had NOT shipped it as of Aug 2026. Local Node is v26.7.0, so after the upgrade vitest unit tests and Playwright Firefox e2e run on native Temporal; only real WebKit browsers exercise the JS implementation in production.

## Confirmed decisions and constraints

- Package choice is decided: stay on `temporal-polyfill`. `@js-temporal/polyfill@0.5.1` was evaluated and rejected (52.1 kB min+gzip vs 19.5 kB, spec conformance dated Mar 2025, no native delegation). Do not switch packages in this task.
- Import sites must not change. The ponyfill entrypoint `import { Temporal } from "temporal-polyfill"` remains a supported 1.x entrypoint. Do not adopt `temporal-polyfill/global` (global types/`lib` changes are out of scope), the `/fns` tree-shakeable API, or a central re-export module in this task.
- 1.0 breaking changes and why they are acceptable here: ESM-only (project is Vite/ESM); BigInt required (Chrome 67+, Firefox 68+, Safari 14+, Node 16+; browser target is modern evergreen only — no legacy support matrix to honor); non-ISO calendar systems moved to `/full/` entrypoints (no calendar usage exists in `frontend/src/`, confirmed by search); `global.min.js` removed from the NPM package (not used here).
- TypeScript: `frontend/tsconfig.app.json` uses target ES2022 with `lib: ["ES2022", "DOM", "DOM.Iterable"]`, which has no built-in Temporal types; package-provided types keep working via the existing `import type` sites. No tsconfig changes are expected; only adjust if typecheck proves otherwise, without changing import sites.
- Known 1.0 behavior deltas vs 0.3 land in this repo's known Temporal risk areas and may legitimately change test outcomes: future-offset `Temporal.ZonedDateTime.from` calculations (e.g. Europe/Berlin 2044, Brazil wall times near offset transitions) and `Duration.round`/`total`/`from` fixes. These are spec-alignment fixes, not regressions; per ADR-004, runtime verification is mandatory because many Temporal regressions are runtime-only.
- Documentation updates belong in this task, not a follow-up.

## Explicitly out of scope (deferred)

- WebKit e2e smoke coverage of the JS-polyfill path (consider as a separate follow-up if polyfill-path confidence is wanted).
- Dropping the dependency entirely once Safari ships native Temporal.
- Adoption of the `temporal-polyfill/fns` tree-shakeable API or bundle-size work beyond the upgrade itself.
- Server-side changes; this task is frontend-only.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 frontend/package.json declares temporal-polyfill at the latest 1.x release available at execution time (^1.0.4 or newer 1.x), and frontend/package-lock.json resolves that 1.x version
- [ ] #2 Frontend and e2e import sites are unchanged: all continue to use the temporal-polyfill ponyfill entrypoint (import { Temporal } from "temporal-polyfill" and import type), with no migration to temporal-polyfill/global, /full/, /implementation, or /fns entrypoints
- [ ] #3 Frontend required checks pass from frontend/: npm run lint, npm run typecheck, npm run build, and npm run test:unit
- [ ] #4 npm run test:e2e -- --project=firefox-desktop passes from frontend/ with Playwright owning the isolated test stack (mongo-test, postgres-test, server-test on 3003, Vite on 4174), never the development API on 3002
- [ ] #5 Any runtime behavior difference vs temporal-polyfill 0.3 observed in unit or e2e tests is either (a) explained by a documented 1.0 conformance or bugfix change from the release notes and accepted with a note in the task, or (b) fixed with a regression test added before completion, following ADR-004's test rules
- [ ] #6 ADR-004 is updated to record the temporal-polyfill 1.x runtime model: one Temporal runtime model with native engine implementation where available (Firefox 139+, Chrome 144+, Node 26+) and the JS polyfill implementation as fallback (notably Safari/WebKit until it ships native Temporal), including the updated_date and the rationale that @js-temporal/polyfill was evaluated and rejected (52.1 kB vs 19.5 kB min+gzip, older spec, no native delegation)
- [ ] #7 frontend/AGENTS.md remains accurate for the upgraded dependency (the "Temporal via temporal-polyfill" line); update its wording if it would mislead a future agent about the 1.x runtime model
- [ ] #8 The task's final summary records the resolved temporal-polyfill version, the check results, and any accepted behavior changes vs 0.3
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
