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
		case asm.OP_LocalLoad:
			val := instr.Args[0]
			Must(0, vm.Stack.Push(
				Must(vm.Functions.Vars.Get(
					Must(SymbolValue(val)),
				)),
			))
		case asm.OP_LocalSave:
			symb := instr.Args[0]
			val := Must(vm.Stack.Pop())
			Must(0, vm.Functions.Vars.Set(
				Must(SymbolValue(symb)),
				val,
			))
		case asm.OP_Br:
			trgt := Must(I64Value(instr.Args[0]))
			vm.Functions.SetBlock(int(trgt))
		case asm.OP_BrTrue:
			thn := Must(I64Value(instr.Args[0]))
			els := Must(I64Value(instr.Args[1]))
			val := Must(vm.Stack.Pop())
			if Must(BoolValue(val)) {
				vm.Functions.SetBlock(int(thn))
			} else {
				vm.Functions.SetBlock(int(els))
			}
		case asm.OP_I64Load:
			val := instr.Args[0]
			Must(0, vm.Stack.Push(val))
		case asm.OP_I64Add:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueI64(
					Must(I64Value(val1))+Must(I64Value(val2)),
				),
			))
		case asm.OP_I64Sub:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueI64(
					Must(I64Value(val1))-Must(I64Value(val2)),
				),
			))
		case asm.OP_I64Mod:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueI64(
					Must(I64Value(val1))%Must(I64Value(val2)),
				),
			))
		case asm.OP_I64Eq:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueBool(
					Must(I64Value(val1)) == Must(I64Value(val2)),
				),
			))
		case asm.OP_I64Neq:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueBool(
					Must(I64Value(val1)) != Must(I64Value(val2)),
				),
			))
		case asm.OP_I64Gt:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueBool(
					Must(I64Value(val1)) > Must(I64Value(val2)),
				),
			))
		case asm.OP_I64Lt:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueBool(
					Must(I64Value(val1)) < Must(I64Value(val2)),
				),
			))
		case asm.OP_BoolAnd:
			val2 := Must(vm.Stack.Pop())
			val1 := Must(vm.Stack.Pop())
			Must(0, vm.Stack.Push(
				asm.ValueBool(
					Must(BoolValue(val1)) && Must(BoolValue(val2)),
				),
			))

		default:
			return fmt.Errorf("unexpected instruction: %s", instr.Inspect())
		}
	}
}
