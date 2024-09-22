.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go build ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean
