package w65c02

import "github.com/ddomurad/dd8/tools/sim/sim"

type W65C02Device struct {
	PHI2_IN   *sim.SignalPort
	PHI1O_OUT *sim.SignalPort
	PHI2O_OUT *sim.SignalPort

	ML_B_OUT *sim.SignalPort
	RW_B_OUT *sim.SignalPort
	SYNC_OUT *sim.SignalPort
	VP_B_OUT *sim.SignalPort

	RES_B_IN *sim.SignalPort

	A_OUT *sim.Bus16Port
	D_BI  *sim.Bus8Port

	// IRQB *sim.SignalPort
	// NMIB *sim.SignalPort
	// SOB  *sim.SignalPort
	// RDY  *sim.SignalPort
	// BE   *sim.SignalPort

	Core W65C02Core
}

func NewW65C02Device() *W65C02Device {
	device := &W65C02Device{
		PHI2_IN:   &sim.SignalPort{},
		PHI1O_OUT: &sim.SignalPort{},
		PHI2O_OUT: &sim.SignalPort{},

		ML_B_OUT: &sim.SignalPort{},
		RW_B_OUT: &sim.SignalPort{},
		RES_B_IN: &sim.SignalPort{},
		SYNC_OUT: &sim.SignalPort{},
		VP_B_OUT: &sim.SignalPort{},

		A_OUT: &sim.Bus16Port{},
		D_BI:  &sim.Bus8Port{},
	}

	device.Core.Init()
	return device
}

func (d *W65C02Device) OnNewStep() {}

func (d *W65C02Device) UpdateSubStep() {
	phi2, _ := d.PHI2_IN.Read()

	d.Core.PHI2 = phi2

	if d.Core.RWB {
		d.Core.D, _ = d.D_BI.Read()
	}

	d.Core.Update()

	if !d.Core.RWB {
		d.D_BI.Write(d.Core.D)
	}
	d.A_OUT.Write(d.Core.A)
	d.PHI1O_OUT.Write(d.Core.PHI1O)
	d.PHI2O_OUT.Write(d.Core.PHI2O)
	d.ML_B_OUT.Write(d.Core.MLB)
	d.RW_B_OUT.Write(d.Core.RWB)
	d.VP_B_OUT.Write(d.Core.VPB)
	d.SYNC_OUT.Write(d.Core.SYNC)
}

func (d *W65C02Device) OnStepFinished() {

}
