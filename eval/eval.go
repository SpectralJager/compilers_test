package runtime

import (
	"fmt"
	"grimlang/ast"
	"grimlang/context"
	"grimlang/dtype"
	"grimlang/object"
	"grimlang/symbol"
)

type EvalState struct {
	GlobalContext context.Context
	Return        object.Object
	IsReturn      bool
}

func (state *EvalState) SetReturn(obj object.Object) {
	state.IsReturn = true
	state.Return = obj
}

func (state *EvalState) GetReturn() object.Object {
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
		if mainSymbol.Kind() != symbol.Function {
			return fmt.Errorf("main should be function")
		}
		mainFunc := mainSymbol.(*symbol.FunctionSymbol)
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

func (state *EvalState) EvalExpression(ctx context.Context, ex ast.Expression) (object.Object, error) {
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

func (state *EvalState) EvalAtom(ctx context.Context, at ast.Atom) (object.Object, error) {
	switch at := at.(type) {
	case *ast.IntAtom:
		return &object.IntObject{Value: at.Value}, nil
	case *ast.BoolAtom:
		return &object.BoolObject{Value: at.Value}, nil
	case *ast.StringAtom:
		return &object.StringObject{Value: at.Value[1 : len(at.Value)-1]}, nil
	case *ast.FloatAtom:
		return &object.FloatObject{Value: at.Value}, nil
	case *ast.ListAtom:
		return state.EvalList(ctx, at)
	default:
		return nil, fmt.Errorf("eval: unexpected atom %T", at)
	}
}

func (state *EvalState) EvalType(ctx context.Context, tp ast.Type) (dtype.Type, error) {
	switch tp := tp.(type) {
	case *ast.IntType:
		return &dtype.IntType{}, nil
	case *ast.BoolType:
		return &dtype.BoolType{}, nil
	case *ast.StringType:
		return &dtype.StringType{}, nil
	case *ast.FloatType:
		return &dtype.FloatType{}, nil
	case *ast.ListType:
		itemType, err := state.EvalType(ctx, tp.Child)
		if err != nil {
			return nil, err
		}
		return &dtype.ListType{Child: itemType}, nil
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

func (state *EvalState) EvalSymbol(ctx context.Context, sm *ast.SymbolExpr) (symbol.Symbol, error) {
	symb := ctx.Search(sm.Identifier)
	if symb == nil {
		return nil, fmt.Errorf("eval: symbol %s not found", sm.Identifier)
	}
	if sm.Next == nil {
		return symb, nil
	}
	switch symb := symb.(type) {
	case *symbol.ModuleSymbol:
		return state.EvalSymbol(symb, sm.Next)
	default:
		return nil, fmt.Errorf("eval: symbol %s not found", sm.Next.Identifier)
	}
}

func (state *EvalState) EvalSymbolExpression(ctx context.Context, sm *ast.SymbolExpr) (object.Object, error) {
	symb, err := state.EvalSymbol(ctx, sm)
	if err != nil {
		return nil, err
	}
	switch symb := symb.(type) {
	case *symbol.ConstantSymbol:
		return symb.Value, nil
	case *symbol.VariableSymbol:
		return symb.Value, nil
	default:
		return nil, fmt.Errorf("eval: can't get object from symbol %s of %T", symb.Name(), symb)
	}
}

func (state *EvalState) EvalFunction(ctx context.Context, fn *ast.FunctionDecl, args []object.Object) (object.Object, error) {
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
			return nil, fmt.Errorf("eval: %s expect %dth argument be %s, got %s", fn.Identifier, i, argType.Name(), args[i].Type().Name())
		}
		err = localContext.Insert(
			&symbol.VariableSymbol{
				Identifier: arg.Identifier,
				Type:       argType,
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
			return nil, fmt.Errorf("eval: function expect return %s, got nothing", retType.Name())
		}
		state.Return = nil
		if !retType.Compare(ret.Type()) {
			return nil, fmt.Errorf("eval: function expect return %s, got %s", retType.Name(), ret.Type().Name())
		}
		return ret, nil
	case fn.Return != nil && !state.IsReturn:
		retType, err := state.EvalType(ctx, fn.Return)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("eval: function expect return %s, got nothing", retType.Name())
	case state.IsReturn:
		state.Return = nil
	}
	return nil, nil
}

func (state *EvalState) EvalSymbolCall(ctx context.Context, sc *ast.SymbolCall) (object.Object, error) {
	tempSymb, err := state.EvalSymbol(ctx, sc.Call)
	if err != nil {
		return nil, err
	}
	params := []object.Object{}
	for _, param := range sc.Arguments {
		obj, err := state.EvalExpression(ctx, param)
		if err != nil {
			return nil, err
		}
		params = append(params, obj)
	}
	switch tempSymb := tempSymb.(type) {
	case *symbol.FunctionSymbol:
		return state.EvalFunction(ctx, tempSymb.Fn, params)
	case *symbol.BuiltinFunctionSymbol:
		return tempSymb.Callee(params...)
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
	if symb.Kind() != symbol.Variable {
		return fmt.Errorf("eval: can't assign new value to non variable symbol %s", symb.Name())
	}
	symbVar := symb.(*symbol.VariableSymbol)
	if !symbVar.Type.Compare(val.Type()) {
		return fmt.Errorf("eval: mismatch types of variable %s of %s and value of %s", symbVar.Identifier, symbVar.Type.Name(), val.Type().Name())
	}
	symbVar.Value = val
	return nil
}

func (state *EvalState) EvalIf(ctx context.Context, st *ast.IfStmt) error {
	type condStruct struct {
		obj  object.Object
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
		if cond.obj.Kind() != object.Bool {
			return fmt.Errorf("eval: condition should be bool value, got %s", cond.obj.Type().Name())
		}
		condBool := cond.obj.(*object.BoolObject)
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
	if cond.Kind() != object.Bool {
		return fmt.Errorf("eval: condition should be bool value, got %s", cond.Type().Name())
	}
	localContext := context.NewContext(ctx.Scope()+"_while", ctx)
	if cond.(*object.BoolObject).Value {
		for cond.(*object.BoolObject).Value {
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

func (state *EvalState) EvalList(ctx context.Context, lst *ast.ListAtom) (object.Object, error) {
	itemType, err := state.EvalType(ctx, lst.Type.Child)
	if err != nil {
		return nil, err
	}
	items := []object.Object{}
	for i, item := range lst.Items {
		obj, err := state.EvalExpression(ctx, item)
		if err != nil {
			return nil, err
		}
		if !itemType.Compare(obj.Type()) {
			return nil, fmt.Errorf("eval: wrong %dth list<%s> item type -> %s", i, itemType.Name(), obj.Type().Name())
		}
		items = append(items, obj)
	}
	return &object.ListObject{ChildType: itemType, Items: items}, nil
}

// ================================================================

func (state *EvalState) NewConstantSymbol(ctx context.Context, cnst *ast.ConstantDecl) error {
	value, err := state.EvalAtom(ctx, cnst.Value)
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
		&symbol.VariableSymbol{
			Identifier: vr.Identifier,
			Type:       typ,
			Value:      value,
		},
	)
}

func (state *EvalState) NewFunctionSymbol(ctx context.Context, fn *ast.FunctionDecl) error {
	args := []dtype.Type{}
	for _, arg := range fn.Arguments {
		typ, err := state.EvalType(ctx, arg.Type)
		if err != nil {
			return err
		}
		args = append(args, typ)
	}
	var ret dtype.Type
	if fn.Return != nil {
		var err error
		ret, err = state.EvalType(ctx, fn.Return)
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
