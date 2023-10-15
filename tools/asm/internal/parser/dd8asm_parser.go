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
		"", "'('", "')'", "':='", "','", "'['", "']'", "':'", "'.def'", "'.org'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "P_DEF", "P_ORG", "REG", "NAME", "HEX_NUM",
		"BIN_NUM", "DEC_NUM", "COMMENT", "EOL", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "statement", "prep_instruction", "prep_def_args", "prep_def_arg_lines",
		"instruction", "arglist_p", "arglist", "argument", "num", "reg", "name",
		"label",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 101, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 53, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 62, 8, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 76, 8, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 85, 8, 7, 1, 8, 1, 8, 1,
		8, 3, 8, 90, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		0, 1, 1, 0, 12, 14, 100, 0, 29, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 52, 1,
		0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 75, 1, 0, 0, 0, 12,
		77, 1, 0, 0, 0, 14, 81, 1, 0, 0, 0, 16, 89, 1, 0, 0, 0, 18, 91, 1, 0, 0,
		0, 20, 93, 1, 0, 0, 0, 22, 95, 1, 0, 0, 0, 24, 97, 1, 0, 0, 0, 26, 28,
		3, 2, 1, 0, 27, 26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0,
		29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5,
		0, 0, 1, 33, 1, 1, 0, 0, 0, 34, 38, 3, 10, 5, 0, 35, 38, 3, 4, 2, 0, 36,
		38, 3, 24, 12, 0, 37, 34, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 36, 1, 0,
		0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 5, 16, 0, 0, 40,
		3, 1, 0, 0, 0, 41, 42, 5, 9, 0, 0, 42, 53, 3, 16, 8, 0, 43, 44, 5, 8, 0,
		0, 44, 45, 5, 1, 0, 0, 45, 46, 5, 16, 0, 0, 46, 47, 3, 8, 4, 0, 47, 48,
		5, 16, 0, 0, 48, 49, 5, 2, 0, 0, 49, 53, 1, 0, 0, 0, 50, 51, 5, 8, 0, 0,
		51, 53, 3, 6, 3, 0, 52, 41, 1, 0, 0, 0, 52, 43, 1, 0, 0, 0, 52, 50, 1,
		0, 0, 0, 53, 5, 1, 0, 0, 0, 54, 55, 3, 22, 11, 0, 55, 56, 5, 3, 0, 0, 56,
		57, 3, 16, 8, 0, 57, 7, 1, 0, 0, 0, 58, 61, 3, 6, 3, 0, 59, 60, 5, 16,
		0, 0, 60, 62, 3, 8, 4, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 9,
		1, 0, 0, 0, 63, 76, 3, 22, 11, 0, 64, 65, 3, 22, 11, 0, 65, 66, 3, 12,
		6, 0, 66, 76, 1, 0, 0, 0, 67, 68, 3, 22, 11, 0, 68, 69, 3, 14, 7, 0, 69,
		76, 1, 0, 0, 0, 70, 71, 3, 22, 11, 0, 71, 72, 3, 12, 6, 0, 72, 73, 5, 4,
		0, 0, 73, 74, 3, 14, 7, 0, 74, 76, 1, 0, 0, 0, 75, 63, 1, 0, 0, 0, 75,
		64, 1, 0, 0, 0, 75, 67, 1, 0, 0, 0, 75, 70, 1, 0, 0, 0, 76, 11, 1, 0, 0,
		0, 77, 78, 5, 5, 0, 0, 78, 79, 3, 14, 7, 0, 79, 80, 5, 6, 0, 0, 80, 13,
		1, 0, 0, 0, 81, 84, 3, 16, 8, 0, 82, 83, 5, 4, 0, 0, 83, 85, 3, 14, 7,
		0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 15, 1, 0, 0, 0, 86, 90,
		3, 18, 9, 0, 87, 90, 3, 20, 10, 0, 88, 90, 3, 22, 11, 0, 89, 86, 1, 0,
		0, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 17, 1, 0, 0, 0, 91, 92,
		7, 0, 0, 0, 92, 19, 1, 0, 0, 0, 93, 94, 5, 10, 0, 0, 94, 21, 1, 0, 0, 0,
		95, 96, 5, 11, 0, 0, 96, 23, 1, 0, 0, 0, 97, 98, 5, 11, 0, 0, 98, 99, 5,
		7, 0, 0, 99, 25, 1, 0, 0, 0, 7, 29, 37, 52, 61, 75, 84, 89,
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
	DD8ASMParserP_DEF   = 8
	DD8ASMParserP_ORG   = 9
	DD8ASMParserREG     = 10
	DD8ASMParserNAME    = 11
	DD8ASMParserHEX_NUM = 12
	DD8ASMParserBIN_NUM = 13
	DD8ASMParserDEC_NUM = 14
	DD8ASMParserCOMMENT = 15
	DD8ASMParserEOL     = 16
	DD8ASMParserWS      = 17
)

