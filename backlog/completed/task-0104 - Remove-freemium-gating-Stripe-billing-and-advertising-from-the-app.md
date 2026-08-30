---
id: TASK-0104
title: 'Remove freemium gating, Stripe billing, and advertising from the app'
status: Done
assignee:
  - opencode
created_date: '2026-08-29 11:40'
updated_date: '2026-08-29 21:21'
labels:
  - freemium
  - ads
  - removal
  - frontend
  - backend
  - docs
dependencies: []
references:
  - docs/requirements/AGENTS.md
  - docs/requirements/README.md
  - docs/design/architecture/adr/ADR-003.md
  - docs/requirements/functional/fr/FR-070.md
  - docs/requirements/functional/fr/FR-071.md
  - docs/requirements/functional/fr/FR-072.md
  - docs/terminology/glossary.md
  - docs/environments.md
  - BACKLOG_WORKFLOW.md
type: chore
ordinal: 110000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
All functionality must be ungated: no freemium limits, no upgrade/paywall surfaces, no advertising, no Stripe billing. Remove the entire freemium and ads apparatus from frontend, server, configuration, and the requirements corpus.

Owner decisions recorded in task notes are binding scope:

- Full Stripe removal on the backend, including routes, model fields, dependency, and STRIPE_* env vars; regenerate swagger and frontend API types.
- Remove all ad plumbing including the advertising cookie-consent category and the Google AdSense mention; keep analytics consent and GTM.
- Keep product-monitoring analytics (PostHog, GTM, numEventsCreated counter, /analytics/monthly-active-event-creators*, /analytics/user/:email, /analytics/scanned-poster). Remove only upgrade-funnel and ad-click analytics.
- Delete requirements records FR-070, FR-071, FR-072, ADR-003, and the glossary Freemium Mode entry per docs/requirements/AGENTS.md conventions.

Stale stripeCustomerId/isPremium fields in existing Mongo user documents are acceptable; no data migration is required since the fields are simply ignored after removal.

The detailed file-by-file implementation plan is in the task plan section; session context is in the notes section.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 No freemium gating remains in frontend or server code: no paywall, upgrade dialog, free-event limit, premium checks, owner-premium lookups, or VITE_ENABLE_FREEMIUM.
- [x] #2 No advertising surfaces or third-party ad plumbing remain: Publift fuse script and zones, AdShield blob, Primis video ad, ads.txt, ad components and assets, and the advertising cookie-consent category (analytics consent stays).
- [x] #3 The Stripe billing stack is fully removed server-side (routes/stripe.go, is-premium endpoint, model fields, stripe-go dependency, STRIPE_* env vars) and swagger plus frontend/src/types/api.ts are regenerated.
- [x] #4 Upgrade-funnel analytics are removed (/analytics/upgrade-dialog-viewed, upgrade-user, downgrade-user, upgrade_* and ad-click PostHog events) while monitoring analytics remain (monthly-active-event-creators*, user/:email, scanned-poster, PostHog, GTM, numEventsCreated).
- [x] #5 Requirements records FR-070, FR-071, FR-072, ADR-003, the glossary Freemium Mode entry, and their README index rows are deleted following docs/requirements/AGENTS.md and terminology conventions.
- [x] #6 Required checks pass: frontend npm run lint, typecheck, build, test:unit; server go build, go vet, and the scoped route suite via the compose.test.yaml overlay.
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
# Implementation plan

## 1. Frontend — delete files

- `frontend/src/utils/freemium.ts` and `frontend/src/utils/freemium.test.ts`
- `frontend/src/utils/services/PricingService.ts`
- `frontend/src/composables/pricing/upgradeDialogModels.ts`
- `frontend/src/composables/pricing/useUpgradeDialogState.ts`
- `frontend/src/components/pricing/UpgradeDialog.vue` and `UpgradeDialog.test.ts`
- `frontend/src/components/pricing/AlreadyDonatedDialog.vue`
- `frontend/src/components/event/Advertisement.vue` (unused cross-promo), `CarbonAd.vue`, `PubliftAd.vue`, `asyncPubliftAd.ts`
- `frontend/public/adshield-bootstrap.js`, `frontend/public/ads.txt`
- `frontend/src/assets/ads/` (tomotime.png, tomotime_mobile.png)

