# Variables
BINARY_NAME := monogo
SRC := cmd/main.go
BIN_DIR := bin
DOCKER_IMAGE := monogo:latest

# Default target: build the Go binary
.PHONY: build
build:  ## Build the Go application binary
	@set -e; \
	mkdir -p $(BIN_DIR); \
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC)

.PHONY: run
run:  ## Run the application directly from source
	@set -e; \
	go run $(SRC)

.PHONY: lint
lint:  ## Run golangci-lint on the codebase
	@set -e; \
	golangci-lint run ./...

.PHONY: docker-build
docker-build:  ## Build the Docker image for the app
	@set -e; \
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-up
docker-up:  ## Start the app and database using docker-compose
	@set -e; \
	docker-compose up -d

.PHONY: docker-down
docker-down:  ## Stop and remove containers with docker-compose
	@set -e; \
	docker-compose down

.PHONY: clean
clean:  ## Remove build artifacts and temporary files
	@set -e; \
	rm -rf $(BIN_DIR) \
		coverage.out \
		*.log \
		dist \
		swagger.json \
		swagger.yaml

.PHONY: swagger
swagger:  ## Generate swagger files (requires swag or similar tool installed)
	@set -e; \
	swag init -g cmd/main.go --output docs --parseDependency

.PHONY: env-setup
env-setup:  ## Setup environment file from .env.example
	@set -e; \
	if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo ".env file created from .env.example"; \
	else \
		echo ".env already exists"; \
	fi

.PHONY: start
start:  ## Setup env, build, and run the app binary
	@set -e; \
	$(MAKE) env-setup; \
	$(MAKE) build; \
	$(BIN_DIR)/$(BINARY_NAME)

.PHONY: start-docker
start-docker:  ## Start the app and database using Docker Compose
	@set -e; \
	$(MAKE) env-setup; \
	$(MAKE) docker-up

.PHONY: all
all: build
