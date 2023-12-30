package tests

import (
	"testing"

	"github.com/ddomurad/dd8/tools/asm/asm"
	"github.com/stretchr/testify/require"
)

func assertListing(t *testing.T, src string, expectedListing string) {
	src += "\n"
	ast, listing := asm.ParseSrcWithListing("test.asm", src)
	require.False(t, ast.Errors.HasErrors())
	// _, err := pkg.Assemble(ast, assemblers.OpcodeAssemblerW65C02S)
	// require.Nil(t, err)
	actuallListing := listing.String()
	require.Equal(t, expectedListing, actuallListing)
}

func Test_ListingGeneration(t *testing.T) {
	t.Run("list_source_lines", func(t *testing.T) {
		assertListing(t, `
      .def TEST := 1 
      ; this is a comment
      `, `Fl# Name
  0 test.asm

File Line  Loc   Code         Source 
   0    0                     .def TEST := 1 
   0    1                     ; this is a comment
`)
	})
}
