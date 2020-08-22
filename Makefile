.PHONY: all build test

build:
	mkdir -p bin
	CGO_ENABLED=0 go build -o bin/data-processor

test:
	go test -race ./...
