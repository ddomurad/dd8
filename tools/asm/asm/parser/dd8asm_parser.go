// Code generated from DD8ASM.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DD8ASM

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DD8ASMParser struct {
	*antlr.BaseParser
}

var DD8ASMParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dd8asmParserInit() {
	staticData := &DD8ASMParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'@'", "'{'", "'}'", "':='", "'['", "']'",
		"'~'", "'.l'", "'.h'", "'*'", "'/'", "'%'", "'+'", "'-'", "'<<'", "'>>'",
		"'&'", "'^'", "'|'", "':'", "'.def'", "'.tmpl'", "'.rep'", "'.org'",
		"'.inc'", "'.db'", "'.dw'", "'.byte'", "'.word'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "P_DEF", "P_TMPL", "P_REP", "P_ORG", "P_INC",
		"P_DB", "P_DW", "P_BYTE", "P_WORD", "REG", "STR", "HEX_NUM", "BIN_NUM",
		"DEC_NUM", "NAME", "COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "statement", "prep_instruction", "tmpl_block", "prep_def_args",
		"prep_def_arg_lines", "instruction", "arglist_p", "arglist", "arglist_lines",
		"argument", "expr", "str", "num", "reg", "namelist", "name", "label",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 41, 230, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 1, 3, 1, 51, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 87,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		111, 8, 2, 1, 2, 1, 2, 3, 2, 115, 8, 2, 3, 2, 117, 8, 2, 1, 3, 1, 3, 5,
		3, 121, 8, 3, 10, 3, 12, 3, 124, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 3, 5, 135, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 154, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8, 163, 8, 8,
		1, 9, 1, 9, 1, 9, 3, 9, 168, 8, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11,
		185, 8, 11, 3, 11, 187, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 5, 11, 209, 8, 11, 10, 11, 12, 11, 212, 9, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 223,
		8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 0, 1, 22, 18, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 0, 5, 1, 0, 13,
		15, 1, 0, 16, 17, 1, 0, 18, 19, 1, 0, 11, 12, 1, 0, 35, 37, 253, 0, 39,
		1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 116, 1, 0, 0, 0, 6, 118, 1, 0, 0, 0,
		8, 127, 1, 0, 0, 0, 10, 131, 1, 0, 0, 0, 12, 153, 1, 0, 0, 0, 14, 155,
		1, 0, 0, 0, 16, 159, 1, 0, 0, 0, 18, 164, 1, 0, 0, 0, 20, 169, 1, 0, 0,
		0, 22, 186, 1, 0, 0, 0, 24, 213, 1, 0, 0, 0, 26, 215, 1, 0, 0, 0, 28, 217,
		1, 0, 0, 0, 30, 219, 1, 0, 0, 0, 32, 224, 1, 0, 0, 0, 34, 226, 1, 0, 0,
		0, 36, 38, 3, 2, 1, 0, 37, 36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37,
		1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0,
		42, 43, 5, 0, 0, 1, 43, 1, 1, 0, 0, 0, 44, 51, 3, 34, 17, 0, 45, 48, 3,
		12, 6, 0, 46, 48, 3, 4, 2, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 47,
		48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 5, 40, 0, 0, 50, 44, 1, 0,
		0, 0, 50, 47, 1, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 53, 5, 27, 0, 0, 53, 117,
		3, 20, 10, 0, 54, 55, 5, 28, 0, 0, 55, 117, 3, 24, 12, 0, 56, 57, 5, 24,
		0, 0, 57, 58, 5, 1, 0, 0, 58, 59, 5, 40, 0, 0, 59, 60, 3, 10, 5, 0, 60,
		61, 5, 40, 0, 0, 61, 62, 5, 2, 0, 0, 62, 117, 1, 0, 0, 0, 63, 64, 5, 24,
		0, 0, 64, 117, 3, 8, 4, 0, 65, 66, 5, 25, 0, 0, 66, 67, 3, 32, 16, 0, 67,
		69, 5, 1, 0, 0, 68, 70, 3, 30, 15, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0,
		0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 5, 2, 0, 0, 72, 73, 3, 6, 3, 0, 73, 117,
		1, 0, 0, 0, 74, 75, 5, 26, 0, 0, 75, 76, 3, 32, 16, 0, 76, 77, 5, 3, 0,
		0, 77, 78, 3, 20, 10, 0, 78, 79, 5, 3, 0, 0, 79, 80, 3, 20, 10, 0, 80,
		81, 3, 6, 3, 0, 81, 117, 1, 0, 0, 0, 82, 83, 5, 4, 0, 0, 83, 84, 3, 32,
		16, 0, 84, 86, 5, 1, 0, 0, 85, 87, 3, 16, 8, 0, 86, 85, 1, 0, 0, 0, 86,
		87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5, 2, 0, 0, 89, 117, 1, 0,
		0, 0, 90, 91, 5, 29, 0, 0, 91, 117, 3, 16, 8, 0, 92, 93, 5, 29, 0, 0, 93,
		94, 5, 1, 0, 0, 94, 95, 5, 40, 0, 0, 95, 96, 3, 18, 9, 0, 96, 97, 5, 40,
		0, 0, 97, 98, 5, 2, 0, 0, 98, 117, 1, 0, 0, 0, 99, 100, 5, 30, 0, 0, 100,
		117, 3, 16, 8, 0, 101, 102, 5, 30, 0, 0, 102, 103, 5, 1, 0, 0, 103, 104,
		5, 40, 0, 0, 104, 105, 3, 18, 9, 0, 105, 106, 5, 40, 0, 0, 106, 107, 5,
		2, 0, 0, 107, 117, 1, 0, 0, 0, 108, 110, 5, 31, 0, 0, 109, 111, 3, 26,
		13, 0, 110, 109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 117, 1, 0, 0, 0,
		112, 114, 5, 32, 0, 0, 113, 115, 3, 26, 13, 0, 114, 113, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116, 52, 1, 0, 0, 0, 116, 54, 1,
		0, 0, 0, 116, 56, 1, 0, 0, 0, 116, 63, 1, 0, 0, 0, 116, 65, 1, 0, 0, 0,
		116, 74, 1, 0, 0, 0, 116, 82, 1, 0, 0, 0, 116, 90, 1, 0, 0, 0, 116, 92,
		1, 0, 0, 0, 116, 99, 1, 0, 0, 0, 116, 101, 1, 0, 0, 0, 116, 108, 1, 0,
		0, 0, 116, 112, 1, 0, 0, 0, 117, 5, 1, 0, 0, 0, 118, 122, 5, 5, 0, 0, 119,
		121, 3, 2, 1, 0, 120, 119, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120,
		1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 122, 1, 0,
		0, 0, 125, 126, 5, 6, 0, 0, 126, 7, 1, 0, 0, 0, 127, 128, 3, 32, 16, 0,
		128, 129, 5, 7, 0, 0, 129, 130, 3, 20, 10, 0, 130, 9, 1, 0, 0, 0, 131,
		134, 3, 8, 4, 0, 132, 133, 5, 40, 0, 0, 133, 135, 3, 10, 5, 0, 134, 132,
		1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 11, 1, 0, 0, 0, 136, 154, 3, 32,
		16, 0, 137, 138, 3, 32, 16, 0, 138, 139, 3, 14, 7, 0, 139, 154, 1, 0, 0,
		0, 140, 141, 3, 32, 16, 0, 141, 142, 3, 16, 8, 0, 142, 154, 1, 0, 0, 0,
		143, 144, 3, 32, 16, 0, 144, 145, 3, 14, 7, 0, 145, 146, 5, 3, 0, 0, 146,
		147, 3, 16, 8, 0, 147, 154, 1, 0, 0, 0, 148, 149, 3, 32, 16, 0, 149, 150,
		3, 16, 8, 0, 150, 151, 5, 3, 0, 0, 151, 152, 3, 14, 7, 0, 152, 154, 1,
		0, 0, 0, 153, 136, 1, 0, 0, 0, 153, 137, 1, 0, 0, 0, 153, 140, 1, 0, 0,
		0, 153, 143, 1, 0, 0, 0, 153, 148, 1, 0, 0, 0, 154, 13, 1, 0, 0, 0, 155,
		156, 5, 8, 0, 0, 156, 157, 3, 16, 8, 0, 157, 158, 5, 9, 0, 0, 158, 15,
		1, 0, 0, 0, 159, 162, 3, 20, 10, 0, 160, 161, 5, 3, 0, 0, 161, 163, 3,
		16, 8, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 17, 1, 0, 0,
		0, 164, 167, 3, 16, 8, 0, 165, 166, 5, 40, 0, 0, 166, 168, 3, 18, 9, 0,
		167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 19, 1, 0, 0, 0, 169, 170,
		3, 22, 11, 0, 170, 21, 1, 0, 0, 0, 171, 172, 6, 11, -1, 0, 172, 173, 5,
		10, 0, 0, 173, 187, 3, 22, 11, 11, 174, 175, 5, 17, 0, 0, 175, 187, 3,
		22, 11, 7, 176, 177, 5, 1, 0, 0, 177, 178, 3, 22, 11, 0, 178, 179, 5, 2,
		0, 0, 179, 187, 1, 0, 0, 0, 180, 185, 3, 26, 13, 0, 181, 185, 3, 32, 16,
		0, 182, 185, 3, 28, 14, 0, 183, 185, 3, 24, 12, 0, 184, 180, 1, 0, 0, 0,
		184, 181, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 183, 1, 0, 0, 0, 185,
		187, 1, 0, 0, 0, 186, 171, 1, 0, 0, 0, 186, 174, 1, 0, 0, 0, 186, 176,
		1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 210, 1, 0, 0, 0, 188, 189, 10, 9,
		0, 0, 189, 190, 7, 0, 0, 0, 190, 209, 3, 22, 11, 10, 191, 192, 10, 8, 0,
		0, 192, 193, 7, 1, 0, 0, 193, 209, 3, 22, 11, 9, 194, 195, 10, 6, 0, 0,
		195, 196, 7, 2, 0, 0, 196, 209, 3, 22, 11, 7, 197, 198, 10, 5, 0, 0, 198,
		199, 5, 20, 0, 0, 199, 209, 3, 22, 11, 6, 200, 201, 10, 4, 0, 0, 201, 202,
		5, 21, 0, 0, 202, 209, 3, 22, 11, 5, 203, 204, 10, 3, 0, 0, 204, 205, 5,
		22, 0, 0, 205, 209, 3, 22, 11, 4, 206, 207, 10, 10, 0, 0, 207, 209, 7,
		3, 0, 0, 208, 188, 1, 0, 0, 0, 208, 191, 1, 0, 0, 0, 208, 194, 1, 0, 0,
		0, 208, 197, 1, 0, 0, 0, 208, 200, 1, 0, 0, 0, 208, 203, 1, 0, 0, 0, 208,
		206, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211,
		1, 0, 0, 0, 211, 23, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214, 5, 34,
		0, 0, 214, 25, 1, 0, 0, 0, 215, 216, 7, 4, 0, 0, 216, 27, 1, 0, 0, 0, 217,
		218, 5, 33, 0, 0, 218, 29, 1, 0, 0, 0, 219, 222, 3, 32, 16, 0, 220, 221,
		5, 3, 0, 0, 221, 223, 3, 30, 15, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1,
		0, 0, 0, 223, 31, 1, 0, 0, 0, 224, 225, 5, 38, 0, 0, 225, 33, 1, 0, 0,
		0, 226, 227, 5, 38, 0, 0, 227, 228, 5, 23, 0, 0, 228, 35, 1, 0, 0, 0, 18,
		39, 47, 50, 69, 86, 110, 114, 116, 122, 134, 153, 162, 167, 184, 186, 208,
		210, 222,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DD8ASMParserInit initializes any static state used to implement DD8ASMParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDD8ASMParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DD8ASMParserInit() {
	staticData := &DD8ASMParserStaticData
	staticData.once.Do(dd8asmParserInit)
}

// NewDD8ASMParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDD8ASMParser(input antlr.TokenStream) *DD8ASMParser {
	DD8ASMParserInit()
	this := new(DD8ASMParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DD8ASMParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DD8ASM.g4"

	return this
}

// DD8ASMParser tokens.
const (
	DD8ASMParserEOF     = antlr.TokenEOF
	DD8ASMParserT__0    = 1
	DD8ASMParserT__1    = 2
	DD8ASMParserT__2    = 3
	DD8ASMParserT__3    = 4
	DD8ASMParserT__4    = 5
	DD8ASMParserT__5    = 6
	DD8ASMParserT__6    = 7
	DD8ASMParserT__7    = 8
	DD8ASMParserT__8    = 9
	DD8ASMParserT__9    = 10
	DD8ASMParserT__10   = 11
	DD8ASMParserT__11   = 12
	DD8ASMParserT__12   = 13
	DD8ASMParserT__13   = 14
	DD8ASMParserT__14   = 15
	DD8ASMParserT__15   = 16
	DD8ASMParserT__16   = 17
	DD8ASMParserT__17   = 18
	DD8ASMParserT__18   = 19
	DD8ASMParserT__19   = 20
	DD8ASMParserT__20   = 21
	DD8ASMParserT__21   = 22
	DD8ASMParserT__22   = 23
	DD8ASMParserP_DEF   = 24
	DD8ASMParserP_TMPL  = 25
	DD8ASMParserP_REP   = 26
	DD8ASMParserP_ORG   = 27
	DD8ASMParserP_INC   = 28
	DD8ASMParserP_DB    = 29
	DD8ASMParserP_DW    = 30
	DD8ASMParserP_BYTE  = 31
	DD8ASMParserP_WORD  = 32
	DD8ASMParserREG     = 33
	DD8ASMParserSTR     = 34
	DD8ASMParserHEX_NUM = 35
	DD8ASMParserBIN_NUM = 36
	DD8ASMParserDEC_NUM = 37
	DD8ASMParserNAME    = 38
	DD8ASMParserCOMMENT = 39
	DD8ASMParserEOL     = 40
	DD8ASMParserWS      = 41
)

// DD8ASMParser rules.
const (
	DD8ASMParserRULE_prog               = 0
	DD8ASMParserRULE_statement          = 1
	DD8ASMParserRULE_prep_instruction   = 2
	DD8ASMParserRULE_tmpl_block         = 3
	DD8ASMParserRULE_prep_def_args      = 4
	DD8ASMParserRULE_prep_def_arg_lines = 5
	DD8ASMParserRULE_instruction        = 6
	DD8ASMParserRULE_arglist_p          = 7
	DD8ASMParserRULE_arglist            = 8
	DD8ASMParserRULE_arglist_lines      = 9
	DD8ASMParserRULE_argument           = 10
	DD8ASMParserRULE_expr               = 11
	DD8ASMParserRULE_str                = 12
	DD8ASMParserRULE_num                = 13
	DD8ASMParserRULE_reg                = 14
	DD8ASMParserRULE_namelist           = 15
	DD8ASMParserRULE_name               = 16
	DD8ASMParserRULE_label              = 17
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserEOF, 0)
}

