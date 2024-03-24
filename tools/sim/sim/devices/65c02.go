package devices

import "github.com/ddomurad/dd8/tools/sim/sim"

// RAMDevice is a device that represents a RAM chip.
type Cpu6502Device struct {
	// inputs
	PHI2 *sim.SignalPort

	// outputs
	RW *sim.SignalPort
	A  *sim.Bus16Port

	// bidirectional
	D *sim.Bus8Port
}

func (d *Cpu6502Device) Init() {
}

func (d *Cpu6502Device) Update(sim *sim.Simulation) bool {
	return false
}
