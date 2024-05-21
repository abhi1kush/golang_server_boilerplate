# Variables
APP_NAME := myapp
GO_CMD := go
GO_BUILD := $(GO_CMD) build
GO_MOD := $(GO_CMD) mod
DOCKER := docker
DOCKER_BUILD := $(DOCKER) build
DOCKER_RUN := $(DOCKER) run

# Directories
BUILD_DIR := build
CMD_DIR := cmd/server
CONFIG_DIR := configs

# Docker
DOCKER_IMAGE := $(APP_NAME):latest

# Default target
all: build

# Download dependencies
deps: 
	$(GO_MOD) download

# Build the Go binary
build: deps
	$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) ./$(CMD_DIR)

# Build the Docker image
docker-build: build
	$(DOCKER_BUILD) -t $(DOCKER_IMAGE) .

# Run the Docker container
docker-run:
	$(DOCKER_RUN) -e APP_PROFILE=dev -p 8080:8080 $(DOCKER_IMAGE)

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

.PHONY: all deps build docker-build docker-run clean

