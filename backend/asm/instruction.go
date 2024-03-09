package asm

import "fmt"

type Opcode byte

const (
	OP_Nop Opcode = iota
	OP_Halt

	OP_LocalSave
	OP_LocalLoad

	OP_Br
	OP_BrTrue

	OP_Call
	OP_Return

	OP_I64Load
	OP_I64Add
	OP_I64Sub
	OP_I64Mod
	OP_I64Eq
	OP_I64Neq
	OP_I64Gt
	OP_I64Lt

	OP_BoolAnd
)

type Instruction struct {
	Opcode Opcode
	Symbol string
	Args   [4]Value
}

func (instr *Instruction) Inspect() string {
	switch instr.Opcode {
	case OP_Nop:
		return "(nop)"
	case OP_Halt:
		return "(halt)"
	case OP_LocalLoad:
		return fmt.Sprintf("(local.load r%s)", instr.Args[0].Inspect())
	case OP_LocalSave:
		return fmt.Sprintf("(local.save r%s)", instr.Args[0].Inspect())
	case OP_Br:
		return fmt.Sprintf("(br %s)", instr.Args[0].Inspect())
	case OP_BrTrue:
		return fmt.Sprintf("(br.true %s %s)", instr.Args[0].Inspect(), instr.Args[1].Inspect())
	case OP_Call:
		return fmt.Sprintf("(call %s %s)", instr.Symbol, instr.Args[0].Inspect())
	case OP_Return:
		return fmt.Sprintf("(return %s)", instr.Args[0].Inspect())
	case OP_I64Load:
		return fmt.Sprintf("(i64.load %s)", instr.Args[0].Inspect())
	case OP_I64Add:
		return "(i64.add)"
	case OP_I64Sub:
		return "(i64.sub)"
	case OP_I64Mod:
		return "(i64.mod)"
	case OP_I64Eq:
		return "(i64.eq)"
	case OP_I64Neq:
		return "(i64.neq)"
	case OP_I64Gt:
		return "(i64.gt)"
	case OP_I64Lt:
		return "(i64.lt)"
	case OP_BoolAnd:
		return "(bool.and)"
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

func InstructionLocalLoad(regIndex int64) Instruction {
	return Instruction{
		Opcode: OP_LocalLoad,
		Args: [4]Value{
			ValueI64(regIndex),
		},
	}
}

func InstructionLocalSave(regIndex int64) Instruction {
	return Instruction{
		Opcode: OP_LocalSave,
		Args: [4]Value{
			ValueI64(regIndex),
		},
	}
}

func InstructionBr(trgt int64) Instruction {
	return Instruction{
		Opcode: OP_Br,
		Args: [4]Value{
			ValueI64(trgt),
		},
	}
}

func InstructionBrTrue(thn, els int64) Instruction {
	return Instruction{
		Opcode: OP_BrTrue,
		Args: [4]Value{
			ValueI64(thn),
			ValueI64(els),
		},
	}
}

func InstructionReturn(argc int64) Instruction {
	return Instruction{
		Opcode: OP_Return,
		Args: [4]Value{
			ValueI64(argc),
		},
	}
}

func InstructionCall(fnIdent string, argc int64) Instruction {
	return Instruction{
		Opcode: OP_Call,
		Symbol: fnIdent,
		Args: [4]Value{
			ValueI64(argc),
		},
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

func InstructionI64Sub() Instruction {
	return Instruction{
		Opcode: OP_I64Sub,
	}
}

func InstructionI64Mod() Instruction {
	return Instruction{
		Opcode: OP_I64Mod,
	}
}

func InstructionI64Eq() Instruction {
	return Instruction{
		Opcode: OP_I64Eq,
	}
}

func InstructionI64Neq() Instruction {
	return Instruction{
		Opcode: OP_I64Neq,
	}
}

func InstructionI64Gt() Instruction {
	return Instruction{
		Opcode: OP_I64Gt,
	}
}

func InstructionI64Lt() Instruction {
	return Instruction{
		Opcode: OP_I64Lt,
	}
}

func InstructionBoolAnd() Instruction {
	return Instruction{
		Opcode: OP_BoolAnd,
	}
}
