package internal

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ddomurad/dd8/tools/asm/internal/parser"
)

type ASTNumericOperand int64
type ASTStringOperand string
type ASTStatement any

type ASTPreprocessInstruction struct {
	Instruction string
	Parameters  []any
}

type ASTInstruction struct {
	OpCode   string
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

		ast.Statements = append(ast.Statements, result)
	}

	return &ast
}

func (v *progVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	children := ctx.GetChildren()
	// v.VisitChildren(ctx)
	if len(children) != 1 && len(children) != 2 {
		panic("lol !") //todo: ??
	}

	return v.Visit(children[0].(antlr.ParseTree))
}

func (v *progVisitor) VisitPrep_instruction(ctx *parser.Prep_instructionContext) interface{} {
	children := ctx.GetChildren()
	//todo: check for length ?
	opcodeResp := v.Visit(children[0].(antlr.ParseTree))
	instStr, ok := opcodeResp.(string)
	if !ok {
		panic("lol 2") //todo: ??
	}

	inst := ASTPreprocessInstruction{
		Instruction: instStr,
		Parameters:  []any{},
	}

	if len(children) > 2 {
		panic("lol") //todo: ??
	}
	if len(children) == 2 {
		x := v.Visit(children[1].(antlr.ParseTree))
		_ = x
		inst.Parameters, ok = v.Visit(children[1].(antlr.ParseTree)).([]interface{})
		if !ok {
			panic("lol") //todo: ??
		}
	}

	return inst
}

func (v *progVisitor) VisitInstruction(ctx *parser.InstructionContext) interface{} {
	children := ctx.GetChildren()
	//todo: check for length ?
	opcodeResp := v.Visit(children[0].(antlr.ParseTree))
	opcodeStr, ok := opcodeResp.(string)
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
		_ = x
		inst.Operands, ok = v.Visit(children[1].(antlr.ParseTree)).([]interface{})
		if !ok {
			panic("lol") //todo: ??
		}
	}

	return inst
}

func (v *progVisitor) VisitPrep_arglist(ctx *parser.Prep_arglistContext) interface{} {
	childrenRes := v.VisitChildren(ctx).([]interface{})
	args := []interface{}{}

	for _, cr := range childrenRes {
		if cr == nil {
			continue
		}
		switch tr := cr.(type) {
		case []interface{}:
			args = append(args, tr...)
		case string:
			args = append(args, ASTStringOperand(tr))
		case int64:
			args = append(args, ASTNumericOperand(tr))
		case interface{}:
			args = append(args, tr)
		default:
			panic("lol") //todo: ??
		}
	}

	return args
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
		case string:
			args = append(args, ASTStringOperand(tr))
		case int64:
			args = append(args, ASTNumericOperand(tr))
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
	return v.Visit(children[0].(antlr.ParseTree))
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
	return num
}

func (v *progVisitor) VisitName(ctx *parser.NameContext) interface{} {
	return ctx.NAME().GetText()
}

func (v *progVisitor) VisitPrep_name(ctx *parser.Prep_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

// func (v *progVisitor) visitOperands(ctx antlr.ParseTree) []interface{} {
// 	children := ctx.GetChildren()
// 	if len(children) != 1 {
// 		panic("lol") //todo: ??
// 	}
//
// 	return v.Visit(children[0].(antlr.ParseTree))
// }

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
