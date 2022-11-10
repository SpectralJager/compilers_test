package ast

import "grimlang/internal/core/frontend/tokens"

type SyntaxNode interface {
	TokenLiteral() string
}

type SyntaxExpression interface {
	SyntaxNode
	expressionNode()
}

type SyntaxAtom interface {
	SyntaxNode
	atomNode()
}

// ----------------------------------
type Program struct {
	Nodes []SyntaxNode
}

func (p *Program) TokenLiteral() string { return "Program" }

// ----------------------------------
type SExpression struct {
	StartListToken tokens.Token
	EndListToken   tokens.Token
	Arguments      []SyntaxNode
}

func (se *SExpression) expressionNode()
func (se *SExpression) TokenLiteral() string { return "(...)" }

// ----------------------------------
type SymbolAtom struct {
	Symbol tokens.Token
}

func (sa *SymbolAtom) atomNode()
func (sa *SymbolAtom) TokenLiteral() string { return sa.Symbol.Literal }

// ----------------------------------
type NumberAtom struct {
	Number tokens.Token
}

func (na *NumberAtom) atomNode()
func (na *NumberAtom) TokenLiteral() string { return na.Number.Literal }
