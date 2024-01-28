package builtin_io

import (
	"fmt"
	"grimlang/runtime"
	"strings"
)

func NewBuiltinIOEnv() runtime.Environment {
	env := runtime.NewEnvironment("io", nil)
	env.Insert(
		runtime.NewBuiltin(
			"println",
			runtime.NewFunctionType(
				runtime.NewVoidType(),
				runtime.NewVariaticType(runtime.NewStringType()),
			),
			Println,
		),
	)
	return env
}

func Println(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) == 0 {
		panic("io/println: expect atleast 1 input")
	}
	msgs := []string{}
	for _, in := range inputs {
		msgs = append(msgs, in.ValueString())
	}
	fmt.Println(strings.Join(msgs, ""))
	return runtime.NewIntLit(0)
}
