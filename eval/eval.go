package eval

import (
	"fmt"
	"grimlang/ast"
	"grimlang/context"
	"grimlang/object"
)

type EvalState struct {
	GlobalContext context.Context
	Return        object.Litteral
	IsReturn      bool
}

func (state *EvalState) SetReturn(obj object.Litteral) {
	state.IsReturn = true
	state.Return = obj
}

func (state *EvalState) GetReturn() object.Litteral {
	state.IsReturn = false
	ret := state.Return
	state.Return = nil
	return ret
}

func (state *EvalState) EvalModule(ctx context.Context, md *ast.Module) error {
	switch md.Kind {
	case ":main":
		globalContext := context.NewContext("global", ctx)
		state.GlobalContext = globalContext
		for _, gl := range md.Body {
			err := state.EvalGlobal(globalContext, gl)
			if err != nil {
				return err
			}
		}
		mainSymbol := globalContext.Search("main")
		if mainSymbol == nil {
			return fmt.Errorf("function main not found")
		}
		if !object.Is(mainSymbol.Kind(), object.FunctionSymbol) {
			return fmt.Errorf("main should be function")
		}
		mainFunc := mainSymbol.(*object.SymbolFunction)
		_, err := state.EvalFunction(globalContext, mainFunc.Fn, nil)
		return err
	default:
		return fmt.Errorf("eval: unexpected module kind %s", md.Kind)
	}
}

// ================================================================

func (state *EvalState) EvalGlobal(ctx context.Context, gl ast.Global) error {
	switch gl := gl.(type) {
	case *ast.ConstantDecl:
		return state.NewConstantSymbol(ctx, gl)
	case *ast.FunctionDecl:
		return state.NewFunctionSymbol(ctx, gl)
	default:
		return fmt.Errorf("eval: unexpected global statement %T", gl)
	}
}

func (state *EvalState) EvalLocal(ctx context.Context, lc ast.Local) error {
	switch lc := lc.(type) {
	case *ast.SymbolCall:
		_, err := state.EvalSymbolCall(ctx, lc)
		return err
	case *ast.ReturnStmt:
		return state.EvalReturn(ctx, lc)
	case *ast.SetStmt:
		return state.EvalSet(ctx, lc)
	case *ast.IfStmt:
		return state.EvalIf(ctx, lc)
	case *ast.WhileStmt:
		return state.EvalWhile(ctx, lc)
	case *ast.ConstantDecl:
		return state.NewConstantSymbol(ctx, lc)
	case *ast.VariableDecl:
		return state.NewVariableSymbol(ctx, lc)
	default:
		return fmt.Errorf("eval: unexpected local %T", lc)
	}
}

func (state *EvalState) EvalExpression(ctx context.Context, ex ast.Expression) (object.Litteral, error) {
	switch ex := ex.(type) {
	case *ast.IntAtom:
		return state.EvalAtom(ctx, ex)
	case *ast.FloatAtom:
		return state.EvalAtom(ctx, ex)
	case *ast.StringAtom:
		return state.EvalAtom(ctx, ex)
	case *ast.BoolAtom:
		return state.EvalAtom(ctx, ex)
	case *ast.ListAtom:
		return state.EvalAtom(ctx, ex)
	case *ast.SymbolCall:
		return state.EvalSymbolCall(ctx, ex)
	case *ast.SymbolExpr:
		return state.EvalSymbolExpression(ctx, ex)
	default:
		return nil, fmt.Errorf("eval: unexpected expression %T", ex)
	}
}

func (state *EvalState) EvalAtom(ctx context.Context, at ast.Atom) (object.Litteral, error) {
	switch at := at.(type) {
	case *ast.IntAtom:
		return &object.LitteralInt{Value: at.Value}, nil
	case *ast.BoolAtom:
		return &object.LitteralBool{Value: at.Value}, nil
	case *ast.StringAtom:
		return &object.LitteralString{Value: at.Value[1 : len(at.Value)-1]}, nil
	case *ast.FloatAtom:
		return &object.LitteralFloat{Value: at.Value}, nil
	case *ast.ListAtom:
		return state.EvalList(ctx, at)
	default:
		return nil, fmt.Errorf("eval: unexpected atom %T", at)
	}
}

func (state *EvalState) EvalType(ctx context.Context, tp ast.Type) (object.DType, error) {
	switch tp := tp.(type) {
	case *ast.IntType:
		return &object.DTypeInt{}, nil
	case *ast.BoolType:
		return &object.DTypeBool{}, nil
	case *ast.StringType:
		return &object.DTypeString{}, nil
	case *ast.FloatType:
		return &object.DTypeFloat{}, nil
	case *ast.ListType:
		itemType, err := state.EvalType(ctx, tp.Child)
		if err != nil {
			return nil, err
		}
		return &object.DTypeList{ChildType: itemType}, nil
	default:
		return nil, fmt.Errorf("eval: unexpected type %T", tp)
	}
}

