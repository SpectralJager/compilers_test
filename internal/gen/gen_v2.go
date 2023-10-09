package gen

import (
	"context"
	"grimlang/internal/ast"
	"grimlang/internal/ir"
	"log"
)

func _GenModule(prog *ast.ProgramAST) *ir.ModuleIR {
	m := ir.NewModule(prog.Name)
	for _, gl := range prog.Body {
		switch gl := gl.(type) {
		case *ast.FunctionAST:
		default:
			m.WriteInstrs(_GenGlobal(context.TODO(), gl)...)
		}
	}
	return m
}

func _GenGlobal(ctx context.Context, node ast.GLOBAL) []*ir.InstrIR {
	switch node := node.(type) {
	case *ast.VarAST:
		return _GenVar(ctx, *node)
	default:
		log.Fatalf("can't generate global code from %T", node)
	}
	return nil
}

func _GenFunction(fn *ast.FunctionAST) *ir.FunctionIR {
	ctx := context.WithValue(context.TODO(), "fnName", fn.Symbol.String())
	ctx = context.WithValue(ctx, "retTypes", fn.ReturnTypes)
	ctx = context.WithValue(ctx, "token", 0)
	fir := ir.NewFunction(*_GenSymbol(ctx, fn.Symbol))

	for i := len(fn.Args); i >= 0; i-- {
		arg := fn.Args[i]
		sm := _GenSymbol(ctx, arg.Symbol)
		tp := _GenType(ctx, arg.Type)
		fir.WriteInstrs(
			ir.VarNew(sm),
			ir.StackType(tp),
			ir.VarSave(sm),
		)
	}

	for _, lc := range fn.Body {
		fir.WriteInstrs(_GenLocal(ctx, lc)...)
	}

	return fir
}

func _GenLocal(ctx context.Context, node ast.LOCAL) []*ir.InstrIR {
	switch node := node.(type) {
	case *ast.VarAST:
		return _GenVar(ctx, *node)
	case *ast.SCallAST:
		return _GenSCall(ctx, *node)
	case *ast.ReturnAST:
		return _GenRet(ctx, *node)
	}
	return nil
}

func _GenVar(ctx context.Context, vr ast.VarAST) []*ir.InstrIR {
	m := ir.NewModule("")
	sm := _GenSymbol(ctx, vr.Symbol)
	tp := _GenType(ctx, vr.Type)
	m.WriteInstrs(ir.VarNew(sm))
	m.WriteInstrs(_GenExpr(ctx, vr.Expression)...)
	m.WriteInstrs(ir.StackType(tp))
	m.WriteInstrs(ir.VarSave(sm))
	return m.Init
}

func _GenExpr(ctx context.Context, ex ast.EXPR) []*ir.InstrIR

func _GenAtom(ctx context.Context, at ast.ATOM) []*ir.InstrIR

func _GenIf(ctx context.Context, ifel ast.IfAST) []*ir.InstrIR

func _GenSCall(ctx context.Context, sc ast.SCallAST) []*ir.InstrIR

func _GenRet(ctx context.Context, rt ast.ReturnAST) []*ir.InstrIR

func _GenInt(ctx context.Context, in ast.IntAST) *ir.IntIR

func _GenSymbol(ctx context.Context, sm ast.SymbolAST) *ir.SymbolIR

func _GenType(ctx context.Context, tp ast.TypeAST) *ir.TypeIR
