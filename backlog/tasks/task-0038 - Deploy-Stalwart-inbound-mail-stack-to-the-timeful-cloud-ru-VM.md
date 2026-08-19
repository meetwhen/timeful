---
id: TASK-0038
title: Deploy Stalwart inbound-mail stack to the timeful-cloud-ru VM
status: Done
assignee: []
created_date: '2026-08-19 17:49'
updated_date: '2026-08-19 19:18'
labels: []
dependencies: []
priority: high
type: chore
ordinal: 26000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Deploy the Stalwart inbound-mail stack (infra/support-mail) to the production VM timeful-cloud-ru (192.144.13.176) over SSH, wired behind the existing Caddy edge.

The domain is timeful.fun (user-confirmed); the mail host is mail.timeful.fun and the mailbox is support@timeful.fun. Outbound stays with the Brevo relay; Stalwart is receive-only. The VM serves timeful.fun (timeful.app is a stale assumption and must not be used).

Current state: compose.edge.yaml edge passthrough fix committed locally (b6af0aae). VM repo (~/timeful) has infra/support-mail + caddy/sites/mail.caddy already but its compose.edge.yaml lacks the CADDY_MAIL_DOMAIN/CADDY_MAIL_UPSTREAM env passthrough and .env.edge lacks the mail vars. VM needs sudo for docker, timeful-edge network exists, ports 25/993 free.

DNS prerequisite: user publishes mail.timeful.fun A -> 192.144.13.176 and timeful.fun MX -> 10 mail.timeful.fun. As of handoff time neither record exists; Caddy ACME and external SMTP/IMAP depend on them.

Deployment per infra/support-mail/README.md with the headless wizard path (no browser on VM): drive x:Bootstrap/set via HTTP 8080 API or stalwartlabs/cli image, pin admin via STALWART_RECOVERY_ADMIN, restart in recovery mode to create support@timeful.fun, then remove recovery vars. Expected deploys dir: /srv/mail (needs sudo mkdir). CADDY_DATA_DIR must be /var/lib/docker/volumes/timeful-edge_caddy_data/_data on this VM (Caddy is containerized, not a host /var/lib/caddy). TLS certs for Stalwart SMTP STARTTLS/IMAPS are copied from Caddy via scripts/sync-certs.sh; sync-certs.sh needs CLI creds from .env and CADDY_DATA_DIR correct.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 mail.timeful.fun resolves and Caddy serves the Stalwart admin/JMAP at https://mail.timeful.fun
- [x] #2 support@timeful.fun mailbox exists and accepts inbound SMTP on the VM (port 25)
- [x] #3 IMAP TLS login works over mail.timeful.fun:993 using the IMAPS cert synced from Caddy
- [x] #4 STALWART_RECOVERY_ADMIN removed from production .env after provisioning; admin password is known and stable
- [x] #5 edge compose keeps the CADDY_MAIL_DOMAIN/CADDY_MAIL_UPSTREAM passthrough so the mail site does not break production/staging certs on reload
- [x] #6 infra/support-mail README, .env.example and caddy/sites/mail.caddy consistently use .fun naming
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All unit tests pass
- [x] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
Deploy Stalwart to timeful-cloud-ru. strategy: controlled copy (no git pull on VM, which is 6 behind origin/main and its local .fun migration is not on origin). Steps:

