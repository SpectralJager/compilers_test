package ast

import (
	"encoding/json"
	"log"
)

type SExpr interface {
	Expr
	sexpr()
}

type SymbolExpr struct {
	Symb Symbol `json:"symbol"`
	Args []Node `json:"args"`
}

func (sexp *SymbolExpr) expr()        {}
func (sexp *SymbolExpr) sexpr()       {}
func (sexp *SymbolExpr) Type() string { return "s-expr" }
func (sexp *SymbolExpr) String() string {
	out, err := json.MarshalIndent(sexp, "", "-")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}
