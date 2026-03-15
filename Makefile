.PHONY: lint setup-pre-commit

pre-commit:
	pnpm --dir apps/frontend run lint
	$(MAKE) -C apps/backend lint

setup-pre-commit:
	./scripts/setup-pre-commit.sh
