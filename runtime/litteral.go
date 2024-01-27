package runtime

import (
	"fmt"
	"strings"
)

type Litteral interface {
	Kind() Kind
	String() string
	Type() Type
	ValueInt() int64
	ValueFloat() float64
	ValueBool() bool
	ValueString() string
	Len() int
	Item(int) Litteral
	Field(string) Symbol
}

type litteral struct {
	kind   Kind
	typ    Type
	i      int64
	f      float64
	b      bool
	s      string
	items  []Litteral
	fields []Symbol
}

func (lit litteral) Kind() Kind {
	return lit.kind
}

func (lit litteral) Type() Type {
	return lit.typ
}

func (lit litteral) String() string {
	switch lit.kind {
	case LI_Int:
		return fmt.Sprintf("%d", lit.i)
	case LI_Float:
		return fmt.Sprintf("%f", lit.f)
	case LI_Bool:
		return fmt.Sprintf("%v", lit.b)
	case LI_String:
		return lit.s
	case LI_List:
		items := []string{}
		itemType := lit.typ.Item()
		for _, item := range lit.items {
			if !itemType.Compare(item.Type()) {
				panic("can't get string representation of litteral: list items types mismatched")
			}
			items = append(items, item.String())
		}
		return fmt.Sprintf("%s{%s}", lit.typ.String(), strings.Join(items, " "))
	case LI_Record:
		fields := []string{}
		for _, fld := range lit.fields {
			fields = append(fields, fmt.Sprintf("%s::%s", fld.Name(), fld.Value().String()))
		}
		return fmt.Sprintf("%s{%s}", lit.typ.String(), strings.Join(fields, " "))
	default:
		panic("can't get string representation of litteral: unexpected litteral kind")
	}
}

func (lit litteral) ValueInt() int64 {
	if lit.kind != LI_Int {
		panic("can't get int value: litteral should be int")
	}
	return lit.i
}

func (lit litteral) ValueFloat() float64 {
	if lit.kind != LI_Float {
		panic("can't get float value: litteral should be float")
	}
	return lit.f
}

func (lit litteral) ValueBool() bool {
	if lit.kind != LI_Bool {
		panic("can't get bool value: litteral should be bool")
	}
	return lit.b
}

func (lit litteral) ValueString() string {
	if lit.kind != LI_String {
		panic("can't get string value: litteral should be string")
	}
	return lit.s
}

func (lit litteral) Len() int {
	if lit.kind != LI_List {
		panic("can't get item: litteral should be list")
	}
	return len(lit.items)
}

func (lit litteral) Item(index int) Litteral {
	if lit.kind != LI_List {
		panic("can't get item: litteral should be list")
	}
	if index >= lit.Len() && index < 0 {
		panic("can't get item: index out of bounds")
	}
	return lit.items[index]
}

func (lit litteral) Field(name string) Symbol {
	if lit.kind != LI_Record {
		panic("can't get field: litteral should be record")
	}
	fld, i := lit.typ.Field(name)
	if i == -1 {
		panic("can't get field: field not found")
	}
	if !fld.Type().Compare(lit.fields[i].Type()) {
		panic("can't get field: field type mismatched")
	}
	return lit.fields[i]
}

func NewIntLit(value int64) litteral {
	return litteral{
		kind: LI_Int,
		i:    value,
		typ:  NewIntType(),
	}
}

func NewFloatLit(value float64) litteral {
	return litteral{
		kind: LI_Float,
		f:    value,
		typ:  NewFloatType(),
	}
}

func NewBoolLit(value bool) litteral {
	return litteral{
		kind: LI_Bool,
		b:    value,
		typ:  NewBoolType(),
	}
}

func NewStringLit(value string) litteral {
	return litteral{
		kind: LI_String,
		s:    value,
		typ:  NewStringType(),
	}
}

func NewListLit(itemType Type, items ...Litteral) litteral {
	for _, item := range items {
		if !itemType.Compare(item.Type()) {
			panic("can't create list: item types mismatched")
		}
	}
	return litteral{
		kind:  LI_List,
		items: items,
		typ:   NewListType(itemType),
	}
}

func NewRecordLit(typ Type, fields ...Symbol) litteral {
	if typ.Kind() != TY_Record {
		panic("can't create record: type should be record")
	}
	if len(fields) != typ.NumFields() {
		panic("can't create record: number of fields not equal")
	}
	for i, fld := range fields {
		if fld.Kind() != SY_Variable {
			panic("can't create record: field should be variable")
		}
		if !fld.Type().Compare(typ.FieldByIndex(i).Type()) {
			panic("can't create record: field type mismatched")
		}
	}
	return litteral{
		kind:   LI_Record,
		typ:    typ,
		fields: fields,
	}
}
