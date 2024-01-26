package builtin_int

import "grimlang/runtime"

func NewBuiltinIntEnv() runtime.Enviroment {
	env := runtime.NewEnviroment("int", nil)
	env.Insert(
		runtime.NewBuiltin(
			"add",
			runtime.NewFunctionType(
				runtime.NewIntType(),
				runtime.NewVariaticType(runtime.NewIntType()),
			),
			Add,
		),
	)
	return env
}

func Add(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) == 0 {
		panic("int/add: expect atleast 1 input")
	}
	res := int64(0)
	for _, in := range inputs {
		res += in.ValueInt()
	}
	return runtime.NewIntLit(res)
}
