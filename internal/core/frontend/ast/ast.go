package ast

import (
	"grimlang/internal/core/frontend/meta"
	"grimlang/internal/core/frontend/tokens"
)

type Node interface {
	NodeType() string
}

// Programm
type Programm struct {
	Package string
	Body    []Node
}

// Atoms -------------------
type Atom interface {
	Node
	Value() string
	AtomType() string
}

// Symbol
type Symbol struct {
	Token tokens.Token
	Meta  meta.Metadata
}

func (atom *Symbol) NodeType() string {
	return "atom"
}
func (atom *Symbol) Value() string {
	return atom.Token.Value
}
func (atom *Symbol) AtomType() string {
	return atom.Token.Type.String()
}

// Number
type Number struct {
	Token tokens.Token
}

func (atom *Number) NodeType() string {
	return "atom"
}
func (atom *Number) Value() string {
	return atom.Token.Value
}
func (atom *Number) AtomType() string {
	return atom.Token.Type.String()
}

// String
type String struct {
	Token tokens.Token
}

func (atom *String) NodeType() string {
	return "atom"
}
func (atom *String) Value() string {
	return atom.Token.Value
}
func (atom *String) AtomType() string {
	return atom.Token.Type.String()
}

// Bool
type Bool struct {
	Token tokens.Token
}

func (atom *Bool) NodeType() string {
	return "atom"
}
func (atom *Bool) Value() string {
	return atom.Token.Value
}
func (atom *Bool) AtomType() string {
	return atom.Token.Type.String()
}

// Bool
type Nil struct {
	Token tokens.Token
}

func (atom *Nil) NodeType() string {
	return "atom"
}
func (atom *Nil) Value() string {
	return ""
}
func (atom *Nil) AtomType() string {
	return atom.Token.Type.String()
}

// Commands ----------------
type Command interface {
	Node
	CommandType() string
}

// Function command
type FnCom struct {
	Name    Symbol
	RetType Symbol
	Params  []Symbol
	Body    []Node
}

func (com *FnCom) NodeType() string {
	return "command"
}
func (com *FnCom) CommandType() string {
	return "fn"
}

// Expression --------------
type Expression interface {
	Node
	Operator() string
	Args() []string
}
