package devices

import "github.com/ddomurad/dd8/tools/sim/sim"

type ClockDevice struct {
	CLK *sim.SignalPort
}

func (d *ClockDevice) Init() {
}

func (d *ClockDevice) Update(sim *sim.Simulation) bool {
	// todo: add reset step after each iteration
	return false
}
