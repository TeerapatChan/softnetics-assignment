# Makefile
# Config.yaml file
CONFIG_FILE=config/config.yaml

BINARY_NAME=main

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

doc:
	swag init -g cmd/app/main.go