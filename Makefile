OS = Linux
VERSION = 1.0.0
MODULE = bfsm

# git commit id
COMMITID ?= latest
# Image URL to use all building/pushing image targets

ROOT_PACKAGE=github.com/icdb37/bfsm
CURDIR = $(shell pwd)
SOURCEDIR = $(CURDIR)
COVER = $($3)

ECHO = echo
RM = rm -rf
MKDIR = mkdir

.PHONY: test build

default: fmt vet tidy imports cilint

test:
	go test -cover=true $(PACKAGES)

race:
	go test -cover=true -race $(PACKAGES)

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./...

# https://godoc.org/golang.org/x/tools/cmd/goimports
imports:
	goimports -e -d -w -local $(ROOT_PACKAGE) ./

# https://github.com/golangci/golangci-lint/
# Install: go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.34.1
cilint:
	golangci-lint -c ./.golangci.yaml run ./...

tidy:
	go mod tidy

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./...

all: tidy imports fmt test

PACKAGES = $(shell go list ./... | grep -v './vendor/\|./tests\|./mock')
BUILD_PATH = $(shell if [ "$(CI_DEST_DIR)" != "" ]; then echo "$(CI_DEST_DIR)" ; else echo "$(PWD)"; fi)

go-generate:
	 go generate ./...

build-debug:
	@$(ECHO) "Will build on "$(BUILD_PATH)
	CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=on go build -a -gcflags='all=-N -l' -v -o $(BUILD_PATH)/bin/${MODULE} $(ROOT_PACKAGE)

build:
	@$(ECHO) "Will build on "$(BUILD_PATH)
	CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=on go build -a -ldflags "-w -s" -v -o $(BUILD_PATH)/bin/${MODULE} $(ROOT_PACKAGE)

build-amd64:
	@$(ECHO) "Will build on "$(BUILD_PATH)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -ldflags "-w -s" -v -o $(BUILD_PATH)/bin/linux/amd64/${MODULE} $(ROOT_PACKAGE)

build-arm64:
	@$(ECHO) "Will build "$(MODULE)" arm64 on "$(BUILD_PATH)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -a -ldflags "-w -s" -v -o $(BUILD_PATH)/bin/linux/arm64/${MODULE} $(ROOT_PACKAGE)

build-auto-arch:
	@$(ECHO) "Will build all services on "$(BUILD_PATH)
	@$(MAKE) build
	@$(MAKE) build-arm64

run: fmt imports cilint
	go run main.go serve

clean:
	rm -f *.out *.html
	rm -rf bin

check-all: tidy vet imports cilint

compile: test build

gen-grpc:
	rm -rf internal/proto
	buf generate