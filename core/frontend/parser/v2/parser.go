package parser

import (
	"encoding/json"
	"fmt"
	"gl/core/frontend/tokens"
	"strings"
)

// Nodes
type Node interface {
	json.Marshaler
	node()
}

func (n ProgramNode) node()  {}
func (n ConstNode) node()    {}
func (n VarNode) node()      {}
func (n SetNode) node()      {}
func (n FunctionNode) node() {}

func (n ExpressionNode) node() {}
func (n ParamNode) node()      {}

func (n IntegerNode) node() {}
func (n FloatNode) node()   {}
func (n StringNode) node()  {}
func (n BooleanNode) node() {}
func (n SymbolNode) node()  {}

type ProgramNode struct {
	Package string            `json:"package"`
	Body    []Node            `json:"body"`
	Meta    map[string]string `json:"meta"`
}

func (progNode ProgramNode) MarshalJSON() ([]byte, error) {
	type tmp ProgramNode
	return toJson(tmp(progNode), "programNode")
}

type ConstNode struct {
	Name     Node              `json:"name"`
	DataType Node              `json:"dataType"`
	Value    Node              `json:"value"`
	Meta     map[string]string `json:"meta"`
}

func (constNode ConstNode) MarshalJSON() ([]byte, error) {
	type tmp ConstNode
	return toJson(tmp(constNode), "constNode")
}

type VarNode struct {
	Name     Node              `json:"name"`
	DataType Node              `json:"dataType"`
	Value    Node              `json:"value"`
	Meta     map[string]string `json:"meta"`
}

func (varNode VarNode) MarshalJSON() ([]byte, error) {
	type tmp VarNode
	return toJson(tmp(varNode), "varNode")
}

type SetNode struct {
	Name  Node              `json:"name"`
	Value Node              `json:"value"`
	Meta  map[string]string `json:"meta"`
}

func (setNode SetNode) MarshalJSON() ([]byte, error) {
	type tmp SetNode
	return toJson(tmp(setNode), "setNode")
}

type FunctionNode struct {
	Name       Node   `json:"name"`
	ReturnType Node   `json:"returnType"`
	Params     []Node `json:"params"`
	Body       []Node `json:"body"`
}

func (funcNode FunctionNode) MarshalJSON() ([]byte, error) {
	type tmp FunctionNode
	return toJson(tmp(funcNode), "functionNode")
}

type ParamNode struct {
	Name     Node              `json:"name"`
	DataType Node              `json:"dataType"`
	Meta     map[string]string `json:"meta"`
}

func (paramNode ParamNode) MarshalJSON() ([]byte, error) {
	type tmp ParamNode
	return toJson(tmp(paramNode), "paramNode")
}

type ExpressionNode struct {
	Function Node              `json:"function"`
	Args     []Node            `json:"args"`
	Meta     map[string]string `json:"meta"`
}

func (exprNode ExpressionNode) MarshalJSON() ([]byte, error) {
	type tmp ExpressionNode
	return toJson(tmp(exprNode), "expressionNode")
}

type IntegerNode struct {
	Value tokens.Token `json:"value"`
}

func (intNode IntegerNode) MarshalJSON() ([]byte, error) {
	type tmp IntegerNode
	return toJson(tmp(intNode), "integerNode")
}

type FloatNode struct {
	Value tokens.Token `json:"value"`
}

func (floatNode FloatNode) MarshalJSON() ([]byte, error) {
	type tmp FloatNode
	return toJson(tmp(floatNode), "floatNode")
}

type StringNode struct {
	Value tokens.Token `json:"value"`
}

func (stringNode StringNode) MarshalJSON() ([]byte, error) {
	type tmp StringNode
	return toJson(tmp(stringNode), "stringNode")
}

type BooleanNode struct {
	Value tokens.Token `json:"value"`
}

func (booleanNode BooleanNode) MarshalJSON() ([]byte, error) {
	type tmp BooleanNode
	return toJson(tmp(booleanNode), "booleanNode")
}

type SymbolNode struct {
	Value tokens.Token `json:"value"`
}

func (symbolNode SymbolNode) MarshalJSON() ([]byte, error) {
	type tmp SymbolNode
	return toJson(tmp(symbolNode), "symbolNode")
}

// Helpers

func toJson(v any, tp string) ([]byte, error) {
	res, _ := json.Marshal(v)
	var temp map[string]interface{}
	json.Unmarshal(res, &temp)
	temp["type"] = tp
	return json.Marshal(temp)
}

// Parser
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

func (p *Parser) Parse() Node {
	var programNode ProgramNode
	for !p.isEOF() {
		result := p.parseGlobal()
		if result == nil {
			return nil
		}
		programNode.Body = append(programNode.Body, result)
	}
	return programNode
}

// Parse functions
func (p *Parser) parseGlobal() Node {
	tok := p.peek(0)
	switch tok.Type {
	case tokens.TokenConst:
		return p.parseConst()
	case tokens.TokenVar:
		return p.parseVar()
	case tokens.TokenFn:
		return p.parseFunction()
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want GLOBAL", tok.String()))
		return nil
	}
}

func (p *Parser) parseLocal() Node {
	tok := p.peek(0)
	switch tok.Type {
	case tokens.TokenConst:
		return p.parseConst()
	case tokens.TokenVar:
		return p.parseVar()
	case tokens.TokenSet:
		return p.parseSet()
	case tokens.TokenLeftParen:
		return p.parseExpression()
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want LOCAL", tok.String()))
		return nil
	}
}

