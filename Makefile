build:
	@go build -o bin/go-pos-test cmd/main.go

test:
	@go test -v ./...

run: build
	@bin/go-pos-test
