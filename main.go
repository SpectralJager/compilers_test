package main

import (
	"grimlang/builtin"
	builtin_float "grimlang/builtin/float"
	builtin_int "grimlang/builtin/int"
	builtin_io "grimlang/builtin/io"
	builtin_list "grimlang/builtin/list"
	builtin_string "grimlang/builtin/string"
	"grimlang/eval"
)

func main() {
	module, hash := eval.CreateModuleFromFile("examples/test.grim")

	state := eval.NewEvalState(builtin.NewBuiltinEnv())
	state.InsertGlobalEnv(builtin_int.NewBuiltinIntEnv())
	state.InsertGlobalEnv(builtin_io.NewBuiltinIOEnv())
	state.InsertGlobalEnv(builtin_list.NewBuiltinListEnv())
	state.InsertGlobalEnv(builtin_string.NewBuiltinStringEnv())
	state.InsertGlobalEnv(builtin_float.NewBuiltinFloatEnv())

	defer func() {
		// fmt.Println()
		// fmt.Println(state.String())
	}()

	eval.EvalModule(state, module, hash)
}
