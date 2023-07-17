build:
	go build -o bin/godo

run: build
	@./bin/godo

test:
	go test -v ./...