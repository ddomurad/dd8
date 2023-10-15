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
	ast := internal.ParseSrc("test.asm", src)
	require.False(t, ast.Errors.HasErrors())
	internal.PreprocessAST(ast)
	require.False(t, ast.Errors.HasErrors())
	bcode, err := internal.Assemble(ast, assemblers.OpcodeAssemblerW65C02S)
	require.NoError(t, err)
	assertByteCodeWithArray(t, expected, bcode)
}

func TestThatCanBuildBinaryCode(t *testing.T) {
	t.Run("brk", func(t *testing.T) {
		assertAssembler(t, "brk", []byte{0x00})
	})
	t.Run("ora", func(t *testing.T) {
		assertAssembler(t, "ora [0x13,x]", []byte{0x01, 0x13})
		assertAssembler(t, "ora 0x9f", []byte{0x05, 0x9f})
		assertAssembler(t, "orai 0x13", []byte{0x09, 0x13})
		assertAssembler(t, "ora 0x199f", []byte{0x0d, 0x9f, 0x19})
		assertAssembler(t, "ora [0x13],y", []byte{0x11, 0x13})
		assertAssembler(t, "ora [0x13]", []byte{0x12, 0x13})
		assertAssembler(t, "ora 0x13,x", []byte{0x15, 0x13})
		assertAssembler(t, "ora 0x9933,y", []byte{0x19, 0x33, 0x99})
		assertAssembler(t, "ora 0x7,y", []byte{0x19, 0x07, 0x00})
		assertAssembler(t, "ora 0x13c1,x", []byte{0x1d, 0xc1, 0x13})
	})
	t.Run("bpl", func(t *testing.T) {
		assertAssembler(t, "bpl", []byte{0x10})
	})
	t.Run("nop", func(t *testing.T) {
		assertAssembler(t, "nop", []byte{0xea})
	})
	t.Run("php", func(t *testing.T) {
		assertAssembler(t, "php", []byte{0x08})
	})
	t.Run("jsr_a", func(t *testing.T) {
		assertAssembler(t, "jsr 0xab99", []byte{0x20, 0x99, 0xab})
	})
	t.Run("lda", func(t *testing.T) {
		assertAssembler(t, "lda [0xf0,x]", []byte{0xa1, 0xf0})
		assertAssembler(t, "lda 0xf0", []byte{0xa5, 0xf0})
		assertAssembler(t, "ldai 0x9a", []byte{0xa9, 0x9a})
		assertAssembler(t, "lda 0x9a71", []byte{0xad, 0x71, 0x9a})
		assertAssembler(t, "lda [0x34],y", []byte{0xb1, 0x34})
		assertAssembler(t, "lda [0x34]", []byte{0xb2, 0x34})
		assertAssembler(t, "lda 0x34,x", []byte{0xb5, 0x34})
		assertAssembler(t, "lda 0x34,y", []byte{0xb9, 0x34, 0x00})
		assertAssembler(t, "lda 0x3499,y", []byte{0xb9, 0x99, 0x34})
		assertAssembler(t, "lda 0x34,x", []byte{0xb5, 0x34})

		assertAssembler(t, "ldai 0xab", []byte{0xa9, 0xab})
		assertAssembler(t, "ldai 10", []byte{0xa9, 0x0a})
		assertAssembler(t, "ldai 0x01", []byte{0xa9, 0x01})
		assertAssembler(t, "lda 0xab", []byte{0xa5, 0xab})
		assertAssembler(t, "lda 0x10ab", []byte{0xad, 0xab, 0x10})
		assertAssembler(t, "lda 0xab_91", []byte{0xad, 0x91, 0xab})
		assertAssembler(t, "lda 0xf_91", []byte{0xad, 0x91, 0x0f})
		assertAssembler(t, "lda 0x3b, x", []byte{0xb5, 0x3b})
		assertAssembler(t, "lda [0x3b, x]", []byte{0xa1, 0x3b})
		assertAssembler(t, "lda 0x113b, x", []byte{0xbd, 0x3b, 0x11})
		assertAssembler(t, "lda [0x3b], y", []byte{0xb1, 0x3b})
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