// ================================================================

func (state *EvalState) EvalLocals(ctx context.Context, lcs []ast.Local) error {
	for _, lc := range lcs {
		err := state.EvalLocal(ctx, lc)
		if err != nil {
			return err
		}
		if state.IsReturn {
			break
		}
	}
	return nil
}

func (state *EvalState) EvalSymbol(ctx context.Context, sm *ast.SymbolExpr) (object.Symbol, error) {
	symb := ctx.Search(sm.Identifier)
	if symb == nil {
		return nil, fmt.Errorf("eval: symbol %s not found", sm.Identifier)
	}
	if sm.Next == nil {
		return symb, nil
	}
	switch symb := symb.(type) {
	case *object.SymbolModule:
		return state.EvalSymbol(symb, sm.Next)
	default:
		return nil, fmt.Errorf("eval: symbol %s not found", sm.Next.Identifier)
	}
}

func (state *EvalState) EvalSymbolExpression(ctx context.Context, sm *ast.SymbolExpr) (object.Litteral, error) {
	symb, err := state.EvalSymbol(ctx, sm)
	if err != nil {
		return nil, err
	}
	switch symb := symb.(type) {
	case *object.SymbolConstant:
		return symb.Value, nil
	case *object.SymbolVariable:
		return symb.Value, nil
	default:
		return nil, fmt.Errorf("eval: can't get object from symbol %s of %T", symb.Name(), symb)
	}
}

func (state *EvalState) EvalFunction(ctx context.Context, fn *ast.FunctionDecl, args []object.Litteral) (object.Litteral, error) {
	localContext := context.NewContext(fn.Identifier, state.GlobalContext)
	if len(args) != len(fn.Arguments) {
		return nil, fmt.Errorf("eval: %s expect %d arguments, got %d", fn.Identifier, len(fn.Arguments), len(args))
	}
	for i, arg := range fn.Arguments {
		argType, err := state.EvalType(localContext, arg.Type)
		if err != nil {
			return nil, err
		}
		if !args[i].Type().Compare(argType) {
			return nil, fmt.Errorf("eval: %s expect %dth argument be %s, got %s", fn.Identifier, i, argType.Inspect(), args[i].Type().Inspect())
		}
		err = localContext.Insert(
			&object.SymbolVariable{
				Identifier: arg.Identifier,
				ValueType:  argType,
				Value:      args[i],
			},
		)
		if err != nil {
			return nil, err
		}
	}
	err := state.EvalLocals(localContext, fn.Body)
	if err != nil {
		return nil, err
	}
	switch {
	case fn.Return != nil && state.IsReturn:
		retType, err := state.EvalType(ctx, fn.Return)
		if err != nil {
			return nil, err
		}
		ret := state.GetReturn()
		if ret == nil {
			return nil, fmt.Errorf("eval: function expect return %s, got nothing", retType.Inspect())
		}
		state.Return = nil
		if !retType.Compare(ret.Type()) {
			return nil, fmt.Errorf("eval: function expect return %s, got %s", retType.Inspect(), ret.Type().Inspect())
		}
		return ret, nil
	case fn.Return != nil && !state.IsReturn:
		retType, err := state.EvalType(ctx, fn.Return)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("eval: function expect return %s, got nothing", retType.Inspect())
	case state.IsReturn:
		state.Return = nil
	}
	return nil, nil
}

func (state *EvalState) EvalSymbolCall(ctx context.Context, sc *ast.SymbolCall) (object.Litteral, error) {
	tempSymb, err := state.EvalSymbol(ctx, sc.Call)
	if err != nil {
		return nil, err
	}
	params := []object.Litteral{}
	for _, param := range sc.Arguments {
		obj, err := state.EvalExpression(ctx, param)
		if err != nil {
			return nil, err
		}
		params = append(params, obj)
	}
	switch tempSymb := tempSymb.(type) {
	case *object.SymbolFunction:
		return state.EvalFunction(ctx, tempSymb.Fn, params)
	case *object.SymbolBuiltin:
		return tempSymb.Fn(params...)
	default:
		return nil, fmt.Errorf("eval: unexpected %s symbol %T", tempSymb.Name(), tempSymb)
	}
}

func (state *EvalState) EvalReturn(ctx context.Context, rt *ast.ReturnStmt) error {
	if rt.Value == nil {
		state.SetReturn(nil)
		return nil
	}
	ret, err := state.EvalExpression(ctx, rt.Value)
	if err != nil {
		return err
	}
	state.SetReturn(ret)
	return nil
}

