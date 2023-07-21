# Base Go commands.
GO_CMD   := go
GO_FMT   := $(GO_CMD) fmt
GO_CLEAN := $(GO_CMD) clean
GO_BUILD := $(GO_CMD) build

# Base golangci-lint commands.
GCL_CMD := golangci-lint
GCL_RUN := $(GCL_CMD) run

# Project executable file, and its binary.
CMD_PATH    := ./cmd/hayasui
BINARY_NAME := hayasui

# Default makefile target.
.DEFAULT_GOAL := bot

# Standarize go coding style for the whole project.
.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

# Lint go source code.
.PHONY: lint
lint: fmt
	@$(GCL_RUN) -D errcheck --timeout 5m

# Clean project binary, test, and coverage file.
.PHONY: clean
clean:
	@$(GO_CLEAN) ./...

# Build the project executable binary.
.PHONY: build
build: clean fmt
	@cd $(CMD_PATH); \
	$(GO_BUILD) -o $(BINARY_NAME) -v .

# Build and run the discord bot.
.PHONY: bot
bot: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) bot

# Docker base command.
DOCKER_CMD   := docker
DOCKER_IMAGE := $(DOCKER_CMD) image

# Docker-compose base command and docker-compose.yml path.
COMPOSE_CMD   := docker-compose
COMPOSE_BUILD := deployment/build.yml
COMPOSE_BOT   := deployment/bot.yml

# Build docker images and container for the project
# then delete builder image.
.PHONY: docker-build
docker-build: clean fmt
	@$(COMPOSE_CMD) -f $(COMPOSE_BUILD) build
	@$(DOCKER_IMAGE) prune -f --filter label=stage=hayasui_builder

# Start built docker containers.
.PHONY: docker
docker:
	@$(COMPOSE_CMD) -f $(COMPOSE_BOT) up -d
	@$(COMPOSE_CMD) -f $(COMPOSE_BOT) logs --follow --tail 20

# Stop docker container.
.PHONY: docker-stop
docker-stop:
	@$(COMPOSE_CMD) -f $(COMPOSE_BOT) stop