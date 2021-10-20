SHELL = /bin/bash

REPO ?= gateway/blue-dashboard-backend

.PHONY: all
all: proto swg build

.PHONY: run
run:
	@./build/server

.PHONY: fmt
fmt:
	@CGO_ENABLED=0 go fmt ./...

.PHONY: vet
vet: fmt
	@CGO_ENABLED=0 go vet ./...

.PHONE: test
test: vet
	CGO_ENABLE=0 go test ./...

.PHONY: build
build: vet
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o build/server ./cmd/server/

.PHONY: proto
proto:
	@PWD=$(shell pwd)
	docker run --rm -i -v $$PWD/api:/go/api -v $$PWD/proto.sh:/go/proto.sh ankrnetwork/ankr-grpc /bin/bash proto.sh

.PHONY: docker-build
docker-build:
	@echo "build docker image"
	@BRANCH_NAME=$(shell git rev-parse --abbrev-ref HEAD); \
	if [[ $$BRANCH_NAME == "develop" ]]; then \
		ENV="stage"; \
	elif [[ $$BRANCH_NAME == "main" ]]; then \
		ENV="prod"; \
	else \
		ENV="feat"; \
	fi; \
	docker build -t $(REPO):$$ENV .

.PHONY: docker-push
docker-push: docker-build
	@echo "tag & push image"
	@BRANCH_NAME=$(shell git rev-parse --abbrev-ref HEAD); SHA1_SHORT=$(shell git rev-parse --short HEAD); \
	if [[ $$BRANCH_NAME == "develop" ]]; then \
		ENV="stage"; \
	elif [[ $$BRANCH_NAME == "master" ]]; then \
		ENV="prod"; \
	else \
		ENV="feat"; \
	fi;  \
	docker tag $(REPO):$$ENV  $(REPO):$$SHA1_SHORT; \
	docker push $(REPO):$$SHA1_SHORT; \
	docker push $(REPO):$$ENV;


.PHONY: clean
clean:
	@rm -f build/server

.PHONY: swg
swg:
	@swagger serve -F=swagger ./api/protos/index/public_index.swagger.json