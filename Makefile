CONFIG_FILE ?= ./configs/apiserver.yaml
APP_DSN ?= $(shell sed -n 's/^database_url:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
MIGRATE := docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate:v4.10.0 -path=/migrations/ -database "$(APP_DSN)"

.PHONY build:
build:
		go build -v ./cmd/apiserver

.DEFAULT_GOAL := build

.PHONY: db-start
db-start:
	@mkdir -p testdata/postgres
	docker run --rm --name postgres -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=go_restful -d -p 5432:5432 postgres

.PHONY: db-stop
db-stop:
	docker stop postgres

.PHONY: migrate
migrate:
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down:
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new:
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir /migrations/ $${name// /_}

.PHONY: migrate-reset
migrate-reset:
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up
