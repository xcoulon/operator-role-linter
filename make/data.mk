.PHONY: download
download:
	@echo "downloading API Resources assets..."
	@go run download/main.go \
	  --cluster-url=$(CLUSTER_URL) \
	  --token=$(TOKEN) \
	  --skippedGroups=toolchain.dev.openshift.com \
	  --output-dir=pkg/apiresources/data

.PHONY: generate
generate:
	@echo "generating API Resources assets..."
	@go install github.com/go-bindata/go-bindata
	@$(GOPATH)/bin/go-bindata \
	  -pkg data \
	  -o ./pkg/apiresources/data/apiresources_asset.go \
	  -nocompress \
	  -prefix pkg/apiresources/data \
	  pkg/apiresources/data