## 2. Frontend — edits

- `frontend/src/constants.ts`: remove `numFreeEvents`, `upgradeDialogTypes` + `UpgradeDialogType`, `authTypes.UPGRADE`
- `frontend/src/stores/main.ts`: remove freemium/premium imports; `enablePaywall`, `setEnablePaywall`, `isPremiumUser`, `viewerHasPremiumAccess`; upgrade-dialog state and actions (`upgradeDialogVisible/Type/Data`, `setUpgradeDialog*`, `showUpgradeDialog`, `hideUpgradeDialog`); the free-events gate inside `createNew`
- `frontend/src/utils/index.ts`: drop the `./freemium` re-export
- `frontend/src/App.vue`: remove UpgradeDialog mount + `handleUpgradeDialogInput`, the `isPremiumUser` Premium badge, related store refs
- `frontend/src/components/home/Dashboard.vue`: remove "free events created" usage row + `openUpgradeDialog`
- `frontend/src/components/EventType.vue`: remove usage row + `openUpgradeDialog` + freemium imports
- `frontend/src/views/Event.vue`: remove `showAds` computed, all four AsyncPubliftAd blocks (desktop side rails `meet_vrec_lhs`/`meet_vrec_rhs`, in-content `meet_incontent_md`, mobile + fixed desktop bottom ads), `videoAdContainer` + `loadVideoAd` + Primis script, `initFusetag` + its call, the 115px/125px ad spacer divs and `bottom: '115px'` sticky offsets, `ownerIsPremium`/`ownerPremiumChecked` props passed to ScheduleOverlap, `checkOwnerPremium` call + `ownerPremiumChecked` watch, commented-out CarbonAd, `freemiumEnabled` import
- `frontend/src/composables/event/useEventLoader.ts`: remove `checkOwnerPremium`, `ownerIsPremium`, `ownerPremiumChecked` (drops the `/users/{id}/is-premium` call)
- `frontend/src/components/schedule_overlap/ScheduleOverlap.vue`: remove `showAds` computed, `ownerIsPremium`/`ownerPremiumChecked` props + defaults, view-model exposure
- `frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue`: remove phone `meet_incontent` ad block + AsyncPubliftAd import
- `frontend/src/components/schedule_overlap/scheduleOverlapTestUtils.ts`: remove `AsyncPubliftAd` stub
- `frontend/src/views/Settings.vue`: remove Billing section + `openBillingPortal`
- `frontend/src/views/Auth.vue`: remove `authTypes.UPGRADE` case + `UpgradeParams` interface
- `frontend/src/views/SignIn.vue`: remove `upgradeRedirect` handling in `signIn` + `handlePostAuthRedirect`, `UpgradeParams` interface
- `frontend/src/views/StripeRedirect.vue`: delete file; `frontend/src/router/index.ts`: remove `/stripe-redirect` route and the `fusetag.destroySticky` beforeEach block
- `frontend/index.html`: remove Publift fuse.js script block, AdShield block, and the Quantcast `qc-cmp2` local-overlay style block (keep GTM)
- Cookie consent: `frontend/src/utils/cookie_utils.ts` remove `advertising` preference, `hasAdvertisingConsent`, `ad_consent` in `initializeGTMConsent` (keep `analytics` + `analytics_consent`); `frontend/src/components/CookieConsent.vue` and `CookieSettings.vue` remove the Advertising category and "Services used: Google AdSense" mention
- `frontend/src/types/transport.ts`: drop `RawUser` `stripeCustomerId`/`isPremium` extras
- Unit tests to update: `App.test.ts` (UpgradeDialog stubs/store fields), `EventType.test.ts`, `components/home/Dashboard.test.ts`, `views/Settings.test.ts`, `utils/cookie_utils.test.ts`, `CookieConsent.test.ts`, `CookieSettings.test.ts`; delete `UpgradeDialog.test.ts`, `freemium.test.ts`

