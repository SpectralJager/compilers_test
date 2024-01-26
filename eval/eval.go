package eval

import (
	"fmt"
	"grimlang/ast"
	"grimlang/runtime"
)

func EvalModule(state EvalState, md *ast.Module, hash string) {
	switch md.Kind {
	case ":main":
		main := runtime.NewEnviroment("main", state.GetBuiltinEnv())
		state.InsertGlobalEnv(main)
		for _, gl := range md.Body {
			EvalGlobal(state, main, gl)
		}
		mainFn := main.SearchLocal("main")
		if mainFn == nil {
			panic("can't eval main module: should contains main function")
		}
		if mainFn.Kind() != runtime.SY_Function {
			panic("can't eval main module: main should be function")
		}
		EvalFunction(state, main, mainFn)
	case ":code":
		global := runtime.NewEnviroment(hash, state.GetBuiltinEnv())
		state.InsertGlobalEnv(global)
		for _, gl := range md.Body {
			EvalGlobal(state, global, gl)
		}
	default:
		panic("can't eval module: unexpected module kind")
	}
}

func EvalGlobal(state EvalState, env runtime.Enviroment, gl ast.Global) {
	switch gl := gl.(type) {
	case *ast.ConstantDecl:
		CreateConstant(state, env, gl)
	case *ast.FunctionDecl:
		CreateFunction(state, env, gl)
	case *ast.ImportDecl:
		EvalImport(state, env, gl)
	case *ast.RecordDefn:
		CreateRecord(state, env, gl)
	default:
		panic("can't eval global: unexpected global node")
	}
}

func EvalAtom(state EvalState, env runtime.Enviroment, atm ast.Atom) runtime.Litteral {
	switch atm := atm.(type) {
	case *ast.IntAtom:
		return runtime.NewIntLit(int64(atm.Value))
	case *ast.FloatAtom:
		return runtime.NewFloatLit(atm.Value)
	case *ast.BoolAtom:
		return runtime.NewBoolLit(atm.Value)
	case *ast.StringAtom:
		return runtime.NewStringLit(atm.Value[1 : len(atm.Value)-1])
	default:
		panic("can't eval atom: unexpected atom node")
	}
}

func EvalType(state EvalState, env runtime.Enviroment, typ ast.Type) runtime.Type {
	switch tp := typ.(type) {
	case *ast.IntType:
		return runtime.NewIntType()
	case *ast.BoolType:
		return runtime.NewBoolType()
	case *ast.FloatType:
		return runtime.NewFloatType()
	case *ast.StringType:
		return runtime.NewStringType()
	case *ast.ListType:
		return runtime.NewListType(EvalType(state, env, tp.Child))
	case *ast.RecordType:
		recSymb := EvalSymbol(state, env, tp.Symbol)
		return recSymb.Type()
	case nil:
		return runtime.NewVoidType()
	default:
		panic("can't eval type: unexpected type node")
	}
}

func EvalExpression(state EvalState, env runtime.Enviroment, exp ast.Expression) runtime.Litteral {
	switch exp := exp.(type) {
	case ast.Atom:
		return EvalAtom(state, env, exp)
	case *ast.SymbolExpr:
		return EvalSymbol(state, env, exp).Value()
	case *ast.SymbolCall:
		EvalSymbolCall(state, env, exp)
		return state.GetReturn()
	case *ast.NewExpr:
		return EvalNew(state, env, exp)
	default:
		panic("can't eval expression: unexpected expression node")
	}
}

func EvalLocal(state EvalState, env runtime.Enviroment, lc ast.Local) {
	switch lc := lc.(type) {
	case *ast.ConstantDecl:
		CreateConstant(state, env, lc)
	case *ast.VariableDecl:
		CreateVariable(state, env, lc)
	case *ast.SymbolCall:
		EvalSymbolCall(state, env, lc)
	case *ast.SetStmt:
		EvalSet(state, env, lc)
	case *ast.IfStmt:
		EvalIf(state, env, lc)
	case *ast.WhileStmt:
		EvalWhile(state, env, lc)
	case *ast.ReturnStmt:
		EvalReturn(state, env, lc)
	}
}

