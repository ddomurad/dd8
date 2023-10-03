package parser

import (
	"fmt"
	"strconv"
)

type lexer struct {
	input string

	codePointer CodePointer
	source      Source
}

func (i *lexer) Error(e string) {
	panic(fmt.Sprintf("%v before %v:%v:%v", e, i.codePointer.Uri, i.codePointer.Line, i.codePointer.Column))
}

func (l *lexer) getNext() byte {
	if curSymbol != 0 {
		buf = fmt.Sprintf("%v%c", buf, curSymbol)
	}

	curSymbol = 0

	if len(l.input) == 0 {
		return 0
	}

	curSymbol = l.input[0]
	l.input = l.input[1:]

	l.codePointer.Column += 1
	if curSymbol == '\n' {
		l.codePointer.Line += 1
		l.codePointer.Column = 0
	}

	return curSymbol
}

func mustParseNumber(str string, base int) int64 {
	num, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		panic(err)
	}
	return num
}
