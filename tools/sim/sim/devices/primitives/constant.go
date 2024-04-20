package primitives

import "github.com/ddomurad/dd8/tools/sim/sim"

type ConstantDevice[T sim.PortType, OUT_T sim.IPort[T]] struct {
	OUT OUT_T

	value T
}

func NewConstantDevice[T sim.PortType, OUT_T sim.IPort[T]](port OUT_T, value T) *ConstantDevice[T, OUT_T] {
	return &ConstantDevice[T, OUT_T]{
		OUT:   port,
		value: value,
	}
}

func (d *ConstantDevice[T, OUT_T]) OnNewStep() {
	d.OUT.Write(d.value)
}

func (d *ConstantDevice[T, OUT_T]) UpdateSubStep() {
}

func (d *ConstantDevice[T, OUT_T]) OnStepFinished() {}
