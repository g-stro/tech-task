PROJECT_NAME = tech
ENV_FILE = .env.example
COMPOSE_TEST_FILE = docker-compose.test.yml

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

# Docker commands
.PHONY: build
build:
	@echo "Building Docker images..."
	docker-compose --env-file $(ENV_FILE) -p $(PROJECT_NAME) up --build -d

.PHONY: up
up:
	@echo "Starting containers..."
	docker-compose --env-file $(ENV_FILE) -p $(PROJECT_NAME) up -d

.PHONY: down
down:
	@echo "Stopping and removing containers..."
	docker-compose --env-file $(ENV_FILE) -p $(PROJECT_NAME) down

.PHONY: clean
clean:
	@echo "Removing containers, images, and volumes..."
	docker-compose --env-file $(ENV_FILE) -p $(PROJECT_NAME) down --rmi all --volumes --remove-orphans

.PHONY: restart
restart: down up

.PHONY: logs
logs:
	@echo "Tailing logs..."
	docker-compose --env-file $(ENV_FILE) -p $(PROJECT_NAME) logs -f

# Run all tests
.PHONY: tests
tests: unit-tests integration-tests

# Run unit tests
.PHONY: unit-tests
unit-tests:
	@echo "Running unit tests..."
	go test -v ./...

# Run integration tests
.PHONY: integration-tests
integration-tests:
	@echo "Running integration tests..."
	@echo "Setting up test environment..."
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_TEST_FILE) -p $(PROJECT_NAME)_test up --build -d
	@echo "Running tests..."
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_TEST_FILE) -p $(PROJECT_NAME)_test exec -T $(PROJECT_NAME)-test go test -v -tags=integration ./...
	@echo "Cleaning up test environment..."
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_TEST_FILE) -p $(PROJECT_NAME)_test down
