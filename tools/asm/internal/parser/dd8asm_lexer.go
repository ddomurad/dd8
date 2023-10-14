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
		"", "", "", "", "", "", "P_DEF", "P_ORG", "REG", "NAME", "HEX_NUM",
		"BIN_NUM", "DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "P_DEF", "P_ORG", "REG", "NAME",
		"HEX_NUM", "BIN_NUM", "DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 106, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 4, 8,
		56, 8, 8, 11, 8, 12, 8, 57, 1, 8, 4, 8, 61, 8, 8, 11, 8, 12, 8, 62, 1,
		9, 1, 9, 1, 9, 1, 9, 4, 9, 69, 8, 9, 11, 9, 12, 9, 70, 1, 10, 1, 10, 1,
		10, 1, 10, 4, 10, 77, 8, 10, 11, 10, 12, 10, 78, 1, 11, 3, 11, 82, 8, 11,
		1, 11, 4, 11, 85, 8, 11, 11, 11, 12, 11, 86, 1, 12, 1, 12, 5, 12, 91, 8,
		12, 10, 12, 12, 12, 94, 9, 12, 1, 12, 1, 12, 1, 13, 4, 13, 99, 8, 13, 11,
		13, 12, 13, 100, 1, 14, 1, 14, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 1, 0, 16, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101,
		2, 0, 70, 70, 102, 102, 2, 0, 79, 79, 111, 111, 2, 0, 82, 82, 114, 114,
		2, 0, 71, 71, 103, 103, 2, 0, 65, 90, 97, 122, 4, 0, 45, 45, 65, 90, 95,
		95, 97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 88, 88,
		120, 120, 4, 0, 48, 57, 65, 70, 95, 95, 97, 102, 2, 0, 66, 66, 98, 98,
		2, 0, 48, 49, 95, 95, 2, 0, 48, 57, 95, 95, 2, 0, 10, 10, 13, 13, 2, 0,
		9, 9, 32, 32, 113, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 33, 1, 0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 38,
		1, 0, 0, 0, 9, 40, 1, 0, 0, 0, 11, 42, 1, 0, 0, 0, 13, 47, 1, 0, 0, 0,
		15, 52, 1, 0, 0, 0, 17, 55, 1, 0, 0, 0, 19, 64, 1, 0, 0, 0, 21, 72, 1,
		0, 0, 0, 23, 81, 1, 0, 0, 0, 25, 88, 1, 0, 0, 0, 27, 98, 1, 0, 0, 0, 29,
		102, 1, 0, 0, 0, 31, 32, 5, 40, 0, 0, 32, 2, 1, 0, 0, 0, 33, 34, 5, 41,
		0, 0, 34, 4, 1, 0, 0, 0, 35, 36, 5, 58, 0, 0, 36, 37, 5, 61, 0, 0, 37,
		6, 1, 0, 0, 0, 38, 39, 5, 44, 0, 0, 39, 8, 1, 0, 0, 0, 40, 41, 5, 58, 0,
		0, 41, 10, 1, 0, 0, 0, 42, 43, 5, 46, 0, 0, 43, 44, 7, 0, 0, 0, 44, 45,
		7, 1, 0, 0, 45, 46, 7, 2, 0, 0, 46, 12, 1, 0, 0, 0, 47, 48, 5, 46, 0, 0,
		48, 49, 7, 3, 0, 0, 49, 50, 7, 4, 0, 0, 50, 51, 7, 5, 0, 0, 51, 14, 1,
		0, 0, 0, 52, 53, 7, 6, 0, 0, 53, 16, 1, 0, 0, 0, 54, 56, 7, 7, 0, 0, 55,
		54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0,
		0, 58, 60, 1, 0, 0, 0, 59, 61, 7, 8, 0, 0, 60, 59, 1, 0, 0, 0, 61, 62,
		1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 18, 1, 0, 0, 0,
		64, 65, 5, 48, 0, 0, 65, 66, 7, 9, 0, 0, 66, 68, 1, 0, 0, 0, 67, 69, 7,
		10, 0, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70,
		71, 1, 0, 0, 0, 71, 20, 1, 0, 0, 0, 72, 73, 5, 48, 0, 0, 73, 74, 7, 11,
		0, 0, 74, 76, 1, 0, 0, 0, 75, 77, 7, 12, 0, 0, 76, 75, 1, 0, 0, 0, 77,
		78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 22, 1, 0, 0,
		0, 80, 82, 5, 45, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 84,
		1, 0, 0, 0, 83, 85, 7, 13, 0, 0, 84, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 24, 1, 0, 0, 0, 88, 92, 5,
		59, 0, 0, 89, 91, 8, 14, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0,
		92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1,
		0, 0, 0, 95, 96, 6, 12, 0, 0, 96, 26, 1, 0, 0, 0, 97, 99, 7, 14, 0, 0,
		98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 28, 1, 0, 0, 0, 102, 103, 7, 15, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 105, 6, 14, 0, 0, 105, 30, 1, 0, 0, 0, 9, 0, 57, 62, 70, 78,
		81, 86, 92, 100, 1, 6, 0, 0,
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
	DD8ASMLexerREG     = 8
	DD8ASMLexerNAME    = 9
	DD8ASMLexerHEX_NUM = 10
	DD8ASMLexerBIN_NUM = 11
	DD8ASMLexerDEC_NUM = 12
	DD8ASMLexerCOMMENT = 13
	DD8ASMLexerEOL     = 14
	DD8ASMLexerWS      = 15
)
