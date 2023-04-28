package ast

type Node interface {
	ast()
}

type Locals interface {
	locals()
}

type Globals interface {
	globals()
}

type ExpressionArg interface {
	exprArg()
}

type Program struct {
	Name string
	Body []Globals
}
