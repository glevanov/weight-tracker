SHELL := /bin/bash

COMPOSE := docker compose -f docker-compose.yaml --env-file .env
BACKUP_DIR := backups
BACKUP_FILE := backups/weight-tracker.sql

.PHONY: up down restart logs backup restore

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

restart:
	$(COMPOSE) down
	$(COMPOSE) up -d

logs:
	$(COMPOSE) logs -f postgres

backup:
	mkdir -p $(BACKUP_DIR)
	$(COMPOSE) exec -T postgres sh -c 'pg_dump -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"' > $(BACKUP_FILE)

restore:
	@test -f $(BACKUP_FILE) || { printf "Backup file not found: %s\n" "$(BACKUP_FILE)"; exit 1; }
	$(COMPOSE) exec -T postgres sh -c 'psql -v ON_ERROR_STOP=1 -U "$$POSTGRES_USER" -d "$$POSTGRES_DB" -c "DROP SCHEMA IF EXISTS public CASCADE; CREATE SCHEMA public;"'
	$(COMPOSE) exec -T postgres sh -c 'psql -v ON_ERROR_STOP=1 -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"' < $(BACKUP_FILE)

# Example dump from a local Postgres:
# PGPASSWORD=password pg_dump -h localhost -p 5432 -U user -d weight-tracker > backups/weight-tracker.sql
# Example restore to a local Postgres:
# PGPASSWORD=password psql -h localhost -p 5432 -U user -d weight-tracker -v ON_ERROR_STOP=1 -f backups/weight-tracker.sql
