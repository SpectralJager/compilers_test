package ir

import (
	"fmt"
	"strings"
)

type Instruction struct {
	Operand OperandType
	Args    []InstrArg
}

func NewInstruction(op OperandType, args ...InstrArg) *Instruction {
	return &Instruction{Operand: op, Args: args}
}

func (i Instruction) String() string {
	var buf strings.Builder
	if s, ok := OperandString[i.Operand]; ok {
		fmt.Fprintf(&buf, "%s", s)
	} else {
		fmt.Fprint(&buf, "ILLEGAL")
	}
	for _, arg := range i.Args {
		fmt.Fprintf(&buf, " %s", arg)
	}
	return buf.String()
}

type OperandType byte

const (
	OP_NULL OperandType = iota
	OP_VAR_NEW
	OP_VAR_SET
	OP_VAR_POP
	OP_VAR_PUSH
	OP_VAR_FREE

	OP_CONST_PUSH

	OP_CALL
	OP_CALLB

	OP_RET

	OP_STACK_DUP
	OP_STACK_TYPE
	OP_STACK_POP

	OP_BR
	OP_BR_TRUE
	OP_BR_FALSE
)

var OperandString = map[OperandType]string{
	OP_NULL: "NULL",

	OP_VAR_NEW:  "VAR.NEW",
	OP_VAR_SET:  "VAR.SET",
	OP_VAR_POP:  "VAR.POP",
	OP_VAR_PUSH: "VAR.PUSH",
	OP_VAR_FREE: "VAR.FREE",

	OP_CONST_PUSH: "CONST.PUSH",

	OP_CALL:  "CALL",
	OP_CALLB: "CALLB",

	OP_RET: "RET",

	OP_STACK_DUP:  "STACK.DUP",
	OP_STACK_POP:  "STACK.POP",
	OP_STACK_TYPE: "STACK.TYPE",

	OP_BR:       "BR",
	OP_BR_TRUE:  "BR.TRUE",
	OP_BR_FALSE: "BR.FALSE",
}

func VarNew(symbol *Symbol, tp *Symbol) *Instruction {
	return NewInstruction(OP_VAR_NEW, symbol, tp)
}
func VarSet(symbol *Symbol, val InstrArg) *Instruction {
	return NewInstruction(OP_VAR_SET, symbol, val)
}
func VarPop(symbol *Symbol) *Instruction {
	return NewInstruction(OP_VAR_POP, symbol)
}
func VarPush(symbol *Symbol) *Instruction {
	return NewInstruction(OP_VAR_PUSH, symbol)
}
func VarFree(symbol *Symbol) *Instruction {
	return NewInstruction(OP_VAR_FREE, symbol)
}
func Call(symbol *Symbol, argsCount *Integer) *Instruction {
	return NewInstruction(OP_CALL, symbol, argsCount)
}
func CallBuiltin(symbol *Symbol, argsCount *Integer) *Instruction {
	return NewInstruction(OP_CALLB, symbol, argsCount)
}
func Return(count *Integer) *Instruction {
	return NewInstruction(OP_RET, count)
}
func ConstPush(constant InstrArg) *Instruction {
	return NewInstruction(OP_CONST_PUSH, constant)
}
func StackPop() *Instruction {
	return NewInstruction(OP_STACK_POP)
}
func StackDup() *Instruction {
	return NewInstruction(OP_STACK_DUP)
}
func StackType(tp *Symbol) *Instruction {
	return NewInstruction(OP_STACK_TYPE, tp)
}
func Br(label *Symbol) *Instruction {
	return NewInstruction(OP_BR, label)
}
func BrTrue(labelThen *Symbol, labelElse *Symbol) *Instruction {
	return NewInstruction(OP_BR_TRUE, labelThen, labelElse)
}
func BrFalse(labelThen *Symbol, labelElse *Symbol) *Instruction {
	return NewInstruction(OP_BR_FALSE, labelThen, labelElse)
}
