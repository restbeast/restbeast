PROJECT_NAME= "restbeast"
PKG = "gitlab.com/restbeast/cli"
VERSION := $(CI_COMMIT_TAG)
CGO_ENABLED = 0
BUILD_CMD := CGO_ENABLED=0 go build -i -o $(PROJECT_NAME) -ldflags "-X 'main.version=$(VERSION)' -X 'main.sentryDsn=$(SENTRY_DSN)'" -v -buildmode=exe $(PKG)
PREFIX = /usr/local

define GITLAB_REQUEST_BODY
{
	"name": "release-$(VERSION)",
	"tag_name": "$(VERSION)",
	"ref": "$(CI_COMMIT_SHA)",
	"description": "release-$(VERSION)",
	"assets": {
		"links": [
			{ "name": "binary linux/386", "link_type": "package", "url": "https://cli-releases.restbeast.com/restbeast_$(CI_COMMIT_TAG)_linux_386.tar.gz" },
			{ "name": "binary linux/amd64", "link_type": "package", "url": "https://cli-releases.restbeast.com/restbeast_$(CI_COMMIT_TAG)_linux_amd64.tar.gz" },
			{ "name": "binary linux/arm64", "link_type": "package", "url": "https://cli-releases.restbeast.com/restbeast_$(CI_COMMIT_TAG)_linux_arm64.tar.gz" },
			{ "name": "binary darwin/amd64", "link_type": "package", "url": "https://cli-releases.restbeast.com/restbeast_$(CI_COMMIT_TAG)_darwin_amd64.tar.gz" },
			{ "name": "binary freebsd/amd64", "link_type": "package", "url": "https://cli-releases.restbeast.com/restbeast_$(CI_COMMIT_TAG)_freebsd_amd64.tar.gz" },
			{ "name": "binary freebsd/arm", "link_type": "package", "url": "https://cli-releases.restbeast.com/restbeast_$(CI_COMMIT_TAG)_freebsd_arm.tar.gz" }
		]
	}
}
endef

.PHONY: all dep build upload release clean linux linux-386 linux-amd64 linux-arm64 darwin darwin-amd64 freebsd freebsd-amd64 freebsd-arm

build: clean dep
	@$(BUILD_CMD)

dep: ## Get the dependencies
	@go get -d ./...

linux-386: clean dep ## Build linux/386
	@GOOS=linux GOARCH=386 $(BUILD_CMD)

pack-linux-386: linux-386
	@tar zcf restbeast_$(VERSION)_linux_386.tar.gz restbeast
	@rm -rf restbeast

linux-amd64: clean dep ## Build linux/amd64
	@GOOS=linux GOARCH=amd64 $(BUILD_CMD)

pack-linux-amd64: linux-amd64
	@tar zcf restbeast_$(VERSION)_linux_amd64.tar.gz restbeast
	@rm -rf restbeast

linux-arm64: clean dep ## Build linux/arm64
	@GOOS=linux GOARCH=arm64 $(BUILD_CMD)

pack-linux-arm64: linux-arm64
	@tar zcf restbeast_$(VERSION)_linux_arm64.tar.gz restbeast
	@rm -rf restbeast

linux: pack-linux-386 pack-linux-amd64 pack-linux-arm64 ## Build all linux

darwin-amd64: clean dep ## Build darwin amd64
	@GOOS=darwin GOARCH=amd64 $(BUILD_CMD)

pack-darwin-amd64: darwin-amd64
	@tar zcf restbeast_$(VERSION)_darwin_amd64.tar.gz restbeast
	@rm -rf restbeast

darwin: pack-darwin-amd64 ## Build all darwin

freebsd-amd64: clean dep ## Build freebsd/amd64
	@GOOS=freebsd GOARCH=amd64 $(BUILD_CMD)

pack-freebsd-amd64: freebsd-amd64
	@tar zcf restbeast_$(VERSION)_freebsd_amd64.tar.gz restbeast
	@rm -rf restbeast

freebsd-arm: clean dep ## Build freebsd/arm
	@GOOS=freebsd GOARCH=arm $(BUILD_CMD)

pack-freebsd-arm: freebsd-arm
	@tar zcf restbeast_$(VERSION)_freebsd_arm.tar.gz restbeast
	@rm -rf restbeast

freebsd: pack-freebsd-amd64 pack-freebsd-arm ## Build all freebsd

all: linux darwin freebsd ## Build all

upload: ## Upload binaries to S3
	@aws s3 cp ./ s3://restbeast-cli-release/ --region eu-central-1 --exclude "*" --include "restbeast_$(VERSION)_*.tar.gz" --recursive

gitlab-release: ## Release given tag in gitlab
	curl -i \
		 --header "content-type: application/json" --header "JOB-TOKEN: $(CI_JOB_TOKEN)" \
		 --data '$(GITLAB_REQUEST_BODY:\n=)' \
		 --request POST $(CI_SERVER_URL)/api/v4/projects/$(CI_PROJECT_ID)/releases

release: all upload gitlab-release

install:
	@cp restbeast $(PREFIX)/bin/

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
