package builtin

import (
	"grimlang/runtime"
	"os"
)

func NewBuiltinEnv() runtime.Enviroment {
	env := runtime.NewEnviroment("builtin", nil)
	env.Insert(
		runtime.NewBuiltin(
			"exit",
			runtime.NewFunctionType(
				runtime.NewVoidType(),
				runtime.NewIntType(),
			),
			Exit,
		),
	)
	env.Insert(
		runtime.NewImport(
			"int",
			"int",
		),
	)
	return env
}

func Exit(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) != 1 {
		panic("builtin(exit): expect 1 input")
	}
	code := inputs[0].ValueInt()
	os.Exit(int(code))
	return runtime.NewIntLit(0)
}
