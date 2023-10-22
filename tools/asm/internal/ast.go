package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ddomurad/dd8/tools/asm/internal/parser"
)

type ASTNumber int64
type ASTRegister string
type ASTName string
type ASTParentheses []any
type ASTStatement any

type SrcPointer struct {
	Name string
	Line int
}

func (p SrcPointer) String() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Line)
}

type ASTLabel struct {
	Name string
}

type ASTOrigin struct {
	Address any
}

type ASTPrepDefine struct {
	Name  string
	Value any
}

type ASTInstruction struct {
	OpCode   ASTName
	Operands []any

	SrcPointer SrcPointer
}

type AST struct {
	Statements []ASTStatement
	Errors     AssemblerError
}

type progVisitor struct {
	parser.BaseDD8ASMVisitor

	srcName       string
	errorListener *ErrorListener
}

func newProgVisitor(srcName string, errorListener *ErrorListener) *progVisitor {
	return &progVisitor{
		srcName:       srcName,
		errorListener: errorListener,
	}
}

func (v *progVisitor) Visit(tree antlr.ParseTree) interface{} { return tree.Accept(v) }
func (v *progVisitor) VisitChildren(ctx antlr.RuleNode) interface{} {
	children := []any{}
	for _, n := range ctx.GetChildren() {
		children = append(children, v.Visit(n.(antlr.ParseTree)))
	}
	return children
}

func (v *progVisitor) VisitTerminal(ctx antlr.TerminalNode) interface{} {
	return nil
}

func (v *progVisitor) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	return nil
}

func (v *progVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	ast := AST{}
	for _, n := range ctx.GetChildren() {
		result := v.Visit(n.(antlr.ParseTree))
		if result == nil {
			continue
		}

		switch tres := result.(type) {
		case []interface{}:
			for _, st := range tres {
				ast.Statements = append(ast.Statements, st)
			}
		case interface{}:
			ast.Statements = append(ast.Statements, tres)
		}
	}

	return &ast
}

func (v *progVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	children := ctx.GetChildren()
	if len(children) != 1 && len(children) != 2 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	return v.Visit(children[0].(antlr.ParseTree))
}

