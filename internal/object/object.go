package object

import (
	"fmt"
	"grimlang/internal/datatype"
)

type Integer struct {
	Datatype datatype.Int
	Value    int
}
type Boolean struct {
	Datatype datatype.Bool
	Value    bool
}
type String struct {
	Datatype datatype.String
	Value    string
}
type Float struct {
	Datatype datatype.Float
	Value    float64
}

type Object interface {
	fmt.Stringer
	obj()
}

func (ob *Integer) obj() {}
func (ob *Boolean) obj() {}
func (ob *Float) obj()   {}
func (ob *String) obj()  {}

func (ob *Integer) String() string {
	return fmt.Sprintf("%d", ob.Value)
}
func (ob *Boolean) String() string {
	return fmt.Sprintf("%v", ob.Value)
}
func (ob *Float) String() string {
	return fmt.Sprintf("%f", ob.Value)
}
func (ob *String) String() string {
	return fmt.Sprintf("\"%s\"", ob.Value)
}
