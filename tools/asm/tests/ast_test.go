package tests

import (
	"reflect"
	"testing"

	pkg "github.com/ddomurad/dd8/tools/asm/asm"
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

func exprOperand(l any, op string, r any) pkg.ASTOperand {
	return pkg.ASTOperand{
		Value: expr(l, op, r),
	}
}

func nameExpresion(n pkg.ASTName) pkg.ASTExpr {
	return pkg.ASTExpr{
		Left:      n,
		Right:     nil,
		Operation: "",
	}
}

func expr(l any, op string, r any) pkg.ASTExpr {
	return pkg.ASTExpr{
		Left:      l,
		Right:     r,
		Operation: op,
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
	// t.Run("apply_defaults", func(t *testing.T) {
	// 	ast := pkg.ParseSrc("", `
	//      .def TEST_VAL := 0x59 + 1
	//      .def OTHER_VAL := 100 - 1
	//      .def TEST_ORG := 0x100*2
	//      lda TEST_VAL,x
	//      .org TEST_ORG
	//      ldai OTHER_VAL
	//      `)
	// 	require.False(t, ast.Errors.HasErrors())
	// 	pkg.PreprocessDefinitions(ast)
	// 	pkg.PreprocessExpressions(ast)
	// 	require.False(t, ast.Errors.HasErrors())
	// 	assertAST(t, &pkg.AST{
	// 		Statements: []pkg.ASTStatement{
	// 			{
	// 				Type:     pkg.ASTStatementTypeInstruction,
	// 				OpCode:   "lda",
	// 				Operands: pkg.ASTOperands{numOperand(0x5a), regOperand("x")}},
	// 			{
	// 				Type:     pkg.ASTStatementTypeOrigin,
	// 				Operands: pkg.ASTOperands{numOperand(0x200)},
	// 			},
	// 			{
	// 				Type:     pkg.ASTStatementTypeInstruction,
	// 				OpCode:   "ldai",
	// 				Operands: pkg.ASTOperands{numOperand(99)}},
	// 		},
	// 	}, ast)
	// })
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
      inline_label: opcc 0x30
      label_1: label_2: opcd 0x40
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
				{
					Type: pkg.ASTStatementTypeLabel,
					Name: "inline_label",
				},
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcc",
					Operands: []pkg.ASTOperand{numOperand(0x30)},
				},
				{
					Type: pkg.ASTStatementTypeLabel,
					Name: "label_1",
				},
				{
					Type: pkg.ASTStatementTypeLabel,
					Name: "label_2",
				},
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcd",
					Operands: []pkg.ASTOperand{numOperand(0x40)},
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

	t.Run(".db_ml", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .db (
         0x10, "test data", 0xff
         0x11, "other data", 0xff
      )
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{
						numOperand(0x10), strOperand("test data"), numOperand(0xff),
						numOperand(0x11), strOperand("other data"), numOperand(0xff)},
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

	t.Run(".dw_ml", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .dw ( 
      0x10, "test data", 0xff 
      "test", 0x00
      )
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypeDataWord,
					Operands: pkg.ASTOperands{
						numOperand(0x10), strOperand("test data"), numOperand(0xff),
						strOperand("test"), numOperand(0x00),
					},
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

	t.Run(".byte", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .byte
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeSkipBytes,
					Operands: pkg.ASTOperands{numOperand(1)},
				},
			},
		}, ast)
	})

	t.Run(".byte_10", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .byte 10
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeSkipBytes,
					Operands: pkg.ASTOperands{numOperand(10)},
				},
			},
		}, ast)
	})

	t.Run(".word", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .word
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeSkipWords,
					Operands: pkg.ASTOperands{numOperand(1)},
				},
			},
		}, ast)
	})

	t.Run(".word_10", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .word 10
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeSkipWords,
					Operands: pkg.ASTOperands{numOperand(10)},
				},
			},
		}, ast)
	})

	t.Run("num_expr_add", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .db 0x10 + 0x20
      .db 0x10 - 0x20
      .db NAME_1 * NAME_2
      .db 0x10 / TEST_NAME
      .db 0x10 << 0x20
      .db 0x10 >> 0x20
      .db 1 + 2 * 3
      .db 1 * 2 - 3
      .db ~1
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(pkg.ASTNumber(0x10), "+", pkg.ASTNumber(0x20))},
				},
				{
					Type:     pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(pkg.ASTNumber(0x10), "-", pkg.ASTNumber(0x20))},
				},
				{
					Type:     pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(nameExpresion("NAME_1"), "*", nameExpresion("NAME_2"))},
				},
				{
					Type:     pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(pkg.ASTNumber(0x10), "/", nameExpresion("TEST_NAME"))},
				},
				{
					Type:     pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(pkg.ASTNumber(0x10), "<<", pkg.ASTNumber(0x20))},
				},
				{
					Type:     pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(pkg.ASTNumber(0x10), ">>", pkg.ASTNumber(0x20))},
				},
				{
					Type: pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(
						pkg.ASTNumber(1), "+", expr(pkg.ASTNumber(2), "*", pkg.ASTNumber(3)))},
				},
				{
					Type: pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(
						expr(pkg.ASTNumber(1), "*", pkg.ASTNumber(2)), "-", pkg.ASTNumber(3))},
				},
				{
					Type: pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{exprOperand(
						nil, "~", pkg.ASTNumber(1))},
				},
			},
		}, ast)
	})
}
