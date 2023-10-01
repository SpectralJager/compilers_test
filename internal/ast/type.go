package ast

import (
	"fmt"
	"strings"
)

type TypeAST struct {
	Primary SymbolAST `parser:"@@"`
	Generic []TypeAST `parser:"('<' @@+ '>')?"`
}

func (t TypeAST) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s", &t.Primary)
	if t.Generic != nil {
		fmt.Fprint(&buf, "< ")
		for _, g := range t.Generic {
			fmt.Fprintf(&buf, "%s ", &g)
		}
		fmt.Fprint(&buf, ">")
	}
	return buf.String()
}
