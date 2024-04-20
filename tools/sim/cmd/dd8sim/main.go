package main

import (
	"fmt"

	"github.com/ddomurad/dd8/tools/sim/sim"
	"github.com/ddomurad/dd8/tools/sim/sim/devices"
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
	sim := buildSimulation()
	for {
		fmt.Println("Step")
		err := sim.Step()
		if err != nil {
			panic(err)
		}
	}
}
