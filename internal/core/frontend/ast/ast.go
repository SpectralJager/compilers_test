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
	Value int
}

func (number *Number) atom()

// String atom
type String struct {
	Token tokens.Token
	Value string
}

func (str *String) atom()

// Bool atom
type Bool struct {
	Token tokens.Token
	Value bool
}

func (bl *Bool) atom()

// Nil atom
type Nil struct {
	Token tokens.Token
}

func (nl *Nil) atom()

// Symbol atom
type Symbol struct {
	Token tokens.Token
	Value Node
}

func (symbol *Symbol) atom()

// List atom
type List struct {
	TokL  tokens.Token
	TokR  tokens.Token
	Atoms []Atom
}

func (list *List) atom()

// Vector atom
type Vector struct {
	TokL  tokens.Token
	TokR  tokens.Token
	Atoms []Atom
}

func (vector *Vector) atom()

// List atom
type HashMap struct {
	TokSt  tokens.Token
	TokEn  tokens.Token
	Keys   []Atom // string or number
	Values []Atom
}

func (hashmap *HashMap) atom()

// Prefix-op s-expr
type PrefixOp struct {
	Operator tokens.Token
	Args     []Node
}

func (prefixOp *PrefixOp) sexpr()

// Prefix-op s-expr
type UnaryOp struct {
	Operator tokens.Token
	Arg      Node
}

func (unaryOp *UnaryOp) sexpr()

// Prefix-op s-expr
type BinOp struct {
	Operator tokens.Token
	Arg1     Node
	Arg2     Node
}

func (binOp *BinOp) sexpr()

// SymbolExpr
type SymbolExpr struct {
	Symb Symbol
	Args []Node
}

func (s *SymbolExpr) sexpr()

// Def expr
type DefExpr struct {
	DefToken  tokens.Token
	Smb       Symbol
	BindValue Node
}

func (defexpr *DefExpr) sexpr()

// Fn expr
type FnExpr struct {
	FnTok tokens.Token
	Args  []Atom // vector
	Doc   Atom   // string
	Body  Node
}

func (fnexpr *FnExpr) sexpr()

// Do expr
type DoExpr struct {
	DoTok tokens.Token
	Body  []SExpr
}

func (doexpr *DoExpr) sexpr()
