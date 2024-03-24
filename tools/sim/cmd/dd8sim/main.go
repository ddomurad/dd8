package main

import (
	idevices "github.com/ddomurad/dd8/tools/sim/cmd/dd8sim/devices"
	"github.com/ddomurad/dd8/tools/sim/sim"
	"github.com/ddomurad/dd8/tools/sim/sim/devices"
)

func main() {
	CLK := sim.NewSignalPort()
	RW := sim.NewSignalPort()

	A := sim.NewSignal16Port()
	D := sim.NewSignal8Port()
	OE_B := sim.NewSignalPort()
	WE_B := sim.NewSignalPort()

	RAM_CE := sim.NewSignalPort()

	clock := devices.ClockDevice{
		CLK: CLK,
	}

	icSelector := idevices.ICSelector0V2Device{
		RW: RW,
		A:  A,

		OE_B:   OE_B,
		WE_B:   WE_B,
		RAM_CE: RAM_CE,
	}

	ram := devices.RAMDevice{
		CE_B: RAM_CE,
		OE_B: OE_B,
		WE_B: WE_B,

		D: D,
		A: A,
	}

	// cpu := devices.Cpu6502Device{
	// 	PHI2: CLK,
	// 	RW:   RW,
	// 	A:    A,
	// 	D:    D,
	// }

	cpu := idevices.TestDevice{
		PHI2: CLK,
		RW:   RW,
		A:    A,
		D:    D,
	}

	board := sim.Board{
		Devices: []sim.Device{
			&clock,
			&icSelector,
			&ram,
			&cpu,
		},
		Ports: []sim.IPort{
			CLK,
			RW,
			OE_B,
			WE_B,
			RAM_CE,
			A,
			D,
		},
	}

	sim := sim.NewSimulation(&board)
	sim.Init()

	for {
		if !sim.Update(10) {
			panic("simulation did not converge")
		}

		sim.Reset()
	}
}
