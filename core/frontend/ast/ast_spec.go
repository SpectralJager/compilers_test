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

type SetSP struct {
	Symbol tokens.Token
	Value  ExpressionArg
}

func (sp *SetSP) spec()  {}
func (l *SetSP) locals() {}

type IfSP struct {
	Then ExpressionArg
	Body []Locals
	ElIf []ElIfSP
	Else *ElseSP
}

func (sp *IfSP) spec()  {}
func (l *IfSP) locals() {}

type ElIfSP struct {
	Then ExpressionArg
	Body []Locals
}

func (sp *ElIfSP) spec()  {}
func (l *ElIfSP) locals() {}

type ElseSP struct {
	Body []Locals
}

func (sp *ElseSP) spec()  {}
func (l *ElseSP) locals() {}

type WhileSP struct {
	Then ExpressionArg
	Body []Locals
	Else *ElseSP
}

func (sp *WhileSP) spec()  {}
func (l *WhileSP) locals() {}

type ForSP struct {
	Iterator tokens.Token
	From     tokens.Token
	To       tokens.Token
	Body     []Locals
}

func (sp *ForSP) spec()  {}
func (l *ForSP) locals() {}
