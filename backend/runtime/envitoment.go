package runtime

import (
	"fmt"
	"grimlang/backend/asm"
)

type Register struct {
	Ident string
	Value asm.Value
}

type Enviroment []*Register

func NewEnviroment(vars asm.Vars) Enviroment {
	env := make(Enviroment, len(vars))
	for i := 0; i < len(vars); i++ {
		val := vars[i]
		ident := val.Ident
		val.Ident = ""
		env[i] = &Register{
			Ident: ident,
			Value: val,
		}
	}
	return env
}

func (env Enviroment) Set(ident string, value asm.Value) error {
	var reg *Register = nil
	for _, r := range env {
		if r.Ident == ident {
			reg = r
			break
		}
	}
	if reg == nil {
		panic(fmt.Errorf("symbol '%s' not exists", ident))
	}
	err := reg.Value.Compare(value)
	if err != nil {
		panic(err)
	}
	reg.Value = value
	return nil
}

func (env Enviroment) Get(ident string) asm.Value {
	for _, r := range env {
		if r.Ident == ident {
			return r.Value
		}
	}
	panic(fmt.Errorf("symbol '%s' not exists", ident))
}
