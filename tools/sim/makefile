all: clean build test

clean:
	rm ./dd8sim -f

build:
	go build ./...

test:
	go test ./...

install:
	go install ./cmd/dd8sim/...
