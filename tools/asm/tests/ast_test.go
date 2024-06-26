package tests

import (
	"reflect"
	"testing"

	pkg "github.com/ddomurad/dd8/tools/asm/asm"
	"github.com/stretchr/testify/require"
)

func clearSrcPointer(statements []pkg.ASTStatement) []pkg.ASTStatement {
	for i := range statements {
		statements[i].SrcPointer = pkg.SrcPointer{}

		for j, op := range statements[i].Operands {
			blockOperand, ok := op.Value.(pkg.ASTBlock)
			if ok {
				blockOperand.Statements = clearSrcPointer(blockOperand.Statements)
				statements[i].Operands[j].Value = blockOperand
			}
		}
	}

	return statements
}

func assertASTStatement(t *testing.T, expected, actual pkg.ASTStatement) {
	expectedType := reflect.TypeOf(expected)
	require.Equal(t, expectedType, reflect.TypeOf(actual))
	require.Equal(t, expected, actual, "instruction")
}

func assertAST(t *testing.T, expected, actual *pkg.AST) {
	require.Len(t, actual.Statements, len(expected.Statements), "statements count")

	actual.Statements = clearSrcPointer(actual.Statements)
	expected.Statements = clearSrcPointer(expected.Statements)

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

func nameListOperand(v ...pkg.ASTName) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v}
}

func regOperand(v pkg.ASTRegister) pkg.ASTOperand {
	return pkg.ASTOperand{Value: v}
}

func numArrayOperand(v ...pkg.ASTNumber) pkg.ASTOperand {
	arr := make(pkg.ASTArray, len(v))
	for i, n := range v {
		arr[i] = n
	}
	return pkg.ASTOperand{Value: arr}
}

func strArrayOperand(v ...pkg.ASTString) pkg.ASTOperand {
	arr := make(pkg.ASTArray, len(v))
	for i, n := range v {
		arr[i] = n
	}
	return pkg.ASTOperand{Value: arr}
}

func anyArrayOperand(v ...any) pkg.ASTOperand {
	arr := make(pkg.ASTArray, len(v))
	for i, n := range v {
		arr[i] = n
	}
	return pkg.ASTOperand{Value: arr}
}

