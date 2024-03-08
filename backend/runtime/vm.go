package runtime

import (
	"errors"
	"fmt"
	"grimlang/backend/asm"
)

var (
	ErrNoEntryPoint = errors.New("program haven't entry point")
)

type VM struct {
	Program   *asm.Program
	Functions *asm.Function
	Stack     *Stack
}

func NewVM() VM {
	return VM{
		Stack: NewStack(),
	}
}

func (vm *VM) LoadProgram(prog *asm.Program) error {
	main := prog.Functions["main_main"]
	if main == nil {
		return ErrNoEntryPoint
	}
	vm.Program = prog
	vm.Functions = main
	return nil
}

func (vm *VM) RunBlock() error {
	for {
		block := vm.Functions.Block()
		instr := Must(block.Next())
		switch instr.Opcode {
		case asm.OP_Halt:
			return nil
		case asm.OP_Nop:
		case asm.OP_I64Load:
			val := instr.Args[0]
			err := vm.Stack.Push(val)
			if err != nil {
				return err
			}
		case asm.OP_I64Add:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueI64(
					Must(I64Value(val1))+Must(I64Value(val2)),
				),
			))
		default:
			return fmt.Errorf("unexpected instruction: %s", instr.Inspect())
		}
	}
}
