GOPACKAGES?=$(shell find ./internal -name '*.go' -exec dirname {} \; | sort | uniq)
GOFILES?=$(shell find ./internal -name '*.go')

.PHONY: test fmt lint

test:
	go test $(GOPACKAGES)

fmt:
	go fmt $(GOPACKAGES)
	goimports -local github.com/tyz910/ray-tracer-challenge/ -ungroup -w ./cmd/ ./examples/ ./internal/

lint:
	ls $(GOFILES) | xargs -L1 golint
