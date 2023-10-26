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

func (c *SemanticContext) String() string {
	var buf strings.Builder
	fmt.Fprint(&buf, "SymbolTable =>\n")
	for _, sm := range c.SymbolTable {
		fmt.Fprintf(&buf, "\t%s\n", sm.String())
	}
	return buf.String()
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

func SemanticModule(ctx *SemanticContext, module ir.Module) error {
	for _, fn := range module.Functions {
		fs := NewFunctionSymbol(fn.Name, fn.Meta.ArgTypes, fn.Meta.Type)
		err := ctx.AppendSymbol(fs)
		if err != nil {
			return err
		}
	}
	for _, instr := range module.Global.Body {
		switch instr.Kind {
		case ir.OP_VAR_NEW:
			vr := NewGlobalSymbol(instr.Identifier, instr.Type)
			err := ctx.AppendSymbol(vr)
			if err != nil {
				return err
			}
		case ir.OP_VAR_SAVE:

		}
	}
	return nil
}
