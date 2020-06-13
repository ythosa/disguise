.PHONY: build
build:
	go build -o disguise.exe -v ./main.go 

.PHONY: run
run:
	go run ./main.go

.DEFAULT_GOAL := build