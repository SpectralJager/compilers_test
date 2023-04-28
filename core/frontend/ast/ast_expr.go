package ast

import "gl/core/frontend/tokens"

// type Expression interface {
// 	expr()
// 	GetSymbol() string
// 	GetArgs() []ExpressionArg
// }

type Expression struct {
	Symbol tokens.Token
	Args   []ExpressionArg
}

func (ea *Expression) exprArg() {}
func (l *Expression) locals()   {}
func (e *Expression) GetSymbol() string {
	return e.Symbol.Value
}
func (e *Expression) GetArgs() []ExpressionArg {
	return e.Args
}
