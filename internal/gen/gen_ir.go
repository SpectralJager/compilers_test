package gen

import (
	"grimlang/internal/ast"
	"grimlang/internal/ir"
	"log"
)

type GeneratorModuleIR struct {
	Program *ast.ProgramAST
	Module  ir.ModuleIR
	rec     chan ir.FunctionIR
}

type GenerateFunctionIR struct {
	Fast *ast.FunctionAST
	Fir  *ir.FunctionIR
}

func NewGenIr(prog *ast.ProgramAST) *GeneratorModuleIR {
	return &GeneratorModuleIR{
		Program: prog,
		Module:  ir.NewModule(prog.Name),
		rec:     make(chan ir.FunctionIR),
	}
}

func (g *GeneratorModuleIR) Start() *ir.ModuleIR {
	cnt := 0
	for _, gl := range g.Program.Body {
		switch gl := gl.(type) {
		case *ast.FunctionAST:
			cnt += 1
			go func() {
				res := g.genFunc(gl)
				g.rec <- res
			}()
		default:
			code := g.genCode(gl)
			g.Module.WriteInstrs(code...)
		}
	}
	if cnt != 0 {
		for res := range g.rec {
			g.Module.WriteFunctions(res)
			cnt -= 1
			if cnt <= 0 {
				break
			}
		}
	}
	close(g.rec)
	return &g.Module
}

func (g *GeneratorIR) genFunc(f *ast.FunctionAST) ir.FunctionIR {
	fn := ir.NewFunction(*g.genAtom(&f.Symbol).(*ir.SymbolIR))
	for i := len(f.Args) - 1; i >= 0; i-- {
		arg := f.Args[i]
		sm := g.genAtom(&arg.Symbol).(*ir.SymbolIR)
		tp := g.genAtom(&arg.Type).(*ir.TypeIR)
		fn.WriteInstrs(
			ir.VarNew(*sm),
			ir.StackType(*tp),
			ir.VarSave(*sm),
		)
	}
	for _, v := range f.Body {
		if f.ReturnTypes != nil {
			fn.WriteInstrs(g.genCode(v, &f.ReturnTypes[0])...)
		} else {
			fn.WriteInstrs(g.genCode(v)...)
		}
	}
	return fn
}

func (g *GeneratorIR) genCode(node ast.AST, ops ...ast.AST) []ir.InstrIR {
	code := make([]ir.InstrIR, 0)
	switch node := node.(type) {
	case *ast.IntAST:
		v := g.genAtom(node).(*ir.IntIR)
		code = append(code, ir.ConstLoadInt(*v))
	case *ast.SymbolAST:
		sb := g.genAtom(node).(*ir.SymbolIR)
		code = append(code, ir.VarLoad(*sb))
	case *ast.VarAST:
		sb := g.genAtom(&node.Symbol).(*ir.SymbolIR)
		tp := g.genAtom(&node.Type).(*ir.TypeIR)
		exp := g.genCode(node.Expression)
		code = append(code, ir.VarNew(*sb))
		code = append(code, exp...)
		code = append(code, ir.StackType(*tp))
		code = append(code, ir.VarSave(*sb))
	case *ast.SCallAST:
		for _, arg := range node.Arguments {
			code = append(code, g.genCode(arg)...)
		}
		count := ir.NewInt(len(node.Arguments))
		sb := g.genAtom(&node.Function).(*ir.SymbolIR)
		code = append(code, ir.Call(*sb, count))
	case *ast.ReturnAST:
		code = append(code, g.genCode(node.Value)...)
		tp := g.genAtom(ops[0].(*ast.TypeAST)).(*ir.TypeIR)
		cnt := 0
		if node.Value != nil {
			cnt = 1
		}
		code = append(code, ir.StackType(*tp))
		code = append(code, ir.Return(ir.NewInt(cnt)))
	case *ast.IfAST:
		code = append(code, g.genCode(node.IfCondition)...)
		var th, el []ir.InstrIR
		var tlab, ellab, end ir.SymbolIR
		for _, c := range node.IfBody {
			th = append(th, g.genCode(c, ops...)...)
		}
		if node.ElseBody != nil {
			for _, c := range node.ElseBody {
				el = append(el, g.genCode(c, ops...)...)
			}
		}
		tlab = ir.NewSymbol("if_{}", nil)
		end = ir.NewSymbol("endif_{}", nil)
		if el != nil {
			ellab = ir.NewSymbol("else_{}", nil)
			code = append(code, ir.BrTrue(tlab, ellab))
			code = append(code, ir.Lable(tlab))
			code = append(code, th...)
			code = append(code, ir.Br(end))
			code = append(code, ir.Lable(ellab))
			code = append(code, el...)
			code = append(code, ir.Br(end))
			code = append(code, ir.Lable(end))
		} else {
			code = append(code, ir.BrTrue(tlab, end))
			code = append(code, ir.Lable(tlab))
			code = append(code, th...)
			code = append(code, ir.Br(end))
			code = append(code, ir.Lable(end))
		}

	default:
		log.Fatalf("GEN_ERR: can't generate code from: %T", node)
	}
	return code
}

func (g *GeneratorIR) genAtom(node ast.ATOM) ir.IR {
	switch node := node.(type) {
	case *ast.SymbolAST:
		var sm *ir.SymbolIR
		if node.Additional == nil {
			tmp := ir.NewSymbol(node.Primary, nil)
			sm = &tmp
		} else {
			tmp := ir.NewSymbol(node.Primary, g.genAtom(node.Additional).(*ir.SymbolIR))
			sm = &tmp
		}
		return sm
	case *ast.TypeAST:
		sm := g.genAtom(&node.Primary).(*ir.SymbolIR)
		tp := ir.NewType(*sm)
		var gns []ir.TypeIR
		for _, st := range node.Generic {
			t := g.genAtom(&st).(*ir.TypeIR)
			gns = append(gns, *t)
		}
		tp.Generic = gns
		return &tp
	case *ast.IntAST:
		v := ir.NewInt(node.Value)
		return &v
	default:
		log.Fatalf("GEN_ERR: can't generate from atom: %T", node)
		return nil
	}
}
