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

func (a *Integer) atom() {}
func (a *Integer) GetValue() string {
	return a.Token.Value
}

// Float
type Float struct {
	Token tokens.Token
}

func (a *Float) atom() {}
func (a *Float) GetValue() string {
	return a.Token.Value
}

// Boolean
type Boolean struct {
	Token tokens.Token
}

func (a *Boolean) atom() {}
func (a *Boolean) GetValue() string {
	return a.Token.Value
}

// String
type String struct {
	Token tokens.Token
}

func (a *String) atom() {}
func (a *String) GetValue() string {
	return a.Token.Value
}
