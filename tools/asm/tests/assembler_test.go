package tests

import (
	"fmt"
	"testing"

	pkg "github.com/ddomurad/dd8/tools/asm/pkg"
	"github.com/ddomurad/dd8/tools/asm/pkg/assemblers"
	"github.com/stretchr/testify/require"
)

func assertByteCodeWithArray(t *testing.T, expected []byte, acctual pkg.ByteCode) {
	if len(acctual) > len(expected) {
		require.Fail(t, "len(acctual) > len(expected)")
	}
	require.Equal(t, expected, acctual.ToByteArray(0x00, len(expected)), "bype code")
}

func assertAssembler(t *testing.T, src string, expected []byte) {
	src += "\n"
	ast := pkg.ParseSrc("test.asm", src)
	for _, e := range ast.Errors.Errors {
		fmt.Printf("AST ERROR: %s at line %d\n", e.Msg, e.Line)
	}
	require.False(t, ast.Errors.HasErrors())
	// pkg.PreprocessAST(ast, nil)
	// require.False(t, ast.Errors.HasErrors())
	bcode, err := pkg.Assemble(ast, assemblers.OpcodeAssemblerW65C02S)
	require.NoError(t, err)
	assertByteCodeWithArray(t, expected, bcode)
}

func assertAssemblerBC(t *testing.T, src string, expected pkg.ByteCode) {
	src += "\n"
	ast := pkg.ParseSrc("test.asm", src)
	require.False(t, ast.Errors.HasErrors())
	// pkg.PreprocessAST(ast, nil)
	// require.False(t, ast.Errors.HasErrors())
	bcode, err := pkg.Assemble(ast, assemblers.OpcodeAssemblerW65C02S)
	require.NoError(t, err)
	require.Equal(t, expected, bcode)
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
	t.Run("tsb", func(t *testing.T) {
		assertAssembler(t, "tsb 0x32", []byte{0x04, 0x32})
		assertAssembler(t, "tsb 0x3292", []byte{0x0c, 0x92, 0x32})
	})
	t.Run("asl", func(t *testing.T) {
		assertAssembler(t, "asl 0x32", []byte{0x06, 0x32})
		assertAssembler(t, "asl a", []byte{0x0a})
		assertAssembler(t, "asl 0x3cc2", []byte{0x0e, 0xc2, 0x3c})
		assertAssembler(t, "asl 0x32,x", []byte{0x16, 0x32})
		assertAssembler(t, "asl 0x32ff,x", []byte{0x1e, 0xff, 0x32})
	})
	t.Run("rmb", func(t *testing.T) {
		assertAssembler(t, "rmb 0, 0x19", []byte{0x07, 0x19})
		assertAssembler(t, "rmb 1, 0x19", []byte{0x17, 0x19})
		assertAssembler(t, "rmb 2, 0x19", []byte{0x27, 0x19})
		assertAssembler(t, "rmb 3, 0x19", []byte{0x37, 0x19})
		assertAssembler(t, "rmb 4, 0x19", []byte{0x47, 0x19})
		assertAssembler(t, "rmb 5, 0x19", []byte{0x57, 0x19})
		assertAssembler(t, "rmb 6, 0x19", []byte{0x67, 0x19})
		assertAssembler(t, "rmb 7, 0x19", []byte{0x77, 0x19})
	})
	t.Run("php", func(t *testing.T) {
		assertAssembler(t, "php", []byte{0x08})
	})
	t.Run("bbr", func(t *testing.T) {
		assertAssembler(t, "bbr 0, 0x80, 0x10", []byte{0x0f, 0x80, 0x0d})
		assertAssembler(t, "bbr 1, 0x80, 0x10", []byte{0x1f, 0x80, 0x0d})
		assertAssembler(t, "bbr 2, 0x80, 0x10", []byte{0x2f, 0x80, 0x0d})
		assertAssembler(t, "bbr 3, 0x80, 0x10", []byte{0x3f, 0x80, 0x0d})
		assertAssembler(t, "bbr 4, 0x80, 0x10", []byte{0x4f, 0x80, 0x0d})
		assertAssembler(t, "bbr 5, 0x80, 0x10", []byte{0x5f, 0x80, 0x0d})
		assertAssembler(t, "bbr 6, 0x80, 0x10", []byte{0x6f, 0x80, 0x0d})
		assertAssembler(t, "bbr 7, 0x80, 0x10", []byte{0x7f, 0x80, 0x0d})
	})
	t.Run("bpl", func(t *testing.T) {
		assertAssembler(t, "bpl 0x10", []byte{0x10, 0x0e})
		assertAssembler(t, "bpl 0x00", []byte{0x10, 0xfe})
	})
	t.Run("trb", func(t *testing.T) {
		assertAssembler(t, "trb 0x32", []byte{0x14, 0x32})
		assertAssembler(t, "trb 0x3292", []byte{0x1c, 0x92, 0x32})
	})
	t.Run("clc", func(t *testing.T) {
		assertAssembler(t, "clc", []byte{0x18})
	})
	t.Run("inc", func(t *testing.T) {
		assertAssembler(t, "inc a", []byte{0x1a})
		assertAssembler(t, "inc 0xfe", []byte{0xe6, 0xfe})
		assertAssembler(t, "inc 0xfe00", []byte{0xee, 0x00, 0xfe})
		assertAssembler(t, "inc 0xfe,x", []byte{0xf6, 0xfe})
		assertAssembler(t, "inc 0xfe00,x", []byte{0xfe, 0x00, 0xfe})
	})
	t.Run("jsr", func(t *testing.T) {
		assertAssembler(t, "jsr 0xab99", []byte{0x20, 0x99, 0xab})
	})
	t.Run("and", func(t *testing.T) {
		assertAssembler(t, "and [0xca,x]", []byte{0x21, 0xca})
		assertAssembler(t, "and 0xca", []byte{0x25, 0xca})
		assertAssembler(t, "andi 0xda", []byte{0x29, 0xda})
		assertAssembler(t, "and 0x8ac0", []byte{0x2d, 0xc0, 0x8a})
		assertAssembler(t, "and [0x1c], y", []byte{0x31, 0x1c})
		assertAssembler(t, "and [0x1c]", []byte{0x32, 0x1c})
		assertAssembler(t, "and 0x1c, x", []byte{0x35, 0x1c})
		assertAssembler(t, "and 0x1c34, y", []byte{0x39, 0x34, 0x1c})
		assertAssembler(t, "and 0x1c34, x", []byte{0x3d, 0x34, 0x1c})
	})
	t.Run("bit", func(t *testing.T) {
		assertAssembler(t, "bit 0x9a", []byte{0x24, 0x9a})
		assertAssembler(t, "bit 0x9a88", []byte{0x2c, 0x88, 0x9a})
		assertAssembler(t, "bit 0x9a, x", []byte{0x34, 0x9a})
		assertAssembler(t, "bit 0x9a1f, x", []byte{0x3c, 0x1f, 0x9a})
		assertAssembler(t, "biti 0x9f", []byte{0x89, 0x9f})
	})
	t.Run("rol", func(t *testing.T) {
		assertAssembler(t, "rol 0x09", []byte{0x26, 0x09})
		assertAssembler(t, "rol a", []byte{0x2a})
		assertAssembler(t, "rol 0x1234", []byte{0x2e, 0x34, 0x12})
		assertAssembler(t, "rol 0x34, x", []byte{0x36, 0x34})
		assertAssembler(t, "rol 0x1234, x", []byte{0x3e, 0x34, 0x12})
	})
	t.Run("plp", func(t *testing.T) {
		assertAssembler(t, "plp", []byte{0x28})
	})
	t.Run("bmi", func(t *testing.T) {
		assertAssembler(t, "bmi 0x10", []byte{0x30, 0x0e})
	})
	t.Run("plp", func(t *testing.T) {
		assertAssembler(t, "plp", []byte{0x28})
	})
	t.Run("sec", func(t *testing.T) {
		assertAssembler(t, "sec", []byte{0x38})
	})
	t.Run("dec", func(t *testing.T) {
		assertAssembler(t, "dec a", []byte{0x3a})
		assertAssembler(t, "dec 0x95", []byte{0xc6, 0x95})
		assertAssembler(t, "dec 0x9599", []byte{0xce, 0x99, 0x95})
		assertAssembler(t, "dec 0x08,x", []byte{0xd6, 0x08})
		assertAssembler(t, "dec 0x0809,x", []byte{0xde, 0x09, 0x08})
	})
	t.Run("rti", func(t *testing.T) {
		assertAssembler(t, "rti", []byte{0x40})
	})
	t.Run("eor", func(t *testing.T) {
		assertAssembler(t, "eor [0xe1, x]", []byte{0x41, 0xe1})
		assertAssembler(t, "eor 0xe2", []byte{0x45, 0xe2})
		assertAssembler(t, "eori 0xe2", []byte{0x49, 0xe2})
		assertAssembler(t, "eor 0xe2f0", []byte{0x4d, 0xf0, 0xe2})
		assertAssembler(t, "eor [0xe2], y", []byte{0x51, 0xe2})
		assertAssembler(t, "eor [0xe2]", []byte{0x52, 0xe2})
		assertAssembler(t, "eor 0xbc, x", []byte{0x55, 0xbc})
		assertAssembler(t, "eor 0xbc00, y", []byte{0x59, 0x00, 0xbc})
		assertAssembler(t, "eor 0xbc00, x", []byte{0x5d, 0x00, 0xbc})
	})
	t.Run("lsr", func(t *testing.T) {
		assertAssembler(t, "lsr 0x09", []byte{0x46, 0x09})
		assertAssembler(t, "lsr a", []byte{0x4a})
		assertAssembler(t, "lsr 0x1234", []byte{0x4e, 0x34, 0x12})
		assertAssembler(t, "lsr 0x34, x", []byte{0x56, 0x34})
		assertAssembler(t, "lsr 0x1234, x", []byte{0x5e, 0x34, 0x12})
	})
	t.Run("nop", func(t *testing.T) {
		assertAssembler(t, "nop", []byte{0xea})
	})
	t.Run("pha", func(t *testing.T) {
		assertAssembler(t, "pha", []byte{0x48})
	})
	t.Run("jmp", func(t *testing.T) {
		assertAssembler(t, "jmp 0xb00b", []byte{0x4c, 0x0b, 0xb0})
		assertAssembler(t, "jmp [0xb00b]", []byte{0x6c, 0x0b, 0xb0})
		assertAssembler(t, "jmp [0x11]", []byte{0x6c, 0x11, 0x00})
		assertAssembler(t, "jmp [0x11, x]", []byte{0x7c, 0x11, 0x00})
		assertAssembler(t, "jmp [0x11ff, x]", []byte{0x7c, 0xff, 0x11})
	})
	t.Run("bvc", func(t *testing.T) {
		assertAssembler(t, "bvc 0x10", []byte{0x50, 0x0e})
	})
	t.Run("cli", func(t *testing.T) {
		assertAssembler(t, "cli", []byte{0x58})
	})
	t.Run("phy", func(t *testing.T) {
		assertAssembler(t, "phy", []byte{0x5a})
	})
	t.Run("rts", func(t *testing.T) {
		assertAssembler(t, "rts", []byte{0x60})
	})
	t.Run("adc", func(t *testing.T) {
		assertAssembler(t, "adc [0xe1, x]", []byte{0x61, 0xe1})
		assertAssembler(t, "adc 0xe2", []byte{0x65, 0xe2})
		assertAssembler(t, "adci 0xe2", []byte{0x69, 0xe2})
		assertAssembler(t, "adc 0xe2f0", []byte{0x6d, 0xf0, 0xe2})
		assertAssembler(t, "adc [0xe2], y", []byte{0x71, 0xe2})
		assertAssembler(t, "adc [0xe2]", []byte{0x72, 0xe2})
		assertAssembler(t, "adc 0xbc, x", []byte{0x75, 0xbc})
		assertAssembler(t, "adc 0xbc00, y", []byte{0x79, 0x00, 0xbc})
		assertAssembler(t, "adc 0xbc00, x", []byte{0x7d, 0x00, 0xbc})
	})
	t.Run("ror", func(t *testing.T) {
		assertAssembler(t, "ror 0x09", []byte{0x66, 0x09})
		assertAssembler(t, "ror a", []byte{0x6a})
		assertAssembler(t, "ror 0x1234", []byte{0x6e, 0x34, 0x12})
		assertAssembler(t, "ror 0x34, x", []byte{0x76, 0x34})
		assertAssembler(t, "ror 0x1234, x", []byte{0x7e, 0x34, 0x12})
	})
	t.Run("stz", func(t *testing.T) {
		assertAssembler(t, "stz 0x09", []byte{0x64, 0x09})
		assertAssembler(t, "stz 0x09,x", []byte{0x74, 0x09})
		assertAssembler(t, "stz 0x09cc", []byte{0x9c, 0xcc, 0x09})
		assertAssembler(t, "stz 0x09cc, x", []byte{0x9e, 0xcc, 0x09})
	})
	t.Run("pla", func(t *testing.T) {
		assertAssembler(t, "pla", []byte{0x68})
	})
	t.Run("bvs", func(t *testing.T) {
		assertAssembler(t, "bvs 0x10", []byte{0x70, 0x0e})
	})
	t.Run("sei", func(t *testing.T) {
		assertAssembler(t, "sei", []byte{0x78})
	})
	t.Run("ply", func(t *testing.T) {
		assertAssembler(t, "ply", []byte{0x7a})
	})
	t.Run("bra", func(t *testing.T) {
		assertAssembler(t, "bra 0x10", []byte{0x80, 0x0e})
	})
	t.Run("sta", func(t *testing.T) {
		assertAssembler(t, "sta [0xe1, x]", []byte{0x81, 0xe1})
		assertAssembler(t, "sta 0xe2", []byte{0x85, 0xe2})
		assertAssembler(t, "sta 0xe2f0", []byte{0x8d, 0xf0, 0xe2})
		assertAssembler(t, "sta [0xe2], y", []byte{0x91, 0xe2})
		assertAssembler(t, "sta [0xe2]", []byte{0x92, 0xe2})
		assertAssembler(t, "sta 0xbc, x", []byte{0x95, 0xbc})
		assertAssembler(t, "sta 0xbc00, y", []byte{0x99, 0x00, 0xbc})
		assertAssembler(t, "sta 0xbc00, x", []byte{0x9d, 0x00, 0xbc})
	})
	t.Run("sty", func(t *testing.T) {
		assertAssembler(t, "sty 0x09", []byte{0x84, 0x09})
		assertAssembler(t, "sty 0x09cc", []byte{0x8c, 0xcc, 0x09})
		assertAssembler(t, "sty 0x09,x", []byte{0x94, 0x09})
	})
	t.Run("stx", func(t *testing.T) {
		assertAssembler(t, "stx 0x09", []byte{0x86, 0x09})
		assertAssembler(t, "stx 0x09cc", []byte{0x8e, 0xcc, 0x09})
		assertAssembler(t, "stx 0x37,y", []byte{0x96, 0x37})
	})
	t.Run("smb", func(t *testing.T) {
		assertAssembler(t, "smb 0, 0x19", []byte{0x87, 0x19})
		assertAssembler(t, "smb 1, 0x19", []byte{0x97, 0x19})
		assertAssembler(t, "smb 2, 0x19", []byte{0xa7, 0x19})
		assertAssembler(t, "smb 3, 0x19", []byte{0xb7, 0x19})
		assertAssembler(t, "smb 4, 0x19", []byte{0xc7, 0x19})
		assertAssembler(t, "smb 5, 0x19", []byte{0xd7, 0x19})
		assertAssembler(t, "smb 6, 0x19", []byte{0xe7, 0x19})
		assertAssembler(t, "smb 7, 0x19", []byte{0xf7, 0x19})
	})
	t.Run("bbs", func(t *testing.T) {
		assertAssembler(t, "bbs 0, 0x10, [0x10]", []byte{0x8f, 0x10, 0x10})
		assertAssembler(t, "bbs 1, 0x10, [0x10]", []byte{0x9f, 0x10, 0x10})
		assertAssembler(t, "bbs 2, 0x10, [0x10]", []byte{0xaf, 0x10, 0x10})
		assertAssembler(t, "bbs 3, 0x10, [0x10]", []byte{0xbf, 0x10, 0x10})
		assertAssembler(t, "bbs 4, 0x10, [0x10]", []byte{0xcf, 0x10, 0x10})
		assertAssembler(t, "bbs 5, 0x10, [0x10]", []byte{0xdf, 0x10, 0x10})
		assertAssembler(t, "bbs 6, 0x10, [0x10]", []byte{0xef, 0x10, 0x10})
		assertAssembler(t, "bbs 7, 0x10, [0x10]", []byte{0xff, 0x10, 0x10})

		assertAssembler(t, "bbs 0, 0x10, 0x10", []byte{0x8f, 0x10, 0x0d})
		assertAssembler(t, "bbs 1, 0x10, 0x10", []byte{0x9f, 0x10, 0x0d})
		assertAssembler(t, "bbs 2, 0x10, 0x10", []byte{0xaf, 0x10, 0x0d})
		assertAssembler(t, "bbs 3, 0x10, 0x10", []byte{0xbf, 0x10, 0x0d})
		assertAssembler(t, "bbs 4, 0x10, 0x10", []byte{0xcf, 0x10, 0x0d})
		assertAssembler(t, "bbs 5, 0x10, 0x10", []byte{0xdf, 0x10, 0x0d})
		assertAssembler(t, "bbs 6, 0x10, 0x10", []byte{0xef, 0x10, 0x0d})
		assertAssembler(t, "bbs 7, 0x10, 0x10", []byte{0xff, 0x10, 0x0d})
	})
	t.Run("dey", func(t *testing.T) {
		assertAssembler(t, "dey", []byte{0x88})
	})
	t.Run("txa", func(t *testing.T) {
		assertAssembler(t, "txa", []byte{0x8a})
	})
	t.Run("bcc", func(t *testing.T) {
		assertAssembler(t, "bcc 0x10", []byte{0x90, 0x0e})
	})
	t.Run("tya", func(t *testing.T) {
		assertAssembler(t, "tya", []byte{0x98})
	})
	t.Run("txs", func(t *testing.T) {
		assertAssembler(t, "txs", []byte{0x9a})
	})
	t.Run("ldy", func(t *testing.T) {
		assertAssembler(t, "ldyi 0x72", []byte{0xa0, 0x72})
		assertAssembler(t, "ldy 0x72", []byte{0xa4, 0x72})
		assertAssembler(t, "ldy 0x7299", []byte{0xac, 0x99, 0x72})
		assertAssembler(t, "ldy 0x72, x", []byte{0xb4, 0x72})
		assertAssembler(t, "ldy 0x72ff, x", []byte{0xbc, 0xff, 0x72})
	})
	t.Run("ldx", func(t *testing.T) {
		assertAssembler(t, "ldxi 0x72", []byte{0xa2, 0x72})
		assertAssembler(t, "ldx 0x72", []byte{0xa6, 0x72})
		assertAssembler(t, "ldx 0x7299", []byte{0xae, 0x99, 0x72})
		assertAssembler(t, "ldx 0x8c,y", []byte{0xb6, 0x8c})
		assertAssembler(t, "ldx 0x8c9e,y", []byte{0xbe, 0x9e, 0x8c})
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
		assertAssembler(t, "ldai 10", []byte{0xa9, 0x0a})
		assertAssembler(t, "ldai 0x01", []byte{0xa9, 0x01})
		assertAssembler(t, "lda 0xab", []byte{0xa5, 0xab})
		assertAssembler(t, "lda 0x10ab", []byte{0xad, 0xab, 0x10})
		assertAssembler(t, "lda 0xab_91", []byte{0xad, 0x91, 0xab})
		assertAssembler(t, "lda 0xf_91", []byte{0xad, 0x91, 0x0f})
		assertAssembler(t, "lda 0x3b, x", []byte{0xb5, 0x3b})
		assertAssembler(t, "lda [0b0011_1011 , x]", []byte{0xa1, 0x3b})
		assertAssembler(t, "lda 0x11_3b, x", []byte{0xbd, 0x3b, 0x11})
		assertAssembler(t, "lda [0x3_b], y", []byte{0xb1, 0x3b})
	})
	t.Run("tay", func(t *testing.T) {
		assertAssembler(t, "tay", []byte{0xa8})
	})
	t.Run("tax", func(t *testing.T) {
		assertAssembler(t, "tax", []byte{0xaa})
	})
	t.Run("bcs", func(t *testing.T) {
		assertAssembler(t, "bcs 0x10", []byte{0xb0, 0x0e})
	})
	t.Run("clv", func(t *testing.T) {
		assertAssembler(t, "clv", []byte{0xb8})
	})
	t.Run("tsx", func(t *testing.T) {
		assertAssembler(t, "tsx", []byte{0xba})
	})
	t.Run("cpy", func(t *testing.T) {
		assertAssembler(t, "cpyi 0x8f", []byte{0xc0, 0x8f})
		assertAssembler(t, "cpy 0x8f", []byte{0xc4, 0x8f})
		assertAssembler(t, "cpy 0x8f87", []byte{0xcc, 0x87, 0x8f})
	})
	t.Run("cpm", func(t *testing.T) {
		assertAssembler(t, "cpm [0xe1, x]", []byte{0xc1, 0xe1})
		assertAssembler(t, "cpm 0xe2", []byte{0xc5, 0xe2})
		assertAssembler(t, "cpmi 0xe2", []byte{0xc9, 0xe2})
		assertAssembler(t, "cpm 0xe2f0", []byte{0xcd, 0xf0, 0xe2})
		assertAssembler(t, "cpm [0xe2], y", []byte{0xd1, 0xe2})
		assertAssembler(t, "cpm [0xe2]", []byte{0xd2, 0xe2})
		assertAssembler(t, "cpm 0xbc, x", []byte{0xd5, 0xbc})
		assertAssembler(t, "cpm 0xbc00, y", []byte{0xd9, 0x00, 0xbc})
		assertAssembler(t, "cpm 0xbc00, x", []byte{0xdd, 0x00, 0xbc})
	})
	t.Run("iny", func(t *testing.T) {
		assertAssembler(t, "iny", []byte{0xc8})
	})
	t.Run("dex", func(t *testing.T) {
		assertAssembler(t, "dex", []byte{0xca})
	})
	t.Run("wai", func(t *testing.T) {
		assertAssembler(t, "wai", []byte{0xcb})
	})
	t.Run("bne", func(t *testing.T) {
		assertAssembler(t, "bne 0x10", []byte{0xd0, 0x0e})
	})
	t.Run("cld", func(t *testing.T) {
		assertAssembler(t, "cld", []byte{0xd8})
	})
	t.Run("phx", func(t *testing.T) {
		assertAssembler(t, "phx", []byte{0xda})
	})
	t.Run("stp", func(t *testing.T) {
		assertAssembler(t, "stp", []byte{0xdb})
	})
	t.Run("cpx", func(t *testing.T) {
		assertAssembler(t, "cpxi 0x8f", []byte{0xe0, 0x8f})
		assertAssembler(t, "cpx 0x8f", []byte{0xe4, 0x8f})
		assertAssembler(t, "cpx 0x8f87", []byte{0xec, 0x87, 0x8f})
	})
	t.Run("cmp", func(t *testing.T) {
		assertAssembler(t, "cmp [0xe1, x]", []byte{0xc1, 0xe1})
		assertAssembler(t, "cmp 0xe2", []byte{0xc5, 0xe2})
		assertAssembler(t, "cmpi 0xe2", []byte{0xc9, 0xe2})
		assertAssembler(t, "cmp 0xe2f0", []byte{0xcd, 0xf0, 0xe2})
		assertAssembler(t, "cmp [0xe2], y", []byte{0xd1, 0xe2})
		assertAssembler(t, "cmp [0xe2]", []byte{0xd2, 0xe2})
		assertAssembler(t, "cmp 0xbc, x", []byte{0xd5, 0xbc})
		assertAssembler(t, "cmp 0xbc00, y", []byte{0xd9, 0x00, 0xbc})
		assertAssembler(t, "cmp 0xbc00, x", []byte{0xdd, 0x00, 0xbc})
	})
	t.Run("sbc", func(t *testing.T) {
		assertAssembler(t, "sbc [0xe1, x]", []byte{0xe1, 0xe1})
		assertAssembler(t, "sbc 0xe2", []byte{0xe5, 0xe2})
		assertAssembler(t, "sbci 0xe2", []byte{0xe9, 0xe2})
		assertAssembler(t, "sbc 0xe2f0", []byte{0xed, 0xf0, 0xe2})
		assertAssembler(t, "sbc [0xe2], y", []byte{0xf1, 0xe2})
		assertAssembler(t, "sbc [0xe2]", []byte{0xf2, 0xe2})
		assertAssembler(t, "sbc 0xbc, x", []byte{0xf5, 0xbc})
		assertAssembler(t, "sbc 0xbc00, y", []byte{0xf9, 0x00, 0xbc})
		assertAssembler(t, "sbc 0xbc00, x", []byte{0xfd, 0x00, 0xbc})
	})
	t.Run("beq", func(t *testing.T) {
		assertAssembler(t, "beq 0x10", []byte{0xf0, 0x0e})
	})
	t.Run("sed", func(t *testing.T) {
		assertAssembler(t, "sed", []byte{0xf8})
	})
	t.Run("plx", func(t *testing.T) {
		assertAssembler(t, "plx", []byte{0xfa})
	})
	t.Run("inx", func(t *testing.T) {
		assertAssembler(t, "inx", []byte{0xe8})
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
	t.Run("reverse_label_test", func(t *testing.T) {
		assertAssembler(t, `
	     nop
	     lda test_label
	     nop
	     nop
	     test_label:
	     nop
	     nop
	     `,
			[]byte{0xea, 0xa5, 0x05, 0xea, 0xea, 0xea, 0xea},
		)
	})
	t.Run("reverse_label_test_page_change", func(t *testing.T) {
		eb := pkg.ByteCode{}
		_ = eb.SetBytes(0x00, []byte{0xea, 0xad, 0xff, 0x01, 0xea, 0xea})
		_ = eb.SetBytes(0x01ff, []byte{0xea, 0xea})
		assertAssemblerBC(t, `
      nop
      lda test_label
      nop
      nop
      .org 0x01ff
      test_label:
      nop
      nop
	     `,
			eb,
		)
	})
	t.Run("reverse_label_test_page_change2", func(t *testing.T) {
		eb := pkg.ByteCode{}
		_ = eb.SetBytes(0x00, []byte{0x10, 0x04, 0xad, 0xff, 0x01, 0xea, 0xea})
		_ = eb.SetBytes(0x01ff, []byte{0xea, 0xea})
		assertAssemblerBC(t, `
      bpl im_label
      lda test_label
      nop
      im_label:
      nop
      .org 0x01ff
      test_label:
      nop
      nop
	     `,
			eb,
		)
	})
	t.Run("reverse_label_test_page_change3", func(t *testing.T) {
		eb := pkg.ByteCode{}
		_ = eb.SetBytes(0xf9, []byte{0x10, 0x06, 0xad, 0xff, 0x01, 0xad, 0x01, 0x01, 0xea})
		_ = eb.SetBytes(0x01ff, []byte{0xea, 0xea})
		assertAssemblerBC(t, `
      .org 0xf9
      bpl im_label
      lda test_label
      lda  im_label
      im_label:
      nop
      .org 0x01ff
      test_label:
      nop
      nop
	     `,
			eb,
		)
	})
	t.Run("reverse_label_test_page_change4", func(t *testing.T) {
		eb := pkg.ByteCode{}
		_ = eb.SetBytes(0xf8, []byte{0x10, 0x05, 0xad, 0xff, 0x01, 0xa5, 0xff, 0xea})
		_ = eb.SetBytes(0x01ff, []byte{0xea, 0xea})
		assertAssemblerBC(t, `
      .org 0xf8
      bpl im_label
      lda test_label
      lda  im_label
      im_label:
      nop
      .org 0x01ff
      test_label:
      nop
      nop
	     `,
			eb,
		)
	})
	t.Run("definde_bytes", func(t *testing.T) {
		assertAssembler(t, `.db 0x0d, 0x0a, 0x0f, "abc"`, []byte{0x0d, 0x0a, 0x0f, 'a', 'b', 'c'})
		assertAssembler(t, `.dw 0x0d, 0x0a, 0x0f, "abc"`, []byte{0x0d, 0x00, 0x0a, 0x00, 0x0f, 0x00, 'a', 'b', 'c', 0x00})
		assertAssembler(t, `.dw 0xf90d, 0x10a, 0xf`, []byte{0x0d, 0xf9, 0x0a, 0x01, 0x0f, 0x00})
		assertAssembler(t, `.dw "x"`, []byte{'x', 0x00})
	})
	t.Run("data_strings", func(t *testing.T) {
		assertAssembler(t, `.db "d\n\r\t\\\n\""`, []byte{'d', '\n', '\r', '\t', '\\', '\n', '"'})
		assertAssembler(t, `.dw "", "", ""`, []byte{})
		assertAssembler(t, `.db "\x00\x01\xff\x111"`, []byte{0x00, 0x01, 0xff, 0x11, '1'})
	})
	t.Run("skip_bytes", func(t *testing.T) {
		assertAssembler(t, `
      .org 0x00 
      .byte 
      .db 0xaa 
      .byte 3 
      .db 0x11
      .word 1
      .db 0xff
      .word 2 
      .db 0x01
      `, []byte{0x00, 0xaa, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x01})
	})

	t.Run("variable_test", func(t *testing.T) {
		assertAssembler(t, `
      .org 0xaaaa
      var_1: .byte 
      var_2: .word 2 
      var_3: .byte 
      .org 0x0000 
      stx var_1 
      stx var_2
      stx var_3
      `, []byte{0x8e, 0xaa, 0xaa, 0x8e, 0xab, 0xaa, 0x8e, 0xaf, 0xaa})
	})

	t.Run("expr_tests", func(t *testing.T) {
		assertAssembler(t, `
      .db 0x01 + 1
      .db 0x10 * 2
      .db 13 % 2 + 1
      .db 2 << 1 + 1
      .db 10 >> 1
      .db 11 / 3
      .db 0x10 | 0x01
      .db 0x15 & 0x03
      .db 0xaa ^ 0xff
      .dw ~0xa5
      .dw ~0x1a5
      .dw ~0x1a5+1
      .db 0x11aa.l
      .db 0x11aa.h
      .db ~0xaa.l
      .db ~0xaa.h
      .db ~0xaa.l>>1*2+1 
      .db (2+2)*2
      .db ~((0x11a9+(3>>1)).l>>1*2+1).l
      `, []byte{0x02, 0x20, 0x02, 0x08, 0x05, 0x03, 0x11, 0x01, 0x55, 0x5a,
			0xff, 0x5a, 0xfe, 0x5b, 0xfe, 0xaa, 0x11, 0x55, 0xff, 0x0a, 0x08, 0xea})
	})

	t.Run("expr_with_const_tests", func(t *testing.T) {
		assertAssembler(t, `
        .def N1 := 0xa0 | 0x0a 
        .def N2 := N1 >> 1
        .db N1
        .db N2  
      `, []byte{0xaa, 0x55})
	})

	t.Run("use_labels_in_statements", func(t *testing.T) {
		bc := pkg.ByteCode{}
		bc.SetBytes(0x1000, []byte{0xea, 0xea, 0xa9, 0x01})
		assertAssemblerBC(t, `
        .def (
          CSTART := 0x1000
          OTHER  := test_label + CSTART
        )
        .org CSTART
        nop
        test_label:
          nop
          ldai OTHER.l
      `, bc)
	})

	t.Run("register_definitions", func(t *testing.T) {
		bc := pkg.ByteCode{}
		bc.SetBytes(0x00, []byte{0xb5, 0x41})
		assertAssemblerBC(t, `
        .def (
        TEST_REG := x
        )
        lda 0x41, TEST_REG 
      `, bc)
	})
}
