.PHONY: build/bin build/docker clean run

BUILD_ENV        := CGO_ENABLED=1
LDFLAGS          := -s -w

SRC_INVOICE_FILE := ./cmd/invoice
BIN := bin/invoice
TEST := grc go test -v -failfast -cover -race
TEST_MODULES = ./...

all: test

help: ## Print help for each target
	$(info invoice Makefile help.)
	$(info ====================================)
	$(info )
	$(info Available commands:)
	$(info )
	@grep '^[[:alnum:]_/]*:.* ##' $(MAKEFILE_LIST) \
		| sort | awk 'BEGIN {FS=":.* ## "}; {printf "%-25s %s\n", $$1, $$2};'

run/docker: ## Run docker compose with air
	@docker compose up --abort-on-container-exit air

deps: ## Install and update go dependicies
	@go mod download && go mod tidy

build/bin: deps ## Build project
	@mkdir -p bin
	@$(BUILD_ENV) go build -o $(BIN) -ldflags="$(LDFLAGS)" $(SRC_INVOICE_FILE)

build/docker: ## Build project with docker
	DOCKER_BUILDKIT=1 docker build -f Dockerfile -t invoice .

test: ## Run all tests
	@${TEST} ${TEST_MODULES}

clean: ## Clean build
	@rm -rf ./bin
