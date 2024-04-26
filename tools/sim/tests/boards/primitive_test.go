package boards

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ddomurad/dd8/tools/sim/tests"
	"github.com/ddomurad/dd8/tools/sim/tests/boards/sources"
)

func Test_PrimitiveBoard(t *testing.T) {
	board, sim := NewPrimitiveBoard(tests.AssembleSource(t, sources.TestAsm))

	// sim.RunTill(t, func() bool {
	// 	addr, _ := board.CPU_DEV.A_OUT.Read()
	// 	return addr == 0xa006
	// }, 100)

	start := time.Now()

	for i := 0; i < 10_000_000; i++ {
		_ = sim.Step()
	}

	elapsed := time.Since(start)

	fmt.Printf("The loop took %s to complete\n", elapsed)

	assert.Equal(t, uint8(0xaa), board.RAM_DEV.Data[0x8080])
	assert.Equal(t, uint8(0xbb), board.RAM_DEV.Data[0x8090])
}
