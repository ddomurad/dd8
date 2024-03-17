package asm

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/google/uuid"
)

type ByteCode map[int]byte

type TemplateGenerator func(s ASTStatement, context *AssemblyContext, tmpl Template, args ASTOperands) *AssemblerError

type Template struct {
	Name          string
	Arguments     []ASTName
	Generate      TemplateGenerator
	TemplateBlock ASTBlock
}

type AssemblyContext struct {
	ProgramCounter  int
	Labels          ContextMap[int]
	Deffinitions    ContextMap[any]
	Templates       map[string]Template
	PrevByteCode    ByteCode
	ByteCode        ByteCode
	PassCnt         int
	OpcodeAssembler OpcodeAssembler
	SourceListing   *SourceListing
}

func (c ByteCode) SetBytes(sa int, byteCode []byte) error {
	for i := 0; i < len(byteCode); i++ {
		addr := sa + i
		if _, ok := c[addr]; ok {
			return fmt.Errorf("byte at %v already occupied", addr)
		}

		c[addr] = byteCode[i]
	}

	return nil
}

func (c ByteCode) ToByteArray(startAddr, endAddr int) []byte {
	out := make([]byte, endAddr-startAddr)
	for i := startAddr; i < endAddr; i++ {
		if b, ok := c[i]; ok {
			out[i] = b
		} else {
			out[i] = 0x00
		}
	}
	return out
}

func (c ByteCode) GetAddresses() []int {
	addr := make([]int, 0, len(c))
	for a := range c {
		addr = append(addr, a)
	}

	sort.Ints(addr)

	return addr
}

type OpcodeAssembler func(inst ASTStatement, actx AssemblyContext, ignoreRelJmp bool) ([]byte, error)

func EnsRegister(operands []any, index int) (string, error) {
	if len(operands) <= index {
		return "", fmt.Errorf("missing opperand: ind: %d, len: %d", index, len(operands))
	}

	sv, ok := operands[index].(ASTRegister)
	if !ok {
		return "", fmt.Errorf("expected a register oprand, go: '%s'", reflect.TypeOf(operands[index]).String())
	}

	return string(sv), nil
}

func Ens16bit(operands []ASTOperand, index int, actx AssemblyContext) (uint16, error) {
	if len(operands) <= index {
		return 0, fmt.Errorf("expected at least %d operands", index)
	}

	v, ok := operands[0].Number(actx)
	if !ok || v > 0xffff || v < 0 {
		return 0, fmt.Errorf("expected 16bit value, got: %x (hex)", v)
	}

	return uint16(v), nil
}

func TryGetTypedOperand[T any](operands []any, index int) (T, bool) {
	var def T
	if len(operands) <= index {
		return def, false
	}

	v, ok := operands[index].(T)
	return v, ok
}

func Ens8bit(num ASTNumber) (uint8, error) {
	if num > 0xff || num < 0 {
		return 0, fmt.Errorf("expected 8bit value, got: %x (hex)", num)
	}

	return uint8(num), nil
}

func getBytes(st ASTStatement, actx AssemblyContext) ([]byte, error) {
	outBytes := make([]byte, 0)

	for _, op := range st.Operands {
		if numValue, ok := op.Number(actx); ok {
			if st.Type == ASTStatementTypeDataWord {
				if numValue < 0x00 || numValue > 0xffff {
					return nil, fmt.Errorf("expected 16bit value got: '%v'", numValue)
				}
				outBytes = append(outBytes, byte(numValue), byte(numValue>>8))
			} else {
				if numValue < 0x00 || numValue > 0xff {
					return nil, fmt.Errorf("expected 8bit value got: '%v'", numValue)
				}
				outBytes = append(outBytes, byte(numValue))
			}
			continue
		}
		if strValue, ok := op.String(actx); ok {
			for _, s := range []byte(strValue) {
				outBytes = append(outBytes, byte(s))
			}
			if st.Type == ASTStatementTypeDataWord && len(strValue)%2 != 0 {
				outBytes = append(outBytes, 0x00)
			}
			continue
		}

		return nil, fmt.Errorf("expected number or string, got: '%v'", reflect.TypeOf(op.Value))
	}

	return outBytes, nil
}

func toAssemblyError(s ASTStatement, err string) *AssemblerError {
	return &AssemblerError{
		Errors: []SourceError{
			{
				Type:    SourceErrorTypeAssemblyError,
				SrcName: s.SrcPointer.Name,
				Line:    s.SrcPointer.Line,
				Msg:     err,
			},
		},
	}
}

// todo: move to a better place and refactor
func sprintOperands(actx AssemblyContext, ops ASTOperands, names ...ASTName) string {
	out := ""
	for i, op := range ops {
		if i > 0 {
			out += ", "
		}
		if names != nil && i < len(names) {
			out += fmt.Sprintf("%s=", (names)[i])
		}

		if v, ok := op.Number(actx); ok {
			out += fmt.Sprintf("0x%x", v)
		} else if v, ok := op.String(actx); ok {
			vs := strings.ReplaceAll(string(v), "\n", "\\n")
			out += fmt.Sprintf("\"%v\"", vs)
		} else if v, ok := op.Register(actx); ok {
			out += fmt.Sprintf("%v", v)
		} else {
			out += fmt.Sprintf("%v", v)
		}
	}
	return out
}

