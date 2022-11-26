package parser

import (
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
			panic("Incorect expression")
		}
	}
	return &program
}

func (p *Parser) parseSExpr() ast.SExpr {
	var sexpr ast.SExpr
	tok := p.nextToken()
	switch tok.Type {
	case tokens.Not, tokens.Neg:
		sexpr = p.parseUnaryExpr()
	case tokens.Add, tokens.Sub, tokens.Mul, tokens.Div, tokens.And, tokens.Or:
		sexpr = p.parsePrefixExpr()
	// case tokens.Lt, tokens.Gt, tokens.Leq, tokens.Geq:
	// sexpr = p.parseBinExpr()
	// case tokens.Symbol:
	// sexpr = p.parseSymbolExpr()
	// case tokens.Def:
	// sexpr = p.parseDefExpr()
	// case tokens.Fn:
	// sexpr = p.parseFnExpr()
	case tokens.Number, tokens.Float, tokens.String:
		sexpr = p.parseAtomExpr()
	case tokens.LParen:
		sexpr = p.parseSExpr()
	default:
		panic("Unixpected token")
	}
	return sexpr
}

func (p *Parser) parseUnaryExpr() ast.SExpr {
	sexpr := ast.UnaryExpr{}
	sexpr.Operator = p.peek(0)
	tok := p.nextToken()
	switch tok.Type {
	case tokens.Number:
		sexpr.Arg = &ast.Number{Token: tok}
	case tokens.True, tokens.False:
		sexpr.Arg = &ast.Bool{Token: tok}
	default:
		panic("Unixpected token")
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
		case tokens.LParen:
			sexpr.Args = append(sexpr.Args, p.parseSExpr())
		default:
			panic("Unixpected token")
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
		panic("Unixpected token")
	}
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
