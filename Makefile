BINARY=api
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"
LINT=$(shell command -v golangci-lint 2> /dev/null)
GOCMD=go
GOTEST=$(GOCMD) test

run:
	go run main.go

build:
	go mod tidy
	go build ${LDFLAGS} -o ${BINARY} main.go

test: 
	$(GOTEST) -v -coverprofile coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	open coverage.html

check-lint:
ifndef golangci-lint
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.45.2
endif

lint: ## run static checks
	@$(MAKE) lint-impl
lint-impl: | check-lint
	@echo "Running linter..."
	@$(LINT) run -E gofmt -E golint -E goimports -E staticcheck -E revive
