PROJECT_NAME= "restbeast"
PKG = "gitlab.com/restbeast/cli"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
VERSION := $(shell git describe --tags)

.PHONY: all dep build clean

all: build

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file
	GOOS=$$GOOS GOARCH=$$GOARCH go build -i -o $(PROJECT_NAME) -ldflags "-w -s -X main.Version=$(VERSION)" -v -buildmode=exe $(PKG)
	tar zcvf restbeast_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz restbeast

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
