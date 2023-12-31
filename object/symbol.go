package object

import (
	"fmt"
	"grimlang/ast"
	"strings"
)

type Symbol interface {
	Object
	Name() string
}

type SymbolVariable struct {
	Identifier string
	ValueType  DType
	Value      Litteral
}

type SymbolConstant struct {
	Identifier string
	Value      Litteral
}

type SymbolFunction struct {
	Identifier   string
	FunctionType DTypeFunction
	Fn           *ast.FunctionDecl
}

type SymbolBuiltin struct {
	Identifier   string
	FunctionType DTypeFunction
	Fn           func(...Litteral) (Litteral, error)
}

type SymbolModule struct {
	Identifier string
	Symbols    []Symbol
}

type SymbolRecord struct {
	Identifier string
	RecordType DTypeRecord
	Fields     []Symbol
}

func (*SymbolVariable) Kind() ObjectKind { return VariableSymbol }
func (*SymbolConstant) Kind() ObjectKind { return ConstantSymbol }
func (*SymbolFunction) Kind() ObjectKind { return FunctionSymbol }
func (*SymbolBuiltin) Kind() ObjectKind  { return BuiltinSymbol }
func (*SymbolModule) Kind() ObjectKind   { return ModuleSymbol }
func (*SymbolRecord) Kind() ObjectKind   { return RecordSymbol }

func (sm *SymbolVariable) Inspect() string {
	return fmt.Sprintf("%s -> var (%s) %s", sm.Identifier, sm.ValueType.Inspect(), sm.Value.Inspect())
}
func (sm *SymbolConstant) Inspect() string {
	return fmt.Sprintf("%s -> const %s", sm.Identifier, sm.Value.Inspect())
}
func (sm *SymbolFunction) Inspect() string {
	return fmt.Sprintf("%s -> %s", sm.Identifier, sm.FunctionType.Inspect())
}
func (sm *SymbolBuiltin) Inspect() string {
	return fmt.Sprintf("%s -> %s", sm.Identifier, sm.FunctionType.Inspect())
}
func (sm *SymbolModule) Inspect() string {
	symbols := []string{}
	for _, symb := range sm.Symbols {
		symbols = append(symbols, symb.Inspect())
	}
	return fmt.Sprintf("===%s\n\t%s\n", sm.Identifier, strings.Join(symbols, "\n\t"))
}
func (sm *SymbolRecord) Inspect() string {
	return fmt.Sprintf("%s -> %s", sm.Identifier, sm.RecordType.Inspect())
}

func (sm *SymbolVariable) Name() string { return sm.Identifier }
func (sm *SymbolConstant) Name() string { return sm.Identifier }
func (sm *SymbolFunction) Name() string { return sm.Identifier }
func (sm *SymbolBuiltin) Name() string  { return sm.Identifier }
func (sm *SymbolModule) Name() string   { return sm.Identifier }
func (sm *SymbolRecord) Name() string   { return sm.Identifier }

func (ctx *SymbolModule) Scope() string {
	return ctx.Identifier
}
func (ctx *SymbolRecord) Scope() string {
	return ctx.Identifier
}

func (ctx *SymbolModule) Search(ident string) Symbol {
	for _, symb := range ctx.Symbols {
		if symb.Name() == ident {
			return symb
		}
	}
	return nil
}
func (ctx *SymbolRecord) Search(ident string) Symbol {
	for _, symb := range ctx.Fields {
		if symb.Name() == ident {
			return symb
		}
	}
	return nil
}

func (ctx *SymbolModule) Insert(symbol Symbol) error {
	if ctx.Search(symbol.Name()) != nil {
		return fmt.Errorf("symbols %s already defined in %s", symbol.Name(), ctx.Identifier)
	}
	ctx.Symbols = append(ctx.Symbols, symbol)
	return nil
}
func (ctx *SymbolRecord) Insert(symbol Symbol) error {
	if ctx.Search(symbol.Name()) != nil {
		return fmt.Errorf("symbols %s already defined in %s", symbol.Name(), ctx.Identifier)
	}
	ctx.Fields = append(ctx.Fields, symbol)
	return nil
}
