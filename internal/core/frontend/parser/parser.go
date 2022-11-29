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
	program := ast.Program{
		PkgName:  "main",
		FileName: "main",
	}
	for p.peek(0).Type != tokens.EOF {
		tok := p.nextToken()
		switch tok.Type {
		case tokens.LParen:
			program.Expresions = append(program.Expresions, p.parseSExpr())
		default:
			fmt.Println("Cant parse expresion. Should start with '(', get: " + tok.String())
		}
	}
	return &program
}

func (p *Parser) parseSExpr() ast.SExpr {
	var sexpr ast.SExpr
	tok := p.nextToken()
	switch tok.Type {
	case tokens.Add, tokens.Sub, tokens.Mul, tokens.Div:
		sexpr = p.parsePrefixExpr()
	case tokens.Number, tokens.Float, tokens.String:
		sexpr = p.parseAtomExpr()
	case tokens.Def:
		sexpr = p.parseDefExpr()
	case tokens.Symbol:
		sexpr = p.parseSymbolExpr()
	case tokens.LParen:
		sexpr = p.parseSExpr()
	default:
		fmt.Println("Unixpected token " + tok.String())
	}
	return sexpr
}

func (p *Parser) parseDefExpr() ast.SExpr {
	sexpr := ast.DefExpr{}
	sexpr.Symbol = p.nextToken()
	if sexpr.Symbol.Type != tokens.Symbol {
		fmt.Println("expected symbol, got " + sexpr.Symbol.Type.String())
	}
	tok := p.nextToken()
	switch tok.Type {
	case tokens.Number:
		sexpr.Value = &ast.Number{Token: tok}
	case tokens.Float:
		sexpr.Value = &ast.Float{Token: tok}
	case tokens.String:
		sexpr.Value = &ast.String{Token: tok}
	}
	p.nextToken()
	return &sexpr
}

func (p *Parser) parseSymbolExpr() ast.SExpr {
	sexpr := ast.SymbolExpr{}
	sexpr.Symbol = p.peek(-1)
	for p.peek(0).Type != tokens.RParen {
		tok := p.nextToken()
		switch tok.Type {
		case tokens.Number:
			sexpr.Args = append(sexpr.Args, &ast.Number{Token: tok})
		case tokens.LParen:
			sexpr.Args = append(sexpr.Args, p.parseSExpr())
		default:
			fmt.Println("Unixpected token " + tok.String())
		}
	}
	p.nextToken()
	return &sexpr
}

func (p *Parser) parsePrefixExpr() ast.SExpr {
	sexpr := ast.PrefixExpr{}
	sexpr.Operator = p.peek(-1)
	for p.peek(0).Type != tokens.RParen {
		tok := p.nextToken()
		switch tok.Type {
		case tokens.Number:
			sexpr.Args = append(sexpr.Args, &ast.Number{Token: tok})
		case tokens.Float:
			sexpr.Args = append(sexpr.Args, &ast.Float{Token: tok})
		case tokens.String:
			sexpr.Args = append(sexpr.Args, &ast.String{Token: tok})
		case tokens.Symbol:
			sexpr.Args = append(sexpr.Args, &ast.SymbolExpr{Symbol: tok})
		case tokens.LParen:
			sexpr.Args = append(sexpr.Args, p.parseSExpr())
		default:
			fmt.Println("Unixpected token " + tok.String())
		}
	}
	p.nextToken()
	return &sexpr
}

func (p *Parser) parseAtomExpr() ast.SExpr {
	sexpr := ast.AtomExpr{}
	tok := p.peek(-1)
	switch tok.Type {
	case tokens.Number:
		sexpr.Atm = &ast.Number{Token: tok}
	default:
		fmt.Println("Unixpected token " + tok.String())
	}
	p.nextToken()
	return &sexpr
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
