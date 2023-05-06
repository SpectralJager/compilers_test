package ast

import "gl/core/frontend/tokens"

type SpecForm interface {
	spec()
}

type ConstSP struct {
	Symbol tokens.Token
	Type   TypeSymbol
	Value  Atom
}

func (sp *ConstSP) spec()   {}
func (g *ConstSP) globals() {}
func (l *ConstSP) locals()  {}

type VarSP struct {
	Symbol tokens.Token
	Type   TypeSymbol
	Value  ExpressionArg
}

func (sp *VarSP) spec()   {}
func (g *VarSP) globals() {}
func (l *VarSP) locals()  {}

type FnSP struct {
	Symbol tokens.Token
	Type   TypeSymbol
	Args   []FnParams
	Body   []Locals
}

func (sp *FnSP) spec()   {}
func (g *FnSP) globals() {}
