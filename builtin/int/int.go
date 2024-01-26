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
	env.Insert(
		runtime.NewBuiltin(
			"sub",
			runtime.NewFunctionType(
				runtime.NewIntType(),
				runtime.NewVariaticType(runtime.NewIntType()),
			),
			Sub,
		),
	)
	env.Insert(
		runtime.NewBuiltin(
			"lt",
			runtime.NewFunctionType(
				runtime.NewBoolType(),
				runtime.NewVariaticType(runtime.NewIntType()),
			),
			Lt,
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

func Sub(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) == 0 {
		panic("int/sub: expect atleast 1 input")
	}
	res := inputs[0].ValueInt()
	for _, in := range inputs[1:] {
		res -= in.ValueInt()
	}
	return runtime.NewIntLit(res)
}

func Lt(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) <= 1 {
		panic("int/lt: expect atleast 2 input")
	}
	res := inputs[0].ValueInt()
	for _, in := range inputs[1:] {
		if res >= in.ValueInt() {
			return runtime.NewBoolLit(false)
		}
		res = in.ValueInt()
	}
	return runtime.NewBoolLit(true)
}
