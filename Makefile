NAME := slqls
VERSION := $(shell git describe --tags `git rev-list --tags --max-count=1`)
REVISION := $(shell git rev-parse --short HEAD)

LDFLAGS := -ldflags="-s -w -X \"main.version=$(VERSION)\" -X \"main.revision=$(REVISION)\""

.PHONY: test
test:
	go test ./... -v -vet=off

.PHONY: test-race
test-race:
	go test ./... -v -race -vet=off

.PHONY: install
install:
	go install $(LDFLAGS) ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run

.PHONY: coverage
coverage:
	go test ./... -v -coverprofile=coverage.out -vet=off
	go tool cover -html=coverage.out

.PHONY: snapshot
snapshot:
	go run github.com/goreleaser/goreleaser --snapshot --rm-dist
