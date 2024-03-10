package asm

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/google/uuid"
)

type ByteCode map[int]byte

type TemplateGenerator func(s ASTStatement, context *AssemblyContext, tmpl Template, args ASTOperands) *AssemblerError

type Template struct {
	Arguments     []ASTName
	Generate      TemplateGenerator
	TemplateBlock ASTBlock
}

type contextDefintion struct {
	ctx   string
	name  string
	value any
}
type PrepDeffinitions struct {
	global  map[string]any
	context []contextDefintion
}

func NewPrepDeffinitions() PrepDeffinitions {
	return PrepDeffinitions{
		global:  map[string]any{},
		context: []contextDefintion{},
	}
}

func (d *PrepDeffinitions) SetGlobal(name string, value any) {
	d.global[name] = value
}

func (d *PrepDeffinitions) Get(name string) (any, bool) {
	for i := len(d.context) - 1; i >= 0; i-- {
		if d.context[i].name == name {
			return d.context[i].value, true
		}
	}

	if v, ok := d.global[name]; ok {
		return v, true
	}

	return nil, false
}

func (d *PrepDeffinitions) PushContext(ctx, name string, value any) {
	d.context = append(d.context, contextDefintion{
		ctx:   ctx,
		name:  name,
		value: value,
	})
}

func (d *PrepDeffinitions) PopContext(ctx string) {
	newContext := make([]contextDefintion, 0, len(d.context))
	for i := 0; i < len(d.context); i++ {
		if d.context[i].ctx != ctx {
			newContext = append(newContext, d.context[i])
		}
	}
	d.context = newContext
}

