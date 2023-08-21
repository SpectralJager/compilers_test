package compiler

import (
	"fmt"
	"gl/internal/ir"
)

var builtins = map[string][]byte{
	"iadd": ir.IntFunc(0),
	"isub": ir.IntFunc(1),
	"imul": ir.IntFunc(2),
	"idiv": ir.IntFunc(3),
	"ilt":  ir.IntFunc(4),
	"igt":  ir.IntFunc(5),
	"igeq": ir.IntFunc(6),
	"ileq": ir.IntFunc(7),
	"ieq":  ir.IntFunc(8),
	"itof": ir.IntFunc(9),
	"itos": ir.IntFunc(10),
	"itob": ir.IntFunc(11),
}

type Generator struct {
	Program  ir.Program
	Function ir.Function
}

func (g *Generator) GenerateProgram(program ProgramAST) (ir.Program, error) {
	g.Program = ir.Program{
		Name:      program.Name,
		Constants: []ir.IConstant{},
		Globals:   []ir.ISymbolDef{},
		InitCode:  ir.NewCode(),
		Functions: map[string]*ir.Function{},
	}
	for _, node := range program.Body {
		switch node := node.(type) {
		case *FunctionAST:
			def, err := g.GenerateDefenition(*node)
			if err != nil {
				return g.Program, err
			}
			g.Program.Globals = append(g.Program.Globals, def)

			irfn, err := g.GenerateFunction(*node)
			if err != nil {
				return g.Program, err
			}
			g.Program.Functions[node.Ident] = &irfn
		case *VaribleAST:
			var def ir.ISymbolDef
			def, err := g.GenerateDefenition(node.Symbol)
			if err != nil {
				return g.Program, err
			}
			ind := len(g.Program.Globals)
			g.Program.Globals = append(g.Program.Globals, def)
			c, err := g.GenerateExpression(node.Value)
			if err != nil {
				return g.Program, err
			}
			code := g.Program.InitCode
			code = code.WriteBytes(*c...)
			code = code.WriteBytes(ir.GlobalSave(uint32(ind))...)
			g.Program.InitCode = code
		default:
			return g.Program, fmt.Errorf("unexpected global ast node: %T", node)
		}
	}
	for i, def := range g.Program.Globals {
		if def, ok := def.(*ir.FunctionDef); ok {
			if def.Name == "main" {
				g.Program.InitCode = g.Program.InitCode.WriteBytes(ir.Call(uint32(i))...)
			}
		}
	}
	return g.Program, nil
}

func (g *Generator) GenerateDataType(dt DataType) (ir.IDataType, error) {
	var irdt ir.IDataType
	switch dt := dt.(type) {
	case *SimpleDataTypeAST:
		irdt = &ir.Primitive{Name: dt.Value}
	default:
		return irdt, fmt.Errorf("unexpected datatype %T", dt)
	}
	return irdt, nil
}

func (g *Generator) GenerateDefenition(node AST) (ir.ISymbolDef, error) {
	var def ir.ISymbolDef
	switch node := node.(type) {
	case FunctionAST:
		var fnDef ir.FunctionDef
		fnDef.Name = node.Ident
		for _, v := range node.Arguments {
			dt, err := g.GenerateDataType(v.DataType)
			if err != nil {
				return def, err
			}
			fnDef.Arguments = append(fnDef.Arguments, dt)
		}
		for _, v := range node.RetTypes {
			dt, err := g.GenerateDataType(v)
			if err != nil {
				return def, err
			}
			fnDef.Returns = append(fnDef.Returns, dt)
		}
		def = &fnDef
	case SymbolDeclAST:
		var vrDef ir.VaribleDef
		vrDef.Name = node.Ident
		dt, err := g.GenerateDataType(node.DataType)
		if err != nil {
			return def, err
		}
		vrDef.Type = dt
		def = &vrDef
	default:
		return def, fmt.Errorf("unexpected ast node for defenition: %T", node)
	}
	return def, nil
}

func (g *Generator) GenerateFunction(fn FunctionAST) (ir.Function, error) {
	g.Function = ir.Function{
		Name:     fn.Ident,
		Locals:   []ir.ISymbolDef{},
		BodyCode: ir.NewCode(),
	}
	for i, v := range fn.Arguments {
		df, err := g.GenerateDefenition(v)
		if err != nil {
			return g.Function, err
		}
		g.Function.Locals = append(g.Function.Locals, df)

		g.Function.BodyCode = g.Function.BodyCode.WriteBytes(ir.LocalSave(uint32(i))...)
	}
	for _, v := range fn.Body {
		code, err := g.GenerateLocal(v)
		if err != nil {
			return g.Function, err
		}
		g.Function.BodyCode = g.Function.BodyCode.WriteBytes(*code...)
	}
	return g.Function, nil
}

