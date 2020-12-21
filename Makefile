# ----------------------------------------------------------------------------
#  STANDARD MAKEFILE VARIABLES
# ----------------------------------------------------------------------------
# PROJECT - the name of the project
PROJECT := $(shell basename `pwd`)

# TEAM - the name of the project
TEAM := $(shell cd .. && basename `pwd`)

# DEBUG - enable debug logging in the service
DEBUG ?= false

# ENV - the environment as read from the env var ENV
ENV ?= Local

# GIT_COMMIT - the git commit
GIT_COMMIT := $(shell git rev-parse --short HEAD)

# GOLANG SPECIFICS
GOVERSION   := 1.12
GO111MODULE := on
GOPATH      ?= $(shell go env GOPATH)
GOMAXPROCS  ?= 4
GOTAGS      ?=

# PROJECT_DIR - the project directory in the $GOPATH
PROJECT_DIR  := $(GOPATH)/src/github.com/$(TEAM)/$(PROJECT)

# KERNEL - the kernel name as provided by uname -s
KERNEL := $(shell uname -s)

# VERSION - the version as read from branch version file version
VERSION := $(shell cat version)

# ----------------------------------------------------------------------------
# Standard Targets
# ----------------------------------------------------------------------------
# Usage - displays a usage based of comments starting with ##
all: usage
usage: Makefile
	@echo
	@echo "$(PROJECT) supports the following:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: usage

# Help - see usage
help: usage

# Check - checks if required variables are set
check:
	@if [ -z "$(ENV)" ]; then \
		echo 'please set $$ENV in Makefile or set an environment variable'; \
		exit 1; \
	fi

	@echo "[INFO]: make check passed"

.PHONY: check

## build - build the local docker sandbox from docker compose
build:
ifeq ($(DEBUG), true)
	PROJECT=$(PROJECT) TEAM=$(TEAM) ENV=$(ENV) \
	docker-compose -f deployments/local/docker-compose.yml up --build --abort-on-container-exit --remove-orphans --force-recreate
else
	PROJECT=$(PROJECT) TEAM=$(TEAM) ENV=$(ENV) \
	docker-compose -f deployments/local/docker-compose.yml up --build --remove-orphans --force-recreate -d
	docker attach $(PROJECT)
endif
.PHONY: build

## clean - will remove all build artifacts and configuration
clean:
	@echo "[INFO]: removing local build artifacts"
	@rm -rf vendor Godeps build

.PHONY: clean

## deps - will download and vendor dependencies
deps:
	@echo "[INFO]: installing dependencies"
	go mod init; go mod tidy; go mod download; go mod vendor

## fmt - will execute go fmt
fmt:
	go fmt

## install - will install the binary
install:
	go install -v

## lint - will lint the code
lint:
	@if [ ! -f $(GOPATH)/bin/golint ]; then \
		go get -u golang.org/x/lint/golint; \
	fi
	@echo [INFO]: linting...
	golint

## test - will execute any tests
test: fmt lint deps
	ENV=$(ENV) go test -v

.PHONY: test

## watch - will watch the local code for changes and rebuild testing container
watch:
	ENV=$(ENV) watcher
.PHONY: watch

