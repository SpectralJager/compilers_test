package object

import (
	"fmt"
	"strings"
)

type Litteral interface {
	Object
	Type() DType
}

type LitteralNull struct{}

type LitteralInt struct {
	Value int
}

type LitteralFloat struct {
	Value float64
}

type LitteralString struct {
	Value string
}

type LitteralBool struct {
	Value bool
}

type LitteralList struct {
	ItemType DType
	Items    []Litteral
}

type LitteralRecord struct {
	RecordType DTypeRecord
	Fields     []Symbol
}

func (*LitteralNull) Kind() ObjectKind   { return NullLitteral }
func (*LitteralInt) Kind() ObjectKind    { return IntLitteral }
func (*LitteralFloat) Kind() ObjectKind  { return FloatLitteral }
func (*LitteralString) Kind() ObjectKind { return StringLitteral }
func (*LitteralBool) Kind() ObjectKind   { return BoolLitteral }
func (*LitteralList) Kind() ObjectKind   { return ListLitteral }
func (*LitteralRecord) Kind() ObjectKind { return RecordLitteral }

func (*LitteralNull) Inspect() string       { return "null" }
func (lit *LitteralInt) Inspect() string    { return fmt.Sprintf("%d", lit.Value) }
func (lit *LitteralFloat) Inspect() string  { return fmt.Sprintf("%f", lit.Value) }
func (lit *LitteralString) Inspect() string { return lit.Value }
func (lit *LitteralBool) Inspect() string   { return fmt.Sprintf("%v", lit.Value) }
func (lit *LitteralList) Inspect() string {
	items := []string{}
	for _, item := range lit.Items {
		items = append(items, item.Inspect())
	}
	return fmt.Sprintf("'(%s)", strings.Join(items, " "))
}
func (lit *LitteralRecord) Inspect() string {
	fields := []string{}
	for _, field := range lit.Fields {
		fields = append(fields, field.Inspect())
	}
	return fmt.Sprintf("%s/%s{%s}", lit.RecordType.Scope, lit.RecordType.Identifier, strings.Join(fields, " "))
}

func (*LitteralNull) Type() DType {
	return &DTypeAny{}
}
func (*LitteralInt) Type() DType {
	return &DTypeInt{}
}
func (*LitteralFloat) Type() DType {
	return &DTypeFloat{}
}
func (*LitteralString) Type() DType {
	return &DTypeString{}
}
func (*LitteralBool) Type() DType {
	return &DTypeBool{}
}
func (lt *LitteralList) Type() DType {
	return &DTypeList{
		ChildType: lt.ItemType,
	}
}
func (lt *LitteralRecord) Type() DType {
	return &lt.RecordType
}

func (ctx *LitteralRecord) Scope() string {
	return ctx.RecordType.Identifier
}

func (ctx *LitteralRecord) Search(ident string) Symbol {
	for _, symb := range ctx.Fields {
		if symb.Name() == ident {
			return symb
		}
	}
	return nil
}

func (ctx *LitteralRecord) Insert(symbol Symbol) error {
	return nil
}
