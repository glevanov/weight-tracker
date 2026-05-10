#!/usr/bin/env bash

set -euo pipefail

postgres_bin=${POSTGRES_BIN:-$(command -v postgres || true)}
initdb_bin=${INITDB_BIN:-$(command -v initdb || true)}

if [[ -z "$postgres_bin" || -z "$initdb_bin" ]]; then
  echo "postgres and initdb must be available in PATH (or set POSTGRES_BIN/INITDB_BIN)"
  exit 1
fi

PGDATA_DIR=${WEIGHT_TRACKER_PGDATA:-/var/lib/weight-tracker-postgres/data}
PGHOST_VALUE=${WEIGHT_TRACKER_DB_HOST:-127.0.0.1}
PGPORT_VALUE=${WEIGHT_TRACKER_DB_PORT:-5432}
PGSOCKET_DIR=${WEIGHT_TRACKER_PG_SOCKET_DIR:-/var/lib/weight-tracker-postgres}

install -d -m 0700 "$PGDATA_DIR"
install -d -m 0770 "$PGSOCKET_DIR"

if [[ ! -f "$PGDATA_DIR/PG_VERSION" ]]; then
  "$initdb_bin" -D "$PGDATA_DIR" --encoding=UTF8 --locale=C.UTF-8
fi

exec "$postgres_bin" -D "$PGDATA_DIR" -h "$PGHOST_VALUE" -p "$PGPORT_VALUE" -k "$PGSOCKET_DIR"
