#!/usr/bin/env bash
# Export the legacy account deletion tombstones before migration
# 20260912120000_account_identity_platform_uuid.sql drops them. The migration
# clears these rows because they name pre-cutover 24-hex identities with no live
# platform identity, so this is the only export point once audit retention of
# account deletions matters.
#
# Usage:
#   POSTGRES_MIGRATOR_URI=postgres://... ./export.sh legacy-tombstones.csv
#
# Run it before applying migration 20260912120000 to a database that still has
# account_deletion_tombstones.external_user_id.
set -euo pipefail

uri="${POSTGRES_MIGRATOR_URI:?POSTGRES_MIGRATOR_URI is required}"
output="${1:?usage: export.sh <output.csv>}"

psql "$uri" --set=ON_ERROR_STOP=1 --command "\copy (SELECT external_user_id, deleted_at FROM account_deletion_tombstones ORDER BY deleted_at) TO '$output' WITH (FORMAT csv, HEADER true)"
