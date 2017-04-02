TEST ?= $(shell go list ./... | grep -v vendor)

default: build

deps:
	go get -u github.com/golang/dep/...
	dep ensure
build:
	go build -v ./cmd/pipeline
test:
	go vet $(TEST)
	test -z "$(gofmt -s -l . 2>&1 | grep -v vendor | tee /dev/stderr)"
	go test -race $(TEST)
