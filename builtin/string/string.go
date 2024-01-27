package builtin_string

import (
	"grimlang/runtime"
	"strings"
)

func NewBuiltinStringEnv() runtime.Enviroment {
	env := runtime.NewEnviroment("string", nil)
	env.Insert(
		runtime.NewBuiltin(
			"format",
			runtime.NewFunctionType(
				runtime.NewStringType(),
				runtime.NewStringType(),
				runtime.NewVariaticType(runtime.NewStringType()),
			),
			Format,
		),
	)
	return env
}

func Format(inputs ...runtime.Litteral) runtime.Litteral {
	if len(inputs) <= 1 {
		panic("string/format: expect atleast 2 inputs")
	}
	format := inputs[0].ValueString()
	if strings.Count(format, "$$") != len(inputs[1:]) {
		panic("string/format: mismatched number of format entries and inputs")
	}
	for _, input := range inputs[1:] {
		format = strings.Replace(format, "$$", input.ValueString(), 1)
	}
	return runtime.NewStringLit(format)
}
