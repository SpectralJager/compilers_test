package eval

import (
	"fmt"
)

type Env map[string]interface{}

func (env *Env) Lookup(symbol string) (interface{}, error) {
	result, ok := (*env)[symbol]
	if !ok {
		return nil, fmt.Errorf("symbol %s undefined", symbol)
	}
	return result, nil
}

func (env *Env) Set(symbol string, value interface{}) error {
	(*env)[symbol] = value
	return nil
}