## 3. Server

- Delete `server/routes/stripe.go`
- `server/routes/users.go`: remove `/:userId/is-premium` route + `getIsUserPremium`
- `server/routes/analytics.go`: remove `upgradeDialogViewed` handler + route, `upgradeUser`, `downgradeUser` handlers + routes (keep `scanned-poster`, `monthly-active-event-creators*`, `user/:email`)
- `server/routes/events.go`: remove `user.StripeCustomerId = nil` sanitize line (~2261); keep both `numEventsCreated` `$inc` updates
- `server/db/users.go`: remove `GetUserByStripeCustomerId`
- `server/models/user.go`: remove `StripeCustomerId`, `IsPremium` fields
- `server/main.go`: remove stripe import, `routes.InitStripe(apiRouter)` call, `stripe.Key` env read
- `server/go.mod`/`go.sum`: drop `github.com/stripe/stripe-go/v82`
- Regenerate swagger from `server/`: `go run github.com/swaggo/swag/cmd/swag@v1.16.6 init --parseDependency`, then from `frontend/`: `npm run gen:api` (regenerates `frontend/src/types/api.ts`)

## 4. Env/config

- `compose.yaml`: remove `VITE_ENABLE_FREEMIUM` (frontend service) and all `STRIPE_*` vars (server service)
- Env files: `.env.development` + `.env.development.example`, `.env.test` + `.env.test.example`, `.env.staging.example`, `.env.production.example`: remove `VITE_ENABLE_FREEMIUM` and the STRIPE_* block
- `docs/environments.md`: remove `VITE_ENABLE_FREEMIUM` and `STRIPE_*` entries
- `.graphifyignore`: remove the `frontend/public/ads.txt` line

## 5. Docs (read docs/requirements/AGENTS.md + docs/terminology/README.md first)

- Delete `docs/requirements/functional/fr/FR-070.md`, `FR-071.md`, `FR-072.md`
- Delete `docs/design/architecture/adr/ADR-003.md`
- `docs/requirements/README.md`: remove the FR-070/071/072 index rows (and any other freemium references)
- `docs/terminology/glossary.md`: remove the Freemium Mode entry and its TOC link; check for dangling cross-references to the term elsewhere in docs
- `docs/design/README.md`: remove the ADR-003 reference

## 6. Verification

- Frontend: `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit`
- Server: `go build ./...`, `go vet ./...`; Mongo route tests via the compose overlay: `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml up -d mongo-test` then `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test`
- Postman-style sanity: freemium greps over frontend/src and server return zero (excluding backlog history and generated docs until regen)
- `graphify update .`

# Remaining work — session 2 checklist (frontend source edits DONE; start here)

## A. Frontend unit tests (do first, before checks)
- App.test.ts: remove upgradeDialogVisible/isPremiumUser from storeToRefs mock; remove upgradeDialogVisible/isPremiumUser refs and hideUpgradeDialog from useMainStore mock; remove all `UpgradeDialog: true` stub entries.
- EventType.test.ts: drop numFreeEvents/upgradeDialogTypes imports, showUpgradeDialogMock + its store mock entry, enablePaywall mock field; delete/update usage-row tests (~lines 89-121).
- Dashboard.test.ts: remove showUpgradeDialogMock (definition + store mock entry + reset).
- Settings.test.ts: remove stripeCustomerId user fixtures and any Billing-section assertions.
- cookie_utils.test.ts / CookieConsent.test.ts / CookieSettings.test.ts: remove advertising preferences/expectations.
- Event.test.ts: remove checkOwnerPremiumMock, ownerIsPremium/ownerPremiumChecked/checkOwnerPremium from useEventLoader mock, all `AsyncPubliftAd: true` stubs (~40), and the checkOwnerPremiumMock expectation at ~3575.
- ScheduleOverlapSidebar.test.ts: remove showAds mock fields and AsyncPubliftAd stubs; drop/adapt ad-rendering tests (~355-393).