func blockOperand(v ...pkg.ASTStatement) pkg.ASTOperand {
	return pkg.ASTOperand{Value: pkg.ASTBlock{Statements: v}}
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

	t.Run("prep_sl_def", func(t *testing.T) {
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

	t.Run("prep_sl_def_single_caracter_name", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def I := 10
      .def R := 12
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "I",
					Operands: pkg.ASTOperands{numOperand(10)},
				},
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "R",
					Operands: pkg.ASTOperands{numOperand(12)},
				},
			},
		}, ast)
	})

	t.Run("prep_sl_def_a_is_forbidden", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def a := 10
      `)
		require.True(t, ast.Errors.HasErrors())
	})

	t.Run("prep_sl_def_x_is_forbidden", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def x := 10
      `)
		require.True(t, ast.Errors.HasErrors())
	})

	t.Run("prep_sl_def_y_is_forbidden", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def y := 10
      `)
		require.True(t, ast.Errors.HasErrors())
	})

	t.Run("prep_sl_def_Z_is_forbidden", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def Z := 10
      `)
		require.True(t, ast.Errors.HasErrors())
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

	t.Run("prep_def_arrays", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .def TEST := <10, 15>
      .def (
        TEST2 := <1, 2, 3>
        TEST3 := <"hello", "world", "!">
        TEST4 := <1, "test">
        TEST5 := <1*2, TEST2*TEST3>
      )
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST",
					Operands: pkg.ASTOperands{numArrayOperand(10, 15)},
				},
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST2",
					Operands: pkg.ASTOperands{numArrayOperand(1, 2, 3)},
				},
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST3",
					Operands: pkg.ASTOperands{strArrayOperand("hello", "world", "!")},
				},
				{
					Type:     pkg.ASTStatementTypePrepDefine,
					Name:     "TEST4",
					Operands: pkg.ASTOperands{anyArrayOperand(pkg.ASTNumber(1), pkg.ASTString("test"))},
				},
				{
					Type: pkg.ASTStatementTypePrepDefine,
					Name: "TEST5",
					Operands: pkg.ASTOperands{anyArrayOperand(
						expr(pkg.ASTNumber(1), "*", pkg.ASTNumber(2)),
						expr(nameExpresion("TEST2"), "*", nameExpresion("TEST3")),
					)},
				},
			},
		}, ast)
	})

	t.Run("prep_def_lenght_of_array", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
        opcode TEST_ARRAY.len
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcode",
					Operands: []pkg.ASTOperand{exprOperand(nameExpresion("TEST_ARRAY"), ".len", nil)},
				},
			},
		}, ast)
	})

	t.Run("prep_def_arrays_usage", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      opcode1 TEST[0]
      opcode2 TEST[1]
      opcode3 TEST[OTER[2*ELSE]]*TEST[2]
      `)

		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcode1",
					Operands: []pkg.ASTOperand{exprOperand(nameExpresion("TEST"), "[]", pkg.ASTNumber(0))},
				},
				{
					Type:     pkg.ASTStatementTypeInstruction,
					OpCode:   "opcode2",
					Operands: []pkg.ASTOperand{exprOperand(nameExpresion("TEST"), "[]", pkg.ASTNumber(1))},
				},
				{
					Type:   pkg.ASTStatementTypeInstruction,
					OpCode: "opcode3",
					Operands: []pkg.ASTOperand{
						exprOperand(
							expr(nameExpresion("TEST"), "[]", expr(nameExpresion("OTER"), "[]", expr(pkg.ASTNumber(2), "*", nameExpresion("ELSE")))),
							"*",
							expr(nameExpresion("TEST"), "[]", pkg.ASTNumber(2)),
						),
					},
				},
			},
		}, ast)
	})

	t.Run("prep_ml_def_with_empty_lines", func(t *testing.T) {
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
      .db NAME_3
      .db 0x10 / TEST_NAME
      .db 0x10 << 0x20
      .db 0x10 >> 0x20
      .db 1 + 2 * 3
      .db 1 * 2 - 3
      .db ~1
      .db 'd'
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
					Operands: pkg.ASTOperands{exprOperand(pkg.ASTName("NAME_3"), "", nil)},
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
				{
					Type: pkg.ASTStatementTypeDataByte,
					Operands: pkg.ASTOperands{
						numOperand(100),
					},
				},
			},
		}, ast)
	})

	t.Run(".tmpl_no_params", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .tmpl test_template () {
        opcode v1, v2
        opcode v3
      }
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypePrepTemplateDef,
					Name: "test_template",
					Operands: []pkg.ASTOperand{
						nameListOperand([]pkg.ASTName{}...),
						blockOperand(
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("v1"), "", nil),
									exprOperand(pkg.ASTName("v2"), "", nil),
								},
							},
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("v3"), "", nil),
								},
							},
						),
					},
				},
			},
		}, ast)
	})

	t.Run(".tmpl_simple_template", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .tmpl test_template (v1, v2, v3) {
        opcode v1, v2
        opcode v3
      }
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypePrepTemplateDef,
					Name: "test_template",
					Operands: []pkg.ASTOperand{
						nameListOperand("v1", "v2", "v3"),
						blockOperand(
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("v1"), "", nil),
									exprOperand(pkg.ASTName("v2"), "", nil),
								},
							},
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("v3"), "", nil),
								},
							},
						),
					},
				},
			},
		}, ast)
	})

	t.Run(".tmpl_use_template", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      @test_template(0x01, x, "test", VAR, TEST*2)
    `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypePrepTemplateUse,
					Name: "test_template",
					Operands: []pkg.ASTOperand{
						numOperand(0x01),
						regOperand("x"),
						strOperand("test"),
						exprOperand(pkg.ASTName("VAR"), "", nil),
						exprOperand(nameExpresion("TEST"), "*", pkg.ASTNumber(2)),
					},
				},
			},
		}, ast)
	})

	t.Run(".tmpl_use_template_no_params", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      @test_template()
    `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type:     pkg.ASTStatementTypePrepTemplateUse,
					Name:     "test_template",
					Operands: []pkg.ASTOperand{},
				},
			},
		}, ast)
	})

	t.Run(".tmpl_in_tmpl", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .tmpl test_template (v1, v2, v3) {
        opcode v1, v2
        opcode v3
        .tmpl nested_template (v4, v5) {
          opcode v10, v11
        }
      }
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypePrepTemplateDef,
					Name: "test_template",
					Operands: []pkg.ASTOperand{
						nameListOperand("v1", "v2", "v3"),
						blockOperand(
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("v1"), "", nil),
									exprOperand(pkg.ASTName("v2"), "", nil),
								},
							},
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("v3"), "", nil),
								},
							},
							pkg.ASTStatement{
								Type: pkg.ASTStatementTypePrepTemplateDef,
								Name: "nested_template",
								Operands: []pkg.ASTOperand{
									nameListOperand("v4", "v5"),
									blockOperand(
										pkg.ASTStatement{
											Type:   pkg.ASTStatementTypeInstruction,
											OpCode: "opcode",
											Operands: pkg.ASTOperands{
												exprOperand(pkg.ASTName("v10"), "", nil),
												exprOperand(pkg.ASTName("v11"), "", nil),
											},
										},
									),
								},
							},
						),
					},
				},
			},
		}, ast)
	})

	t.Run(".rep_test", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      .rep i, 0, 10 {
        opcode i
      }
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypePrepRepeat,
					Name: "",
					Operands: pkg.ASTOperands{
						nameOperand("i"),
						numOperand(0),
						numOperand(10),
						blockOperand(
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("i"), "", nil),
								},
							},
						),
					},
				},
			},
		}, ast)
	})

	t.Run("code_block_test", func(t *testing.T) {
		ast := pkg.ParseSrc("", `
      {
        opcode i
        opcode j
        {
          test k 
        }
      }
      `)
		require.False(t, ast.Errors.HasErrors())
		assertAST(t, &pkg.AST{
			Statements: []pkg.ASTStatement{
				{
					Type: pkg.ASTStatementTypePrepBlock,
					Name: "",
					Operands: pkg.ASTOperands{
						blockOperand(
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("i"), "", nil),
								},
							},
							pkg.ASTStatement{
								Type:   pkg.ASTStatementTypeInstruction,
								OpCode: "opcode",
								Operands: pkg.ASTOperands{
									exprOperand(pkg.ASTName("j"), "", nil),
								},
							},
							pkg.ASTStatement{
								Type: pkg.ASTStatementTypePrepBlock,
								Operands: pkg.ASTOperands{
									blockOperand(
										pkg.ASTStatement{
											Type:   pkg.ASTStatementTypeInstruction,
											OpCode: "test",
											Operands: pkg.ASTOperands{
												exprOperand(pkg.ASTName("k"), "", nil),
											},
										}),
								},
							},
						),
					},
				},
			},
		}, ast)
	})
}
