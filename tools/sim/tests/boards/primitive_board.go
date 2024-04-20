package boards

import (
	"github.com/ddomurad/dd8/tools/asm/asm"
	"github.com/ddomurad/dd8/tools/sim/sim"
	simdev "github.com/ddomurad/dd8/tools/sim/sim/devices"
	simdevprim "github.com/ddomurad/dd8/tools/sim/sim/devices/primitives"
	simw65 "github.com/ddomurad/dd8/tools/sim/sim/devices/w65c02"
)

type primitiveBoardController struct {
	CLK_IN *sim.SignalPort
	RWB_IN *sim.SignalPort

	RAM_CE_B_OUT *sim.SignalPort
	OE_B_OUT     *sim.SignalPort
	WE_B_OUT     *sim.SignalPort
}

func newPrimitiveBoardController() *primitiveBoardController {
	return &primitiveBoardController{
		CLK_IN: &sim.SignalPort{},
		RWB_IN: &sim.SignalPort{},

		RAM_CE_B_OUT: &sim.SignalPort{},
		OE_B_OUT:     &sim.SignalPort{},
		WE_B_OUT:     &sim.SignalPort{},
	}
}

func (d *primitiveBoardController) OnNewStep() {}

func (d *primitiveBoardController) UpdateSubStep() {
	clk, clk_dr := d.CLK_IN.Read()
	rwb, rwb_dr := d.RWB_IN.Read()

	if !clk_dr && !rwb_dr {
		return
	}

	d.RAM_CE_B_OUT.Write(!clk)
	d.OE_B_OUT.Write(!rwb)
	d.WE_B_OUT.Write(rwb)
}

func (d *primitiveBoardController) OnStepFinished() {}

type PrimitiveBoard struct {
	CLK_DEV *simdev.ClockDevice
	CPU_DEV *simw65.W65C02Device
	RAM_DEV *simdev.MemoryDevice[uint16, *sim.Bus16Port]

	CTRL_DEV *primitiveBoardController
}

func NewPrimitiveBoard(byteCode asm.ByteCode) (*PrimitiveBoard, *sim.Simulation) {
	constHigh := simdevprim.NewConstantDevice(&sim.SignalPort{}, true)
	clk := simdev.NewClockDevice()
	cpu := simw65.NewW65C02Device()
	ram := simdev.NewMemoryDevice(&sim.Bus16Port{}, 16)
	ctrl := newPrimitiveBoardController()

	board := &PrimitiveBoard{
		CLK_DEV:  clk,
		CPU_DEV:  cpu,
		RAM_DEV:  ram,
		CTRL_DEV: ctrl,
	}

	sim := sim.NewSimulation(
		[]sim.IDevice{constHigh, clk, cpu, ram, ctrl},
		[]sim.IPortConnection{
			sim.NewConnection(constHigh.OUT, cpu.RES_B_IN),

			sim.NewConnection(clk.CLK_OUT, cpu.PHI2_IN, ctrl.CLK_IN),
			sim.NewConnection(cpu.RW_B_OUT, ctrl.RWB_IN),

			sim.NewConnection(ctrl.RAM_CE_B_OUT, ram.CE_B_IN),
			sim.NewConnection(ctrl.OE_B_OUT, ram.OE_B_IN),
			sim.NewConnection(ctrl.WE_B_OUT, ram.WE_B_IN),

			sim.NewConnection(cpu.A_OUT, ram.ADDR_IN),
			sim.NewConnection(cpu.D_BI, ram.DATA_BI),
			sim.NewConnection(cpu.PHI1O_OUT),
			sim.NewConnection(cpu.PHI2O_OUT),
			sim.NewConnection(cpu.ML_B_OUT),
			sim.NewConnection(cpu.VP_B_OUT),
			sim.NewConnection(cpu.SYNC_OUT),
		},
	)

	for a, d := range byteCode {
		ram.Data[a] = d
	}

	return board, sim
}
