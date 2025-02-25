include .env
export

PWD   := $(shell pwd)
POSTGRESQL_URL := postgres://$(USER):$(USER)@$(HOST):$(PORT)/manager?sslmode=disable



.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: up-docker
up-docker: ## Run docker in silent mode
	$(call print-target)
	@docker compose up -d

.PHONY: stop-docker
stop-docker: ## Stop docker
	$(call print-target)
	@docker-compose stop

.PHONY: fix
fix: ## Format code
	$(call print-target)
	@go fmt ./...

.PHONY: migrate-up
migrate-up: ## Up migration
	$(call print-target)
	@migrate -database $(POSTGRESQL_URL) -path db/migrations up

.PHONY: migrate-down
migrate-down: ## DOWN migration
	$(call print-target)
	@migrate -database $(POSTGRESQL_URL) -path db/migrations down
.PHONY: migrate-force
migrate-force: ## Migrate version force version
	$(call print-target)
	@migrate -database $(POSTGRESQL_URL) -path db/migrations force $(i)

.PHONY: sql
show-table: ##Simple table struct view 
	$(call print-target)
	@psql  $(POSTGRESQL_URL) -c "\d $(table)"


define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
