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
		return _GenVar(ctx, node)
	default:
		log.Fatalf("can't generate global code from %T", node)
	}
	return nil
}

func _GenFunction(fn *ast.FunctionAST) *ir.FunctionIR {
	ctx := context.WithValue(context.TODO(), "fnName", fn.Symbol.String())
	ctx = context.WithValue(ctx, "retTypes", fn.ReturnTypes)
	tok := 0
	ctx = context.WithValue(ctx, "token", &tok)
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
		return _GenVar(ctx, node)
	case *ast.SCallAST:
		return _GenSCall(ctx, node)
	case *ast.ReturnAST:
		return _GenRet(ctx, node)
	case *ast.IfAST:
		return _GenIf(ctx, node)
	}
	return nil
}

func _GenExpr(ctx context.Context, ex ast.EXPR) []*ir.InstrIR {
	switch ex := ex.(type) {
	case *ast.SCallAST:
		return _GenSCall(ctx, ex)
	case *ast.IntAST:
		code := make([]*ir.InstrIR, 0)
		code = append(code, ir.ConstLoadInt(_GenInt(ctx, *ex)))
		return code
	case *ast.SymbolAST:
		sm := _GenSymbol(ctx, *ex)
		code := make([]*ir.InstrIR, 0)
		code = append(code, ir.VarLoad(sm))
	}
	return nil
}

func _GenVar(ctx context.Context, vr *ast.VarAST) []*ir.InstrIR {
	m := ir.NewModule("")
	sm := _GenSymbol(ctx, vr.Symbol)
	tp := _GenType(ctx, vr.Type)
	m.WriteInstrs(ir.VarNew(sm))
	m.WriteInstrs(_GenExpr(ctx, vr.Expression)...)
	m.WriteInstrs(ir.StackType(tp))
	m.WriteInstrs(ir.VarSave(sm))
	return m.Init
}

func _GenSCall(ctx context.Context, sc *ast.SCallAST) []*ir.InstrIR {
	code := make([]*ir.InstrIR, 0)
	for _, arg := range sc.Arguments {
		code = append(code, _GenExpr(ctx, arg)...)
	}
	l := ir.NewInt(len(sc.Arguments))
	sm := _GenSymbol(ctx, sc.Function)
	code = append(code, ir.Call(sm, l))
	return code
}

func _GenRet(ctx context.Context, rt *ast.ReturnAST) []*ir.InstrIR {
	code := make([]*ir.InstrIR, 0)
	code = append(code, _GenExpr(ctx, rt.Value)...)
	if retTypes, ok := ctx.Value("retTypes").([]ast.TypeAST); ok {
		l := ir.NewInt(1)
		tp := _GenType(ctx, retTypes[0])
		code = append(code, ir.StackType(tp))
		code = append(code, ir.Return(l))
	} else {
		l := ir.NewInt(0)
		code = append(code, ir.Return(l))
	}
	return code
}

func _GenIf(ctx context.Context, ifel *ast.IfAST) []*ir.InstrIR

func _GenInt(ctx context.Context, in ast.IntAST) *ir.IntIR

func _GenSymbol(ctx context.Context, sm ast.SymbolAST) *ir.SymbolIR

func _GenType(ctx context.Context, tp ast.TypeAST) *ir.TypeIR
