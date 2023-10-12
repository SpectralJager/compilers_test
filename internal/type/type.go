package tp

import "fmt"

type Type interface {
	fmt.Stringer
	tp()
}

func (t IntegerType) tp() {}

type IntegerType struct{}

func (i IntegerType) String() string { return "int" }
