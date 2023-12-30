package object

import (
	"fmt"
	"strings"
)

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
	ItemType Object
	Items    []Object
}

type LitteralRecord struct {
	Identifier string
	Fields     []Object
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
	for i, field := range lit.Fields {
		fields = append(fields, fmt.Sprintf("%d::%s", i, field.Inspect()))
	}
	return fmt.Sprintf("%s{%s}", lit.Identifier, strings.Join(fields, " "))
}
