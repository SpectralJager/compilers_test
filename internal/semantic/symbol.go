package semantic

import (
	"fmt"
	tp "grimlang/internal/type"
	"strings"
)

type SymbolKind string

const (
	Function SymbolKind = "function"
	Local    SymbolKind = "local"
	Global   SymbolKind = "global"
	Type     SymbolKind = "type"
)

type Symbol struct {
	Name  string
	Kind  SymbolKind
	Scope string
	Args  []tp.Type
	Type  *tp.Type
}

func NewTypeSymbol(name string) Symbol {
	return Symbol{
		Name: name,
		Kind: Type,
	}
}

func NewLocalSymbol(name string, tp tp.Type, scope string) Symbol {
	return Symbol{
		Name:  name,
		Type:  &tp,
		Kind:  Local,
		Scope: scope,
	}
}

func NewGlobalSymbol(name string, tp tp.Type) Symbol {
	return Symbol{
		Name: name,
		Type: &tp,
		Kind: Global,
	}
}

func NewFunctionSymbol(name string, args []tp.Type, ret *tp.Type) Symbol {
	return Symbol{
		Name: name,
		Type: ret,
		Args: args,
		Kind: Function,
	}
}

func (s *Symbol) String() string {
	switch s.Kind {
	case Type:
		return fmt.Sprintf("type:%s", s.Name)
	case Local:
		return fmt.Sprintf("local:%s-%s -> %s", s.Name, s.Scope, s.Type)
	case Global:
		return fmt.Sprintf("global:%s -> %s", s.Name, s.Type)
	case Function:
		var buf strings.Builder
		fmt.Fprintf(&buf, "func:%s => ", s.Name)
		if s.Args != nil {
			for _, arg := range s.Args {
				fmt.Fprintf(&buf, " %s", arg)
			}
			fmt.Fprint(&buf, " ->")
			if s.Type != nil {
				fmt.Fprintf(&buf, " %s", s.Type)
			} else {
				fmt.Fprint(&buf, " void")
			}
		} else {
			fmt.Fprint(&buf, " void ->")
			if s.Type != nil {
				fmt.Fprintf(&buf, " %s", s.Type)
			} else {
				fmt.Fprint(&buf, " void")
			}
		}
		return buf.String()
	default:
		return fmt.Sprintf("undefined:%s", s.Name)
	}
}
