package asm

import (
	"fmt"
	"strings"
)

type Program struct {
	Functions map[string]Function
}

func NewProgram(fns ...Function) (*Program, error) {
	prog := &Program{
		Functions: map[string]Function{},
	}
	for _, fn := range fns {
		_, ok := prog.Functions[fn.Ident]
		if ok {
			return prog, fmt.Errorf("function '%s' alredy exists", fn.Ident)
		}
		prog.Functions[fn.Ident] = fn
	}
	return prog, nil
}

func (pr *Program) InspectIndent(indent int) string {
	fns := []string{}
	for _, fn := range pr.Functions {
		fns = append(fns, fn.InspectIndent(indent))
	}
	return strings.Join(fns, "\n")
}
