package tests

import (
	"testing"

	"github.com/ddomurad/dd8/tools/asm/internal"
	"github.com/ddomurad/dd8/tools/asm/internal/assemblers"
	"github.com/stretchr/testify/require"
)

func assertByteCodeWithArray(t *testing.T, expected []byte, acctual internal.ByteCode) {
	require.Equal(t, expected, acctual.ToByteArray(0x00, len(expected)), "bype code")
}

func assertAssembler(t *testing.T, src string, expected []byte) {
	src += "\n"
	ast := internal.ParseSrc(src)
	err := internal.PreprocessAST(ast)
	require.NoError(t, err)
	bcode, err := internal.Assemble(ast, assemblers.OpcodeAssemblerW65C02S)
	require.NoError(t, err)
	assertByteCodeWithArray(t, expected, bcode)
}

func TestThatCanBuildBinaryCode(t *testing.T) {
	t.Run("brk", func(t *testing.T) {
		assertAssembler(t, "brk", []byte{0x00})
	})
	t.Run("nop", func(t *testing.T) {
		assertAssembler(t, "nop", []byte{0xea})
	})
	t.Run("php", func(t *testing.T) {
		assertAssembler(t, "php", []byte{0x08})
	})
	t.Run("ldai", func(t *testing.T) {
		assertAssembler(t, "ldai 0xab", []byte{0xa9, 0xab})
		assertAssembler(t, "ldai 10", []byte{0xa9, 0x0a})
		assertAssembler(t, "ldai 0x01", []byte{0xa9, 0x01})
	})
	t.Run("lda_zp", func(t *testing.T) {
		assertAssembler(t, "lda 0xab", []byte{0xa5, 0xab})
	})
	t.Run("lda_a", func(t *testing.T) {
		assertAssembler(t, "lda 0x10ab", []byte{0xad, 0xab, 0x10})
		assertAssembler(t, "lda 0xab_91", []byte{0xad, 0x91, 0xab})
		assertAssembler(t, "lda 0xf_91", []byte{0xad, 0x91, 0x0f})
	})
	t.Run("lda_ax_zp", func(t *testing.T) {
		assertAssembler(t, "lda 0x3b, x", []byte{0xb5, 0x3b})
	})
	t.Run("lda_ax", func(t *testing.T) {
		assertAssembler(t, "lda 0x113b, x", []byte{0xbd, 0x3b, 0x11})
	})
	t.Run("label_test", func(t *testing.T) {
		assertAssembler(t, `
      nop
      test_label:
      nop
      nop
      lda test_label `,
			[]byte{0xea, 0xea, 0xea, 0xa5, 0x01},
		)
	})
	t.Run("origin", func(t *testing.T) {
		assertAssembler(t, `
      .org 0x02
      nop
      nop
      .org 10
      nop
      .org 0x00
      nop`,
			[]byte{0xea, 0x00, 0xea, 0xea, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xea},
		)
	})
	t.Run("define", func(t *testing.T) {
		assertAssembler(t, `
      .def VAR1 := 10
      .def VAR2 := 0x20
      .def VAR3 := VAR2
      .org VAR1 
      nop
      .org 0x00
      ldai VAR3`,
			[]byte{0xa9, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xea},
		)
	})
}
