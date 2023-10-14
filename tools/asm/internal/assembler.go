package internal

import (
	"fmt"
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

type OpcodeAssembler func(inst ASTInstruction) ([]byte, error)

func substituteLabel(inst *ASTInstruction, labels map[string]int) {
	for i, op := range inst.Operands {
		strVal, ok := op.(ASTName)
		if !ok {
			continue
		}
		lblAddr, ok := labels[string(strVal)]
		if ok {
			inst.Operands[i] = ASTNumber(lblAddr)
		}
	}
}

func Assemble(ast *AST, opcodeAssembler OpcodeAssembler) (ByteCode, error) {
	programCounter := 0x00
	byteCode := make(ByteCode)
	labels := map[string]int{}

	for _, s := range ast.Statements {
		lbl, ok := s.(ASTLabel)
		if ok {
			labels[lbl.Name] = programCounter
			continue
		}

		inst, ok := s.(ASTInstruction)
		if !ok {
			panic("elo !") //todo: ??
		}

		substituteLabel(&inst, labels)
		opBytes, err := opcodeAssembler(inst)
		if err != nil {
			panic(err) //todo: ??
		}
		err = byteCode.SetBytes(programCounter, opBytes)
		if err != nil {
			panic(err) //todo: ??
		}
		programCounter += len(opBytes)
	}

	return byteCode, nil
}

func AssembleSrc(srcName string, reader SourceReader, opcoOpcodeAssembler OpcodeAssembler) (ByteCode, error) {
	ast, err := CompileAST(srcName, reader)
	if err != nil {
		return nil, err
	}
	return Assemble(ast, opcoOpcodeAssembler)
}
