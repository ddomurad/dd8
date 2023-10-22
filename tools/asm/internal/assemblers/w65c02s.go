package assemblers

import (
	"fmt"

	"github.com/ddomurad/dd8/tools/asm/internal"
)

var (
	oneByteInstructions = map[string]byte{
		"brk": 0x00,
		"php": 0x08,
		"clc": 0x18,
		"plp": 0x28,
		"sec": 0x38,
		"rti": 0x40,
		"pha": 0x48,
		"cli": 0x58,
		"phy": 0x5a,
		"rts": 0x60,
		"pla": 0x68,
		"sei": 0x78,
		"ply": 0x7a,
		"dey": 0x88,
		"txa": 0x8a,
		"tya": 0x98,
		"txs": 0x9a,
		"tay": 0xa8,
		"tax": 0xaa,
		"clv": 0xb8,
		"tsx": 0xba,
		"iny": 0xc8,
		"dex": 0xca,
		"wai": 0xcb,
		"cld": 0xd8,
		"phx": 0xda,
		"stp": 0xdb,
		"inx": 0xe8,
		"nop": 0xea,
		"sed": 0xf8,
		"plx": 0xfa,
	}
	immediateInstructions = map[string]byte{
		"orai": 0x09,
		"andi": 0x29,
		"eori": 0x49,
		"adci": 0x69,
		"biti": 0x89,
		"ldyi": 0xa0,
		"ldai": 0xa9,
		"ldxi": 0xa2,
		"cpyi": 0xc0,
		"cpmi": 0xc9,
		"cpxi": 0xe0,
		"sbci": 0xe9,
		"cmpi": 0xc9,
	}
	zpInstructions = map[string]byte{
		"tsb": 0x04,
		"ora": 0x05,
		"asl": 0x06,
		"trb": 0x14,
		"bit": 0x24,
		"and": 0x25,
		"rol": 0x26,
		"eor": 0x45,
		"lsr": 0x46,
		"adc": 0x65,
		"stz": 0x64,
		"ror": 0x66,
		"sty": 0x84,
		"sta": 0x85,
		"stx": 0x86,
		"ldy": 0xa4,
		"lda": 0xa5,
		"ldx": 0xa6,
		"cpy": 0xc4,
		"cpm": 0xc5,
		"dec": 0xc6,
		"sbc": 0xe5,
		"cmp": 0xc5,
		"cpx": 0xe4,
		"inc": 0xe6,
	}
	absoluteInstructions = map[string]byte{
		"ora": 0x0d,
		"tsb": 0x0c,
		"trb": 0x1c,
		"bit": 0x2c,
		"and": 0x2d,
		"asl": 0x0e,
		"jsr": 0x20,
		"rol": 0x2e,
		"eor": 0x4d,
		"lsr": 0x4e,
		"jmp": 0x4c,
		"ror": 0x6e,
		"adc": 0x6d,
		"sty": 0x8c,
		"sta": 0x8d,
		"stx": 0x8e,
		"stz": 0x9c,
		"lda": 0xad,
		"ldy": 0xac,
		"ldx": 0xae,
		"cpy": 0xcc,
		"cpm": 0xcd,
		"dec": 0xce,
		"cpx": 0xec,
		"sbc": 0xed,
		"cmp": 0xcd,
		"inc": 0xee,
	}
	indirectAInstructions = map[string]byte{
		"jmp": 0x6c,
	}
	indirectAXInstructions = map[string]byte{
		"jmp": 0x7c,
	}
	indirectZpXInstructions = map[string]byte{
		"ora": 0x01,
		"and": 0x21,
		"eor": 0x41,
		"adc": 0x61,
		"sta": 0x81,
		"lda": 0xa1,
		"cpm": 0xc1,
		"sbc": 0xe1,
		"cmp": 0xc1,
	}
	indirectZpYInstructions = map[string]byte{
		"ora": 0x11,
		"and": 0x31,
		"eor": 0x51,
		"adc": 0x71,
		"sta": 0x91,
		"lda": 0xb1,
		"cpm": 0xd1,
		"sbc": 0xf1,
		"cmp": 0xd1,
	}
	indirectZpInstructions = map[string]byte{
		"ora": 0x12,
		"and": 0x32,
		"eor": 0x52,
		"adc": 0x72,
		"sta": 0x92,
		"lda": 0xb2,
		"cpm": 0xd2,
		"sbc": 0xf2,
		"cmp": 0xd2,
	}
	indexZpXInstructions = map[string]byte{
		"ora": 0x15,
		"asl": 0x16,
		"bit": 0x34,
		"and": 0x35,
		"rol": 0x36,
		"eor": 0x55,
		"lsr": 0x56,
		"stz": 0x74,
		"adc": 0x75,
		"ror": 0x76,
		"sty": 0x94,
		"sta": 0x95,
		"ldy": 0xb4,
		"lda": 0xb5,
		"cpm": 0xd5,
		"dec": 0xd6,
		"sbc": 0xf5,
		"cmp": 0xd5,
		"inc": 0xf6,
	}
	indexZpYInstructions = map[string]byte{
		"stx": 0x96,
		"ldx": 0xb6,
	}
	indexAXInstructions = map[string]byte{
		"ora": 0x1d,
		"asl": 0x1e,
		"bit": 0x3c,
		"and": 0x3d,
		"rol": 0x3e,
		"eor": 0x5d,
		"lsr": 0x5e,
		"adc": 0x7d,
		"ror": 0x7e,
		"sta": 0x9d,
		"stz": 0x9e,
		"ldy": 0xbc,
		"lda": 0xbd,
		"cpm": 0xdd,
		"dec": 0xde,
		"sbc": 0xfd,
		"cmp": 0xdd,
		"inc": 0xfe,
	}
	indexAYInstructions = map[string]byte{
		"ora": 0x19,
		"and": 0x39,
		"eor": 0x59,
		"adc": 0x79,
		"cpm": 0xd9,
		"sbc": 0xf9,
		"cmp": 0xd9,
		"sta": 0x99,
		"lda": 0xb9,
		"ldx": 0xbe,
	}
	acumulatorInstructions = map[string]byte{
		"asl": 0x0a,
		"inc": 0x1a,
		"rol": 0x2a,
		"dec": 0x3a,
		"lsr": 0x4a,
		"ror": 0x6a,
	}
	relativeBitInstructions = map[string]byte{
		"bbr": 0x0f,
		"bbs": 0x8f,
	}
	zpBitInstructions = map[string]byte{
		"rmb": 0x07,
		"smb": 0x87,
	}
	relativeIntructions = map[string]byte{
		"bpl": 0x10,
		"bmi": 0x30,
		"bvc": 0x50,
		"bra": 0x80,
		"bcc": 0x90,
		"bcs": 0xb0,
		"bne": 0xd0,
		"beq": 0xf0,
		"bvs": 0x70,
	}
)

