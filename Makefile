

all: lint test

test:
	go test

lint:
	go fmt ./...

clean:
	go clean -v
