package sim

type Signal byte
type Bus8 uint8
type Bus16 uint16

const (
// Low   Signal = 0
// High  Signal = 1
// Float Signal = 'z'
// PullUp   Signal = 'h'
// PullDown Signal = 'l'
)

type IPort interface {
	IsDiven() bool
	Reset()
}

type Port[T Signal | Bus8 | Bus16] struct {
	signal    T
	drivernBy Device
}

func (p *Port[T]) Drive(d Device, v T) {
	if p.drivernBy != nil && p.drivernBy != d {
		panic("port is already driven by other device")
	}

	p.drivernBy = d
	p.signal = v
}

func (p *Port[T]) Read() T {
	return p.signal
}

func (p *Port[T]) IsDiven() bool {
	return p.drivernBy != nil
}

func (p *Port[T]) Reset() {
	p.drivernBy = nil
}

type SignalPort = Port[Signal]
type Bus8Port = Port[Bus8]
type Bus16Port = Port[Bus16]

func NewSignalPort() *SignalPort {
	return &SignalPort{}
}

func NewSignal8Port() *Bus8Port {
	return &Bus8Port{}
}

func NewSignal16Port() *Bus16Port {
	return &Bus16Port{}
}

type Board struct {
	Devices []Device
	Ports   []IPort
}

type Simulation struct {
	Board *Board
}

type Device interface {
	Init()
	Update(sim *Simulation) bool
}

func NewSimulation(board *Board) *Simulation {
	return &Simulation{Board: board}
}

func (s *Simulation) Init() {
	for _, d := range s.Board.Devices {
		d.Init()
	}
	for _, p := range s.Board.Ports {
		p.Reset()
	}
}

func (s *Simulation) Update(maxCycles int) bool {
	for i := 0; i < maxCycles; i++ {
		updated := false
		for _, d := range s.Board.Devices {
			updated = updated || d.Update(s)
		}

		if !updated {
			return true
		}
	}
	return false
}

func (s *Simulation) Reset() {
	for _, p := range s.Board.Ports {
		p.Reset()
	}
}
