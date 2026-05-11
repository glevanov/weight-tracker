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

All production variables are prefixed with `WEITHG_TRACKER_`.

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

## PostgreSQL major upgrade

### Develompent upgrade
```bash
cd apps/database
make backup
# stop service, bump docker tag, remove volume
make restore
```

### Production upgrade

TODO