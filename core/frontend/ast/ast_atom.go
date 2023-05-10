package ast

import "gl/core/frontend/tokens"

type Atom interface {
	atom()
	GetValue() string
}

// Integer
type Integer struct {
	Token tokens.Token
}

func (a *Integer) atom()     {}
func (a *Integer) ast()      {}
func (ex *Integer) exprArg() {}
func (a *Integer) GetValue() string {
	return a.Token.Value
}

// Float
type Float struct {
	Token tokens.Token
}

func (a *Float) atom()     {}
func (a *Float) ast()      {}
func (ex *Float) exprArg() {}
func (a *Float) GetValue() string {
	return a.Token.Value
}

// Boolean
type Boolean struct {
	Token tokens.Token
}

func (a *Boolean) atom()     {}
func (a *Boolean) ast()      {}
func (ex *Boolean) exprArg() {}
func (a *Boolean) GetValue() string {
	return a.Token.Value
}

// String
type String struct {
	Token tokens.Token
}

func (a *String) atom()     {}
func (a *String) ast()      {}
func (ex *String) exprArg() {}
func (a *String) GetValue() string {
	return a.Token.Value
}

// Symbol
type Symbol struct {
	Token tokens.Token
}

func (a *Symbol) atom()     {}
func (a *Symbol) ast()      {}
func (ex *Symbol) exprArg() {}
func (a *Symbol) GetValue() string {
	return a.Token.Value
}
