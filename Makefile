.PHONY: lint setup-pre-commit dev-frontend dev-backend dev-database dev-database-down build-prod clean-prod

SHELL := /bin/bash

PROD_BINARY := build/weight-tracker
FRONTEND_DIST := apps/frontend/dist
BACKEND_EMBED_DIST := apps/backend/internal/static/dist
BUILD_BASE := /weight-tracker

pre-commit:
	pnpm --dir apps/frontend run lint
	$(MAKE) -C apps/backend lint

setup-pre-commit:
	./scripts/setup-pre-commit.sh

dev-frontend:
	pnpm --dir apps/frontend run dev

dev-backend:
	$(MAKE) -C apps/backend run

dev-database:
	$(MAKE) -C apps/database up

dev-database-down:
	$(MAKE) -C apps/database down

build-prod: clean-prod
	pnpm --dir apps/frontend install --frozen-lockfile
	VITE_API_URL=$(BUILD_BASE)/api pnpm --dir apps/frontend run build --base $(BUILD_BASE)/
	mkdir -p apps/backend/internal/static
	mkdir -p $(BACKEND_EMBED_DIST)
	rm -rf $(BACKEND_EMBED_DIST)/*
	cp -r $(FRONTEND_DIST)/* $(BACKEND_EMBED_DIST)/
	touch $(BACKEND_EMBED_DIST)/.keep
	go build -C apps/backend -o ../../$(PROD_BINARY) ./cmd/server

clean-prod:
	mkdir -p $(BACKEND_EMBED_DIST)
	rm -rf $(BACKEND_EMBED_DIST)/*
	touch $(BACKEND_EMBED_DIST)/.keep
	rm -f $(PROD_BINARY)
