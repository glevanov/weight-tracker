SHELL := /bin/bash

.PHONY: run tidy

run:
	set -a; [ -f .env ] && source ./.env; set +a; go run . $(ARGS)

tidy:
	go mod tidy
