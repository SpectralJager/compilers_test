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
	program := ast.Program{}
	for p.peek(0).Type != tokens.EOF {
		program.Body = append(program.Body, p.parseExpr())
	}
	return &program
}

func (p *Parser) parseExpr() ast.Expr {
	tok := p.peek(0)
	switch tok.Type {
	case tokens.LParen:
		if p.peek(1).Type == tokens.Symbol {
			return p.parseSExpr()
		} else {
			return p.parseSpForm()
		}
	default:
		panic("Cant parse Expression. Should start with '(', got: " + tok.String())
	}
}

func (p *Parser) parseAtom() ast.Atom {
	tok := p.nextToken()
	switch tok.Type {
	case tokens.Int, tokens.Float:
		return &ast.Number{Token: tok}
	case tokens.String:
		return &ast.String{Token: tok}
	case tokens.Bool:
		return &ast.Bool{Token: tok}
	case tokens.Symbol:
		return &ast.Symbol{Token: tok}
	default:
		panic("Unsupported atom type, got: " + tok.Type.String())
	}
}

func (p *Parser) parseSExpr() ast.Expr {
	p.nextToken() // consume '('
	var se ast.SymbolExpr

	se.Symb = p.getSymbol()

	for p.peek(0).Type != tokens.RParen {
		tok := p.peek(0)
		switch tok.Type {
		case tokens.LParen:
			p.nextToken()
			se.Args = append(se.Args, p.parseSExpr())
		case tokens.Int, tokens.Float, tokens.String, tokens.Bool, tokens.Symbol:
			se.Args = append(se.Args, p.parseAtom())
		}
	}
	p.nextToken() // consume ')'
	return &se
}

func (p *Parser) parseSpForm() ast.Expr {
	p.nextToken()
	tok := p.nextToken()
	switch tok.Type {
	case tokens.Def:
		return p.parseDef()
	case tokens.Set:
		return p.parseSet()
	case tokens.Fn:
		return p.parseFn()
	case tokens.Ret:
		return p.parseRet()
	default:
		panic("Unsupported sp-form opperation " + tok.Type.String())
	}
}

func (p *Parser) parseRet() ast.Expr {
	var ret ast.RetSF
	ret.Value = p.parseAtom()
	p.nextToken() // consume ')'
	return &ret
}

func (p *Parser) parseFn() ast.Expr {
	var fn ast.FnSF

	fn.Symb = p.getSymbol()

	if p.peek(0).Type != tokens.LBracket {
		tok := p.nextToken()
		panic("Expected '[' for list of function arguments, got: " + tok.String())
	}
	p.nextToken() // consume '['
	for p.peek(0).Type != tokens.RBracket {
		smb, ok := p.parseAtom().(*ast.Symbol)
		if !ok {
			panic("Argument for function declaration should be symbol, expected symbol, got " + smb.Token.String())
		}
		fn.Args = append(fn.Args, *smb)
	}
	p.nextToken() // consume ']'

	p.nextToken() // consume '('
	for p.peek(0).Type != tokens.RParen {
		fn.Body = append(fn.Body, p.parseExpr())
	}
	p.nextToken() // consume ')'
	p.nextToken() // consume ')'
	return &fn
}

func (p *Parser) parseSet() ast.Expr {
	var set ast.SetSF

	set.Symb = p.getSymbol()

	for p.peek(0).Type != tokens.RParen {
		tok := p.peek(0)
		switch tok.Type {
		case tokens.LParen:
			set.Value = p.parseSExpr()
		case tokens.Int, tokens.Float, tokens.String, tokens.Bool, tokens.Symbol:
			set.Value = p.parseAtom()
		default:
			panic("argument for set should be atom or expression, got " + tok.String())
		}
	}
	p.nextToken() // consume ')'

	return &set
}

func (p *Parser) parseDef() ast.Expr {
	var def ast.DefSF

	def.Symb = p.getSymbol()

	for p.peek(0).Type != tokens.RParen {
		tok := p.peek(0)
		switch tok.Type {
		case tokens.LParen:
			def.Value = p.parseSExpr()
		case tokens.Int, tokens.Float, tokens.String, tokens.Bool, tokens.Symbol:
			def.Value = p.parseAtom()
		default:
			panic("argument for def should be atom or expression, got " + tok.String())
		}
	}
	p.nextToken() //consume ')'
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

func (p *Parser) getSymbol() ast.Symbol {
	smb, ok := p.parseAtom().(*ast.Symbol)
	if !ok {
		panic("Expected Symbol, got: " + smb.String())
	}
	return *smb
}
