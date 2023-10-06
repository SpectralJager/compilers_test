package ast

import (
	"fmt"
	"strings"
)

type ProgramAST struct {
	Name string
	Path string
	Body []GLOBAL `parser:"@@+"`
}

func (p *ProgramAST) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "// program_name:%s\n", p.Name)
	for _, v := range p.Body {
		fmt.Fprintf(&buf, "%s\n", v)
	}
	return buf.String()
}
