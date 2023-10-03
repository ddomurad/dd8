package parser

func ParseSource(src RawSourceFile) Source {
	inter := lexer{
		input: string(src.Source),
		codePointer: CodePointer{
			Uri:    string(src.Uri),
			Line:   1,
			Column: 1,
		},
	}

	inter.getNext()
	yyParse(&inter)

	return inter.source
}
