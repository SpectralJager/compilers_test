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
	Body []Node `json:"body"`
}

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
type Int struct {
	Token tokens.Token `json:"token"`
}

func (atm *Int) atom()                {}
func (atm *Int) TokenLiteral() string { return atm.Token.Value }
func (atm *Int) String() string       { return atm.Token.Value }
