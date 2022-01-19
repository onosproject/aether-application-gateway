export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

APP_GTWY_VERSION := latest

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
	go run cmd/app/main.go

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
aether-application-gateway: # @HELP build aether-application-gateway Docker image
	@go mod vendor
	docker build . -f build/aether-application-gateway/Dockerfile \
		    -t onosproject/aether-application-gateway:${APP_GTWY_VERSION}
	@rm -rf vendor

images: # @HELP build all Docker images (the build happens inside a docker container)
images: aether-application-gateway

docker-build: aether-application-gateway

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
