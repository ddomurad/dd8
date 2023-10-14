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
	case internal.ASTLabel:
		at := actual.(internal.ASTLabel)
		require.Equal(t, et, at, "label")
	case internal.ASTOrigin:
		at := actual.(internal.ASTOrigin)
		require.Equal(t, et, at, "origin")
	case internal.ASTPrepDefine:
		at := actual.(internal.ASTPrepDefine)
		require.Equal(t, et, at, "define")
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

func TestPreprocesor(t *testing.T) {
	t.Run("apply_defaults", func(t *testing.T) {
		ast := internal.ParseSrc(`
      .def TEST_VAL := 0x59
      .def OTHER_VAL := 100
      .def TEST_ORG := 0x100
      lda TEST_VAL,x 
      .org TEST_ORG
      ldai OTHER_VAL
      `)
		err := internal.PreprocessDefinitions(ast)
		require.NoError(t, err)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTInstruction{OpCode: "lda", Operands: []any{
					internal.ASTNumericOperand(0x59),
					internal.ASTNameOperand("x"),
				}},
				internal.ASTOrigin{Address: internal.ASTNumericOperand(0x100)},
				internal.ASTInstruction{OpCode: "ldai", Operands: []any{internal.ASTNumericOperand(100)}},
			},
		}, ast)
	})
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
					internal.ASTNameOperand("x"),
				}},
			},
		}, ast)
	})

	t.Run("label", func(t *testing.T) {
		ast := internal.ParseSrc(`
        opca 0x10
      test_label:
        opcb 0x20
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTInstruction{OpCode: "opca", Operands: []any{internal.ASTNumericOperand(0x10)}},
				internal.ASTLabel{Name: "test_label"},
				internal.ASTInstruction{OpCode: "opcb", Operands: []any{internal.ASTNumericOperand(0x20)}},
			},
		}, ast)
	})

	t.Run(".org", func(t *testing.T) {
		ast := internal.ParseSrc(`
      .org 0x10
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTOrigin{Address: internal.ASTNumericOperand(0x10)},
			},
		}, ast)
	})

	t.Run("prep_ol_def", func(t *testing.T) {
		ast := internal.ParseSrc(`
      .def TEST := 10
      .def TEST-100 := 0b1001_1100
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTPrepDefine{Name: "TEST", Value: internal.ASTNumericOperand(10)},
				internal.ASTPrepDefine{Name: "TEST-100", Value: internal.ASTNumericOperand(0b1001_1100)},
			},
		}, ast)
	})

	t.Run("prep_ml_def", func(t *testing.T) {
		ast := internal.ParseSrc(`
      .def (
        TEST := 10
        TEST-100 := 0b1001_1100
      )
      `)
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				internal.ASTPrepDefine{Name: "TEST", Value: internal.ASTNumericOperand(10)},
				internal.ASTPrepDefine{Name: "TEST-100", Value: internal.ASTNumericOperand(0b1001_1100)},
			},
		}, ast)
	})
}
