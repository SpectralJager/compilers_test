package frontend

import (
	"fmt"
	"strings"
)

// Parser
type Parser struct {
	tokens []Token
	pos    int
	errors []error
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() Node {
	var programNode ProgramNode
	for !p.isEOF() {
		result := p.parseGlobal()
		if result == nil {
			return nil
		}
		programNode.Body = append(programNode.Body, result)
	}
	return &programNode
}

// Parse functions
func (p *Parser) parseGlobal() Node {
	tok := p.peek(0)
	switch tok.Type {
	case TokenConst:
		return p.parseConst()
	case TokenVar:
		return p.parseVar()
	case TokenFn:
		return p.parseFunction()
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want GLOBAL", tok.String()))
		return nil
	}
}

func (p *Parser) parseLocal() Node {
	tok := p.peek(0)
	switch tok.Type {
	case TokenConst:
		return p.parseConst()
	case TokenVar:
		return p.parseVar()
	case TokenSet:
		return p.parseSet()
	case TokenWhile:
		return p.parseWhile()
	case TokenIf:
		return p.parseIf()
	case TokenLeftParen:
		return p.parseExpression()
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want LOCAL", tok.String()))
		return nil
	}
}

func (p *Parser) parseAtom() Node {
	tok := p.next()
	switch tok.Type {
	case TokenNumber:
		if strings.Contains(tok.Value, ".") {
			return &FloatNode{Value: tok, Meta: nil}
		}
		return &IntegerNode{Value: tok, Meta: nil}
	case TokenString:
		return &StringNode{Value: tok, Meta: nil}
	case TokenTrue, TokenFalse:
		return &BooleanNode{Value: tok, Meta: nil}
	case TokenSymbol:
		return &SymbolNode{Value: tok, Meta: nil}
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ATOM", tok.String()))
		return nil
	}
}

func (p *Parser) parseDataType() Node {
	tok := p.peek(0)
	switch tok.Type {
	case TokenSymbol:
		return p.parseAtom()
	case TokenList:
		p.errors = append(p.errors, fmt.Errorf("unimplemented"))
		return nil
	case TokenMap:
		p.errors = append(p.errors, fmt.Errorf("unimplemented"))
		return nil
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want DATA_TYPE", tok.String()))
		return nil
	}
}

func (p *Parser) parseExpressionArgument() Node {
	tok := p.peek(0)
	switch tok.Type {
	case TokenSymbol, TokenNumber, TokenString, TokenTrue, TokenFalse:
		return p.parseAtom()
	case TokenLeftParen:
		return p.parseExpression()
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want EXPR_ARG", tok.String()))
		return nil
	}
}

func (p *Parser) parseIf() Node {
	var ifNode IfNode

	if !p.match(TokenIf) {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'if'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseExpressionArgument()
	if res == nil {
		return nil
	}
	ifNode.ConditionExpressions = res

	if !p.match(TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '{'", p.peek(0).String()))
		return nil
	}
	p.next()

	for p.peek(0).Type != TokenRightBrace {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		ifNode.ThenBody = append(ifNode.ThenBody, res)
	}

	if p.peek(0).Type != TokenRightBrace {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '}'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.match(TokenElse) {
		p.next()

		if !p.match(TokenLeftBrace) {
			p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '{'", p.peek(0).String()))
			return nil
		}
		p.next()

		for p.peek(0).Type != TokenRightBrace {
			res := p.parseLocal()
			if res == nil {
				return nil
			}
			ifNode.ElseBody = append(ifNode.ElseBody, res)
		}

		if p.peek(0).Type != TokenRightBrace {
			p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '}'", p.peek(0).String()))
			return nil
		}
		p.next()
	}

	return &ifNode
}

func (p *Parser) parseWhile() Node {
	var whileNode WhileNode

	if !p.match(TokenWhile) {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'while'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseExpressionArgument()
	if res == nil {
		return nil
	}
	whileNode.ConditionExpressions = res

	if !p.match(TokenLeftBrace) {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '{'", p.peek(0).String()))
		return nil
	}
	p.next()

	for p.peek(0).Type != TokenRightBrace {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		whileNode.ThenBody = append(whileNode.ThenBody, res)
	}

	if p.peek(0).Type != TokenRightBrace {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '}'", p.peek(0).String()))
		return nil
	}
	p.next()

	return &whileNode
}

func (p *Parser) parseExpression() Node {
	var expressionNode ExpressionNode

	if p.peek(0).Type != TokenLeftParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '('", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseAtom()
	if res == nil {
		return nil
	}
	expressionNode.Function = res

	for p.peek(0).Type != TokenRightParen {
		res := p.parseExpressionArgument()
		if res == nil {
			return nil
		}
		expressionNode.Args = append(expressionNode.Args, res)
	}

	if p.peek(0).Type != TokenRightParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ')'", p.peek(0).String()))
		return nil
	}
	p.next()

	return &expressionNode
}
func (p *Parser) parseFunction() Node {
	var functionNode FunctionNode

	if p.peek(0).Type != TokenFn {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'fn'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	functionNode.Name = p.parseAtom()

	if p.peek(0).Type != TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	functionNode.ReturnType = res

	if p.peek(0).Type != TokenLeftParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '('", p.peek(0).String()))
		return nil
	}
	p.next()

	for p.peek(0).Type != TokenRightParen {
		res := p.parseParameterNode()
		if res == nil {
			return nil
		}
		functionNode.Params = append(functionNode.Params, res)
	}

	if p.peek(0).Type != TokenRightParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ')'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != TokenLeftBrace {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '{'", p.peek(0).String()))
		return nil
	}
	p.next()

	for p.peek(0).Type != TokenRightBrace {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		functionNode.Body = append(functionNode.Body, res)
	}

	if p.peek(0).Type != TokenRightBrace {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '}'", p.peek(0).String()))
		return nil
	}
	p.next()

	return &functionNode
}

func (p *Parser) parseParameterNode() Node {
	var paramNode ParamNode

	if p.peek(0).Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	paramNode.Name = p.parseAtom()

	if p.peek(0).Type != TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	paramNode.DataType = res

	return &paramNode
}

func (p *Parser) parseVar() Node {
	var varNode VarNode

	if p.peek(0).Type != TokenVar {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'var'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	varNode.Name = p.parseAtom()

	if p.peek(0).Type != TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	varNode.DataType = res

	if p.peek(0).Type != TokenAssign {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '='", p.peek(0).String()))
		return nil
	}
	p.next()

	res = p.parseExpressionArgument()
	if res == nil {
		return nil
	}
	varNode.Value = res

	if p.peek(0).Type != TokenSemicolon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ';'", p.peek(0).String()))
		return nil
	}
	p.next()

	return &varNode
}

func (p *Parser) parseSet() Node {
	var setNode SetNode

	if p.peek(0).Type != TokenSet {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'var'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	setNode.Name = p.parseAtom()

	if p.peek(0).Type != TokenAssign {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '='", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseExpressionArgument()
	if res == nil {
		return nil
	}
	setNode.Value = res

	if p.peek(0).Type != TokenSemicolon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ';'", p.peek(0).String()))
		return nil
	}
	p.next()

	return &setNode
}

func (p *Parser) parseConst() Node {
	var constNode ConstNode

	if p.peek(0).Type != TokenConst {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'const'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	constNode.Name = p.parseAtom()

	if p.peek(0).Type != TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	constNode.DataType = res

	if p.peek(0).Type != TokenAssign {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '='", p.peek(0).String()))
		return nil
	}
	p.next()

	res = p.parseAtom()
	if res == nil {
		return nil
	}
	constNode.Value = res

	if p.peek(0).Type != TokenSemicolon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ';'", p.peek(0).String()))
		return nil
	}
	p.next()

	return &constNode
}

// parser utilities
func (p *Parser) match(expected TokenType) bool {
	return p.peek(0).Type == expected
}

func (p *Parser) next() Token {
	tok := p.peek(0)
	p.pos += 1
	return tok
}

func (p *Parser) peek(n int) Token {
	if p.pos+n >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}
	return p.tokens[p.pos+n]
}

func (p *Parser) isEOF() bool {
	return p.match(TokenEOF)
}

func (p *Parser) Errors() []error {
	return p.errors
}
