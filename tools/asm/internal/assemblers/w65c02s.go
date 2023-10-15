package assemblers

import (
	"fmt"

	"github.com/ddomurad/dd8/tools/asm/internal"
)

var (
	variatingOpcodes = map[string]byte{
		"ora": 0x01,
		"lda": 0xa1,
	}
	immediateOpcodes = map[string]byte{
		"orai": 0x09,
		"ldai": 0xa9,
	}
	simpleAbsolureOpcodes = map[string]byte{
		"jsr": 0x20,
	}
	simpleOpcodes = map[string]byte{
		"brk": 0x00,
		"bpl": 0x10,
		"nop": 0xea,
		"php": 0x08,
	}
)

func prepInstruction(inst internal.ASTInstruction) (indirect bool, opcode string, arg1 *uint16, arg2 *byte, err error) {
	indirect = false
	opcode = string(inst.OpCode)

	arg1 = nil
	arg2 = nil
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
		arg1 = &v1u16
	}

	if opLen == 2 {
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
		arg2 = &v2b
	}

	return
}

func variationOdas(indirect bool, opcode string, arg1 *uint16, arg2 *byte) (bcode []byte, err error) {
	bcode = nil
	err = nil

	baseByte, ok := variatingOpcodes[opcode]
	if !ok {
		return
	}

	if arg1 == nil {
		err = fmt.Errorf("opcode '%s' requires a numeric operand", opcode)
		return
	}

	if !indirect {
		if arg2 == nil { //zp or a
			if *arg1 <= 0xff { //zp
				baseByte += 0x04
				bcode = []byte{baseByte, byte(*arg1)}
				return
			}
			// a
			baseByte += 0x0c
			bcode = []byte{baseByte, byte(*arg1), byte((*arg1) >> 8)}
			return
		}

		//zp,x or a,x or a,y
		if *arg2 == 'x' {
			if *arg1 <= 0xff { //zp, x
				baseByte += 0x14
				bcode = []byte{baseByte, byte(*arg1)}
				return
			}
			// zp, a
			baseByte += 0x1c
			bcode = []byte{baseByte, byte(*arg1), byte((*arg1) >> 8)}
			return
		}
		if *arg2 == 'y' {
			baseByte += 0x18
			bcode = []byte{baseByte, byte(*arg1), byte((*arg1) >> 8)}
			return
		}
	} else {
		u8, err8 := ens8bit(*arg1)
		if err != nil {
			err = err8
			return
		}
		if arg2 != nil {
			if *arg2 == 'x' {
				baseByte += 0x00
				bcode = []byte{baseByte, byte(u8)}
				return
			}
			// y
			baseByte += 0x10
			bcode = []byte{baseByte, byte(u8)}
			return
		}

		baseByte += 0x11
		bcode = []byte{baseByte, byte(u8)}
		return
	}

	return
}

func ens8bit(num uint16) (uint8, error) {
	if num > 0xff || num < 0 {
		return 0, fmt.Errorf("expected 8bit value, got: %x (hex)", num)
	}

	return uint8(num), nil
}

func OpcodeAssemblerW65C02S(inst internal.ASTInstruction) ([]byte, error) {
	indirect, opcode, arg1, arg2, err := prepInstruction(inst)
	_ = indirect
	_ = opcode
	_ = arg1
	_ = arg2
	_ = err

	sb, ok := simpleOpcodes[opcode]
	if ok {
		if arg1 != nil || arg2 != nil {
			return nil, fmt.Errorf("opcode '%s' does not require operand", opcode)
		}
		return []byte{sb}, nil
	}

	ib, ok := immediateOpcodes[opcode]
	if ok {
		if arg1 == nil {
			return nil, fmt.Errorf("opcode '%s' requires a numeric operand", opcode)
		}
		if arg2 != nil {
			return nil, fmt.Errorf("opcode '%s' requires only a numeric operand", opcode)
		}
		u8, err := ens8bit(*arg1)
		if err != nil {
			return nil, err
		}
		return []byte{ib, u8}, nil
	}

	sa, ok := simpleAbsolureOpcodes[opcode]
	if ok {
		if arg1 == nil {
			return nil, fmt.Errorf("opcode '%s' requires a numeric operand", opcode)
		}
		if arg2 != nil {
			return nil, fmt.Errorf("opcode '%s' requires only one operand", opcode)
		}
		return []byte{sa, byte(*arg1), byte((*arg1) >> 8)}, nil
	}

	bcode, err := variationOdas(indirect, opcode, arg1, arg2)
	if err != nil {
		return nil, err
	}
	if bcode != nil {
		return bcode, nil
	}

	return nil, fmt.Errorf("unsupported instruction")
}

// func __OpcodeAssemblerW65C02S(inst internal.ASTInstruction) ([]byte, error) {
// 	switch inst.OpCode {
// 	case "brk":
// 		return []byte{0x00}, nil
// 	case "bpl":
// 		return []byte{0x10}, nil
// 	case "nop":
// 		return []byte{0xea}, nil
// 	case "php":
// 		return []byte{0x08}, nil
// 	case "jsr":
// 		p0, err := internal.Ens16bit(inst.Operands, 0)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return []byte{0x20, byte(p0), byte(p0 >> 8)}, nil
// 	case "ldai":
// 		p0, err := internal.Ens8bit(inst.Operands, 0)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return []byte{0xa9, byte(p0)}, nil
// 	case "lda":
//
// 		p0, err := internal.Ens16bit(inst.Operands, 0)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		// zero page, absolute
// 		if len(inst.Operands) == 1 {
// 			// zero page
// 			if p0 <= 0xff {
// 				return []byte{0xa5, byte(p0)}, nil
// 			}
// 			// absolute
// 			return []byte{0xad, byte(p0), byte(p0 >> 8)}, nil
// 		}
// 		// absolute indexed, absolute indexed zero_page
// 		if len(inst.Operands) == 2 {
// 			p1, err := internal.EnsRegister(inst.Operands, 1)
// 			if err != nil {
// 				return nil, err
// 			}
//
// 			// zero page
// 			if p0 <= 0xff {
// 				if p1 == "x" {
// 					return []byte{0xb5, byte(p0)}, nil
// 				}
// 			} else {
// 				if p1 == "x" {
// 					return []byte{0xbd, byte(p0), byte(p0 >> 8)}, nil
// 				}
// 			}
// 		}
// 	}
//
// 	return nil, fmt.Errorf("unsupported instruction: '%s' '%v'", inst.OpCode, inst.Operands) //todo: make a src error with file name and line pointer
// }
