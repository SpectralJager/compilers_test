package gen

import (
	"grimlang/internal/ast"
	"grimlang/internal/ir"
	"log"
	"sync"
)

type GeneratorIR struct {
	Program *ast.ProgramAST
	Module  ir.ModuleIR
	rec     chan ir.FunctionIR
}

func (g *GeneratorIR) Init(prog *ast.ProgramAST) *GeneratorIR {
	g.Program = prog
	g.Module = ir.NewModule(prog.Name)
	return g
}

func (g *GeneratorIR) Start() *ir.ModuleIR {
	wg := new(sync.WaitGroup)
	select {
	case res := <-g.rec:
		g.Module.WriteFunctions(res)
	default:
		for _, gl := range g.Program.Body {
			switch gl := gl.(type) {
			case *ast.FunctionAST:
				go func() {
					wg.Add(1)
					res := g.genFunc(gl)
					g.rec <- res
					wg.Done()
				}()
			default:
				code := g.genCode(gl)
				g.Module.WriteInstrs(code...)
			}
		}
	}
	wg.Wait()
	return &g.Module
}

func (g *GeneratorIR) genFunc(f *ast.FunctionAST) ir.FunctionIR {
	fn := ir.NewFunction(*g.genAtom(&f.Symbol).(*ir.SymbolIR))
	return fn
}

func (g *GeneratorIR) genCode(node ast.AST) []ir.InstrIR {
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
		code = append(code, ir.VarNew(*sb))
		exp := g.genCode(node.Expression)
		code = append(code, exp...)
		code = append(code, ir.VarSave(*sb))
	case *ast.SCallAST:

	default:
		log.Fatalf("GEN_ERR: can't generate code from: %T", node)
	}
	return code
}

func (g *GeneratorIR) genAtom(node ast.ATOM) ir.IR {
	switch node := node.(type) {
	case *ast.SymbolAST:
		sm := ir.NewSymbol(
			node.Primary,
			g.genAtom(node.Additional).(*ir.SymbolIR),
		)
		return &sm
	case *ast.IntAST:
		v := ir.NewInt(node.Value)
		return &v
	default:
		log.Fatalf("GEN_ERR: can't generate from atom: %T", node)
		return nil
	}
}
