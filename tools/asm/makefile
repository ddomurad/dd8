all: clean gen_antlr build test

clean:
	rm ./pkg/parser/* -rf

build:
	go build ./...

test:
	go test ./...

install:
	go install ./cmd/dd8asm/...

gen_antlr:
	(cd ./asm/ && antlr4 -no-listener -visitor -Dlanguage=Go -o parser DD8ASM.g4)
