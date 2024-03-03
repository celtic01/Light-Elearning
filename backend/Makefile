.PHONY: start-db stop-db

# PostgreSQL configuration
POSTGRES_USER := postgres
POSTGRES_PASSWORD := postgres
POSTGRES_DB := goapi
POSTGRES_PORT := 5432

# Docker container configuration
CONTAINER_NAME := goapi-db
DOCKER_IMAGE := postgres:latest

start-db:
	@echo "Starting PostgreSQL container..."
	@docker run --name $(CONTAINER_NAME) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-p $(POSTGRES_PORT):$(POSTGRES_PORT) \
		-e TZ=Europe/Bucharest \
		-d $(DOCKER_IMAGE)

stop-db:
	@echo "Stopping PostgreSQL container..."
	@docker stop $(CONTAINER_NAME)
	@docker rm $(CONTAINER_NAME)
