package parser

import (
	"fmt"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/tokens"
)

type Parser struct {
	toks []tokens.Token

	current int
}

func NewParser(toks []tokens.Token) *Parser {
	return &Parser{
		toks: toks,
	}
}

func (p *Parser) Run() *ast.Program {
	program := ast.Program{}
	for p.peek(0).Type != tokens.EOF {
		tok := p.nextToken()
		switch tok.Type {
		case tokens.Int:
			program.Body = append(program.Body, &ast.Int{Token: tok})
		default:
			fmt.Println("Cant parse Statement. Should start with '(' or should be Atom, got: " + tok.String())
		}
	}
	return &program
}

func (p *Parser) nextToken() tokens.Token {
	tok := p.toks[p.current]
	p.current += 1
	return tok
}

func (p *Parser) peek(offset int) tokens.Token {
	if p.current+offset >= len(p.toks) {
		return p.toks[len(p.toks)-1]
	}
	return p.toks[p.current+offset]
}
