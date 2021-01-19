MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
ROOT:=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: build
build:  ## build the tire cli binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o dist/tire cmd/tire/tire.go


.PHONY: install
install:  ## build and install the tire cli binary
	CGO_ENABLED=0 go install ${LDFLAGS}

.PHONY: version
version: ## display the version of the tire cli
	@echo $(VERSION)

.PHONY: goimport
goimport: ## run goimports on all files
	goimports -l $(GOFILES)

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)
