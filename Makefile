BIN := pokecli
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS := "-s -w -X main.revision=$(CURRENT_REVISION)"

.PHONY: build
build:
	go build -ldflags=$(BUILD_LDFLAGS) -o $(BIN) .

.PHONY: install
install:
	go install -ldflags=$(BUILD_LDFLAGS) .

.PHONY: clean
clean:
	go clean

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: coverage
coverage:
	go test -v -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: tidy
tidy:
	go mod tidy
