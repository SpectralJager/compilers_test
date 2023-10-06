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

type Integer struct {
	Value int
}

func NewInteger(value int) *Integer {
	return &Integer{Value: value}
}

func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type Symbol struct {
	Name      string
	SubSymbol *Symbol
}

func NewSymbol(name string, sub *Symbol) *Symbol {
	return &Symbol{Name: name, SubSymbol: sub}
}

func (s *Symbol) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s", s.Name)
	if s.SubSymbol != nil {
		fmt.Fprintf(&buf, "/%s", s.SubSymbol.String())
	}
	return buf.String()
}
