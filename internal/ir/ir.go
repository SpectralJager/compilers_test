package ir

import "fmt"

type IR interface {
	fmt.Stringer
}

type Block interface {
	IR
	block()
}

type Instruction interface {
	IR
	instr()
}
