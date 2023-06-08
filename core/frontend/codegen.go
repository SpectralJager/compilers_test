package frontend

import "fmt"

type Codegenerator struct {
	ctx SemanticContext
}

func NewCodegenerator(ctx SemanticContext) *Codegenerator {
	return &Codegenerator{
		ctx: ctx,
	}
}

type bytecode uint8

const (
	OP_NIL bytecode = iota
	OP_HALT
)

var btString = map[bytecode]string{
	OP_NIL:  "nil",
	OP_HALT: "halt",
}

func (bt *bytecode) String() string {
	if val, ok := btString[*bt]; ok {
		return val
	} else {
		return "undefined"
	}
}

type Instruction struct {
	Operator bytecode
	Arg      []uint8
	ArgSize  int
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%s\t%x", i.Operator.String(), i.Arg)
}

func NewNilInstruction() *Instruction {
	return &Instruction{
		Operator: OP_NIL,
		ArgSize:  0,
	}
}

func NewHaltInstruction() *Instruction {
	return &Instruction{
		Operator: OP_HALT,
		ArgSize:  0,
	}
}
