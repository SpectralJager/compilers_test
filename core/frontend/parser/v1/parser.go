package parser

import (
	"fmt"
	"gl/core/frontend/ast"
	"gl/core/frontend/tokens"
	"strings"
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
	for !p.isEOF() {
		result := p.parseGlobal()
		if result == nil {
			return nil
		}
		program.Body = append(program.Body, result)
	}
	return &program
}

func (p *Parser) parseGlobal() ast.Globals {
	tok := p.peek(0)
	switch tok.Type {
	case tokens.TokenConst:
		return p.parseConst()
	case tokens.TokenVar:
		return p.parseVar()
	case tokens.TokenFn:
		return p.parseFn()
	}
	p.errors = append(p.errors, fmt.Errorf("unexpected token %v", tok.String()))
	return nil
}

func (p *Parser) parseLocal() ast.Locals {
	tok := p.peek(0)
	switch tok.Type {
	case tokens.TokenConst:
		return p.parseConst()
	case tokens.TokenVar:
		return p.parseVar()
	case tokens.TokenSet:
		return p.parseSet()
	case tokens.TokenIf:
		return p.parseIf()
	case tokens.TokenLeftParen:
		return p.parserExpression()
	case tokens.TokenWhile:
		return p.parseWhile()
	case tokens.TokenFor:
		return p.parseFor()
	}
	p.errors = append(p.errors, fmt.Errorf("unexpected token %v", tok.String()))
	return nil
}

