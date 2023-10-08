package ir

import (
	"fmt"
	"strings"
)

type InstrArg interface {
	fmt.Stringer
	instrArg()
}

func (i *Integer) instrArg() {}
func (i *Symbol) instrArg()  {}
func (i *Type) instrArg()    {}

type Integer struct {
	Value int
}

func NewInteger(value int) *Integer {
	return &Integer{Value: value}
}

func (i *Integer) String() string { return fmt.Sprintf("%d", i.Value) }

type Symbol struct {
	Primary   string
	Secondary *Symbol
}

func NewSymbol(p string, s *Symbol) *Symbol {
	return &Symbol{Primary: p, Secondary: s}
}

func (i *Symbol) String() string {
	if i.Secondary == nil {
		return i.Primary
	} else {
		return fmt.Sprintf("%s/%s", i.Primary, i.Secondary)
	}
}

type Type struct {
	Primary *Symbol
	Generic []*Type
}

func NewType(p *Symbol, g ...*Type) *Type {
	return &Type{Primary: p, Generic: g}
}

func (i *Type) String() string {
	if i.Generic == nil {
		return i.Primary.String()
	} else {
		var buf strings.Builder
		fmt.Fprintf(&buf, "%s<", i.Primary.String())
		for _, g := range i.Generic {
			fmt.Fprintf(&buf, " %s", g.String())
		}
		fmt.Fprint(&buf, ">")
		return buf.String()
	}
}
