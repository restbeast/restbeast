PROJECT_NAME= "restbeast"
PKG = "gitlab.com/restbeast/cli"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
VERSION := $(CI_COMMIT_TAG)
BUILD_PAIRS = \
	linux:amd64 \
	linux:386 \
	linux:arm64 \
	freebsd:amd64 \
	darwin:amd64 \
	darwin:arm64
ASSET_LINKS := $(foreach PAIR,$(BUILD_PAIRS), \
	$(eval GOOS = $(word 1,$(subst :, ,$(PAIR)))) \
	$(eval GOARCH = $(word 2,$(subst :, ,$(PAIR)))) \
	{ \"name\": \"binary $(GOOS)/$(GOARCH)\", \"link_type\": \"package\",\"url\": \"https://cli-releases.restbeast.com/restbeast_$(CI_COMMIT_TAG)_$(GOOS)_$(GOARCH).tar.gz\" },\
)

ASSET_LINKS_JSON := $(shell echo $(ASSET_LINKS) | sed -e 's/,$$//')

.PHONY: build-all dep build upload-binaries release-all clean

dep: ## Get the dependencies
	@go get -v -d ./...

build: dep ## Build the binary file with args GOOS and GOARCH
	@GOOS=$$GOOS GOARCH=$$GOARCH go build -i -o $(PROJECT_NAME) -ldflags "-X 'main.version=$(VERSION)' -X 'main.sentryDsn=$(SENTRY_DSN)'" -v -buildmode=exe $(PKG)
	@tar zcvf restbeast_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz restbeast

build-all: dep ## Build all binary files
	@- $(foreach PAIR,$(BUILD_PAIRS), \
		$(eval GOOS = $(word 1,$(subst :, ,$(PAIR)))) \
		$(eval GOARCH = $(word 2,$(subst :, ,$(PAIR)))) \
		\
		GOOS=$$GOOS GOARCH=$$GOARCH go build -i -o $(PROJECT_NAME) -ldflags "-X 'main.version=$(VERSION)' -X 'main.sentryDsn=$(SENTRY_DSN)'" -v -buildmode=exe $(PKG) ; \
		tar zcf restbeast_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz restbeast ; \
		echo "restbeast_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz" ; \
	)

upload-binaries:
	@- $(foreach PAIR,$(BUILD_PAIRS), \
		$(eval GOOS = $(word 1,$(subst :, ,$(PAIR)))) \
		$(eval GOARCH = $(word 2,$(subst :, ,$(PAIR)))) \
		\
		aws s3 cp "restbeast_$(VERSION)_$(GOOS)_$(GOARCH).tar.gz" s3://restbeast-cli-release/ --region eu-central-1 ; \
	)

create-gitlab-release:
	curl -i \
		 --header "content-type: application/json" --header "JOB-TOKEN: $(CI_JOB_TOKEN)" \
         --data '{ "name": "release-$(VERSION)", "tag_name": "$(VERSION)", "ref": "$(CI_COMMIT_SHA)", "description": "release-$(VERSION)", "assets": { "links": [ $(ASSET_LINKS_JSON) ] } }' \
         --request POST $(CI_SERVER_URL)/api/v4/projects/$(CI_PROJECT_ID)/releases

release-all: build-all upload-binaries create-gitlab-release

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
