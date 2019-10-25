# By default the project should be build under GOPATH/src/github.com/<orgname>/<reponame>
GO_PACKAGE_ORG_NAME ?= $(shell basename $$(dirname $$PWD))
GO_PACKAGE_REPO_NAME ?= $(shell basename $$PWD)
GO_PACKAGE_PATH ?= github.com/${GO_PACKAGE_ORG_NAME}/${GO_PACKAGE_REPO_NAME}

GO111MODULE?=on
export GO111MODULE

GO_BINARY_DIR=$(OUT_DIR)/bin
GO_BINARY_NAME=operator-role-linter

.PHONY: build
## Build the linter
build: generate 
	@echo "building in ${GO_PACKAGE_PATH}"
	@-mkdir -p $(OUT_DIR)/bin
	$(Q)CGO_ENABLED=0 \
		go build ${V_FLAG} \
		-ldflags "-X ${GO_PACKAGE_PATH}/version.Commit=${GIT_COMMIT_ID} -X ${GO_PACKAGE_PATH}/version.BuildTime=${BUILD_TIME}" \
		-gcflags=-trimpath=$(GOPATH)/src \
		-asmflags=-trimpath=$(GOPATH)/src \
		-o $(GO_BINARY_DIR)/$(GO_BINARY_NAME) \
		cmd/main.go


.PHONY: install
## Build the linter
install: build 
	@echo "installing..."
	@cp $(GO_BINARY_DIR)/$(GO_BINARY_NAME) $(GOPATH)/bin
	@echo "executable is now available at $(GOPATH)/bin/$(GO_BINARY_NAME)"
	
