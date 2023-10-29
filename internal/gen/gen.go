package gen

import (
	"fmt"
	"grimlang/internal/ast"
	"grimlang/internal/ir"
	"grimlang/internal/object"
	tp "grimlang/internal/type"
)

type GenContext struct {
	CurrentScope string
	ReturnType   *tp.Type
}

func (g GenContext) Copy() GenContext {
	return GenContext{
		CurrentScope: g.CurrentScope,
		ReturnType:   g.ReturnType,
	}
}

type IRGenerator struct{}

func (g IRGenerator) GenerateModule(prog *ast.ProgramAST) (ir.Module, error) {
	module := ir.NewModule(prog.Name)
	global := ir.NewFunction("global")
	ctx := GenContext{CurrentScope: "global"}
	for _, gl := range prog.Body {
		switch gl := gl.(type) {
		case *ast.VarAST:
			res, err := g.GenerateGlobal(ctx, gl)
			if err != nil {
				return module, err
			}
			global.PushInstructions(res...)
		case *ast.FunctionAST:
			fn, err := g.GenerateFunction(ctx, gl)
			if err != nil {
				return module, err
			}
			module.PushFunctions(fn)
		}
	}
	module.SetGlobal(global)
	return module, nil
}

func (g IRGenerator) GenerateFunction(ctx GenContext, fn *ast.FunctionAST) (ir.Function, error) {
	fir := ir.NewFunction(fn.Symbol.String())
	ctx.CurrentScope = "local"
	if fn.ReturnTypes != nil {
		t, err := g.GenerateType(ctx, fn.ReturnTypes)
		if err != nil {
			return fir, err
		}
		fir.Meta.Type = &t
		ctx.ReturnType = &t
	}

	for i := len(fn.Args) - 1; i >= 0; i-- {
		arg := fn.Args[i]
		t, err := g.GenerateType(ctx, arg.Type)
		if err != nil {
			return fir, err
		}
		fir.PushInstructions(
			ir.VarNew(arg.Symbol.String(), t),
			ir.StackType(t),
			ir.VarSave(arg.Symbol.String()),
		)
		fir.Meta.ArgTypes = append(fir.Meta.ArgTypes, t)
	}

	for _, lc := range fn.Body {
		res, err := g.GenerateLocal(ctx, lc)
		if err != nil {
			return fir, err
		}
		fir.PushInstructions(res...)
	}
	return fir, nil
}

func (g IRGenerator) GenerateGlobal(ctx GenContext, gl ast.GLOBAL) ([]ir.Instruction, error) {
	switch gl := gl.(type) {
	case *ast.VarAST:
		return g.GenerateVariable(ctx, gl)
	default:
		return nil, fmt.Errorf("can't generate global instructions for %T", gl)
	}
}

func (g IRGenerator) GenerateLocal(ctx GenContext, lc ast.LOCAL) ([]ir.Instruction, error) {
	switch lc := lc.(type) {
	case *ast.VarAST:
		return g.GenerateVariable(ctx, lc)
	case *ast.SetAST:
		return g.GenerateSet(ctx, lc)
	case *ast.SCallAST:
		return g.GenerateCall(ctx, lc)
	case *ast.ReturnAST:
		return g.GenerateReturn(ctx, lc)
	case *ast.IfAST:
		return g.GenerateIf(ctx, lc)
	default:
		return nil, fmt.Errorf("can't generate local instructions for %T", lc)
	}
}

func (g IRGenerator) GenerateExpression(ctx GenContext, exp ast.EXPR) ([]ir.Instruction, error) {
	switch exp := exp.(type) {
	case *ast.SCallAST:
		return g.GenerateCall(ctx, exp)
	case *ast.SymbolAST:
		code := make([]ir.Instruction, 0)
		code = append(code,
			ir.VarLoad(exp.String()),
		)
		return code, nil
	case *ast.IntAST:
		code := make([]ir.Instruction, 0)
		obj, err := g.GenerateObject(ctx, exp)
		if err != nil {
			return nil, err
		}
		instr := ir.ConstLoad(obj)
		instr.Type = tp.NewInt()
		code = append(code,
			instr,
		)
		return code, nil
	default:
		return nil, fmt.Errorf("can't generate expression from %T", exp)
	}
}

func (g IRGenerator) GenerateObject(ctx GenContext, at ast.ATOM) (object.Object, error) {
	switch at := at.(type) {
	case *ast.IntAST:
		return object.NewInt(at.Value), nil
	default:
		return object.Object{}, fmt.Errorf("can't generate object from %T", at)
	}
}

