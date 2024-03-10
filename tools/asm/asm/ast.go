package asm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ddomurad/dd8/tools/asm/asm/parser"
)

type ASTNumber int64
type ASTRegister string
type ASTString string
type ASTName string
type ASTExpr struct {
	Left, Right any
	Operation   string
}

type ASTBlock struct {
	Statements []ASTStatement
}

type ASTOperand struct {
	Value    any
	Indirect bool
}

type ASTOperands []ASTOperand

func (o ASTOperand) Number(actx AssemblyContext) (ASTNumber, bool) {
	return EvaluateExprTo[ASTNumber](o, actx)
}

func (o ASTOperand) String(actx AssemblyContext) (ASTString, bool) {
	return EvaluateExprTo[ASTString](o, actx)
}

func (o ASTOperand) Register(actx AssemblyContext) (ASTRegister, bool) {
	return EvaluateExprTo[ASTRegister](o, actx)
}

func (o ASTOperand) Name() (ASTName, bool) {
	v, ok := o.Value.(ASTName)
	return v, ok
}

func (o ASTOperand) Names() ([]ASTName, bool) {
	v, ok := o.Value.([]ASTName)
	return v, ok
}

func (o ASTOperand) Expr() (ASTExpr, bool) {
	v, ok := o.Value.(ASTExpr)
	return v, ok
}

func (o ASTOperand) Block() (ASTBlock, bool) {
	v, ok := o.Value.(ASTBlock)
	return v, ok
}

func EvaluateExprTo[T any](o ASTOperand, actx AssemblyContext) (T, bool) {
	var def T
	switch tv := o.Value.(type) {
	case ASTExpr:
		ev, ok := EvaluateExpr(tv, actx)
		if !ok {
			return def, false
		}
		rv, ok := ev.(T)
		if !ok {
			return def, false
		}
		return rv, true
	case T:
		return tv, true
	}

	return def, false
}

func EvaluateExpr(expr ASTExpr, actx AssemblyContext) (any, bool) {
	var success bool
	lv, rv := expr.Left, expr.Right

	if le, ok := lv.(ASTExpr); ok {
		lv, success = EvaluateExpr(le, actx)
		if !success {
			return 0, false
		}
	}

	if re, ok := rv.(ASTExpr); ok {
		rv, success = EvaluateExpr(re, actx)
		if !success {
			return 0, false
		}
	}

	nv, nok := lv.(ASTName)
	if nok && rv == nil && expr.Operation == "" {
		if v, ok := actx.Labels[string(nv)]; ok {
			return ASTNumber(v), true
		}
		// v, ok := actx.Deffinitions[string(nv)]
		v, ok := actx.Deffinitions.Get(string(nv))
		return v, ok
	}

	ln, lnok := lv.(ASTNumber)
	ls, lsok := lv.(ASTString)
	rn, rnok := rv.(ASTNumber)
	rs, rsok := rv.(ASTString)

	// if (lnok && rsok) || (lsok && rnok) {
	// 	return 0, false //todo: better error reporting
	// }

	if lsok && rnok {
		switch expr.Operation {
		case "*":
			return ASTString(strings.Repeat(string(ls), int(rn))), true
		default:
			return "", false
		}
	}

	if lsok && !rsok {
		return ls, true
	}
	if lsok && rsok {
		switch expr.Operation {
		case "+":
			return ASTString(ls + rs), true
		default:
			return "", false
		}
	}

	if lnok && !rnok {
		switch expr.Operation {
		case ".l":
			return ASTNumber(ln & 0xff), true
		case ".h":
			return ASTNumber((ln >> 8) & 0xff), true
		case "()":
			return ln, true
		}
	}

	if !lnok && rnok {
		switch expr.Operation {
		case "~":
			return ASTNumber(^rn & 0xffff), true
		case "-":
			return -1 * ASTNumber(rn), true
		}
	}

	if lnok && rnok {
		switch expr.Operation {
		case "*":
			return ASTNumber(ln * rn), true
		case "/":
			return ASTNumber(ln / rn), true
		case "%":
			return ASTNumber(ln % rn), true
		case "+":
			return ASTNumber(ln + rn), true
		case "-":
			return ASTNumber(ln - rn), true
		case ">>":
			return ASTNumber(ln >> rn), true
		case "<<":
			return ASTNumber(ln << rn), true
		case "&":
			return ASTNumber(ln & rn), true
		case "^":
			return ASTNumber(ln ^ rn), true
		case "|":
			return ASTNumber(ln | rn), true
		}
	}

	return 0, true
}

func (ol ASTOperands) Number(index int, actx AssemblyContext) (ASTNumber, bool) {
	if len(ol) <= index {
		return 0, false
	}
	return ol[index].Number(actx)
}

func (ol ASTOperands) Register(index int, actx AssemblyContext) (ASTRegister, bool) {
	if len(ol) <= index {
		return "", false
	}
	return ol[index].Register(actx)
}

