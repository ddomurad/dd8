// Code generated from DD8ASM.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type DD8ASMLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var DD8ASMLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dd8asmlexerLexerInit() {
	staticData := &DD8ASMLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "':='", "','", "'['", "']'", "':'", "'.def'", "'.org'",
		"'.inc'", "'.db'", "'.dw'", "'.byte'", "'.word'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "P_DEF", "P_ORG", "P_INC", "P_DB", "P_DW",
		"P_BYTE", "P_WORD", "REG", "STR", "HEX_NUM", "BIN_NUM", "DEC_NUM", "NAME",
		"COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "P_DEF", "P_ORG",
		"P_INC", "P_DB", "P_DW", "P_BYTE", "P_WORD", "REG", "STR", "HEX_NUM",
		"BIN_NUM", "DEC_NUM", "NAME", "COMMENT", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 23, 162, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 104, 8, 15, 10, 15, 12, 15, 107,
		9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 115, 8, 16, 11,
		16, 12, 16, 116, 1, 17, 1, 17, 1, 17, 1, 17, 4, 17, 123, 8, 17, 11, 17,
		12, 17, 124, 1, 18, 3, 18, 128, 8, 18, 1, 18, 4, 18, 131, 8, 18, 11, 18,
		12, 18, 132, 1, 19, 4, 19, 136, 8, 19, 11, 19, 12, 19, 137, 1, 19, 4, 19,
		141, 8, 19, 11, 19, 12, 19, 142, 1, 20, 1, 20, 5, 20, 147, 8, 20, 10, 20,
		12, 20, 150, 9, 20, 1, 20, 1, 20, 1, 21, 4, 21, 155, 8, 21, 11, 21, 12,
		21, 156, 1, 22, 1, 22, 1, 22, 1, 22, 0, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		1, 0, 23, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70,
		102, 102, 2, 0, 79, 79, 111, 111, 2, 0, 82, 82, 114, 114, 2, 0, 71, 71,
		103, 103, 2, 0, 73, 73, 105, 105, 2, 0, 78, 78, 110, 110, 2, 0, 67, 67,
		99, 99, 2, 0, 66, 66, 98, 98, 2, 0, 87, 87, 119, 119, 2, 0, 89, 89, 121,
		121, 2, 0, 84, 84, 116, 116, 2, 0, 65, 90, 97, 122, 3, 0, 10, 10, 13, 13,
		34, 34, 2, 0, 88, 88, 120, 120, 4, 0, 48, 57, 65, 70, 95, 95, 97, 102,
		2, 0, 48, 49, 95, 95, 2, 0, 48, 57, 95, 95, 4, 0, 45, 45, 65, 90, 95, 95,
		97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13,
		13, 2, 0, 9, 9, 32, 32, 171, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 49, 1, 0, 0, 0, 5, 51, 1,
		0, 0, 0, 7, 54, 1, 0, 0, 0, 9, 56, 1, 0, 0, 0, 11, 58, 1, 0, 0, 0, 13,
		60, 1, 0, 0, 0, 15, 62, 1, 0, 0, 0, 17, 67, 1, 0, 0, 0, 19, 72, 1, 0, 0,
		0, 21, 77, 1, 0, 0, 0, 23, 81, 1, 0, 0, 0, 25, 85, 1, 0, 0, 0, 27, 91,
		1, 0, 0, 0, 29, 97, 1, 0, 0, 0, 31, 99, 1, 0, 0, 0, 33, 110, 1, 0, 0, 0,
		35, 118, 1, 0, 0, 0, 37, 127, 1, 0, 0, 0, 39, 135, 1, 0, 0, 0, 41, 144,
		1, 0, 0, 0, 43, 154, 1, 0, 0, 0, 45, 158, 1, 0, 0, 0, 47, 48, 5, 40, 0,
		0, 48, 2, 1, 0, 0, 0, 49, 50, 5, 41, 0, 0, 50, 4, 1, 0, 0, 0, 51, 52, 5,
		58, 0, 0, 52, 53, 5, 61, 0, 0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 44, 0, 0,
		55, 8, 1, 0, 0, 0, 56, 57, 5, 91, 0, 0, 57, 10, 1, 0, 0, 0, 58, 59, 5,
		93, 0, 0, 59, 12, 1, 0, 0, 0, 60, 61, 5, 58, 0, 0, 61, 14, 1, 0, 0, 0,
		62, 63, 5, 46, 0, 0, 63, 64, 7, 0, 0, 0, 64, 65, 7, 1, 0, 0, 65, 66, 7,
		2, 0, 0, 66, 16, 1, 0, 0, 0, 67, 68, 5, 46, 0, 0, 68, 69, 7, 3, 0, 0, 69,
		70, 7, 4, 0, 0, 70, 71, 7, 5, 0, 0, 71, 18, 1, 0, 0, 0, 72, 73, 5, 46,
		0, 0, 73, 74, 7, 6, 0, 0, 74, 75, 7, 7, 0, 0, 75, 76, 7, 8, 0, 0, 76, 20,
		1, 0, 0, 0, 77, 78, 5, 46, 0, 0, 78, 79, 7, 0, 0, 0, 79, 80, 7, 9, 0, 0,
		80, 22, 1, 0, 0, 0, 81, 82, 5, 46, 0, 0, 82, 83, 7, 0, 0, 0, 83, 84, 7,
		10, 0, 0, 84, 24, 1, 0, 0, 0, 85, 86, 5, 46, 0, 0, 86, 87, 7, 9, 0, 0,
		87, 88, 7, 11, 0, 0, 88, 89, 7, 12, 0, 0, 89, 90, 7, 1, 0, 0, 90, 26, 1,
		0, 0, 0, 91, 92, 5, 46, 0, 0, 92, 93, 7, 10, 0, 0, 93, 94, 7, 3, 0, 0,
		94, 95, 7, 4, 0, 0, 95, 96, 7, 0, 0, 0, 96, 28, 1, 0, 0, 0, 97, 98, 7,
		13, 0, 0, 98, 30, 1, 0, 0, 0, 99, 105, 5, 34, 0, 0, 100, 101, 5, 92, 0,
		0, 101, 104, 5, 34, 0, 0, 102, 104, 8, 14, 0, 0, 103, 100, 1, 0, 0, 0,
		103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105,
		106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109,
		5, 34, 0, 0, 109, 32, 1, 0, 0, 0, 110, 111, 5, 48, 0, 0, 111, 112, 7, 15,
		0, 0, 112, 114, 1, 0, 0, 0, 113, 115, 7, 16, 0, 0, 114, 113, 1, 0, 0, 0,
		115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		34, 1, 0, 0, 0, 118, 119, 5, 48, 0, 0, 119, 120, 7, 9, 0, 0, 120, 122,
		1, 0, 0, 0, 121, 123, 7, 17, 0, 0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 36, 1, 0, 0, 0,
		126, 128, 5, 45, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		130, 1, 0, 0, 0, 129, 131, 7, 18, 0, 0, 130, 129, 1, 0, 0, 0, 131, 132,
		1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 38, 1, 0,
		0, 0, 134, 136, 7, 19, 0, 0, 135, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0,
		137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0, 139,
		141, 7, 20, 0, 0, 140, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 140,
		1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 40, 1, 0, 0, 0, 144, 148, 5, 59,
		0, 0, 145, 147, 8, 21, 0, 0, 146, 145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0,
		148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150,
		148, 1, 0, 0, 0, 151, 152, 6, 20, 0, 0, 152, 42, 1, 0, 0, 0, 153, 155,
		7, 21, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 154, 1, 0,
		0, 0, 156, 157, 1, 0, 0, 0, 157, 44, 1, 0, 0, 0, 158, 159, 7, 22, 0, 0,
		159, 160, 1, 0, 0, 0, 160, 161, 6, 22, 0, 0, 161, 46, 1, 0, 0, 0, 11, 0,
		103, 105, 116, 124, 127, 132, 137, 142, 148, 156, 1, 6, 0, 0,
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

