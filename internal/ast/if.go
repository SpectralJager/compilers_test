package ast

import (
	"fmt"
	"strings"
)

type IfAST struct {
	IfCondition EXPR    `parser:"'@if' @@ "`
	IfBody      []LOCAL `parser:"'{' @@+ "`
	Elif        []struct {
		ElifCondition EXPR    `parser:"@@ '=>'"`
		ElifBody      []LOCAL `parser:" @@+"`
	} `parser:"('elif' @@)*"`
	ElseBody []LOCAL `parser:"('else' '=>' @@+)? '}'"`
}

func (f *IfAST) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "@if %s {\n", f.IfCondition)
	for _, l := range f.IfBody {
		fmt.Fprintf(&buf, "\t%s\n", l)
	}
	for _, e := range f.Elif {
		fmt.Fprintf(&buf, "elif %s =>\n", e.ElifCondition)
		for _, l := range e.ElifBody {
			fmt.Fprintf(&buf, "\t%s\n", l)
		}
	}
	if f.IfCondition != nil {
		fmt.Fprint(&buf, "else =>\n")
		for _, l := range f.ElseBody {
			fmt.Fprintf(&buf, "\t%s\n", l)
		}
	}

	fmt.Fprint(&buf, "}\n")
	return buf.String()
}

// part of ...
func (f *IfAST) locl() {}
