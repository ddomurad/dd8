package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ddomurad/dd8/tools/asm/pkg/parser"
)

type ASTNumber int64
type ASTRegister string
type ASTString string
type ASTName string

type ASTOperand struct {
	Value    any
	Indirect bool
}

type ASTOperands []ASTOperand

func (o ASTOperand) IsNumber() bool {
	_, ok := o.Value.(ASTNumber)
	return ok
}

func (o ASTOperand) IsRegister() bool {
	_, ok := o.Value.(ASTRegister)
	return ok
}

func (o ASTOperand) IsName() bool {
	_, ok := o.Value.(ASTName)
	return ok
}

func (o ASTOperand) Number() (ASTNumber, bool) {
	v, ok := o.Value.(ASTNumber)
	return v, ok
}

func (o ASTOperand) String() (ASTString, bool) {
	v, ok := o.Value.(ASTString)
	return v, ok
}

func (o ASTOperand) Register() (ASTRegister, bool) {
	v, ok := o.Value.(ASTRegister)
	return v, ok
}

func (o ASTOperand) Name() (ASTName, bool) {
	v, ok := o.Value.(ASTName)
	return v, ok
}

func (ol ASTOperands) Number(index int) (ASTNumber, bool) {
	if len(ol) <= index {
		return ASTNumber(0), false
	}
	v, ok := ol[index].Value.(ASTNumber)
	return v, ok
}

func (ol ASTOperands) Register(index int) (ASTRegister, bool) {
	if len(ol) <= index {
		return ASTRegister('0'), false
	}
	v, ok := ol[index].Value.(ASTRegister)
	return v, ok
}

func (ol ASTOperands) Name(index int) (ASTName, bool) {
	if len(ol) <= index {
		return ASTName(""), false
	}
	v, ok := ol[index].Value.(ASTName)
	return v, ok
}

func (ol ASTOperands) String(index int) (ASTString, bool) {
	if len(ol) <= index {
		return ASTString(""), false
	}
	v, ok := ol[index].Value.(ASTString)
	return v, ok
}

func (ol ASTOperands) Indirect(index int) bool {
	if len(ol) <= index {
		return false
	}
	return ol[index].Indirect
}

type SrcPointer struct {
	Name string
	Line int
}

func (p SrcPointer) String() string {
	return fmt.Sprintf("%s:%d", p.Name, p.Line)
}

type ASTStatementType string

const (
	ASTStatementTypeInstruction ASTStatementType = "instr"
	ASTStatementTypeOrigin      ASTStatementType = ".org"
	ASTStatementTypePrepDefine  ASTStatementType = ".def"
	ASTStatementTypeInclude     ASTStatementType = ".inc"
	ASTStatementTypeDataByte    ASTStatementType = ".db"
	ASTStatementTypeDataWord    ASTStatementType = ".dw"
	ASTStatementTypeLabel       ASTStatementType = "label"
)

type ASTStatement struct {
	Type ASTStatementType

	Name     string
	OpCode   ASTName
	Operands ASTOperands

	SrcPointer SrcPointer
}

func (s ASTStatement) Copy() ASTStatement {
	to := s.Operands
	s.Operands = make(ASTOperands, len(s.Operands))
	copy(s.Operands, to)
	return s
}

type AST struct {
	Statements []ASTStatement
	Errors     AssemblerError
}

func (ast *AST) HasLabel(name string) bool {
	for _, s := range ast.Statements {
		if s.Type == ASTStatementTypeLabel {
			if s.Name == name {
				return true
			}
		}
	}
	return false
}

