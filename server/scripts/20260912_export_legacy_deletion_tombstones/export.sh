#!/usr/bin/env bash
# Export the legacy account deletion tombstones from a pre-baseline database.
# The baseline schema does not carry account_deletion_tombstones.external_user_id
# because those rows name pre-cutover 24-hex identities with no live platform
# identity, and a pre-baseline database is recreated from the baseline rather
# than upgraded. This is the only export point once audit retention of account
# deletions matters.
#
# Usage:
#   POSTGRES_MIGRATOR_URI=postgres://... ./export.sh legacy-tombstones.csv
#
# Run it against a database that still has
# account_deletion_tombstones.external_user_id, before recreating it from the
# baseline.
set -euo pipefail

uri="${POSTGRES_MIGRATOR_URI:?POSTGRES_MIGRATOR_URI is required}"
output="${1:?usage: export.sh <output.csv>}"

psql "$uri" --set=ON_ERROR_STOP=1 --command "\copy (SELECT external_user_id, deleted_at FROM account_deletion_tombstones ORDER BY deleted_at) TO '$output' WITH (FORMAT csv, HEADER true)"
