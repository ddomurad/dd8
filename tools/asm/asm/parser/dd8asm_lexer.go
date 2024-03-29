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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "P_DEF", "P_TMPL",
		"P_REP", "P_ORG", "P_INC", "P_DB", "P_DW", "P_BYTE", "P_WORD", "REG",
		"STR", "HEX_NUM", "BIN_NUM", "DEC_NUM", "NAME", "COMMENT", "EOL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 241, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 187, 8, 33, 10, 33, 12, 33, 190,
		9, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 4, 34, 198, 8, 34, 11,
		34, 12, 34, 199, 1, 35, 1, 35, 1, 35, 1, 35, 4, 35, 206, 8, 35, 11, 35,
		12, 35, 207, 1, 36, 4, 36, 211, 8, 36, 11, 36, 12, 36, 212, 1, 37, 1, 37,
		4, 37, 217, 8, 37, 11, 37, 12, 37, 218, 1, 37, 3, 37, 222, 8, 37, 1, 38,
		1, 38, 5, 38, 226, 8, 38, 10, 38, 12, 38, 229, 9, 38, 1, 38, 1, 38, 1,
		39, 4, 39, 234, 8, 39, 11, 39, 12, 39, 235, 1, 40, 1, 40, 1, 40, 1, 40,
		0, 0, 41, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 1, 0, 28, 2, 0, 76, 76, 108, 108, 2,
		0, 72, 72, 104, 104, 2, 0, 68, 68, 100, 100, 2, 0, 69, 69, 101, 101, 2,
		0, 70, 70, 102, 102, 2, 0, 84, 84, 116, 116, 2, 0, 77, 77, 109, 109, 2,
		0, 80, 80, 112, 112, 2, 0, 82, 82, 114, 114, 2, 0, 79, 79, 111, 111, 2,
		0, 71, 71, 103, 103, 2, 0, 73, 73, 105, 105, 2, 0, 78, 78, 110, 110, 2,
		0, 67, 67, 99, 99, 2, 0, 66, 66, 98, 98, 2, 0, 87, 87, 119, 119, 2, 0,
		89, 89, 121, 121, 4, 0, 65, 65, 88, 90, 97, 97, 120, 122, 3, 0, 10, 10,
		13, 13, 34, 34, 2, 0, 88, 88, 120, 120, 4, 0, 48, 57, 65, 70, 95, 95, 97,
		102, 2, 0, 48, 49, 95, 95, 2, 0, 48, 57, 95, 95, 3, 0, 65, 90, 95, 95,
		97, 122, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 66, 87, 98,
		119, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 249, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0,
		0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0,
		0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0,
		0, 0, 0, 81, 1, 0, 0, 0, 1, 83, 1, 0, 0, 0, 3, 85, 1, 0, 0, 0, 5, 87, 1,
		0, 0, 0, 7, 89, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 93, 1, 0, 0, 0, 13,
		95, 1, 0, 0, 0, 15, 98, 1, 0, 0, 0, 17, 100, 1, 0, 0, 0, 19, 102, 1, 0,
		0, 0, 21, 104, 1, 0, 0, 0, 23, 107, 1, 0, 0, 0, 25, 110, 1, 0, 0, 0, 27,
		112, 1, 0, 0, 0, 29, 114, 1, 0, 0, 0, 31, 116, 1, 0, 0, 0, 33, 118, 1,
		0, 0, 0, 35, 120, 1, 0, 0, 0, 37, 123, 1, 0, 0, 0, 39, 126, 1, 0, 0, 0,
		41, 128, 1, 0, 0, 0, 43, 130, 1, 0, 0, 0, 45, 132, 1, 0, 0, 0, 47, 134,
		1, 0, 0, 0, 49, 139, 1, 0, 0, 0, 51, 145, 1, 0, 0, 0, 53, 150, 1, 0, 0,
		0, 55, 155, 1, 0, 0, 0, 57, 160, 1, 0, 0, 0, 59, 164, 1, 0, 0, 0, 61, 168,
		1, 0, 0, 0, 63, 174, 1, 0, 0, 0, 65, 180, 1, 0, 0, 0, 67, 182, 1, 0, 0,
		0, 69, 193, 1, 0, 0, 0, 71, 201, 1, 0, 0, 0, 73, 210, 1, 0, 0, 0, 75, 221,
		1, 0, 0, 0, 77, 223, 1, 0, 0, 0, 79, 233, 1, 0, 0, 0, 81, 237, 1, 0, 0,
		0, 83, 84, 5, 40, 0, 0, 84, 2, 1, 0, 0, 0, 85, 86, 5, 41, 0, 0, 86, 4,
		1, 0, 0, 0, 87, 88, 5, 44, 0, 0, 88, 6, 1, 0, 0, 0, 89, 90, 5, 64, 0, 0,
		90, 8, 1, 0, 0, 0, 91, 92, 5, 123, 0, 0, 92, 10, 1, 0, 0, 0, 93, 94, 5,
		125, 0, 0, 94, 12, 1, 0, 0, 0, 95, 96, 5, 58, 0, 0, 96, 97, 5, 61, 0, 0,
		97, 14, 1, 0, 0, 0, 98, 99, 5, 91, 0, 0, 99, 16, 1, 0, 0, 0, 100, 101,
		5, 93, 0, 0, 101, 18, 1, 0, 0, 0, 102, 103, 5, 126, 0, 0, 103, 20, 1, 0,
		0, 0, 104, 105, 5, 46, 0, 0, 105, 106, 7, 0, 0, 0, 106, 22, 1, 0, 0, 0,
		107, 108, 5, 46, 0, 0, 108, 109, 7, 1, 0, 0, 109, 24, 1, 0, 0, 0, 110,
		111, 5, 42, 0, 0, 111, 26, 1, 0, 0, 0, 112, 113, 5, 47, 0, 0, 113, 28,
		1, 0, 0, 0, 114, 115, 5, 37, 0, 0, 115, 30, 1, 0, 0, 0, 116, 117, 5, 43,
		0, 0, 117, 32, 1, 0, 0, 0, 118, 119, 5, 45, 0, 0, 119, 34, 1, 0, 0, 0,
		120, 121, 5, 60, 0, 0, 121, 122, 5, 60, 0, 0, 122, 36, 1, 0, 0, 0, 123,
		124, 5, 62, 0, 0, 124, 125, 5, 62, 0, 0, 125, 38, 1, 0, 0, 0, 126, 127,
		5, 38, 0, 0, 127, 40, 1, 0, 0, 0, 128, 129, 5, 94, 0, 0, 129, 42, 1, 0,
		0, 0, 130, 131, 5, 124, 0, 0, 131, 44, 1, 0, 0, 0, 132, 133, 5, 58, 0,
		0, 133, 46, 1, 0, 0, 0, 134, 135, 5, 46, 0, 0, 135, 136, 7, 2, 0, 0, 136,
		137, 7, 3, 0, 0, 137, 138, 7, 4, 0, 0, 138, 48, 1, 0, 0, 0, 139, 140, 5,
		46, 0, 0, 140, 141, 7, 5, 0, 0, 141, 142, 7, 6, 0, 0, 142, 143, 7, 7, 0,
		0, 143, 144, 7, 0, 0, 0, 144, 50, 1, 0, 0, 0, 145, 146, 5, 46, 0, 0, 146,
		147, 7, 8, 0, 0, 147, 148, 7, 3, 0, 0, 148, 149, 7, 7, 0, 0, 149, 52, 1,
		0, 0, 0, 150, 151, 5, 46, 0, 0, 151, 152, 7, 9, 0, 0, 152, 153, 7, 8, 0,
		0, 153, 154, 7, 10, 0, 0, 154, 54, 1, 0, 0, 0, 155, 156, 5, 46, 0, 0, 156,
		157, 7, 11, 0, 0, 157, 158, 7, 12, 0, 0, 158, 159, 7, 13, 0, 0, 159, 56,
		1, 0, 0, 0, 160, 161, 5, 46, 0, 0, 161, 162, 7, 2, 0, 0, 162, 163, 7, 14,
		0, 0, 163, 58, 1, 0, 0, 0, 164, 165, 5, 46, 0, 0, 165, 166, 7, 2, 0, 0,
		166, 167, 7, 15, 0, 0, 167, 60, 1, 0, 0, 0, 168, 169, 5, 46, 0, 0, 169,
		170, 7, 14, 0, 0, 170, 171, 7, 16, 0, 0, 171, 172, 7, 5, 0, 0, 172, 173,
		7, 3, 0, 0, 173, 62, 1, 0, 0, 0, 174, 175, 5, 46, 0, 0, 175, 176, 7, 15,
		0, 0, 176, 177, 7, 9, 0, 0, 177, 178, 7, 8, 0, 0, 178, 179, 7, 2, 0, 0,
		179, 64, 1, 0, 0, 0, 180, 181, 7, 17, 0, 0, 181, 66, 1, 0, 0, 0, 182, 188,
		5, 34, 0, 0, 183, 184, 5, 92, 0, 0, 184, 187, 5, 34, 0, 0, 185, 187, 8,
		18, 0, 0, 186, 183, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0,
		0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0, 190,
		188, 1, 0, 0, 0, 191, 192, 5, 34, 0, 0, 192, 68, 1, 0, 0, 0, 193, 194,
		5, 48, 0, 0, 194, 195, 7, 19, 0, 0, 195, 197, 1, 0, 0, 0, 196, 198, 7,
		20, 0, 0, 197, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 197, 1, 0, 0,
		0, 199, 200, 1, 0, 0, 0, 200, 70, 1, 0, 0, 0, 201, 202, 5, 48, 0, 0, 202,
		203, 7, 14, 0, 0, 203, 205, 1, 0, 0, 0, 204, 206, 7, 21, 0, 0, 205, 204,
		1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0,
		0, 0, 208, 72, 1, 0, 0, 0, 209, 211, 7, 22, 0, 0, 210, 209, 1, 0, 0, 0,
		211, 212, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213,
		74, 1, 0, 0, 0, 214, 216, 7, 23, 0, 0, 215, 217, 7, 24, 0, 0, 216, 215,
		1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0,
		0, 0, 219, 222, 1, 0, 0, 0, 220, 222, 7, 25, 0, 0, 221, 214, 1, 0, 0, 0,
		221, 220, 1, 0, 0, 0, 222, 76, 1, 0, 0, 0, 223, 227, 5, 59, 0, 0, 224,
		226, 8, 26, 0, 0, 225, 224, 1, 0, 0, 0, 226, 229, 1, 0, 0, 0, 227, 225,
		1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230, 1, 0, 0, 0, 229, 227, 1, 0,
		0, 0, 230, 231, 6, 38, 0, 0, 231, 78, 1, 0, 0, 0, 232, 234, 7, 26, 0, 0,
		233, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235,
		236, 1, 0, 0, 0, 236, 80, 1, 0, 0, 0, 237, 238, 7, 27, 0, 0, 238, 239,
		1, 0, 0, 0, 239, 240, 6, 40, 0, 0, 240, 82, 1, 0, 0, 0, 10, 0, 186, 188,
		199, 207, 212, 218, 221, 227, 235, 1, 6, 0, 0,
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
	DD8ASMLexerT__7    = 8
	DD8ASMLexerT__8    = 9
	DD8ASMLexerT__9    = 10
	DD8ASMLexerT__10   = 11
	DD8ASMLexerT__11   = 12
	DD8ASMLexerT__12   = 13
	DD8ASMLexerT__13   = 14
	DD8ASMLexerT__14   = 15
	DD8ASMLexerT__15   = 16
	DD8ASMLexerT__16   = 17
	DD8ASMLexerT__17   = 18
	DD8ASMLexerT__18   = 19
	DD8ASMLexerT__19   = 20
	DD8ASMLexerT__20   = 21
	DD8ASMLexerT__21   = 22
	DD8ASMLexerT__22   = 23
	DD8ASMLexerP_DEF   = 24
	DD8ASMLexerP_TMPL  = 25
	DD8ASMLexerP_REP   = 26
	DD8ASMLexerP_ORG   = 27
	DD8ASMLexerP_INC   = 28
	DD8ASMLexerP_DB    = 29
	DD8ASMLexerP_DW    = 30
	DD8ASMLexerP_BYTE  = 31
	DD8ASMLexerP_WORD  = 32
	DD8ASMLexerREG     = 33
	DD8ASMLexerSTR     = 34
	DD8ASMLexerHEX_NUM = 35
	DD8ASMLexerBIN_NUM = 36
	DD8ASMLexerDEC_NUM = 37
	DD8ASMLexerNAME    = 38
	DD8ASMLexerCOMMENT = 39
	DD8ASMLexerEOL     = 40
	DD8ASMLexerWS      = 41
)
