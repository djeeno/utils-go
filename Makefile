.DEFAULT_GOAL := help
.PHONY: help
help:  ## display this doc
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'

.PHONY: tidy
tidy:
	go mod tidy && git diff --exit-code

.PHONY: lint
lint:  ## lint
	command -v golangci-lint || go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	golangci-lint run

.PHONY: test
test:  ## test
	go test -race -timeout 60s -cover -coverprofile=./coverage.txt ./...
	go tool cover -func=./coverage.txt

.PHONY: ci
ci: tidy lint test ## ci
