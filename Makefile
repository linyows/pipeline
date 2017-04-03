TEST ?= $(shell go list ./... | grep -v vendor)
COMMIT = $$(git describe --always)
NAME = "$(shell awk -F\" '/^const Name/ { print $$2; exit }' cmd/pipeline/version.go)"
VERSION = "$(shell awk -F\" '/^const Version/ { print $$2; exit }' cmd/pipeline/version.go)"

default: build

deps:
	go get -u github.com/golang/dep/...
	dep ensure

depsdev: deps
	go get github.com/mitchellh/gox
	go get github.com/tcnksm/ghr
	go get github.com/golang/lint/golint
	go get github.com/pierrre/gotestcover
	go get github.com/mattn/goveralls

build:
	go build -v ./cmd/pipeline

ci: depsdev test cover

test:
	go vet $(TEST)
	test -z "$(gofmt -s -l . 2>&1 | grep -v vendor | tee /dev/stderr)"
	go test -race $(TEST)

cover:
	gotestcover -v -covermode=count -coverprofile=coverage.out -parallelpackages=4 $(TEST)

bin: depsdev
	@sh -c "'$(CURDIR)/scripts/build.sh' $(NAME)"

dist: bin
	ghr v$(VERSION) pkg

.PHONY: default bin dist test deps
