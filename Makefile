# Makefile
# Config.yaml file
CONFIG_FILE=config/config.yaml

# Use yq to extract values from the YAML file
DB_USER=$(shell yq '.database.user' $(CONFIG_FILE))
DB_PASSWORD=$(shell yq '.database.userpassword' $(CONFIG_FILE))
DB_HOST=$(shell yq '.database.host' $(CONFIG_FILE))
DB_PORT=$(shell yq '.database.port' $(CONFIG_FILE))
DB_NAME=$(shell yq '.database.database_name' $(CONFIG_FILE))
DB_OPTIONS=$(shell yq '.database.options' $(CONFIG_FILE))

# Combine into full database URL
DATABASE_URL='$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)$(DB_OPTIONS)'


# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
MIGRATION_DIR = internal/db/migrations
DB_DRIVER = mysql

# Default target executed when no arguments are given to make
all: build

compose-up : 
	docker-compose up -d
compose-down : 
	docker-compose down

tidy:
	go mod tidy

dev:
	goreload --build cmd/app
	
# Build the project
build:
	go build -o bin/$(BINARY_NAME) cmd/app/main.go

format:
	gofumpt -w .
	swag fmt .
lint:
	go run -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0 run -v ./...

schedule:
	go run internal/schedule/server/server.go	