# By default the project should be build under GOPATH/src/github.com/<orgname>/<reponame>
GO_PACKAGE_ORG_NAME ?= $(shell basename $$(dirname $$PWD))
GO_PACKAGE_REPO_NAME ?= $(shell basename $$PWD)
GO_PACKAGE_PATH ?= github.com/${GO_PACKAGE_ORG_NAME}/${GO_PACKAGE_REPO_NAME}

GO111MODULE?=on
export GO111MODULE

.PHONY: build
## Build the linter
build: generate $(OUT_DIR)/bin/linter

$(OUT_DIR)/linter:
	@echo "building in ${GO_PACKAGE_PATH}"
	@-mkdir -p $(OUT_DIR)/bin
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
		go build ${V_FLAG} \
		-ldflags "-X ${GO_PACKAGE_PATH}/version.Commit=${GIT_COMMIT_ID} -X ${GO_PACKAGE_PATH}/version.BuildTime=${BUILD_TIME}" \
		-gcflags=-trimpath=$(GOPATH)/src \
		-asmflags=-trimpath=$(GOPATH)/src \
		-o $(OUT_DIR)/bin/host-operator \
		cmd/manager/main.go

