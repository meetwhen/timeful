#!/usr/bin/env bash
set -euo pipefail

# Copies the Caddy-issued certificate for the mail hostname into ./certs
# (mounted into the stalwart container at /certs) and asks Stalwart to
# reload TLS certificates. Caddy renews roughly every 60 days, so this is
# run nightly from cron (see README).

cd "$(dirname "$0")/.."

if [[ ! -f .env ]]; then
    echo "error: .env not found next to compose.yaml" >&2
    exit 1
fi
set -a
# shellcheck disable=SC1091
source .env
set +a

caddy_cert_dir="${CADDY_DATA_DIR}/certificates/acme-v02.api.letsencrypt.org-directory/${MAIL_DOMAIN}"
[[ -d "${caddy_cert_dir}" ]] || {
    echo "error: caddy certificate dir not found: ${caddy_cert_dir}" >&2
    exit 1
}

install -m 0644 "${caddy_cert_dir}/${MAIL_DOMAIN}.crt" certs/
install -m 0644 "${caddy_cert_dir}/${MAIL_DOMAIN}.key" certs/

docker run --rm \
    -e STALWART_URL \
    -e STALWART_USER \
    -e STALWART_PASSWORD \
    stalwartlabs/cli create action/ReloadTlsCertificates