// DD8ASMParser rules.
const (
	DD8ASMParserRULE_prog               = 0
	DD8ASMParserRULE_statement          = 1
	DD8ASMParserRULE_prep_instruction   = 2
	DD8ASMParserRULE_prep_def_args      = 3
	DD8ASMParserRULE_prep_def_arg_lines = 4
	DD8ASMParserRULE_instruction        = 5
	DD8ASMParserRULE_arglist_p          = 6
	DD8ASMParserRULE_arglist            = 7
	DD8ASMParserRULE_argument           = 8
	DD8ASMParserRULE_num                = 9
	DD8ASMParserRULE_reg                = 10
	DD8ASMParserRULE_name               = 11
	DD8ASMParserRULE_label              = 12
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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68352) != 0 {
		{
			p.SetState(26)
			p.Statement()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
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
	EOL() antlr.TerminalNode
	Instruction() IInstructionContext
	Prep_instruction() IPrep_instructionContext
	Label() ILabelContext

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
	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(34)
			p.Instruction()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(35)
			p.Prep_instruction()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 3 {
		{
			p.SetState(36)
			p.Label()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(39)
		p.Match(DD8ASMParserEOL)
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

// IPrep_instructionContext is an interface to support dynamic dispatch.
type IPrep_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	P_ORG() antlr.TerminalNode
	Argument() IArgumentContext
	P_DEF() antlr.TerminalNode
	AllEOL() []antlr.TerminalNode
	EOL(i int) antlr.TerminalNode
	Prep_def_arg_lines() IPrep_def_arg_linesContext
	Prep_def_args() IPrep_def_argsContext

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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(DD8ASMParserP_ORG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Argument()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Match(DD8ASMParserP_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(DD8ASMParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Prep_def_arg_lines()
		}
		{
			p.SetState(47)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(DD8ASMParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(DD8ASMParserP_DEF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Prep_def_args()
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
	p.EnterRule(localctx, 6, DD8ASMParserRULE_prep_def_args)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Name()
	}
	{
		p.SetState(55)
		p.Match(DD8ASMParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
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
	p.EnterRule(localctx, 8, DD8ASMParserRULE_prep_def_arg_lines)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Prep_def_args()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(59)
			p.Match(DD8ASMParserEOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
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
	p.EnterRule(localctx, 10, DD8ASMParserRULE_instruction)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Name()
		}
		{
			p.SetState(65)
			p.Arglist_p()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Name()
		}
		{
			p.SetState(68)
			p.Arglist()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.Name()
		}
		{
			p.SetState(71)
			p.Arglist_p()
		}
		{
			p.SetState(72)
			p.Match(DD8ASMParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Arglist()
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
	p.EnterRule(localctx, 12, DD8ASMParserRULE_arglist_p)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(DD8ASMParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Arglist()
	}
	{
		p.SetState(79)
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
	p.EnterRule(localctx, 14, DD8ASMParserRULE_arglist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Argument()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DD8ASMParserT__3 {
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
			p.Arglist()
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

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Num() INumContext
	Reg() IRegContext
	Name() INameContext

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

func (s *ArgumentContext) Num() INumContext {
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

func (s *ArgumentContext) Reg() IRegContext {
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

func (s *ArgumentContext) Name() INameContext {
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
	p.EnterRule(localctx, 16, DD8ASMParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DD8ASMParserHEX_NUM, DD8ASMParserBIN_NUM, DD8ASMParserDEC_NUM:
		{
			p.SetState(86)
			p.Num()
		}

	case DD8ASMParserREG:
		{
			p.SetState(87)
			p.Reg()
		}

	case DD8ASMParserNAME:
		{
			p.SetState(88)
			p.Name()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 18, DD8ASMParserRULE_num)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28672) != 0) {
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
	p.EnterRule(localctx, 20, DD8ASMParserRULE_reg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
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
	p.EnterRule(localctx, 22, DD8ASMParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
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
	p.EnterRule(localctx, 24, DD8ASMParserRULE_label)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(DD8ASMParserNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(DD8ASMParserT__6)
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
