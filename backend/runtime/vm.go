package runtime

import (
	"errors"
	"fmt"
	"grimlang/backend/asm"
	"strings"
)

var (
	ErrNoEntryPoint = errors.New("program haven't entry point")
)

type VM struct {
	Program *asm.Program
	Calls   *CallStack
	Stack   *Stack
	Cache   *Cache
}

func NewVM() VM {
	return VM{
		Stack: NewStack(),
		Calls: NewCallStack(),
		Cache: NewCache(),
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

func (vm *VM) PushFunc(fn asm.Function, argc int) {
	args := make([]string, argc)
	argv := make([]asm.Value, argc)
	for i := 0; i < argc; i++ {
		val := vm.Stack.Pop()
		args[i] = fmt.Sprintf("(%s)%s", val.Type.Inspect(), val.Inspect())
		argv[i] = val
	}
	hash := fmt.Sprintf("%s[%s]", fn.Ident, strings.Join(args, " "))
	if fn.Cacheble {
		if cacheItem, err := vm.Cache.Get(hash); err == nil {
			for i := 0; i < len(cacheItem.Return); i++ {
				vm.Stack.Push(cacheItem.Return[i])
			}
			return
		}
		for i := 0; i < argc; i++ {
			vm.Stack.Push(argv[i])
		}
	}
	fr := NewFrame(hash, fn, 0, 0, vm.Stack.Sp-argc)
	vm.Calls.Push(fr)
}

func (vm *VM) PopFunc(argc int) {
	argv := make([]asm.Value, argc)
	for i := 0; i < argc; i++ {
		val := vm.Stack.Pop()
		argv[i] = val
	}
	vm.Stack.Sp = vm.Calls.Top().Sp
	for i := 0; i < argc; i++ {
		vm.Stack.Push(argv[i])
	}
	fr := vm.Calls.Pop()
	if fr.Function.Cacheble {
		vm.Cache.Set(fr.Hash, CacheItem{
			Return: argv,
		})
	}
}

func (vm *VM) RunBlock() error {
	for {
		instr := vm.Calls.Top().NextInstruction()
		switch instr.Opcode {
		case asm.OP_Halt:
			return nil
		case asm.OP_Nop:
		case asm.OP_LocalLoad:
			val := instr.Args[0]
			vm.Stack.Push(
				vm.Calls.Top().Enviroment.Get(
					I64Value(val),
				),
			)
		case asm.OP_LocalSave:
			symb := instr.Args[0]
			val := vm.Stack.Pop()
			vm.Calls.Top().Enviroment.Set(
				I64Value(symb),
				val,
			)
		case asm.OP_Br:
			trgt := I64Value(instr.Args[0])
			vm.Calls.Top().SetBlock(int(trgt))
		case asm.OP_BrTrue:
			thn := I64Value(instr.Args[0])
			els := I64Value(instr.Args[1])
			val := vm.Stack.Pop()
			if BoolValue(val) {
				vm.Calls.Top().SetBlock(int(thn))
			} else {
				vm.Calls.Top().SetBlock(int(els))
			}
		case asm.OP_Call:
			ident := instr.Symbol
			fn, err := vm.Program.Function(ident)
			if err != nil {
				return err
			}
			argc := I64Value(instr.Args[0])
			vm.PushFunc(fn, int(argc))
		case asm.OP_Return:
			argc := I64Value(instr.Args[0])
			vm.PopFunc(int(argc))
		case asm.OP_I64Load:
			val := instr.Args[0]
			vm.Stack.Push(val)
		case asm.OP_I64Add:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueI64(
					I64Value(val1) + I64Value(val2),
				),
			)
		case asm.OP_I64Sub:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueI64(
					I64Value(val1) - I64Value(val2),
				),
			)
		case asm.OP_I64Mod:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueI64(
					I64Value(val1) % I64Value(val2),
				),
			)
		case asm.OP_I64Eq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					I64Value(val1) == I64Value(val2),
				),
			)
		case asm.OP_I64Neq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					I64Value(val1) != I64Value(val2),
				),
			)
		case asm.OP_I64Gt:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					I64Value(val1) > I64Value(val2),
				),
			)
		case asm.OP_I64Lt:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					I64Value(val1) < I64Value(val2),
				),
			)
		case asm.OP_BoolAnd:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					BoolValue(val1) && BoolValue(val2),
				),
			)

		default:
			return fmt.Errorf("unexpected instruction: %s", instr.Inspect())
		}
	}
}