func (v *progVisitor) VisitInstruction(ctx *parser.InstructionContext) interface{} {
	children := ctx.GetChildren()
	if len(children) == 0 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}
	opcodeResp := v.Visit(children[0].(antlr.ParseTree))
	opcodeStr, ok := opcodeResp.(ASTName)
	if !ok {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	inst := ASTInstruction{
		OpCode:   opcodeStr,
		Operands: []any{},
		SrcPointer: SrcPointer{
			Name: v.srcName,
			Line: ctx.GetStart().GetLine(),
		},
	}

	for _, child := range children[1:] {
		visited := v.Visit(child.(antlr.ParseTree))
		if visited == nil {
			continue
		}
		operands, ok := visited.([]interface{})
		if !ok {
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
		inst.Operands = append(inst.Operands, operands...)
	}

	// if len(children) > 2 {
	// 	v.statementStructureError(ctx.GetStart().GetLine())
	// 	return nil
	// }
	//
	// if len(children) == 2 {
	// 	inst.Operands, ok = v.Visit(children[1].(antlr.ParseTree)).([]interface{})
	// 	if !ok {
	// 		v.statementStructureError(ctx.GetStart().GetLine())
	// 		return nil
	// 	}
	// }

	return inst
}

func (v *progVisitor) VisitArglist_p(ctx *parser.Arglist_pContext) interface{} {
	args, ok := v.Visit(ctx.Arglist().(*parser.ArglistContext)).([]interface{})
	if !ok {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}
	pargs := make(ASTParentheses, len(args))
	copy(pargs, args)
	return []interface{}{pargs}
}

func (v *progVisitor) VisitArglist(ctx *parser.ArglistContext) interface{} {
	childrenRes := v.VisitChildren(ctx).([]interface{})
	args := []interface{}{}

	for _, cr := range childrenRes {
		if cr == nil {
			continue
		}
		switch tr := cr.(type) {
		case []interface{}:
			args = append(args, tr...)
		case interface{}:
			args = append(args, tr)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	return args
}

func (v *progVisitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	children := ctx.GetChildren()
	if len(children) != 1 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}
	cr := v.Visit(children[0].(antlr.ParseTree))
	return cr
}

func (v *progVisitor) VisitNum(ctx *parser.NumContext) interface{} {
	children := ctx.GetChildren()
	if len(children) != 1 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}
	terminalNode, ok := children[0].(antlr.TerminalNode)
	if !ok {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	nodeStr := terminalNode.GetText()
	numBytes := 0

	switch terminalNode.GetSymbol().GetTokenType() {
	case parser.DD8ASMParserHEX_NUM:
		nodeStr = nodeStr[2:] // skip 0x prefix
		numBytes = 16
	case parser.DD8ASMParserDEC_NUM:
		numBytes = 10
	case parser.DD8ASMParserBIN_NUM:
		nodeStr = nodeStr[2:] // skip 0b prefix
		numBytes = 2
	}

	nodeStr = strings.ReplaceAll(nodeStr, "_", "")
	num, err := strconv.ParseInt(nodeStr, numBytes, 64)
	if err != nil {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}
	return ASTNumber(num)
}

func (v *progVisitor) VisitName(ctx *parser.NameContext) interface{} {
	return ASTName(ctx.NAME().GetText())
}

func (v *progVisitor) VisitReg(ctx *parser.RegContext) interface{} {
	return ASTRegister(ctx.REG().GetText())
}

func (v *progVisitor) VisitLabel(ctx *parser.LabelContext) interface{} {
	return ASTLabel{Name: ctx.NAME().GetText()}
}

func (v *progVisitor) VisitPrep_instruction(ctx *parser.Prep_instructionContext) interface{} {
	children := ctx.GetChildren()
	if len(children) == 0 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	insType, ok := children[0].(antlr.TerminalNode)
	if !ok {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	prepInst := insType.GetText()
	switch prepInst {
	case ".org":
		return v.buildPrepOrigin(ctx, children[1:])
	case ".def":
		return v.buildPrepDefine(ctx, children[1:])
	}

	v.statementStructureError(ctx.GetStart().GetLine())
	return nil
}

func (v *progVisitor) VisitPrep_def_args(ctx *parser.Prep_def_argsContext) interface{} {
	argCtx, ok := ctx.Argument().(*parser.ArgumentContext)
	if !ok {
		v.errorListener.ProgramError(ctx.GetStart().GetLine(), "preprocessor argument missing")
		return nil
	}

	return []interface{}{
		ASTPrepDefine{
			Name:  ctx.Name().GetText(),
			Value: v.VisitArgument(argCtx),
		}}
}

func (v *progVisitor) VisitPrep_def_arg_lines(ctx *parser.Prep_def_arg_linesContext) interface{} {
	children := ctx.GetChildren()
	defs := []interface{}{}

	for _, child := range children {
		vc := v.Visit(child.(antlr.ParseTree))
		if vc == nil {
			continue
		}
		switch tc := vc.(type) {
		case []interface{}:
			defs = append(defs, tc...)
		case ASTPrepDefine:
			defs = append(defs, tc)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	return defs
}

func (v *progVisitor) buildPrepOrigin(ctx *parser.Prep_instructionContext, children []antlr.Tree) interface{} {
	if len(children) != 1 {
		v.errorListener.ProgramError(ctx.GetStart().GetLine(), "preprocessor argument missing")
	}

	vc := v.Visit(children[0].(antlr.ParseTree))
	return ASTOrigin{Address: vc}
}

func (v *progVisitor) buildPrepDefine(ctx *parser.Prep_instructionContext, children []antlr.Tree) interface{} {
	defs := make([]interface{}, 0, len(children))
	for _, c := range children {
		vc := v.Visit(c.(antlr.ParseTree))
		if vc == nil {
			continue
		}

		switch tv := vc.(type) {
		case []interface{}:
			defs = append(defs, tv...)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	return defs
}

func (v *progVisitor) statementStructureError(line int) {
	v.errorListener.ProgramError(line, "unexpected statement structure")
}

func applyDefsToOperand(operand any, defs map[string]any) (any, bool) {
	nameOp, ok := operand.(ASTName)
	if !ok {
		return nil, false
	}

	nv, ok := defs[string(nameOp)]
	return nv, ok
}

func applyDefsToOperands(operands []any, defs map[string]any) bool {
	rok := false
	for i, op := range operands {
		nv, ok := applyDefsToOperand(op, defs)
		if ok {
			operands[i] = nv
			rok = true
		}
	}

	return rok
}

func applyDefinitions(ast *AST, defs map[string]any) bool {
	rok := false
	for i, st := range ast.Statements {
		switch tst := st.(type) {
		case ASTOrigin:
			nv, ok := applyDefsToOperand(tst.Address, defs)
			if ok {
				tst.Address = nv
				ast.Statements[i] = tst
				rok = ok
			}
		case ASTInstruction:
			rok = rok || applyDefsToOperands(tst.Operands, defs)
			ast.Statements[i] = tst
		default:
			continue
		}
	}

	return rok
}

type ErrorListener struct {
	srcName string
	errors  AssemblerError
}

func NewErrorListener(srcName string) *ErrorListener {
	return &ErrorListener{
		srcName: srcName,
		errors:  NewAssemblerError(),
	}
}

// ReportAmbiguity implements antlr.ErrorListener.
func (*ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	panic("unimplemented")
}

// ReportAttemptingFullContext implements antlr.ErrorListener.
func (*ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	panic("unimplemented")
}

// ReportContextSensitivity implements antlr.ErrorListener.
func (*ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	panic("unimplemented")
}

// SyntaxError implements antlr.ErrorListener.
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, col int, msg string, e antlr.RecognitionException) {
	l.errors.Append(SourceError{
		Type:    SourceErrorTypeSyntaxError,
		SrcName: l.srcName,
		Line:    line,
		Col:     col,
		Msg:     msg,
	})
}

func (l *ErrorListener) ProgramError(line int, msg string) {
	l.errors.Append(SourceError{
		Type:    SourceErrorTypeUnexpectedStructureError,
		SrcName: l.srcName,
		Line:    line,
		Msg:     msg,
	})
}

func PreprocessDefinitions(ast *AST) {
	filtSts := make([]ASTStatement, 0, len(ast.Statements))
	defs := map[string]any{}
	for _, st := range ast.Statements {
		def, ok := st.(ASTPrepDefine)
		if !ok {
			filtSts = append(filtSts, st)
			continue
		}
		_, ok = defs[def.Name]
		if ok {
			ast.Errors.Append(SourceError{
				Type: SourceErrorTypeProgramError,
				Msg:  fmt.Sprintf("label redefinition: '%s'", def.Name),
			})
		}

		defs[def.Name] = def.Value
	}

	ast.Statements = filtSts
	for applyDefinitions(ast, defs) {
	}
}

func PreprocessAST(ast *AST) {
	PreprocessDefinitions(ast)
}

func ParseSrc(srcName string, src string) *AST {
	errorListener := NewErrorListener(srcName)
	input := antlr.NewInputStream(src)
	lexer := parser.NewDD8ASMLexer(input)
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewDD8ASMParser(stream)
	parser.AddErrorListener(errorListener)
	tree := parser.Prog()

	ast := newProgVisitor(srcName, errorListener).Visit(tree).(*AST)
	ast.Errors = errorListener.errors
	return ast
}

func CompileAST(srcName string, reader SourceReader) (*AST, error) {
	src, err := reader.ReadSourceFile(srcName)
	if err != nil {
		return nil, err
	}

	ast := ParseSrc(srcName, src)
	if ast.Errors.HasErrors() {
		return ast, ast.Errors
	}

	return ast, nil
}
