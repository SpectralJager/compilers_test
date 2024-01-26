package main

import (
	"fmt"
	"grimlang/builtin"
	builtin_int "grimlang/builtin/int"
	"grimlang/eval"
)

func main() {
	module, hash := eval.CreateModuleFromFile("examples/test.grim")

	state := eval.NewEvalState(builtin.NewBuiltinEnv())
	state.InsertGlobalEnv(builtin_int.NewBuiltinIntEnv())

	defer func() {
		fmt.Println(state.String())
	}()

	eval.EvalModule(state, module, hash)
}
