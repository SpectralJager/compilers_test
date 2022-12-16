package ast

import (
	"encoding/json"
)

type SpForm interface {
	Expr
	spform()
}

// ---------------- SP-Form ------------------------
type DefSF struct {
	Symb  Symbol `json:"symb"`
	Value Node   `json:"value"`
}

func (sp *DefSF) expr()   {}
func (sp *DefSF) spform() {}
func (sp *DefSF) Type() string {
	return "def"
}
func (sp *DefSF) String() string {
	res, err := json.Marshal(sp)
	if err != nil {
		return ""
	}
	return string(res)
}

type SetSF struct {
	Symb  Symbol `json:"symb"`
	Value Node   `json:"value"`
}

func (sp *SetSF) expr()   {}
func (sp *SetSF) spform() {}
func (sp *SetSF) Type() string {
	return "set"
}
func (sp *SetSF) String() string {
	res, err := json.Marshal(sp)
	if err != nil {
		return ""
	}
	return string(res)
}

type FnSF struct {
	Symb Symbol   `json:"sumb"`
	Args []Symbol `json:"args"`
	Body []Expr   `json:"body"`
}

func (fn *FnSF) expr()   {}
func (fn *FnSF) spform() {}
func (fn *FnSF) Type() string {
	return "fn"
}
func (fn *FnSF) String() string {
	res, err := json.Marshal(fn)
	if err != nil {
		return ""
	}
	return string(res)
}

type RetSF struct {
	Value Atom `json:"value"`
}

func (ret *RetSF) expr()   {}
func (ret *RetSF) spform() {}
func (ret *RetSF) Type() string {
	return "ret"
}
func (ret *RetSF) String() string {
	res, err := json.Marshal(ret)
	if err != nil {
		return ""
	}
	return string(res)
}
