package assemblers

import (
	"errors"

	"github.com/ddomurad/dd8/tools/asm/internal"
)

func OpcodeAssemblerW65C02S(inst internal.ASTInstruction) ([]byte, error) {
	switch inst.OpCode {
	case "brk":
		return []byte{0x00}, nil
	case "nop":
		return []byte{0xea}, nil
	case "php":
		return []byte{0x08}, nil
	case "ldai":
		p0, err := internal.Ens8bit(inst.Operands, 0)
		if err != nil {
			return nil, err
		}
		return []byte{0xa9, byte(p0)}, nil
	case "lda":
		p0, err := internal.Ens16bit(inst.Operands, 0)
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
			p1, err := internal.EnsRegister(inst.Operands, 1)
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
