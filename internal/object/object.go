package object

import "fmt"

type Object struct {
	Kind ObjectKind
	I    int
	B    bool
}

func NewInt(val int) Object {
	return Object{Kind: Int, I: val}
}

func NewBoolean(val bool) Object {
	return Object{Kind: Boolean, B: val}
}

func (o Object) Value() (any, error) {
	switch o.Kind {
	case Int:
		return o.I, nil
	case Boolean:
		return o.B, nil
	default:
		return nil, fmt.Errorf("can't get value for kind %s", o.Kind)
	}
}

func (o Object) String() string {
	switch o.Kind {
	case Int:
		return fmt.Sprintf("%d", o.I)
	case Boolean:
		return fmt.Sprintf("%v", o.B)
	default:
		return fmt.Sprintf("undefined{%s}", o.Kind)
	}
}

type ObjectKind string

const (
	Int     ObjectKind = "int"
	Boolean ObjectKind = "bool"
)
