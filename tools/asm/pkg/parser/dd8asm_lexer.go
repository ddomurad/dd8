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
		"'.inc'", "'.db'", "'.dw'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "P_DEF", "P_ORG", "P_INC", "P_DB", "P_DW",
		"REG", "STR", "HEX_NUM", "BIN_NUM", "DEC_NUM", "NAME", "COMMENT", "EOL",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "P_DEF", "P_ORG",
		"P_INC", "P_DB", "P_DW", "REG", "STR", "HEX_NUM", "BIN_NUM", "DEC_NUM",
		"NAME", "COMMENT", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 146, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 88,
		8, 13, 10, 13, 12, 13, 91, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 4, 14, 99, 8, 14, 11, 14, 12, 14, 100, 1, 15, 1, 15, 1, 15, 1, 15,
		4, 15, 107, 8, 15, 11, 15, 12, 15, 108, 1, 16, 3, 16, 112, 8, 16, 1, 16,
		4, 16, 115, 8, 16, 11, 16, 12, 16, 116, 1, 17, 4, 17, 120, 8, 17, 11, 17,
		12, 17, 121, 1, 17, 4, 17, 125, 8, 17, 11, 17, 12, 17, 126, 1, 18, 1, 18,
		5, 18, 131, 8, 18, 10, 18, 12, 18, 134, 9, 18, 1, 18, 1, 18, 1, 19, 4,
		19, 139, 8, 19, 11, 19, 12, 19, 140, 1, 20, 1, 20, 1, 20, 1, 20, 0, 0,
		21, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 1, 0, 21, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2,
		0, 70, 70, 102, 102, 2, 0, 79, 79, 111, 111, 2, 0, 82, 82, 114, 114, 2,
		0, 71, 71, 103, 103, 2, 0, 73, 73, 105, 105, 2, 0, 78, 78, 110, 110, 2,
		0, 67, 67, 99, 99, 2, 0, 66, 66, 98, 98, 2, 0, 87, 87, 119, 119, 2, 0,
		65, 90, 97, 122, 3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 88, 88, 120, 120,
		4, 0, 48, 57, 65, 70, 95, 95, 97, 102, 2, 0, 48, 49, 95, 95, 2, 0, 48,
		57, 95, 95, 4, 0, 45, 45, 65, 90, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57,
		65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 155,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 1, 43, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5,
		47, 1, 0, 0, 0, 7, 50, 1, 0, 0, 0, 9, 52, 1, 0, 0, 0, 11, 54, 1, 0, 0,
		0, 13, 56, 1, 0, 0, 0, 15, 58, 1, 0, 0, 0, 17, 63, 1, 0, 0, 0, 19, 68,
		1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 77, 1, 0, 0, 0, 25, 81, 1, 0, 0, 0,
		27, 83, 1, 0, 0, 0, 29, 94, 1, 0, 0, 0, 31, 102, 1, 0, 0, 0, 33, 111, 1,
		0, 0, 0, 35, 119, 1, 0, 0, 0, 37, 128, 1, 0, 0, 0, 39, 138, 1, 0, 0, 0,
		41, 142, 1, 0, 0, 0, 43, 44, 5, 40, 0, 0, 44, 2, 1, 0, 0, 0, 45, 46, 5,
		41, 0, 0, 46, 4, 1, 0, 0, 0, 47, 48, 5, 58, 0, 0, 48, 49, 5, 61, 0, 0,
		49, 6, 1, 0, 0, 0, 50, 51, 5, 44, 0, 0, 51, 8, 1, 0, 0, 0, 52, 53, 5, 91,
		0, 0, 53, 10, 1, 0, 0, 0, 54, 55, 5, 93, 0, 0, 55, 12, 1, 0, 0, 0, 56,
		57, 5, 58, 0, 0, 57, 14, 1, 0, 0, 0, 58, 59, 5, 46, 0, 0, 59, 60, 7, 0,
		0, 0, 60, 61, 7, 1, 0, 0, 61, 62, 7, 2, 0, 0, 62, 16, 1, 0, 0, 0, 63, 64,
		5, 46, 0, 0, 64, 65, 7, 3, 0, 0, 65, 66, 7, 4, 0, 0, 66, 67, 7, 5, 0, 0,
		67, 18, 1, 0, 0, 0, 68, 69, 5, 46, 0, 0, 69, 70, 7, 6, 0, 0, 70, 71, 7,
		7, 0, 0, 71, 72, 7, 8, 0, 0, 72, 20, 1, 0, 0, 0, 73, 74, 5, 46, 0, 0, 74,
		75, 7, 0, 0, 0, 75, 76, 7, 9, 0, 0, 76, 22, 1, 0, 0, 0, 77, 78, 5, 46,
		0, 0, 78, 79, 7, 0, 0, 0, 79, 80, 7, 10, 0, 0, 80, 24, 1, 0, 0, 0, 81,
		82, 7, 11, 0, 0, 82, 26, 1, 0, 0, 0, 83, 89, 5, 34, 0, 0, 84, 85, 5, 34,
		0, 0, 85, 88, 5, 34, 0, 0, 86, 88, 8, 12, 0, 0, 87, 84, 1, 0, 0, 0, 87,
		86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0,
		0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 34, 0, 0, 93, 28,
		1, 0, 0, 0, 94, 95, 5, 48, 0, 0, 95, 96, 7, 13, 0, 0, 96, 98, 1, 0, 0,
		0, 97, 99, 7, 14, 0, 0, 98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98,
		1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 30, 1, 0, 0, 0, 102, 103, 5, 48,
		0, 0, 103, 104, 7, 9, 0, 0, 104, 106, 1, 0, 0, 0, 105, 107, 7, 15, 0, 0,
		106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		109, 1, 0, 0, 0, 109, 32, 1, 0, 0, 0, 110, 112, 5, 45, 0, 0, 111, 110,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 115, 7, 16,
		0, 0, 114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 34, 1, 0, 0, 0, 118, 120, 7, 17, 0, 0, 119,
		118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122,
		1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 125, 7, 18, 0, 0, 124, 123, 1, 0,
		0, 0, 125, 126, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0,
		127, 36, 1, 0, 0, 0, 128, 132, 5, 59, 0, 0, 129, 131, 8, 19, 0, 0, 130,
		129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133,
		1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 6, 18,
		0, 0, 136, 38, 1, 0, 0, 0, 137, 139, 7, 19, 0, 0, 138, 137, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		40, 1, 0, 0, 0, 142, 143, 7, 20, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145,
		6, 20, 0, 0, 145, 42, 1, 0, 0, 0, 11, 0, 87, 89, 100, 108, 111, 116, 121,
		126, 132, 140, 1, 6, 0, 0,
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
	DD8ASMLexerREG     = 13
	DD8ASMLexerSTR     = 14
	DD8ASMLexerHEX_NUM = 15
	DD8ASMLexerBIN_NUM = 16
	DD8ASMLexerDEC_NUM = 17
	DD8ASMLexerNAME    = 18
	DD8ASMLexerCOMMENT = 19
	DD8ASMLexerEOL     = 20
	DD8ASMLexerWS      = 21
)