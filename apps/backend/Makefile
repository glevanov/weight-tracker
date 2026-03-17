.PHONY: build run clean fmt lint test-unit test-integration tidy download docker-build

SHELL := /bin/bash

BINARY_NAME := build/server
CMD_DIR := ./cmd/server
DOCKERFILE := Dockerfile.test
TEST_ENV := TESTCONTAINERS_RYUK_DISABLED=true


build:
	go build -o $(BINARY_NAME) $(CMD_DIR)

run:
	set -a; [ -f .env ] && source ./.env; set +a; go run $(CMD_DIR)

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt ./...	

lint:
	gofmt -w .
	go vet ./...
	
test-unit:
	go test -v ./internal/...

test-integration:
	$(TEST_ENV) go test -v -count=1 ./tests/...

tidy:
	go mod tidy

download:
	go mod download

docker-build:
	docker build -f $(DOCKERFILE) -t weight-tracker-service:test .
