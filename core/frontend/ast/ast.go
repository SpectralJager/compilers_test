package ast

import "gl/core/frontend/tokens"

type Node interface {
	ast()
}

type Locals interface {
	locals()
}

type Globals interface {
	globals()
}

type ExpressionArg interface {
	exprArg()
}

type Program struct {
	Name string
	Body []Globals
}

func (a *Program) ast() {}

type TypeSymbol interface {
	typeSymb()
}

type SimpleType struct {
	Symbol tokens.Token
}

func (a *SimpleType) ast()       {}
func (ts *SimpleType) typeSymb() {}

type FnParams struct {
	Symbol tokens.Token
	Type   TypeSymbol
}

func (a *FnParams) ast() {}
