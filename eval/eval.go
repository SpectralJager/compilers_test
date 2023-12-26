package eval

import (
	"fmt"
	"grimlang/ast"
	"grimlang/context"
	"grimlang/dtype"
	"grimlang/object"
	"grimlang/symbol"
)

var globalContext = context.NewContext("", nil)

func EvalModule(ctx context.Context, md *ast.Module) error {
	switch md.Kind {
	case ":main":
		globalContext = context.NewContext("global", ctx)
		for _, gl := range md.Body {
			err := EvalGlobal(globalContext, gl)
			if err != nil {
				return err
			}
		}
		mainSymbol := globalContext.Search("main")
		if mainSymbol == nil {
			return fmt.Errorf("function main not found")
		}
		if mainSymbol.Kind() != symbol.Function {
			return fmt.Errorf("main should be function")
		}
		mainFunc := mainSymbol.(*symbol.FunctionSymbol)
		_, err := EvalFunction(globalContext, mainFunc.Fn)
		return err
	default:
		return fmt.Errorf("eval: unexpected module kind %s", md.Kind)
	}
}

// ================================================================

func EvalGlobal(ctx context.Context, gl ast.Global) error {
	switch gl := gl.(type) {
	case *ast.ConstantDecl:
		return NewConstantSymbol(ctx, gl)
	case *ast.FunctionDecl:
		return NewFunctionSymbol(ctx, gl)
	default:
		return fmt.Errorf("eval: unexpected global statement %T", gl)
	}
}

func EvalLocal(ctx context.Context, lc ast.Local) error {
	switch lc := lc.(type) {
	case *ast.SymbolCall:
		_, err := EvalSymbolCall(ctx, lc)
		return err
	default:
		return fmt.Errorf("eval: unexpected local %T", lc)
	}
}

func EvalExpression(ctx context.Context, ex ast.Expression) (object.Object, error) { return nil, nil }

func EvalAtom(ctx context.Context, at ast.Atom) (object.Object, error) {
	switch at := at.(type) {
	case *ast.IntAtom:
		return &object.IntObject{Value: at.Value}, nil
	default:
		return nil, fmt.Errorf("eval: unexpected atom %T", at)
	}
}

func EvalType(ctx context.Context, tp ast.Type) (dtype.Type, error) {
	switch tp := tp.(type) {
	case *ast.PrimitiveType:
		return EvalPrimitiveType(ctx, tp)
	default:
		return nil, fmt.Errorf("eval: unexpected type %T", tp)
	}
}

// ================================================================

func EvalSymbol(ctx context.Context, sm *ast.SymbolExpr) (symbol.Symbol, error) {
	symb := ctx.Search(sm.Identifier)
	if symb == nil {
		return nil, fmt.Errorf("eval: symbol %s not found", sm.Identifier)
	}
	if sm.Next == nil {
		return symb, nil
	}
	switch symb := symb.(type) {
	case *symbol.ModuleSymbol:
		return EvalSymbol(symb, sm.Next)
	case *symbol.BuiltinTypeSymbol:
		return EvalSymbol(symb, sm.Next)
	default:
		return nil, fmt.Errorf("eval: symbol %s not found", sm.Next.Identifier)
	}
}

func EvalFunction(ctx context.Context, sm *ast.FunctionDecl) (object.Object, error) { return nil, nil }

func EvalPrimitiveType(ctx context.Context, tp *ast.PrimitiveType) (dtype.Type, error) {
	typeSymbol, err := EvalSymbol(ctx, &tp.Symbol)
	if err != nil {
		return nil, err
	}
	if typeSymbol.Kind() != symbol.BuiltinType {
		return nil, fmt.Errorf("eval: %s should be type", typeSymbol.Name())
	}
	switch typeSymbol := typeSymbol.(type) {
	case *symbol.BuiltinTypeSymbol:
		return typeSymbol.Type, nil
	default:
		return nil, fmt.Errorf("eval: unexpected symbol of type %T", typeSymbol)
	}
}

func EvalSymbolCall(ctx context.Context, sc *ast.SymbolCall) (object.Object, error) { return nil, nil }

// ================================================================

func NewConstantSymbol(ctx context.Context, cnst *ast.ConstantDecl) error {
	value, err := EvalAtom(ctx, cnst.Value)
	if err != nil {
		return err
	}
	return ctx.Insert(
		&symbol.ConstantSymbol{
			Identifier: cnst.Identifier,
			Value:      value,
		},
	)
}

func NewFunctionSymbol(ctx context.Context, fn *ast.FunctionDecl) error {
	args := []dtype.Type{}
	for _, arg := range fn.Arguments {
		typ, err := EvalType(ctx, arg.Type)
		if err != nil {
			return err
		}
		args = append(args, typ)
	}
	var ret dtype.Type
	if fn.Return != nil {
		var err error
		ret, err = EvalType(ctx, fn.Return)
		if err != nil {
			return err
		}
	}
	return ctx.Insert(
		&symbol.FunctionSymbol{
			Identifier: fn.Identifier,
			Fn:         fn,
			Type: &dtype.FunctionType{
				Args:   args,
				Return: ret,
			},
		},
	)
}
