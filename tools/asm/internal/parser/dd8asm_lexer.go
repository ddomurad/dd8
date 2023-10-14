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
		"", "'('", "')'", "','", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "NAME", "PREP_NAME", "HEX_NUM", "BIN_NUM", "DEC_NUM",
		"COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "NAME", "PREP_NAME", "HEX_NUM", "BIN_NUM",
		"DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 86, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 4, 4, 35, 8, 4, 11, 4, 12, 4, 36, 1, 5, 1, 5, 4, 5, 41, 8, 5, 11,
		5, 12, 5, 42, 1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 49, 8, 6, 11, 6, 12, 6, 50,
		1, 7, 1, 7, 1, 7, 1, 7, 4, 7, 57, 8, 7, 11, 7, 12, 7, 58, 1, 8, 3, 8, 62,
		8, 8, 1, 8, 4, 8, 65, 8, 8, 11, 8, 12, 8, 66, 1, 9, 1, 9, 5, 9, 71, 8,
		9, 10, 9, 12, 9, 74, 9, 9, 1, 9, 1, 9, 1, 10, 4, 10, 79, 8, 10, 11, 10,
		12, 10, 80, 1, 11, 1, 11, 1, 11, 1, 11, 0, 0, 12, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 1, 0, 8, 4,
		0, 45, 45, 65, 90, 95, 95, 97, 122, 2, 0, 88, 88, 120, 120, 4, 0, 48, 57,
		65, 70, 95, 95, 97, 102, 2, 0, 66, 66, 98, 98, 2, 0, 48, 49, 95, 95, 2,
		0, 48, 57, 95, 95, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 93, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		1, 25, 1, 0, 0, 0, 3, 27, 1, 0, 0, 0, 5, 29, 1, 0, 0, 0, 7, 31, 1, 0, 0,
		0, 9, 34, 1, 0, 0, 0, 11, 38, 1, 0, 0, 0, 13, 44, 1, 0, 0, 0, 15, 52, 1,
		0, 0, 0, 17, 61, 1, 0, 0, 0, 19, 68, 1, 0, 0, 0, 21, 78, 1, 0, 0, 0, 23,
		82, 1, 0, 0, 0, 25, 26, 5, 40, 0, 0, 26, 2, 1, 0, 0, 0, 27, 28, 5, 41,
		0, 0, 28, 4, 1, 0, 0, 0, 29, 30, 5, 44, 0, 0, 30, 6, 1, 0, 0, 0, 31, 32,
		5, 61, 0, 0, 32, 8, 1, 0, 0, 0, 33, 35, 7, 0, 0, 0, 34, 33, 1, 0, 0, 0,
		35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 10, 1,
		0, 0, 0, 38, 40, 5, 46, 0, 0, 39, 41, 7, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41,
		42, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 12, 1, 0, 0,
		0, 44, 45, 5, 48, 0, 0, 45, 46, 7, 1, 0, 0, 46, 48, 1, 0, 0, 0, 47, 49,
		7, 2, 0, 0, 48, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0,
		50, 51, 1, 0, 0, 0, 51, 14, 1, 0, 0, 0, 52, 53, 5, 48, 0, 0, 53, 54, 7,
		3, 0, 0, 54, 56, 1, 0, 0, 0, 55, 57, 7, 4, 0, 0, 56, 55, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 16, 1, 0, 0,
		0, 60, 62, 5, 45, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64,
		1, 0, 0, 0, 63, 65, 7, 5, 0, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0,
		66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 18, 1, 0, 0, 0, 68, 72, 5,
		59, 0, 0, 69, 71, 8, 6, 0, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72,
		70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 72, 1, 0, 0,
		0, 75, 76, 6, 9, 0, 0, 76, 20, 1, 0, 0, 0, 77, 79, 7, 6, 0, 0, 78, 77,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0,
		81, 22, 1, 0, 0, 0, 82, 83, 7, 7, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 6,
		11, 0, 0, 85, 24, 1, 0, 0, 0, 9, 0, 36, 42, 50, 58, 61, 66, 72, 80, 1,
		6, 0, 0,
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
	DD8ASMLexerT__0      = 1
	DD8ASMLexerT__1      = 2
	DD8ASMLexerT__2      = 3
	DD8ASMLexerT__3      = 4
	DD8ASMLexerNAME      = 5
	DD8ASMLexerPREP_NAME = 6
	DD8ASMLexerHEX_NUM   = 7
	DD8ASMLexerBIN_NUM   = 8
	DD8ASMLexerDEC_NUM   = 9
	DD8ASMLexerCOMMENT   = 10
	DD8ASMLexerEOL       = 11
	DD8ASMLexerWS        = 12
)
