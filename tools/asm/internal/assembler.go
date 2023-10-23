package internal

import (
	"fmt"
	"reflect"
	"sort"
)

type ByteCode map[int]byte

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

type OpcodeAssembler func(pcc int, inst ASTStatement, ignoreRelJmp bool) ([]byte, error)

func substituteLabel(ast *AST, inst ASTStatement, labels map[string]int) ASTStatement {
	for i, op := range inst.Operands {
		strVal, ok := op.Name()
		if !ok {
			continue
		}
		lblAddr, ok := labels[string(strVal)]
		if ok {
			inst.Operands[i].Value = ASTNumber(lblAddr)
		} else if ast.HasLabel(string(strVal)) {
			inst.Operands[i].Value = ASTNumber(0x00) //note: temp. replace future label with 0x00
		}
	}
	return inst
}

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

func Ens16bit(operands []ASTOperand, index int) (uint16, error) {
	if len(operands) <= index {
		return 0, fmt.Errorf("expected at least %d operands", index)
	}

	v, ok := operands[0].Number()
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

func Assemble(oast *AST, opcodeAssembler OpcodeAssembler) (ByteCode, error) {
	labels := map[string]int{}
	prevByteCode := make(ByteCode)
	passCnt := 0

	for {
		programCounter := 0x00
		byteCode := make(ByteCode)
		astc := oast.Copy()
		for _, s := range astc.Statements {
			if s.Type == ASTStatementTypeLabel {
				labels[s.Name] = programCounter
				continue
			}

			if s.Type == ASTStatementTypeOrigin {
				p0, ok := s.Operands.Number(0)
				if !ok {
					return byteCode, fmt.Errorf("expected exaclty 1 operand")
				}
				programCounter = int(p0)
				continue
			}

			if s.Type != ASTStatementTypeInstruction {
				return nil, fmt.Errorf("unexpected statement: %v", s.Type)
			}

			s = substituteLabel(astc, s, labels)
			opBytes, err := opcodeAssembler(programCounter, s, passCnt == 0)
			if err != nil {
				return nil, err
			}
			err = byteCode.SetBytes(programCounter, opBytes)
			if err != nil {
				return nil, err
			}
			programCounter += len(opBytes)
		}

		if reflect.DeepEqual(prevByteCode, byteCode) {
			return byteCode, nil
		}

		prevByteCode = byteCode
		passCnt++
	}
}

func AssembleSrc(srcName string, reader SourceReader, opcoOpcodeAssembler OpcodeAssembler) (ByteCode, error) {
	ast, err := CompileAST(srcName, reader)
	if err != nil {
		return nil, err
	}

	PreprocessAST(ast, reader)
	if ast.Errors.HasErrors() {
		return nil, ast.Errors
	}

	return Assemble(ast, opcoOpcodeAssembler)
}
