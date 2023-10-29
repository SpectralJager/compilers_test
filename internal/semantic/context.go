package semantic

import (
	"fmt"
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
					tp.NewInt(),
					tp.NewInt(),
					tp.NewVariatic(tp.NewInt()),
				},
				tp.ToPtr(tp.NewInt()),
			),
			NewFunctionSymbol(
				"int/sub",
				[]tp.Type{
					tp.NewInt(),
					tp.NewInt(),
					tp.NewVariatic(tp.NewInt()),
				},
				tp.ToPtr(tp.NewInt()),
			),
			NewFunctionSymbol(
				"int/mul",
				[]tp.Type{
					tp.NewInt(),
					tp.NewInt(),
					tp.NewVariatic(tp.NewInt()),
				},
				tp.ToPtr(tp.NewInt()),
			),
			NewFunctionSymbol(
				"int/div",
				[]tp.Type{
					tp.NewInt(),
					tp.NewInt(),
					tp.NewVariatic(tp.NewInt()),
				},
				tp.ToPtr(tp.NewInt()),
			),
		},
	}
}

func (c *SemanticContext) IsExist(sm Symbol) bool {
	for _, s := range c.SymbolTable {
		switch s.Kind {
		case Function, Type:
			if s.Name == sm.Name {
				return true
			}
		case Global:
			if s.Name == sm.Name {
				return sm.Kind != Local
			}
		}
	}
	return false
}

func (c *SemanticContext) AppendSymbol(sm Symbol) error {
	switch sm.Kind {
	case Function, Global:
		if c.IsExist(sm) {
			return fmt.Errorf("symbol %s already exists", sm.String())
		}
		c.SymbolTable = append(c.SymbolTable, sm)
	default:
		return fmt.Errorf("can't append symbol to table, symbol kind is %s", sm.Kind)
	}
	return nil
}

func (c *SemanticContext) FindFunction(name string) (Symbol, error) {
	for _, sm := range c.SymbolTable {
		for sm.Kind == Function && sm.Name == name {
			return sm, nil
		}
	}
	return Symbol{}, fmt.Errorf("can't find function with name %s", name)
}

func (c *SemanticContext) PushType(t tp.Type) {
	c.Stack = append(c.Stack, t)
}

func (c *SemanticContext) PopType() (tp.Type, error) {
	if len(c.Stack) == 0 {
		return tp.Type{}, fmt.Errorf("stack is empty")
	}
	t := c.Stack[]
}

func (c *SemanticContext) String() string {
	var buf strings.Builder
	fmt.Fprint(&buf, "SymbolTable =>\n")
	for _, sm := range c.SymbolTable {
		fmt.Fprintf(&buf, "\t%s\n", sm.String())
	}
	fmt.Fprint(&buf, "Stack =>\n")
	for _, sm := range c.Stack {
		fmt.Fprintf(&buf, "\t%s\n", sm.String())
	}
	return buf.String()
}
