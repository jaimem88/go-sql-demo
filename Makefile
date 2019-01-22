# Base service vars
export LOG_FORMATTER := $(if $(LOG_FORMATTER),$(LOG_FORMATTER),text)
export LOG_LEVEL := $(if $(LOG_LEVEL),$(LOG_LEVEL),DEBUG)
export ENVIRONMENT := $(if $(ENVIRONMENT),$(ENVIRONMENT),local)
export GOCACHE := $(if $(GOCACHE),$(GOCACHE),off)

# Base DB vars
export DB_HOST := $(if $(DB_HOST),$(DB_HOST),localhost)
export DB_PORT := $(if $(DB_PORT),$(DB_PORT),5000)
export DB_NAME := $(if $(DB_NAME),$(DB_NAME),postgres)
export DB_USER := $(if $(DB_USER),$(DB_USER),postgres)
export DB_PASS := $(if $(DB_PASS),$(DB_PASS),postgres)

PROJECT_NAME				:= $(shell basename -s .git `git config --get remote.origin.url`)

# Use richgo for pretty test output if you have it installed
# https://github.com/kyoh86/richgo
# if you have richgo installed set this to 'richgo'
export RICHGO := $(if $(RICHGO),$(RICHGO),richgo)

.PHONY: run
run:: build
	./bin/$(PROJECT_NAME) ${ARGS}

.PHONY: build
build:
	@echo $(PROJECT_NAME)
	go build  -o ./bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)/

.PHONY:
check:
	golangci-lint run `go list ./... | cut -d'/' -f4-` test/endtoend/...

.PHONY: static
static::
	docker run -v $(PWD):/tmp -w /tmp \
	devlube/packr:0.0.1 packr -i /tmp/pkg/pgsql
	go fmt $(PWD)/pkg/pgsql/a_pgsql-packr.go

.PHONY: test
test:
	GOCACHE=$(GOCACHE) $(RICHGO) test ./... -timeout 120s -race ${ARGS}

.PHONY: migrations
migrations:
	go build  -o ./bin/migrations ./cmd/migrations/
	./bin/migrations
