GO ?= go
SHELL := /bin/sh
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

build:
	@./scripts/build.sh

clean:
	$(GO) clean ./...
	rm -rf build

# Run all tests for the project.
test:
	@./scripts/test.sh

# Install md2pdf to $GOBIN.
install:
	@./scripts/install.sh

# Remove md2pdf from $GOBIN.
uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

.PHONY: build clean test install uninstall
