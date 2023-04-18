PROJECT_NAME= "restbeast"
PKG = "github.com/restbeast/restbeast"
CGO_ENABLED = 0
SENTRY_DSN = "https://9c22d42ad3594172b24f9f5fc4db9d73@o421531.ingest.sentry.io/5341476"
BUILD_CMD := CGO_ENABLED=0 go build -o $(PROJECT_NAME) -ldflags "-X 'main.version=$(VERSION)' -X 'main.sentryDsn=$(SENTRY_DSN)'" -v -buildmode=exe $(PKG)
PREFIX = /usr/local

.PHONY: all dep build upload release clean linux linux-386 linux-amd64 linux-arm64 darwin darwin-amd64 freebsd freebsd-amd64 freebsd-arm clean test help

build: clean dep
	@$(BUILD_CMD)

dep: ## Get the dependencies
	@go get -d ./...

linux-386: clean dep ## Build linux/386
	@GOOS=linux GOARCH=386 $(BUILD_CMD)

linux-amd64: clean dep ## Build linux/amd64
	@GOOS=linux GOARCH=amd64 $(BUILD_CMD)

linux-arm64: clean dep ## Build linux/arm64
	@GOOS=linux GOARCH=arm64 $(BUILD_CMD)

linux: linux-386 linux-amd64 linux-arm64 ## Build all linux

darwin-amd64: clean dep ## Build darwin amd64
	@GOOS=darwin GOARCH=amd64 $(BUILD_CMD)

darwin-arm64: clean dep ## Build darwin arm64
	@GOOS=darwin GOARCH=arm64

darwin: darwin-amd64 darwin-arm64  ## Build all darwin

freebsd-amd64: clean dep ## Build freebsd/amd64
	@GOOS=freebsd GOARCH=amd64 $(BUILD_CMD)

freebsd-arm: clean dep ## Build freebsd/arm
	@GOOS=freebsd GOARCH=arm $(BUILD_CMD)

freebsd: freebsd-amd64 freebsd-arm ## Build all freebsd

all: linux darwin freebsd ## Build all

install:
	@rm -rf $(PREFIX)/bin/restbeast
	@cp restbeast $(PREFIX)/bin/

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

test:
	@go test -coverpkg=./... -coverprofile=profile.cov ./...
	@go tool cover -func profile.cov

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
