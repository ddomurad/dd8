package devices

import "github.com/ddomurad/dd8/tools/sim/sim"

type ICSelector0V2Device struct {
	// inputs
	RW *sim.SignalPort
	A  *sim.Bus16Port

	lastRW sim.Signal
	lastA  sim.Bus16

	// outputs
	RAM_CE *sim.SignalPort
	OE_B   *sim.SignalPort
	WE_B   *sim.SignalPort
}

func (d *ICSelector0V2Device) Init() {
}

func (d *ICSelector0V2Device) Update(sim *sim.Simulation) bool {
	if d.RW.Read() == d.lastRW && d.A.Read() == d.lastA {
		return false
	}

	if d.RW.Read() == 1 {
		d.OE_B.Drive(d, 0)
		d.WE_B.Drive(d, 1)
	} else {
		d.OE_B.Drive(d, 1)
		d.WE_B.Drive(d, 0)
	}

	if d.A.Read() < 0x1000 {
		d.RAM_CE.Drive(d, 1)
	} else {
		d.RAM_CE.Drive(d, 0)
	}

	return false
}
