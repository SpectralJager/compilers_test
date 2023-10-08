package gen

import (
	"grimlang/internal/ast"
	"grimlang/internal/ir"
)

func GenIrModule(prog *ast.ProgramAST) *ir.Module {
	m := ir.NewModule(prog.Name, prog.Path)
	m.AddChunk(GenIrGlobalChunk(prog.Body))
	for _, v := range prog.Body {
		switch v := v.(type) {
		case *ast.FunctionAST:
			m.AddChunk(GenIrChunk(v))
		}
	}
	return m
}

func GenIrChunk(f *ast.FunctionAST) *ir.Chunk {
	ch := ir.NewChunk(GenIrSymbol(&f.Symbol).String())
	bl := ir.NewBlock("")
	ch.AddBlock(bl)
	for i := len(f.Args) - 1; i >= 0; i-- {
		v := f.Args[i]
		bl.AddInstructions(ir.VarNew(GenIrSymbol(&v.Symbol), GenIrType(&v.Type)))
		bl.AddInstructions(ir.VarPop(GenIrSymbol(&v.Symbol)))
	}
	for _, v := range f.Body {
		switch v := v.(type) {
		case *ast.IfAST:
		default:
			bl.AddInstructions(GenIrLocal(v)...)
		}
	}
	return ch
}

func GenIrLocal(l ast.LOCAL) []*ir.Instruction {
	bl := ir.NewBlock("")
	switch l := l.(type) {
	case *ast.VarAST:
		bl.AddInstructions(GenIrVariable(l)...)
	case *ast.SCallAST:
		bl.AddInstructions(GenIrSCall(l)...)
	case *ast.ReturnAST:
		bl.AddInstructions(ir.Return(ir.NewInteger()))
	}
	return bl.Code
}

func GenIrGlobalChunk(node []ast.GLOBAL) *ir.Chunk {
	ch := ir.NewChunk("global")
	bl := ir.NewBlock("")
	ch.AddBlock(bl)
	for _, v := range node {
		switch v := v.(type) {
		case *ast.VarAST:
			ins := GenIrVariable(v)
			bl.AddInstructions(ins...)
		}
	}
	return ch
}

func GenIrVariable(vr *ast.VarAST) []*ir.Instruction {
	bl := ir.NewBlock("")
	nm := GenIrSymbol(&vr.Symbol)
	tp := GenIrType(&vr.Type)
	bl.AddInstructions(ir.VarNew(nm, tp))
	bl.AddInstructions(GenIrExpression(vr.Expression)...)
	bl.AddInstructions(ir.VarPop(nm))
	return bl.Code
}

func GenIrSymbol(s *ast.SymbolAST) *ir.Symbol {
	if s == nil {
		return nil
	}
	return ir.NewSymbol(s.Primary, GenIrSymbol(s.Additional))
}

func GenIrType(s *ast.TypeAST) *ir.Type {
	if s == nil {
		return nil
	}
	var tps []*ir.Type
	for _, tp := range s.Generic {
		tps = append(tps, GenIrType(&tp))
	}
	return ir.NewType(GenIrSymbol(&s.Primary), tps...)
}

func GenIrExpression(ex ast.EXPR) []*ir.Instruction {
	bl := ir.NewBlock("")
	switch ex := ex.(type) {
	case *ast.SCallAST:
		bl.AddInstructions(GenIrSCall(ex)...)
	case *ast.IntAST:
		bl.AddInstructions(
			ir.ConstPush(ir.NewInteger(ex.Value)),
		)
	case *ast.SymbolAST:
		bl.AddInstructions(
			ir.VarPush(GenIrSymbol(ex)),
		)
	}
	return bl.Code
}

func GenIrSCall(sc *ast.SCallAST) []*ir.Instruction {
	bl := ir.NewBlock("")
	l := len(sc.Arguments)
	for _, arg := range sc.Arguments {
		bl.AddInstructions(GenIrExpression(arg)...)
	}
	bl.AddInstructions(ir.Call(GenIrSymbol(&sc.Function), ir.NewInteger(l)))
	return bl.Code
}
