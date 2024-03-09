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
