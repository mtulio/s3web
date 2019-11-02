# ###############
# Defauts
export PATH := $(PWD)/bin:$(PATH)
export GO111MODULE := on

APP_NAME ?= s3web-app

# Build envs
GOOS := linux
GOARCH := amd64

GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_DESCRIBE := $(shell git describe --tags --always)
VERSION := $(shell cat ./cmd/$(APP_NAME)/VERSION 2>/dev/null || echo "N/A")
BIN_NAME := $(PWD)/bin/$(APP_NAME)

LDFLAGS :=
LDFLAGS += -X main.VersionCommit=$(GIT_COMMIT)
LDFLAGS += -X main.VersionTag=$(GIT_DESCRIBE)
LDFLAGS += -X main.VersionFull=$(VERSION)
LDFLAGS += -X main.VersionEnv=$(ENV)
LDFLAGS += -X main.AppName=$(APP_NAME)

# ###############
# Main targets
# Build a version
.PHONY: builder
builder: clean module
	@test -d $(PWD)/bin || mkdir $(PWD)/bin
	cd cmd/$(APP_NAME); \
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) \
		go build \
		-ldflags "$(LDFLAGS)" \
		$(BUILD_TAGS) \
		-o $(BIN_NAME) && strip $(BIN_NAME)

.PHONY: build
build: builder bin-ls

.PHONY: clean
clean:
	@rm -f bin/$(BIN_NAME)

.PHONY: module
module:
	@test -f go.mod || go mod init

.PHONY: bin-ls
bin-ls:
	@echo -e "\n#> Bin file: "
	@ls bin/$(APP_NAME)
