package devices

import (
	"github.com/ddomurad/dd8/tools/sim/sim"
)

type TestDevice struct {
	// inputs
	PHI2     *sim.SignalPort
	lastPhi2 sim.Signal

	// outputs
	RW *sim.SignalPort
	A  *sim.Bus16Port
	D  *sim.Bus8Port
}

func (d *TestDevice) Init() {
}

func (d *TestDevice) Update(s *sim.Simulation) bool {
	if d.PHI2.Read() == 1 && d.lastPhi2 == 0 {
		d.RW.Drive(d, 0) // write
		d.A.Drive(d, 0x0000)
		d.D.Drive(d, 0xaa)
		d.lastPhi2 = 1
		return true
	}

	if d.PHI2.Read() == 0 && d.lastPhi2 == 1 {
		d.RW.Drive(d, 1)
		d.A.Drive(d, 0x0000)
		d.D.Drive(d, 0x00)
		d.lastPhi2 = 0
		return true
	}

	return false
}
