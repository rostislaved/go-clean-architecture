include .env
export

PROJECT_ROOT := $(shell git rev-parse --show-toplevel)
COVERAGE_FILE := $(PROJECT_ROOT)/$(NAME_COVERAGE_FILE)


dev:
	docker compose -f docker/dev/dev.yml up --no-log-prefix --attach $(APP_NAME)
.PHONY: run

prod:
	docker compose -f docker/prod/prod.yml up --build --no-log-prefix --attach $(APP_NAME)
.PHONY: prod

lint:
	golangci-lint run -c .golangci.yml --fix
.PHONY: lint

remove:
	docker rm $(APP_NAME)
.PHONY: remove

stop:
	docker stop $(APP_NAME)
.PHONY: stop

clean: stop remove


test:
	# Cоветую присмотреться к ginko
	go test ./... --cover  --coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE) -o coverage.html

show_cover:
	open coverage.html
