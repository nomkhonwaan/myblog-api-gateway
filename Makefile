# Version control options
GIT 	 := git
VERSION  := $(shell $(GIT) describe --match 'v[0-9]*' --dirty='.m' --always --tags)
REVISION := $(shell $(GIT) rev-parse HEAD)$(shell if ! $(GIT) diff --no-ext-diff --quiet --exit-code; then echo .m; fi)

# Golang options
GO       ?= go
BINDATA  ?= bindata
MOCKGEN  ?= mockgen
PKG      := github.com/nomkhonwaan/myblog-api-gateway
TAGS     :=
LDFLAGS  :=
GOFLAGS  :=
BINDIR   := $(CURDIR)/bin

.DEFAULT_GOAL := run

.PHONY: all
all: clean install build
	
.PHONY: install
install:
	$(GO) mod download

.PHONY: clean
clean:
	rm -rf $(BINDIR)/* && \
	rm -rf $(CURDIR)/coverage.out && \
	rm -rf $(CURDIR)/vendor

.PHONY: run 
run: 
	$(GO) run cmd/main.go \
		serve \
		--listen-address=:8080

.PHONY: generate
generate:
	$(GO) generate ./...

.PHONY: test
test:
	$(GO) test -v ./... -race -coverprofile=coverage.out -covermode=atomic
	
.PHONY: bindata
bindata:
	$(BINDATA) -o ./pkg/data/data.go ./data/...

.PHONY: build
build:
	$(GO) build \
		$(GOFLAGS) \
		-tags '$(TAGS)' \
		-ldflags '-X main.Version=$(VERSION) -X main.Revision=$(REVISION) $(LDFLAGS)' \
		-o $(BINDIR)/myblog-api-gateway cmd/main.go