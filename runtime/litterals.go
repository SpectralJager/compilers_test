package runtime

import (
	"fmt"
	"grimlang/ast"
	"strings"
)

type _litteral struct {
	kind Kind
	typ  Type

	i int64
	f float64
	b bool
	s string

	items []Litteral

	fn         *ast.FunctionDecl
	fn_builtin func(...Litteral) (Litteral, error)

	fields []Litteral
}

func (lit _litteral) Kind() Kind {
	return lit.kind
}

func (lit _litteral) String() string {
	switch lit.kind {
	case Int:
		return fmt.Sprintf("%d", lit.i)
	case Float:
		return fmt.Sprintf("%f", lit.f)
	case Bool:
		return fmt.Sprintf("%v", lit.b)
	case String:
		return lit.s
	case List:
		items := []string{}
		for _, item := range lit.items {
			items = append(items, item.String())
		}
		return fmt.Sprintf("'(%s)", strings.Join(items, " "))
	case Record:
		flds := []string{}
		for i, fld := range lit.fields {
			fldType := lit.typ.Field(i)
			flds = append(flds, fmt.Sprintf("%s::%s", fldType.Name(), fld.String()))
		}
		return fmt.Sprintf("%s{%s}", lit.typ.Name(), strings.Join(flds, " "))
	case Function:
		return lit.typ.String()
	default:
		panic("can't get value of " + lit.kind.String())
	}
}

func (lit _litteral) Type() Type {
	return lit.typ
}

func (lit _litteral) Int() int64 {
	if lit.kind != Int {
		panic("can't get int value for " + lit.kind.String())
	}
	return lit.i
}

func (lit _litteral) Float() float64 {
	if lit.kind != Float {
		panic("can't get float value for " + lit.kind.String())
	}
	return lit.f
}

func (lit _litteral) Bool() bool {
	if lit.kind != Bool {
		panic("can't get bool value for " + lit.kind.String())
	}
	return lit.b
}

func (lit _litteral) Str() string {
	if lit.kind != String {
		panic("can't get string value for " + lit.kind.String())
	}
	return lit.s
}

func (lit _litteral) Item(index int) Litteral {
	if lit.kind != List {
		panic("can't get item for " + lit.kind.String())
	}
	if index < 0 || index >= lit.Len() {
		panic("index out of bounds")
	}
	return lit.items[index]
}

func (lit _litteral) Len() int {
	if lit.kind != List {
		panic("can't get len for " + lit.kind.String())
	}
	return len(lit.items)
}

func (lit _litteral) Field(index int) Litteral {
	if lit.kind != Record {
		panic("can't get field for " + lit.kind.String())
	}
	if index < 0 || index >= lit.typ.NumField() {
		panic("index out of bounds")
	}
	return lit.fields[index]
}

func (lit _litteral) FieldByName(name string) Litteral {
	if lit.kind != Record {
		panic("can't get field for " + lit.kind.String())
	}
	index := lit.typ.FieldIndex(name)
	return lit.Field(index)
}

func (lit _litteral) Call(ctx Context, args ...Litteral) (Litteral, error) {
	switch lit.kind {
	case Builtin:
		return lit.fn_builtin(args...)
	case Function:
		panic("unimplemented")
	default:
		panic("can't call litteral of " + lit.kind.String())
	}
}

func NewIntLit(value int64) _litteral {
	return _litteral{
		kind: Int,
		typ:  NewIntType(),
		i:    value,
	}
}

func NewFloatLit(value float64) _litteral {
	return _litteral{
		kind: Float,
		typ:  NewFloatType(),
		f:    value,
	}
}

func NewBoolLit(value bool) _litteral {
	return _litteral{
		kind: Bool,
		typ:  NewBoolType(),
		b:    value,
	}
}

func NewStringLit(value string) _litteral {
	return _litteral{
		kind: String,
		typ:  NewStringType(),
		s:    value,
	}
}

func NewListLit(itemType Type, items ...Litteral) _litteral {
	for i, item := range items {
		if !itemType.Compare(item.Type()) {
			panic(fmt.Sprintf("%dth item should be type of %s", i, itemType.String()))
		}
	}
	return _litteral{
		kind:  List,
		typ:   NewListType(itemType),
		items: items,
	}
}

func NewFunctionLit(typ Type, fn *ast.FunctionDecl) _litteral {
	return _litteral{
		kind: Function,
		typ:  typ,
		fn:   fn,
	}
}

func NewBuiltinLit(typ Type, fn func(...Litteral) (Litteral, error)) _litteral {
	return _litteral{
		kind:       Builtin,
		typ:        typ,
		fn_builtin: fn,
	}
}

func NewRecordLit(typ Type, flds ...Litteral) _litteral {
	if typ.NumField() != len(flds) {
		panic("can't create record litteral: mismatch fields")
	}
	for i, fld := range flds {
		fldType := typ.Field(i)
		if !fldType.Type().Compare(fld.Type()) {
			panic("can't create record litteral: mismatch fields")
		}
	}
	return _litteral{
		kind:   Record,
		typ:    typ,
		fields: flds,
	}
}
