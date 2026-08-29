---
id: TASK-0104
title: 'Remove freemium gating, Stripe billing, and advertising from the app'
status: To Do
assignee: []
created_date: '2026-08-29 11:40'
updated_date: '2026-08-29 11:40'
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
- [ ] #1 No freemium gating remains in frontend or server code: no paywall, upgrade dialog, free-event limit, premium checks, owner-premium lookups, or VITE_ENABLE_FREEMIUM.
- [ ] #2 No advertising surfaces or third-party ad plumbing remain: Publift fuse script and zones, AdShield blob, Primis video ad, ads.txt, ad components and assets, and the advertising cookie-consent category (analytics consent stays).
- [ ] #3 The Stripe billing stack is fully removed server-side (routes/stripe.go, is-premium endpoint, model fields, stripe-go dependency, STRIPE_* env vars) and swagger plus frontend/src/types/api.ts are regenerated.
- [ ] #4 Upgrade-funnel analytics are removed (/analytics/upgrade-dialog-viewed, upgrade-user, downgrade-user, upgrade_* and ad-click PostHog events) while monitoring analytics remain (monthly-active-event-creators*, user/:email, scanned-poster, PostHog, GTM, numEventsCreated).
- [ ] #5 Requirements records FR-070, FR-071, FR-072, ADR-003, the glossary Freemium Mode entry, and their README index rows are deleted following docs/requirements/AGENTS.md and terminology conventions.
- [ ] #6 Required checks pass: frontend npm run lint, typecheck, build, test:unit; server go build, go vet, and the scoped route suite via the compose.test.yaml overlay.
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
<!-- SECTION:NOTES:END -->
