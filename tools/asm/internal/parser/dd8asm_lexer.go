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
		"", "'('", "')'", "':='", "','", "':'", "'.def'", "'.org'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "P_DEF", "P_ORG", "NAME", "HEX_NUM", "BIN_NUM",
		"DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "P_DEF", "P_ORG", "NAME", "HEX_NUM",
		"BIN_NUM", "DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 103, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 4, 7, 52, 8, 7, 11, 7, 12, 7,
		53, 1, 7, 5, 7, 57, 8, 7, 10, 7, 12, 7, 60, 9, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 4, 8, 66, 8, 8, 11, 8, 12, 8, 67, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 74,
		8, 9, 11, 9, 12, 9, 75, 1, 10, 3, 10, 79, 8, 10, 1, 10, 4, 10, 82, 8, 10,
		11, 10, 12, 10, 83, 1, 11, 1, 11, 5, 11, 88, 8, 11, 10, 11, 12, 11, 91,
		9, 11, 1, 11, 1, 11, 1, 12, 4, 12, 96, 8, 12, 11, 12, 12, 12, 97, 1, 13,
		1, 13, 1, 13, 1, 13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 1, 0, 15, 2, 0,
		68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0,
		79, 79, 111, 111, 2, 0, 82, 82, 114, 114, 2, 0, 71, 71, 103, 103, 4, 0,
		45, 45, 65, 90, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95,
		97, 122, 2, 0, 88, 88, 120, 120, 4, 0, 48, 57, 65, 70, 95, 95, 97, 102,
		2, 0, 66, 66, 98, 98, 2, 0, 48, 49, 95, 95, 2, 0, 48, 57, 95, 95, 2, 0,
		10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 110, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 1, 29, 1, 0, 0, 0, 3, 31, 1, 0, 0, 0, 5, 33, 1, 0, 0, 0, 7,
		36, 1, 0, 0, 0, 9, 38, 1, 0, 0, 0, 11, 40, 1, 0, 0, 0, 13, 45, 1, 0, 0,
		0, 15, 51, 1, 0, 0, 0, 17, 61, 1, 0, 0, 0, 19, 69, 1, 0, 0, 0, 21, 78,
		1, 0, 0, 0, 23, 85, 1, 0, 0, 0, 25, 95, 1, 0, 0, 0, 27, 99, 1, 0, 0, 0,
		29, 30, 5, 40, 0, 0, 30, 2, 1, 0, 0, 0, 31, 32, 5, 41, 0, 0, 32, 4, 1,
		0, 0, 0, 33, 34, 5, 58, 0, 0, 34, 35, 5, 61, 0, 0, 35, 6, 1, 0, 0, 0, 36,
		37, 5, 44, 0, 0, 37, 8, 1, 0, 0, 0, 38, 39, 5, 58, 0, 0, 39, 10, 1, 0,
		0, 0, 40, 41, 5, 46, 0, 0, 41, 42, 7, 0, 0, 0, 42, 43, 7, 1, 0, 0, 43,
		44, 7, 2, 0, 0, 44, 12, 1, 0, 0, 0, 45, 46, 5, 46, 0, 0, 46, 47, 7, 3,
		0, 0, 47, 48, 7, 4, 0, 0, 48, 49, 7, 5, 0, 0, 49, 14, 1, 0, 0, 0, 50, 52,
		7, 6, 0, 0, 51, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0,
		53, 54, 1, 0, 0, 0, 54, 58, 1, 0, 0, 0, 55, 57, 7, 7, 0, 0, 56, 55, 1,
		0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59,
		16, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62, 5, 48, 0, 0, 62, 63, 7, 8,
		0, 0, 63, 65, 1, 0, 0, 0, 64, 66, 7, 9, 0, 0, 65, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 18, 1, 0, 0, 0,
		69, 70, 5, 48, 0, 0, 70, 71, 7, 10, 0, 0, 71, 73, 1, 0, 0, 0, 72, 74, 7,
		11, 0, 0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75,
		76, 1, 0, 0, 0, 76, 20, 1, 0, 0, 0, 77, 79, 5, 45, 0, 0, 78, 77, 1, 0,
		0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 1, 0, 0, 0, 80, 82, 7, 12, 0, 0, 81,
		80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 22, 1, 0, 0, 0, 85, 89, 5, 59, 0, 0, 86, 88, 8, 13, 0, 0, 87, 86,
		1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 6, 11, 0, 0, 93, 24, 1,
		0, 0, 0, 94, 96, 7, 13, 0, 0, 95, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97,
		95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 26, 1, 0, 0, 0, 99, 100, 7, 14,
		0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 6, 13, 0, 0, 102, 28, 1, 0, 0, 0,
		9, 0, 53, 58, 67, 75, 78, 83, 89, 97, 1, 6, 0, 0,
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
	DD8ASMLexerP_DEF   = 6
	DD8ASMLexerP_ORG   = 7
	DD8ASMLexerNAME    = 8
	DD8ASMLexerHEX_NUM = 9
	DD8ASMLexerBIN_NUM = 10
	DD8ASMLexerDEC_NUM = 11
	DD8ASMLexerCOMMENT = 12
	DD8ASMLexerEOL     = 13
	DD8ASMLexerWS      = 14
)
