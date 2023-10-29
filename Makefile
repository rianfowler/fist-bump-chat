BINARY_NAME=fist-bump-chat

.PHONY: all
all: test build

.PHONY: build
build:
	go build -o $(BINARY_NAME) cmd/fistbump/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)

.PHONY: run
run:
	go run cmd/fistbump/main.go

.PHONY: deps
deps:
	go get -v -t ./...

.PHONY: watch-run
watch-run:
	watchexec -e go,html -r -- make run

.PHONY: watch-test
watch-test:
	watchexec -e go,html -r -- make test

.PHONEY: list-todos
list-todos:
	@go list -f '{{.Dir}}' ./... | while read -r d; do grep -rn "TODO" $$d/*.go; done