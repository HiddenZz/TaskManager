PWD   := $(shell pwd)

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


define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
