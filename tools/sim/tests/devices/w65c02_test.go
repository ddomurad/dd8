package devices

import (
	"testing"

	"github.com/ddomurad/dd8/tools/sim/sim/devices/w65c02"
	"github.com/stretchr/testify/assert"
)

func Test_W65C02Device(t *testing.T) {
	t.Run("full_test_program", func(t *testing.T) {
		//todo: extend reset sequence
		sim := buildTestSimulation(map[uint16][]uint8{
			0xffff: {0xff, 0xff},
			0xfffc: {0x00, 0xf0},
			0xf000: {0x18, 0x20, 0x17, 0xf0, 0x20, 0x46, 0xf0},
		})
		sim.RunSteps(t, 12)
		sim.AssertResults(t, [][]int{
			{0, 1, 0xfffc, 0x00, 0, 1},
			{1, 1, 0xfffc, 0x00, 0, 1},

			{0, 1, 0xfffd, 0x00, 0, 1},
			{1, 1, 0xfffd, 0xf0, 0, 1},

			{0, 1, 0xf000, 0xf0, 1, 1}, //2
			{1, 1, 0xf000, 0x18, 1, 1},

			{0, 1, 0xf001, 0x18, 0, 1}, //4
			{1, 1, 0xf001, 0x20, 0, 1},

			{0, 1, 0xf001, 0x20, 1, 1}, //6
			{1, 1, 0xf001, 0x20, 1, 1},

			{0, 1, 0xf002, 0x20, 0, 1}, //8
			{1, 1, 0xf002, 0x17, 0, 1},
		})
	})

	t.Run("clc", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x18, 0x20, 0x17, 0xf0},
		})
		sim.board.CPU_DEV.Core.State.SetSR(w65c02.SR_CARRY, true)
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x18, 1, 1},

			{0, 1, 0xf001, 0x18, 0, 1},
			{1, 1, 0xf001, 0x20, 0, 1},

			{0, 1, 0xf001, 0x20, 1, 1}, //fetch next instr
			{1, 1, 0xf001, 0x20, 1, 1},
		})

		assert.False(t, sim.board.CPU_DEV.Core.State.GetSR(w65c02.SR_CARRY), "SR_CARRY")
	})

	t.Run("jsr", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x20, 0x17, 0xf0, 0x20},
			0xf017: {0x18},
		})
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x20, 1, 1},

			{0, 1, 0xf001, 0x20, 0, 1},
			{1, 1, 0xf001, 0x17, 0, 1},

			{0, 1, 0x0100, 0x17, 0, 1},
			{1, 1, 0x0100, 0x00, 0, 1},

			{0, 1, 0x0100, 0x00, 0, 1},
			{1, 0, 0x0100, 0xf0, 0, 1},

			{0, 1, 0x01ff, 0xf0, 0, 1},
			{1, 0, 0x01ff, 0x02, 0, 1},

			{0, 1, 0xf002, 0x02, 0, 1},
			{1, 1, 0xf002, 0xf0, 0, 1},

			{0, 1, 0xf017, 0xf0, 1, 1},
			{1, 1, 0xf017, 0x18, 1, 1},
		})
	})

	t.Run("stz", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x9c, 0x01, 0x80, 0x18},
		})
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x9c, 1, 1},

			{0, 1, 0xf001, 0x9c, 0, 1},
			{1, 1, 0xf001, 0x01, 0, 1},

			{0, 1, 0xf002, 0x01, 0, 1},
			{1, 1, 0xf002, 0x80, 0, 1},

			{0, 1, 0x8001, 0x80, 0, 1},
			{1, 0, 0x8001, 0x00, 0, 1},

			{0, 1, 0xf003, 0x00, 1, 1},
			{1, 1, 0xf003, 0x18, 1, 1},
		})
	})

	t.Run("ldai", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0xa9, 0x42, 0x18},
		})
		sim.board.CPU_DEV.Core.State.SetSR(w65c02.SR_ZERO, true)
		sim.board.CPU_DEV.Core.State.SetSR(w65c02.SR_NEGATIVE, true)
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0xa9, 1, 1},

			{0, 1, 0xf001, 0xa9, 0, 1},
			{1, 1, 0xf001, 0x42, 0, 1},

			{0, 1, 0xf002, 0x42, 1, 1},
			{1, 1, 0xf002, 0x18, 1, 1},
		})

		assert.Equal(t, uint8(0x42), sim.board.CPU_DEV.Core.State.A, "A")
		assert.Equal(t, uint8(0x00), sim.board.CPU_DEV.Core.State.P, "P")
	})

	t.Run("ldai_z", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0xa9, 0x00, 0x18},
		})
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0xa9, 1, 1},

			{0, 1, 0xf001, 0xa9, 0, 1},
			{1, 1, 0xf001, 0x00, 0, 1},

			{0, 1, 0xf002, 0x00, 1, 1},
			{1, 1, 0xf002, 0x18, 1, 1},
		})

		assert.Equal(t, uint8(0x00), sim.board.CPU_DEV.Core.State.A, "A")
		assert.Equal(t, uint8(0x40), sim.board.CPU_DEV.Core.State.P, "P") // check for zero flag
	})

	t.Run("ldai_n", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0xa9, 0xf0, 0x18},
		})
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0xa9, 1, 1},

			{0, 1, 0xf001, 0xa9, 0, 1},
			{1, 1, 0xf001, 0xf0, 0, 1},

			{0, 1, 0xf002, 0xf0, 1, 1},
			{1, 1, 0xf002, 0x18, 1, 1},
		})

		assert.Equal(t, uint8(0xf0), sim.board.CPU_DEV.Core.State.A, "A")
		assert.Equal(t, uint8(0x01), sim.board.CPU_DEV.Core.State.P, "P") // check for negative flag
	})

	// test sta
	t.Run("sta", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x8d, 0x02, 0x80, 0x18},
		})
		sim.board.CPU_DEV.Core.State.A = 0x42
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x8d, 1, 1},

			{0, 1, 0xf001, 0x8d, 0, 1},
			{1, 1, 0xf001, 0x02, 0, 1},

			{0, 1, 0xf002, 0x02, 0, 1},
			{1, 1, 0xf002, 0x80, 0, 1},

			{0, 1, 0x8002, 0x80, 0, 1},
			{1, 0, 0x8002, 0x42, 0, 1},

			{0, 1, 0xf003, 0x42, 1, 1}, // next intruction
			{1, 1, 0xf003, 0x18, 1, 1},
		})

		assert.Equal(t, uint8(0x42), sim.board.CPU_DEV.Core.State.A, "A")
		assert.Equal(t, uint8(0x00), sim.board.CPU_DEV.Core.State.P, "P")
		sim.AssertMemory(t, 0x8002, 0x42)
	})

	t.Run("jmp", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x4c, 0x17, 0xf0},
			0xf017: {0x18, 0x18},
		})
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x4c, 1, 1},

			{0, 1, 0xf001, 0x4c, 0, 1},
			{1, 1, 0xf001, 0x17, 0, 1},

			{0, 1, 0xf002, 0x17, 0, 1},
			{1, 1, 0xf002, 0xf0, 0, 1},

			{0, 1, 0xf017, 0xf0, 1, 1}, // next opcode after jump
			{1, 1, 0xf017, 0x18, 1, 1},

			{0, 1, 0xf018, 0x18, 0, 1},
			{1, 1, 0xf018, 0x18, 0, 1},
		})
	})

	// test for rts
	t.Run("rts", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0x0100: {0x11, 0xaa, 0xbb},
			0xbbaa: {0x22, 0x18, 0x20},
			0xf000: {0x60, 0x18},
		})
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x60, 1, 1}, //

			{0, 1, 0xf001, 0x60, 0, 1},
			{1, 1, 0xf001, 0x18, 0, 1},

			{0, 1, 0x0100, 0x18, 0, 1},
			{1, 1, 0x0100, 0x11, 0, 1},

			{0, 1, 0x0101, 0x11, 0, 1},
			{1, 1, 0x0101, 0xaa, 0, 1},

			{0, 1, 0x0102, 0xaa, 0, 1},
			{1, 1, 0x0102, 0xbb, 0, 1}, // 5

			{0, 1, 0xbbaa, 0xbb, 0, 1},
			{1, 1, 0xbbaa, 0x22, 0, 1}, // 6

			{0, 1, 0xbbab, 0x22, 1, 1},
			{1, 1, 0xbbab, 0x18, 1, 1},

			{0, 1, 0xbbac, 0x18, 0, 1},
			{1, 1, 0xbbac, 0x20, 0, 1},
		})
	})

	t.Run("and", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x2d, 0x01, 0x80, 0x18},
			0x8001: {0x54},
		})
		sim.board.CPU_DEV.Core.State.A = 0xf0
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x2d, 1, 1},

			{0, 1, 0xf001, 0x2d, 0, 1},
			{1, 1, 0xf001, 0x01, 0, 1},

			{0, 1, 0xf002, 0x01, 0, 1},
			{1, 1, 0xf002, 0x80, 0, 1},

			{0, 1, 0x8001, 0x80, 0, 1},
			{1, 1, 0x8001, 0x54, 0, 1},

			{0, 1, 0xf003, 0x54, 1, 1}, // next intruction
			{1, 1, 0xf003, 0x18, 1, 1},
		})

		assert.Equal(t, uint8(0x50), sim.board.CPU_DEV.Core.State.A, "A")
		assert.Equal(t, uint8(0x00), sim.board.CPU_DEV.Core.State.P, "P")
	})

	t.Run("and_z", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x2d, 0x01, 0x80, 0x18},
			0x8001: {0x11},
		})
		sim.board.CPU_DEV.Core.State.A = 0xee
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x2d, 1, 1},

			{0, 1, 0xf001, 0x2d, 0, 1},
			{1, 1, 0xf001, 0x01, 0, 1},

			{0, 1, 0xf002, 0x01, 0, 1},
			{1, 1, 0xf002, 0x80, 0, 1},

			{0, 1, 0x8001, 0x80, 0, 1},
			{1, 1, 0x8001, 0x11, 0, 1},

			{0, 1, 0xf003, 0x11, 1, 1}, // next intruction
			{1, 1, 0xf003, 0x18, 1, 1},
		})

		assert.Equal(t, uint8(0x00), sim.board.CPU_DEV.Core.State.A, "A")
		assert.Equal(t, uint8(0x40), sim.board.CPU_DEV.Core.State.P, "P")
	})

	t.Run("and_n", func(t *testing.T) {
		sim := buildTestSimulation(map[uint16][]uint8{
			0xf000: {0x2d, 0x01, 0x80, 0x18},
			0x8001: {0x91},
		})
		sim.board.CPU_DEV.Core.State.A = 0xee
		sim.AssertSets(t, [][]int{
			{0, 1, 0xf000, 0x00, 1, 1},
			{1, 1, 0xf000, 0x2d, 1, 1},

			{0, 1, 0xf001, 0x2d, 0, 1},
			{1, 1, 0xf001, 0x01, 0, 1},

			{0, 1, 0xf002, 0x01, 0, 1},
			{1, 1, 0xf002, 0x80, 0, 1},

			{0, 1, 0x8001, 0x80, 0, 1},
			{1, 1, 0x8001, 0x91, 0, 1},

			{0, 1, 0xf003, 0x91, 1, 1}, // next intruction
			{1, 1, 0xf003, 0x18, 1, 1},
		})

		assert.Equal(t, uint8(0x80), sim.board.CPU_DEV.Core.State.A, "A")
		assert.Equal(t, uint8(0x01), sim.board.CPU_DEV.Core.State.P, "P")
	})

}

// clk, rwb, addr, data, sync, mlb

/*
 * Memory Lock (MLB)  1
 * Read/Write RWB     2
 * Reset RESB         3
 * Synchorize with opcode SYNC 4
 * Vector Pull VSS    5
 * */

// 0x07 => 0b00000111
