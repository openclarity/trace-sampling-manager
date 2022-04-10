e_Y=\033[1;33m
C_C=\033[0;36m
C_M=\033[0;35m
C_R=\033[0;41m
C_N=\033[0m
SHELL=/bin/bash

# Project variables
BINARY_NAME ?= trace-sampling-manager
DOCKER_REGISTRY ?= gcr.io/eticloud/k8sec
VERSION ?= $(shell git rev-parse HEAD)
DOCKER_IMAGE ?= $(DOCKER_REGISTRY)/$(BINARY_NAME)
DOCKER_TAG ?= ${VERSION}

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: build
build: ## Build Trace Sampling Manager
	@(echo "Building Trace Sampling Manager ..." )
	@(cd manager && go mod tidy && go build -o bin/trace-sampling-manager cmd/manager/main.go && ls -l bin/)

.PHONY: build_ci
build_ci:
	@(cd manager && go build -o $(WORKSPACE)/artifacts/trace-sampling-manager cmd/manager/main.go)

.PHONY: docker_build
docker_build: ## Build Trace Sampling Manager docker image
	@(echo "Building Trace Sampling Manager docker image [${DOCKER_IMAGE}:${DOCKER_TAG}] ..." )
	@(cd manager && GOOS=linux go build -o bin/trace-sampling-manager cmd/manager/main.go)
	@(mkdir docker/artifacts && mv manager/bin/trace-sampling-manager docker/artifacts)
	@(cd docker && docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} . && rm -rf ./artifacts)

.PHONY: docker_push
docker_push: docker_build ## Build Trace Sampling Manager docker image and push it to remote
	@(echo "Pushing Trace Sampling Manager docker image [${DOCKER_IMAGE}:${DOCKER_TAG}] ..." )
	@(docker push ${DOCKER_IMAGE}:${DOCKER_TAG})

.PHONY: api
api: ## Generating API code
	@(echo "Generating API code ..." )
	@(cd api; ./generate.sh)

.PHONY: test
test: ## Run Unit Tests
	@(cd manager && go test ./pkg/...)

.PHONY: clean
clean: ## Clean all build artifacts
	@(rm -rf manager/bin/* ; echo "Build artifacts cleanup done" )

