package assemblers

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ddomurad/dd8/tools/asm/internal"
)

func eStr(operands []any, index int) (string, error) {
	if len(operands) <= index {
		return "", fmt.Errorf("lolA") //todo: ??
	}

	sv, ok := operands[index].(internal.ASTNameOperand)
	if !ok {

		return "", fmt.Errorf("lolA") //todo: ??
	}

	return string(sv), nil
}

func e16bit(operands []any, index int) (uint16, error) {
	if len(operands) <= index {
		return 0, fmt.Errorf("lolA") //todo: ??
	}

	v := operands[index]
	switch tv := v.(type) {
	case internal.ASTNumericOperand:
		if tv > 0xffff || tv < 0 {
			return 0, fmt.Errorf("expected 16bit value, got: %x (hex)", tv)
		}
		return uint16(tv), nil
	}

	return 0, fmt.Errorf("unexpected value type, expected int64 got %v", reflect.TypeOf(v))
}

func e8bit(operands []any, index int) (uint8, error) {
	if len(operands) <= index {
		return 0, fmt.Errorf("lolA") //todo: ??
	}

	v := operands[index]
	switch tv := v.(type) {
	case internal.ASTNumericOperand:
		if tv > 0xff || tv < 0 {
			return 0, fmt.Errorf("expected 8bit value, got: %x (hex)", tv)
		}
		return uint8(tv), nil
	}

	return 0, fmt.Errorf("unexpected value type, expected int64 got %v", reflect.TypeOf(v))
}

func OpcodeAssemblerW65C02S(inst internal.ASTInstruction) ([]byte, error) {
	switch inst.OpCode {
	case "brk":
		return []byte{0x00}, nil
	case "nop":
		return []byte{0xea}, nil
	case "php":
		return []byte{0x08}, nil
	case "ldai":
		p0, err := e8bit(inst.Operands, 0)
		if err != nil {
			return nil, err
		}
		return []byte{0xa9, byte(p0)}, nil
	case "lda":
		p0, err := e16bit(inst.Operands, 0)
		if err != nil {
			return nil, err
		}

		// zero page, absolute
		if len(inst.Operands) == 1 {
			// zero page
			if p0 <= 0xff {
				return []byte{0xa5, byte(p0)}, nil
			}
			// absolute
			return []byte{0xad, byte(p0), byte(p0 >> 8)}, nil
		}
		// absolute indexed, absolute indexed zero_page
		if len(inst.Operands) == 2 {
			p1, err := eStr(inst.Operands, 1)
			if err != nil {
				return nil, err
			}

			// zero page
			if p0 <= 0xff {
				if p1 == "x" {
					return []byte{0xb5, byte(p0)}, nil
				}
			} else {
				if p1 == "x" {
					return []byte{0xbd, byte(p0), byte(p0 >> 8)}, nil
				}
			}
		}
	}

	return nil, errors.New("unsupported instruction") //todo: ??
}
