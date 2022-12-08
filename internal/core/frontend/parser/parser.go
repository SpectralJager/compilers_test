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
		case tokens.Int, tokens.Float:
			program.Body = append(program.Body, &ast.Number{Token: tok})
		case tokens.LParen:
			program.Body = append(program.Body, p.parseLExpr())
		default:
			fmt.Println("Cant parse Statement. Should start with '(' or should be Atom, got: " + tok.String())
		}
	}
	return &program
}

func (p *Parser) parseLExpr() ast.Node {
	tok := p.nextToken()
	switch tok.Type {
	case tokens.Def:
		return p.parseDef()
	default:
		panic("unsupported opperation " + tok.Type.String())
	}
}

func (p *Parser) parseDef() ast.SPForm {
	var def ast.DefSF

	tok := p.nextToken()
	if tok.Type != tokens.Symbol {
		return nil
	}
	def.Symb = tok

	tok = p.nextToken()
	switch tok.Type {
	case tokens.Int, tokens.Float:
		def.Value = &ast.Number{Token: tok}
	case tokens.String:
		def.Value = &ast.String{Token: tok}
	default:
		fmt.Println("Cant parse def value, got: " + tok.String())
	}
	p.nextToken()
	return &def
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
