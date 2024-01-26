package runtime

import (
	"fmt"
	"grimlang/ast"
)

type Symbol interface {
	String() string
	Kind() Kind
	Name() string
	Type() Type
	Value() Litteral
	Set(Litteral)
	Fn() *ast.FunctionDecl
	Builtin() func(...Litteral) Litteral
	Path() string
}

type symbol struct {
	kind    Kind
	name    string
	typ     Type
	value   Litteral
	fn      *ast.FunctionDecl
	builtin func(...Litteral) Litteral
	path    string
}

func (symb *symbol) Kind() Kind {
	return symb.kind
}

func (symb *symbol) Name() string {
	return symb.name
}

func (symb *symbol) Type() Type {
	switch symb.kind {
	case SY_Builtin, SY_Function, SY_Variable:
		return symb.typ
	case SY_Constant:
		return symb.value.Type()
	default:
		panic("can't get type: unexpected symbol kind")
	}
}

func (symb *symbol) Value() Litteral {
	switch symb.kind {
	case SY_Variable, SY_Constant:
		return symb.value
	default:
		panic("can't get value: unexpected symbol kind")
	}
}

func (symb *symbol) Set(val Litteral) {
	if symb.kind != SY_Variable {
		panic("can't set new value to non variable symbol")
	}
	if !symb.typ.Compare(val.Type()) {
		panic("can't set new value: type mismatched")
	}
	symb.value = val
}

func (symb *symbol) Fn() *ast.FunctionDecl {
	if symb.kind != SY_Function {
		panic("can't get function: symbol should be function")
	}
	return symb.fn
}

func (symb *symbol) Builtin() func(...Litteral) Litteral {
	if symb.kind != SY_Builtin {
		panic("can't get builtin: symbol should be builtin")
	}
	return symb.builtin
}

func (symb *symbol) Path() string {
	if symb.kind != SY_Import {
		panic("can't get path: symbol should be import")
	}
	return symb.path
}

func (symb *symbol) String() string {
	switch symb.kind {
	case SY_Constant:
		return fmt.Sprintf("%s -> const %s", symb.name, symb.value.String())
	case SY_Variable:
		return fmt.Sprintf("%s -> var %s::%s", symb.name, symb.typ.String(), symb.value.String())
	case SY_Function, SY_Builtin:
		return fmt.Sprintf("%s -> %s", symb.name, symb.typ.String())
	case SY_Import:
		return fmt.Sprintf("%s -> import(%s)", symb.name, symb.path)
	default:
		panic("can't get string representation of symbol: unexpected symbol kind")
	}
}

func NewVariable(name string, typ Type, value Litteral) *symbol {
	if !typ.Compare(value.Type()) {
		panic("can't create variable: value type mismatched")
	}
	return &symbol{
		kind:  SY_Variable,
		name:  name,
		typ:   typ,
		value: value,
	}
}

func NewConstant(name string, value Litteral) *symbol {
	return &symbol{
		kind:  SY_Constant,
		name:  name,
		value: value,
	}
}

func NewFunction(name string, typ Type, fn *ast.FunctionDecl) *symbol {
	return &symbol{
		kind: SY_Function,
		name: name,
		typ:  typ,
		fn:   fn,
	}
}

func NewBuiltin(name string, typ Type, fn func(...Litteral) Litteral) *symbol {
	return &symbol{
		kind:    SY_Builtin,
		name:    name,
		typ:     typ,
		builtin: fn,
	}
}

func NewImport(name string, path string) *symbol {
	return &symbol{
		kind: SY_Import,
		name: name,
		path: path,
	}
}
