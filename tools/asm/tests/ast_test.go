package tests

import (
	"reflect"
	"testing"

	pkg "github.com/ddomurad/dd8/tools/asm/pkg"
	"github.com/stretchr/testify/require"
)

func assertASTStatement(t *testing.T, expected, actual pkg.ASTStatement) {
	expectedType := reflect.TypeOf(expected)
	require.Equal(t, expectedType, reflect.TypeOf(actual))

	actual.SrcPointer = pkg.SrcPointer{}
	expected.SrcPointer = pkg.SrcPointer{}
	require.Equal(t, expected, actual, "instruction")
}

func assertAST(t *testing.T, expected, actual *pkg.AST) {
	require.Len(t, actual.Statements, len(expected.Statements), "statements count")
	for i, expectedStatement := range expected.Statements {
		assertASTStatement(t, expectedStatement, actual.Statements[i])
	}
}

func numOperand(v pkg.ASTNumber) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v}
}

func strOperand(v pkg.ASTString) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v}
}

func nameOperand(v pkg.ASTName) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v}
}

func regOperand(v pkg.ASTRegister) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v}
}

func pnumOperand(v pkg.ASTNumber) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v, Indirect: true}
}

func pstrOperand(v pkg.ASTString) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v, Indirect: true}
}

func pnameOperand(v pkg.ASTName) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v, Indirect: true}
}

func pregOperand(v pkg.ASTRegister) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v, Indirect: true}
}

func TestPreprocesor(t *testing.T) {
	t.Run("apply_defaults", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def TEST_VAL := 0x59
      .def OTHER_VAL := 100
      .def TEST_ORG := 0x100
      lda TEST_VAL,x 
      .org TEST_ORG
      ldai OTHER_VAL
      `)
		require.False(t, ast.Errors.HasErrors())
		pkg.PreprocessDefinitions(ast)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "lda",
					Operands: pkg.ASTOperands{numOperand(0x59), regOperand("x")}},
				{
					Type:     pkg.ASTStatementTypeOrigin,
					Operands: pkg.ASTOperands{numOperand(0x100)},
				},
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "ldai",
					Operands: pkg.ASTOperands{numOperand(100)}},
			},
		}, ast)
	})
}

func TestThatCanParseSource(t *testing.T) {
	t.Run("simple_opcodes", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      opcodea
      opcodeb
      opcodec
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{Type: pkg.ASTStatementTypeInstruction, OpCode: "opcodea", Operands: []pkg.ASTOperand{}},
				{Type: pkg.ASTStatementTypeInstruction, OpCode: "opcodeb", Operands: []pkg.ASTOperand{}},
				{Type: pkg.ASTStatementTypeInstruction, OpCode: "opcodec", Operands: []pkg.ASTOperand{}},
			},
		}, ast)
	})

	t.Run("opcodes_with_single_num_param", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      opcodea 10
      opcodeb 0x1_a
      opcodec 0b10_1_1
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{Type: pkg.ASTStatementTypeInstruction, OpCode: "opcodea", Operands: []pkg.ASTOperand{numOperand(10)}},
				{Type: pkg.ASTStatementTypeInstruction, OpCode: "opcodeb", Operands: []pkg.ASTOperand{numOperand(0x1a)}},
				{Type: pkg.ASTStatementTypeInstruction, OpCode: "opcodec", Operands: []pkg.ASTOperand{numOperand(0xb)}},
			},
		}, ast)
	})

	t.Run("opcodes_with_single_param_and_reg", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
			     opcodea 10, x
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcodea",
					Operands: []pkg.ASTOperand{numOperand(10), regOperand("x")}},
			},
		}, ast)
	})

	t.Run("opcodes_with_agrs_in_par", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
			     opcodea [10, x]
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcodea",
					Operands: []pkg.ASTOperand{pnumOperand(10), pregOperand("x")}},
			},
		}, ast)
	})

	t.Run("opcodes_with_agrs_in_first_par", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
			     opcodea [10], x
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcodea",
					Operands: []pkg.ASTOperand{pnumOperand(10), regOperand("x")}},
			},
		}, ast)
	})

	t.Run("opcodes_with_one_par", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
			     opcodea [10]
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:   pkg.ASTStatementTypeInstruction,
					OpCode: "opcodea", Operands: []pkg.ASTOperand{
						pnumOperand(10),
					}},
			},
		}, ast)
	})

	t.Run("opcodes_with_three_operands", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
			     opcodea 3, 10, 35 
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:   pkg.ASTStatementTypeInstruction,
					OpCode: "opcodea", Operands: []pkg.ASTOperand{
						numOperand(3),
						numOperand(10),
						numOperand(35),
					}},
			},
		}, ast)
	})

	t.Run("label", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
        opca 0x10
      test_label:
        opcb 0x20
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opca",
					Operands: []pkg.ASTOperand{numOperand(0x10)},
				},
				{
					Type: pkg.ASTStatementTypeLabel,
					Name: "test_label",
				},
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcb",
					Operands: []pkg.ASTOperand{numOperand(0x20)},
				},
			},
		}, ast)
	})

	t.Run(".db", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .db 0x10, "test data", 0xff
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{
						numOperand(0x10), strOperand("test data"), numOperand(0xff)},
				},
			},
		}, ast)
	})

	t.Run(".dw", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .dw 0x10, "test data", 0xff
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypeDataWord,
					Operands: pkg.ASTOperands{
						numOperand(0x10), strOperand("test data"), numOperand(0xff)},
				},
			},
		}, ast)
	})

	t.Run(".org", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .org 0x10
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeOrigin,
					Operands: pkg.ASTOperands{numOperand(0x10)},
				},
			},
		}, ast)
	})

	t.Run("prep_ol_def", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def TEST := 10
      .def TEST-100 := 0b1001_1100
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST",
					Operands: pkg.ASTOperands{numOperand(10)},
				},
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST-100",
					Operands: pkg.ASTOperands{numOperand(0b1001_1100)},
				},
			},
		}, ast)
	})

	t.Run("prep_ml_def", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def (
        TEST := 10
        TEST-100 := 0b1001_1100
      )
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST",
					Operands: pkg.ASTOperands{numOperand(10)},
				},
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST-100",
					Operands: pkg.ASTOperands{numOperand(0b1001_1100)},
				},
			},
		}, ast)
	})
}
