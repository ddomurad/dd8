package tests

import (
	"testing"

	"github.com/ddomurad/dd8/tools/asm/internal/parser"
	"github.com/stretchr/testify/assert"
)

func assertOpcoedeEntry(t *testing.T, entry parser.SourceEntry, opCode string) {
	assert.Equal(t, parser.SourceEntryTypeOpCode, entry.Type)
	assert.Equal(t, opCode, entry.OpCode)
}

func assertOpcoedeNumParEntry(t *testing.T, entry parser.SourceEntry, opCode string, numParam int64) {
	assert.Equal(t, parser.SourceEntryTypeOpCode, entry.Type)
	assert.Equal(t, opCode, entry.OpCode)
	assert.Equal(t, numParam, entry.NumParam)
}

func TestThatCanParseSource(t *testing.T) {
	t.Run("simple_opcodes", func(tt *testing.T) {
		src := parser.ParseSource(parser.RawSourceFile{
			Source: `
      opcodea
      opcodeb
      opcodec
      `,
			Uri: "test.asm",
		})
		assertOpcoedeEntry(t, src.Entries[0], "opcodea")
		assertOpcoedeEntry(t, src.Entries[1], "opcodeb")
		assertOpcoedeEntry(t, src.Entries[2], "opcodec")
	})

	t.Run("opcodes_with_num_param", func(tt *testing.T) {
		src := parser.ParseSource(parser.RawSourceFile{
			Source: `
      opcodea 10
      opcodeb 0x1a
      opcodec 0b1010
      `,
			Uri: "test.asm",
		})
		assertOpcoedeNumParEntry(t, src.Entries[0], "opcodea", 10)
		assertOpcoedeNumParEntry(t, src.Entries[1], "opcodeb", 0x1a)
		assertOpcoedeNumParEntry(t, src.Entries[2], "opcodec", 0b1010)
	})
}

// func Test_that_file_can_be_parssed(t *testing.T) {
// 	entires, err := os.ReadDir("./")
// 	require.NoError(t, err)
//
// 	srcReader := internal.NewFileSourceReader("./")
//
// 	for _, entry := range entires {
// 		if path.Ext(entry.Name()) != ".asm" {
// 			continue
// 		}
//
// 		t.Run("test_"+entry.Name(), func(tt *testing.T) {
// 			tree := parser.ParseSource(parser.RawSourceFile{})
// 		})
// 	}
// }
