package builtin

import (
	"grimlang/runtime"
	"os"
)

func NewBuiltinEnv() runtime.Environment {
	env := runtime.NewEnvironment("builtin", nil)
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
	env.Insert(
		runtime.NewImport(
			"io",
			"io",
		),
	)
	env.Insert(
		runtime.NewImport(
			"list",
			"list",
		),
	)
	env.Insert(
		runtime.NewImport(
			"string",
			"string",
		),
	)
	env.Insert(
		runtime.NewImport(
			"float",
			"float",
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
