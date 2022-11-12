package parser

import (
	"fmt"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/tokens"
)

type Parser struct {
	tokenChan chan tokens.Token

	syntaxTree ast.Program
	// currentNode ast.SyntaxNode

	prevToken    tokens.Token
	currentToken tokens.Token
}

func NewParser(tokenChan chan tokens.Token) *Parser {
	parser := Parser{
		tokenChan: tokenChan,
		syntaxTree: ast.Program{
			Nodes: []ast.SyntaxNode{},
		},
	}

	return &parser
}

func (p *Parser) readToken() {
	p.prevToken = p.currentToken
	p.currentToken = <-p.tokenChan
}

func (p *Parser) ParseProgram() ast.Program {
	var node ast.SyntaxNode
	for {
		p.readToken()
		switch p.currentToken.Type {
		case tokens.LeftParen:
			node = p.parseSExpression()
		case tokens.Quote:
			node = p.parseQuoteExpression()
		case tokens.EOF:
			return p.syntaxTree
		default:
			panic(fmt.Sprintf("expected '(', got (%q)", &p.currentToken))
		}
		p.syntaxTree.Nodes = append(p.syntaxTree.Nodes, node)
	}
}

func (p *Parser) parseSExpression() ast.SyntaxNode {
	var SExpr ast.SExpression
	SExpr.StartListToken = p.currentToken
	p.readToken()
	if isOperationOrKeyword(p.currentToken) {
		SExpr.Operation = p.currentToken
	} else {
		panic(fmt.Sprintf("expected symbol, operation or keyword at (_ ...), got (%q)", &p.currentToken))
	}
	for {
		p.readToken()
		switch p.currentToken.Type {
		case tokens.Identifier:
			SExpr.Arguments = append(SExpr.Arguments, &ast.SymbolAtom{Symbol: p.currentToken})
		case tokens.Int:
			SExpr.Arguments = append(SExpr.Arguments, &ast.NumberAtom{Number: p.currentToken})
		case tokens.RightParen:
			SExpr.EndListToken = p.currentToken
			return &SExpr
		case tokens.LeftParen:
			SExpr.Arguments = append(SExpr.Arguments, p.parseSExpression())
		default:
			panic(fmt.Sprintf("expected data or new s-expr, got (%q)", p.currentToken))
		}
	}
}

func (p *Parser) parseQuoteExpression() ast.SyntaxNode {
	var QExpr ast.QuoteExpression
	QExpr.QuoteToken = p.currentToken
	for {
		p.readToken()
		switch p.currentToken.Type {
		case tokens.Identifier:
			QExpr.Body = &ast.SymbolAtom{Symbol: p.currentToken}
			return &QExpr
		default:
			panic(fmt.Sprintf("expected symbol, got (%q)", p.currentToken))
		}
	}
}

func (p *Parser) parseList() ast.SyntaxNode {
	var LAtom ast.ListAtom
	LAtom.StartListToken = p.currentToken
	for {
		p.readToken()
		switch p.currentToken.Type {
		case tokens.Identifier:
			LAtom.Arguments = append(LAtom.Arguments, &ast.SymbolAtom{Symbol: p.currentToken})
		case tokens.Int:
			LAtom.Arguments = append(LAtom.Arguments, &ast.NumberAtom{Number: p.currentToken})
		case tokens.RightParen:
			LAtom.EndListToken = p.currentToken
			return &LAtom
		default:
			panic(fmt.Sprintf("expected data or new s-expr, got (%q)", p.currentToken))
		}
	}
}

// utils
func isOperationOrKeyword(tok tokens.Token) bool {
	if _, ok := tokens.Operations[tok.Literal]; ok {
		return true
	} else if _, ok := tokens.Keywords[tok.Literal]; ok {
		return true
	} else if tok.Type == tokens.Identifier {
		return true
	} else {
		return false
	}
}