func (g *Generator) GenerateLocal(l Local) (*ir.Code, error) {
	code := ir.NewCode()
	switch l := l.(type) {
	case *SymbolExpressionAST:
		c, err := g.GenerateExpression(l)
		if err != nil {
			return code, err
		}
		code = code.WriteBytes(*c...)
	case *ReturnAST:
		c, err := g.GenerateExpression(l.Expression)
		if err != nil {
			return code, err
		}
		code = code.WriteBytes(*c...)
		code = code.WriteBytes(ir.Return(1)...)
	case *VaribleAST:
		var def ir.ISymbolDef
		def, err := g.GenerateDefenition(l.Symbol)
		if err != nil {
			return code, err
		}
		ind := len(g.Function.Locals)
		g.Function.Locals = append(g.Function.Locals, def)
		c, err := g.GenerateExpression(l.Value)
		if err != nil {
			return code, err
		}
		code = code.WriteBytes(*c...)
		code = code.WriteBytes(ir.LocalSave(uint32(ind))...)
	case *SetAST:
		c, err := g.GenerateExpression(l.Value)
		if err != nil {
			return code, err
		}
		code = code.WriteBytes(*c...)
		ind := -1
		for i, v := range g.Function.Locals {
			if v.(*ir.VaribleDef).Name == l.Ident {
				ind = i
				break
			}
		}
		if ind == -1 {
			for i, v := range g.Program.Globals {
				if v, ok := v.(*ir.VaribleDef); ok && v.Name == l.Ident {
					ind = i
					break
				}
			}
			if ind == -1 {
				return code, fmt.Errorf("symbol '%s' not found for function '%s'", l.Ident, g.Function.Name)
			}
			code = code.WriteBytes(ir.GlobalSave(uint32(ind))...)
		} else {
			code = code.WriteBytes(ir.LocalSave(uint32(ind))...)
		}
	case *IfAST:
		exprCode, err := g.GenerateExpression(l.Expr)
		if err != nil {
			return code, err
		}
		code = code.WriteBytes(*exprCode...)
		start := len(*g.Function.BodyCode) + len(*exprCode)
		thenCode := ir.NewCode()
		for _, v := range l.ThenBody {
			temp, err := g.GenerateLocal(v)
			if err != nil {
				return code, err
			}
			thenCode = thenCode.WriteBytes(*temp...)
		}
		thenOffset := len(*thenCode)
		elseCode := ir.NewCode()
		for _, v := range l.ElseBody {
			temp, err := g.GenerateLocal(v)
			if err != nil {
				return code, err
			}
			elseCode = elseCode.WriteBytes(*temp...)
		}
		elseOffset := len(*elseCode)
		elseCode = elseCode.WriteBytes(ir.Jump(uint32(start + thenOffset + elseOffset + 5 + 5))...)
		elseOffset = len(*elseCode)
		code = code.WriteBytes(ir.JumpCondition(uint32(start + elseOffset + 5))...)
		code = code.WriteBytes(*elseCode...)
		code = code.WriteBytes(*thenCode...)
		code = code.WriteByte(ir.Null())
	default:
		return code, fmt.Errorf("unexpected local ast type %T", l)
	}
	return code, nil
}

func (g *Generator) GenerateExpression(expr Expression) (*ir.Code, error) {
	code := ir.NewCode()
	switch expr := expr.(type) {
	case *IntegerAST:
		intConst := ir.Integer{Value: expr.Integer}
		ind := -1
		for i, v := range g.Program.Constants {
			if v.String() == intConst.String() {
				ind = i
			}
		}
		if ind == -1 {
			ind = len(g.Program.Constants)
			g.Program.Constants = append(g.Program.Constants, &intConst)
		}
		code = code.WriteBytes(ir.Load(uint32(ind))...)
	case *SymbolAST:
		ind := -1
		for i, s := range g.Function.Locals {
			if s.(*ir.VaribleDef).Name == expr.Symbol {
				ind = i
				break
			}
		}
		if ind == -1 {
			for i, v := range g.Program.Globals {
				if v, ok := v.(*ir.VaribleDef); ok && v.Name == expr.Symbol {
					ind = i
					break
				}
			}
			if ind == -1 {
				return code, fmt.Errorf("symbol '%s' not found for function '%s'", expr.Symbol, g.Function.Name)
			}
			code = code.WriteBytes(ir.GlobalLoad(uint32(ind))...)
		} else {
			code = code.WriteBytes(ir.LocalLoad(uint32(ind))...)
		}
	case *SymbolExpressionAST:
		for i := len(expr.Arguments) - 1; i >= 0; i-- {
			c, err := g.GenerateExpression(expr.Arguments[i])
			if err != nil {
				return code, err
			}
			code = code.WriteBytes(*c...)
		}
		ind := -1
		for i, s := range g.Program.Globals {
			if s, ok := s.(*ir.FunctionDef); ok && s.Name == expr.Symbol {
				ind = i
				break
			}
		}
		if ind == -1 {
			if bt, ok := builtins[expr.Symbol]; ok {
				code = code.WriteBytes(bt...)
				return code, nil
			}
			return code, fmt.Errorf("function '%s' not found", expr.Symbol)
		}
		code = code.WriteBytes(ir.Call(uint32(ind))...)
	default:
		return code, fmt.Errorf("unexpected argument type %T", expr)
	}
	return code, nil
}
