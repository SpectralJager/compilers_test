package repr

import "fmt"

type ObjectKind string

const (
	ObjectInt ObjectKind = "int"
)

type Object struct {
	Kind  ObjectKind
	Value any
}

func NewObject(value any) *Object {
	return &Object{Value: value, Kind: ObjectInt}
}

func (o Object) String() string {
	return fmt.Sprintf("%s{%v}", o.Kind, o.Value)
}
