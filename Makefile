SHELL := /bin/bash
GO_PROJECT := github.com/djeeno/utils-go
REPOSITORY_ROOT := ~/go/src/${GO_PROJECT}
VERSION := v0.0.3
REVISION := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell TZ=UTC date +%Y%m%d%H%M%S)
GO_VERSION := $(shell go version)
BUILD_OPTS := -ldflags '-X "main.version=${VERSION}" -X "main.hash=${REVISION}" -X "main.builddate=${BUILD_DATE}" -X "main.goversion=${GO_VERSION}"'

##
# ターゲット
##
.PHONY: help
.DEFAULT_GOAL := help
help:  ## display help docs
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

test:  ## run test
	# run test
	go test -cover *.go

release:  ## release as ${VERSION}
	# release as ${VERSION}
	@if [ "`git diff; git diff --staged`" != "" ]; then echo "Uncommitted changes."; false; fi
	git tag -a "${VERSION}" -m "release ${VERSION}"
	git push
	git push --tags
