export

GOBIN := $(PWD)/bin
PATH := $(GOBIN):$(PATH)
LDFLAGS := -w -s

.PHONY: tools
tools:
	go install github.com/bufbuild/buf/cmd/buf@v1.32.2

.PHONY: fmt
fmt: tidy fmt/buf

.PHONY: tidy
tidy:
	go mod tidy

fmt/buf:
	$(GOBIN)/buf format --write

.PHONY: generate
generate: build generate/buf

generate/buf:
	buf generate

.PHONY: build
build: build/protoc-gen-nori

build/protoc-gen-nori:
	go build -ldflags "$(LDFLAGS)" -o $(GOBIN)/protoc-gen-nori ./cmd/protoc-gen-nori
