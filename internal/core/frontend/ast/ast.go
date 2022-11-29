package ast

import (
	"encoding/json"
	"fmt"
	"grimlang/internal/core/frontend/tokens"
	"log"
	"strconv"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Atom interface {
	Node
	atom()
	Value() (interface{}, error)
}

type SExpr interface {
	Node
	sexpr()
}

// --------------- Starting point ---------------
// Program node
type Program struct {
	PkgName    string  `json:"package"`
	FileName   string  `json:"file"`
	Expresions []SExpr `json:"expressions"`
}

func (program *Program) TokenLiteral() string { return program.PkgName + " " + program.FileName }
func (program *Program) String() string {
	out, err := json.MarshalIndent(program, "", "-")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}

// ---------------- Atoms ------------------------
// Number atom
type Number struct {
	Token tokens.Token `json:"token"`
}

func (number *Number) atom() {}
func (number *Number) Value() (interface{}, error) {
	switch number.Token.Type {
	case tokens.Number:
		return strconv.Atoi(number.Token.Value)
	default:
		return nil, fmt.Errorf("incorect number type, got '%s'", number.Token.Type.String())
	}
}
func (number *Number) TokenLiteral() string { return number.Token.Value }
func (number *Number) String() string       { return number.Token.Value }

// Float atom
type Float struct {
	Token tokens.Token `json:"token"`
}

func (float *Float) atom() {}
func (float *Float) Value() (interface{}, error) {
	switch float.Token.Type {
	case tokens.Float:
		return strconv.ParseFloat(float.Token.Value, 64)
	default:
		return nil, fmt.Errorf("incorect float type, got '%s'", float.Token.Type.String())
	}
}
func (float *Float) TokenLiteral() string { return float.Token.Value }
func (float *Float) String() string       { return float.Token.Value }

// String atom
type String struct {
	Token tokens.Token `json:"token"`
}

func (str *String) atom() {}
func (str *String) Value() (interface{}, error) {
	return str.Token.Value, nil
}
func (str *String) TokenLiteral() string { return str.Token.Value }
func (str *String) String() string       { return str.Token.Value }

// Symbol atom

// ---------------- S-Expressions ----------------
// Prefix-op s-expr
type PrefixExpr struct {
	Operator tokens.Token `json:"operator"`
	Args     []Node       `json:"args"`
}

func (prefixExpr *PrefixExpr) sexpr()               {}
func (prefixExpr *PrefixExpr) TokenLiteral() string { return prefixExpr.Operator.Value }
func (prefixExpr *PrefixExpr) String() string {
	out, err := json.Marshal(prefixExpr)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}

// Atom expr
type AtomExpr struct {
	Atm Atom `json:"atom"`
}

func (atom *AtomExpr) sexpr()               {}
func (atom *AtomExpr) TokenLiteral() string { return atom.Atm.TokenLiteral() }
func (atom *AtomExpr) String() string       { return atom.Atm.String() }

// Def expr
type DefExpr struct {
	Symbol tokens.Token `json:"symbol"`
	Value  Node         `json:"value"`
}

func (def *DefExpr) sexpr()               {}
func (def *DefExpr) TokenLiteral() string { return "def" }
func (def *DefExpr) String() string {
	out, err := json.Marshal(def)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}

// Symbol expr
type SymbolExpr struct {
	Symbol tokens.Token `json:"symbol"`
	Args   []Node       `json:"args"`
}

func (se *SymbolExpr) sexpr()               {}
func (se *SymbolExpr) TokenLiteral() string { return se.Symbol.Value }
func (se *SymbolExpr) String() string {
	out, err := json.Marshal(se)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(out)
}
