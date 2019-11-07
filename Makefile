GO_PACKAGE_BASE = github.com/yeyus/dmr-entities
PACKAGE = dmr-entities
ENTRYPOINT = main.go

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

all: clean proto build ## clean and build for local environment

proto: ## recreate proto bindings for go
	protoc -I api/protobuf api/protobuf/entities.proto --go_out=plugins=grpc:pkg/api

build: ## build the go binary for the current environment
	@mkdir -p ./bin
	CGO_ENABLED=0 go build -o ./bin/syncjob $(GO_PACKAGE_BASE)/cmd/syncjob

build-linux-amd64: ## build the go binary for linux-amd64 systems
	@mkdir -p ./bin
	env GOOD=linux GOARCH=amd64 go build -i -o ./bin/syncjob-linux-adm64 -v $(GO_PACKAGE_BASE)/cmd/syncjob

container:
	docker build -t yeyus/dmr-entities .

publish: container ## publish to docker hub
	docker push yeyus/dmr-entities:latest

clean: ## clean all files created by this makefile
	@rm -rf ./bin
	@rm pkg/api/entities.pb.go

run-syncjob: ## run locally
	go run $(GO_PACKAGE_BASE)/cmd/syncjob

run-server: ## run locally
	go run $(GO_PACKAGE_BASE)/cmd/server	
