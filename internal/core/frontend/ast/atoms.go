package ast

import (
	"encoding/json"
	"grimlang/internal/core/frontend/tokens"
)

// ---------------- Atoms ------------------------
// Number atom (int, float)
type Number struct {
	Token tokens.Token `json:"token"`
}

func (atm *Number) atom()                {}
func (atm *Number) TokenLiteral() string { return atm.Token.Value }
func (atm *Number) String() string {
	res, err := json.Marshal(atm)
	if err != nil {
		return ""
	}
	return string(res)
}

// String atom
type String struct {
	Token tokens.Token `json:"token"`
}

func (atm *String) atom()                {}
func (atm *String) TokenLiteral() string { return atm.Token.Value }
func (atm *String) String() string {
	res, err := json.Marshal(atm)
	if err != nil {
		return ""
	}
	return string(res)
}

// Bool atom
type Bool struct {
	Token tokens.Token `json:"token"`
}

func (atm *Bool) atom()                {}
func (atm *Bool) TokenLiteral() string { return atm.Token.Value }
func (atm *Bool) String() string {
	res, err := json.Marshal(atm)
	if err != nil {
		return ""
	}
	return string(res)
}
