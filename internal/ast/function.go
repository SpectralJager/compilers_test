package ast

import (
	"fmt"
	"strings"
)

type FunctionAST struct {
	Symbol SymbolAST `parser:"'@fn' @@"`
	Args   []struct {
		Symbol SymbolAST `parser:"@@"`
		Type   TypeAST   `parser:"':'@@"`
	} `parser:"'(' @@* ')'"`
	ReturnTypes []TypeAST `parser:"('<' @@+ '>')?"`
	Body        []LOCAL   `parser:"'{' @@+ '}'"`
}

func (f *FunctionAST) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "@fn %s(", &f.Symbol)
	for _, s := range f.Args {
		fmt.Fprintf(&buf, " %s:%s", &s.Symbol, s.Type)
	}
	fmt.Fprint(&buf, ")")
	if f.ReturnTypes != nil {
		fmt.Fprint(&buf, "<")
		for _, s := range f.ReturnTypes {
			fmt.Fprintf(&buf, " %s", &s)
		}
		fmt.Fprint(&buf, ">")
	}
	fmt.Fprint(&buf, "{\n")
	for _, s := range f.Body {
		fmt.Fprintf(&buf, "\t%s\n", s)
	}
	fmt.Fprint(&buf, "}")
	return buf.String()
}

// part of ...
func (f *FunctionAST) glob() {}
