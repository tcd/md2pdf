.DEFAULT_GOAL := help

SHELL := /bin/bash
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

test:
	go test -v ./...

cmd:	
	@cd ./cmd/$(PROJECT_NAME) && go run *.go

all:
	@echo "Nothing to do for all"

build:
	@echo "Nothing to do for build"

clean:
	@echo "Nothing to do for clean"

help:	
	@echo
	@echo "  test – run 'go test' for the entire project"
	@echo


.PHONY: all build clean test help cmd