func (p *Parser) parseFor() *ast.ForSP {
	forSP := new(ast.ForSP)

	if !p.match(tokens.TokenFor) {
		p.errors = append(p.errors, fmt.Errorf("expected @for, got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenSymbol) {
		p.errors = append(p.errors, fmt.Errorf("expected SYMBOL, got %s", p.peek(0)))
		return nil
	}
	forSP.Iterator = p.next()

	if !p.match(tokens.TokenFrom) {
		p.errors = append(p.errors, fmt.Errorf("expected 'from', got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenNumber) {
		p.errors = append(p.errors, fmt.Errorf("expected NUMBER, got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenTo) {
		p.errors = append(p.errors, fmt.Errorf("expected 'to', got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenNumber) {
		p.errors = append(p.errors, fmt.Errorf("expected NUMBER, got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '{', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for !p.match(tokens.TokenRightBrace) {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		forSP.Body = append(forSP.Body, res)
	}

	if !p.match(tokens.TokenRightBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '}', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return forSP
}

func (p *Parser) parseWhile() *ast.WhileSP {
	whileSP := new(ast.WhileSP)

	if !p.match(tokens.TokenWhile) {
		p.errors = append(p.errors, fmt.Errorf("expected @while, got %s", p.peek(0)))
		return nil
	}
	p.next()

	exprArg := p.parseExprArg()
	if exprArg == nil {
		p.errors = append(p.errors, fmt.Errorf("expected ExpressionArg, got %s", p.peek(0)))
		return nil
	}
	whileSP.Then = exprArg

	if !p.match(tokens.TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '{', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for !p.match(tokens.TokenRightBrace) {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		whileSP.Body = append(whileSP.Body, res)
	}

	if !p.match(tokens.TokenRightBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '}', got %s", p.peek(0)))
		return nil
	}
	p.next()

	if p.match(tokens.TokenElse) {
		res := p.parseElse()
		if res == nil {
			return nil
		}
		whileSP.Else = res
	}

	return whileSP
}

func (p *Parser) parseIf() *ast.IfSP {
	ifSP := new(ast.IfSP)

	if !p.match(tokens.TokenIf) {
		p.errors = append(p.errors, fmt.Errorf("expected @if, got %s", p.peek(0)))
		return nil
	}
	p.next()

	exprArg := p.parseExprArg()
	if exprArg == nil {
		p.errors = append(p.errors, fmt.Errorf("expected ExpressionArg, got %s", p.peek(0)))
		return nil
	}
	ifSP.Then = exprArg

	if !p.match(tokens.TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '{', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for !p.match(tokens.TokenRightBrace) {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		ifSP.Body = append(ifSP.Body, res)
	}

	if !p.match(tokens.TokenRightBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '}', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for p.match(tokens.TokenElif) {
		res := p.parseElIf()
		if res == nil {
			return nil
		}
		ifSP.ElIf = append(ifSP.ElIf, *res)
	}

	if p.match(tokens.TokenElse) {
		res := p.parseElse()
		if res == nil {
			return nil
		}
		ifSP.Else = res
	}

	return ifSP
}

func (p *Parser) parseElIf() *ast.ElIfSP {
	elifSP := new(ast.ElIfSP)

	if !p.match(tokens.TokenElif) {
		p.errors = append(p.errors, fmt.Errorf("expected 'elif', got %s", p.peek(0)))
		return nil
	}
	p.next()

	exprArg := p.parseExprArg()
	if exprArg == nil {
		p.errors = append(p.errors, fmt.Errorf("expected ExpressionArg, got %s", p.peek(0)))
		return nil
	}
	elifSP.Then = exprArg

	if !p.match(tokens.TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '{', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for !p.match(tokens.TokenRightBrace) {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		elifSP.Body = append(elifSP.Body, res)
	}

	if !p.match(tokens.TokenRightBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '}', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return elifSP
}

func (p *Parser) parseElse() *ast.ElseSP {
	elseSP := new(ast.ElseSP)

	if !p.match(tokens.TokenElse) {
		p.errors = append(p.errors, fmt.Errorf("expected 'else', got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '{', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for !p.match(tokens.TokenRightBrace) {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		elseSP.Body = append(elseSP.Body, res)
	}

	if !p.match(tokens.TokenRightBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '}', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return elseSP
}
func (p *Parser) parseSet() *ast.SetSP {
	setSP := new(ast.SetSP)

	if !p.match(tokens.TokenSet) {
		p.errors = append(p.errors, fmt.Errorf("expected @set, got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenSymbol) {
		p.errors = append(p.errors, fmt.Errorf("expected SYMBOL, got %s", p.peek(0)))
		return nil
	}
	setSP.Symbol = p.next()

	if !p.match(tokens.TokenAssign) {
		p.errors = append(p.errors, fmt.Errorf("expected '=', got %s", p.peek(0)))
		return nil
	}
	p.next()

	exprArg := p.parseExprArg()
	if exprArg == nil {
		p.errors = append(p.errors, fmt.Errorf("expected ExpressionArg, got %s", p.peek(0)))
		return nil
	}
	setSP.Value = exprArg

	if !p.match(tokens.TokenSemicolon) {
		p.errors = append(p.errors, fmt.Errorf("expected ';', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return setSP
}

func (p *Parser) parseFn() *ast.FnSP {
	fnSP := new(ast.FnSP)

	if !p.match(tokens.TokenFn) {
		p.errors = append(p.errors, fmt.Errorf("expected @fn, got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenSymbol) {
		p.errors = append(p.errors, fmt.Errorf("expected SYMBOL, got %s", p.peek(0)))
		return nil
	}
	fnSP.Symbol = p.next()

	if !p.match(tokens.TokenColon) {
		p.errors = append(p.errors, fmt.Errorf("expected ':', got %s", p.peek(0)))
		return nil
	}
	p.next()

	ts := p.parseTypeSymbol()
	if ts == nil {
		p.errors = append(p.errors, fmt.Errorf("expected TYPE, got %s", p.peek(0)))
		return nil
	}
	fnSP.Type = ts

	if !p.match(tokens.TokenLeftParen) {
		p.errors = append(p.errors, fmt.Errorf("expected '(', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for !p.match(tokens.TokenRightParen) {
		res := p.parseFnParam()
		if res == nil {
			return nil
		}
		fnSP.Args = append(fnSP.Args, *res)
	}

	if !p.match(tokens.TokenRightParen) {
		p.errors = append(p.errors, fmt.Errorf("expected ')', got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '{', got %s", p.peek(0)))
		return nil
	}
	p.next()

	for !p.match(tokens.TokenRightBrace) {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		fnSP.Body = append(fnSP.Body, res)
	}

	if !p.match(tokens.TokenRightBrace) {
		p.errors = append(p.errors, fmt.Errorf("expected '}', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return fnSP
}

func (p *Parser) parseVar() *ast.VarSP {
	varSP := new(ast.VarSP)

	if !p.match(tokens.TokenVar) {
		p.errors = append(p.errors, fmt.Errorf("expected @var, got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenSymbol) {
		p.errors = append(p.errors, fmt.Errorf("expected SYMBOL, got %s", p.peek(0)))
		return nil
	}
	varSP.Symbol = p.next()

	if !p.match(tokens.TokenColon) {
		p.errors = append(p.errors, fmt.Errorf("expected ':', got %s", p.peek(0)))
		return nil
	}
	p.next()

	ts := p.parseTypeSymbol()
	if ts == nil {
		p.errors = append(p.errors, fmt.Errorf("expected TYPE, got %s", p.peek(0)))
		return nil
	}
	varSP.Type = ts

	if !p.match(tokens.TokenAssign) {
		p.errors = append(p.errors, fmt.Errorf("expected '=', got %s", p.peek(0)))
		return nil
	}
	p.next()

	exprArg := p.parseExprArg()
	if exprArg == nil {
		p.errors = append(p.errors, fmt.Errorf("expected ExpressionArg, got %s", p.peek(0)))
		return nil
	}
	varSP.Value = exprArg

	if !p.match(tokens.TokenSemicolon) {
		p.errors = append(p.errors, fmt.Errorf("expected ';', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return varSP
}

func (p *Parser) parseConst() *ast.ConstSP {
	constSP := new(ast.ConstSP)

	if !p.match(tokens.TokenConst) {
		p.errors = append(p.errors, fmt.Errorf("expected @const, got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenSymbol) {
		p.errors = append(p.errors, fmt.Errorf("expected SYMBOL, got %s", p.peek(0)))
		return nil
	}
	constSP.Symbol = p.next()

	if !p.match(tokens.TokenColon) {
		p.errors = append(p.errors, fmt.Errorf("expected ':', got %s", p.peek(0)))
		return nil
	}
	p.next()

	ts := p.parseTypeSymbol()
	if ts == nil {
		p.errors = append(p.errors, fmt.Errorf("expected TYPE, got %s", p.peek(0)))
		return nil
	}
	constSP.Type = ts

	if !p.match(tokens.TokenAssign) {
		p.errors = append(p.errors, fmt.Errorf("expected '=', got %s", p.peek(0)))
		return nil
	}
	p.next()

	atom := p.parseAtom()
	if atom == nil {
		p.errors = append(p.errors, fmt.Errorf("expected ATOM, got %s", p.peek(0)))
		return nil
	}
	constSP.Value = atom

	if !p.match(tokens.TokenSemicolon) {
		p.errors = append(p.errors, fmt.Errorf("expected ';', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return constSP
}

func (p *Parser) parserExpression() *ast.Expression {
	expr := new(ast.Expression)

	if !p.match(tokens.TokenLeftParen) {
		p.errors = append(p.errors, fmt.Errorf("expected '(', got %s", p.peek(0)))
		return nil
	}
	p.next()

	if !p.match(tokens.TokenSymbol) {
		p.errors = append(p.errors, fmt.Errorf("expected SYMBOL, got %s", p.peek(0)))
		return nil
	}
	expr.Symbol = p.next()

	for !p.match(tokens.TokenRightParen) {
		res := p.parseExprArg()
		if res == nil {
			return nil
		}
		expr.Args = append(expr.Args, res)
	}

	if !p.match(tokens.TokenRightParen) {
		p.errors = append(p.errors, fmt.Errorf("expected ')', got %s", p.peek(0)))
		return nil
	}
	p.next()

	return expr
}

func (p *Parser) parseFnParam() *ast.FnParams {
	fnParam := new(ast.FnParams)
	if !p.match(tokens.TokenSymbol) {
		p.errors = append(p.errors, fmt.Errorf("expected SYMBOL, got %s", p.peek(0)))
		return nil
	}
	fnParam.Symbol = p.next()

	if !p.match(tokens.TokenColon) {
		p.errors = append(p.errors, fmt.Errorf("expected ':', got %s", p.peek(0)))
		return nil
	}
	p.next()

	ts := p.parseTypeSymbol()
	if ts == nil {
		p.errors = append(p.errors, fmt.Errorf("expected TYPE, got %s", p.peek(0)))
		return nil
	}
	fnParam.Type = ts

	return fnParam
}

func (p *Parser) parseExprArg() ast.ExpressionArg {
	tok := p.peek(0)
	switch tok.Type {
	case tokens.TokenLeftParen:
		return p.parserExpression()
	case tokens.TokenString, tokens.TokenTrue, tokens.TokenFalse, tokens.TokenNumber, tokens.TokenSymbol:
		return p.parseAtom().(ast.ExpressionArg)
	}
	p.errors = append(p.errors, fmt.Errorf("unexpected token %v", tok.String()))
	return nil
}

func (p *Parser) parseAtom() ast.Atom {
	tok := p.next()
	switch tok.Type {
	case tokens.TokenNumber:
		if strings.Contains(tok.Value, ".") {
			return &ast.Float{Token: tok}
		}
		return &ast.Integer{Token: tok}
	case tokens.TokenString:
		return &ast.String{Token: tok}
	case tokens.TokenTrue, tokens.TokenFalse:
		return &ast.Boolean{Token: tok}
	case tokens.TokenSymbol:
		return &ast.Symbol{Token: tok}
	}
	p.errors = append(p.errors, fmt.Errorf("unexpected token %v", tok.String()))
	return nil
}

func (p *Parser) parseTypeSymbol() ast.TypeSymbol {
	tok := p.next()
	switch tok.Type {
	case tokens.TokenSymbol:
		return &ast.SimpleType{Symbol: tok}
	}
	p.errors = append(p.errors, fmt.Errorf("unexpected token %v", tok.String()))
	return nil
}

func (p *Parser) match(expected tokens.TokenType) bool {
	return p.peek(0).Type == expected
}

func (p *Parser) next() tokens.Token {
	tok := p.peek(0)
	p.pos += 1
	return tok
}

func (p *Parser) peek(n int) tokens.Token {
	if p.pos+n >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}
	return p.tokens[p.pos+n]
}

func (p *Parser) isEOF() bool {
	return p.match(tokens.TokenEOF)
}

func (p *Parser) Errors() []error {
	return p.errors
}
