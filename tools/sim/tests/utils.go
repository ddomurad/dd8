package tests

import (
	"testing"

	"github.com/ddomurad/dd8/tools/sim/sim"
	"github.com/stretchr/testify/require"

	"github.com/ddomurad/dd8/tools/asm/asm"
	"github.com/ddomurad/dd8/tools/asm/asm/assemblers"
)

type MemorySourceReaderWriter struct {
	*asm.MemorySourceReader
	*asm.MemorySourceWriter
}

func NewMemorySourceReaderWriter() *MemorySourceReaderWriter {
	buf := make(map[string][]byte)
	return &MemorySourceReaderWriter{
		MemorySourceReader: asm.NewMemorySourceReader(buf),
		MemorySourceWriter: asm.NewMemorySourceWriter(buf),
	}
}

func (m *MemorySourceReaderWriter) LoadSourceFile(name string, data []byte) error {
	return m.WriteSourceFile(name, string(data))
}

func AssembleSource(t *testing.T, src []byte) asm.ByteCode {
	fs := NewMemorySourceReaderWriter()
	_ = fs.LoadSourceFile("test.asm", src)

	byteCode, _, _, aerr, err := asm.AssembleSrc(
		"test.asm", fs, assemblers.OpcodeAssemblerW65C02S, false, false)

	require.Nil(t, aerr)
	require.NoError(t, err)

	return byteCode
}

func AssertPortValue[T sim.PortType](t *testing.T, expected T, port sim.IPort[T]) {
	v, _ := port.Read()
	if v != expected {
		t.Fatalf("expected %v, got %v", expected, v)
	}
}
