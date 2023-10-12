package ast

import "fmt"

type AST interface {
	fmt.Stringer
	ast()
}

// unions

type GLOBAL interface {
	AST
	glob()
}
type LOCAL interface {
	AST
	locl()
}
type EXPR interface {
	AST
	expr()
}
type ATOM interface {
	AST
	atom()
}
