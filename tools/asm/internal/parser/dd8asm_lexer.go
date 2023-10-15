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
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "P_DEF", "P_ORG", "REG", "NAME", "HEX_NUM",
		"BIN_NUM", "DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "P_DEF", "P_ORG",
		"REG", "NAME", "HEX_NUM", "BIN_NUM", "DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 114, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 64, 8, 10, 11, 10,
		12, 10, 65, 1, 10, 4, 10, 69, 8, 10, 11, 10, 12, 10, 70, 1, 11, 1, 11,
		1, 11, 1, 11, 4, 11, 77, 8, 11, 11, 11, 12, 11, 78, 1, 12, 1, 12, 1, 12,
		1, 12, 4, 12, 85, 8, 12, 11, 12, 12, 12, 86, 1, 13, 3, 13, 90, 8, 13, 1,
		13, 4, 13, 93, 8, 13, 11, 13, 12, 13, 94, 1, 14, 1, 14, 5, 14, 99, 8, 14,
		10, 14, 12, 14, 102, 9, 14, 1, 14, 1, 14, 1, 15, 4, 15, 107, 8, 15, 11,
		15, 12, 15, 108, 1, 16, 1, 16, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 16, 2, 0, 68, 68, 100, 100, 2, 0,
		69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 79, 79, 111, 111, 2, 0,
		82, 82, 114, 114, 2, 0, 71, 71, 103, 103, 2, 0, 65, 90, 97, 122, 4, 0,
		45, 45, 65, 90, 95, 95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95,
		97, 122, 2, 0, 88, 88, 120, 120, 4, 0, 48, 57, 65, 70, 95, 95, 97, 102,
		2, 0, 66, 66, 98, 98, 2, 0, 48, 49, 95, 95, 2, 0, 48, 57, 95, 95, 2, 0,
		10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 121, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1,
		35, 1, 0, 0, 0, 3, 37, 1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 42, 1, 0, 0, 0,
		9, 44, 1, 0, 0, 0, 11, 46, 1, 0, 0, 0, 13, 48, 1, 0, 0, 0, 15, 50, 1, 0,
		0, 0, 17, 55, 1, 0, 0, 0, 19, 60, 1, 0, 0, 0, 21, 63, 1, 0, 0, 0, 23, 72,
		1, 0, 0, 0, 25, 80, 1, 0, 0, 0, 27, 89, 1, 0, 0, 0, 29, 96, 1, 0, 0, 0,
		31, 106, 1, 0, 0, 0, 33, 110, 1, 0, 0, 0, 35, 36, 5, 40, 0, 0, 36, 2, 1,
		0, 0, 0, 37, 38, 5, 41, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 58, 0, 0, 40,
		41, 5, 61, 0, 0, 41, 6, 1, 0, 0, 0, 42, 43, 5, 44, 0, 0, 43, 8, 1, 0, 0,
		0, 44, 45, 5, 91, 0, 0, 45, 10, 1, 0, 0, 0, 46, 47, 5, 93, 0, 0, 47, 12,
		1, 0, 0, 0, 48, 49, 5, 58, 0, 0, 49, 14, 1, 0, 0, 0, 50, 51, 5, 46, 0,
		0, 51, 52, 7, 0, 0, 0, 52, 53, 7, 1, 0, 0, 53, 54, 7, 2, 0, 0, 54, 16,
		1, 0, 0, 0, 55, 56, 5, 46, 0, 0, 56, 57, 7, 3, 0, 0, 57, 58, 7, 4, 0, 0,
		58, 59, 7, 5, 0, 0, 59, 18, 1, 0, 0, 0, 60, 61, 7, 6, 0, 0, 61, 20, 1,
		0, 0, 0, 62, 64, 7, 7, 0, 0, 63, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65,
		63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 69, 7, 8, 0,
		0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71,
		1, 0, 0, 0, 71, 22, 1, 0, 0, 0, 72, 73, 5, 48, 0, 0, 73, 74, 7, 9, 0, 0,
		74, 76, 1, 0, 0, 0, 75, 77, 7, 10, 0, 0, 76, 75, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 24, 1, 0, 0, 0, 80,
		81, 5, 48, 0, 0, 81, 82, 7, 11, 0, 0, 82, 84, 1, 0, 0, 0, 83, 85, 7, 12,
		0, 0, 84, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 87,
		1, 0, 0, 0, 87, 26, 1, 0, 0, 0, 88, 90, 5, 45, 0, 0, 89, 88, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 93, 7, 13, 0, 0, 92, 91, 1,
		0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95,
		28, 1, 0, 0, 0, 96, 100, 5, 59, 0, 0, 97, 99, 8, 14, 0, 0, 98, 97, 1, 0,
		0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101,
		103, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 104, 6, 14, 0, 0, 104, 30,
		1, 0, 0, 0, 105, 107, 7, 14, 0, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 32, 1, 0, 0, 0,
		110, 111, 7, 15, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 6, 16, 0, 0, 113,
		34, 1, 0, 0, 0, 9, 0, 65, 70, 78, 86, 89, 94, 100, 108, 1, 6, 0, 0,
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
	DD8ASMLexerREG     = 10
	DD8ASMLexerNAME    = 11
	DD8ASMLexerHEX_NUM = 12
	DD8ASMLexerBIN_NUM = 13
	DD8ASMLexerDEC_NUM = 14
	DD8ASMLexerCOMMENT = 15
	DD8ASMLexerEOL     = 16
	DD8ASMLexerWS      = 17
)
