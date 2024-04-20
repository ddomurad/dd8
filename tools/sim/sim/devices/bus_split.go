package devices

import "github.com/ddomurad/dd8/tools/sim/sim"

/* implementes
type IDevice interface {
	OnNewStep()
	UpdateSubStep()
	OnStepFinished()
}
*/

type BusSplit[
	TIN sim.IntPortType,
	BUS_IN_T sim.IPort[TIN],
	TOUT sim.IntPortType,
	BUS_OUT_T sim.IPort[TIN],
] struct {
	IN  BUS_IN_T
	OUT BUS_OUT_T
}

func NewBusSplit[
	TIN sim.IntPortType,
	BUS_IN_T sim.IPort[TIN],
	TOUT sim.IntPortType,
	BUS_OUT_T sim.IPort[TIN],
](
	in BUS_IN_T,
	out BUS_OUT_T,
) *BusSplit[TIN, BUS_IN_T, TOUT, BUS_OUT_T] {
	return &BusSplit[TIN, BUS_IN_T, TOUT, BUS_OUT_T]{
		IN:  in,
		OUT: out,
	}
}

func (d *BusSplit[TIN, BUS_IN_T, TOUT, BUS_OUT_T]) OnNewStep() {}

func (d *BusSplit[TIN, BUS_IN_T, TOUT, BUS_OUT_T]) UpdateSubStep() {
	// d.OUT.Write(d.IN.Read())
}

func (d *BusSplit[TIN, BUS_IN_T, TOUT, BUS_OUT_T]) OnStepFinished() {}