type AssemblyContext struct {
	ProgramCounter  int
	Labels          map[string]int
	Deffinitions    PrepDeffinitions
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

func AssembleTemplate(s ASTStatement, context *AssemblyContext, tmpl Template, args ASTOperands) *AssemblerError {
	rndCtxName := uuid.New().String()
	if len(tmpl.Arguments) != len(args) {
		return toAssemblyError(ASTStatement{}, fmt.Sprintf("template expected %d arguments, got: %d", len(tmpl.Arguments), len(args)))
	}

	if context.SourceListing != nil {
		context.SourceListing.SetOverrideLine(s.SrcPointer.Name, s.SrcPointer.Line)
	}

	for i, arg := range tmpl.Arguments {
		// todo: this kida shares same code as with the  `else if s.Type == ASTStatementTypePrepDefine`
		expr, ok := args.Expr(i)
		if !ok {
			context.Deffinitions.SetGlobal(string(arg), args[i].Value)
			continue
		}
		v, ok := EvaluateExpr(expr, *context)
		if !ok {
			return toAssemblyError(s, "failed to evaluate expression")
		}
		context.Deffinitions.PushContext(rndCtxName, string(arg), v)
	}

	err := AssempleStatements(context, tmpl.TemplateBlock.Statements)
	context.Deffinitions.PopContext(rndCtxName)

	if context.SourceListing != nil {
		context.SourceListing.ClearOverride()
	}

	return err
}

func AssempleStatements(context *AssemblyContext, statements []ASTStatement) *AssemblerError {
	for _, s := range statements {
		if s.Type == ASTStatementTypeLabel {
			context.Labels[s.Name] = context.ProgramCounter
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
				context.Deffinitions.SetGlobal(s.Name, s.Operands[0].Value)
				continue
			}
			v, ok := EvaluateExpr(expr, *context)
			if !ok {
				return toAssemblyError(s, "failed to evaluate expression")
			}
			context.Deffinitions.SetGlobal(s.Name, v)
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
		Labels:          map[string]int{},
		Deffinitions:    NewPrepDeffinitions(),
		Templates:       map[string]Template{},
		PrevByteCode:    map[int]byte{},
		ByteCode:        map[int]byte{},
		PassCnt:         0,
		OpcodeAssembler: opcodeAssembler,
		SourceListing:   sourceListing,
	}

	for _, s := range ast.Statements {
		if s.Type == ASTStatementTypeLabel {
			context.Labels[s.Name] = 0x00
		}
	}

	for {
		context.ProgramCounter = 0
		context.ByteCode = make(ByteCode)
		context.Deffinitions = NewPrepDeffinitions()
		context.Templates = map[string]Template{}

		if sourceListing != nil {
			sourceListing.ClearCodePass()
		}

		err := AssempleStatements(&context, ast.Statements)
		if err != nil {
			return nil, err
		}

		// for _, s := range ast.Statements {
		// 	if s.Type == ASTStatementTypeLabel {
		// 		context.Labels[s.Name] = context.ProgramCounter
		// 		continue
		// 	} else if s.Type == ASTStatementTypeOrigin {
		// 		p0, ok := s.Operands.Number(0, context)
		// 		if !ok {
		// 			return nil, toAssemblyError(s, "expected exaclty 1 number operand")
		// 		}
		// 		context.ProgramCounter = int(p0)
		// 		continue
		// 	} else if s.Type == ASTStatementTypeDataByte || s.Type == ASTStatementTypeDataWord {
		// 		bs, err := getBytes(s, context)
		// 		if err != nil {
		// 			return nil, toAssemblyError(s, err.Error())
		// 		}
		// 		err = context.ByteCode.SetBytes(context.ProgramCounter, bs)
		// 		if err != nil {
		// 			return nil, toAssemblyError(s, err.Error())
		// 		}
		// 		if sourceListing != nil {
		// 			sourceListing.AddCode(s.SrcPointer.Name, s.SrcPointer.Line, context.ProgramCounter, bs...)
		// 		}
		// 		context.ProgramCounter += len(bs)
		// 		continue
		// 	} else if s.Type == ASTStatementTypeSkipBytes || s.Type == ASTStatementTypeSkipWords {
		// 		n, ok := s.Operands[0].Number(context)
		// 		if !ok {
		// 			return nil, toAssemblyError(s, fmt.Sprintf("expected number, got: '%v'", reflect.TypeOf(s.Operands[0])))
		// 		}
		// 		if s.Type == ASTStatementTypeSkipWords {
		// 			n *= 2
		// 		}
		// 		context.ProgramCounter += int(n)
		// 		continue
		// 	} else if s.Type == ASTStatementTypePrepDefine {
		// 		expr, ok := s.Operands.Expr(0)
		// 		if !ok {
		// 			context.Deffinitions[s.Name] = s.Operands[0].Value
		// 			continue
		// 		}
		// 		v, ok := EvaluateExpr(expr, context)
		// 		if !ok {
		// 			return nil, toAssemblyError(s, "failed to evaluate expression")
		// 		}
		// 		context.Deffinitions[s.Name] = v
		// 		continue
		// 	} else if s.Type == ASTStatementTypePrepTemplateDef {
		// 		argsNames, ok := s.Operands[0].Names()
		// 		if !ok {
		// 			return nil, toAssemblyError(s, fmt.Sprintf("expected a list of names, got: '%v'", reflect.TypeOf(s.Operands[0])))
		// 		}
		// 		tmplBlock, ok := s.Operands[1].Block()
		// 		if !ok {
		// 			return nil, toAssemblyError(s, fmt.Sprintf("expected a block, got: '%v'", reflect.TypeOf(s.Operands[1])))
		// 		}
		// 		context.Templates[s.Name] = Template{
		// 			Arguments:     argsNames,
		// 			Generate:      AssembleTemplate,
		// 			TemplateBlock: tmplBlock,
		// 		}
		// 		continue
		// 	} else if s.Type == ASTStatementTypePrepTemplateUse {
		// 		tmpl, ok := context.Templates[s.Name]
		// 		if !ok {
		// 			return nil, toAssemblyError(s, fmt.Sprintf("template '%s' not found", s.Name))
		// 		}
		// 		tmpl.Generate(tmpl, s.Operands, &context)
		// 		_, _ = tmpl, ok
		// 	}
		//
		// 	if s.Type != ASTStatementTypeInstruction {
		// 		return nil, toAssemblyError(s, fmt.Sprintf("unexpected statement: %v", s.Type))
		// 	}
		//
		// 	opBytes, err := opcodeAssembler(s, context, context.PassCnt == 0)
		// 	if err != nil {
		// 		return nil, toAssemblyError(s, err.Error())
		// 	}
		// 	err = context.ByteCode.SetBytes(context.ProgramCounter, opBytes)
		// 	if err != nil {
		// 		return nil, toAssemblyError(s, err.Error())
		// 	}
		// 	if sourceListing != nil {
		// 		sourceListing.AddCode(s.SrcPointer.Name, s.SrcPointer.Line, context.ProgramCounter, opBytes...)
		// 	}
		//
		// 	context.ProgramCounter += len(opBytes)
		// }

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
