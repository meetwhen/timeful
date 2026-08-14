# PostgreSQL Staging Rollout

## Preconditions

- Revoke the Microsoft Graph credential previously committed in the legacy test.
- Merge the PostgreSQL-aware application and run the CI gates.
- Populate every required PostgreSQL role credential and URI in `.env.staging`.
- Keep `POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=false` for the first deploy.

## Deploy With Creation Disabled

```sh
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml config --quiet
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml up -d --build
docker compose --project-name timeful-staging --env-file .env.staging -f compose.yaml -f compose.staging.yaml ps
curl -fsS https://staging.timeful.fun/api/health/live
curl -fsS https://staging.timeful.fun/api/health
```

Confirm `postgres-migrate` completed successfully, Mongo event read/write behavior remains intact, and both readiness endpoints return successfully.

## Enable And Smoke Test

Set `POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true` in `.env.staging`, redeploy the server, then create one anonymous timed poll and one dates-only poll. Confirm each has `p_` long and short IDs and exercise guest response mutation, selected schedule save/clear, plugin `set-slots`/`get-slots`, and a signed-in response. Confirm PostgreSQL account responses remain absent from the dashboard and `/api/user/events`.

## Rollback

Set `POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=false` and redeploy the server. This returns new creation to MongoDB without moving or disabling existing PostgreSQL events. Do not roll back the additive SQL schema.

PostgreSQL backup and recovery automation remain deferred in phase one.