// DD8ASMLexerInit initializes any static state used to implement DD8ASMLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewDD8ASMLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func DD8ASMLexerInit() {
	staticData := &DD8ASMLexerLexerStaticData
	staticData.once.Do(dd8asmlexerLexerInit)
}

// NewDD8ASMLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewDD8ASMLexer(input antlr.CharStream) *DD8ASMLexer {
	DD8ASMLexerInit()
	l := new(DD8ASMLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &DD8ASMLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "DD8ASM.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DD8ASMLexer tokens.
const (
	DD8ASMLexerT__0    = 1
	DD8ASMLexerT__1    = 2
	DD8ASMLexerT__2    = 3
	DD8ASMLexerT__3    = 4
	DD8ASMLexerT__4    = 5
	DD8ASMLexerT__5    = 6
	DD8ASMLexerT__6    = 7
	DD8ASMLexerP_DEF   = 8
	DD8ASMLexerP_ORG   = 9
	DD8ASMLexerP_INC   = 10
	DD8ASMLexerP_DB    = 11
	DD8ASMLexerP_DW    = 12
	DD8ASMLexerP_BYTE  = 13
	DD8ASMLexerP_WORD  = 14
	DD8ASMLexerREG     = 15
	DD8ASMLexerSTR     = 16
	DD8ASMLexerHEX_NUM = 17
	DD8ASMLexerBIN_NUM = 18
	DD8ASMLexerDEC_NUM = 19
	DD8ASMLexerNAME    = 20
	DD8ASMLexerCOMMENT = 21
	DD8ASMLexerEOL     = 22
	DD8ASMLexerWS      = 23
)
