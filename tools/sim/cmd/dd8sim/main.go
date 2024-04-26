package main

import (
	"fmt"
	"time"

	"github.com/ddomurad/dd8/tools/asm/asm"
	"github.com/ddomurad/dd8/tools/asm/asm/assemblers"
	"github.com/ddomurad/dd8/tools/sim/sim"
	"github.com/ddomurad/dd8/tools/sim/sim/devices"
	"github.com/ddomurad/dd8/tools/sim/tests"
	"github.com/ddomurad/dd8/tools/sim/tests/boards"
	"github.com/ddomurad/dd8/tools/sim/tests/boards/sources"
)

func buildSimulation() sim.Simulation {
	clkDevice := devices.NewClockDevice()
	ram32KbDev := devices.NewMemoryDevice(&sim.Bus15Port{}, 15)
	rom8KbDev := devices.NewMemoryDevice(&sim.Bus13Port{}, 13)

	return *sim.NewSimulation(
		[]sim.IDevice{
			clkDevice,
			ram32KbDev,
		},
		[]sim.IPortConnection{
			sim.NewConnection(clkDevice.CLK_OUT, ram32KbDev.CE_B_IN, rom8KbDev.CE_B_IN),
		},
	)
}

func main() {

	fs := tests.NewMemorySourceReaderWriter()
	_ = fs.LoadSourceFile("test.asm", sources.TestAsm)

	byteCode, _, _, _, err := asm.AssembleSrc(
		"test.asm", fs, assemblers.OpcodeAssemblerW65C02S, false, false)
	if err != nil {
		panic(err)
	}

	_, sim := boards.NewPrimitiveBoard(byteCode)

	start := time.Now()

	for i := 0; i < 10_000_000; i++ {
		_ = sim.Step()
	}

	elapsed := time.Since(start)

	fmt.Printf("The loop took %s to complete\n", elapsed)
	fmt.Printf("%d\n", 10_000_000/elapsed.Milliseconds())
}
