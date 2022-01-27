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

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

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

all: images

publish:
	./../build-tools/publish-version ${VERSION} onosproject/aether-appliction-gateway

jenkins-publish: build-tools docker-build docker-push # @HELP Jenkins calls this to publish artifacts
	../build-tools/release-merge-commit

license_check: # @HELP examine and ensure license headers exist
license_check: build-tools
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR} --boilerplate LicenseRef-ONF-Member-1.0

golang-ci: # @HELP install golang-ci if not present
	golangci-lint --version || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b `go env GOPATH`/bin v1.42.0

linters: golang-ci # @HELP examines Go source code and reports coding problems
	golangci-lint run --timeout 5m

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build-tools deps license_check linters # openapi-linters
	CGO_ENABLED=1 TEST_PACKAGES=github.com/onosproject/aether-application-gateway/... ./../build-tools/build/jenkins/make-unit
