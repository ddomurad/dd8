package tests

import (
	"testing"

	"github.com/ddomurad/dd8/tools/asm/asm"
	"github.com/ddomurad/dd8/tools/asm/asm/assemblers"
	"github.com/stretchr/testify/require"
)

func assertListing(t *testing.T, src string, expectedListing string) {
	src += "\n"
	ast := asm.ParseSrc("test.asm", src)
	require.False(t, ast.Errors.HasErrors())
	listing := asm.NewSourceListing()
	listing.AddSource("test.asm", src)
	_, err := asm.Assemble(ast, assemblers.OpcodeAssemblerW65C02S, listing)
	require.Nil(t, err)
	actuallListing := listing.FileListings["test.asm"].String()
	require.Equal(t, expectedListing, actuallListing)
}

func Test_ListingGeneration(t *testing.T) {
	t.Run("list_source_lines", func(t *testing.T) {
		assertListing(t, `.def TEST := 1 
; this is a comment`,
			`Line  Loc   Code         Source
    1                   .def TEST := 1 
    2                   ; this is a comment
    3                   
`)
	})
	t.Run("list_opcodes_from_ast", func(t *testing.T) {
		assertListing(t, `ldai 0xaa`,
			`Line  Loc   Code         Source
    1  0000  a9aa      ldai 0xaa
    2                   
`)
	})
	t.Run("list_opcodes_from_ast_multiple_bytes", func(t *testing.T) {
		assertListing(t, `.db 1, 2, 3, 4, 5, 6, 7, 8, 9, 10`,
			`Line  Loc   Code         Source
    1  0000  01020304  .db 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
    1  0004  05060708
    1  0008  090a    
    2                   
`)
	})
}
