# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=ectolinq

# Linter
GOLINT=golangci-lint

.PHONY: all build clean test coverage lint fmt mod-tidy help

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

lint:
	$(GOLINT) run

fmt:
	$(GOCMD) fmt ./...

mod-tidy:
	$(GOMOD) tidy

help:
	@echo "Available commands:"
	@echo "  make build      - Build the binary"
	@echo "  make clean      - Remove binary and clean project"
	@echo "  make test       - Run tests"
	@echo "  make coverage   - Run tests with coverage"
	@echo "  make lint       - Run linter"
	@echo "  make fmt        - Format code"
	@echo "  make mod-tidy   - Tidy go modules"
	@echo "  make help       - Print this help message"