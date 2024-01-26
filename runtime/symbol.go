package runtime

import "grimlang/ast"

type Symbol interface {
	Kind() Kind
	Name() string
	Type() Type
	Value() litteral
	Fn() *ast.FunctionDecl
	Builtin() func(...Litteral) Litteral
}
