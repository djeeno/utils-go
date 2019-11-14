SHELL := /bin/bash
GO_PROJECT := github.com/djeeno/utils-go
REPOSITORY_ROOT := ~/go/src/${GO_PROJECT}
VERSION := v0.1.1
REVISION := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell TZ=UTC date +%Y%m%d%H%M%S)
GO_VERSION := $(shell go version)
BUILD_OPTS := -ldflags '-X "main.version=${VERSION}" -X "main.hash=${REVISION}" -X "main.builddate=${BUILD_DATE}" -X "main.goversion=${GO_VERSION}"'
OPEN_CMD := $(shell if command -v explorer.exe; then true; elif command -v open; then true; else echo echo; fi)

##
# targets
##
.PHONY: help
.DEFAULT_GOAL := help
help:  ## display help docs
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

init:  ## init
	if ! command -v golangci-lint 1>/dev/null; then\
		go install github.com/golangci/golangci-lint/cmd/golangci-lint;\
	fi

lint: init ## run lint
	# go fmt
	go fmt ./...
	# run lint
	golangci-lint run

test: lint ## run test
	# run test
	go test -race -covermode=atomic ./...

test-v: lint ## run test
	# run test
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html
	${OPEN_CMD} coverage.html

check-uncommitted:
	@if [ "`git diff; git diff --staged`" != "" ]; then\
		echo "Uncommitted changes. Execute the following command:";\
		echo "git commit -m 'release ${VERSION}'";\
		false;\
	fi

release: test check-uncommitted ## release as ${VERSION}
	# release ${VERSION}
	git tag -a "${VERSION}" -m "release ${VERSION}"
	git push
	git push --tags
