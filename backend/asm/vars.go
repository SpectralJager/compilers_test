package asm

import (
	"fmt"
	"strings"
)

const (
	vrsStr = "(vars\n%s%s)"
)

type Vars []Value

func Var(ident string, value Value) Value {
	value.Ident = ident
	return value
}

func NewVars(vars ...Value) Vars {
	return vars
}

func (vr Vars) InspectIndent(indent int) string {
	indentStr := strings.Repeat(" ", indent)
	vars := []string{}
	for _, v := range vr {
		vars = append(vars, fmt.Sprintf("[%s %s]", v.Ident, v.Type.Inspect()))
	}
	return fmt.Sprintf(vrsStr, indentStr, strings.Join(vars, "\n"+indentStr))
}
