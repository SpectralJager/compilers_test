package models

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

type Atom interface {
	Node
	statementNode()
}

// ----------------------------------
// starting node of program
type Program struct {
	Nodes []Node
}

func (p *Program) TokenLiteral() string { return "Program" }

type Operator struct {
	OpToken Token
	Args    []Node
}
