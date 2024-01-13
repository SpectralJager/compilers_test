package runtime

import (
	"errors"
	"fmt"
)

type Enviroment struct {
	Current *Symbol
	Parent  *Enviroment
}

func NewEnviroment(current *Symbol, parent *Enviroment) (*Enviroment, error) {
	if current.Kind != SK_Module && current.Kind != SK_Record {
		return nil, errors.New("can't create new enviroment, current symbol should be module or record")
	}
	return &Enviroment{
		Current: current,
		Parent:  parent,
	}, nil
}

func (env *Enviroment) SearchSymbolInCurrent(ident string) *Symbol {
	for _, symb := range env.Current.Items {
		if symb.Identifier == ident {
			return symb
		}
	}
	return nil
}

func (env *Enviroment) AppendSymbolInCurrent(symb *Symbol) error {
	if env.SearchSymbolInCurrent(symb.Identifier) != nil {
		return fmt.Errorf("can't append symbol '%s' to enviroment '%s', already exists", symb.Identifier, env.Current.Identifier)
	}
	env.Current.Items = append(env.Current.Items, symb)
	return nil
}

func (env *Enviroment) SearchSymbol(ident string) *Symbol {
	if symb := env.SearchSymbolInCurrent(ident); symb != nil {
		return symb
	}
	if env.Parent != nil {
		return env.Parent.SearchSymbol(ident)
	}
	return nil
}
