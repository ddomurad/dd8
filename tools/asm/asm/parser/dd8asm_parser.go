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
		"", "'('", "')'", "'@'", "'{'", "'}'", "':='", "','", "'['", "']'",
		"'~'", "'.l'", "'.h'", "'*'", "'/'", "'%'", "'+'", "'-'", "'<<'", "'>>'",
		"'&'", "'^'", "'|'", "':'", "'.def'", "'.tmpl'", "'.org'", "'.inc'",
		"'.db'", "'.dw'", "'.byte'", "'.word'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "P_DEF", "P_TMPL", "P_ORG", "P_INC", "P_DB",
		"P_DW", "P_BYTE", "P_WORD", "REG", "STR", "HEX_NUM", "BIN_NUM", "DEC_NUM",
		"NAME", "COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "statement", "prep_instruction", "tmpl_block", "prep_def_args",
		"prep_def_arg_lines", "instruction", "arglist_p", "arglist", "arglist_lines",
		"argument", "expr", "str", "num", "reg", "namelist", "name", "label",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 222, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 1, 3, 1, 51, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 79, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 103, 8, 2, 1, 2, 1, 2, 3, 2, 107, 8, 2, 3, 2, 109,
		8, 2, 1, 3, 1, 3, 5, 3, 113, 8, 3, 10, 3, 12, 3, 116, 9, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 127, 8, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 146, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 3, 8, 155, 8, 8, 1, 9, 1, 9, 1, 9, 3, 9, 160, 8, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 177, 8, 11, 3, 11, 179, 8, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 201, 8, 11, 10,
		11, 12, 11, 204, 9, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 3, 15, 215, 8, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 0, 1, 22, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 0, 5, 1, 0, 13, 15, 1, 0, 16, 17, 1, 0, 18, 19, 1, 0, 11, 12,
		1, 0, 34, 36, 244, 0, 39, 1, 0, 0, 0, 2, 50, 1, 0, 0, 0, 4, 108, 1, 0,
		0, 0, 6, 110, 1, 0, 0, 0, 8, 119, 1, 0, 0, 0, 10, 123, 1, 0, 0, 0, 12,
		145, 1, 0, 0, 0, 14, 147, 1, 0, 0, 0, 16, 151, 1, 0, 0, 0, 18, 156, 1,
		0, 0, 0, 20, 161, 1, 0, 0, 0, 22, 178, 1, 0, 0, 0, 24, 205, 1, 0, 0, 0,
		26, 207, 1, 0, 0, 0, 28, 209, 1, 0, 0, 0, 30, 211, 1, 0, 0, 0, 32, 216,
		1, 0, 0, 0, 34, 218, 1, 0, 0, 0, 36, 38, 3, 2, 1, 0, 37, 36, 1, 0, 0, 0,
		38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1,
		0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 0, 0, 1, 43, 1, 1, 0, 0, 0, 44,
		51, 3, 34, 17, 0, 45, 48, 3, 12, 6, 0, 46, 48, 3, 4, 2, 0, 47, 45, 1, 0,
		0, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51,
		5, 39, 0, 0, 50, 44, 1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 51, 3, 1, 0, 0, 0,
		52, 53, 5, 26, 0, 0, 53, 109, 3, 20, 10, 0, 54, 55, 5, 27, 0, 0, 55, 109,
		3, 24, 12, 0, 56, 57, 5, 24, 0, 0, 57, 58, 5, 1, 0, 0, 58, 59, 5, 39, 0,
		0, 59, 60, 3, 10, 5, 0, 60, 61, 5, 39, 0, 0, 61, 62, 5, 2, 0, 0, 62, 109,
		1, 0, 0, 0, 63, 64, 5, 24, 0, 0, 64, 109, 3, 8, 4, 0, 65, 66, 5, 25, 0,
		0, 66, 67, 3, 32, 16, 0, 67, 69, 5, 1, 0, 0, 68, 70, 3, 30, 15, 0, 69,
		68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 5, 2, 0,
		0, 72, 73, 3, 6, 3, 0, 73, 109, 1, 0, 0, 0, 74, 75, 5, 3, 0, 0, 75, 76,
		3, 32, 16, 0, 76, 78, 5, 1, 0, 0, 77, 79, 3, 16, 8, 0, 78, 77, 1, 0, 0,
		0, 78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 5, 2, 0, 0, 81, 109,
		1, 0, 0, 0, 82, 83, 5, 28, 0, 0, 83, 109, 3, 16, 8, 0, 84, 85, 5, 28, 0,
		0, 85, 86, 5, 1, 0, 0, 86, 87, 5, 39, 0, 0, 87, 88, 3, 18, 9, 0, 88, 89,
		5, 39, 0, 0, 89, 90, 5, 2, 0, 0, 90, 109, 1, 0, 0, 0, 91, 92, 5, 29, 0,
		0, 92, 109, 3, 16, 8, 0, 93, 94, 5, 29, 0, 0, 94, 95, 5, 1, 0, 0, 95, 96,
		5, 39, 0, 0, 96, 97, 3, 18, 9, 0, 97, 98, 5, 39, 0, 0, 98, 99, 5, 2, 0,
		0, 99, 109, 1, 0, 0, 0, 100, 102, 5, 30, 0, 0, 101, 103, 3, 26, 13, 0,
		102, 101, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 109, 1, 0, 0, 0, 104,
		106, 5, 31, 0, 0, 105, 107, 3, 26, 13, 0, 106, 105, 1, 0, 0, 0, 106, 107,
		1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 52, 1, 0, 0, 0, 108, 54, 1, 0, 0,
		0, 108, 56, 1, 0, 0, 0, 108, 63, 1, 0, 0, 0, 108, 65, 1, 0, 0, 0, 108,
		74, 1, 0, 0, 0, 108, 82, 1, 0, 0, 0, 108, 84, 1, 0, 0, 0, 108, 91, 1, 0,
		0, 0, 108, 93, 1, 0, 0, 0, 108, 100, 1, 0, 0, 0, 108, 104, 1, 0, 0, 0,
		109, 5, 1, 0, 0, 0, 110, 114, 5, 4, 0, 0, 111, 113, 3, 2, 1, 0, 112, 111,
		1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 5, 0, 0,
		118, 7, 1, 0, 0, 0, 119, 120, 3, 32, 16, 0, 120, 121, 5, 6, 0, 0, 121,
		122, 3, 20, 10, 0, 122, 9, 1, 0, 0, 0, 123, 126, 3, 8, 4, 0, 124, 125,
		5, 39, 0, 0, 125, 127, 3, 10, 5, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1,
		0, 0, 0, 127, 11, 1, 0, 0, 0, 128, 146, 3, 32, 16, 0, 129, 130, 3, 32,
		16, 0, 130, 131, 3, 14, 7, 0, 131, 146, 1, 0, 0, 0, 132, 133, 3, 32, 16,
		0, 133, 134, 3, 16, 8, 0, 134, 146, 1, 0, 0, 0, 135, 136, 3, 32, 16, 0,
		136, 137, 3, 14, 7, 0, 137, 138, 5, 7, 0, 0, 138, 139, 3, 16, 8, 0, 139,
		146, 1, 0, 0, 0, 140, 141, 3, 32, 16, 0, 141, 142, 3, 16, 8, 0, 142, 143,
		5, 7, 0, 0, 143, 144, 3, 14, 7, 0, 144, 146, 1, 0, 0, 0, 145, 128, 1, 0,
		0, 0, 145, 129, 1, 0, 0, 0, 145, 132, 1, 0, 0, 0, 145, 135, 1, 0, 0, 0,
		145, 140, 1, 0, 0, 0, 146, 13, 1, 0, 0, 0, 147, 148, 5, 8, 0, 0, 148, 149,
		3, 16, 8, 0, 149, 150, 5, 9, 0, 0, 150, 15, 1, 0, 0, 0, 151, 154, 3, 20,
		10, 0, 152, 153, 5, 7, 0, 0, 153, 155, 3, 16, 8, 0, 154, 152, 1, 0, 0,
		0, 154, 155, 1, 0, 0, 0, 155, 17, 1, 0, 0, 0, 156, 159, 3, 16, 8, 0, 157,
		158, 5, 39, 0, 0, 158, 160, 3, 18, 9, 0, 159, 157, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 19, 1, 0, 0, 0, 161, 162, 3, 22, 11, 0, 162, 21, 1, 0,
		0, 0, 163, 164, 6, 11, -1, 0, 164, 165, 5, 10, 0, 0, 165, 179, 3, 22, 11,
		11, 166, 167, 5, 17, 0, 0, 167, 179, 3, 22, 11, 7, 168, 169, 5, 1, 0, 0,
		169, 170, 3, 22, 11, 0, 170, 171, 5, 2, 0, 0, 171, 179, 1, 0, 0, 0, 172,
		177, 3, 26, 13, 0, 173, 177, 3, 32, 16, 0, 174, 177, 3, 28, 14, 0, 175,
		177, 3, 24, 12, 0, 176, 172, 1, 0, 0, 0, 176, 173, 1, 0, 0, 0, 176, 174,
		1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0, 178, 163, 1, 0,
		0, 0, 178, 166, 1, 0, 0, 0, 178, 168, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0,
		179, 202, 1, 0, 0, 0, 180, 181, 10, 9, 0, 0, 181, 182, 7, 0, 0, 0, 182,
		201, 3, 22, 11, 10, 183, 184, 10, 8, 0, 0, 184, 185, 7, 1, 0, 0, 185, 201,
		3, 22, 11, 9, 186, 187, 10, 6, 0, 0, 187, 188, 7, 2, 0, 0, 188, 201, 3,
		22, 11, 7, 189, 190, 10, 5, 0, 0, 190, 191, 5, 20, 0, 0, 191, 201, 3, 22,
		11, 6, 192, 193, 10, 4, 0, 0, 193, 194, 5, 21, 0, 0, 194, 201, 3, 22, 11,
		5, 195, 196, 10, 3, 0, 0, 196, 197, 5, 22, 0, 0, 197, 201, 3, 22, 11, 4,
		198, 199, 10, 10, 0, 0, 199, 201, 7, 3, 0, 0, 200, 180, 1, 0, 0, 0, 200,
		183, 1, 0, 0, 0, 200, 186, 1, 0, 0, 0, 200, 189, 1, 0, 0, 0, 200, 192,
		1, 0, 0, 0, 200, 195, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201, 204, 1, 0,
		0, 0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 23, 1, 0, 0, 0,
		204, 202, 1, 0, 0, 0, 205, 206, 5, 33, 0, 0, 206, 25, 1, 0, 0, 0, 207,
		208, 7, 4, 0, 0, 208, 27, 1, 0, 0, 0, 209, 210, 5, 32, 0, 0, 210, 29, 1,
		0, 0, 0, 211, 214, 3, 32, 16, 0, 212, 213, 5, 7, 0, 0, 213, 215, 3, 30,
		15, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 31, 1, 0, 0, 0,
		216, 217, 5, 37, 0, 0, 217, 33, 1, 0, 0, 0, 218, 219, 5, 37, 0, 0, 219,
		220, 5, 23, 0, 0, 220, 35, 1, 0, 0, 0, 18, 39, 47, 50, 69, 78, 102, 106,
		108, 114, 126, 145, 154, 159, 176, 178, 200, 202, 214,
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
	DD8ASMParserP_ORG   = 26
	DD8ASMParserP_INC   = 27
	DD8ASMParserP_DB    = 28
	DD8ASMParserP_DW    = 29
	DD8ASMParserP_BYTE  = 30
	DD8ASMParserP_WORD  = 31
	DD8ASMParserREG     = 32
	DD8ASMParserSTR     = 33
	DD8ASMParserHEX_NUM = 34
	DD8ASMParserBIN_NUM = 35
	DD8ASMParserDEC_NUM = 36
	DD8ASMParserNAME    = 37
	DD8ASMParserCOMMENT = 38
	DD8ASMParserEOL     = 39
	DD8ASMParserWS      = 40
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&691472957448) != 0 {
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

		case DD8ASMParserT__2, DD8ASMParserP_DEF, DD8ASMParserP_TMPL, DD8ASMParserP_ORG, DD8ASMParserP_INC, DD8ASMParserP_DB, DD8ASMParserP_DW, DD8ASMParserP_BYTE, DD8ASMParserP_WORD:
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
	Argument() IArgumentContext
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

func (s *Prep_instructionContext) Argument() IArgumentContext {
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

	p.SetState(108)
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
			p.Match(DD8ASMParserT__2)
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
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&270583071746) != 0 {
			{
				p.SetState(77)
				p.Arglist()
			}

		}
		{
			p.SetState(80)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(82)
			p.Match(DD8ASMParserP_DB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Arglist()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(84)
			p.Match(DD8ASMParserP_DB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Arglist_lines()
		}
		{
			p.SetState(88)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(91)
			p.Match(DD8ASMParserP_DW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Arglist()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(93)
			p.Match(DD8ASMParserP_DW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Arglist_lines()
		}
		{
			p.SetState(97)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(100)
			p.Match(DD8ASMParserP_BYTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120259084288) != 0 {
			{
				p.SetState(101)
				p.Num()
			}

		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(104)
			p.Match(DD8ASMParserP_WORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120259084288) != 0 {
			{
				p.SetState(105)
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
		p.SetState(110)
		p.Match(DD8ASMParserT__3)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&691472957448) != 0 {
		{
			p.SetState(111)
			p.Statement()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(DD8ASMParserT__4)
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
		p.SetState(119)
		p.Name()
	}
	{
		p.SetState(120)
		p.Match(DD8ASMParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
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
		p.SetState(123)
		p.Prep_def_args()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(124)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
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
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Name()
		}
		{
			p.SetState(130)
			p.Arglist_p()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Name()
		}
		{
			p.SetState(133)
			p.Arglist()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(135)
			p.Name()
		}
		{
			p.SetState(136)
			p.Arglist_p()
		}
		{
			p.SetState(137)
			p.Match(DD8ASMParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Arglist()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(140)
			p.Name()
		}
		{
			p.SetState(141)
			p.Arglist()
		}
		{
			p.SetState(142)
			p.Match(DD8ASMParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
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
		p.SetState(147)
		p.Match(DD8ASMParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Arglist()
	}
	{
		p.SetState(149)
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
		p.SetState(151)
		p.Argument()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(152)
			p.Match(DD8ASMParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
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
		p.SetState(156)
		p.Arglist()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(157)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
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
		p.SetState(161)
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
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DD8ASMParserT__9:
		{
			p.SetState(164)
			p.Match(DD8ASMParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.expr(11)
		}

	case DD8ASMParserT__16:
		{
			p.SetState(166)
			p.Match(DD8ASMParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.expr(7)
		}

	case DD8ASMParserT__0:
		{
			p.SetState(168)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.expr(0)
		}
		{
			p.SetState(170)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DD8ASMParserREG, DD8ASMParserSTR, DD8ASMParserHEX_NUM, DD8ASMParserBIN_NUM, DD8ASMParserDEC_NUM, DD8ASMParserNAME:
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case DD8ASMParserHEX_NUM, DD8ASMParserBIN_NUM, DD8ASMParserDEC_NUM:
			{
				p.SetState(172)
				p.Num()
			}

		case DD8ASMParserNAME:
			{
				p.SetState(173)
				p.Name()
			}

		case DD8ASMParserREG:
			{
				p.SetState(174)
				p.Reg()
			}

		case DD8ASMParserSTR:
			{
				p.SetState(175)
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
	p.SetState(202)
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
			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(181)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&57344) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(182)
					p.expr(10)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(184)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DD8ASMParserT__15 || _la == DD8ASMParserT__16) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(185)
					p.expr(9)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(187)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DD8ASMParserT__17 || _la == DD8ASMParserT__18) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(188)
					p.expr(7)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(190)
					p.Match(DD8ASMParserT__19)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(191)
					p.expr(6)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(193)
					p.Match(DD8ASMParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(194)
					p.expr(5)
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(196)
					p.Match(DD8ASMParserT__21)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(197)
					p.expr(4)
				}

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, DD8ASMParserRULE_expr)
				p.SetState(198)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(199)
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
		p.SetState(204)
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
		p.SetState(205)
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
		p.SetState(207)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120259084288) != 0) {
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
		p.SetState(209)
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
		p.SetState(211)
		p.Name()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DD8ASMParserT__6 {
		{
			p.SetState(212)
			p.Match(DD8ASMParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
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
		p.SetState(216)
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
		p.SetState(218)
		p.Match(DD8ASMParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
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
