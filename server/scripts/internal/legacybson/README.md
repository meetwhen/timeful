# legacybson

This package holds the BSON storage structs the one-off MongoDB backfill scripts in `server/scripts` still decode.
It is script-only legacy code, and runtime server code must use `timeful/server/models` instead.
TASK-0199.09.03 deletes this package together with the scripts that import it.
