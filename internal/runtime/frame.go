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
	Stack    Stack
	Code     []ir.IInstruction
}

func NewFrame(name string, varibles map[string]ir.ISymbolDef, code []ir.IInstruction) *Frame {
	vars := make(map[string]Object)
	for k := range varibles {
		vars[k] = nil
	}
	return &Frame{
		Name:     name,
		Varibles: vars,
		Code:     code,
	}
}

func (f *Frame) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "= Frame: %s\n", f.Name)
	fmt.Fprint(&buf, "varibles:\n")
	for k, v := range f.Varibles {
		if v != nil {
			fmt.Fprintf(&buf, "\t%s -> %s\n", k, v.String())
		}
	}
	fmt.Fprintf(&buf, "%s", f.Stack.StackTrace())
	fmt.Fprint(&buf, "===================\n")
	return buf.String()
}
