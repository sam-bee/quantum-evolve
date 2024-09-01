.PHONY: build test

build:
	@echo "Building..."
	@go build -o quantum-evolve main.go

test:
	@go fmt ./...;
	@go build ./...;
	@go test -v ./...;
	@echo "\n\n";
