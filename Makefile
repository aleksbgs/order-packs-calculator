# Makefile for Order Packs Calculator

# Variables
BINARY_NAME=order-packs-calculator
GO=go
MOCKGEN=$(GO) run github.com/golang/mock/mockgen@v1.6.0

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	$(GO) build -o $(BINARY_NAME) ./cmd/api

# Run the application locally
.PHONY: run
run: build
	./$(BINARY_NAME)

# Run the application with Docker Compose
.PHONY: docker-up
docker-up:
	docker-compose up --build

# Stop Docker Compose containers
.PHONY: docker-down
docker-down:
	docker-compose down

# Generate mocks for testing
.PHONY: generate-mocks
generate-mocks:
	$(MOCKGEN) -source=internal/infrastructure/repository/pack_repository.go -destination=internal/infrastructure/repository/mocks/pack_repository_mock.go -package=mocks
	$(MOCKGEN) -source=internal/service/calculate_packs.go -destination=internal/service/mocks/calculate_packs_mock.go -package=mocks

# Run all tests
.PHONY: test
test: test-unit test-integration

# Run unit tests
.PHONY: test-unit
test-unit:
	$(GO) test -v ./internal/...

# Run integration tests
.PHONY: test-integration
test-integration:
	$(GO) test -v ./tests/...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	$(GO) test ./internal/... -coverprofile=coverage_internal.out
	$(GO) test ./tests/... -coverprofile=coverage_tests.out

# Clean up generated files
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	$(GO) clean
