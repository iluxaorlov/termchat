.PHONY build:
build:
	go build ./cmd/client
	go build ./cmd/server

.PHONY server:
server:
	go build ./cmd/server

.PHONY client:
client:
	go build ./cmd/client

.DEFAULT_GOAL := build
