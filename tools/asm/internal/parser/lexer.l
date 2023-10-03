%{
package parser

import (
	"fmt"
)

var (
    buf string
    curSymbol byte
)

const EOF = 0

func (l *lexer) Lex(lval *yySymType) int {
    c := curSymbol
%}

%yyc c
%yyn c = l.getNext()

OP_CODE [a-zA-Z]+
DEC_NUM [0-9]+
HEX_NUM 0x[0-9aAbBcCdDeEfF]+
BIN_NUM 0b[01]+

SYMBOLS [{};]

%%
    buf = buf[:0]

[ \t\n\r]+          // Ignore whitespace
;.*$

{OP_CODE}          lval.String = buf; return OP_CODE
{DEC_NUM}          lval.Number = mustParseNumber(buf, 10); return DEC_NUM
{HEX_NUM}          lval.Number = mustParseNumber(buf[2:len(buf)], 16); return HEX_NUM
{BIN_NUM}          lval.Number = mustParseNumber(buf[2:len(buf)], 2); return HEX_NUM
{SYMBOLS}          return int(buf[0])
\0                 return EOF
%%


panic(fmt.Sprintf("Unrecognized symbol: '%c' at %s", l.input[0],  l.codePointer.String()))
}

