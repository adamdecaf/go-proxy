build:
	go tool vet .
	go build

test: build
	go test -v ./...