func (s *ProgContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DD8ASMParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1382962692112) != 0 {
		{
			p.SetState(36)
			p.Statement()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(DD8ASMParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Label() ILabelContext
	EOL() antlr.TerminalNode
	Instruction() IInstructionContext
	Prep_instruction() IPrep_instructionContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *StatementContext) EOL() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserEOL, 0)
}

func (s *StatementContext) Instruction() IInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *StatementContext) Prep_instruction() IPrep_instructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrep_instructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrep_instructionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DD8ASMParserRULE_statement)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Label()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		switch p.GetTokenStream().LA(1) {
		case DD8ASMParserNAME:
			{
				p.SetState(45)
				p.Instruction()
			}

		case DD8ASMParserT__3, DD8ASMParserP_DEF, DD8ASMParserP_TMPL, DD8ASMParserP_REP, DD8ASMParserP_ORG, DD8ASMParserP_INC, DD8ASMParserP_DB, DD8ASMParserP_DW, DD8ASMParserP_BYTE, DD8ASMParserP_WORD:
			{
				p.SetState(46)
				p.Prep_instruction()
			}

		case DD8ASMParserEOL:

		default:
		}
		{
			p.SetState(49)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrep_instructionContext is an interface to support dynamic dispatch.
type IPrep_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	P_ORG() antlr.TerminalNode
	AllArgument() []IArgumentContext
	Argument(i int) IArgumentContext
	P_INC() antlr.TerminalNode
	Str() IStrContext
	P_DEF() antlr.TerminalNode
	AllEOL() []antlr.TerminalNode
	EOL(i int) antlr.TerminalNode
	Prep_def_arg_lines() IPrep_def_arg_linesContext
	Prep_def_args() IPrep_def_argsContext
	P_TMPL() antlr.TerminalNode
	Name() INameContext
	Tmpl_block() ITmpl_blockContext
	Namelist() INamelistContext
	P_REP() antlr.TerminalNode
	Arglist() IArglistContext
	P_DB() antlr.TerminalNode
	Arglist_lines() IArglist_linesContext
	P_DW() antlr.TerminalNode
	P_BYTE() antlr.TerminalNode
	Num() INumContext
	P_WORD() antlr.TerminalNode

	// IsPrep_instructionContext differentiates from other interfaces.
	IsPrep_instructionContext()
}

