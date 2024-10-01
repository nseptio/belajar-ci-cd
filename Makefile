# Makefile

# Load environment variables from .env file
include .env

# Variables
DB_URL := mongodb://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  run              : Run the web service"
	@echo "  watch            : Run the web service with live-reload"

.PHONY: run
run:
	go run main.go

.PHONY: watch
watch:
	air -c air.conf