package tests

import (
	"testing"

	"github.com/ddomurad/dd8/tools/asm/internal"
	"github.com/ddomurad/dd8/tools/asm/internal/assemblers"
	"github.com/ddomurad/dd8/tools/asm/internal/output"
	"github.com/stretchr/testify/require"
)

func TestThatCanParseSimpleAssemlyFile(t *testing.T) {
	srcReader := internal.NewFileSourceReader("./res/simple_tests/")
	testCases := map[string]string{
		"test.asm":     "test.out",
		"test2.asm":    "test2.out",
		"all_inst.asm": "all_inst.out",
		"example.asm":  "example.out",
	}

	for srcName, expectedOutputFile := range testCases {
		t.Run("test_file_"+srcName, func(t *testing.T) {
			expectedOutput, err := srcReader.ReadSourceFile(expectedOutputFile)
			require.NoError(t, err, "read expected output error")
			bcode, err := internal.AssembleSrc(srcName, srcReader, assemblers.OpcodeAssemblerW65C02S)
			require.NoError(t, err, "assemble src")
			hex := output.GetIntelHEX(bcode)
			require.Equal(t, expectedOutput, string(hex))
		})
	}
}

func TestThatCanParseComplexAssemlyFile(t *testing.T) {
	srcReader := internal.NewFileSourceReader("./res/crazy_test/")
	testCases := map[string]string{
		"main.asm": "main.out",
	}

	for srcName, expectedOutputFile := range testCases {
		t.Run("test_file_"+srcName, func(t *testing.T) {
			expectedOutput, err := srcReader.ReadSourceFile(expectedOutputFile)
			require.NoError(t, err, "read expected output error")
			bcode, err := internal.AssembleSrc(srcName, srcReader, assemblers.OpcodeAssemblerW65C02S)
			require.NoError(t, err, "assemble src")
			hex := output.GetIntelHEX(bcode)
			require.Equal(t, expectedOutput, string(hex))
		})
	}
}
