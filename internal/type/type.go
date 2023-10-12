package tp

import "fmt"

type Type interface {
	fmt.Stringer
	tp()
}

type IntegerType struct{}

func (i IntegerType) String() string { return "int" }
func (t IntegerType) tp()            {}
