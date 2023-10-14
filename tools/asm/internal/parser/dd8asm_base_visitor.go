// Code generated from DD8ASM.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DD8ASM

import "github.com/antlr4-go/antlr/v4"

type BaseDD8ASMVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDD8ASMVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitPrep_instruction(ctx *Prep_instructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitPrep_def_args(ctx *Prep_def_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitPrep_def_arg_lines(ctx *Prep_def_arg_linesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitInstruction(ctx *InstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitArglist(ctx *ArglistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDD8ASMVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}
