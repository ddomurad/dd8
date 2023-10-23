package tests

import (
	"testing"
)

func TestASTErrorHandling(t *testing.T) {
	// t.Run("simple_sytax_err", func(t *testing.T) {
	// 	ast := internal.ParseSrc("test.asm", "!!")
	// 	require.True(t, ast.Errors.HasErrors())
	// 	require.Equal(t, "[syntax_err] test.asm:1:0 token recognition error at: '!'", ast.Errors.Errors[0].Error())
	// 	require.Equal(t, "[syntax_err] test.asm:1:1 token recognition error at: '!'", ast.Errors.Errors[1].Error())
	// })
	//
	// t.Run("program_error", func(t *testing.T) {
	// 	ast := internal.ParseSrc("test.asm", `
	//      .def TEST :x = 2
	//      `)
	// 	require.Len(t, ast.Errors.Errors, 3)
	// 	require.Equal(t, "[syntax_err] test.asm:2:16 mismatched input ':' expecting ':='", ast.Errors.Errors[0].Error())
	// 	require.Equal(t, "[syntax_err] test.asm:2:19 token recognition error at: '='", ast.Errors.Errors[1].Error())
	// 	require.Equal(t, "[struc_err] test.asm:2 preprocessor argument missing", ast.Errors.Errors[2].Error())
	// })
	// t.Run("test", func(t *testing.T) {
	// 	ast := internal.ParseSrc("test.asm", `
	//      ldai 0x10,x,
	//      ldai 0x10,x,
	//      `)
	// 	require.Len(t, ast.Errors.Errors, 3)
	// 	require.Equal(t, "[syntax_err] test.asm:2:18 extraneous input '\\n' expecting {REG, STR, NAME, HEX_NUM, BIN_NUM, DEC_NUM}", ast.Errors.Errors[0].Error())
	// 	require.Equal(t, "[syntax_err] test.asm:3:11 mismatched input '0x10' expecting EOL", ast.Errors.Errors[1].Error())
	// 	require.Equal(t, "[struc_err] test.asm:2 unexpected statement structure", ast.Errors.Errors[2].Error())
	// })
}
