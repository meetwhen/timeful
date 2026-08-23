---
id: TASK-0037
title: Enable POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED on staging and production
status: Done
assignee: []
created_date: '2026-08-19 14:56'
updated_date: '2026-08-19 15:03'
labels: []
milestone: Postgres rollout
dependencies: []
documentation:
  - docs/postgres-staging-rollout.md
  - DEPLOYMENT.md
priority: high
type: chore
ordinal: 25000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Enable anonymous event creation in PostgreSQL for both staging and production by setting POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true in the remote env files (/home/user1/timeful/.env.staging and /home/user1/timeful/.env.production on timeful-cloud-ru) and redeploying the respective server containers.

Context: The server routes new event creation to MongoDB unless postgresCreationEnabled returns true (server/routes/postgres_event_routes.go:687). That requires the flag true AND an anonymous (signed-in users go to Mongo regardless) timed or dates-only poll payload. Currently both environments have the flag false, so ALL events, anonymous or signed-in, are created in MongoDB (m_ public IDs). The phase-1 rollout doc is docs/postgres-staging-rollout.md; phase 2 (enable + smoke test) applies to staging first, then production.

Rollback: set the flag back to false and redeploy; existing PostgreSQL events remain intact (additive schema, no migration rollback).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Flag POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true set in remote .env.staging and container env of running staging server
- [x] #2 Staging smoke test: anonymous timed poll and dates-only poll both return bare 8-char Crockford shortIds (no prefix) and readable at /e/, guest response mutation works, existing m_ prefixed Mongo event still readable, signed-in creation still goes to Mongo
- [x] #3 Flag POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true set in remote .env.production and container env of running production server
- [x] #4 Production smoke test: anonymous event creation returns bare Crockford ID and readable at https://timeful.fun/e/{id}
- [x] #5 No MongoDB data loss; backups of both env files retained
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 15:00
---
Staging deployed with POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true (container env verified). Smoke passed: anonymous timed poll created bare Crockford ID R5SD23MY, ids endpoint confirms shortId==longId==R5SD23MY, GET /e readable, guest response POST returned credentials, schedule save/clear 200. Days-only creation and GET /responses needed full RFC3339 dates and timeMin/timeMax params respectively (script payload issues, not server bugs).
---

created: 2026-08-19 15:03
---
Production deployed with POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true (container env verified), health ok. Smoke passed: anonymous timed poll created bare Crockford MYQ2KMTE, days-only poll S9TY9YFD, ids shortId==longId, both readable with correct names, guest response POST returned credentials and GET responses shows the guest availability, schedule save/clear 200. Legacy Mongo event m_B6NXVASB still resolves (200). Backups retained: .env.staging.bak.before-postgres-enable-20260819, .env.production.bak.before-postgres-enable-20260819.
---
<!-- COMMENTS:END -->
