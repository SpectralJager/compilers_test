package parser

import (
	"gl/core/frontend/ast"
	"gl/core/frontend/tokens"
)

type Parser struct {
	tokens []tokens.Token
	pos    int
	errors []error
}

func NewParser(tokens []tokens.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() *ast.Program {
	var program ast.Program
	return &program
}