func (state *EvalState) EvalSet(ctx context.Context, st *ast.SetStmt) error {
	val, err := state.EvalExpression(ctx, st.Value)
	if err != nil {
		return err
	}
	symb, err := state.EvalSymbol(ctx, st.Symbol)
	if err != nil {
		return err
	}
	if object.Is(symb.Kind(), object.VariableSymbol) {
		return fmt.Errorf("eval: can't assign new value to non variable symbol %s", symb.Name())
	}
	symbVar := symb.(*object.SymbolVariable)
	if !symbVar.ValueType.Compare(val.Type()) {
		return fmt.Errorf("eval: mismatch types of variable %s of %s and value of %s", symbVar.Identifier, symbVar.Value.Inspect(), val.Type().Inspect())
	}
	symbVar.Value = val
	return nil
}

func (state *EvalState) EvalIf(ctx context.Context, st *ast.IfStmt) error {
	type condStruct struct {
		obj  object.Litteral
		body []ast.Local
	}
	conditions := []condStruct{}

	cond, err := state.EvalExpression(ctx, st.Condition)
	if err != nil {
		return err
	}
	conditions = append(conditions, condStruct{
		obj:  cond,
		body: st.ThenBody,
	})

	for _, elif := range st.Elif {
		cond, err = state.EvalExpression(ctx, elif.Condition)
		if err != nil {
			return err
		}
		conditions = append(conditions, condStruct{
			obj:  cond,
			body: elif.Body,
		})
	}
	localContext := context.NewContext(ctx.Scope()+"_if", ctx)
	for _, cond := range conditions {
		if !object.Is(cond.obj.Kind(), object.BoolLitteral) {
			return fmt.Errorf("eval: condition should be bool value, got %s", cond.obj.Type().Inspect())
		}
		condBool := cond.obj.(*object.LitteralBool)
		if !condBool.Value {
			continue
		}
		err := state.EvalLocals(localContext, cond.body)
		if err != nil {
			return err
		}
		return nil
	}
	if len(st.ElseBody) != 0 {
		err := state.EvalLocals(localContext, st.ElseBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func (state *EvalState) EvalWhile(ctx context.Context, st *ast.WhileStmt) error {
	cond, err := state.EvalExpression(ctx, st.Condition)
	if err != nil {
		return err
	}
	if object.Is(cond.Kind(), object.BoolLitteral) {
		return fmt.Errorf("eval: condition should be bool value, got %s", cond.Type().Inspect())
	}
	localContext := context.NewContext(ctx.Scope()+"_while", ctx)
	if cond.(*object.LitteralBool).Value {
		for cond.(*object.LitteralBool).Value {
			err = state.EvalLocals(localContext, st.ThenBody)
			if err != nil {
				return err
			}
			cond, err = state.EvalExpression(ctx, st.Condition)
			if err != nil {
				return err
			}
		}
	} else {
		err = state.EvalLocals(localContext, st.ElseBody)
		if err != nil {
			return err
		}
	}
	return nil
}

func (state *EvalState) EvalList(ctx context.Context, lst *ast.ListAtom) (object.Litteral, error) {
	itemType, err := state.EvalType(ctx, lst.Type.Child)
	if err != nil {
		return nil, err
	}
	items := []object.Litteral{}
	for i, item := range lst.Items {
		obj, err := state.EvalExpression(ctx, item)
		if err != nil {
			return nil, err
		}
		if !itemType.Compare(obj.Type()) {
			return nil, fmt.Errorf("eval: wrong %dth list<%s> item type -> %s", i, itemType.Inspect(), obj.Type().Inspect())
		}
		items = append(items, obj)
	}
	return &object.LitteralList{ItemType: itemType, Items: items}, nil
}

// ================================================================

func (state *EvalState) NewConstantSymbol(ctx context.Context, cnst *ast.ConstantDecl) error {
	value, err := state.EvalAtom(ctx, cnst.Value)
	if err != nil {
		return err
	}
	return ctx.Insert(
		&object.SymbolConstant{
			Identifier: cnst.Identifier,
			Value:      value,
		},
	)
}

func (state *EvalState) NewVariableSymbol(ctx context.Context, vr *ast.VariableDecl) error {
	value, err := state.EvalExpression(ctx, vr.Value)
	if err != nil {
		return err
	}
	typ, err := state.EvalType(ctx, vr.Type)
	if err != nil {
		return err
	}
	return ctx.Insert(
		&object.SymbolVariable{
			Identifier: vr.Identifier,
			ValueType:  typ,
			Value:      value,
		},
	)
}

func (state *EvalState) NewFunctionSymbol(ctx context.Context, fn *ast.FunctionDecl) error {
	args := []object.DType{}
	for _, arg := range fn.Arguments {
		typ, err := state.EvalType(ctx, arg.Type)
		if err != nil {
			return err
		}
		args = append(args, typ)
	}
	var ret object.DType
	if fn.Return != nil {
		var err error
		ret, err = state.EvalType(ctx, fn.Return)
		if err != nil {
			return err
		}
	}
	return ctx.Insert(
		&object.SymbolFunction{
			Identifier: fn.Identifier,
			Fn:         fn,
			FunctionType: object.DTypeFunction{
				ArgumentsType: args,
				ReturnType:    ret,
			},
		},
	)
}
