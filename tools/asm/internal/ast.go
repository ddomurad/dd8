package internal

import (
	"errors"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ddomurad/dd8/tools/asm/internal/parser"
)

type ASTNumber int64
type ASTRegister string
type ASTName string
type ASTStatement any

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
}

type AST struct {
	Statements []ASTStatement
}

type progVisitor struct {
	parser.BaseDD8ASMVisitor
}

func newProgVisitor() *progVisitor {
	return &progVisitor{}
}

func (v *progVisitor) Visit(tree antlr.ParseTree) interface{} { return tree.Accept(v) }
func (v *progVisitor) VisitChildren(ctx antlr.RuleNode) interface{} {
	children := []any{}
	for _, n := range ctx.GetChildren() {
		children = append(children, v.Visit(n.(antlr.ParseTree)))
	}
	return children
}

func (v *progVisitor) VisitTerminal(_ antlr.TerminalNode) interface{} {
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
		panic("lol !") //todo: ??
	}

	return v.Visit(children[0].(antlr.ParseTree))
}

func (v *progVisitor) VisitInstruction(ctx *parser.InstructionContext) interface{} {
	children := ctx.GetChildren()
	//todo: check for length ?
	opcodeResp := v.Visit(children[0].(antlr.ParseTree))
	opcodeStr, ok := opcodeResp.(ASTName)
	if !ok {
		panic("lol 2") //todo: ??
	}

	inst := ASTInstruction{
		OpCode:   opcodeStr,
		Operands: []any{},
	}

	if len(children) > 2 {
		panic("lol") //todo: ??
	}
	if len(children) == 2 {
		x := v.Visit(children[1].(antlr.ParseTree))
		_ = x //todo: clean up pls ?
		inst.Operands, ok = v.Visit(children[1].(antlr.ParseTree)).([]interface{})
		if !ok {
			panic("lol") //todo: ??
		}
	}

	return inst
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
			panic("lol") //todo: ??
		}
	}

	return args
}

func (v *progVisitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	children := ctx.GetChildren()
	if len(children) != 1 {
		panic("lol") //todo: ???
	}
	cr := v.Visit(children[0].(antlr.ParseTree))
	return cr
	// switch tr := cr.(type) {
	// case string:
	// 	return ASTNameOperand(tr)
	// case int64:
	// 	return ASTNumericOperand(tr)
	// case interface{}:
	// 	return tr
	// default:
	// 	panic("lol") //todo: ??
	// }
}

func (v *progVisitor) VisitNum(ctx *parser.NumContext) interface{} {
	children := ctx.GetChildren()
	if len(children) != 1 {
		panic("lol") //todo: ??
	}
	terminalNode, ok := children[0].(antlr.TerminalNode)
	if !ok {
		panic("lol") //todo: ??
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
		panic(err) //todo: ??
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
		panic("elo 123") //todo: ??
	}

	insType, ok := children[0].(antlr.TerminalNode)
	if !ok {
		panic("elo ppp") //todo: ??
	}

	prepInst := insType.GetText()
	switch prepInst {
	case ".org":
		return v.buildPrepOrigin(children[1:])
	case ".def":
		return v.buildPrepDefine(children[1:])
	}

	panic("lul") //todo: ??
}

func (v *progVisitor) VisitPrep_def_args(ctx *parser.Prep_def_argsContext) interface{} {
	argCtx, ok := ctx.Argument().(*parser.ArgumentContext)
	if !ok {
		panic("lol") //todo: lol
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
			panic("ulala") //todo:: ??
		}
	}

	return defs
}

func (v *progVisitor) buildPrepOrigin(children []antlr.Tree) interface{} {
	if len(children) != 1 {
		panic("wolo wolo") //todo: ??
	}

	vc := v.Visit(children[0].(antlr.ParseTree))
	return ASTOrigin{Address: vc}
}

func (v *progVisitor) buildPrepDefine(children []antlr.Tree) interface{} {
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
			panic("lol 00-") //todo: ??
		}
	}

	return defs
}

func applyDefToOperand(operand any, defs map[string]any) (any, bool) {
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
		nv, ok := applyDefToOperand(op, defs)
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
			nv, ok := applyDefToOperand(tst.Address, defs)
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

func PreprocessDefinitions(ast *AST) error {
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
			return errors.New("00000A ") //todo: ??
		}

		defs[def.Name] = def.Value
	}

	ast.Statements = filtSts
	for applyDefinitions(ast, defs) {
	}

	return nil
}

func ParseSrc(src string) *AST {
	input := antlr.NewInputStream(src)
	lexer := parser.NewDD8ASMLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewDD8ASMParser(stream)
	tree := parser.Prog()

	return newProgVisitor().Visit(tree).(*AST)
}

func CompileAST(srcName string, reader SourceReader) (*AST, error) {
	src, err := reader.ReadSourceFile(srcName)
	if err != nil {
		return nil, err
	}

	return ParseSrc(src), nil
}
