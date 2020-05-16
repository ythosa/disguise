.PHONY: build
build:
	go build -o disguise.exe -v ./main.go 

.DEFAULT_GOAL := build