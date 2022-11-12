package ast

import (
	"fmt"
	"grimlang/internal/core/frontend/tokens"
)

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
	Operation      tokens.Token
	EndListToken   tokens.Token
	Arguments      []SyntaxNode
}

func (se *SExpression) expressionNode() {}
func (se *SExpression) TokenLiteral() string {
	lit := fmt.Sprintf("(%s", se.Operation.Literal)
	for _, arg := range se.Arguments {
		lit += " "
		lit += arg.TokenLiteral()
	}
	lit += ")"
	return lit
}

// ----------------------------------
type QuoteExpression struct {
	QuoteToken tokens.Token
	Body       SyntaxNode
}

func (qe *QuoteExpression) expressionNode() {}
func (qe *QuoteExpression) TokenLiteral() string {
	return fmt.Sprintf("'%q", qe.Body.TokenLiteral())
}

// ----------------------------------
type ListAtom struct {
	StartListToken tokens.Token
	EndListToken   tokens.Token
	Arguments      []SyntaxAtom
}

func (la *ListAtom) atomNode() {}
func (la *ListAtom) TokenLiteral() string {
	lit := "["
	for _, arg := range la.Arguments {
		lit += " "
		lit += arg.TokenLiteral()
	}
	lit += "]"
	return lit
}

// ----------------------------------
type SymbolAtom struct {
	Symbol tokens.Token
}

func (sa *SymbolAtom) atomNode()            {}
func (sa *SymbolAtom) TokenLiteral() string { return sa.Symbol.Literal }

// ----------------------------------
type NumberAtom struct {
	Number tokens.Token
}

func (na *NumberAtom) atomNode()            {}
func (na *NumberAtom) TokenLiteral() string { return na.Number.Literal }
