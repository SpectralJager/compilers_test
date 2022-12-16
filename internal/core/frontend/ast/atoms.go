package ast

import (
	"encoding/json"
	"grimlang/internal/core/frontend/tokens"
)

// ---------------- Atoms ------------------------
type Atom interface {
	Node
	atom()
}

// Number atom (int, float)
type Number struct {
	Token tokens.Token `json:"number"`
}

func (atm *Number) atom()        {}
func (atm *Number) Type() string { return "number" }
func (atm *Number) String() string {
	res, err := json.Marshal(atm)
	if err != nil {
		return ""
	}
	return string(res)
}

// String atom
type String struct {
	Token tokens.Token `json:"string"`
}

func (atm *String) atom()        {}
func (atm *String) Type() string { return "string" }
func (atm *String) String() string {
	res, err := json.Marshal(atm)
	if err != nil {
		return ""
	}
	return string(res)
}

// Bool atom
type Bool struct {
	Token tokens.Token `json:"bool"`
}

func (atm *Bool) atom()        {}
func (atm *Bool) Type() string { return "bool" }
func (atm *Bool) String() string {
	res, err := json.Marshal(atm)
	if err != nil {
		return ""
	}
	return string(res)
}

// Symbol atom
type Symbol struct {
	Token tokens.Token `json:"symbol"`
}

func (atm *Symbol) atom()        {}
func (atm *Symbol) Type() string { return "symbol" }
func (atm *Symbol) String() string {
	res, err := json.Marshal(atm)
	if err != nil {
		return ""
	}
	return string(res)
}
