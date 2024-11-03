BINARY=go-svc
VERSION := $(shell git describe --tags --always --dirty)
GIT_COMMIT := $(shell git rev-list -1 HEAD)

.DEFAULT_GOAL := build
.PHONY: version lint test build clean

version: ## Show version
	@echo $(VERSION) \(git commit: $(GIT_COMMIT)\)

lint: ## Run linters
	golangci-lint run --color always ./...

test: ## Run only quick tests
	go test -short  ./...

build: ## Build binary
	go build -o ${BINARY} main.go

clean: ## Remove binary
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