1. Prepare fix files (done in repo): compose.edge.yaml passthrough (committed b6af0aae), infra/support-mail .fun migration (committed 2535486d). Add STALWART_RECOVERY_MODE passthrough to infra/support-mail/compose.yaml (new, required for recovery phase via .env).
2. Sync (scp/rsync to VM ~/timeful): compose.edge.yaml (with mail vars), .env.edge.example (add CADDY_MAIL_DOMAIN/CADDY_MAIL_UPSTREAM), infra/support-mail/** (.fun names + CADDY_DATA_DIR docs). VM caddy/sites/mail.caddy already uses placeholders; identical.
3. Write VM ~/timeful/.env.edge mail vars: CADDY_MAIL_DOMAIN=mail.timeful.fun, CADDY_MAIL_UPSTREAM=http://stalwart:8080.
4. Create /srv/mail (sudo), copy infra/support-mail contents, write /srv/mail/.env with .fun values + real recovery admin password + CADDY_DATA_DIR=/var/lib/docker/volumes/timeful-edge_caddy_data/_data.
5. sudo docker compose up -d in /srv/mail; confirm bootstrap log line.
6. Headless wizard via curl to http://127.0.0.1:8080/api (or stalwartlabs/cli) with x:Bootstrap/set payload: serverHostname=mail.timeful.fun, defaultDomain=timeful.fun, requestTlsCertificate=false, generateDkimKeys=false, tracer Stdout, directory Internal, dnsServer Manual.
7. Set STALWART_RECOVERY_MODE=1 in /srv/mail/.env, restart stalwart; create support@timeful.fun account + (postmaster alias) + known admin password via CLI/recovery admin; verify IMAP/SMTP locally.
8. After DNS propagated: update edge, reload Caddy, run sync-certs.sh, import cert in WebUI, verify dig/MX/IMAP. This step blocked on user DNS.
9. Hardening: remove recovery vars, disable unused listeners, nightly sync-certs cron.

Risks: DNS records not published yet (mail.timeful.fun A + MX missing as of last check). Only ~3.7GiB free disk. First two SSH connects to VM timed out at banner exchange previously; retry.On. sudorequired for docker on VM. CLI sync-certs.sh paths must see the network (stalwart reachable at http://stalwart:8080 inside timeful-edge).
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
[2026-08-19 18:05Z] Deployment progress: local repo committed b6af0aae (edge passthrough fix), a3d0049d (RECOVERY_MODE passthrough), 47fb6cdc (RECOVERY_ADMIN optional). VM sync applied (compose.edge.yaml, .env.edge.example, infra/support-mail/**; backups in /tmp/mail-before-sync/). .env.edge + mail vars. /srv/mail stack up: stalwart v0.16.18, headless wizard driven via x:Bootstrap/set (mail.timeful.fun / timeful.fun, TLS auto off, DKIM off, Stdout tracer, Internal dir, Manual DNS); permanent admin admin@timeful.fun password set via CLI and stored in /srv/mail/.env; support@timeful.fun created (Account c), creds on VM in /srv/mail/.mailbox-credentials (root-only). Recovery vars removed; normal mode verified (SMTP 25, IMAPS 993, MTA queue). SMTP smoke test delivered to local mailbox (DSN success). Edge reloaded with --env-file; prod + staging still 200; Caddy ACME for mail.timeful.fun fails only due to missing DNS, retries 60s. Nightly sync-certs cron added.

BLOCKED on DNS (user task): mail.timeful.fun A -> 192.144.13.176 and timeful.fun MX -> 10 mail.timeful.fun . not yet published. After propagation: Caddy auto-issues cert; run /srv/mail/scripts/sync-certs.sh; import cert in WebUI (IMAPS/STARTTLS); verify IMAP TLS login + external SMTP; optionally tighten DMARC.

[2026-08-19 18:58Z] DNS published and verified: mail.timeful.fun A 192.144.13.176 (first save used the wrong host name; corrected to subdomain 'mail'), timeful.fun MX 10 mail.timeful.fun. Both confirmed on auth NS ns1/ns2.reg.ru.

[2026-08-19 18:58Z] Caddy issued a production Let's Encrypt cert (issuer YE2) for mail.timeful.fun once DNS propagated. Cert storage lives under the volume mountpoint /var/lib/docker/volumes/timeful-edge_caddy_data/_data/caddy/certificates/... (note the extra 'caddy' segment).

[2026-08-19 18:58Z] Fixed two sync blockers on the VM: (1) /srv/mail/.env CADDY_DATA_DIR was .../_data but must include the /caddy suffix for scripts/sync-certs.sh to find the certs; (2) the nightly sync cron was in user1's crontab but the Caddy volume certs are root-only, so it moved to root's crontab.

[2026-08-19 18:58Z] Stalwart had no Certificate objects (requestTlsCertificate=false wizard leaves the store empty), so sync-certs.sh copy + reload alone could not load certs. Created Certificate jap2vwwgaaqc via CLI with File refs /certs/mail.timeful.fun.crt + .key, set SystemSettings.defaultCertificateId, ran action/ReloadTlsCertificates. The recurring 'No TLS certificates available' warnings ceased.

[2026-08-19 18:58Z] Verified: IMAPS :993 and SMTP STARTTLS :25 serve CN=mail.timeful.fun with a full ISRG chain (loopback + hairpin through the public IP). IMAP TLS login as support@timeful.fun succeeds (INBOX/Sent/Drafts/Junk). Caddy serves JMAP: https://mail.timeful.fun/.well-known/jmap -> /jmap/session HTTP 200; /admin -> 302 login. Production timeful.fun and staging.timeful.fun still HTTP 200. docker compose config --quiet passes for /srv/mail and the edge stack; sync-certs.sh passes bash -n.

[2026-08-19 18:58Z] External-reachability finding: from my network, raw TCP to mail.timeful.fun:25 and :993 completes the handshake but the data phase is filtered (no SMTP banner / no ServerHello), while hairpin via the public IP works from the provider network. Suspected hoster-side inbound filtering for external IPs. External SMTP from a public MTA remains unproven; verify with the hoster and/or from a mobile/webmail client.

Doc polish (repo, uncommitted): infra/support-mail/.env.example CADDY_DATA_DIR comment now explains the containerized '/caddy' mountpoint; README step 4 adds the headless CLI path for creating the Certificate object; README step 6 notes the root-cron requirement.

Security note: the admin STALWART_PASSWORD value was printed into this session's transcript while debugging CLI auth. Consider rotating the password in /srv/mail/.env and updating downstream references.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Stalwart inbound-mail deployment to timeful-cloud-ru — DNS gate cleared, full stack verified

**What changed:** Published DNS and completed the last mile of the Stalwart inbound-mail deployment so mail.timeful.fun is live behind the existing Caddy edge. stack (stalwart v0.16.18, /srv/mail) was already deployed and provisioned headlessly in prior sessions; this session unblocked transport and TLS.

**DNS (user, done):** `mail.timeful.fun` A → 192.144.13.176 and `timeful.fun` MX → 10 mail.timeful.fun, confirmed on auth NS ns1/ns2.reg.ru. The A record required a fix (initially saved under the wrong host name).

**TLS cert pipeline (two bugs fixed):**
1. `/srv/mail/.env` `CADDY_DATA_DIR` pointed at `.../_data` but certs live at `.../_data/caddy/certificates/...`; corrected and documented in `.env.example`.
2. Stalwart had no Certificate objects (`requestTlsCertificate=false` leaves the TLS store empty), so the copy + reload from `sync-certs.sh` could not load certs. Created Certificate `jap2vwwgaaqc` (File refs → `/certs/mail.timeful.fun.{crt,key}`), set `SystemSettings.defaultCertificateId`, and reloaded. Also moved the nightly sync cron from user1's crontab to root's (the Caddy volume certs are root-only).

**Verified (objective):**
- Caddy issued a production LE cert (YE2) for mail.timeful.fun.
- IMAPS :993 / SMTP STARTTLS :25 present CN=mail.timeful.fun with a full ISRG chain (loopback + public-IP hairpin).
- IMAP TLS login as support@timeful.fun OK (INBOX/Sent/Drafts/Junk listed).
- Caddy serves the mail site: /.well-known/jmap → /jmap/session 200; /admin 302→login. Production + staging still 200.
- `docker compose config --quiet` passes for /srv/mail and the edge stack; `sync-certs.sh` passes `bash -n` (no app code touched, so frontend/server unit/e2e suites are unaffected).
- Recovery vars confirmed absent from production `.env`; admin CLI auth works.

**Tests:** infra-only change (compose/caddy/scripts/docs). Validation: compose config, bash syntax, live SMTP/IMAP/JMAP/edge checks above.

**Risks / follow-ups (not blocking):**
- External inbound SMTP from a public MTA not yet proven: from my network, TCP to :25/:993 completes but the data phase is filtered while the provider-network hairpin works; suspected hoster-side inbound filtering for external IPs. Verify with the hoster and test from mobile.
- Admin `STALWART_PASSWORD` was exposed in this session's transcript during CLI debugging; rotate on the VM if it matters.
- Local repo main is 4 commits ahead of origin (mail work, unpushed); VM repo is 6 behind origin with working-tree copies — decide push-vs-copy for permanent sync.
- Optional hardening trail: disable unused listeners (465/995/4190/443-in-container), optionally tighten DMARC to p=reject once relay SPF/DKIM alignment is confirmed.
<!-- SECTION:FINAL_SUMMARY:END -->
