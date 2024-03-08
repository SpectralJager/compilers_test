package asm

import (
	"fmt"
	"strings"
)

const (
	vrsStr = "(vars\n%s%s)"
)

type Vars map[string]Value

func (vr Vars) InspectIndent(indent int) string {
	indentStr := strings.Repeat(" ", indent)
	vars := []string{}
	for k, v := range vr {
		vars = append(vars, fmt.Sprintf("[%s %s]", k, v.Type.Inspect()))
	}
	return fmt.Sprintf(vrsStr, indentStr, strings.Join(vars, "\n"+indentStr))
}

func (vr Vars) Set(ident string, value Value) error {
	vl, ok := vr[ident]
	if !ok {
		return fmt.Errorf("variable '%s' not exists", ident)
	}
	err := vl.Compare(value)
	if err != nil {
		return err
	}
	vr[ident] = value
	return nil
}

func (vr Vars) Get(ident string) (Value, error) {
	vl, ok := vr[ident]
	if !ok {
		return Value{}, fmt.Errorf("variable '%s' not exists", ident)
	}
	return vl, nil
}