## B. Server
- Delete server/routes/stripe.go.
- routes/users.go: remove /:userId/is-premium route + getIsUserPremium handler.
- routes/analytics.go: remove upgradeDialogViewed handler+route, upgradeUser, downgradeUser (KEEP scanned-poster, monthly-active-event-creators*, user/:email).
- routes/events.go: remove `user.StripeCustomerId = nil` sanitize (~2261); KEEP both numEventsCreated $inc updates (events.go ~419, ~2205).
- db/users.go: remove GetUserByStripeCustomerId.
- models/user.go: remove StripeCustomerId, IsPremium fields.
- main.go: remove stripe import, routes.InitStripe(apiRouter), stripe.Key env read.
- go.mod/go.sum: drop github.com/stripe/stripe-go/v82 (run go mod tidy).
- Regenerate: from server/ run `go run github.com/swaggo/swag/cmd/swag@v1.16.6 init --parseDependency`, then from frontend/ `npm run gen:api` (regenerates src/types/api.ts; do not hand-edit).

## C. Env/config + docs
- compose.yaml: remove VITE_ENABLE_FREEMIUM (frontend service) and all STRIPE_* vars (server service).
- Remove VITE_ENABLE_FREEMIUM + STRIPE_* from .env.development, .env.development.example, .env.test, .env.test.example, .env.staging.example, .env.production.example.
- docs/environments.md: remove VITE_ENABLE_FREEMIUM and STRIPE_* entries.
- .graphifyignore: remove frontend/public/ads.txt line.
- Docs (read docs/requirements/AGENTS.md + docs/terminology/README.md first): delete FR-070.md, FR-071.md, FR-072.md, ADR-003.md; remove their rows from docs/requirements/README.md; remove glossary Freemium Mode entry + TOC link and fix dangling cross-references; remove ADR-003 reference from docs/design/README.md.
- Run `npm run format:markdown` on changed Markdown files (repo root tooling).

## D. Verification (all required)
- frontend/: npm run lint, typecheck, build, test:unit.
- server/: go build ./..., go vet ./...
- Mongo route tests via overlay: `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml up -d mongo-test` then `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test`.
- Sanity greps: freemium|isPremium|is-premium|paywall|UpgradeDialog|numFreeEvents|stripe|adshield|fusetag|publift|primis|adsense|ad_consent|hasAdvertisingConsent|VITE_ENABLE_FREEMIUM → zero in frontend/src + server (backlog history excluded).
- `graphify update .`
- Then finalize per BACKLOG_WORKFLOW.md: verify each AC with evidence, check AC/DoD boxes, write final summary, set status Done.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Session decisions (owner-approved 2026-08-29, binding scope)

- Backend scope: full Stripe removal (not frontend-only). Stale `stripeCustomerId`/`isPremium` fields in existing Mongo documents are acceptable; do not write a migration.
- Ad-adjacent plumbing: remove everything, including the advertising cookie-consent category and the Google AdSense mention. Keep the analytics consent category and GTM.
- Docs: delete FR-070/071/072 and ADR-003 records outright (not mark-as-removed); remove the glossary Freemium Mode entry and README index rows.
- Analytics: keep monitoring (PostHog, GTM, `numEventsCreated` counter, `/analytics/monthly-active-event-creators*`, `/analytics/user/:email`, `/analytics/scanned-poster`). Remove only upgrade-funnel and ad-click analytics. This implies removing `upgrade-user`/`downgrade-user` (premium management, meaningless without premium state) and `/analytics/upgrade-dialog-viewed`.

## Implementation-session guidance

