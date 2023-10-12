package ast

import (
	"fmt"
	"strings"
)

type SCallAST struct {
	Function  SymbolAST `parser:"'(' @@"`
	Arguments []EXPR    `parser:"@@* ')'"`
}

func (a *SCallAST) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "(%s", &a.Function)
	for _, arg := range a.Arguments {
		fmt.Fprintf(&buf, " %s", arg)
	}
	fmt.Fprint(&buf, ")")
	return buf.String()
}

// part of ...
func (a *SCallAST) expr() {}
func (a *SCallAST) locl() {}
func (a *SCallAST) ast()  {}
