# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# --- Tooling & Variables ----------------------------------------------------------------
include ./misc/make/tools.Makefile

install-deps: gotestsum mockery
deps: $(GOTESTSUM) $(MOCKERY)
deps:
	@ echo "Required Tools Are Available"

# action: up | down - usage make migrate action=down|up
.PHONY: migrate
migrate:
	@ if [ "$(action)" ]; then \
    	migrate -database "postgresql://postgres:@127.0.0.1:5432/posbe?sslmode=disable" -path db/migrations $(action); \
    	echo "migration ${action} done"; \
   	else \
    	echo "missing action (use: make migrate action=up or action=down)"; \
    fi

.PHONY: api-specs
api-specs:
	@ swag init --parseDependency --parseInternal \
		--parseDepth 4 -g ./cmd/api/main.go
	@ echo "api spec generated"

.PHONY: tests
tests: $(MOCKERY) $(GOTESTSUM)
	@ gotestsum --format pkgname-and-test-fails \
		--hide-summary=skipped \
		-- -coverprofile=cover.out ./internal/account ./internal/utils
	@ rm cover.out
	@ echo "testing completed"

.PHONY: lint
lint: $(GOLANGCI)
	@ golangci-lint cache clean
	@ golangci-lint run -c .golangci.yaml ./...
	@ echo "linting completed"

.PHONY: mocks
mocks: $(MOCKERY)
	mockery --config .mockery.yml

.PHONY: run
run:
	go mod tidy -compat=1.26
	go run -race ./cmd/api/main.go

.PHONY: watch
watch:
	go mod tidy -compat=1.26
	air

.PHONY: binary
binary: tests api-specs
	@ mkdir -p ./build
	@ cp ./misc/conf/.example.env ./build/.env
	@ go mod tidy -compat=1.26
	@ go build -o ./build/posbe ./cmd/api/main.go
	@ echo "binary generated"
