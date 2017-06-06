TEST ?= $(shell go list ./... | grep -v vendor)
COMMIT = $$(git describe --always)
NAME = "$(shell awk -F\" '/^const Name/ { print $$2; exit }' cmd/pipeline/version.go)"
VERSION = "$(shell awk -F\" '/^const Version/ { print $$2; exit }' cmd/pipeline/version.go)"

INFO_COLOR=\033[1;34m
RESET=\033[0m
BOLD=\033[1m

default: build

deps: ## Installing dependencies
	go get -u github.com/golang/dep/...
	dep ensure

depsdev: deps ## Installing dependencies with development
	go get github.com/mitchellh/gox
	go get github.com/tcnksm/ghr
	go get github.com/golang/lint/golint
	go get github.com/pierrre/gotestcover
	go get github.com/mattn/goveralls

build: ## Build for bin
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Building$(RESET)"
	go build -v ./cmd/pipeline

ci: depsdev vet lint test cover ## Run test and more...

test: ## Run Test
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Testing$(RESET)"
	go test -v $(TEST) -timeout=30s -parallel=4
	go test -race $(TEST)

vet: ## Exec go vet
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Vetting$(RESET)"
	go vet $(TEST)

lint: ## Exec golint
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Linting$(RESET)"
	golint -set_exit_status $(TEST)

cover: ## Exec gotestcover
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)k$(RESET)"
	gotestcover -v -covermode=count -coverprofile=coverage.out -parallelpackages=4 $(TEST)

bin: depsdev
	@sh -c "'$(CURDIR)/scripts/build.sh' $(NAME)"

ghr: ## Upload to Github releases without token check
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Releasing for Github$(RESET)"
	ghr v$(VERSION) pkg

dist: bin ## Upload to Github releases without token check
	@test -z $(GITHUB_TOKEN) || test -z $(GITHUB_API) || $(MAKE) ghr

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(INFO_COLOR)%-30s$(RESET) %s\n", $$1, $$2}'

.PHONY: default bin dist test deps
