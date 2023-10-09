package gen

import (
	"context"
	"grimlang/internal/ast"
	"grimlang/internal/ir"
)

func _GenModule(prog *ast.ProgramAST) *ir.ModuleIR

func _GenFunction(fn *ast.FunctionAST) *ir.FunctionIR

func _GenGlobal(ctx context.Context, node *ast.GLOBAL) []*ir.InstrIR

func _GenLocal(ctx context.Context, node *ast.LOCAL) []*ir.InstrIR

func _GenExpr(ctx context.Context, ex *ast.EXPR) []*ir.InstrIR

func _GenAtom(ctx context.Context, at *ast.ATOM) []*ir.InstrIR

func _GenVar(ctx context.Context, vr *ast.VarAST) []*ir.InstrIR

func _GenRet(ctx context.Context, rt *ast.ReturnAST) []*ir.InstrIR

func _GenInt(ctx context.Context, in *ast.IntAST) []*ir.InstrIR

func _GenSymbol(ctx context.Context, in *ast.IntAST) []*ir.InstrIR

func _GenType(ctx context.Context, in *ast.IntAST) []*ir.InstrIR
