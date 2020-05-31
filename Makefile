.PHONY: build
build:
	go build -v ./cmd/todo

.DEFAULT_GOAL := build