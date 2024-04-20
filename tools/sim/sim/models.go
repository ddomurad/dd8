package sim

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

type IntPortType interface {
	constraints.Integer
}

type PortType interface {
	IntPortType | bool
}

type IPort[T PortType] interface {
	Read() (T, bool)
	Write(value T)

	// Size() int
}

type Port[T PortType] struct {
	connection *PortConnection[T]
}

type SignalPort = Port[bool]
type Bus8Port = Port[uint8]
type Bus12Port = Port[uint16]
type Bus13Port = Port[uint16]
type Bus14Port = Port[uint16]
type Bus15Port = Port[uint16]
type Bus16Port = Port[uint16]

func (p *Port[T]) Read() (T, bool) {
	// note: this will crash iw the port is not connected
	// this is fine for now
	return p.connection.value, p.connection.driver != nil
}

func (p *Port[T]) Write(value T) {
	// note: this will crash iw the port is not connected
	// this is fine for now
	p.connection.value = value

	if p.connection.driver != p && p.connection.driver != nil {
		p.connection.multipleDrivers = true
	}
	if p.connection.driver != p {
		p.connection.changed = true
	}
	p.connection.driver = p
}

type IPortConnection interface {
	OnNewStep()
	BeforeSubStep()
	UpdateSubStep() bool
}

type PortConnection[T PortType] struct {
	Ports []*Port[T]

	value     T
	lastValue T
	changed   bool

	driver          *Port[T]
	multipleDrivers bool
}

func NewConnection[T PortType](ports ...*Port[T]) IPortConnection {
	connection := &PortConnection[T]{
		Ports: ports,
	}

	for _, p := range connection.Ports {
		p.connection = connection
	}

	return connection
}

func (c *PortConnection[T]) OnNewStep() {
	c.changed = true
	c.driver = nil
}

func (c *PortConnection[T]) BeforeSubStep() {
	c.multipleDrivers = false
	c.changed = false
}

func (c *PortConnection[T]) UpdateSubStep() bool {
	if c.driver != nil {
		if c.value != c.lastValue {
			c.changed = true
			c.lastValue = c.value
		}
	}

	return c.changed
}

type IDevice interface {
	OnNewStep()
	UpdateSubStep()
	OnStepFinished()
}

type Simulation struct {
	Devices     []IDevice
	Connections []IPortConnection
}

func NewSimulation(devices []IDevice, connections []IPortConnection) *Simulation {
	return &Simulation{
		Devices:     devices,
		Connections: connections,
	}
}

func (s *Simulation) Step() error {
	for _, c := range s.Connections {
		c.OnNewStep()
	}

	for _, d := range s.Devices {
		d.OnNewStep()
	}

	for i := 0; ; i++ {
		for _, c := range s.Connections {
			c.BeforeSubStep()
		}

		for _, d := range s.Devices {
			d.UpdateSubStep()
		}

		signalChange := false
		for _, c := range s.Connections {
			signalChange = c.UpdateSubStep() || signalChange
		}
		if !signalChange {
			break
		}

		if i > 10 {
			return errors.New("simulation did not converge")
		}
	}

	for _, d := range s.Devices {
		d.OnStepFinished()
	}
	return nil
}

func (s *Simulation) RunTill(t *testing.T, condition func() bool, maxSteps int) {
	for i := 0; i < maxSteps; i++ {
		err := s.Step()
		require.NoError(t, err)
		if condition() {
			err := s.Step()
			require.NoError(t, err)
			return
		}
	}

	require.Fail(t, "condition not met")
}
