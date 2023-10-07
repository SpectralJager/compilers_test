package ir

import "fmt"

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
