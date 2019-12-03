# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY=justicia
BINARY_UNIX=$(BINARY)_unix
PLATFORMS := windows linux darwin
os = $(word 1, $@)

VERSION ?= $(shell git describe master)
VERSION ?= vlatest
DATE=$(shell date "+(%d %B %Y)")

.PHONY: all

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin

all:
	@echo " make <cmd>"
	@echo ""
	@echo "commands:"
	@echo " build           - runs go build"
	@echo " build_version   - runs go build with ldflags version=${VERSION} & date=${DATE}"
	@echo " install_version - runs go install with ldflags version=${VERSION} & date=${DATE}"
	@echo ""

build: clean
	@go build -v -o ${BINARY}

build_version: clean
	@go build -v -ldflags='-s -w -X "main.version=${VERSION}" -X "main.date=${DATE}"' -o ${BINARY}

install_version:
	@go install -v -ldflags='-s -w -X "main.version=${VERSION}" -X "main.date=${DATE}"'

clean:
	@rm -f ${BINARY}

run:
		$(GOBUILD) -o $(BINARY) -v
		./$(BINARY)

check_version:
	@if [ -a "${BINARY}" ]; then \
		echo "${BINARY} already exists"; \
		exit 1; \
	fi;

# Cross compilation
build-linux:
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
