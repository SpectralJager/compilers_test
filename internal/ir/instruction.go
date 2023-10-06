package ir

import (
	"fmt"
	"strings"
)

type Instruction struct {
	Operator OpType
	Args     []InstrArg
}

func (i *Instruction) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s ", i.Operator)
	for _, arg := range i.Args {
		fmt.Fprintf(&buf, "%s ", arg.String())
	}
	return buf.String()
}

type OpType string

const (
	NULL OpType = "null"

	VarNew  = "var.new"
	VarSet  = "var.set"
	VarPop  = "var.pop"
	VarPush = "var.push"

	ConstPush = "const.push"

	Call  = "call"
	Callb = "callb"

	BlockBegin = "block.begin"
	BlockEnd   = "block.end"

	StackType = "stack.type"

	Jump      = "jump"
	JumpTrue  = "jump.true"
	JumpFalse = "jump.false"

	Label = "label"
	Ret   = "ret"
)

func OPVarNew(varName *Symbol, varType *Symbol) Instruction {
	return Instruction{
		Operator: VarNew,
		Args:     []InstrArg{varName, varType},
	}
}

func OPVarSet(varName *Symbol, value InstrArg) Instruction {
	return Instruction{
		Operator: VarSet,
		Args:     []InstrArg{varName, value},
	}
}

func OPVarPop(varName *Symbol) Instruction {
	return Instruction{
		Operator: VarPop,
		Args:     []InstrArg{varName},
	}
}

func OPVarPush(varName *Symbol) Instruction {
	return Instruction{
		Operator: VarPush,
		Args:     []InstrArg{varName},
	}
}

func OPConstPush(value InstrArg) Instruction {
	return Instruction{
		Operator: ConstPush,
		Args:     []InstrArg{value},
	}
}

func OPCall(chunkName *Symbol, argSize *Integer) Instruction {
	return Instruction{
		Operator: Call,
		Args:     []InstrArg{chunkName, argSize},
	}
}

func OPCallb(symbol *Symbol, argSize *Integer) Instruction {
	return Instruction{
		Operator: Callb,
		Args:     []InstrArg{symbol, argSize},
	}

}

func OPBlockBegin(tp *Symbol) Instruction {
	return Instruction{
		Operator: BlockBegin,
		Args:     []InstrArg{tp},
	}
}

func OPBlockEnd() Instruction {
	return Instruction{
		Operator: BlockEnd,
	}
}

func OPJump(lbl *Symbol) Instruction {
	return Instruction{
		Operator: Jump,
		Args:     []InstrArg{lbl},
	}
}

func OPJumpTrue(lbl *Symbol) Instruction {
	return Instruction{
		Operator: JumpTrue,
		Args:     []InstrArg{lbl},
	}
}

func OPJumpFalse(lbl *Symbol) Instruction {
	return Instruction{
		Operator: JumpFalse,
		Args:     []InstrArg{lbl},
	}
}

func OPLabel(name *Symbol) Instruction {
	return Instruction{
		Operator: Label,
		Args:     []InstrArg{name},
	}
}

func OPStackType(tp *Symbol) Instruction {
	return Instruction{
		Operator: StackType,
		Args:     []InstrArg{tp},
	}
}

func OPReturn(cnt *Integer) Instruction {
	return Instruction{
		Operator: Ret,
		Args:     []InstrArg{cnt},
	}
}