type Prep_instructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrep_instructionContext() *Prep_instructionContext {
	var p = new(Prep_instructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prep_instruction
	return p
}

func InitEmptyPrep_instructionContext(p *Prep_instructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prep_instruction
}

func (*Prep_instructionContext) IsPrep_instructionContext() {}

func NewPrep_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prep_instructionContext {
	var p = new(Prep_instructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_prep_instruction

	return p
}

func (s *Prep_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Prep_instructionContext) P_ORG() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_ORG, 0)
}

func (s *Prep_instructionContext) AllArgument() []IArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentContext); ok {
			len++
		}
	}

	tst := make([]IArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentContext); ok {
			tst[i] = t.(IArgumentContext)
			i++
		}
	}

	return tst
}

func (s *Prep_instructionContext) Argument(i int) IArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *Prep_instructionContext) P_INC() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_INC, 0)
}

func (s *Prep_instructionContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *Prep_instructionContext) P_DEF() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_DEF, 0)
}

func (s *Prep_instructionContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(DD8ASMParserEOL)
}

func (s *Prep_instructionContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(DD8ASMParserEOL, i)
}

func (s *Prep_instructionContext) Prep_def_arg_lines() IPrep_def_arg_linesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrep_def_arg_linesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrep_def_arg_linesContext)
}

