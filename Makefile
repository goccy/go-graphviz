export

CONTAINER_NAME := graphviz-wasm
IMAGE_NAME := graphviz-wasm
GOBIN := $(PWD)/bin
PATH := $(GOBIN):internal/tools/nori/bin:$(PATH)

.PHONY: tools
tools: nori
	go install github.com/bufbuild/buf/cmd/buf@v1.32.2

fmt/buf:
	buf format --write

generate/wasm: container/build
	$(eval CONTAINER_ID := $(shell docker create graphviz-wasm))
	docker cp "$(CONTAINER_ID):/work/graphviz.wasm" ./internal/wasm/graphviz.wasm

container/build:
	docker build ./internal/wasm/build -t $(IMAGE_NAME) --build-arg GRAPHVIZ_VERSION=$(shell cat graphviz.version)

container/prune:
	docker container prune

.PHONY: generate/buf
generate/buf:
	$(GOBIN)/buf generate
	mv bind.c internal/wasm/build
	mv bind.go internal/wasm/

.PHONY: nori
nori:
	make build -C ./internal/tools/nori
