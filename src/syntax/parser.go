package syntax

import "fmt"

type Parser struct {
	tokens []Token
	errors []error

	pos int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Run() *Programm {
	var programm Programm
	for !p.isEOF() {
		tok := p.readToken()
		switch tok.Type {
		default:
			p.errors = append(p.errors, fmt.Errorf("unknown token: %s", tok))
		}
	}
	return &programm
}

func (p *Parser) readToken() Token {
	tok := p.tokens[p.pos]
	p.pos += 1
	return tok
}

func (p *Parser) isEOF() bool {
	return len(p.tokens) == p.pos
}

func (p *Parser) Errors() []error {
	return p.errors
}
