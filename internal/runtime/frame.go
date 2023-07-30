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
	Varibles map[string]Object
	Code     []ir.IInstruction
	Bp       int
	Ip       int
}

func NewFrame(name string, varibles map[string]ir.ISymbolDef, code []ir.IInstruction, bp int) Frame {
	vars := make(map[string]Object)
	for k := range varibles {
		vars[k] = nil
	}
	return Frame{
		Name:     name,
		Varibles: vars,
		Code:     code,
		Bp:       bp,
	}
}

func (f *Frame) Instruction() ir.IInstruction {
	return f.Code[f.Ip]
}

func (f *Frame) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "= Frame: %s\n", f.Name)
	fmt.Fprintf(&buf, "base pointer: %d\n", f.Bp)
	fmt.Fprint(&buf, "varibles:\n")
	for k, v := range f.Varibles {
		if v != nil {
			fmt.Fprintf(&buf, "\t%s -> %s\n", k, v.String())
		}
	}
	fmt.Fprint(&buf, "===================\n")
	return buf.String()
}
