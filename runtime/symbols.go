package runtime

import (
	"fmt"
	"strings"
)

type _symbol struct {
	kind  Kind
	name  string
	typ   Type
	value Litteral

	symbols map[string]Symbol
}

func (symb _symbol) Kind() Kind {
	return symb.kind
}

func (symb _symbol) Name() string {
	return symb.name
}

func (symb _symbol) String() string {
	switch symb.kind {
	case Variable:
		return fmt.Sprintf("%s -> var::%s(%s)", symb.name, symb.typ, symb.value)
	case Constant:
		return fmt.Sprintf("%s -> const::%s", symb.name, symb.value)
	case Function:
		return fmt.Sprintf("%s -> %s", symb.name, symb.typ)
	case Record:
		return fmt.Sprintf("%s", symb.typ)
	case Module:
		symbols := []string{}
		for _, symbol := range symb.symbols {
			symbols = append(symbols, symbol.String())
		}
		return fmt.Sprintf("=== %s \n\t%s\n", symb.name, strings.Join(symbols, "\n\t"))
	default:
		panic("can't get string for " + symb.kind.String())
	}
}

func (symb _symbol) Type() Type {
	switch symb.kind {
	case Variable, Record:
		return symb.typ
	case Constant, Function:
		return symb.value.Type()
	default:
		panic("can't get type for " + symb.kind.String())
	}
}

func (symb _symbol) Value() Litteral {
	switch symb.kind {
	case Constant, Variable, Function:
		return symb.value
	default:
		panic("can't get value for " + symb.kind.String())
	}
}

func (symb _symbol) Set(value Litteral) error {
	if symb.kind != Variable {
		return fmt.Errorf("can't set value to %s", symb.kind.String())
	}
	if !symb.typ.Compare(value.Type()) {
		return fmt.Errorf("symbol and value types mismatch")
	}
	symb.value = value
	return nil
}

func (ctx _symbol) Scope() string {
	if ctx.kind != Module && ctx.kind != Record {
		panic("can't get scope for " + ctx.kind.String())
	}
	return ctx.name
}

func (ctx _symbol) Search(name string) Symbol {
	if ctx.kind != Module && ctx.kind != Record {
		panic("can't search symbol for " + ctx.kind.String())
	}
	symb, ok := ctx.symbols[name]
	if !ok {
		return nil
	}
	return symb
}

func (ctx _symbol) Insert(symb Symbol) error {
	if ctx.kind != Module {
		panic("can't insert symbol for " + ctx.kind.String())
	}
	if ctx.Search(symb.Name()) != nil {
		return fmt.Errorf("symbol '%s' already defined in scope '%s'", symb.Name(), ctx.Scope())
	}
	ctx.symbols[symb.Name()] = symb
	return nil
}

func (ctx _symbol) Update(symb Symbol) error {
	if ctx.kind != Module {
		panic("can't update symbol for " + ctx.kind.String())
	}
	old := ctx.Search(symb.Name())
	if old == nil {
		return fmt.Errorf("symbols '%s' not defined in scope '%s'", symb.Name(), ctx.Scope())
	}
	if old.Kind() != Variable {
		return fmt.Errorf("can't update non variable symbol '%s'", old.Name())
	}
	if !old.Type().Compare(symb.Type()) {
		return fmt.Errorf("old and updated variable '%s' types mismatched", old.Name())
	}
	ctx.symbols[old.Name()] = symb
	return nil
}

func NewVariableSymbol(name string, typ Type, value Litteral) _symbol {
	return _symbol{
		kind:  Variable,
		name:  name,
		typ:   typ,
		value: value,
	}
}

func NewConstantSymbol(name string, value Litteral) _symbol {
	return _symbol{
		kind:  Constant,
		name:  name,
		value: value,
	}
}

func NewFunctionSymbol(name string, typ Type, fn Litteral) _symbol {
	return _symbol{
		kind:  Function,
		name:  name,
		typ:   typ,
		value: fn,
	}
}

func NewRecordSymbol(typ Type) _symbol {
	return _symbol{
		kind: Record,
		name: typ.Name(),
		typ:  typ,
	}
}

func NewModuleSymbol(name string) _symbol {
	return _symbol{
		kind:    Module,
		name:    name,
		symbols: map[string]Symbol{},
	}
}

func NewContext(name string, parent Context) Context {
	return _symbol{
		kind:    Module,
		name:    name,
		symbols: map[string]Symbol{},
	}
}
