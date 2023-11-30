
APP_PKG = $(shell go list github.com/brcodingdev/go-crud-users/internal/...)

lint:
	@echo "Linting"
	@golint -set_exit_status $(APP_PKG)

test:
	@echo "Testing"
	@go test ./... -v -count=1 -race

build:
	@echo "Building docker image"
	@docker-compose build

run:
	@echo "Starting go crud users"
	@docker-compose up -d
