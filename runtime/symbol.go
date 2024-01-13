package runtime

import (
	"fmt"
	"strings"
)

type SymbolKind uint

const (
	SK_Variable SymbolKind = iota
	SK_Constant
	SK_Function
	SK_Module
	SK_Record
	SK_Field
)

type Symbol struct {
	// general
	Scope      string
	Identifier string
	Kind       SymbolKind
	// variable, function, field
	Type *Type
	// variable, constant, function
	Value *Object
	// module, record
	Items []*Symbol
}

func (sm *Symbol) String() string {
	switch sm.Kind {
	case SK_Variable:
		return fmt.Sprintf("%s/%s -> var (%s)%s", sm.Scope, sm.Identifier, sm.Type.String(), sm.Value.String())
	case SK_Constant:
		return fmt.Sprintf("%s/%s -> const %s", sm.Scope, sm.Identifier, sm.Value.String())
	case SK_Function:
		return fmt.Sprintf("%s/%s -> %s", sm.Scope, sm.Identifier, sm.Type.String())
	case SK_Record:
		fields := []string{}
		for _, fld := range sm.Items {
			fields = append(fields, fld.String())
		}
		return fmt.Sprintf("%s/%s -> record{%s}", sm.Scope, sm.Identifier, strings.Join(fields, " "))
	case SK_Field:
		return fmt.Sprintf("%s::%s", sm.Identifier, sm.Type.String())
	case SK_Module:
		symbols := []string{}
		for _, symbol := range sm.Items {
			symbols = append(symbols, symbol.String())
		}
		return fmt.Sprintf("%s |>\n\t%s", sm.Identifier, strings.Join(symbols, "\n\t"))
	}
	panic("can't get string for symbol " + sm.Identifier)
}

func NewVaribaleSymbol(scope string, identifier string, typ *Type, value *Object) *Symbol {
	return &Symbol{
		Kind:       SK_Variable,
		Scope:      scope,
		Identifier: identifier,
		Type:       typ,
		Value:      value,
	}
}

func NewConstantSymbol(scope string, identifier string, value *Object) *Symbol {
	return &Symbol{
		Kind:       SK_Constant,
		Scope:      scope,
		Identifier: identifier,
		Value:      value,
	}
}

func NewFunctionSymbol(scope string, identifier string, typ *Type, fn *Object) *Symbol {
	return &Symbol{
		Kind:       SK_Function,
		Scope:      scope,
		Identifier: identifier,
		Type:       typ,
		Value:      fn,
	}
}

func NewRecordSymbol(scope string, identifier string) *Symbol {
	return &Symbol{
		Kind:       SK_Record,
		Scope:      scope,
		Identifier: identifier,
	}
}

func NewModuleSymbol(scope string, identifier string) *Symbol {
	return &Symbol{
		Kind:       SK_Module,
		Scope:      scope,
		Identifier: identifier,
	}
}
