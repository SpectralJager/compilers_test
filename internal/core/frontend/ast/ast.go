package ast

import (
	"fmt"
	"grimlang/internal/core/frontend/tokens"
)

type Node interface {
	TokenLiteral() string
}

type Atom interface {
	Node
	atom()
}

type SExpr interface {
	Node
	sexpr()
}

// Program node
type Program struct {
	PkgName    string
	FileName   string
	Expresions []SExpr
}

func (program *Program) TokenLiteral() string {
	return fmt.Sprintf("%s %s, %d", program.PkgName, program.FileName, len(program.Expresions))
}

// Number atom
type Number struct {
	Token tokens.Token
}

func (number *Number) TokenLiteral() string {
	return number.Token.String()
}
func (number *Number) atom() {}

// String atom
type String struct {
	Token tokens.Token
}

func (str *String) TokenLiteral() string {
	return str.Token.String()
}
func (str *String) atom() {}

// Bool atom
type Bool struct {
	Token tokens.Token
}

func (bl *Bool) TokenLiteral() string {
	return bl.Token.String()
}
func (bl *Bool) atom() {}

// Nil atom
type Nil struct {
	Token tokens.Token
}

func (nl *Nil) TokenLiteral() string {
	return nl.Token.String()
}
func (nl *Nil) atom() {}

// Symbol atom
type Symbol struct {
	Token tokens.Token
	Value Node
}

func (symbol *Symbol) TokenLiteral() string {
	return fmt.Sprintf("%s %s", symbol.Token.String(), symbol.Value.TokenLiteral())
}
func (symbol *Symbol) atom() {}

// List atom
type List struct {
	Atoms []Atom
}

func (list *List) TokenLiteral() string {
	return fmt.Sprintf("'(%s ... %s)", list.Atoms[0], list.Atoms[len(list.Atoms)-1])
}
func (list *List) atom() {}

// Vector atom
type Vector struct {
	Atoms []Atom
}

func (vector *Vector) TokenLiteral() string {
	return fmt.Sprintf("[%v ... %v]", vector.Atoms[0], vector.Atoms[len(vector.Atoms)-1])
}
func (vector *Vector) atom() {}

// List atom
type HashMap struct {
	Keys   []Atom // string or number
	Values []Atom
}

func (hashmap *HashMap) TokenLiteral() string {
	return fmt.Sprintf("{%v %v ...}", hashmap.Keys[0], hashmap.Values[0])
}
func (hashmap *HashMap) atom() {}

// Prefix-op s-expr
type PrefixExpr struct {
	Operator tokens.Token
	Args     []Node
}

func (prefixExpr *PrefixExpr) TokenLiteral() string {
	return "(" + prefixExpr.Operator.Value + ")"
}
func (prefixExpr *PrefixExpr) sexpr() {}

// Unary-op s-expr
type UnaryExpr struct {
	Operator tokens.Token
	Arg      Node
}

func (unaryExpr *UnaryExpr) TokenLiteral() string {
	return "(" + unaryExpr.Operator.String() + ")"
}
func (unaryExpr *UnaryExpr) sexpr() {}

// Bin-op s-expr
type BinExpr struct {
	Operator tokens.Token
	Arg1     Node
	Arg2     Node
}

func (binExpr *BinExpr) TokenLiteral() string {
	return "(" + binExpr.Operator.String() + ")"
}
func (binExpr *BinExpr) sexpr() {}

// SymbolExpr
type SymbolExpr struct {
	Symb Symbol
	Args []Node
}

func (s *SymbolExpr) TokenLiteral() string {
	return "(" + s.Symb.Token.String() + ")"
}
func (s *SymbolExpr) sexpr() {}

// Def expr
type DefExpr struct {
	DefToken  tokens.Token
	Smb       Symbol
	BindValue Node
}

func (defexpr *DefExpr) TokenLiteral() string {
	return "(" + defexpr.DefToken.String() + ")"
}
func (defexpr *DefExpr) sexpr() {}

// Fn expr
type FnExpr struct {
	FnTok tokens.Token
	Args  []Atom // vector
	Doc   Atom   // string
	Body  Node
}

func (fnexpr *FnExpr) TokenLiteral() string {
	return "(" + fnexpr.FnTok.String() + ")"
}
func (fnexpr *FnExpr) sexpr() {}

// Do expr
type DoExpr struct {
	DoTok tokens.Token
	Body  []SExpr
}

func (doexpr *DoExpr) TokenLiteral() string {
	return "(" + doexpr.DoTok.String() + ")"
}
func (doexpr *DoExpr) sexpr() {}

// Atom expr
type AtomExpr struct {
	Atm Atom
}

func (atom *AtomExpr) TokenLiteral() string {
	return "(" + atom.Atm.TokenLiteral() + ")"
}
func (atom *AtomExpr) sexpr() {}
