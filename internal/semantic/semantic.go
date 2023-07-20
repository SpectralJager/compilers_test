package semantic

import "bytes"

type ScopeType string

const (
	global   ScopeType = "global"
	function ScopeType = "function"
	local    ScopeType = "local"
)

type Scope struct {
	Type        ScopeType
	Name        string
	SymbolTable *SymbolTable
	Scopes      []*Scope
}

type SymbolTable struct {
	Definitions []*SymbolDefinitions
}

type SymbolDefinitions interface {
	sd() string
}

type FunctionDefinition struct {
	Ident       string
	ReturnTypes []TypeDefenition
	Args        []TypeDefenition
}

func (s *FunctionDefinition) sd() string {
	var buf bytes.Buffer
	buf.WriteString("fn " + s.Ident + ": ")
	for _, def := range s.ReturnTypes {
		buf.WriteString(def.String() + " ")
	}
	buf.WriteString("->")
	return buf.String()
}

type VaribleDefinition struct {
	Ident string
	Type  TypeDefenition
}

type TypeDefenition struct {
	Name     string
	Subtypes []TypeDefenition
}

func (s *TypeDefenition) String() string {
	var buf bytes.Buffer

	return buf.String()
}
