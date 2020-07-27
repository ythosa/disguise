.PHONY: build
build:
	go build -o disguise -v ./src/main.go

.PHONY: windows-build
windows-build:
	env GOOS=windows GOARCH=amd64 go build -o build/disguise_win64.exe -v ./src/main.go

.PHONY: linux-build
linux-build:
	env GOOS=linux GOARCH=amd64 go build -o build/disguise_lin64 -v ./src/main.go

.PHONY: run
run:
	go run ./src/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v ./src/...

.PHONY: pipeline
pipeline:
	make lint && make test && make

.DEFAULT_GOAL := build
