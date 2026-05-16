#!/usr/bin/env bash

set -euo pipefail

usage() {
  echo "Usage: $0 <backup-file.sql>"
  echo
  echo "Restores SQL backup into configured Weight Tracker database."
  echo "Stops app service before restore and starts it again afterwards."
}

if [[ ${1:-} == "-h" || ${1:-} == "--help" ]]; then
  usage
  exit 0
fi

if [[ $# -ne 1 ]]; then
  usage
  exit 1
fi

backup_file=$1
if [[ ! -f "$backup_file" ]]; then
  echo "Backup file not found: $backup_file"
  exit 1
fi

if ! command -v psql >/dev/null 2>&1; then
  echo "psql must be available in PATH"
  exit 1
fi

env_file=${WEIGHT_TRACKER_ENV_FILE:-/etc/weight-tracker/weight-tracker.env}
app_service=${WEIGHT_TRACKER_APP_SERVICE:-weight-tracker.service}
postgres_service=${WEIGHT_TRACKER_POSTGRES_SERVICE:-weight-tracker-postgres.service}

read_env_value() {
  local key=$1
  local file=$2
  local line current_key current_value first_char last_char

  while IFS= read -r line || [[ -n "$line" ]]; do
    line=${line%$'\r'}
    [[ -z "$line" || ${line:0:1} == "#" ]] && continue
    [[ "$line" != *=* ]] && continue

    current_key=${line%%=*}
    current_value=${line#*=}

    if [[ "$current_key" == "$key" ]]; then
      if [[ ${#current_value} -ge 2 ]]; then
        first_char=${current_value:0:1}
        last_char=${current_value: -1}
        if [[ ("$first_char" == '"' && "$last_char" == '"') || ("$first_char" == "'" && "$last_char" == "'") ]]; then
          current_value=${current_value:1:-1}
        fi
      fi

      printf '%s' "$current_value"
      return 0
    fi
  done < "$file"

  return 0
}

if [[ ! -f "$env_file" ]]; then
  echo "Environment file not found: $env_file"
  exit 1
fi

db_host=$(read_env_value "WEIGHT_TRACKER_DB_HOST" "$env_file")
db_port=$(read_env_value "WEIGHT_TRACKER_DB_PORT" "$env_file")
db_name=$(read_env_value "WEIGHT_TRACKER_DB_NAME" "$env_file")
db_user=$(read_env_value "WEIGHT_TRACKER_DB_USER" "$env_file")
db_password=$(read_env_value "WEIGHT_TRACKER_DB_PASSWORD" "$env_file")

db_host=${db_host:-127.0.0.1}
db_port=${db_port:-5432}

if [[ -z "$db_name" || -z "$db_user" || -z "$db_password" ]]; then
  echo "WEIGHT_TRACKER_DB_NAME, WEIGHT_TRACKER_DB_USER and WEIGHT_TRACKER_DB_PASSWORD are required"
  exit 1
fi

echo "Stopping $app_service"
sudo systemctl stop "$app_service"

echo "Starting $postgres_service"
sudo systemctl start "$postgres_service"

echo "Dropping and recreating public schema in $db_name"
PGPASSWORD="$db_password" psql -h "$db_host" -p "$db_port" -U "$db_user" -d "$db_name" -v ON_ERROR_STOP=1 -c "DROP SCHEMA IF EXISTS public CASCADE; CREATE SCHEMA public AUTHORIZATION \"$db_user\";"

echo "Restoring from $backup_file"
tmp_restore_file=$(mktemp)
trap 'rm -f "$tmp_restore_file"' EXIT

while IFS= read -r line || [[ -n "$line" ]]; do
  case "$line" in
    "\\restrict "*|"\\unrestrict "*)
      continue
      ;;
    "SET SESSION AUTHORIZATION "*";"|*" OWNER TO "*)
      continue
      ;;
  esac

  printf '%s\n' "$line" >> "$tmp_restore_file"
done < "$backup_file"

PGPASSWORD="$db_password" psql -h "$db_host" -p "$db_port" -U "$db_user" -d "$db_name" -v ON_ERROR_STOP=1 -f "$tmp_restore_file"

echo "Starting $app_service"
sudo systemctl start "$app_service"

echo "Restore completed"
