package runtime

import (
	"fmt"
	"grimlang/backend/asm"
	"maps"
)

type Enviroment map[string]asm.Value

func NewEnviroment(vars asm.Vars) Enviroment {
	// env := Enviroment{}
	// for k, v := range vars {
	// 	env[k] = v
	// }
	return Enviroment(maps.Clone(vars))
}

func (env Enviroment) Set(ident string, value asm.Value) error {
	vl, ok := env[ident]
	if !ok {
		return fmt.Errorf("symbol '%s' not exists", ident)
	}
	err := vl.Compare(value)
	if err != nil {
		return err
	}
	env[ident] = value
	return nil
}

func (env Enviroment) Get(ident string) (asm.Value, error) {
	vl, ok := env[ident]
	if !ok {
		return asm.Value{}, fmt.Errorf("symbol '%s' not exists", ident)
	}
	return vl, nil
}
