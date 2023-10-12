package ast

import "fmt"

type SetAST struct {
	Symbol     SymbolAST `parser:"'@set' @@"`
	Expression EXPR      `parser:"'=' @@ ';'"`
}

func (s SetAST) String() string {
	return fmt.Sprintf("@set %s = %s;", &s.Symbol, s.Expression)
}

// part of ...
func (v SetAST) locl() {}
func (v SetAST) ast()  {}
