.PHONY: proto dep lint build vendor

# load env variables from .env
ENV_PATH ?= ./.env
ifneq ($(wildcard $(ENV_PATH)),)
    include .env
    export
endif

# service code
SERVICE = decision

# current version
DOCKER_TAG ?= latest
# docker registry url
DOCKER_URL = gitlab.secreate.dev:5050/ev2go/server

# database migrations
DB_ADMIN_USER ?= admin
DB_ADMIN_PASSWORD ?= admin
DB_HOST ?= localhost
DB_NAME ?= ev2go
DB_PORT ?= 35432
DB_OCPI_USER ?= $(SERVICE)
DB_OCPI_PASSWORD ?= $(SERVICE)

DB_DRIVER = postgres
DB_STRING = "user=$(DB_OCPI_USER) password=$(DB_OCPI_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable"
DB_ADMIN_STRING = "user=$(DB_ADMIN_USER) password=$(DB_ADMIN_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable"
DB_INIT_FOLDER = "./db/init"
DB_MIG_FOLDER = "./db/migrations"

export GOFLAGS=-mod=vendor

# Build commands =======================================================================================================

vendor:
	go mod vendor

dep:
	go env -w GO111MODULE=on
	go env -w GOPRIVATE=gitlab.secreate.dev/*
	go mod tidy

check-lint-installed:
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed"; \
		exit 1; \
	fi; \

lint: check-lint-installed
	@echo Running golangci-lint
	golangci-lint --modules-download-mode vendor run --timeout 5m0s --skip-dirs-use-default ./...
	go fmt ./...

mock: ## generate mocks for all the directories except root and vendor
	@rm -R ./mocks 2> /dev/null; \
	find . -maxdepth 1 -type d \( ! -path ./*vendor ! -name . \) -exec mockery --all --dir {} \;

build: lint ## builds the main
	@mkdir -p bin
	go build -o bin/ cmd/main.go

artifacts: dep vendor mock build ## builds and generates all artifacts

run: ## run the service
	./bin/main

# Database commands ====================================================================================================

check-goose-installed:
	@if ! [ -x "$$(command -v goose)" ]; then \
		echo "goose is not installed"; \
		exit 1; \
	fi; \

db-init-schema:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_ADMIN_STRING) goose -dir $(DB_INIT_FOLDER) up

db-status: check-goose-installed
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir $(DB_MIG_FOLDER) status

db-up: check-goose-installed
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir $(DB_MIG_FOLDER) up

db-down: check-goose-installed
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir $(DB_MIG_FOLDER) down

db-create: check-goose-installed
	@if [ -z $(name) ]; then \
      	echo "usage: make db-create name=<you-migration-name>"; \
    else \
		GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_STRING) goose -dir $(DB_MIG_FOLDER) create $(name) sql; \
	fi

# Tests commands =======================================================================================================

test: ## run the tests
	@echo "running tests (skipping stateful)"
	go test -count=1 ./...

test-with-coverage: ## run the tests with coverage
	@echo "running tests with coverage file creation (skipping integration)"
	go test -count=1 -coverprofile .testCoverage.txt -v ./...

test-integration: ## run the integration tests
	@echo "running integration tests"
	go test -count=1 -tags integration ./...

build-test-bin: ## recursively go through folders and build integration tests to binary files
	mkdir -p bin
	@echo Bulding test binary
	for path in $$(find . -name "*_test.go" -printf '%h\n' | sort -u ); do \
  		echo $$path; \
  		fn=$$(echo $$path | sed 's+/+_+g' | sed -e 's/\.//g'); \
  		fn=$$fn"_test"; \
  		go test -c -o ./bin/$$fn -tags integration $$path ; \
  		errorCode="$$?"; \
  		if [ "$$errorCode" -gt "0" ] ; then \
  			echo "\033[31mTest build failed!\033[0m" ; \
  			exit 1 ; \
  		fi; \
  	done

# Docker commands =======================================================================================================

docker-build: ## Build the docker images for all services (build inside)
	@echo Building images
	docker build . -f ./Dockerfile -t $(DOCKER_URL)/$(SERVICE):$(DOCKER_TAG) --build-arg SERVICE=$(SERVICE)

docker-build-test: ## Build the docker images for all services (build inside)
	@echo Building images
	docker build . -f ./Dockerfile-test -t $(DOCKER_URL)/$(SERVICE):$(DOCKER_TAG)_test --build-arg SERVICE=$(SERVICE)

docker-push: docker-build ## Build and push docker images to the repository
	@echo Pushing images
	docker push $(DOCKER_URL)/$(SERVICE):$(DOCKER_TAG)

docker-push-test: docker-build-test ## Build and push docker images to the repository
	@echo Pushing images
	docker push $(DOCKER_URL)/$(SERVICE):$(DOCKER_TAG)_test

docker-run:
	@echo Running container
	docker run $(DOCKER_URL)/$(SERVICE):$(DOCKER_TAG)

# CI/CD gitlab commands =================================================================================================
ci-check-go-mod:
	invalidRequire=$$(go mod edit -json | jq '.Require | .[] | select(.Path|test("^gitlab\\.secreate\\.dev\\/ev2go\\/server\\/.")) | select(.Version|test("^v\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}-\\d{8}-\\d{4}$$")|not) |.Path'); \
    if ! [ "$$invalidRequire" = "" ] ; then \
        echo "\033[31mCheck go.mod! Contains invalid modules:\033[0m" ; \
        echo "$$invalidRequire"; \
        exit 1 ; \
    fi; \
    invalidReplace=$$(go mod edit -json | jq '.Replace | .[] | .Old | select(.Path|test("^gitlab\\.secreate\\.dev\\/ev2go\\/server\\/.")) | .Path'); \
    if ! [ "$$invalidReplace" = "" ] ; then \
        echo "\033[31mCheck go.mod! Contains invalid modules:\033[0m" ; \
        echo "$$invalidReplace"; \
        exit 1 ; \
    fi

ci-check-mocks:
	@mv ./mocks ./mocks-init
	find . -maxdepth 1 -type d \( ! -path ./*vendor ! -name . \) -exec mockery --all --dir {} \;
	mockshash=$$(find ./mocks -type f -print0 | sort -z | xargs -r0 md5sum | awk '{print $$1}' | md5sum | awk '{print $$1}'); \
	mocksinithash=$$(find ./mocks-init -type f -print0 | sort -z | xargs -r0 md5sum | awk '{print $$1}' | md5sum | awk '{print $$1}'); \
	rm -fr ./mocks-init; \
	echo $$mockshash $$mocksinithash; \
	if ! [ "$$mockshash" = "$$mocksinithash" ] ; then \
	  echo "\033[31mMocks should be updated!\033[0m" ; \
	  exit 1 ; \
	fi

ci-check: ci-check-go-mod ci-check-mocks

ci-build: test-with-coverage docker-push docker-push-test

ci-build-mr: test-with-coverage docker-build

ci-run-test-bin: ## executes all tests binary from inside test container
	@find . -name "*_test" -type f -exec sh -c {} \;
