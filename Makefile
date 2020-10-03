

all: lint test

test:
	go test

bench:
	go test -bench=.

lint:
	go fmt ./...

clean:
	go clean -v
