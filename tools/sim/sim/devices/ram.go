package devices

import "github.com/ddomurad/dd8/tools/sim/sim"

type RAMDevice struct {
	CE_B *sim.SignalPort
	OE_B *sim.SignalPort
	WE_B *sim.SignalPort

	lastCE_B sim.Signal
	lastOE_B sim.Signal
	lastWE_B sim.Signal

	D *sim.Bus8Port
	A *sim.Bus16Port

	ram [65536]byte
}

func (d *RAMDevice) Init() {
}

func (d *RAMDevice) Update(sim *sim.Simulation) bool {
	if d.CE_B.Read() == d.lastCE_B &&
		d.OE_B.Read() == d.lastOE_B &&
		d.WE_B.Read() == d.lastWE_B {
		return false
	}

	if d.CE_B.Read() == 1 {
		return false
	}

	if d.WE_B.Read() == 0 && d.OE_B.Read() == 0 {
		panic("RAMDevice: both WE_B and OE_B are low")
	}

	if d.WE_B.Read() == 0 {
		d.ram[d.A.Read()] = byte(d.D.Read())
		return true
	}

	return false
}
