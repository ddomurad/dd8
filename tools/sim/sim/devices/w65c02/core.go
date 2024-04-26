package w65c02

import "fmt"

type W65C02StatusRegister int

const (
	SR_NEGATIVE W65C02StatusRegister = 0
	SR_OVERFLOW W65C02StatusRegister = 1
	// skip
	SR_BRK         W65C02StatusRegister = 3
	SR_DECIMAL     W65C02StatusRegister = 4
	SR_IRQ_DISABLE W65C02StatusRegister = 5
	SR_ZERO        W65C02StatusRegister = 6
	SR_CARRY       W65C02StatusRegister = 7
)

type W65C02CoreState struct {
	PC      uint16
	A, X, Y uint8
	P, S    uint8
	TCU     uint8

	RstSeq uint8

	nextRST uint8
	nextPC  uint16
	nextTCU uint8
	nextS   uint8
	operand uint16

	opCode uint8
}

type W65C02Core struct {
	PHI2  bool
	PHI1O bool
	PHI2O bool

	RESB bool
	MLB  bool
	RWB  bool
	SYNC bool
	VPB  bool

	A uint16
	D uint8

	nextRWB bool
	D_buf   uint8
	D_out   uint8
	State   W65C02CoreState
}

func (c W65C02CoreState) GetSR(sr W65C02StatusRegister) bool {
	return c.P&(1<<sr) == 1<<sr
}

func (c *W65C02CoreState) SetSR(sr W65C02StatusRegister, v bool) {
	if v {
		c.P |= 1 << sr
	} else {
		c.P &= ^(1 << sr)
	}
}

func (c *W65C02CoreState) UpdateSR_ZN(v uint8) {
	z := 0
	if v == 0 {
		z = 1
	}
	n := int(v) >> 7

	c.P = c.P&0xbe | uint8(n<<SR_NEGATIVE) | uint8(z<<SR_ZERO)
}

func (mcu *W65C02Core) Init() {
	mcu.MLB = true
	mcu.RWB = true
	mcu.VPB = true
	mcu.SYNC = false

	mcu.A = 0xfffc
	mcu.D = 0x00

}

func (m *W65C02Core) Update() {
	if m.PHI2 {
		m.updateClkHigh()
	} else {
		m.updateClkLow()
	}
}

func (m *W65C02Core) updateClkHigh() {
	m.RWB = m.nextRWB
	if !m.RWB {
		m.D = m.D_out
	}
	m.D_buf = m.D

	m.State.PC = m.State.nextPC
	m.State.TCU = m.State.nextTCU
	m.State.S = m.State.nextS
	m.State.RstSeq = m.State.nextRST
}

func (m *W65C02Core) updateClkLow() {
	m.nextRWB = true
	m.RWB = true

	if m.State.RstSeq == 0 {
		m.A = 0xfffc
		m.State.nextRST = 1
		return
	}
	if m.State.RstSeq == 1 {
		m.A = 0xfffd
		m.State.PC = uint16(m.D_buf)
		m.State.nextRST = 2
		return
	}
	if m.State.RstSeq == 2 { // fix: please refactore me, this is a hack
		m.State.TCU = 0
		m.State.PC |= uint16(m.D_buf) << 8
		m.State.nextRST = 3
	}

	m.SYNC = m.State.TCU == 0 // note: fetch opcode
	if m.State.TCU == 0 {
		m.State.nextPC = m.State.PC + 1
		m.A = m.State.PC
		m.State.nextTCU = 1
		return //todo: ????
	}

	if m.State.TCU == 1 { //note: store opcode
		m.State.opCode = m.D_buf
	}

	switch m.State.opCode {
	case OpCode_CLC:
		m.handle_CLC()
	case OpCode_JSR:
		m.handle_JSR()
	case OpCode_STZ:
		m.handle_STx_a(nil)
	case OpCode_LDAI:
		m.handle_LDAI()
	case OpCode_STA:
		m.handle_STx_a(&m.State.A)
	case OpCode_JMP:
		m.handle_JMP()
	case OpCode_RTS:
		m.handle_RTS()
	case OPCode_AND:
		m.handle_AND()
	case OpCode_BEQ:
		m.handle_BEQ()
	default:
		panic(fmt.Sprintf("unexpected opcode, %x", m.State.opCode))
	}
}

func (m *W65C02Core) handle_CLC() {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC
		m.State.nextTCU = 2
	case 2:
		m.State.SetSR(SR_CARRY, false)
		m.State.TCU = 0
		m.State.nextTCU = 1
	default:
		panic("unexpected TCU")
	}
}

