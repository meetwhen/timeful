# Inbound mail for support@timeful.app via Stalwart (Docker)

Receive-only mail server for the timeful.app domain. Sending stays with the
listmonk relay; this stack never sends outbound mail from the domain.

## Topology

- `stalwart` container publishes ports **25** (inbound SMTP) and **993**
  (IMAPS) directly. IMAP clients connect to `mail.timeful.app:993`.
- The container HTTP listener (`8080`) is bound to `127.0.0.1` on the VM
  only. Caddy fronts `https://mail.timeful.app` (admin UI, JMAP, autoconfig).
  `STALWART_PUBLIC_URL` makes discovery documents publish that public URL.
- TLS certificates come from Caddy's Let's Encrypt issuance for
  `mail.timeful.app` (Caddy owns ACME on 443). The certificate and key are
  copied into `./certs` and imported into Stalwart for SMTP STARTTLS and
  IMAPS; see `scripts/sync-certs.sh`.
- Storage: RocksDB on the `stalwart-data` volume (fine for a single inbox).

## Prerequisites

- VM with Docker, a static public IP, and TCP **25**, **993**, **443**, **80**
  reachable from the internet. Some providers block inbound 25 by default
  (check the firewall; open a ticket with the provider if needed).
- Caddy running on the VM so `https://mail.timeful.app` terminates there.
  This repo's Caddy auto-imports the new site file via `import sites/*.caddy`.

## 1. DNS (manual, at the timeful.app provider)

| Type | Name | Value | Purpose |
|---|---|---|---|
| A | `mail.timeful.app` | VM public IP | server identity |
| MX | `timeful.app` | `10 mail.timeful.app.` | receive mail |
| PTR | reverse zone (VM provider) | `mail.timeful.app` | inbound credibility |
| CAA | `timeful.app` | `0 issue "letsencrypt.org"` | allow Let's Encrypt certs |
| SPF | `timeful.app` | leave as-is (relay's infra) | do NOT list the VM |
| DMARC | `_dmarc.timeful.app` | keep current policy; tighten to `p=reject` only after relay DKIM/SPF alignment is verified | |
| MTA-STS (optional) | `_mta-sts.timeful.app` + `mta-sts.timeful.app` CNAME | see Stalwart docs | only add once TLS works |

Wait for propagation (`dig MX timeful.app`, `dig A mail.timeful.app`).

## 2. Deploy

On the VM (e.g. `/srv/mail`):

```
cp -r infra/support-mail /srv/mail
cd /srv/mail
cp .env.example .env
# edit .env: STALWART_RECOVERY_ADMIN, STALWART_URL/PASSWORD later
docker compose up -d
docker compose logs stalwart 2>&1 | grep -A8 'bootstrap mode'
```

Env vars for Caddy (from `.env.edge`):
`CADDY_MAIL_DOMAIN=mail.timeful.app` and
`CADDY_MAIL_UPSTREAM=http://stalwart:8080` (Stalwart joins the `timeful-edge`
network). Reload Caddy and verify `https://mail.timeful.app` serves the
admin UI (cert first issued on the first request).

## 3. Setup wizard

Open `http://127.0.0.1:8080/admin` on the VM (or SSH tunnel) and sign in
with the bootstrap credentials. Answers:

1. Server hostname `mail.timeful.app`; domain `timeful.app`; **disable**
   automatic TLS certificates (Caddy owns ACME on 443); disable "generate
   DKIM keys" since the relay signs DKIM (keep disabled unless you plan to
   import the keys into the relay).
2. Storage: keep the defaults (RocksDB everywhere).
3. Directory: *Internal Directory*.
4. Logging: **Console** (Docker captures stdout).
5. DNS: *Manual DNS Server Management*.

The final screen prints the permanent administrator credential
(`admin@timeful.app` + password) - save it, then `docker compose restart
stalwart`.

## 4. Load Caddy's certificate into Stalwart

```
/srv/mail/scripts/sync-certs.sh   # copies cert to ./certs + reloads
```

In the WebUI (`https://mail.timeful.app/admin`): Settings -> TLS ->
Certificates -> New certificate: certificate chain file
`/certs/mail.timeful.app.crt`, private key file `/certs/mail.timeful.app.key`.
Set it as the default certificate. No restart needed - reload happens via
the reload action.

## 5. Mailboxes

Management -> Accounts: create `support@timeful.app` (mailbox). Optional
`postmaster@timeful.app` alias. IMAP client settings:

- host `mail.timeful.app`, port **993** (SSL/TLS)
- login `support@timeful.app`

## 6. Certificate renewal

Caddy renews automatically; Stalwart must re-read the files:

```
30 4 * * * /srv/mail/scripts/sync-certs.sh >> /srv/mail/sync-certs.log 2>&1
```

(Less frequent is fine; certs live ~90 days. Alternatively trigger
`ReloadTlsCertificates` from the WebUI Actions panel, or restart the
container.)

## 7. Hardening (per Stalwart security docs)

- Remove `STALWART_RECOVERY_ADMIN` from `.env` and `docker compose up -d`
  once the wizard administrator works.
- In the WebUI, disable listeners that are not used: plain IMAP 143, POP3
  110/995, submission 465/587, ManageSieve 4190, and any HTTP listener
  besides 8080 (which Caddy needs as upstream).
- Keep `STALWART_ADMIN_BIND_HOST=127.0.0.1` so the admin UI never binds
  publicly.

## 8. Backup

The mailbox data is the `timeful-support-mail-stalwart-data` volume plus the
`/etc/stalwart` config volume. Back up both, e.g.:

```
docker run --rm -v timeful-support-mail-stalwart-data:/data -v /srv/backups:/backup \
    alpine tar czf /backup/stalwart-data-$(date +%F).tar.gz /data
```

## Verification

- `dig MX timeful.app` -> mail host; `dig A mail.timeful.app` -> VM IP
- external: `swaks --to support@timeful.app --from someone@example.org
  --server mail.timeful.app` (or send from any webmail provider)
- check `docker compose logs stalwart` for delivery, then read the message
  via IMAP
- admin UI and discovery: `curl https://mail.timeful.app/.well-known/jmap`

## Notes / gotchas

- Spam/abuse: the mailbox will attract spam; Stalwart's built-in spam
  filter handles it (see Docs -> Spamfilter rules in the WebUI).
- Don't add the VM to the domain SPF record - it doesn't send.
- Self-served mail to `support@timeful.app` from the listmon relay is
  delivered to this server; ensure the relay's DKIM selector stays
  published, otherwise; DMARC verification may fail for those messages and
  (with a strict policy) they can be rejected.
- Docs: https://stalw.art - Installation, DNS, Security, Reverse proxy,
  Actions (`ReloadTlsCertificates`).