func (g IRGenerator) GenerateType(ctx GenContext, t ast.Type) (tp.Type, error) {
	switch t := t.(type) {
	case *ast.IntType:
		return tp.NewInt(), nil
	default:
		return tp.Type{}, fmt.Errorf("can't generate type from %T", t)
	}
}

func (g IRGenerator) GenerateVariable(ctx GenContext, vr *ast.VarAST) ([]ir.Instruction, error) {
	res, err := g.GenerateExpression(ctx, vr.Expression)
	if err != nil {
		return nil, err
	}

	t, err := g.GenerateType(ctx, vr.Type)
	if err != nil {
		return nil, err
	}

	code := make([]ir.Instruction, 0)
	code = append(code, res...)
	code = append(code, ir.VarNew(vr.Symbol.String(), t))
	code = append(code, ir.StackType(t))
	code = append(code, ir.VarSave(vr.Symbol.String()))
	return code, nil
}

func (g IRGenerator) GenerateSet(ctx GenContext, st *ast.SetAST) ([]ir.Instruction, error) {
	res, err := g.GenerateExpression(ctx, st.Expression)
	if err != nil {
		return nil, err
	}

	code := make([]ir.Instruction, 0)
	code = append(code, res...)
	code = append(code, ir.VarSave(st.Symbol.String()))
	return code, nil
}

func (g IRGenerator) GenerateCall(ctx GenContext, call *ast.SCallAST) ([]ir.Instruction, error) {
	code := make([]ir.Instruction, 0)
	for _, arg := range call.Arguments {
		res, err := g.GenerateExpression(ctx, arg)
		if err != nil {
			return nil, err
		}
		code = append(code, res...)
	}
	code = append(code, ir.Call(call.Function.String(), object.NewInt(len(call.Arguments))))
	return code, nil
}

func (g IRGenerator) GenerateReturn(ctx GenContext, ret *ast.ReturnAST) ([]ir.Instruction, error) {
	code := make([]ir.Instruction, 0)
	if ret.Value != nil {
		res, err := g.GenerateExpression(ctx, ret.Value)
		if err != nil {
			return nil, err
		}
		code = append(code, res...)
	}
	count := 0
	if ctx.ReturnType != nil {
		count = 1
		code = append(code, ir.StackType(*ctx.ReturnType))
	}
	code = append(code, ir.Return(object.NewInt(count)))
	return code, nil
}

func (g IRGenerator) GenerateIf(ctx GenContext, bl *ast.IfAST) ([]ir.Instruction, error) {
	ctx = ctx.Copy()
	ctx.CurrentScope = RandStringBytes(8)
	code := make([]ir.Instruction, 0)
	thenLabel := fmt.Sprintf("begin_if_%s", ctx.CurrentScope)
	endLabel := fmt.Sprintf("end_if_%s", ctx.CurrentScope)
	res, err := g.GenerateExpression(ctx, bl.IfCondition)
	if err != nil {
		return nil, err
	}
	code = append(code, res...)
	thenCode, err := g.GenerateBlock(ctx, bl.IfBody)
	if err != nil {
		return nil, err
	}
	if bl.ElseBody != nil {
		elseLabel := fmt.Sprintf("else_%s", ctx.CurrentScope)
		code = append(code, ir.BrTrue(thenLabel, elseLabel))
		code = append(code, ir.Label(thenLabel))
		code = append(code, thenCode...)
		code = append(code, ir.Br(endLabel))
		elseCode := make([]ir.Instruction, 0)
		elseCode = append(elseCode, ir.Label(elseLabel))
		res, err := g.GenerateBlock(ctx, bl.ElseBody)
		if err != nil {
			return nil, err
		}
		elseCode = append(elseCode, res...)
		// elseCode = append(elseCode, ir.Br(endLabel))
		code = append(code, elseCode...)
		code = append(code, ir.Label(endLabel))
	} else {
		code = append(code, ir.BrTrue(thenLabel, endLabel))
		code = append(code, ir.Label(thenLabel))
		code = append(code, thenCode...)
		code = append(code, ir.Label(endLabel))
	}

	return code, nil
}

func (g IRGenerator) GenerateBlock(ctx GenContext, locals []ast.LOCAL) ([]ir.Instruction, error) {
	code := make([]ir.Instruction, 0)
	for _, lc := range locals {
		res, err := g.GenerateLocal(ctx, lc)
		if err != nil {
			return nil, err
		}
		code = append(code, res...)
	}
	return code, nil
}
