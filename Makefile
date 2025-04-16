.PHONY: build test clean run help

## build: Build the application

build:
	go build -o bin/ ./cmd/app

## run: Build and run the application
run: build
	./bin/app

## test: Run tests
test:
	go test -v ./...

## clean: Remove build artifacts
clean:
	rm -rf bin/

## help: Display available commands
help:
	@echo "Available targets:"
	@awk -F':.*?## ' '/^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort