#!/usr/bin/env bash

set -euo pipefail

db_port=${WEIGHT_TRACKER_DB_PORT:-5432}
db_name=${WEIGHT_TRACKER_DB_NAME:-}
db_user=${WEIGHT_TRACKER_DB_USER:-}
db_password=${WEIGHT_TRACKER_DB_PASSWORD:-}
admin_host=${WEIGHT_TRACKER_PG_ADMIN_HOST:-/var/lib/weight-tracker-postgres}

if [[ -z "$db_name" || -z "$db_user" || -z "$db_password" ]]; then
  echo "WEIGHT_TRACKER_DB_NAME, WEIGHT_TRACKER_DB_USER and WEIGHT_TRACKER_DB_PASSWORD are required"
  exit 1
fi

psql_cmd=(psql -h "$admin_host" -p "$db_port" -d postgres -v ON_ERROR_STOP=1)

script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
schema_file="$script_dir/../sql/schema.sql"

if [[ ! -f "$schema_file" ]]; then
  echo "Schema file not found: $schema_file"
  exit 1
fi

"${psql_cmd[@]}" --set=app_user="$db_user" --set=app_password="$db_password" <<'SQL'
DO
$$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = :'app_user') THEN
    EXECUTE format('CREATE ROLE %I LOGIN PASSWORD %L', :'app_user', :'app_password');
  ELSE
    EXECUTE format('ALTER ROLE %I WITH LOGIN PASSWORD %L', :'app_user', :'app_password');
  END IF;
END
$$;
SQL

"${psql_cmd[@]}" --set=app_db="$db_name" --set=app_user="$db_user" <<'SQL'
DO
$$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = :'app_db') THEN
    EXECUTE format('CREATE DATABASE %I OWNER %I', :'app_db', :'app_user');
  END IF;
END
$$;
SQL

psql_app_cmd=(psql -h "$admin_host" -p "$db_port" -d "$db_name" -v ON_ERROR_STOP=1)

"${psql_app_cmd[@]}" -f "$schema_file"
