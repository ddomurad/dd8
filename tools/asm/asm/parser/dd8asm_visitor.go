// Code generated from DD8ASM.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DD8ASM

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DD8ASMParser.
type DD8ASMVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DD8ASMParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#prep_instruction.
	VisitPrep_instruction(ctx *Prep_instructionContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#tmpl_block.
	VisitTmpl_block(ctx *Tmpl_blockContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#prep_def_args.
	VisitPrep_def_args(ctx *Prep_def_argsContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#prep_def_arg_lines.
	VisitPrep_def_arg_lines(ctx *Prep_def_arg_linesContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#instruction.
	VisitInstruction(ctx *InstructionContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#arglist_p.
	VisitArglist_p(ctx *Arglist_pContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#arglist.
	VisitArglist(ctx *ArglistContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#arglist_lines.
	VisitArglist_lines(ctx *Arglist_linesContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#str.
	VisitStr(ctx *StrContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#num.
	VisitNum(ctx *NumContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#reg.
	VisitReg(ctx *RegContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#namelist.
	VisitNamelist(ctx *NamelistContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by DD8ASMParser#label.
	VisitLabel(ctx *LabelContext) interface{}
}
