DOCKERHUB_USER_NAME=lptonussi

##############
# playground #
##############


PROJECT_NAME := "github.com/tonussi/hermes"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

build: dep ## Build the binary file
	@go build -i -o build/main $(PKG)

clean: ## Remove previous build
	@rm -f ./build

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'






build_debug_server: ## build debug simple server in go (study)
	docker-compose -f debug-server.docker-compose.yml up --build

build_debug_client: ## build debug simple client in go (study)
	docker-compose -f debug-client.docker-compose.yml up --build

build_debug_hermes: ## build debug hermes
	docker-compose -f debug-hermes.docker-compose.yml up --build

build_server: ## build simple server in go without delve (study)
	docker-compose -f server.docker-compose.yml up --build

build_hermes: ## build hermes without delve
	docker-compose -f hermes.docker-compose.yml up --build

build_client: ## build simple client in go without delve (study)
	docker-compose -f client.docker-compose.yml up --build









run_debug_client: ## run debug simple client in go (study)
	docker-compose -f debug-client.docker-compose.yml up

run_debug_hermes: ## run debug hermes
	docker-compose -f debug-hermes.docker-compose.yml up

run_debug_server: ## run debug simple server in go (study)
	docker-compose -f debug-server.docker-compose.yml up

run_server: ## run simple server in go (study)
	docker-compose -f server.docker-compose.yml up

run_client: ## run simple client in go (study)
	docker-compose -f client.docker-compose.yml up






docker_down: # down running dockers by docker compose
	docker-compose -f client.docker-compose.yml down
	docker-compose -f server.docker-compose.yml down
	docker-compose -f hermes.docker-compose.yml down
