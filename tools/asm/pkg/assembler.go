package internal

import (
	"fmt"
	"reflect"
	"sort"
)

type ByteCode map[int]byte
type AssemblyContext struct {
	ProgramCounter int
	Labels         map[string]int
	Deffinitions   map[string]any
	PrevByteCode   ByteCode
	ByteCode       ByteCode
	Errors         AssemblerError
	PassCnt        int
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

// type OpcodeAssembler func(pcc int, inst ASTStatement, ignoreRelJmp bool) ([]byte, error)
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
		if strValue, ok := op.String(); ok {
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
			SourceError{
				Type:    SourceErrorTypeAssemblyError,
				SrcName: s.SrcPointer.Name,
				Line:    s.SrcPointer.Line,
				Msg:     err,
			},
		},
	}
}

func Assemble(ast *AST, opcodeAssembler OpcodeAssembler) (ByteCode, *AssemblerError) {
	context := AssemblyContext{
		ProgramCounter: 0,
		Labels:         map[string]int{},
		Deffinitions:   map[string]any{},
		PrevByteCode:   map[int]byte{},
		ByteCode:       map[int]byte{},
		PassCnt:        0,
	}

	for _, s := range ast.Statements {
		if s.Type == ASTStatementTypeLabel {
			context.Labels[s.Name] = 0x00
		}
	}

	for {
		context.ProgramCounter = 0
		context.ByteCode = make(ByteCode)
		context.Deffinitions = map[string]any{}

		for _, s := range ast.Statements {
			if s.Type == ASTStatementTypeLabel {
				context.Labels[s.Name] = context.ProgramCounter
				continue
			} else if s.Type == ASTStatementTypeOrigin {
				p0, ok := s.Operands.Number(0, context)
				if !ok {
					return nil, toAssemblyError(s, "expected exaclty 1 number operand")
				}
				context.ProgramCounter = int(p0)
				continue
			} else if s.Type == ASTStatementTypeDataByte || s.Type == ASTStatementTypeDataWord {
				bs, err := getBytes(s, context)
				if err != nil {
					return nil, toAssemblyError(s, err.Error())
				}
				err = context.ByteCode.SetBytes(context.ProgramCounter, bs)
				if err != nil {
					return nil, toAssemblyError(s, err.Error())
				}
				context.ProgramCounter += len(bs)
				continue
			} else if s.Type == ASTStatementTypeSkipBytes || s.Type == ASTStatementTypeSkipWords {
				n, ok := s.Operands[0].Number(context)
				if !ok {
					return nil, toAssemblyError(s, fmt.Sprintf("expected number, got: '%v'", reflect.TypeOf(s.Operands[0])))
				}
				if s.Type == ASTStatementTypeSkipWords {
					n *= 2
				}
				context.ProgramCounter += int(n)
				continue
			} else if s.Type == ASTStatementTypePrepDefine {
				expr, ok := s.Operands.Expr(0)
				if !ok {
					context.Deffinitions[s.Name] = s.Operands[0].Value
					continue
				}
				v, ok := EvaluateExpr(expr, context)
				if !ok {
					return nil, toAssemblyError(s, "failed to evaluate expression")
				}
				context.Deffinitions[s.Name] = v
				continue
			}

			if s.Type != ASTStatementTypeInstruction {
				return nil, toAssemblyError(s, fmt.Sprintf("unexpected statement: %v", s.Type))
			}

			opBytes, err := opcodeAssembler(s, context, context.PassCnt == 0)
			if err != nil {
				return nil, toAssemblyError(s, err.Error())
			}
			err = context.ByteCode.SetBytes(context.ProgramCounter, opBytes)
			if err != nil {
				return nil, toAssemblyError(s, err.Error())
			}
			context.ProgramCounter += len(opBytes)
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

func AssembleSrc(srcName string, reader SourceReader, opcoOpcodeAssembler OpcodeAssembler) (ByteCode, *AssemblerError, error) {
	ast, err := CompileAST(srcName, reader)
	if err != nil {
		return nil, nil, err
	}

	PreprocessAllIncludes(ast, reader)
	if ast.Errors.HasErrors() {
		return nil, &ast.Errors, nil
	}

	bc, aerr := Assemble(ast, opcoOpcodeAssembler)
	return bc, aerr, err
}
