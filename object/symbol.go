package object

import (
	"fmt"
	"grimlang/ast"
	"strings"
)

type SymbolVariable struct {
	Identifier string
	ValueType  Object
	Value      Object
}

type SymbolConstant struct {
	Identifier string
	Value      Object
}

type SymbolFunction struct {
	Identifier   string
	FunctionType Object
	Fn           *ast.FunctionDecl
}

type SymbolBuiltin struct {
	Identifier   string
	FunctionType Object
	Fn           func(...Object) (Object, error)
}

type SymbolModule struct {
	Identifier string
	Symbols    []Object
}

type SymbolRecord struct {
	Identifier string
	RecordType Object
	Fields     []Object
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
