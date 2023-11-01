build:
	@go build -o bin/go-data-access

run: build
	@./bin/go-data-access

test:
	@go test -v ./...
