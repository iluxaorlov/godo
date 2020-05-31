.PHONY: build
build:
	go build -v ./cmd/godo

.DEFAULT_GOAL := build