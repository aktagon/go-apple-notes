.PHONY: build test clean run-example run-cli

build:
	go build -o bin/notes-cli ./cmd/notes-cli

test:
	go test ./...

clean:
	rm -rf bin/

run-example:
	go run ./examples/basic

run-cli:
	go run ./cmd/notes-cli

install:
	go install ./cmd/notes-cli

fmt:
	go fmt ./...

vet:
	go vet ./...

mod-tidy:
	go mod tidy