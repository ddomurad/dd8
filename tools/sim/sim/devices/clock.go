package devices

import "github.com/ddomurad/dd8/tools/sim/sim"

type ClockDevice struct {
	CLK_OUT *sim.SignalPort

	clk int
}

func NewClockDevice() *ClockDevice {
	return &ClockDevice{
		CLK_OUT: &sim.SignalPort{},
		clk:     1,
	}
}

func (d *ClockDevice) OnNewStep() {
	d.clk++
	d.CLK_OUT.Write(d.clk%2 == 1)
}

func (d *ClockDevice) UpdateSubStep() {}

func (d *ClockDevice) OnStepFinished() {}
