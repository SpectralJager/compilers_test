package runtime

import (
	"fmt"
	"grimlang/backend/asm"
)

type Enviroment [16]asm.Value

func (env *Enviroment) Set(index int64, value asm.Value) {
	if index >= int64(len(env)) {
		panic(fmt.Errorf("register '%d' not exists", index))
	}
	env[index] = value
}

func (env *Enviroment) Get(index int64) asm.Value {
	if index >= int64(len(env)) {
		panic(fmt.Errorf("register '%d' not exists", index))
	}
	return env[index]
}
