COMPOSE := docker compose -f docker-compose.yaml --env-file .env

.PHONY: up down restart logs

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

restart:
	$(COMPOSE) down
	$(COMPOSE) up -d

logs:
	$(COMPOSE) logs -f postgres