- Implementation is planned for a later session; this task was created from a planning-only session that mapped every reference site via graphify plus greps (grep terms used: freemium, isPremium, premium, paywall, upgrade, stripe, adshield, fusetag, publift, primis, adsense, numFreeEvents, VITE_ENABLE_FREEMIUM).
- Read BACKLOG_WORKFLOW.md before implementation: search before creating, mark In Progress and assign when starting, keep plan/notes current, record a final summary.
- Regeneration chain matters: server Swag annotations change → from `server/` run `go run github.com/swaggo/swag/cmd/swag@v1.16.6 init --parseDependency` → from `frontend/` run `npm run gen:api` (regenerates `frontend/src/types/api.ts`; do not hand-edit it).
- `frontend/src/types/transport.ts` `RawUser` extends the generated `Schemas["models.User"]`; after regen the manual `stripeCustomerId`/`isPremium` extras can be dropped.
- Gating is currently frontend-only: the server never enforces a creation limit, it only `$inc`s `numEventsCreated` (events.go:419, events.go:2205). Keep those increments.
- Advertisement.vue (`EventAdvertisement`, toomtime cross-promo) is already unreferenced; CarbonAd.vue is only referenced from a commented-out line in Event.vue.
- e2e has no freemium/ads coverage today, so no e2e expectations change; unit tests to update are listed in the plan.
- Markdown edits must follow sentence-per-line and glossary-canonicalization rules (see AGENTS.md Documentation Authoring section).

## Verification recap for the executing session

- Frontend: `npm run lint`, `npm run typecheck`, `npm run build`, `npm run test:unit` from `frontend/`.
- Server: `go build ./...`, `go vet ./...` from `server/`; Mongo-backed route tests via `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml up -d mongo-test` then `... run --rm server-route-test`.
- Run `graphify update .` after code changes.

## Session 1 progress (2026-08-29, implementation started)

Status: frontend source edits are COMPLETE; frontend unit tests NOT yet updated; server, env/config, docs, swagger regen, and all verification still pending.