func AssembleTemplate(s ASTStatement, context *AssemblyContext, tmpl Template, args ASTOperands) *AssemblerError {
	rndCtxName := uuid.New().String()
	if len(tmpl.Arguments) != len(args) {
		return toAssemblyError(ASTStatement{}, fmt.Sprintf("template expected %d arguments, got: %d", len(tmpl.Arguments), len(args)))
	}

	srcListingOverriden := false
	if context.SourceListing != nil {
		srcListingOverriden = context.SourceListing.OverrideLine(s.SrcPointer.Name, s.SrcPointer.Line)
	}

	context.Deffinitions.PushContext(rndCtxName)
	context.Labels.PushContext(rndCtxName)

	for i, arg := range tmpl.Arguments {
		expr, ok := args.Expr(i)
		if !ok {
			context.Deffinitions.Set(string(arg), args[i].Value)
			continue
		}
		v, ok := EvaluateExpr(expr, *context)
		if !ok {
			return toAssemblyError(s, "failed to evaluate expression")
		}
		context.Deffinitions.Set(string(arg), v)
	}

	if context.SourceListing != nil {
		context.SourceListing.InsertVirtualLine(
			s.SrcPointer.Name,
			s.SrcPointer.Line,
			fmt.Sprintf("@%s(%s)", s.Name, sprintOperands(*context, args, tmpl.Arguments...)),
		)
	}
	err := AssempleStatements(context, tmpl.TemplateBlock.Statements)
	context.Deffinitions.PopContext(rndCtxName)
	context.Labels.PopContext(rndCtxName)

	if srcListingOverriden {
		context.SourceListing.ClearOverride()
	}

	return err
}

func AssembleRepeat(s ASTStatement, context *AssemblyContext, args ASTOperands) *AssemblerError {
	if len(args) != 4 {
		return toAssemblyError(s, "repeate expected 4 arguments")
	}

	vname, ok := args[0].Name()
	if !ok {
		return toAssemblyError(s, "repeate expected a name as first argument")
	}

	vstart, ok := args[1].Number(*context)
	if !ok {
		return toAssemblyError(s, "repeate expected a number as second argument")
	}

	vend, ok := args[2].Number(*context)
	if !ok {
		return toAssemblyError(s, "repeate expected a number as third argument")
	}

	vblock, ok := args[3].Block()
	if !ok {
		return toAssemblyError(s, "repeate expected a code block as fourth argument")
	}

	srcListingOverriden := false
	if context.SourceListing != nil {
		srcListingOverriden = context.SourceListing.OverrideLine(s.SrcPointer.Name, s.SrcPointer.Line)
	}

	if context.SourceListing != nil {
		context.SourceListing.InsertVirtualLine(
			s.SrcPointer.Name,
			s.SrcPointer.Line,
			fmt.Sprintf("@repeate %s=%d to %d", vname, vstart, vend))
	}

	i := vstart
	for {
		rndCtxName := uuid.New().String()
		context.Deffinitions.PushContext(rndCtxName)
		context.Labels.PushContext(rndCtxName)

		context.Deffinitions.Set(string(vname), ASTNumber(i))

		if context.SourceListing != nil {
			context.SourceListing.InsertVirtualLine(
				s.SrcPointer.Name,
				s.SrcPointer.Line,
				fmt.Sprintf("@repeate %s=%d", vname, i))
		}

		err := AssempleStatements(context, vblock.Statements)

		context.Deffinitions.PopContext(rndCtxName)
		context.Labels.PopContext(rndCtxName)

		if err != nil {
			return err
		}

		if vstart > vend {
			i--
			if i < vend {
				break
			}
		} else {
			i++
			if i > vend {
				break
			}
		}
	}

	if srcListingOverriden {
		context.SourceListing.ClearOverride()
	}

	return nil
}

