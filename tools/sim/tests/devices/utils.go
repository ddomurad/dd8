package devices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ddomurad/dd8/tools/sim/sim"
	"github.com/ddomurad/dd8/tools/sim/tests/boards"
)

type stepResult struct {
	clk  bool
	rwb  bool
	addr uint16
	data uint8
	sync bool
	mlb  bool
}

type simulation struct {
	board   *boards.PrimitiveBoard
	sim     *sim.Simulation
	results []stepResult
}

func buildTestSimulation(memoryMap map[uint16][]uint8) *simulation {
	board, sim := boards.NewPrimitiveBoard(nil)

	for addr, data := range memoryMap {
		for i, d := range data {
			board.RAM_DEV.Data[addr+uint16(i)] = d
		}
	}

	return &simulation{
		board: board,
		sim:   sim,
	}
}

func (s *simulation) RunSteps(t *testing.T, n int) {
	for i := 0; i < n; i++ {
		err := s.sim.Step()
		require.NoError(t, err)

		clk_v, clk_d := s.board.CLK_DEV.CLK_OUT.Read()
		rwb_v, rwb_d := s.board.CPU_DEV.RW_B_OUT.Read()
		addr_v, addr_d := s.board.CPU_DEV.A_OUT.Read()
		data_v, _ := s.board.CPU_DEV.D_BI.Read()
		mlb_v, mlb_d := s.board.CPU_DEV.ML_B_OUT.Read()
		sync_v, sync_d := s.board.CPU_DEV.SYNC_OUT.Read()

		require.True(t, clk_d, "clk expected to be driven")
		require.True(t, rwb_d, "rwb expected to be driven")
		require.True(t, addr_d, "addr expected to be driven")
		require.True(t, mlb_d, "mlb expected to be driven")
		require.True(t, sync_d, "sync expected to be driven")
		// require.True(t, data_d, "data expected to be driven")

		s.results = append(s.results, stepResult{
			clk:  clk_v,
			rwb:  rwb_v,
			addr: addr_v,
			data: data_v,
			mlb:  mlb_v,
			sync: sync_v,
		})
	}
}

func (s *simulation) AssertResults(t *testing.T, expected [][]int) {
	for i, r := range s.results {
		t.Logf("%d: %v\n", i, r)
	}

	for i, e := range expected {
		r := s.results[i]
		require.Equal(t, e[0] != 0, r.clk, "clk")
		require.Equal(t, e[1] != 0, r.rwb, "rwb")
		require.Equal(t, e[2], int(r.addr), "addr")
		require.Equal(t, e[3], int(r.data), "data")
		require.Equal(t, e[4] != 0, r.sync, "sync")
		require.Equal(t, e[5] != 0, r.mlb, "mlb")
	}
}

func (s *simulation) AssertSets(t *testing.T, expected [][]int) {
	s.board.CPU_DEV.Core.State.RstSeq = 2
	s.board.CPU_DEV.Core.State.PC = uint16(expected[0][2])

	s.RunSteps(t, len(expected)+1)
	s.AssertResults(t, expected)
}

func (s *simulation) AssertMemory(t *testing.T, addr uint16, expected ...uint8) {
	for i, e := range expected {
		require.Equal(t, e, s.board.RAM_DEV.Data[addr+uint16(i)], fmt.Sprintf("memory[%04x]", addr+uint16(i)))
	}
}
