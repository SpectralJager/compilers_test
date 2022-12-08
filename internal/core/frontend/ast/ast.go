package ast

import (
	"encoding/json"
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

type SPForm interface {
	Node
	spform()
}

// --------------- Starting point ---------------
// Program node
type Program struct {
	Body []Node `json:"body"`
}

func (program *Program) TokenLiteral() string { return "program" }
func (program *Program) String() string {
	out, err := json.MarshalIndent(program, "", "-")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}
