SHELL := /bin/bash
GO_PROJECT := github.com/djeeno/utils-go
REPOSITORY_ROOT := ~/go/src/${GO_PROJECT}
VERSION := v0.0.0
REVISION := $(shell git rev-parse HEAD)
BUILD_DATE := $(shell TZ=UTC date +%Y%m%d%H%M%S)
GO_VERSION := $(shell go version)
BUILD_OPTS := -ldflags '-X "main.version=${VERSION}" -X "main.hash=${REVISION}" -X "main.builddate=${BUILD_DATE}" -X "main.goversion=${GO_VERSION}"'

##
# ターゲット
##
.PHONY: help
.DEFAULT_GOAL := help
help:  ## このドキュメントを表示します。
	# ドキュメントを表示します。
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

test:  ## テストを実行します。
	# テストを実行します。
	go test -cover *.go

