SHELL := /bin/sh
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

run:	
	@cd ./cmd/$(PROJECT_NAME) && go run *.go

all:
	@echo "Nothing to do for all"

build:
	@echo "Nothing to do for build"

clean:
	go clean ./...

test:
	go test -v ./...


.PHONY: all build clean test run
