package parser

import (
	"fmt"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/lexer"
	"grimlang/internal/core/frontend/tokens"
)

type Parser struct {
	lex       lexer.Lexer
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
		case tokens.Plus:
			SExpr.Arguments = append(SExpr.Arguments, &ast.SymbolAtom{Symbol: p.currentToken})
		default:
			panic(fmt.Sprintf("expected symbol, number or new s-expr, got (%q)", &p.currentToken))
		}
	}
}
