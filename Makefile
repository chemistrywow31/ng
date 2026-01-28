.PHONY: help build run test lint clean swagger deps

# Default target
help:
	@echo "Available commands:"
	@echo "  make deps     - Download dependencies"
	@echo "  make swagger  - Generate Swagger docs"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make test     - Run tests"
	@echo "  make lint     - Run linter"
	@echo "  make clean    - Clean build artifacts"
	@echo ""
	@echo "Quick start:"
	@echo "  make deps && make swagger && make run"

# Download dependencies
deps:
	go mod tidy
	go mod download

# Generate Swagger documentation
swagger:
	@which swag > /dev/null || go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/api/main.go -o docs/swagger --parseDependency --parseInternal

# Build the application
build: swagger
	go build -ldflags="-s -w" -o bin/api ./cmd/api

# Run the application
run:
	go run ./cmd/api

# Run with live reload (requires air)
dev:
	@which air > /dev/null || go install github.com/cosmtrek/air@latest
	air

# Run tests
test:
	go test -v -race -cover ./...

# Run tests with coverage report
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Run linter
lint:
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf docs/swagger/
	rm -f coverage.out coverage.html