func (s *Prep_instructionContext) Prep_def_args() IPrep_def_argsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrep_def_argsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrep_def_argsContext)
}

func (s *Prep_instructionContext) P_TMPL() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_TMPL, 0)
}

func (s *Prep_instructionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Prep_instructionContext) Tmpl_block() ITmpl_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITmpl_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITmpl_blockContext)
}

func (s *Prep_instructionContext) Namelist() INamelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *Prep_instructionContext) P_REP() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_REP, 0)
}

func (s *Prep_instructionContext) Arglist() IArglistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Prep_instructionContext) P_DB() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_DB, 0)
}

func (s *Prep_instructionContext) Arglist_lines() IArglist_linesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglist_linesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglist_linesContext)
}

func (s *Prep_instructionContext) P_DW() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_DW, 0)
}

func (s *Prep_instructionContext) P_BYTE() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_BYTE, 0)
}

func (s *Prep_instructionContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *Prep_instructionContext) P_WORD() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserP_WORD, 0)
}

func (s *Prep_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prep_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prep_instructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitPrep_instruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Prep_instruction() (localctx IPrep_instructionContext) {
	localctx = NewPrep_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DD8ASMParserRULE_prep_instruction)
	var _la int

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(DD8ASMParserP_ORG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Argument()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(DD8ASMParserP_INC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.Str()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Match(DD8ASMParserP_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Prep_def_arg_lines()
		}
		{
			p.SetState(60)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.Match(DD8ASMParserP_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Prep_def_args()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(65)
			p.Match(DD8ASMParserP_TMPL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Name()
		}
		{
			p.SetState(67)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == DD8ASMParserNAME {
			{
				p.SetState(68)
				p.Namelist()
			}

		}
		{
			p.SetState(71)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Tmpl_block()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(74)
			p.Match(DD8ASMParserP_REP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Name()
		}
		{
			p.SetState(76)
			p.Match(DD8ASMParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Argument()
		}
		{
			p.SetState(78)
			p.Match(DD8ASMParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Argument()
		}
		{
			p.SetState(80)
			p.Tmpl_block()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(82)
			p.Match(DD8ASMParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Name()
		}
		{
			p.SetState(84)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541166011394) != 0 {
			{
				p.SetState(85)
				p.Arglist()
			}

		}
		{
			p.SetState(88)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(90)
			p.Match(DD8ASMParserP_DB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Arglist()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(92)
			p.Match(DD8ASMParserP_DB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Arglist_lines()
		}
		{
			p.SetState(96)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(99)
			p.Match(DD8ASMParserP_DW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Arglist()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(101)
			p.Match(DD8ASMParserP_DW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Arglist_lines()
		}
		{
			p.SetState(105)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(108)
			p.Match(DD8ASMParserP_BYTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168576) != 0 {
			{
				p.SetState(109)
				p.Num()
			}

		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(112)
			p.Match(DD8ASMParserP_WORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168576) != 0 {
			{
				p.SetState(113)
				p.Num()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITmpl_blockContext is an interface to support dynamic dispatch.
type ITmpl_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsTmpl_blockContext differentiates from other interfaces.
	IsTmpl_blockContext()
}

type Tmpl_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmpl_blockContext() *Tmpl_blockContext {
	var p = new(Tmpl_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_tmpl_block
	return p
}

func InitEmptyTmpl_blockContext(p *Tmpl_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_tmpl_block
}

func (*Tmpl_blockContext) IsTmpl_blockContext() {}

func NewTmpl_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tmpl_blockContext {
	var p = new(Tmpl_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_tmpl_block

	return p
}

func (s *Tmpl_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Tmpl_blockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Tmpl_blockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Tmpl_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tmpl_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tmpl_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitTmpl_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Tmpl_block() (localctx ITmpl_blockContext) {
	localctx = NewTmpl_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DD8ASMParserRULE_tmpl_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(DD8ASMParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1382962692112) != 0 {
		{
			p.SetState(119)
			p.Statement()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(125)
		p.Match(DD8ASMParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrep_def_argsContext is an interface to support dynamic dispatch.
type IPrep_def_argsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Argument() IArgumentContext

	// IsPrep_def_argsContext differentiates from other interfaces.
	IsPrep_def_argsContext()
}

type Prep_def_argsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrep_def_argsContext() *Prep_def_argsContext {
	var p = new(Prep_def_argsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prep_def_args
	return p
}

func InitEmptyPrep_def_argsContext(p *Prep_def_argsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prep_def_args
}

func (*Prep_def_argsContext) IsPrep_def_argsContext() {}

func NewPrep_def_argsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prep_def_argsContext {
	var p = new(Prep_def_argsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_prep_def_args

	return p
}

func (s *Prep_def_argsContext) GetParser() antlr.Parser { return s.parser }

func (s *Prep_def_argsContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Prep_def_argsContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *Prep_def_argsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prep_def_argsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prep_def_argsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitPrep_def_args(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Prep_def_args() (localctx IPrep_def_argsContext) {
	localctx = NewPrep_def_argsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DD8ASMParserRULE_prep_def_args)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Name()
	}
	{
		p.SetState(128)
		p.Match(DD8ASMParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Argument()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrep_def_arg_linesContext is an interface to support dynamic dispatch.
type IPrep_def_arg_linesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prep_def_args() IPrep_def_argsContext
	EOL() antlr.TerminalNode
	Prep_def_arg_lines() IPrep_def_arg_linesContext

	// IsPrep_def_arg_linesContext differentiates from other interfaces.
	IsPrep_def_arg_linesContext()
}

type Prep_def_arg_linesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrep_def_arg_linesContext() *Prep_def_arg_linesContext {
	var p = new(Prep_def_arg_linesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prep_def_arg_lines
	return p
}

func InitEmptyPrep_def_arg_linesContext(p *Prep_def_arg_linesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_prep_def_arg_lines
}

func (*Prep_def_arg_linesContext) IsPrep_def_arg_linesContext() {}

func NewPrep_def_arg_linesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prep_def_arg_linesContext {
	var p = new(Prep_def_arg_linesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_prep_def_arg_lines

	return p
}

func (s *Prep_def_arg_linesContext) GetParser() antlr.Parser { return s.parser }

func (s *Prep_def_arg_linesContext) Prep_def_args() IPrep_def_argsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrep_def_argsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrep_def_argsContext)
}

func (s *Prep_def_arg_linesContext) EOL() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserEOL, 0)
}

func (s *Prep_def_arg_linesContext) Prep_def_arg_lines() IPrep_def_arg_linesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrep_def_arg_linesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrep_def_arg_linesContext)
}

func (s *Prep_def_arg_linesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prep_def_arg_linesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prep_def_arg_linesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitPrep_def_arg_lines(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Prep_def_arg_lines() (localctx IPrep_def_arg_linesContext) {
	localctx = NewPrep_def_arg_linesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DD8ASMParserRULE_prep_def_arg_lines)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Prep_def_args()
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(132)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Prep_def_arg_lines()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Arglist_p() IArglist_pContext
	Arglist() IArglistContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InstructionContext) Arglist_p() IArglist_pContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglist_pContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglist_pContext)
}

func (s *InstructionContext) Arglist() IArglistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DD8ASMParserRULE_instruction)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Name()
		}
		{
			p.SetState(138)
			p.Arglist_p()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.Name()
		}
		{
			p.SetState(141)
			p.Arglist()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(143)
			p.Name()
		}
		{
			p.SetState(144)
			p.Arglist_p()
		}
		{
			p.SetState(145)
			p.Match(DD8ASMParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Arglist()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(148)
			p.Name()
		}
		{
			p.SetState(149)
			p.Arglist()
		}
		{
			p.SetState(150)
			p.Match(DD8ASMParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Arglist_p()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArglist_pContext is an interface to support dynamic dispatch.
type IArglist_pContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Arglist() IArglistContext

	// IsArglist_pContext differentiates from other interfaces.
	IsArglist_pContext()
}

type Arglist_pContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArglist_pContext() *Arglist_pContext {
	var p = new(Arglist_pContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_arglist_p
	return p
}

func InitEmptyArglist_pContext(p *Arglist_pContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_arglist_p
}

func (*Arglist_pContext) IsArglist_pContext() {}

func NewArglist_pContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arglist_pContext {
	var p = new(Arglist_pContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_arglist_p

	return p
}

func (s *Arglist_pContext) GetParser() antlr.Parser { return s.parser }

func (s *Arglist_pContext) Arglist() IArglistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Arglist_pContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arglist_pContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arglist_pContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitArglist_p(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Arglist_p() (localctx IArglist_pContext) {
	localctx = NewArglist_pContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DD8ASMParserRULE_arglist_p)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(DD8ASMParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Arglist()
	}
	{
		p.SetState(157)
		p.Match(DD8ASMParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArglistContext is an interface to support dynamic dispatch.
type IArglistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Argument() IArgumentContext
	Arglist() IArglistContext

	// IsArglistContext differentiates from other interfaces.
	IsArglistContext()
}

type ArglistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArglistContext() *ArglistContext {
	var p = new(ArglistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_arglist
	return p
}

func InitEmptyArglistContext(p *ArglistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_arglist
}

func (*ArglistContext) IsArglistContext() {}

func NewArglistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArglistContext {
	var p = new(ArglistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_arglist

	return p
}

func (s *ArglistContext) GetParser() antlr.Parser { return s.parser }

func (s *ArglistContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArglistContext) Arglist() IArglistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *ArglistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArglistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArglistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitArglist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Arglist() (localctx IArglistContext) {
	localctx = NewArglistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DD8ASMParserRULE_arglist)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Argument()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(160)
			p.Match(DD8ASMParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Arglist()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArglist_linesContext is an interface to support dynamic dispatch.
type IArglist_linesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Arglist() IArglistContext
	EOL() antlr.TerminalNode
	Arglist_lines() IArglist_linesContext

	// IsArglist_linesContext differentiates from other interfaces.
	IsArglist_linesContext()
}

type Arglist_linesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArglist_linesContext() *Arglist_linesContext {
	var p = new(Arglist_linesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_arglist_lines
	return p
}

func InitEmptyArglist_linesContext(p *Arglist_linesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_arglist_lines
}

func (*Arglist_linesContext) IsArglist_linesContext() {}

func NewArglist_linesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arglist_linesContext {
	var p = new(Arglist_linesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_arglist_lines

	return p
}

func (s *Arglist_linesContext) GetParser() antlr.Parser { return s.parser }

func (s *Arglist_linesContext) Arglist() IArglistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglistContext)
}

func (s *Arglist_linesContext) EOL() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserEOL, 0)
}

func (s *Arglist_linesContext) Arglist_lines() IArglist_linesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArglist_linesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArglist_linesContext)
}

func (s *Arglist_linesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arglist_linesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arglist_linesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitArglist_lines(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Arglist_lines() (localctx IArglist_linesContext) {
	localctx = NewArglist_linesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DD8ASMParserRULE_arglist_lines)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Arglist()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(165)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Arglist_lines()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DD8ASMParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Num() INumContext
	Name() INameContext
	Reg() IRegContext
	Str() IStrContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Num() INumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *ExprContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExprContext) Reg() IRegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *ExprContext) Str() IStrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *DD8ASMParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, DD8ASMParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DD8ASMParserT__9:
		{
			p.SetState(172)
			p.Match(DD8ASMParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.expr(11)
		}

	case DD8ASMParserT__16:
		{
			p.SetState(174)
			p.Match(DD8ASMParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.expr(7)
		}

	case DD8ASMParserT__0:
		{
			p.SetState(176)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.expr(0)
		}
		{
			p.SetState(178)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DD8ASMParserREG, DD8ASMParserSTR, DD8ASMParserHEX_NUM, DD8ASMParserBIN_NUM, DD8ASMParserDEC_NUM, DD8ASMParserNAME:
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case DD8ASMParserHEX_NUM, DD8ASMParserBIN_NUM, DD8ASMParserDEC_NUM:
			{
				p.SetState(180)
				p.Num()
			}

		case DD8ASMParserNAME:
			{
				p.SetState(181)
				p.Name()
			}

		case DD8ASMParserREG:
			{
				p.SetState(182)
				p.Reg()
			}

		case DD8ASMParserSTR:
			{
				p.SetState(183)
				p.Str()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(189)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57344) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(190)
					p.expr(10)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(191)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(192)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DD8ASMParserT__15 || _la == DD8ASMParserT__16) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(193)
					p.expr(9)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(195)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DD8ASMParserT__17 || _la == DD8ASMParserT__18) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(196)
					p.expr(7)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(198)
					p.Match(DD8ASMParserT__19)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(199)
					p.expr(6)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(201)
					p.Match(DD8ASMParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(202)
					p.expr(5)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(204)
					p.Match(DD8ASMParserT__21)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(205)
					p.expr(4)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(207)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DD8ASMParserT__10 || _la == DD8ASMParserT__11) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR() antlr.TerminalNode

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_str
	return p
}

func InitEmptyStrContext(p *StrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_str
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) STR() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserSTR, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitStr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DD8ASMParserRULE_str)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(DD8ASMParserSTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEC_NUM() antlr.TerminalNode
	HEX_NUM() antlr.TerminalNode
	BIN_NUM() antlr.TerminalNode

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_num
	return p
}

func InitEmptyNumContext(p *NumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_num
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) DEC_NUM() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserDEC_NUM, 0)
}

func (s *NumContext) HEX_NUM() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserHEX_NUM, 0)
}

func (s *NumContext) BIN_NUM() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserBIN_NUM, 0)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DD8ASMParserRULE_num)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168576) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegContext is an interface to support dynamic dispatch.
type IRegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REG() antlr.TerminalNode

	// IsRegContext differentiates from other interfaces.
	IsRegContext()
}

type RegContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegContext() *RegContext {
	var p = new(RegContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_reg
	return p
}

func InitEmptyRegContext(p *RegContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_reg
}

func (*RegContext) IsRegContext() {}

func NewRegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegContext {
	var p = new(RegContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_reg

	return p
}

func (s *RegContext) GetParser() antlr.Parser { return s.parser }

func (s *RegContext) REG() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserREG, 0)
}

func (s *RegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitReg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Reg() (localctx IRegContext) {
	localctx = NewRegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DD8ASMParserRULE_reg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(DD8ASMParserREG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamelistContext is an interface to support dynamic dispatch.
type INamelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Namelist() INamelistContext

	// IsNamelistContext differentiates from other interfaces.
	IsNamelistContext()
}

type NamelistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamelistContext() *NamelistContext {
	var p = new(NamelistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_namelist
	return p
}

func InitEmptyNamelistContext(p *NamelistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_namelist
}

func (*NamelistContext) IsNamelistContext() {}

func NewNamelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamelistContext {
	var p = new(NamelistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_namelist

	return p
}

func (s *NamelistContext) GetParser() antlr.Parser { return s.parser }

func (s *NamelistContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NamelistContext) Namelist() INamelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamelistContext)
}

func (s *NamelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamelistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitNamelist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Namelist() (localctx INamelistContext) {
	localctx = NewNamelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DD8ASMParserRULE_namelist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Name()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DD8ASMParserT__2 {
		{
			p.SetState(220)
			p.Match(DD8ASMParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Namelist()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DD8ASMParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(DD8ASMParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAME() antlr.TerminalNode

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DD8ASMParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DD8ASMParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) NAME() antlr.TerminalNode {
	return s.GetToken(DD8ASMParserNAME, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DD8ASMVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DD8ASMParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DD8ASMParserRULE_label)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(DD8ASMParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Match(DD8ASMParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *DD8ASMParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DD8ASMParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
