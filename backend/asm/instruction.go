package asm

import "fmt"

type Opcode byte

const (
	OP_Nop Opcode = iota
	OP_Halt

	OP_I64Load
	OP_I64Add
)

type Instruction struct {
	Opcode Opcode
	Args   [4]Value
}

func (instr *Instruction) Inspect() string {
	switch instr.Opcode {
	case OP_Nop:
		return "(nop)"
	case OP_Halt:
		return "(halt)"
	case OP_I64Load:
		return fmt.Sprintf("(i64.load %s)", instr.Args[0].Inspect())
	case OP_I64Add:
		return "(i64.add)"
	default:
		return "(unenxpected or illigal instruction)"
	}
}

func InstructionNop() Instruction {
	return Instruction{
		Opcode: OP_Nop,
	}
}

func InstructionHalt() Instruction {
	return Instruction{
		Opcode: OP_Halt,
	}
}

func InstructionI64Load(val int64) Instruction {
	return Instruction{
		Opcode: OP_I64Load,
		Args: [4]Value{
			ValueI64(val),
		},
	}
}

func InstructionI64Add() Instruction {
	return Instruction{
		Opcode: OP_I64Add,
	}
}
