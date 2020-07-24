.PHONY: build
build:
	go build -o disguise -v ./src/main.go

.PHONY: run
run:
	go run ./src/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v ./src/...

.DEFAULT_GOAL := build
