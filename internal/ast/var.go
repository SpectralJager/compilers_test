package ast

import (
	"fmt"
)

type VarAST struct {
	Symbol     SymbolAST `parser:"'@var' @@"`
	Type       TypeAST   `parser:"':'@@"`
	Expression EXPR      `parser:"'=' @@ ';'"`
}

func (v VarAST) String() string {
	return fmt.Sprintf("@var %s:%s = %s;", &v.Symbol, &v.Type, v.Expression)
}

// part of ...
func (v VarAST) glob() {}
func (v VarAST) locl() {}
