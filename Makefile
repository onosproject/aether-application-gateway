# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

export CGO_ENABLED=1
export GO111MODULE=on

APP_GTWY_VERSION := latest

# Variables
VERSION                  ?= $(shell cat ./VERSION)

## Docker related
DOCKER_USER              ?=
DOCKER_PASSWORD          ?=
DOCKER_REGISTRY          ?=
DOCKER_REPOSITORY        ?= onosproject/
DOCKER_BUILD_ARGS        ?=
DOCKER_TAG               ?= latest
DOCKER_IMAGENAME         := ${DOCKER_REGISTRY}${DOCKER_REPOSITORY}aether-application-gateway:${DOCKER_TAG}

.PHONY: build

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/app: run the cmd/app application
.PHONY: run/app
run/app:
	go run cmd/aether-application-gateway/main.go

## tidy: run go mod tidy and vendor
.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

## deps-upgrade: get latest dependency versions and run go mod tidy and vendor
deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #
build: # @HELP build the Go binaries and run all validations (default)
build:
	CGO_ENABLED=1 go build -o build/_output/aether-application-gateway ./cmd/aether-application-gateway

test: # @HELP run the unit tests and source code validation
test: build deps linters license #openapi-linters
	CGO_ENABLED=1 go test -race github.com/onosproject/aether-application-gateway/cmd/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: deps license linters # openapi-linters
	CGO_ENABLED=1 TEST_PACKAGES=github.com/onosproject/aether-application-gateway/... ./build/build-tools/build/jenkins/make-unit


aether-application-gateway-docker: # @HELP build aether-application-gateway Docker image
	@go mod vendor
	docker build . -f build/aether-application-gateway/Dockerfile \
		    -t onosproject/aether-application-gateway:${APP_GTWY_VERSION}
	@rm -rf vendor

images: # @HELP build all Docker images (the build happens inside a docker container)
images: aether-application-gateway-docker

docker-build: aether-application-gateway-docker

docker-push: # push to docker registry: use DOCKER_REGISTRY, DOCKER_REPOSITORY and DOCKER_TAG to customize
ifdef DOCKER_USER
ifdef DOCKER_PASSWORD
	echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USER} --password-stdin
else
	@echo "DOCKER_USER is specified but DOCKER_PASSWORD is missing"
	@exit 1
endif
endif
	docker push ${DOCKER_IMAGENAME}

openapi-spec-validator: # @HELP install openapi-spec-validator
	openapi-spec-validator -h || python -m pip install openapi-spec-validator==0.3.1

openapi-linters: # @HELP lints the Open API specifications
openapi-linters: openapi-spec-validator
	openapi-spec-validator api/app-gtwy-openapi3.yaml

all: images

publish:
	./build/build-tools/publish-version ${VERSION} onosproject/aether-appliction-gateway

jenkins-publish: build-tools docker-build docker-push # @HELP Jenkins calls this to publish artifacts
	./build/build-tools/release-merge-commit