func AssempleStatements(context *AssemblyContext, statements []ASTStatement) *AssemblerError {
	for _, s := range statements {
		if s.Type == ASTStatementTypeLabel {
			context.Labels.Set(s.Name, context.ProgramCounter)
			continue
		} else if s.Type == ASTStatementTypeOrigin {
			p0, ok := s.Operands.Number(0, *context)
			if !ok {
				return toAssemblyError(s, "expected exaclty 1 number operand")
			}
			context.ProgramCounter = int(p0)
			continue
		} else if s.Type == ASTStatementTypeDataByte || s.Type == ASTStatementTypeDataWord {
			bs, err := getBytes(s, *context)
			if err != nil {
				return toAssemblyError(s, err.Error())
			}
			err = context.ByteCode.SetBytes(context.ProgramCounter, bs)
			if err != nil {
				return toAssemblyError(s, err.Error())
			}
			if context.SourceListing != nil {
				context.SourceListing.AddCode(s.SrcPointer.Name, s.SrcPointer.Line, context.ProgramCounter, bs...)
			}
			context.ProgramCounter += len(bs)
			continue
		} else if s.Type == ASTStatementTypeSkipBytes || s.Type == ASTStatementTypeSkipWords {
			n, ok := s.Operands[0].Number(*context)
			if !ok {
				return toAssemblyError(s, fmt.Sprintf("expected number, got: '%v'", reflect.TypeOf(s.Operands[0])))
			}
			if s.Type == ASTStatementTypeSkipWords {
				n *= 2
			}
			context.ProgramCounter += int(n)
			continue
		} else if s.Type == ASTStatementTypePrepDefine {
			expr, ok := s.Operands.Expr(0)
			if !ok {
				context.Deffinitions.Set(s.Name, s.Operands[0].Value)
				continue
			}
			v, ok := EvaluateExpr(expr, *context)
			if !ok {
				return toAssemblyError(s, "failed to evaluate expression")
			}
			context.Deffinitions.Set(s.Name, v)
			continue
		} else if s.Type == ASTStatementTypePrepTemplateDef {
			argsNames, ok := s.Operands[0].Names()
			if !ok {
				return toAssemblyError(s, fmt.Sprintf("expected a list of names, got: '%v'", reflect.TypeOf(s.Operands[0])))
			}
			tmplBlock, ok := s.Operands[1].Block()
			if !ok {
				return toAssemblyError(s, fmt.Sprintf("expected a block, got: '%v'", reflect.TypeOf(s.Operands[1])))
			}
			context.Templates[s.Name] = Template{
				Name:          s.Name,
				Arguments:     argsNames,
				Generate:      AssembleTemplate,
				TemplateBlock: tmplBlock,
			}
			continue
		} else if s.Type == ASTStatementTypePrepTemplateUse {
			tmpl, ok := context.Templates[s.Name]
			if !ok {
				return toAssemblyError(s, fmt.Sprintf("template '%s' not found", s.Name))
			}
			err := tmpl.Generate(s, context, tmpl, s.Operands)
			if err != nil {
				return err
			}
			continue
		} else if s.Type == ASTStatementTypePrepRepeate {
			err := AssembleRepeat(s, context, s.Operands)
			if err != nil {
				return err
			}
			continue
		}

		if s.Type != ASTStatementTypeInstruction {
			return toAssemblyError(s, fmt.Sprintf("unexpected statement: %v", s.Type))
		}

		opBytes, err := context.OpcodeAssembler(s, *context, context.PassCnt == 0)
		if err != nil {
			return toAssemblyError(s, err.Error())
		}
		err = context.ByteCode.SetBytes(context.ProgramCounter, opBytes)
		if err != nil {
			return toAssemblyError(s, err.Error())
		}
		if context.SourceListing != nil {
			context.SourceListing.AddCode(s.SrcPointer.Name, s.SrcPointer.Line, context.ProgramCounter, opBytes...)
		}

		context.ProgramCounter += len(opBytes)
	}

	return nil
}

func Assemble(ast *AST, opcodeAssembler OpcodeAssembler, sourceListing *SourceListing) (ByteCode, *AssemblerError) {
	context := AssemblyContext{
		ProgramCounter:  0,
		Labels:          NewContextMap[int](),
		Deffinitions:    NewContextMap[any](),
		Templates:       map[string]Template{},
		PrevByteCode:    map[int]byte{},
		ByteCode:        map[int]byte{},
		PassCnt:         0,
		OpcodeAssembler: opcodeAssembler,
		SourceListing:   sourceListing,
	}

	for _, s := range ast.Statements {
		if s.Type == ASTStatementTypeLabel {
			context.Labels.Set(s.Name, 0x00)
		}
	}

	for {
		context.ProgramCounter = 0
		context.ByteCode = make(ByteCode)
		context.Deffinitions = NewContextMap[any]()
		context.Templates = map[string]Template{}

		if sourceListing != nil {
			sourceListing.ClearCodePass()
		}

		err := AssempleStatements(&context, ast.Statements)
		if err != nil {
			return nil, err
		}

		if reflect.DeepEqual(context.PrevByteCode, context.ByteCode) {
			return context.ByteCode, nil
		}

		context.PrevByteCode = context.ByteCode
		context.PassCnt++

		if context.PassCnt > 5 {
			panic("to many assembly passes") //todo: make this a popper error
		}
	}
}

func AssembleSrc(srcName string, reader SourceReader, opcoOpcodeAssembler OpcodeAssembler, createListing bool) (ByteCode, *SourceListing, *AssemblerError, error) {
	var sourceListing *SourceListing
	if createListing {
		sourceListing = NewSourceListing()
	}

	ast, err := CompileAST(srcName, sourceListing, reader)
	if err != nil {
		return nil, nil, nil, err
	}

	PreprocessAllIncludes(ast, sourceListing, reader)
	if ast.Errors.HasErrors() {
		return nil, nil, &ast.Errors, nil
	}

	bc, aerr := Assemble(ast, opcoOpcodeAssembler, sourceListing)
	return bc, sourceListing, aerr, err
}
