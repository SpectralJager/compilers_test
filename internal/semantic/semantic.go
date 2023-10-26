package semantic

import (
	"fmt"
	"grimlang/internal/ir"
	tp "grimlang/internal/type"
	"strings"
)

type SemanticContext struct {
	Stack       []tp.Type
	SymbolTable []Symbol
}

func NewSemanticContext() *SemanticContext {
	return &SemanticContext{
		Stack: make([]tp.Type, 0),
		SymbolTable: []Symbol{
			NewTypeSymbol("int"),
			NewTypeSymbol("bool"),
			NewFunctionSymbol(
				"int/add",
				[]tp.Type{
					tp.NewVariatic(tp.NewInt()),
				},
				tp.ToPtr(tp.NewInt()),
			),
		},
	}
}

func (c *SemanticContext) AppendSymbol(sm Symbol) error {
	switch sm.Kind {
	default:
		return fmt.Errorf("can't append symbol to table, symbol kind is %s", sm.Kind)
	}
}

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
		for _, tp := range s.Args {
			fmt.Fprintf(&buf, "%s ", tp)
		}
		fmt.Fprintf(&buf, "-> %s ", s.Type)
		return buf.String()
	default:
		return fmt.Sprintf("undefined:%s", s.Name)
	}
}

func SemanticModule(ctx *SemanticContext, module ir.Module) error {
	for _, fn := range module.Functions {
		fs := NewFunctionSymbol(fn.Name, fn.Meta.ArgTypes, fn.Meta.Type)
		err := ctx.AppendSymbol(fs)
		if err != nil {
			return err
		}
	}
	return nil
}
