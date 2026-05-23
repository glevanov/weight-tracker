# weight-tracker
Personal weight tracking app.

![cat is not happy about weighing 100 kilos](./apps/frontend/public/android-chrome-512x512.png)

## App demo
https://github.com/user-attachments/assets/93f8ff24-d9c4-456a-ae0d-7f525da89009

## Structure

- `apps/frontend` - Svelte frontend
- `apps/backend` - Go backend service
- `apps/database` - database for development

There used to be a mongo migrator, see commit 153b631d69de.

## Production build

Build frontend and backend into one Linux binary:

```bash
make build-prod
```

The output binary is `build/weight-tracker`.

## Global runtime config

Use a single global env file for both app and database:

```bash
sudo mkdir -p /etc/weight-tracker
sudo cp deploy/config/weight-tracker.env.example /etc/weight-tracker/weight-tracker.env
sudo chmod 600 /etc/weight-tracker/weight-tracker.env
```

All production variables are prefixed with `WEIGHT_TRACKER_`.

## postgres dependencies

Make sure you have postgres before starting the process:

```bash
# Check where binaries are
command -v postgres || true
command -v initdb || true
command -v psql || true

# If you do not have postgres, install the dependencies
sudo dnf install -y postgresql-server postgresql
```

## systemd services

This repository provides native Postgres and app services:

- `deploy/systemd/weight-tracker-postgres.service`
- `deploy/systemd/weight-tracker-postgres-init.service`
- `deploy/systemd/weight-tracker.service`

Install units and binary:

```bash
sudo cp deploy/systemd/*.service /etc/systemd/system/
sudo install -d /opt/weight-tracker
sudo cp build/weight-tracker /opt/weight-tracker/weight-tracker
sudo cp -r deploy /opt/weight-tracker/
```

Create users:

```bash
sudo useradd --system --home /nonexistent --shell /usr/sbin/nologin weighttracker || true
sudo useradd --system --home /nonexistent --shell /usr/sbin/nologin weighttracker-postgres || true
sudo chown -R weighttracker:weighttracker /opt/weight-tracker
sudo chown -R weighttracker-postgres:weighttracker-postgres /var/lib/weight-tracker-postgres || true
```

Enable and start:

```bash
sudo systemctl daemon-reload
sudo systemctl enable --now weight-tracker-postgres.service
sudo systemctl enable --now weight-tracker-postgres-init.service
sudo systemctl enable --now weight-tracker.service
```

Check status and logs:

```bash
sudo systemctl status weight-tracker.service
sudo journalctl -u weight-tracker.service -f
```

## Backups and restore (production)

Run backup/restore after services are installed and the PostgreSQL service is running.

Create a timestamped backup:

```bash
set -a
source /etc/weight-tracker/weight-tracker.env
set +a

ts=$(date +%Y%m%d-%H%M%S)
backup_file="backups/${WEIGHT_TRACKER_DB_NAME}-${ts}.sql"
mkdir -p backups

PGPASSWORD="$WEIGHT_TRACKER_DB_PASSWORD" pg_dump -h "$WEIGHT_TRACKER_DB_HOST" -p "$WEIGHT_TRACKER_DB_PORT" -U "$WEIGHT_TRACKER_DB_USER" -d "$WEIGHT_TRACKER_DB_NAME" --no-owner --no-privileges > "$backup_file"

printf 'Backup written to %s\n' "$backup_file"
```

Restore from a backup (includes app stop/start):

```bash
sudo /opt/weight-tracker/deploy/scripts/restore-postgres.sh /path/to/weight-tracker-backup.sql
```

Script source:

- `deploy/scripts/restore-postgres.sh`

## Update the binary
Binary update requires some manual steps.

- Stop the service
```bash
sudo systemctl stop weight-tracker.service
```
- Replace the binary
- Run script to re-apply correct permissions
```bash
sudo sh /opt/weight-tracker/deploy/scripts/restore-owner.sh
```
- Start the service
```bash
sudo systemctl start weight-tracker.service
```

## PostgreSQL major upgrade

### Development upgrade
```bash
cd apps/database
make backup
# stop service, bump docker tag, remove volume
make restore
```
