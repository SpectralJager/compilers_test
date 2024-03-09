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
	Program *asm.Program
	Calls   *CallStack
	Stack   *Stack
}

func NewVM() VM {
	return VM{
		Stack: NewStack(),
		Calls: NewCallStack(),
	}
}

func (vm *VM) LoadProgram(prog *asm.Program) error {
	main, err := prog.Function("main_main")
	if err != nil {
		return ErrNoEntryPoint
	}
	vm.Program = prog
	vm.PushFunc(main, 0)
	return nil
}

func (vm *VM) PushFunc(fn asm.Function, argc int) error {
	// argv := []asm.Value{}
	// for i := 0; i < argc; i++ {
	// 	argv = append(argv, Must(vm.Stack.Pop()))
	// }
	fr := NewFrame(fn, 0, 0, vm.Stack.Sp-argc)
	return vm.Calls.Push(fr)
}

func (vm *VM) PopFunc(argc int) error {
	argv := []asm.Value{}
	for i := 0; i < argc; i++ {
		val, err := vm.Stack.Pop()
		if err != nil {
			return err
		}
		argv = append(argv, val)
	}
	vm.Stack.Sp = vm.Calls.Top().Sp
	for _, arg := range argv {
		err := vm.Stack.Push(arg)
		if err != nil {
			return err
		}
	}
	return vm.Calls.Pop()
}

func (vm *VM) RunBlock() error {
	for {
		instr := Must(vm.Calls.Top().NextInstruction())
		switch instr.Opcode {
		case asm.OP_Halt:
			return nil
		case asm.OP_Nop:
		case asm.OP_LocalLoad:
			val := instr.Args[0]
			Must(0, vm.Stack.Push(
				Must(vm.Calls.Top().Enviroment.Get(
					Must(SymbolValue(val)),
				)),
			))
		case asm.OP_LocalSave:
			symb := instr.Args[0]
			val := Must(vm.Stack.Pop())
			Must(0, vm.Calls.Top().Enviroment.Set(
				Must(SymbolValue(symb)),
				val,
			))
		case asm.OP_Br:
			trgt := Must(I64Value(instr.Args[0]))
			vm.Calls.Top().SetBlock(int(trgt))
		case asm.OP_BrTrue:
			thn := Must(I64Value(instr.Args[0]))
			els := Must(I64Value(instr.Args[1]))
			val := Must(vm.Stack.Pop())
			if Must(BoolValue(val)) {
				vm.Calls.Top().SetBlock(int(thn))
			} else {
				vm.Calls.Top().SetBlock(int(els))
			}
		case asm.OP_Call:
			ident := Must(SymbolValue(instr.Args[0]))
			fn, err := vm.Program.Function(ident)
			if err != nil {
				return err
			}
			argc := Must(I64Value(instr.Args[1]))
			Must(0, vm.PushFunc(fn, int(argc)))
		case asm.OP_Return:
			argc := Must(I64Value(instr.Args[0]))
			Must(0, vm.PopFunc(int(argc)))
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