func EvalWhile(state EvalState, env runtime.Enviroment, stm *ast.WhileStmt) {
	cond := EvalExpression(state, env, stm.Condition)
	if cond.Kind() != runtime.LI_Bool {
		panic("can't eval if statement: condition should be bool")
	}
	for cond.ValueBool() {
		local := runtime.NewEnviroment(env.Name()+"_while", env)
		for _, lc := range stm.ThenBody {
			EvalLocal(state, local, lc)
			if state.IsReturn() {
				return
			}
		}
		cond = EvalExpression(state, env, stm.Condition)
		if cond.Kind() != runtime.LI_Bool {
			panic("can't eval if statement: condition should be bool")
		}
	}
	local := runtime.NewEnviroment(env.Name()+"_while", env)
	for _, lc := range stm.ElseBody {
		EvalLocal(state, local, lc)
		if state.IsReturn() {
			return
		}
	}
}

func EvalIf(state EvalState, env runtime.Enviroment, stm *ast.IfStmt) {
	local := runtime.NewEnviroment(env.Name()+"_if", env)
	cond := EvalExpression(state, local, stm.Condition)
	if cond.Kind() != runtime.LI_Bool {
		panic("can't eval if statement: condition should be bool")
	}
	if cond.ValueBool() {
		for _, lc := range stm.ThenBody {
			EvalLocal(state, local, lc)
			if state.IsReturn() {
				break
			}
		}
		return
	}
	for _, elif := range stm.Elif {
		cond = EvalExpression(state, env, elif.Condition)
		if cond.Kind() != runtime.LI_Bool {
			panic("can't eval if statement: condition should be bool")
		}
		if cond.ValueBool() {
			for _, lc := range elif.Body {
				EvalLocal(state, local, lc)
				if state.IsReturn() {
					break
				}
			}
			return
		}
	}
	for _, lc := range stm.ElseBody {
		EvalLocal(state, local, lc)
		if state.IsReturn() {
			break
		}
	}
	return
}

func EvalNew(state EvalState, env runtime.Enviroment, nw *ast.NewExpr) runtime.Litteral {
	tp := EvalType(state, env, nw.Type)
	items := []runtime.Litteral{}
	for _, item := range nw.Items {
		items = append(items, EvalExpression(state, env, item))
	}
	switch tp.Kind() {
	case runtime.TY_Record:
		fields := []runtime.Symbol{}
		for i, item := range items {
			fld := tp.FieldByIndex(i)
			fields = append(fields,
				runtime.NewVariable(
					fld.Name(),
					fld.Type(),
					item,
				),
			)
		}
		return runtime.NewRecordLit(tp, fields...)
	case runtime.TY_List:
		return runtime.NewListLit(tp.Item(), items...)
	default:
		panic("can't eval new expression: unexpected type kind")
	}
}

func EvalReturn(state EvalState, env runtime.Enviroment, ret *ast.ReturnStmt) {
	if ret.Value == nil {
		state.SetReturn(runtime.NewIntLit(0))
	} else {
		state.SetReturn(EvalExpression(state, env, ret.Value))
	}
}

func EvalSet(state EvalState, env runtime.Enviroment, st *ast.SetStmt) {
	sm := EvalSymbol(state, env, st.Symbol)
	val := EvalExpression(state, env, st.Value)
	sm.Set(val)
}

func EvalSymbolCall(state EvalState, env runtime.Enviroment, sc *ast.SymbolCall) {
	args := []runtime.Litteral{}
	for _, arg := range sc.Arguments {
		args = append(args, EvalExpression(state, env, arg))
	}
	fn := EvalSymbol(state, env, sc.Call)
	if state.IsSwitchEnv() {
		env = state.GetSwitchEnv()
	}
	switch fn.Kind() {
	case runtime.SY_Builtin:
		state.SetReturn(fn.Builtin()(args...))
	case runtime.SY_Function:
		EvalFunction(state, env, fn, args...)
	}
}

