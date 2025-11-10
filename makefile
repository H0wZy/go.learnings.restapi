.PHONY: default run gobuild build test docs cls up down logs runwdocs stop

-include .env
export

APP_NAME ?= go.learnings.api.rest

default: runwdocs

run:
	@go run main.go

runwdocs:
	@swag init
	@go run main.go

gobuild:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./...

docs:
	@swag init

cls:
	@rm -f $(APP_NAME)
	@rm -rf ./docs

build:
	@docker compose build

up:
	@docker compose up -d

stop:
	@docker compose stop

down:
	@docker compose down

logs:
	@docker compose logs -f
