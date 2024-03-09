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
	env := make(Enviroment, 0, len(vars))
	for _, val := range vars {
		ident := val.Ident
		val.Ident = ""
		env = append(env, &Register{
			Ident: ident,
			Value: val,
		})
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
		return fmt.Errorf("symbol '%s' not exists", ident)
	}
	err := reg.Value.Compare(value)
	if err != nil {
		return err
	}
	reg.Value = value
	return nil
}

func (env Enviroment) Get(ident string) (asm.Value, error) {
	for _, r := range env {
		if r.Ident == ident {
			return r.Value, nil
		}
	}
	return asm.Value{}, fmt.Errorf("symbol '%s' not exists", ident)
}
