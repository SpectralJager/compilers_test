package ast

import (
	"encoding/json"
	"grimlang/internal/core/frontend/tokens"
	"log"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Atom interface {
	Node
	atom()
}

type SExpr interface {
	Node
	sexpr()
}

// --------------- Starting point ---------------
// Program node
type Program struct {
	PkgName    string  `json:"package"`
	FileName   string  `json:"file"`
	Expresions []SExpr `json:"expressions"`
}

func (program *Program) TokenLiteral() string { return program.PkgName + " " + program.FileName }
func (program *Program) String() string {
	out, err := json.MarshalIndent(program, "", "-")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}

// ---------------- Atoms ------------------------
// Number atom
type Number struct {
	Token tokens.Token `json:"token"`
}

func (number *Number) atom()                {}
func (number *Number) TokenLiteral() string { return number.Token.Value }
func (number *Number) String() string       { return number.Token.String() }

// ---------------- S-Expressions ----------------
// Prefix-op s-expr
type PrefixExpr struct {
	Operator tokens.Token `json:"operator"`
	Args     []Node       `json:"args"`
}

func (prefixExpr *PrefixExpr) sexpr()               {}
func (prefixExpr *PrefixExpr) TokenLiteral() string { return prefixExpr.Operator.Value }
func (prefixExpr *PrefixExpr) String() string       { return prefixExpr.Operator.String() }

// Atom expr
type AtomExpr struct {
	Atm Atom `json:"atom"`
}

func (atom *AtomExpr) sexpr()               {}
func (atom *AtomExpr) TokenLiteral() string { return atom.Atm.TokenLiteral() }
func (atom *AtomExpr) String() string       { return atom.Atm.String() }
