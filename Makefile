.DEFAULT_GOAL := build

.PHONY:fmt vet build clean test
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

test: vet
	go test ./...

clean:
	go clean
