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
	actuallListing := listing.FileListings["test.asm"].String(listing.FileListings)
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
	t.Run("list_for_template", func(t *testing.T) {
		assertListing(t, `.tmpl test(v1) {
  ldai v1
}

@test(0xaa)
@test(0xbb)`,
			`Line  Loc   Code         Source
    1                   .tmpl test(v1) {
    2                     ldai v1
    3                   }
    4                   
    5                   @test(v1=0xaa)
    5  0000  a9aa        ldai v1
    6                   @test(v1=0xbb)
    6  0002  a9bb        ldai v1
    7                   
`)
	})
	t.Run("list_repeat_single_line", func(t *testing.T) {
		assertListing(t, `.rep i,0,2 {
  ldai i
}`,
			`Line  Loc   Code         Source
    1                   @repeat i=0
    1  0000  a900        ldai i
    1                   @repeat i=1
    1  0002  a901    
    1                   @repeat i=2
    1  0004  a902    
    1                   @repeat i=0 to 2 {
    2                     ldai i
    3                   }
    4                   
`)
	})
}
