BINARY_NAME=fist-bump-chat

all: test build

build:
	go build -o $(BINARY_NAME) cmd/fistbump/main.go

test:
	go test -v ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

run:
	go run cmd/fistbump/main.go

deps:
	go get -v -t ./...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)_unix cmd/fistbump/main.go
