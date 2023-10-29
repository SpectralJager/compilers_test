package semantic

import (
	"fmt"
	"grimlang/internal/ir"
)

func SemanticModule(ctx *SemanticContext, module ir.Module) error {
	err := collectFunctions(ctx, module.Functions)
	if err != nil {
		return err
	}
	err = SemanticGlobals(ctx, module.Global.Body)
	if err != nil {
		return err
	}
	return nil
}

func collectFunctions(ctx *SemanticContext, fns []ir.Function) error {
	for _, fn := range fns {
		fs := NewFunctionSymbol(fn.Name, fn.Meta.ArgTypes, fn.Meta.Type)
		err := ctx.AppendSymbol(fs)
		if err != nil {
			return err
		}
	}
	return nil
}

func SemanticGlobals(ctx *SemanticContext, gls []ir.Instruction) error {
	for _, instr := range gls {
		switch instr.Kind {
		case ir.OP_VAR_NEW:
			vr := NewGlobalSymbol(instr.Identifier, instr.Type)
			err := ctx.AppendSymbol(vr)
			if err != nil {
				return err
			}
		case ir.OP_CONST_LOAD:
			ctx.PushType(instr.Type)
		case ir.OP_CALL:
			fn, err := ctx.FindFunction(instr.Identifier)
			if err != nil {
				return err
			}
			fmt.Println(fn.String())
		}
	}
	return nil
}
