package datatype

import (
	"fmt"
)

// Primitive Datatypes
type Int struct{}
type Bool struct{}
type String struct{}
type Float struct{}

// Datatype union interface
type Datatype interface {
	fmt.Stringer
	typ()
}

func (dt *Int) typ()    {}
func (dt *Bool) typ()   {}
func (dt *String) typ() {}
func (dt *Float) typ()  {}

func (dt *Int) String() string {
	return "int"
}
func (dt *Bool) String() string {
	return "bool"
}
func (dt *String) String() string {
	return "string"
}
func (dt *Float) String() string {
	return "float"
}
