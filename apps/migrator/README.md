# Mongo to Postgres Migrator

One-off migration tool for moving data from MongoDB Atlas to local PostgreSQL.

## What it migrates

- `users` collection -> `users` table
- `weight` collection -> `weights` table (`user` field maps to `username`)

## Defaults

- Mongo DB: `weight-tracker`
- Users collection: `users`
- Weights collection: `weight`
- Truncate mode: enabled (`users` and `weights` are truncated before import)

## Setup

Copy env example and fill your Mongo URI:

```bash
cp .env.example .env
```

Required environment variables:

- `MONGO_URI`
- `PG_URI` (or pass `--pg-uri`)

Optional environment variables:

- `MONGO_DB`
- `MONGO_USERS_COLLECTION`
- `MONGO_WEIGHTS_COLLECTION`

## Run

Dry run first:

```bash
make run ARGS="--dry-run"
```

Run migration:

```bash
make run
```

## Flags

- `--dry-run`
- `--truncate` / `--no-truncate`
- `--batch-size=1000`
- `--mongo-uri=...`
- `--mongo-db=...`
- `--users-collection=...`
- `--weights-collection=...`
- `--pg-uri=...`
