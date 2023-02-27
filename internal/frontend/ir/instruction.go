package ir

import (
	"fmt"
)

type Operator string

const (
	OP_LOAD = "load"
	OP_SAVE = "save"
	OP_CALL = "call"
	OP_RET  = "ret"
	OP_JMC  = "jmc"
	OP_JMP  = "jmp"
	OP_INC  = "inc"
)

type Instruction struct {
	Op   Operator
	Arg  string
	Meta map[string]string
}

func (in *Instruction) Decode() string {
	return fmt.Sprintf("%s %v", in.Op, in.Arg)
}
