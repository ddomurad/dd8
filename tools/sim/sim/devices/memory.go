package devices

import "github.com/ddomurad/dd8/tools/sim/sim"

type MemoryDevice[T sim.IntPortType, ADDR_BUS_T sim.IPort[T]] struct {
	ADDR_IN ADDR_BUS_T
	DATA_BI *sim.Bus8Port

	CE_B_IN *sim.SignalPort
	WE_B_IN *sim.SignalPort
	OE_B_IN *sim.SignalPort

	shouldStore bool
	storeAddr   T
	storeData   uint8

	Data []uint8
}

func NewMemoryDevice[T sim.IntPortType, ADDR_BUS_T sim.IPort[T]](port ADDR_BUS_T, bits int) *MemoryDevice[T, ADDR_BUS_T] {
	return &MemoryDevice[T, ADDR_BUS_T]{
		ADDR_IN: port,
		DATA_BI: &sim.Bus8Port{},
		CE_B_IN: &sim.SignalPort{},
		WE_B_IN: &sim.SignalPort{},
		OE_B_IN: &sim.SignalPort{},

		Data: make([]uint8, 1<<bits),
	}
}

func (d *MemoryDevice[T, ADDR_BUS_T]) OnNewStep() {
	d.shouldStore = false
}

func (d *MemoryDevice[T, ADDR_BUS_T]) UpdateSubStep() {
	ce_b, ce_b_d := d.CE_B_IN.Read()
	if !ce_b_d {
		return
	}

	if !ce_b {
		addr, _ := d.ADDR_IN.Read()
		we_b, _ := d.WE_B_IN.Read()
		oe_b, _ := d.OE_B_IN.Read()

		if !we_b {
			data, _ := d.DATA_BI.Read()
			d.shouldStore = true
			d.storeAddr = addr
			d.storeData = data
		}
		if !oe_b {
			// note: for now, let it crash if the address is out of bounds
			d.DATA_BI.Write(d.Data[addr])
			d.shouldStore = false
		}
	}
}

func (d *MemoryDevice[T, ADDR_BUS_T]) OnStepFinished() {
	// note: for now, let it crash if the address is out of bounds
	if d.shouldStore {
		d.Data[d.storeAddr] = d.storeData
	}
}
