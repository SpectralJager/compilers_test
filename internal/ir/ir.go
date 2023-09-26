package ir

import "fmt"

type IR interface {
	fmt.Stringer
}

type INSTR interface {
	IR
	instr()
}

type OBJ interface {
	IR
	obj()
}

type Program struct {
	Name  string
	Types []Type
	Init  []INSTR
	Blocs []Bloc
}

type Type struct {
	Kind    typeKind
	Subtype []Type
}
