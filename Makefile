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

.Phony: build-binary
build-binary: api-docs
	@ echo "Build Binary"
	@ mkdir ./build
	@ cp .example.env ./build/.env
	@ go mod tidy -compat=1.22
	@ go build -o ./build/posbe ./cmd/api/main.go
	@ GOOS=windows GOARCH=amd64 go build -o ./build/posbe.exe ./cmd/api/main.go
	@ echo "generate binary done"

.Phony: api-docs
api-docs: run-tests
	@ echo "Re-generate Swagger File (API Spec docs)"
	@ swag init --parseDependency --parseInternal \
		--parseDepth 4 -g ./cmd/api/main.go
	@ echo "generate swagger file done"

.Phony: run-tests
run-tests: $(MOCKERY) $(GOTESTSUM) run-lint
	@ echo "Run tests"
	@ gotestsum --format pkgname-and-test-fails \
		--hide-summary=skipped \
		-- -coverprofile=cover.out ./...
	@ rm cover.out

.Phony: run-lint
run-lint: $(GOLANGCI)
	@ echo "Applying linter"
	@ golangci-lint cache clean
	@ golangci-lint run -c .golangci.yaml ./...

.Phony: run-api
run-api:
	@echo "Run App"
	go mod tidy -compat=1.22
	go run ./cmd/api/main.go

.Phony: run-watch-api
run-watch-api:
	go mod tidy -compat=1.22
	air

.Phony: run-app
run-app:
	@echo "Run App"
	cd ./web && npm run build && cd ..
	go mod tidy -compat=1.22
	go run ./cmd/api/main.go