func (t *AST) Copy() *AST {
	newAst := AST{
		Statements: make([]ASTStatement, len(t.Statements)),
	}

	for i, s := range t.Statements {
		newAst.Statements[i] = s.Copy()
	}

	return &newAst
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
		case []ASTStatement:
			for _, st := range tres {
				ast.Statements = append(ast.Statements, st)
			}
		case ASTStatement:
			ast.Statements = append(ast.Statements, tres)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
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

	inst := ASTStatement{
		Type:     ASTStatementTypeInstruction,
		OpCode:   opcodeStr,
		Operands: ASTOperands{},
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
		operands, ok := visited.(ASTOperands)
		if !ok {
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
		inst.Operands = append(inst.Operands, operands...)
	}
	return inst
}

func (v *progVisitor) VisitArglist_p(ctx *parser.Arglist_pContext) interface{} {
	args, ok := v.Visit(ctx.Arglist().(*parser.ArglistContext)).(ASTOperands)
	if !ok {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}
	for i := range args {
		args[i].Indirect = true
	}
	return args
}

func (v *progVisitor) VisitArglist(ctx *parser.ArglistContext) interface{} {
	args := ASTOperands{}
	for _, ch := range ctx.GetChildren() {
		cr := v.Visit(ch.(antlr.ParseTree))
		if cr == nil {
			continue
		}

		switch tr := cr.(type) {
		case ASTOperands:
			args = append(args, tr...)
		case ASTOperand:
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
	return ASTOperand{Value: cr}
}

func (v *progVisitor) VisitStr(ctx *parser.StrContext) interface{} {
	str := ctx.STR().GetText()
	str = strings.ReplaceAll(str, "\"", "")
	return ASTString(str)
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
	return ASTStatement{
		Type: ASTStatementTypeLabel,
		Name: ctx.NAME().GetText(),
		SrcPointer: SrcPointer{
			Name: v.srcName,
			Line: ctx.GetStart().GetLine(),
		},
	}
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
	case ".inc":
		return v.buildPrepInc(ctx, children[1:])
	case ".org":
		return v.buildPrepOrigin(ctx, children[1:])
	case ".def":
		return v.buildPrepDefine(ctx, children[1:])
	case ".db":
		return v.buildPrepByte(ctx, 1, children[1:])
	case ".dw":
		return v.buildPrepByte(ctx, 2, children[1:])
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

	return []ASTStatement{
		{
			Type:     ASTStatementTypePrepDefine,
			Name:     ctx.Name().GetText(),
			Operands: ASTOperands{v.VisitArgument(argCtx).(ASTOperand)},
			SrcPointer: SrcPointer{
				Name: v.srcName,
				Line: ctx.GetStart().GetLine(),
			},
		}}
}

func (v *progVisitor) VisitPrep_def_arg_lines(ctx *parser.Prep_def_arg_linesContext) interface{} {
	children := ctx.GetChildren()
	defs := []ASTStatement{}

	for _, child := range children {
		vc := v.Visit(child.(antlr.ParseTree))
		if vc == nil {
			continue
		}
		switch tc := vc.(type) {
		case []ASTStatement:
			defs = append(defs, tc...)
		case ASTStatement:
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
	op, ok := vc.(ASTOperand)
	if !ok {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}
	return ASTStatement{
		Type:     ASTStatementTypeOrigin,
		Operands: ASTOperands{op},
		SrcPointer: SrcPointer{
			Name: v.srcName,
			Line: ctx.GetStart().GetLine(),
		},
	}
}

func (v *progVisitor) buildPrepDefine(ctx *parser.Prep_instructionContext, children []antlr.Tree) interface{} {
	defs := make([]ASTStatement, 0, len(children))
	for _, c := range children {
		vc := v.Visit(c.(antlr.ParseTree))
		if vc == nil {
			continue
		}

		switch tv := vc.(type) {
		case []ASTStatement:
			defs = append(defs, tv...)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	return defs
}

func (v *progVisitor) buildPrepByte(ctx *parser.Prep_instructionContext, size int, children []antlr.Tree) interface{} {
	if len(children) != 1 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	vc := v.Visit(children[0].(antlr.ParseTree))
	operands, ok := vc.(ASTOperands)
	if !ok {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	defType := ASTStatementTypeDataByte
	if size == 2 {
		defType = ASTStatementTypeDataWord
	}

	return ASTStatement{
		Type:     defType,
		Operands: operands,
		SrcPointer: SrcPointer{
			Name: v.srcName,
			Line: ctx.GetStart().GetLine(),
		},
	}
}

func (v *progVisitor) buildPrepInc(ctx *parser.Prep_instructionContext, children []antlr.Tree) interface{} {
	if len(children) != 1 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	vc := v.Visit(children[0].(antlr.ParseTree))
	if vc == nil {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	return []ASTStatement{{
		Type:     ASTStatementTypeInclude,
		Operands: []ASTOperand{{Value: vc}},
		SrcPointer: SrcPointer{
			Name: v.srcName,
			Line: ctx.GetStart().GetLine(),
		}},
	}
}

func (v *progVisitor) statementStructureError(line int) {
	v.errorListener.ProgramError(line, "unexpected statement structure")
}

func applyDefsToOperand(operand ASTOperand, defs map[string]ASTOperand) (ASTOperand, bool) {
	nameOp, ok := operand.Name()
	if !ok {
		return ASTOperand{}, false
	}

	nv, ok := defs[string(nameOp)]
	return nv, ok
}

func applyDefsToOperands(operands ASTOperands, defs map[string]ASTOperand) bool {
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

func applyDefinitions(ast *AST, defs map[string]ASTOperand) bool {
	rok := false
	for i, st := range ast.Statements {
		rok = rok || applyDefsToOperands(st.Operands, defs)
		ast.Statements[i] = st
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
	defs := map[string]ASTOperand{}
	for _, st := range ast.Statements {

		if st.Type != ASTStatementTypePrepDefine {
			filtSts = append(filtSts, st)
			continue
		}
		_, ok := defs[st.Name]
		if ok {
			ast.Errors.Append(SourceError{
				Type:    SourceErrorTypeEvalError,
				SrcName: st.SrcPointer.Name,
				Line:    st.SrcPointer.Line,
				Msg:     fmt.Sprintf("label redefinition: '%s'", st.Name),
			})
		}

		defs[st.Name] = st.Operands[0]
	}

	ast.Statements = filtSts
	for applyDefinitions(ast, defs) {
	}
}

func PreprocessIncludes(ast *AST, reader SourceReader) bool {
	updated := false
	statements := make([]ASTStatement, 0, len(ast.Statements))
	for _, s := range ast.Statements {
		if s.Type == ASTStatementTypeInclude {
			updated = true
			srcName, ok := s.Operands.String(0)
			if !ok {
				ast.Errors.Append(SourceError{
					Type:    SourceErrorTypeEvalError,
					SrcName: s.SrcPointer.Name,
					Line:    s.SrcPointer.Line,
					Msg:     "include source name expected to be a string",
				})
			}
			incAst, err := CompileAST(string(srcName), reader)
			if err != nil {
				ast.Errors.Append(SourceError{
					Type:    SourceErrorTypeEvalError,
					SrcName: s.SrcPointer.Name,
					Line:    s.SrcPointer.Line,
					Msg:     fmt.Sprintf("failed to include src '%s'. err: %v", srcName, err),
				})
				return false
			}
			statements = append(statements, incAst.Statements...)
		} else {
			statements = append(statements, s)
		}
	}

	ast.Statements = statements
	return updated
}

func PreprocessAST(ast *AST, reader SourceReader) {
	for PreprocessIncludes(ast, reader) {
	}

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
