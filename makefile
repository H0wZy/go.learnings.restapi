.PHONY:	default run build test docs cls

APP_NAME=go.learnings.restapi

default: run
run:
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
cls:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
