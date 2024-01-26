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
}

type litteral struct {
	kind  Kind
	typ   Type
	i     int64
	f     float64
	b     bool
	s     string
	items []Litteral
}

func (lit *litteral) Kind() Kind {
	return lit.kind
}

func (lit *litteral) Type() Type {
	return lit.typ
}

func (lit *litteral) String() string {
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
	default:
		panic("can't get string representation of litteral: unexpected litteral kind")
	}
}

func (lit *litteral) ValueInt() int64 {
	if lit.kind != LI_Int {
		panic("can't get int value: litteral should be int")
	}
	return lit.i
}

func (lit *litteral) ValueFloat() float64 {
	if lit.kind != LI_Float {
		panic("can't get float value: litteral should be float")
	}
	return lit.f
}

func (lit *litteral) ValueBool() bool {
	if lit.kind != LI_Bool {
		panic("can't get bool value: litteral should be bool")
	}
	return lit.b
}

func (lit *litteral) ValueString() string {
	if lit.kind != LI_String {
		panic("can't get string value: litteral should be string")
	}
	return lit.s
}

func (lit *litteral) Len() int {
	if lit.kind != LI_List {
		panic("can't get item: litteral should be list")
	}
	return len(lit.items)
}

func (lit *litteral) Item(index int) Litteral {
	if lit.kind != LI_List {
		panic("can't get item: litteral should be list")
	}
	if index >= lit.Len() && index < 0 {
		panic("can't get item: index out of bounds")
	}
	return lit.items[index]
}

func NewIntLit(value int64) *litteral {
	return &litteral{
		kind: LI_Int,
		i:    value,
		typ:  NewIntType(),
	}
}

func NewFloatLit(value float64) *litteral {
	return &litteral{
		kind: LI_Float,
		f:    value,
		typ:  NewFloatType(),
	}
}

func NewBoolLit(value bool) *litteral {
	return &litteral{
		kind: LI_Bool,
		b:    value,
		typ:  NewBoolType(),
	}
}

func NewStringLit(value string) *litteral {
	return &litteral{
		kind: LI_String,
		s:    value,
		typ:  NewStringType(),
	}
}

func NewListLit(itemType Type, items ...Litteral) *litteral {
	for _, item := range items {
		if !itemType.Compare(item.Type()) {
			panic("can't create list: item types mismatched")
		}
	}
	return &litteral{
		kind:  LI_List,
		items: items,
		typ:   NewListType(itemType),
	}
}
