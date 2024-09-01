include .envrc

MIGRATION_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create migrate-up migrate-down
migration:
	@migrate create -seq -ext  sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) up
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))
