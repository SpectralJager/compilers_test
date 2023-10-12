package tp

import "fmt"

type Type interface {
	fmt.Stringer
	tp()
}

type IntegerType struct{}

func (i IntegerType) String() string { return "int" }
func (t IntegerType) tp()            {}

type BooleanType struct{}

func (i BooleanType) String() string { return "bool" }
func (t BooleanType) tp()            {}

type VariaticType struct {
	Subtype Type
}

func (i VariaticType) String() string { return "..." + i.Subtype.String() }
func (t VariaticType) tp()            {}
