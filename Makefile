.PHONY: build
build:
	go build -o disguise -v ./src/main.go

.PHONY: run
run:
	go run ./src/main.go

.DEFAULT_GOAL := build