func EvalSymbol(state EvalState, env runtime.Enviroment, sm *ast.SymbolExpr) runtime.Symbol {
	if sm.Next != nil {
		state.SetSymbolFlag()
		s := env.Search(sm.Identifier)
		switch s.Kind() {
		case runtime.SY_Import:
			env := state.SearchGlobalEnv(s.Path())
			state.SetSwitchEnv(env)
			return EvalSymbol(state, env, sm.Next)
		case runtime.SY_Variable, runtime.SY_Constant:
			if s.Value().Type().Kind() == runtime.TY_Record {
				env := runtime.NewEnviromentFromRecord(s.Value())
				return EvalSymbol(state, env, sm.Next)
			}
			fallthrough
		default:
			panic("can't eval symbol: unexpected symbol kind")
		}
	}
	if state.GetSymbolFlag() {
		state.ClrSymbolFlag()
		s := env.SearchLocal(sm.Identifier)
		if s == nil {
			panic("can't eval symbol: symbol not found")
		}
		return s
	} else {
		s := env.Search(sm.Identifier)
		if s == nil {
			panic("can't eval symbol: symbol not found")
		}
		return s
	}
}

func EvalImport(state EvalState, env runtime.Enviroment, im *ast.ImportDecl) {
	path := EvalAtom(state, env, im.Path)
	module, hash := CreateModuleFromFile(path.ValueString())
	EvalModule(state, module, hash)
	env.Insert(
		runtime.NewImport(
			im.Ident,
			hash,
		),
	)
}

func EvalFunction(state EvalState, env runtime.Enviroment, fn runtime.Symbol, args ...runtime.Litteral) {
	local := runtime.NewEnviroment(fn.Name(), env)
	if len(args) != fn.Type().NumIns() {
		panic("can't eval function: mismatched number of inputs")
	}
	for i, arg := range args {
		if !fn.Type().In(i).Compare(arg.Type()) {
			panic("can't eval function: argument type mismatch")
		}
		local.Insert(runtime.NewConstant(
			fn.Fn().Arguments[i].Identifier,
			arg,
		))
	}
	defer func() {
		fmt.Println(local.String())
		if fn.Type().Out().Compare(runtime.NewVoidType()) {
			state.GetReturn()
			return
		}
		if !state.IsReturn() {
			panic("can't eval function: expected output")
		}
		if !fn.Type().Out().Compare(state.GetReturnType()) {
			panic("can't eval function: output type mismatched")
		}
	}()

	for _, lc := range fn.Fn().Body {
		EvalLocal(state, local, lc)
		if state.IsReturn() {
			return
		}
	}
}

func CreateConstant(state EvalState, env runtime.Enviroment, cnst *ast.ConstantDecl) {
	val := EvalAtom(state, env, cnst.Value)
	env.Insert(runtime.NewConstant(cnst.Identifier, val))
}

func CreateVariable(state EvalState, env runtime.Enviroment, vr *ast.VariableDecl) {
	typ := EvalType(state, env, vr.Type)
	val := EvalExpression(state, env, vr.Value)
	env.Insert(
		runtime.NewVariable(
			vr.Identifier,
			typ,
			val,
		),
	)
}

func CreateFunction(state EvalState, env runtime.Enviroment, fn *ast.FunctionDecl) {
	args := []runtime.Type{}
	for _, arg := range fn.Arguments {
		args = append(args, EvalType(state, env, arg.Type))
	}
	env.Insert(
		runtime.NewFunction(
			fn.Identifier,
			runtime.NewFunctionType(
				EvalType(state, env, fn.Return),
				args...,
			),
			fn,
		),
	)
}

func CreateRecord(state EvalState, env runtime.Enviroment, rec *ast.RecordDefn) {
	fields := []runtime.FieldType{}
	for _, fld := range rec.Fields {
		fields = append(fields, runtime.NewFieldType(
			fld.Identifier,
			EvalType(state, env, fld.Type),
		))
	}
	env.Insert(
		runtime.NewRecord(
			runtime.NewRecordType(
				rec.Identifier,
				fields...,
			),
		),
	)
}
