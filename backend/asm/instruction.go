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

	OP_Rotate
	OP_Duplicate

	OP_I64Load
	OP_I64Add
	OP_I64Sub
	OP_I64Mul
	OP_I64Div
	OP_I64Neg
	OP_I64Mod
	OP_I64Eq
	OP_I64Neq
	OP_I64Gt
	OP_I64Lt
	OP_I64Geq
	OP_I64Leq

	OP_F64Load
	OP_F64Add
	OP_F64Sub
	OP_F64Mul
	OP_F64Div
	OP_F64Neg
	OP_F64Eq
	OP_F64Neq
	OP_F64Gt
	OP_F64Lt
	OP_F64Geq
	OP_F64Leq

	OP_BoolLoad
	OP_BoolEq
	OP_BoolNeq
	OP_BoolAnd
	OP_BoolOr
	OP_BoolNot

	OP_StringLoad
	OP_StringConcat
	OP_StringEq
	OP_StringNeq

	OP_ListLoad
	OP_ListConstruct
	OP_ListGet
	OP_ListSet
	OP_ListInsert
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
	case OP_Rotate:
		return "(rot)"
	case OP_Duplicate:
		return "(dup)"
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
	case OP_I64Mul:
		return "(i64.mul)"
	case OP_I64Div:
		return "(i64.div)"
	case OP_I64Neg:
		return "(i64.neg)"
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
	case OP_I64Geq:
		return "(i64.gt)"
	case OP_I64Leq:
		return "(i64.gt)"
	case OP_F64Load:
		return fmt.Sprintf("(f64.load %s)", instr.Args[0].Inspect())
	case OP_F64Add:
		return "(f64.add)"
	case OP_F64Sub:
		return "(f64.sub)"
	case OP_F64Mul:
		return "(f64.mul)"
	case OP_F64Div:
		return "(f64.div)"
	case OP_F64Neg:
		return "(f64.neg)"
	case OP_F64Eq:
		return "(f64.eq)"
	case OP_F64Neq:
		return "(f64.neq)"
	case OP_F64Gt:
		return "(f64.gt)"
	case OP_F64Lt:
		return "(f64.lt)"
	case OP_F64Geq:
		return "(f64.gt)"
	case OP_F64Leq:
		return "(f64.gt)"
	case OP_BoolLoad:
		return fmt.Sprintf("(bool.load %s)", instr.Args[0].Inspect())
	case OP_BoolEq:
		return "(bool.eq)"
	case OP_BoolNeq:
		return "(bool.neq)"
	case OP_BoolAnd:
		return "(bool.and)"
	case OP_BoolOr:
		return "(bool.or)"
	case OP_BoolNot:
		return "(bool.not)"
	case OP_StringLoad:
		return fmt.Sprintf("(str.load %s)", instr.Args[0].Inspect())
	case OP_StringConcat:
		return "(str.concat)"
	case OP_StringEq:
		return "(str.eq)"
	case OP_StringNeq:
		return "(str.neq)"
	case OP_ListLoad:
		return fmt.Sprintf("(list.load %s)", instr.Args[0].Inspect())
	case OP_ListConstruct:
		return fmt.Sprintf("(list.construct %s)", instr.Args[0].Inspect())
	case OP_ListGet:
		return "(list.get)"
	case OP_ListSet:
		return "(list.set)"
	case OP_ListInsert:
		return "(list.insert)"
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

func InstructionRotate() Instruction {
	return Instruction{
		Opcode: OP_Rotate,
	}
}

func InstructionDuplicate() Instruction {
	return Instruction{
		Opcode: OP_Duplicate,
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

func InstructionI64Mul() Instruction {
	return Instruction{
		Opcode: OP_I64Mul,
	}
}

func InstructionI64Div() Instruction {
	return Instruction{
		Opcode: OP_I64Div,
	}
}

func InstructionI64Neg() Instruction {
	return Instruction{
		Opcode: OP_I64Neg,
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

func InstructionI64Geq() Instruction {
	return Instruction{
		Opcode: OP_I64Geq,
	}
}

func InstructionI64Leq() Instruction {
	return Instruction{
		Opcode: OP_I64Leq,
	}
}

func InstructionF64Load(val float64) Instruction {
	return Instruction{
		Opcode: OP_F64Load,
		Args: [4]Value{
			ValueF64(val),
		},
	}
}

func InstructionF64Add() Instruction {
	return Instruction{
		Opcode: OP_F64Add,
	}
}

func InstructionF64Sub() Instruction {
	return Instruction{
		Opcode: OP_F64Sub,
	}
}

func InstructionF64Mul() Instruction {
	return Instruction{
		Opcode: OP_F64Mul,
	}
}

func InstructionF64Div() Instruction {
	return Instruction{
		Opcode: OP_F64Div,
	}
}

func InstructionF64Neg() Instruction {
	return Instruction{
		Opcode: OP_F64Neg,
	}
}

func InstructionF64Eq() Instruction {
	return Instruction{
		Opcode: OP_F64Eq,
	}
}

func InstructionF64Neq() Instruction {
	return Instruction{
		Opcode: OP_F64Neq,
	}
}

func InstructionF64Gt() Instruction {
	return Instruction{
		Opcode: OP_F64Gt,
	}
}

func InstructionF64Lt() Instruction {
	return Instruction{
		Opcode: OP_F64Lt,
	}
}

func InstructionF64Geq() Instruction {
	return Instruction{
		Opcode: OP_F64Geq,
	}
}

func InstructionF64Leq() Instruction {
	return Instruction{
		Opcode: OP_F64Leq,
	}
}

func InstructionBoolLoad(val bool) Instruction {
	return Instruction{
		Opcode: OP_BoolLoad,
		Args: [4]Value{
			ValueBool(val),
		},
	}
}

func InstructionBoolEq() Instruction {
	return Instruction{
		Opcode: OP_BoolEq,
	}
}

func InstructionBoolNeq() Instruction {
	return Instruction{
		Opcode: OP_BoolNeq,
	}
}

func InstructionBoolAnd() Instruction {
	return Instruction{
		Opcode: OP_BoolAnd,
	}
}

func InstructionBoolOr() Instruction {
	return Instruction{
		Opcode: OP_BoolOr,
	}
}

func InstructionBoolNot() Instruction {
	return Instruction{
		Opcode: OP_BoolNot,
	}
}

func InstructionStringLoad(val string) Instruction {
	return Instruction{
		Opcode: OP_StringLoad,
		Args: [4]Value{
			ValueString(val),
		},
	}
}

func InstructionStringConcat() Instruction {
	return Instruction{
		Opcode: OP_StringConcat,
	}
}

func InstructionStringEq() Instruction {
	return Instruction{
		Opcode: OP_StringEq,
	}
}

func InstructionStringNeq() Instruction {
	return Instruction{
		Opcode: OP_StringNeq,
	}
}

func InstructionListLoad(values ...Value) Instruction {
	return Instruction{
		Opcode: OP_ListLoad,
		Args: [4]Value{
			ValueList(values...),
		},
	}
}

func InstructionListConstruct(itemCount int64) Instruction {
	return Instruction{
		Opcode: OP_ListConstruct,
		Args: [4]Value{
			ValueI64(itemCount),
		},
	}
}
func InstructionListGet() Instruction {
	return Instruction{
		Opcode: OP_ListGet,
	}
}
func InstructionListSet() Instruction {
	return Instruction{
		Opcode: OP_ListSet,
	}
}
func InstructionListInsert() Instruction {
	return Instruction{
		Opcode: OP_ListInsert,
	}
}
