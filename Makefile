.PHONY: lint setup-pre-commit dev-frontend dev-backend dev-database dev-database-down

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
