package tests

import (
	"testing"

	pkg "github.com/ddomurad/dd8/tools/asm/pkg"
	"github.com/ddomurad/dd8/tools/asm/pkg/assemblers"
	"github.com/ddomurad/dd8/tools/asm/pkg/output"
	"github.com/stretchr/testify/require"
)

func TestThatCanParseSimpleAssemlyFile(t *testing.T) {
	srcReader := pkg.NewFileSourceReader("./res/simple_tests/")
	testCases := map[string]string{
		"test.asm":      "test.out",
		"test2.asm":     "test2.out",
		"all_inst.asm":  "all_inst.out",
		"example.asm":   "example.out",
		"data_test.asm": "data_test.out",
	}

	for srcName, expectedOutputFile := range testCases {
		t.Run("test_file_"+srcName, func(t *testing.T) {
			expectedOutput, err := srcReader.ReadSourceFile(expectedOutputFile)
			require.NoError(t, err, "read expected output error")
			bcode, err := pkg.AssembleSrc(srcName, srcReader, assemblers.OpcodeAssemblerW65C02S)
			require.NoError(t, err, "assemble src")
			hex := output.GetIntelHEX(bcode)
			require.Equal(t, expectedOutput, string(hex))
		})
	}
}

func TestThatCanParseComplexAssemlyFile(t *testing.T) {
	srcReader := pkg.NewFileSourceReader("./res/crazy_test/")
	testCases := map[string]string{
		"main.asm": "main.out",
	}

	for srcName, expectedOutputFile := range testCases {
		t.Run("test_file_"+srcName, func(t *testing.T) {
			expectedOutput, err := srcReader.ReadSourceFile(expectedOutputFile)
			require.NoError(t, err, "read expected output error")
			bcode, err := pkg.AssembleSrc(srcName, srcReader, assemblers.OpcodeAssemblerW65C02S)
			require.NoError(t, err, "assemble src")
			hex := output.GetIntelHEX(bcode)
			require.Equal(t, expectedOutput, string(hex))
		})
	}
}

func TestFullSpecRegression(t *testing.T) {
	srcReader := pkg.NewFileSourceReader("./res/full_spec/")
	testCases := map[string]string{
		"main.asm": "main.out",
	}

	for srcName, expectedOutputFile := range testCases {
		t.Run("test_file_"+srcName, func(t *testing.T) {
			expectedOutput, err := srcReader.ReadSourceFile(expectedOutputFile)
			require.NoError(t, err, "read expected output error")
			bcode, err := pkg.AssembleSrc(srcName, srcReader, assemblers.OpcodeAssemblerW65C02S)
			require.NoError(t, err, "assemble src")
			hex := output.GetIntelHEX(bcode)
			require.Equal(t, expectedOutput, string(hex))
		})
	}
}
