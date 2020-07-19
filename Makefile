.PHONY: build
build:
	go build -v ./main.go

.PHONY: run
run:
	go run ./main.go

.DEFAULT_GOAL := build