func prepInstruction(inst internal.ASTInstruction) (indirect bool, opcode string, arg1u *uint16, arg1r *byte, arg2u *uint16, arg2r *byte, opcnt int, err error) {
	indirect = false
	opcode = string(inst.OpCode)

	arg1u = nil
	arg1r = nil
	arg2u = nil
	arg2r = nil
	opcnt = 0
	err = nil

	operands := inst.Operands
	opLen := len(operands)

	pv, ok := internal.TryGetTypedOperand[internal.ASTParentheses](operands, 0)
	if ok {
		if opLen == 1 {
			operands = pv
			opLen = len(operands)
			indirect = true
		} else {
			if len(pv) != 1 {
				err = fmt.Errorf("expected only one operand in parentheses")
				return
			}
			operands[0] = pv[0]
			indirect = true
		}
	}
	if opLen >= 1 {
		opcnt = 1
		v1r, ok := internal.TryGetTypedOperand[internal.ASTRegister](operands, 0)
		if ok {
			v1rb := v1r[0]
			arg1r = &v1rb
		} else {
			v1, ok := internal.TryGetTypedOperand[internal.ASTNumber](operands, 0)
			if !ok {
				err = fmt.Errorf("first operand expected to be a number, got instead: '%v'", operands[0])
				return
			}
			if v1 < 0 || v1 > 0xffff {
				err = fmt.Errorf("numeric operand expected to be 16bits, got instead: %x", v1)
				return
			}
			v1u16 := uint16(v1)
			arg1u = &v1u16
		}
	}

	if opLen == 2 {
		opcnt = 2
		v2n, ok := internal.TryGetTypedOperand[internal.ASTNumber](operands, 1)
		if ok {
			v2u16 := uint16(v2n)
			arg2u = &v2u16
			return
		}
		v2, ok := internal.TryGetTypedOperand[internal.ASTRegister](operands, 1)
		if !ok {
			err = fmt.Errorf("first operand expected to be a register, got instead: '%v'", operands[1])
			return
		}
		if v2 != "x" && v2 != "y" {
			err = fmt.Errorf("second operand expected to be one letter register, got instead: '%v'", v2)
			return
		}
		v2b := v2[0]
		arg2r = &v2b

	}

	return
}

func ens4bit(num uint16) (uint8, error) {
	if num > 0x0f || num < 0 {
		return 0, fmt.Errorf("expected 4bit value, got: %x (hex)", num)
	}

	return uint8(num), nil
}

func ens8bit(num uint16) (uint8, error) {
	if num > 0xff || num < 0 {
		return 0, fmt.Errorf("expected 8bit value, got: %x (hex)", num)
	}

	return uint8(num), nil
}