func (p *Parser) parseAtom() Node {
	tok := p.next()
	switch tok.Type {
	case tokens.TokenNumber:
		if strings.Contains(tok.Value, ".") {
			return FloatNode{tok}
		}
		return IntegerNode{tok}
	case tokens.TokenString:
		return StringNode{tok}
	case tokens.TokenTrue, tokens.TokenFalse:
		return BooleanNode{tok}
	case tokens.TokenSymbol:
		return SymbolNode{tok}
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ATOM", tok.String()))
		return nil
	}
}

func (p *Parser) parseDataType() Node {
	tok := p.peek(0)
	switch tok.Type {
	case tokens.TokenSymbol:
		return p.parseAtom()
	case tokens.TokenList:
		p.errors = append(p.errors, fmt.Errorf("unimplemented"))
		return nil
	case tokens.TokenMap:
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
	case tokens.TokenSymbol, tokens.TokenNumber, tokens.TokenString, tokens.TokenTrue, tokens.TokenFalse:
		return p.parseAtom()
	case tokens.TokenLeftParen:
		return p.parseExpression()
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want EXPR_ARG", tok.String()))
		return nil
	}
}

func (p *Parser) parseExpression() Node {
	var expressionNode ExpressionNode

	if p.peek(0).Type != tokens.TokenLeftParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '('", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseAtom()
	if res == nil {
		return nil
	}
	expressionNode.Function = res

	for p.peek(0).Type != tokens.TokenRightParen {
		res := p.parseExpressionArgument()
		if res == nil {
			return nil
		}
		expressionNode.Args = append(expressionNode.Args, res)
	}

	if p.peek(0).Type != tokens.TokenRightParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ')'", p.peek(0).String()))
		return nil
	}
	p.next()

	return expressionNode
}
func (p *Parser) parseFunction() Node {
	var functionNode FunctionNode

	if p.peek(0).Type != tokens.TokenFn {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'fn'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != tokens.TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	functionNode.Name = p.parseAtom()

	if p.peek(0).Type != tokens.TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	functionNode.ReturnType = res

	if p.peek(0).Type != tokens.TokenLeftParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '('", p.peek(0).String()))
		return nil
	}
	p.next()

	for p.peek(0).Type != tokens.TokenRightParen {
		res := p.parseParameterNode()
		if res == nil {
			return nil
		}
		functionNode.Params = append(functionNode.Params, res)
	}

	if p.peek(0).Type != tokens.TokenRightParen {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ')'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != tokens.TokenLeftBrace {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '{'", p.peek(0).String()))
		return nil
	}
	p.next()

	for p.peek(0).Type != tokens.TokenRightBrace {
		res := p.parseLocal()
		if res == nil {
			return nil
		}
		functionNode.Body = append(functionNode.Params, res)
	}

	if p.peek(0).Type != tokens.TokenRightBrace {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '}'", p.peek(0).String()))
		return nil
	}
	p.next()

	return functionNode
}

func (p *Parser) parseParameterNode() Node {
	var paramNode ParamNode

	if p.peek(0).Type != tokens.TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	paramNode.Name = p.parseAtom()

	if p.peek(0).Type != tokens.TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	paramNode.DataType = res

	return paramNode
}

func (p *Parser) parseVar() Node {
	var varNode VarNode

	if p.peek(0).Type != tokens.TokenVar {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'var'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != tokens.TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	varNode.Name = p.parseAtom()

	if p.peek(0).Type != tokens.TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	varNode.DataType = res

	if p.peek(0).Type != tokens.TokenAssign {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '='", p.peek(0).String()))
		return nil
	}
	p.next()

	res = p.parseExpressionArgument()
	if res == nil {
		return nil
	}
	varNode.Value = res

	if p.peek(0).Type != tokens.TokenSemicolon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ';'", p.peek(0).String()))
		return nil
	}
	p.next()

	return varNode
}

func (p *Parser) parseSet() Node {
	var setNode SetNode

	if p.peek(0).Type != tokens.TokenSet {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'var'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != tokens.TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	setNode.Name = p.parseAtom()

	if p.peek(0).Type != tokens.TokenAssign {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '='", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseExpressionArgument()
	if res == nil {
		return nil
	}
	setNode.Value = res

	if p.peek(0).Type != tokens.TokenSemicolon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ';'", p.peek(0).String()))
		return nil
	}
	p.next()

	return setNode
}

func (p *Parser) parseConst() Node {
	var constNode ConstNode

	if p.peek(0).Type != tokens.TokenConst {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want 'const'", p.peek(0).String()))
		return nil
	}
	p.next()

	if p.peek(0).Type != tokens.TokenSymbol {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want SYMBOL", p.peek(0).String()))
		return nil
	}
	constNode.Name = p.parseAtom()

	if p.peek(0).Type != tokens.TokenColon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ':'", p.peek(0).String()))
		return nil
	}
	p.next()

	res := p.parseDataType()
	if res == nil {
		return nil
	}
	constNode.DataType = res

	if p.peek(0).Type != tokens.TokenAssign {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want '='", p.peek(0).String()))
		return nil
	}
	p.next()

	res = p.parseAtom()
	if res == nil {
		return nil
	}
	constNode.Value = res

	if p.peek(0).Type != tokens.TokenSemicolon {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %s, want ';'", p.peek(0).String()))
		return nil
	}
	p.next()

	return constNode
}

// parser utilities
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
