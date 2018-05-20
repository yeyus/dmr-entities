PACKAGE = dmr-entities
ENTRYPOINT = main.go

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

all: clean proto build ## clean and build for local environment

proto: ## recreate proto bindings for go
	protoc -I model/ model/entities.proto --go_out=plugins=grpc:model

build: ## build the go binary for the current environment
	@mkdir -p ./bin
	go build -i -o ./bin/$(PACKAGE)

build-linux-amd64: ## build the go binary for linux-amd64 systems
	@mkdir -p ./bin
	GOOS=linux GOARCH=amd64 go build -i -o ./bin/$(PACKAGE)-linux-amd64

container: build-linux-amd64 ## build the container for linux-amd64
	docker build -t yeyus/dmr-entities .

publish: container ## publish to docker hub
	docker push yeyus/dmr-entities:latest

clean: ## clean all files created by this makefile
	@rm -rf ./bin
	@rm model/entities.pb.go

run: ## run locally
	go run main.go
