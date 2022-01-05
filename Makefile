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
