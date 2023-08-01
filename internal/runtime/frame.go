package runtime

import (
	"bytes"
	"fmt"
	"gl/internal/ir"
)

/*
name: Name
stack:
	...
varibles:
	a -> int:23
code:
	...
*/

type Frame struct {
	Name     string
	Varibles []Object
	Code     *ir.Code
	Bp       int
	Ip       int
}

func NewFrame(name string, varibles []ir.ISymbolDef, code *ir.Code, bp int) Frame {
	vars := make([]Object, len(varibles))
	return Frame{
		Name:     name,
		Varibles: vars,
		Code:     code,
		Bp:       bp,
	}
}

func (f *Frame) Instruction() byte {
	return (*f.Code)[f.Ip]
}

func (f *Frame) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "= Frame: %s\n", f.Name)
	fmt.Fprintf(&buf, "base pointer: %d\n", f.Bp)
	fmt.Fprint(&buf, "varibles:\n")
	for i, v := range f.Varibles {
		if v != nil {
			fmt.Fprintf(&buf, "\t%d -> %s\n", i, v.String())
		}
	}
	fmt.Fprint(&buf, "===================\n")
	return buf.String()
}