func (ol ASTOperands) Name(index int) (ASTName, bool) {
	if len(ol) <= index {
		return "", false
	}
	v, ok := ol[index].Value.(ASTName)
	return v, ok
}

func (ol ASTOperands) String(index int) (ASTString, bool) {
	if len(ol) <= index {
		return "", false
	}
	v, ok := ol[index].Value.(ASTString)
	return v, ok
}

func (ol ASTOperands) Expr(index int) (ASTExpr, bool) {
	if len(ol) <= index {
		return ASTExpr{}, false
	}
	v, ok := ol[index].Value.(ASTExpr)
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
	ASTStatementTypeInstruction     ASTStatementType = "instr"
	ASTStatementTypeOrigin          ASTStatementType = ".org"
	ASTStatementTypeInclude         ASTStatementType = ".inc"
	ASTStatementTypePrepDefine      ASTStatementType = ".def"
	ASTStatementTypePrepTemplateDef ASTStatementType = ".tmpl"
	ASTStatementTypePrepTemplateUse ASTStatementType = "@"
	ASTStatementTypeDataByte        ASTStatementType = ".db"
	ASTStatementTypeDataWord        ASTStatementType = ".dw"
	ASTStatementTypeLabel           ASTStatementType = "label"
	ASTStatementTypeSkipBytes       ASTStatementType = ".byte"
	ASTStatementTypeSkipWords       ASTStatementType = ".word"
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
			ast.Statements = append(ast.Statements, tres...)
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

func (v *progVisitor) VisitArglist_lines(ctx *parser.Arglist_linesContext) interface{} {
	children := ctx.GetChildren()
	operands := ASTOperands{}

	for _, child := range children {
		vc := v.Visit(child.(antlr.ParseTree))
		if vc == nil {
			continue
		}
		switch tc := vc.(type) {
		case ASTOperands:
			operands = append(operands, tc...)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	return operands
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

func (v *progVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	children := ctx.GetChildren()
	if len(children) == 1 {
		vc := v.Visit(children[0].(antlr.ParseTree))
		if vc == nil {
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
		if nameChild, ok := vc.(ASTName); ok {
			return ASTExpr{
				Left:      nameChild,
				Right:     nil,
				Operation: "",
			}
		}
		return vc
	}

	operands := make([]any, 0)
	operation := ""
	for _, child := range children {
		vc := v.Visit(child.(antlr.ParseTree))
		if vc == nil {
			tn, ok := child.(antlr.TerminalNode)
			if !ok {
				v.statementStructureError(ctx.GetStart().GetLine())
				return nil
			}
			operation = tn.GetText()
			continue
		}

		switch tc := vc.(type) {
		case ASTString:
			operands = append(operands, tc)
		case ASTNumber:
			operands = append(operands, tc)
		case ASTName:
			operands = append(operands, tc)
		case ASTExpr:
			operands = append(operands, tc)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	if len(operands) == 2 && operation != "" {
		return ASTExpr{
			Left:      operands[0],
			Right:     operands[1],
			Operation: operation,
		}
	}
	if len(operands) == 1 && (operation == "~" || operation == "-") {
		return ASTExpr{
			Right:     operands[0],
			Left:      nil,
			Operation: operation,
		}
	}
	if len(operands) == 1 && (operation == ".l" || operation == ".h") {
		return ASTExpr{
			Left:      operands[0],
			Right:     nil,
			Operation: operation,
		}
	}
	if len(operands) == 1 && operation == ")" {
		return ASTExpr{
			Left:      operands[0],
			Right:     nil,
			Operation: "()",
		}
	}
	v.statementStructureError(ctx.GetStart().GetLine())
	return nil
}

func (v *progVisitor) VisitStr(ctx *parser.StrContext) interface{} {
	str := ctx.STR().GetText()
	str, err := strconv.Unquote(str)
	if err != nil {
		v.errorListener.ProgramError(ctx.GetStart().GetLine(), err.Error())
	}
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
	case ".tmpl":
		return v.buildPrepTemplate(ctx, children[1:])
	case "@":
		return v.buildTemplateUsage(ctx, children[1:])
	case ".db":
		return v.buildPrepByte(ctx, ASTStatementTypeDataByte, children[1:])
	case ".dw":
		return v.buildPrepByte(ctx, ASTStatementTypeDataWord, children[1:])
	case ".byte":
		return v.buildSkipBytes(ctx, ASTStatementTypeSkipBytes, children[1:])
	case ".word":
		return v.buildSkipBytes(ctx, ASTStatementTypeSkipWords, children[1:])
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

func (v *progVisitor) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	children := ctx.GetChildren()
	names := make([]ASTName, 0)

	for _, child := range children {
		vc := v.Visit(child.(antlr.ParseTree))
		if vc == nil {
			continue
		}
		switch t := vc.(type) {
		case ASTName:
			names = append(names, t)
		case []ASTName:
			names = append(names, t...)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	return names
}

func (v *progVisitor) VisitTmpl_block(ctx *parser.Tmpl_blockContext) interface{} {
	children := ctx.GetChildren()
	_ = children
	block := ASTBlock{
		Statements: []ASTStatement{},
	}

	// todo: refactor this, similar to VisitProg
	for _, child := range children {
		vc := v.Visit(child.(antlr.ParseTree))
		if vc == nil {
			continue
		}
		switch tc := vc.(type) {
		case []ASTStatement:
			block.Statements = append(block.Statements, tc...)
		case ASTStatement:
			block.Statements = append(block.Statements, tc)
		default:
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
	}

	return block
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

func (v *progVisitor) buildTemplateUsage(ctx *parser.Prep_instructionContext, children []antlr.Tree) interface{} {
	var nameNode interface{}
	var arguments interface{}

	if len(children) == 4 {
		nameNode = v.Visit(children[0].(antlr.ParseTree))
		arguments = v.Visit(children[2].(antlr.ParseTree))
	} else if len(children) == 3 {
		nameNode = v.Visit(children[0].(antlr.ParseTree))
		arguments = ASTOperands{}
	} else {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	// todo: check if namenode is astname
	// todo: check if arguments are astname

	return ASTStatement{
		Type:     ASTStatementTypePrepTemplateUse,
		Name:     string(nameNode.(ASTName)),
		Operands: arguments.(ASTOperands),
		SrcPointer: SrcPointer{
			Name: v.srcName,
			Line: ctx.GetStart().GetLine(),
		},
	}
}

func (v *progVisitor) buildPrepTemplate(ctx *parser.Prep_instructionContext, children []antlr.Tree) interface{} {
	// todo: refactor this function

	var nameNode interface{}
	var arguments interface{}
	var body interface{}

	if len(children) == 5 {
		nameNode = v.Visit(children[0].(antlr.ParseTree))
		arguments = v.Visit(children[2].(antlr.ParseTree))
		body = v.Visit(children[4].(antlr.ParseTree))
	} else if len(children) == 4 {
		nameNode = v.Visit(children[0].(antlr.ParseTree))
		arguments = []ASTName{}
		body = v.Visit(children[3].(antlr.ParseTree))
	} else {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	// todo: check if namenode is astname
	// todo: check if arguments are astname

	return ASTStatement{
		Type: ASTStatementTypePrepTemplateDef,
		Name: string(nameNode.(ASTName)),
		Operands: []ASTOperand{
			{Value: arguments},
			{Value: body},
		},
		SrcPointer: SrcPointer{
			Name: v.srcName,
			Line: ctx.GetStart().GetLine(),
		},
	}
}

func (v *progVisitor) buildPrepByte(ctx *parser.Prep_instructionContext, defType ASTStatementType, children []antlr.Tree) interface{} {
	operands := make(ASTOperands, 0)

	for _, c := range children {
		vc := v.Visit(c.(antlr.ParseTree))
		if vc == nil {
			continue
		}

		co, ok := vc.(ASTOperands)
		if !ok {
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
		operands = append(operands, co...)
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

func (v *progVisitor) buildSkipBytes(ctx *parser.Prep_instructionContext, defType ASTStatementType, children []antlr.Tree) interface{} {
	if len(children) > 1 {
		v.statementStructureError(ctx.GetStart().GetLine())
		return nil
	}

	operands := make(ASTOperands, 1)

	if len(children) == 1 {
		vc := v.Visit(children[0].(antlr.ParseTree))
		if vc == nil {
			v.statementStructureError(ctx.GetStart().GetLine())
			return nil
		}
		operands[0] = ASTOperand{Value: vc}
	} else {
		operands[0] = ASTOperand{Value: ASTNumber(1)}
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
	panic("unimplemented ReportAmbiguity")
}

// ReportAttemptingFullContext implements antlr.ErrorListener.
func (*ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	panic("unimplemented ReportAttemptingFullContext")
}

// ReportContextSensitivity implements antlr.ErrorListener.
func (*ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex int, stopIndex int, prediction int, configs *antlr.ATNConfigSet) {
	panic("unimplemented ReportContextSensitivity")
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

func PreprocessIncludes(ast *AST, sourceListing *SourceListing, reader SourceReader) bool {
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
			incAst, err := CompileAST(string(srcName), sourceListing, reader)
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

func PreprocessAllIncludes(ast *AST, sourceListing *SourceListing, reader SourceReader) {
	for PreprocessIncludes(ast, sourceListing, reader) {
	}
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

func CompileAST(srcName string, sourceListing *SourceListing, reader SourceReader) (*AST, error) {
	src, err := reader.ReadSourceFile(srcName)
	if err != nil {
		return nil, err
	}

	ast := ParseSrc(srcName, src)
	if ast.Errors.HasErrors() {
		return ast, ast.Errors
	}

	if sourceListing != nil {
		sourceListing.AddSource(srcName, src)
	}
	return ast, nil
}
