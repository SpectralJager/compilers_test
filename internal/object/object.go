package object

import "fmt"

type Object interface {
	obj()
}

func (o Integer) obj() {}

type Integer struct {
	Value int
}

func NewInt(val int) *Integer {
	return &Integer{Value: val}
}

func (o Integer) String() string { return fmt.Sprintf("%d", o.Value) }
