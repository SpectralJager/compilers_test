package symbol

import (
	"fmt"
	"grimlang/builtin"
	"grimlang/dtype"
	"grimlang/object"
)

type Symbol interface {
	Kind() SymbolKind
	Name() string
}

type SymbolKind uint

const (
	Constant SymbolKind = 7 << iota
	Variable
	Function
	Module
	BuiltinFunction
	BuiltinType
)

// =======================================================

type ConstantSymbol struct {
	Identifier string
	Value      object.Object
}

func (*ConstantSymbol) Kind() SymbolKind { return Constant }
func (sm *ConstantSymbol) Name() string {
	return sm.Identifier
}

// -------------------------------------------------------
type VariableSymbol struct {
	Identifier string
	Type       dtype.Type
	Value      object.Object
}

func (*VariableSymbol) Kind() SymbolKind { return Variable }
func (sm *VariableSymbol) Name() string {
	return sm.Identifier
}

// -------------------------------------------------------
type FunctionSymbol struct {
	Identifier string
	Type       dtype.Type
	Fn         any
}

func (*FunctionSymbol) Kind() SymbolKind { return Function }
func (sm *FunctionSymbol) Name() string {
	return sm.Identifier
}

// -------------------------------------------------------
type BuiltinFunctionSymbol struct {
	Identifier string
	Type       dtype.Type
	Callee     builtin.BuiltinFunction
}

func (*BuiltinFunctionSymbol) Kind() SymbolKind { return BuiltinFunction }
func (sm *BuiltinFunctionSymbol) Name() string {
	return sm.Identifier
}

// -------------------------------------------------------
type ModuleSymbol struct {
	Identifier string
	Symbols    []Symbol
}

func (*ModuleSymbol) Kind() SymbolKind { return Module }
func (sm *ModuleSymbol) Name() string {
	return sm.Identifier
}
func (ctx *ModuleSymbol) Scope() string {
	return ctx.Identifier
}
func (ctx *ModuleSymbol) Search(ident string) Symbol {
	for _, sym := range ctx.Symbols {
		if sym.Name() == ident {
			return sym
		}
	}
	return nil
}
func (ctx *ModuleSymbol) Insert(sym Symbol) error {
	if sym := ctx.Search(sym.Name()); sym != nil {
		return fmt.Errorf("symbol '%s' already defined in scope '%s'", sym.Name(), ctx.Scope())
	}
	ctx.Symbols = append(ctx.Symbols, sym)
	return nil
}

// -------------------------------------------------------
type BuiltinTypeSymbol struct {
	ModuleSymbol
}

func (*BuiltinTypeSymbol) Kind() SymbolKind { return BuiltinType }