func cmpPtr[T comparable](p *T, v T) bool {
	if p == nil {
		return false
	}

	return *p == v
}

func is4bit(num *uint16) bool {
	return num != nil && *num <= 0x0f
}

func is8bit(num *uint16) bool {
	return num != nil && *num <= 0xff
}

func OpcodeAssemblerW65C02S(pc int, inst internal.ASTInstruction) ([]byte, error) {
	indirect, opcode, arg1, arg1r, arg2, arg2r, opcnt, err := prepInstruction(inst)
	if err != nil {
		return nil, err
	}

	_, _, _, _, _, _ = indirect, opcode, arg1, arg1r, arg2, arg2r

	ibc, isImmediate := immediateInstructions[opcode]
	if isImmediate {
		if is8bit(arg1) && opcnt == 1 {
			return []byte{ibc, byte(*arg1)}, nil
		}
	} else if opcnt == 0 { //one byte inst
		bc, ok := oneByteInstructions[opcode]
		if ok {
			return []byte{bc}, nil
		}
	} else if !indirect && opcnt == 1 && arg1 != nil { // a
		bc, ok := relativeIntructions[opcode]
		if ok {
			rj := (int(*arg1) - pc) - 0x02
			if rj > 127 || rj < -128 {
				return nil, fmt.Errorf("relative jump to large at %s", inst.SrcPointer.String())
			}
			return []byte{bc, byte(rj)}, nil
		}

		if is8bit(arg1) {
			bc, ok := zpInstructions[opcode]
			if ok {
				return []byte{bc, byte(*arg1)}, nil
			}
		}

		bc, ok = absoluteInstructions[opcode]
		if ok {
			return []byte{bc, byte(*arg1), byte(*arg1 >> 8)}, nil
		}
	} else if indirect && arg1 != nil && cmpPtr(arg2r, 'x') { // (zp, x)
		if is8bit(arg1) {
			bc, ok := indirectZpXInstructions[opcode]
			if ok {
				return []byte{bc, byte(*arg1)}, nil
			}
		}
		bc, ok := indirectAXInstructions[opcode]
		if ok {
			return []byte{bc, byte(*arg1), byte(*arg1 >> 8)}, nil
		}
	} else if indirect && is8bit(arg1) && cmpPtr(arg2r, 'y') { // (zp, x)
		bc, ok := indirectZpYInstructions[opcode]
		if ok {
			return []byte{bc, byte(*arg1)}, nil
		}
	} else if indirect && arg1 != nil && opcnt == 1 { // (zp, x)
		if is8bit(arg1) {
			bc, ok := indirectZpInstructions[opcode]
			if ok {
				return []byte{bc, byte(*arg1)}, nil
			}
		}
		bc, ok := indirectAInstructions[opcode]
		if ok {
			return []byte{bc, byte(*arg1), byte(*arg1 >> 8)}, nil
		}
	} else if !indirect && is8bit(arg1) && cmpPtr(arg2r, 'x') { // (zp, x)
		bc, ok := indexZpXInstructions[opcode]
		if ok {
			return []byte{bc, byte(*arg1)}, nil
		}
	} else if !indirect && arg1 != nil && cmpPtr(arg2r, 'y') { // (zp, x)
		if is8bit(arg1) {
			bc, ok := indexZpYInstructions[opcode]
			if ok {
				return []byte{bc, byte(*arg1)}, nil
			}
		}
		bc, ok := indexAYInstructions[opcode]
		if ok {
			return []byte{bc, byte(*arg1), byte(*arg1 >> 8)}, nil
		}
	} else if !indirect && arg1 != nil && cmpPtr(arg2r, 'x') { // (zp, x)
		bc, ok := indexAXInstructions[opcode]
		if ok {
			return []byte{bc, byte(*arg1), byte(*arg1 >> 8)}, nil
		}
	} else if !indirect && cmpPtr(arg1r, 'a') && opcnt == 1 { // reg a
		bc, ok := acumulatorInstructions[opcode]
		if ok {
			return []byte{bc}, nil
		}
	} else if !indirect && arg1 != nil && is4bit(arg2) { // rmb, smb
		if is8bit(arg1) {
			bc, ok := zpBitInstructions[opcode]
			if ok {
				return []byte{bc + byte(*arg2)<<4, byte(*arg1)}, nil
			}
		}

		bc, ok := relativeBitInstructions[opcode]
		if ok {

			rj := (int(*arg1) - pc) - 0x02
			if rj > 127 || rj < -128 {
				return nil, fmt.Errorf("relative jump to large at %s", inst.SrcPointer.String())
			}
			return []byte{bc + byte(*arg2)<<4, byte(rj)}, nil
			// return []byte{bc + byte(*arg1)<<4}, nil
		}
	}

	return nil, fmt.Errorf("unsupported instruction at %s", inst.SrcPointer.String())
}
