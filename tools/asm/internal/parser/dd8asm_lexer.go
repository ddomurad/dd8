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
		"'.inc'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "P_DEF", "P_ORG", "P_INC", "REG", "STR",
		"HEX_NUM", "BIN_NUM", "DEC_NUM", "NAME", "COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "P_DEF", "P_ORG",
		"P_INC", "REG", "STR", "HEX_NUM", "BIN_NUM", "DEC_NUM", "NAME", "COMMENT",
		"EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 134, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 76, 8, 11,
		10, 11, 12, 11, 79, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 4,
		12, 87, 8, 12, 11, 12, 12, 12, 88, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 95,
		8, 13, 11, 13, 12, 13, 96, 1, 14, 3, 14, 100, 8, 14, 1, 14, 4, 14, 103,
		8, 14, 11, 14, 12, 14, 104, 1, 15, 4, 15, 108, 8, 15, 11, 15, 12, 15, 109,
		1, 15, 4, 15, 113, 8, 15, 11, 15, 12, 15, 114, 1, 16, 1, 16, 5, 16, 119,
		8, 16, 10, 16, 12, 16, 122, 9, 16, 1, 16, 1, 16, 1, 17, 4, 17, 127, 8,
		17, 11, 17, 12, 17, 128, 1, 18, 1, 18, 1, 18, 1, 18, 0, 0, 19, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 1, 0, 20, 2, 0,
		68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0,
		79, 79, 111, 111, 2, 0, 82, 82, 114, 114, 2, 0, 71, 71, 103, 103, 2, 0,
		73, 73, 105, 105, 2, 0, 78, 78, 110, 110, 2, 0, 67, 67, 99, 99, 2, 0, 65,
		90, 97, 122, 3, 0, 10, 10, 13, 13, 34, 34, 2, 0, 88, 88, 120, 120, 4, 0,
		48, 57, 65, 70, 95, 95, 97, 102, 2, 0, 66, 66, 98, 98, 2, 0, 48, 49, 95,
		95, 2, 0, 48, 57, 95, 95, 4, 0, 45, 45, 65, 90, 95, 95, 97, 122, 5, 0,
		45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10, 13, 13, 2, 0, 9,
		9, 32, 32, 143, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0,
		0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0,
		0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1,
		0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 46,
		1, 0, 0, 0, 9, 48, 1, 0, 0, 0, 11, 50, 1, 0, 0, 0, 13, 52, 1, 0, 0, 0,
		15, 54, 1, 0, 0, 0, 17, 59, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 69, 1,
		0, 0, 0, 23, 71, 1, 0, 0, 0, 25, 82, 1, 0, 0, 0, 27, 90, 1, 0, 0, 0, 29,
		99, 1, 0, 0, 0, 31, 107, 1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 126, 1, 0,
		0, 0, 37, 130, 1, 0, 0, 0, 39, 40, 5, 40, 0, 0, 40, 2, 1, 0, 0, 0, 41,
		42, 5, 41, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 58, 0, 0, 44, 45, 5, 61,
		0, 0, 45, 6, 1, 0, 0, 0, 46, 47, 5, 44, 0, 0, 47, 8, 1, 0, 0, 0, 48, 49,
		5, 91, 0, 0, 49, 10, 1, 0, 0, 0, 50, 51, 5, 93, 0, 0, 51, 12, 1, 0, 0,
		0, 52, 53, 5, 58, 0, 0, 53, 14, 1, 0, 0, 0, 54, 55, 5, 46, 0, 0, 55, 56,
		7, 0, 0, 0, 56, 57, 7, 1, 0, 0, 57, 58, 7, 2, 0, 0, 58, 16, 1, 0, 0, 0,
		59, 60, 5, 46, 0, 0, 60, 61, 7, 3, 0, 0, 61, 62, 7, 4, 0, 0, 62, 63, 7,
		5, 0, 0, 63, 18, 1, 0, 0, 0, 64, 65, 5, 46, 0, 0, 65, 66, 7, 6, 0, 0, 66,
		67, 7, 7, 0, 0, 67, 68, 7, 8, 0, 0, 68, 20, 1, 0, 0, 0, 69, 70, 7, 9, 0,
		0, 70, 22, 1, 0, 0, 0, 71, 77, 5, 34, 0, 0, 72, 73, 5, 34, 0, 0, 73, 76,
		5, 34, 0, 0, 74, 76, 8, 10, 0, 0, 75, 72, 1, 0, 0, 0, 75, 74, 1, 0, 0,
		0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80,
		1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 34, 0, 0, 81, 24, 1, 0, 0, 0,
		82, 83, 5, 48, 0, 0, 83, 84, 7, 11, 0, 0, 84, 86, 1, 0, 0, 0, 85, 87, 7,
		12, 0, 0, 86, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 26, 1, 0, 0, 0, 90, 91, 5, 48, 0, 0, 91, 92, 7, 13,
		0, 0, 92, 94, 1, 0, 0, 0, 93, 95, 7, 14, 0, 0, 94, 93, 1, 0, 0, 0, 95,
		96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 28, 1, 0, 0,
		0, 98, 100, 5, 45, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100,
		102, 1, 0, 0, 0, 101, 103, 7, 15, 0, 0, 102, 101, 1, 0, 0, 0, 103, 104,
		1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 30, 1, 0,
		0, 0, 106, 108, 7, 16, 0, 0, 107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0,
		109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111,
		113, 7, 17, 0, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 32, 1, 0, 0, 0, 116, 120, 5, 59,
		0, 0, 117, 119, 8, 18, 0, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0,
		120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 123, 124, 6, 16, 0, 0, 124, 34, 1, 0, 0, 0, 125, 127,
		7, 18, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 126, 1, 0,
		0, 0, 128, 129, 1, 0, 0, 0, 129, 36, 1, 0, 0, 0, 130, 131, 7, 19, 0, 0,
		131, 132, 1, 0, 0, 0, 132, 133, 6, 18, 0, 0, 133, 38, 1, 0, 0, 0, 11, 0,
		75, 77, 88, 96, 99, 104, 109, 114, 120, 128, 1, 6, 0, 0,
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
	DD8ASMLexerREG     = 11
	DD8ASMLexerSTR     = 12
	DD8ASMLexerHEX_NUM = 13
	DD8ASMLexerBIN_NUM = 14
	DD8ASMLexerDEC_NUM = 15
	DD8ASMLexerNAME    = 16
	DD8ASMLexerCOMMENT = 17
	DD8ASMLexerEOL     = 18
	DD8ASMLexerWS      = 19
)
