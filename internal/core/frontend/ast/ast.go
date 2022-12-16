package ast

import (
	"encoding/json"
	"log"
)

type Node interface {
	Type() string
	String() string
}

type Expr interface {
	Node
	expr()
}

// Program node
type Program struct {
	Body []Node `json:"body"`
}

func (program *Program) Type() string { return "program" }
func (program *Program) String() string {
	out, err := json.MarshalIndent(program, "", "-")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}
