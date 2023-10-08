package ir

import "fmt"

type InstrArg interface {
	fmt.Stringer
	ia()
}

func (o Type) ia()   {}
func (o Int) ia()    {}
func (o Symbol) ia() {}

type Type string

func NewType(tp string) Type {
	return Type(tp)
}
func (o Type) String() string { return string(o) }

type Int int

func NewInt(val int) Int {
	return Int(val)
}
func (o Int) String() string { return fmt.Sprintf("%d", o) }

type Symbol string

func NewSymbol(val string) Symbol {
	return Symbol(val)
}
func (o Symbol) String() string { return string(o) }