func (m *W65C02Core) handle_JSR() {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 2
	case 2:
		m.A = 0x0100 | uint16(m.State.S)
		m.State.operand = uint16(m.D_buf)
		m.State.nextTCU = 3
	case 3:
		m.A = 0x0100 | uint16(m.State.S)
		m.D_out = uint8(m.State.PC >> 8)
		m.State.nextTCU = 4
		m.State.nextS = m.State.S - 1
		m.nextRWB = false
	case 4:
		m.A = 0x0100 | uint16(m.State.S)
		m.D_out = uint8(m.State.PC & 0xff)
		m.State.nextTCU = 5
		m.State.nextS = m.State.S - 1
		m.nextRWB = false
	case 5:
		m.A = m.State.PC
		m.State.nextTCU = 6
	case 6:
		m.State.PC = m.State.operand | uint16(m.D_buf)<<8
		m.State.nextPC = m.State.PC
		m.A = m.State.PC
		m.State.TCU = 0
		m.State.nextTCU = 1
	default:
		panic("unexpected TCU")
	}
}

func (m *W65C02Core) handle_STx_a(v *uint8) {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 2
	case 2:
		m.A = m.State.PC
		m.State.operand = uint16(m.D_buf)
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 3
	case 3:
		m.A = m.State.operand | uint16(m.D_buf)<<8
		m.nextRWB = false
		if v != nil {
			m.D_out = *v
		} else {
			m.D_out = 0x00
		}
		m.State.nextTCU = 0
	default:
		panic("unexpected TCU")
	}
}

func (m *W65C02Core) handle_LDAI() {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 2
	case 2:
		m.A = m.State.PC
		m.State.A = m.D_buf
		m.State.UpdateSR_ZN(m.State.A)
		m.State.nextPC = m.State.PC + 1
		m.State.TCU = 0
		m.State.nextTCU = 1
	default:
		panic("unexpected TCU")
	}
}

func (m *W65C02Core) handle_JMP() {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 2
	case 2:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.operand = uint16(m.D_buf)
		m.State.nextTCU = 3
	case 3:
		m.State.PC = m.State.operand | uint16(m.D_buf)<<8
		m.State.nextPC = m.State.PC
		m.A = m.State.PC
		m.State.TCU = 0
		m.State.nextTCU = 1
	default:
		panic("unexpected TCU")
	}
}

func (m *W65C02Core) handle_RTS() {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 2
	case 2:
		m.A = 0x0100 | uint16(m.State.S)
		m.State.operand = uint16(m.D_buf)
		m.State.nextS = m.State.S + 1
		m.State.nextTCU = 3
	case 3:
		m.A = 0x0100 | uint16(m.State.S)
		m.State.nextTCU = 4
		m.State.nextS = m.State.S + 1
	case 4:
		m.A = 0x0100 | uint16(m.State.S)
		m.State.operand = uint16(m.D_buf)
		m.State.nextTCU = 5
	case 5:
		m.State.PC = m.State.operand | uint16(m.D_buf)<<8
		m.State.nextPC = m.State.PC + 1
		m.A = m.State.PC
		m.State.nextTCU = 0
	default:
		panic("unexpected TCU")
	}
}

func (m *W65C02Core) handle_AND() {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 2
	case 2:
		m.A = m.State.PC
		m.State.operand = uint16(m.D_buf)
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 3
	case 3:
		m.A = m.State.operand | uint16(m.D_buf)<<8
		m.State.nextPC = m.State.PC
		m.State.nextTCU = 4
	case 4:
		m.A = m.State.PC
		m.State.A = m.State.A & m.D_buf
		m.State.UpdateSR_ZN(m.State.A)
		m.State.TCU = 0
		m.State.nextTCU = 1
	default:
		panic("unexpected TCU")
	}
}

func (m *W65C02Core) handle_BEQ() {
	switch m.State.TCU {
	case 1:
		m.A = m.State.PC
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 2
	case 2:
		m.A = m.State.PC
		m.State.operand = uint16(m.D_buf)
		m.State.nextPC = m.State.PC + 1
		m.State.nextTCU = 3
	case 3:
		m.A = m.State.operand | uint16(m.D_buf)<<8
		m.State.nextPC = m.State.PC
		m.State.nextTCU = 4
	case 4:
		if m.State.GetSR(SR_ZERO) {
			m.State.PC = m.State.operand | uint16(m.D_buf)<<8
		}
		m.State.nextPC = m.State.PC + 1
		m.A = m.State.PC
		m.State.TCU = 0
		m.State.nextTCU = 1
	default:
		panic("unexpected TCU")
	}
}