Frontend deletions done via git rm: freemium.ts + test, PricingService.ts, upgradeDialogModels.ts, useUpgradeDialogState.ts, UpgradeDialog.vue + test, AlreadyDonatedDialog.vue, Advertisement.vue, CarbonAd.vue, PubliftAd.vue, asyncPubliftAd.ts, public/adshield-bootstrap.js, public/ads.txt, src/assets/ads/* (tomotime images), views/StripeRedirect.vue.

Frontend edits done: constants.ts (authTypes.UPGRADE, upgradeDialogTypes/UpgradeDialogType, numFreeEvents removed); stores/main.ts (freemium/premium imports, enablePaywall, isPremiumUser, viewerHasPremiumAccess, upgrade-dialog state/actions, free-events gate in createNew removed; pricingPageConversion kept as out of scope); utils/index.ts (./freemium re-export dropped); App.vue (UpgradeDialog mount, handleUpgradeDialogInput, Premium badge, store refs removed); Dashboard.vue + EventType.vue (usage rows, openUpgradeDialog, freemium imports, unused storeToRefs import removed).

Event.vue edits done: video ad container div, all 4 AsyncPubliftAd blocks (meet_vrec_lhs/rhs, meet_incontent_md, mobile+fixed desktop bottom), commented CarbonAd, showAds computed, videoAdContainer + adsBootstrapped refs, loadVideoAd, initFusetag, ownerPremiumChecked watch, checkOwnerPremium call, 115px/125px spacers and bottom:'115px' sticky style, :owner-is-premium/:owner-premium-checked props on ScheduleOverlap, freemiumEnabled + AsyncPubliftAd + viewerHasPremiumAccess removed.

More frontend edits done: router/index.ts (/stripe-redirect route + fusetag.destroySticky beforeEach block removed); env.d.ts (Window fusetag/enableStickyFooter and VITE_ENABLE_FREEMIUM removed — addition beyond plan, required by AC#1); useEventLoader.ts (checkOwnerPremium/ownerIsPremium/ownerPremiumChecked/getRealOwnerId import removed); ScheduleOverlap.vue (owner-premium props/defaults, showAds, freemiumEnabled, showAds in derived view-model removed); useScheduleOverlapViewModels.ts (showAds from options type, derived Pick, sidebar VM; mobile overlay bottomOffset now constant "4rem"); scheduleOverlapViewModelContracts.ts (showAds removed); ScheduleOverlapSidebar.vue (meet_incontent ad block + import removed); scheduleOverlapTestUtils.ts (AsyncPubliftAd stub + showAds removed); Settings.vue (Billing section, openBillingPortal, now-unused get import removed); Auth.vue (upgradeParams field, UpgradeParams interface, UPGRADE case fully removed); SignIn.vue (upgradeRedirect computed, UPGRADE branch in signIn, UpgradeParams, checkout branch in handlePostAuthRedirect removed).

Cookie consent done: cookie_utils.ts rewritten without advertising preference/hasAdvertisingConsent/ad_consent (analytics + GTM kept); CookieConsent.vue acceptAll no longer sets advertising; CookieSettings.vue Advertising category + Google AdSense mention + advertising in acceptAll/rejectAll removed. transport.ts: RawUser is now plain Schemas["models.User"] (stripeCustomerId/isPremium extras dropped). index.html: fuse.js script, Quantcast qc-cmp2 local-overlay style, and AdShield block removed; GTM kept.

Discovery: freemium/ads plumbing also reaches useScheduleOverlapViewModels.ts and scheduleOverlapViewModelContracts.ts (showAds threading incl. mobile overlay bottomOffset 115px offset) — updated in this session though not itemized in the original plan. Same category as plan items; no scope change.

Gotcha for next session: while editing useEventLoader.ts a temporary duplicate checkOwnerPremium was created and then fully removed; current file verified clean (grep checkOwnerPremium/ownerIsPremium returns nothing).

Test files needing updates (next session): App.test.ts (pinia storeToRefs mock lists upgradeDialogVisible/isPremiumUser; useMainStore mock has upgradeDialogVisible ref, isPremiumUser ref, hideUpgradeDialog; UpgradeDialog: true stubs at ~5 mount sites); EventType.test.ts (numFreeEvents/upgradeDialogTypes imports, showUpgradeDialogMock, enablePaywall mock, usage-row tests ~89-121); components/home/Dashboard.test.ts (showUpgradeDialogMock at 12/15/51/178); views/Settings.test.ts (stripeCustomerId fixtures at 27/92, Billing section tests); utils/cookie_utils.test.ts (advertising expectations at 23-53); CookieConsent.test.ts (advertising at 67/90); CookieSettings.test.ts (advertising at 57/81/92); views/Event.test.ts (checkOwnerPremiumMock hoisted 124/141, useEventLoader mock fields ownerIsPremium/ownerPremiumChecked/checkOwnerPremium at 196-206, ~40 'AsyncPubliftAd: true' stubs lines ~638-3554, line 3575 expects checkOwnerPremiumMock toHaveBeenCalled — delete that expectation/test); ScheduleOverlapSidebar.test.ts (showAds + AsyncPubliftAd stub tests ~355-393; drop or adapt the ad-rendering assertions).

IMPORTANT worktree note: repo has unrelated uncommitted TASK-0107 changes (SlideToggle.vue/.test.ts, AvailabilityTypeToggle.vue/.test.ts + new AvailabilityTypeToggle.test.ts) — do not touch or revert them; frontend checks will include them.

## Session 2 progress (2026-08-30, implementation completed)

Frontend unit tests updated and verified: App.test.ts (storeToRefs + store mocks stripped of upgradeDialogVisible/isPremiumUser/hideUpgradeDialog, all UpgradeDialog stubs removed), EventType.test.ts rewritten (usage-row and upgrade tests deleted; also removed the now-dead unused `mainStore = useMainStore()` leftover in EventType.vue), Dashboard.test.ts (showUpgradeDialogMock + viewerHasPremiumAccess removed), Settings.test.ts (stripeCustomerId fixtures removed), cookie_utils.test.ts / CookieConsent.test.ts / CookieSettings.test.ts (advertising expectations removed; CookieSettings now asserts 2 checkboxes), Event.test.ts (51 AsyncPubliftAd stubs, checkOwnerPremiumMock, ownerIsPremium/ownerPremiumChecked, viewerHasPremiumAccess, and the checkOwnerPremium expectation removed), ScheduleOverlapSidebar.test.ts (ad-wrapper test deleted, sticky test adapted without showAds).

Server removals done: routes/stripe.go deleted; users.go is-premium route + handler removed; analytics.go upgrade-dialog-viewed/upgrade-user/downgrade-user routes and handlers removed (scanned-poster, monthly-active-event-creators*, user/:email kept); events.go StripeCustomerId sanitize line removed (both numEventsCreated $inc kept at events.go:419 and :2205); db/users.go GetUserByStripeCustomerId removed; models/user.go StripeCustomerId/IsPremium fields removed; main.go stripe import, InitStripe call, and stripe.Key env read removed; go mod tidy dropped github.com/stripe/stripe-go/v82.

Regeneration done: swag v1.16.6 init --parseDependency from server/ then npm run gen:api from frontend/ regenerated frontend/src/types/api.ts (163 deletions, 1 benign insertion). Grep confirms zero stripe/premium in server/docs and api.ts.

Env/config done: compose.yaml VITE_ENABLE_FREEMIUM + all STRIPE_* removed; all six env files cleaned; docs/environments.md entries removed (also removed the stale 'Stripe redirects' phrase in the APP_BASE_URL sentence); .graphifyignore ads.txt line removed.

Docs done: FR-070/071/072 and docs/design ADR-003 deleted; requirements README index rows removed; glossary Freemium Mode entry, Product section, and TOC lines removed; the Platform Visitor entry's dangling FR-072 authoritative context redirected to FR-079 (consistent with its sibling visitor entries); docs/design/README.md ADR-003 row removed. npm run format:markdown run on changed Markdown. Note: frontend/adr/ no longer exists (migrated to docs/design in 5ae68018); the graphify semantic cache still holds stale freemium nodes from removed files and one VITE_ENABLE_FREEMIUM node from docs/environments.md; purging them requires a semantic re-extraction with an API key set, and all source files are clean.

Verification evidence: frontend lint, typecheck, build, and test:unit all pass (137 files, 952 tests); server go build and go vet pass; Mongo route suite via compose.test.yaml overlay returns ok timeful/server/routes; sanity greps for freemium|isPremium|is-premium|paywall|UpgradeDialog|numFreeEvents|stripe|adshield|fusetag|publift|primis|adsense|ad_consent|hasAdvertisingConsent|VITE_ENABLE_FREEMIUM and for advertising/CarbonAd/Advertisement/ads.txt/upgrade-dialog/upgrade-user/downgrade-user return zero over frontend/src + server (server/docs included in generated-docs check).

## Out-of-scope observations for the owner (not changed)

- frontend/src/components/TeamsNotReadyDialog.vue still says 'Upgrade Required' and mentions upgrading for the Timeful organization plan (cal.com book-a-call flow). This is an organization-plan sales CTA, not the freemium apparatus; it was not in the binding plan, in the same category as pricingPageConversion which session 1 explicitly kept out of scope.
- slackbot.MONETIZATION and SLACK_MONETIZATION_WEBHOOK_URL are now unused by any handler after the Stripe and upgrade-funnel removals; retained because the binding plan did not include slackbot plumbing or the env var.
- E2E has no freemium/ads coverage today (per session 1 notes), so no e2e expectations changed; the task's verification plan scoped checks to unit + route tests.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
# Remove freemium gating, Stripe billing, and advertising

## What changed and why

All functionality is now ungated: the freemium limit/paywall stack, Stripe billing, advertising plumbing, and their analytics are fully removed from frontend, server, configuration, and the requirements corpus, per the owner-approved binding scope.

### Frontend
- Deleted freemium/premium/ad files: `utils/freemium.ts` (+test), `utils/services/PricingService.ts`, `composables/pricing/upgradeDialogModels.ts` + `useUpgradeDialogState.ts`, `components/pricing/UpgradeDialog.vue` (+test) + `AlreadyDonatedDialog.vue`, `event/Advertisement.vue`, `CarbonAd.vue`, `PubliftAd.vue`, `asyncPubliftAd.ts`, `public/adshield-bootstrap.js`, `public/ads.txt`, `assets/ads/*`, `views/StripeRedirect.vue`.
- Removed freemium/premium state and actions from `stores/main.ts` and `constants.ts` (`numFreeEvents`, `upgradeDialogTypes`, `authTypes.UPGRADE`), the UpgradeDialog mount and Premium badge from `App.vue`, usage rows from Dashboard/EventType, the free-events gate in `createNew`, and all ad blocks/scripts (Publift fuse + zones, AdShield, Primis video ad, 115px/125px spacers), owner-premium lookups via `/users/{id}/is-premium`, the `/stripe-redirect` route, Billing section in Settings, and `authTypes.UPGRADE` auth flows.
- Cookie consent now has only necessary + analytics categories (GTM and analytics consent kept); the advertising category, `hasAdvertisingConsent`, `ad_consent`, and the Google AdSense mention are gone; `transport.ts` `RawUser` dropped `stripeCustomerId`/`isPremium`.
- Unit tests updated across App, EventType, Dashboard, Settings, cookie utils/consent/settings, Event, and ScheduleOverlapSidebar; `main.ts`-adjacent dead code (`mainStore` in EventType.vue) cleaned up.

### Server
- Deleted `routes/stripe.go`; removed the `/:userId/is-premium` endpoint, `upgrade-dialog-viewed`, `upgrade-user`, and `downgrade-user` analytics routes/handlers, `GetUserByStripeCustomerId`, `StripeCustomerId`/`IsPremium` model fields, the stripe import + `InitStripe` call + `stripe.Key` read in `main.go`, and the `stripeCustomerId` sanitize line in `events.go`.
- Kept monitoring analytics: `scanned-poster`, `monthly-active-event-creators*`, `user/:email`, PostHog/GTM, and both `numEventsCreated` `$inc` updates (events.go:419, events.go:2205).
- Dropped `github.com/stripe/stripe-go/v82` via `go mod tidy`; regenerated swagger (`swag v1.16.6 init --parseDependency`) and frontend API types (`npm run gen:api`).

### Config and docs
- Removed `VITE_ENABLE_FREEMIUM` and all `STRIPE_*` vars from `compose.yaml` and all six env files; updated `docs/environments.md` and `.graphifyignore`.
- Deleted `FR-070`, `FR-071`, `FR-072`, and `docs/design/architecture/adr/ADR-003.md` with their index rows; removed the glossary Freemium Mode entry (and its Product section/TOC lines), redirecting the Platform Visitor entry's authoritative context to FR-079; removed the ADR-003 row from `docs/design/README.md`.

## Verification
- `frontend`: `npm run lint`, `npm run typecheck`, `npm run build` pass; `npm run test:unit` passes (137 files, 952 tests).
- `server`: `go build ./...` and `go vet ./...` pass; Mongo route suite via the `compose.test.yaml` overlay returns `ok timeful/server/routes` (test state retained).
- Sanity greps for freemium/isPremium/is-premium/paywall/UpgradeDialog/numFreeEvents/stripe/adshield/fusetag/publift/primis/adsense/ad_consent/hasAdvertisingConsent/VITE_ENABLE_FREEMIUM and advertising/ads.txt/upgrade-funnel analytics return zero over `frontend/src` + `server`; generated swagger and `src/types/api.ts` are clean.
- `graphify update .` run after code changes.

## Risks / follow-ups
- `TeamsNotReadyDialog.vue` still shows the organization-plan "Upgrade Required" book-a-call CTA (sales flow, outside the binding scope, same category as the intentionally kept `pricingPageConversion`).
- `slackbot.MONETIZATION` / `SLACK_MONETIZATION_WEBHOOK_URL` are now unused by handlers and could be removed in a follow-up.
- The graphify semantic cache retains stale freemium nodes for deleted files; a semantic re-extraction with an API key would purge them.
<!-- SECTION:FINAL_SUMMARY:END -->
