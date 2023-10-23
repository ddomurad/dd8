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

	actual.SrcPointer = internal.SrcPointer{}
	expected.SrcPointer = internal.SrcPointer{}
	require.Equal(t, expected, actual, "instruction")
}

func assertAST(t *testing.T, expected, actual *internal.AST) {
	require.Len(t, actual.Statements, len(expected.Statements), "statements count")
	for i, expectedStatement := range expected.Statements {
		assertASTStatement(t, expectedStatement, actual.Statements[i])
	}
}

func operand(v any) internal.ASTOperand {
	switch tv := v.(type) {
	case int:
		return internal.ASTOperand{Value: internal.ASTNumber(tv)}
	case string:
		if len(tv) == 1 {
			return internal.ASTOperand{Value: internal.ASTRegister(tv)}
		}
		return internal.ASTOperand{Value: internal.ASTName(tv)}
	}

	panic("unsuported operand in test")
}

func poperand(v any) internal.ASTOperand {
	op := operand(v)
	op.Indirect = true
	return op
}

func TestPreprocesor(t *testing.T) {
	t.Run("apply_defaults", func(t *testing.T) {
		ast := internal.ParseSrc("", `
      .def TEST_VAL := 0x59
      .def OTHER_VAL := 100
      .def TEST_ORG := 0x100
      lda TEST_VAL,x 
      .org TEST_ORG
      ldai OTHER_VAL
      `)
		require.False(t, ast.Errors.HasErrors())
		internal.PreprocessDefinitions(ast)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypeInstruction,
					OpCode:   "lda",
					Operands: internal.ASTOperands{operand(0x59), operand("x")}},
				{
					Type:     internal.ASTStatementTypeOrigin,
					Operands: internal.ASTOperands{operand(0x100)},
				},
				{
					Type:     internal.ASTStatementTypeInstruction,
					OpCode:   "ldai",
					Operands: internal.ASTOperands{operand(100)}},
			},
		}, ast)
	})
}

func TestThatCanParseSource(t *testing.T) {
	t.Run("simple_opcodes", func(t *testing.T) {
		ast := internal.ParseSrc("", `
      opcodea
      opcodeb
      opcodec
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{Type: internal.ASTStatementTypeInstruction, OpCode: "opcodea", Operands: []internal.ASTOperand{}},
				{Type: internal.ASTStatementTypeInstruction, OpCode: "opcodeb", Operands: []internal.ASTOperand{}},
				{Type: internal.ASTStatementTypeInstruction, OpCode: "opcodec", Operands: []internal.ASTOperand{}},
			},
		}, ast)
	})

	t.Run("opcodes_with_single_num_param", func(t *testing.T) {
		ast := internal.ParseSrc("", `
      opcodea 10
      opcodeb 0x1_a
      opcodec 0b10_1_1
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{Type: internal.ASTStatementTypeInstruction, OpCode: "opcodea", Operands: []internal.ASTOperand{operand(10)}},
				{Type: internal.ASTStatementTypeInstruction, OpCode: "opcodeb", Operands: []internal.ASTOperand{operand(0x1a)}},
				{Type: internal.ASTStatementTypeInstruction, OpCode: "opcodec", Operands: []internal.ASTOperand{operand(0xb)}},
			},
		}, ast)
	})

	t.Run("opcodes_with_single_param_and_reg", func(t *testing.T) {
		ast := internal.ParseSrc("", `
			     opcodea 10, x
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypeInstruction,
					OpCode:   "opcodea",
					Operands: []internal.ASTOperand{operand(10), operand("x")}},
			},
		}, ast)
	})

	t.Run("opcodes_with_agrs_in_par", func(t *testing.T) {
		ast := internal.ParseSrc("", `
			     opcodea [10, x]
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypeInstruction,
					OpCode:   "opcodea",
					Operands: []internal.ASTOperand{poperand(10), poperand("x")}},
			},
		}, ast)
	})

	t.Run("opcodes_with_agrs_in_first_par", func(t *testing.T) {
		ast := internal.ParseSrc("", `
			     opcodea [10], x
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypeInstruction,
					OpCode:   "opcodea",
					Operands: []internal.ASTOperand{poperand(10), operand("x")}},
			},
		}, ast)
	})

	t.Run("opcodes_with_one_par", func(t *testing.T) {
		ast := internal.ParseSrc("", `
			     opcodea [10]
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:   internal.ASTStatementTypeInstruction,
					OpCode: "opcodea", Operands: []internal.ASTOperand{
						poperand(10),
					}},
			},
		}, ast)
	})

	t.Run("label", func(t *testing.T) {
		ast := internal.ParseSrc("", `
        opca 0x10
      test_label:
        opcb 0x20
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypeInstruction,
					OpCode:   "opca",
					Operands: []internal.ASTOperand{operand(0x10)},
				},
				{
					Type: internal.ASTStatementTypeLabel,
					Name: "test_label",
				},
				{
					Type:     internal.ASTStatementTypeInstruction,
					OpCode:   "opcb",
					Operands: []internal.ASTOperand{operand(0x20)},
				},
			},
		}, ast)
	})

	t.Run(".org", func(t *testing.T) {
		ast := internal.ParseSrc("", `
      .org 0x10
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypeOrigin,
					Operands: internal.ASTOperands{operand(0x10)},
				},
			},
		}, ast)
	})

	t.Run("prep_ol_def", func(t *testing.T) {
		ast := internal.ParseSrc("", `
      .def TEST := 10
      .def TEST-100 := 0b1001_1100
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypePrepDefine,
					Name:     "TEST",
					Operands: internal.ASTOperands{operand(10)},
				},
				{
					Type:     internal.ASTStatementTypePrepDefine,
					Name:     "TEST-100",
					Operands: internal.ASTOperands{operand(0b1001_1100)},
				},
			},
		}, ast)
	})

	t.Run("prep_ml_def", func(t *testing.T) {
		ast := internal.ParseSrc("", `
      .def (
        TEST := 10
        TEST-100 := 0b1001_1100
      )
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &internal.AST{
			Statements: []internal.ASTStatement{
				{
					Type:     internal.ASTStatementTypePrepDefine,
					Name:     "TEST",
					Operands: internal.ASTOperands{operand(10)},
				},
				{
					Type:     internal.ASTStatementTypePrepDefine,
					Name:     "TEST-100",
					Operands: internal.ASTOperands{operand(0b1001_1100)},
				},
			},
		}, ast)
	})
}
