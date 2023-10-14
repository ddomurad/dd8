package tests

import (
	"reflect"
	"testing"

	"github.com/ddomurad/dd8/tools/asm/internal"
	"github.com/stretchr/testify/require"
)

func assertASTStatement(t *testing.T, expected, actual internal.ASTStatement) {
	expectedType := reflect.TypeOf(expected)
	require.Equal(t, expectedType, reflect.TypeOf(actual))
	switch et := expected.(type) {
	case internal.ASTInstruction:
		at := actual.(internal.ASTInstruction)
		require.Equal(t, et, at, "instruction")
	default:
		require.Failf(t, "unsupported instruction", "type: %v", expectedType)
	}
}

func assertAST(t *testing.T, expected, actual *internal.AST) {
	require.Len(t, actual.Statements, len(expected.Statements), "statements count")
	for i, expectedStatement := range expected.Statements {
		assertASTStatement(t, expectedStatement, actual.Statements[i])
	}
}

func TestThatCanParseSource(t *testing.T) {
	t.Run("simple_opcodes", func(tt *testing.T) {
		ast := internal.ParseSrc(`
      opcodea
      opcodeb
      opcodec
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTInstruction{OpCode: "opcodea", Operands: []any{}},
				internal.ASTInstruction{OpCode: "opcodeb", Operands: []any{}},
				internal.ASTInstruction{OpCode: "opcodec", Operands: []any{}},
			},
		}, ast)
	})

	t.Run("opcodes_with_single_num_param", func(tt *testing.T) {
		ast := internal.ParseSrc(`
      opcodea 10
      opcodeb 0x1_a
      opcodec 0b10_1_1
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTInstruction{OpCode: "opcodea", Operands: []any{internal.ASTNumericOperand(10)}},
				internal.ASTInstruction{OpCode: "opcodeb", Operands: []any{internal.ASTNumericOperand(0x1a)}},
				internal.ASTInstruction{OpCode: "opcodec", Operands: []any{internal.ASTNumericOperand(0xb)}},
			},
		}, ast)
	})

	t.Run("opcodes_with_single_param_and_reg", func(tt *testing.T) {
		ast := internal.ParseSrc(`
			     opcodea 10, x
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTInstruction{OpCode: "opcodea", Operands: []any{
					internal.ASTNumericOperand(10),
					internal.ASTStringOperand("x"),
				}},
			},
		}, ast)
	})
	return
	t.Run("eq_definition", func(t *testing.T) {

		ast := internal.ParseSrc(`
      .test test_name, 0x10
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTPreprocessInstruction{
					Instruction: "test",
					Parameters: []any{
						internal.ASTStringOperand("test_name"),
						internal.ASTNumericOperand(0x10),
					},
				},
			},
		}, ast)
	})
}
