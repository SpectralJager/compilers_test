package ast

type Type interface {
	AST
	tp()
}

type IntType struct {
	T struct{} `parser:"'int'"`
}

func (i *IntType) String() string {
	return "int"
}

// part of ...
func (i *IntType) tp()  {}
func (i *IntType) ast() {}
