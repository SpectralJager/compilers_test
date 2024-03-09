package asm

import (
	"fmt"
	"strings"
)

type Program struct {
	Functions []Function
}

func NewProgram(fns ...Function) (*Program, error) {
	prog := &Program{
		Functions: make([]Function, 0, len(fns)),
	}
	for _, fn := range fns {
		_, err := prog.Function(fn.Ident)
		if err == nil {
			return prog, fmt.Errorf("function '%s' alredy exists", fn.Ident)
		}
		prog.Functions = append(prog.Functions, fn)
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

func (pr *Program) Function(ident string) (Function, error) {
	for _, fn := range pr.Functions {
		if fn.Ident == ident {
			return fn, nil
		}
	}
	return Function{}, fmt.Errorf("function '%s' not exists", ident)
}
