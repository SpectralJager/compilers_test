package builtin_float

import (
	"fmt"
	"grimlang/runtime"
)

func NewBuiltinFloatEnv() runtime.Environment {
	env := runtime.NewEnvironment("float", nil)
	env.Insert(
		runtime.NewBuiltin(
			"add",
			runtime.NewFunctionType(
				runtime.NewFloatType(),
				runtime.NewVariaticType(runtime.NewFloatType()),
			),
			Add,
		),
	)
	env.Insert(
		runtime.NewBuiltin(
			"sub",
			runtime.NewFunctionType(
				runtime.NewFloatType(),
				runtime.NewVariaticType(runtime.NewFloatType()),
			),
			Sub,
		),
	)
	env.Insert(
		runtime.NewBuiltin(
			"lt",
			runtime.NewFunctionType(
				runtime.NewBoolType(),
				runtime.NewVariaticType(runtime.NewFloatType()),
			),
			Lt,
		),
	)
	env.Insert(
		runtime.NewBuiltin(
			"toString",
			runtime.NewFunctionType(
				runtime.NewStringType(),
				runtime.NewFloatType(),
			),
			ToString,
		),
	)
	return env
}

func Add(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) == 0 {
		panic("float/add: expect atleast 1 input")
	}
	res := float64(0)
	for _, in := range inputs {
		res += in.ValueFloat()
	}
	return runtime.NewFloatLit(res)
}

func Sub(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) == 0 {
		panic("float/sub: expect atleast 1 input")
	}
	res := inputs[0].ValueFloat()
	for _, in := range inputs[1:] {
		res -= in.ValueFloat()
	}
	return runtime.NewFloatLit(res)
}

func Lt(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) <= 1 {
		panic("float/lt: expect atleast 2 input")
	}
	res := inputs[0].ValueFloat()
	for _, in := range inputs[1:] {
		if res >= in.ValueFloat() {
			return runtime.NewBoolLit(false)
		}
		res = in.ValueFloat()
	}
	return runtime.NewBoolLit(true)
}

func ToString(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) != 1 {
		panic("float/toString: expect 1 input")
	}
	return runtime.NewStringLit(
		fmt.Sprintf("%f", inputs[0].ValueFloat()),
	)
}
