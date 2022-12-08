package ast

import (
	"encoding/json"
	"grimlang/internal/core/frontend/tokens"
)

// ---------------- SP-Form ------------------------
type DefSF struct {
	Symb  tokens.Token `json:"symb"`
	Value Node         `json:"value"`
}

func (sp *DefSF) spform() {}
func (sp *DefSF) TokenLiteral() string {
	return "def " + sp.Symb.Value
}
func (sp *DefSF) String() string {
	res, err := json.Marshal(sp)
	if err != nil {
		return ""
	}
	return string(res)
}
