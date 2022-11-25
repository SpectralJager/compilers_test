package parser

import (
	"fmt"
)

type Node interface {
	Literal() string
}

type Atom interface {
	Node
	atomStmt()
}

type SExpression interface {
	Node
	sExpression()
}

// ----
type Programm struct {
	Name  string
	Nodes []Node
}

func (p *Programm) Literal() string {
	temp := fmt.Sprintf("{\n\tname: %v\n\tnodes: ", p.Name)
	for _, n := range p.Nodes {
		temp += n.Literal()
	}
	temp += "}\n"
	return temp
}

// ----
type Number struct {
	Value int
}

func (n *Number) atomStmt() {}
func (n *Number) Literal() string {
	return fmt.Sprintf("%d", n.Value)
}

// ----
type Float struct {
	Value float64
}

func (f *Float) atomStmt() {}
func (f *Float) Literal() string {
	return fmt.Sprintf("%f", f.Value)
}

// ----
type String struct {
	Value string
}

func (s *String) atomStmt() {}
func (s *String) Literal() string {
	return s.Value
}

// ----
type Symbol struct {
	Lit string
}
