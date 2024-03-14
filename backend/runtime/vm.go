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
	for i := 0; i < argc; i++ {
		val := vm.Stack.Memory[vm.Stack.Sp-argc+i]
		args[i] = fmt.Sprintf("(%s)%s", val.Type.Inspect(), val.Inspect())
	}
	hash := fmt.Sprintf("%s[%s]", fn.Ident, strings.Join(args, " "))
	if fn.Cacheble {
		if cacheItem, err := vm.Cache.Get(hash); err == nil {
			vm.Stack.Sp -= argc
			for i := 0; i < len(cacheItem.Return); i++ {
				vm.Stack.Push(cacheItem.Return[i])
			}
			return
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
		case asm.OP_Rotate:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(val2)
			vm.Stack.Push(val1)
		case asm.OP_Duplicate:
			val := vm.Stack.Pop()
			vm.Stack.Push(val)
			vm.Stack.Push(val)
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
		case asm.OP_I64Mul:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueI64(
					I64Value(val1) * I64Value(val2),
				),
			)
		case asm.OP_I64Div:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueI64(
					I64Value(val1) / I64Value(val2),
				),
			)
		case asm.OP_I64Neg:
			val := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueI64(
					-I64Value(val),
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
		case asm.OP_I64Geq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					I64Value(val1) >= I64Value(val2),
				),
			)
		case asm.OP_I64Leq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					I64Value(val1) <= I64Value(val2),
				),
			)
		case asm.OP_F64Load:
			val := instr.Args[0]
			vm.Stack.Push(val)
		case asm.OP_F64Add:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueF64(
					F64Value(val1) + F64Value(val2),
				),
			)
		case asm.OP_F64Sub:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueF64(
					F64Value(val1) - F64Value(val2),
				),
			)
		case asm.OP_F64Mul:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueF64(
					F64Value(val1) * F64Value(val2),
				),
			)
		case asm.OP_F64Div:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueF64(
					F64Value(val1) / F64Value(val2),
				),
			)
		case asm.OP_F64Neg:
			val := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueF64(
					-F64Value(val),
				),
			)
		case asm.OP_F64Eq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					F64Value(val1) == F64Value(val2),
				),
			)
		case asm.OP_F64Neq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					F64Value(val1) != F64Value(val2),
				),
			)
		case asm.OP_F64Gt:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					F64Value(val1) > F64Value(val2),
				),
			)
		case asm.OP_F64Lt:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					F64Value(val1) < F64Value(val2),
				),
			)
		case asm.OP_F64Geq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					F64Value(val1) >= F64Value(val2),
				),
			)
		case asm.OP_F64Leq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					F64Value(val1) <= F64Value(val2),
				),
			)
		case asm.OP_BoolLoad:
			val := instr.Args[0]
			vm.Stack.Push(
				val,
			)
		case asm.OP_BoolEq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					BoolValue(val1) == BoolValue(val2),
				),
			)
		case asm.OP_BoolNeq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					BoolValue(val1) != BoolValue(val2),
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
		case asm.OP_BoolOr:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					BoolValue(val1) || BoolValue(val2),
				),
			)
		case asm.OP_BoolNot:
			val := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					!BoolValue(val),
				),
			)
		case asm.OP_StringLoad:
			val := instr.Args[0]
			vm.Stack.Push(
				val,
			)
		case asm.OP_StringConcat:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueString(
					StringValue(val1) + StringValue(val2),
				),
			)
		case asm.OP_StringEq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					StringValue(val1) == StringValue(val2),
				),
			)
		case asm.OP_StringNeq:
			val2 := vm.Stack.Pop()
			val1 := vm.Stack.Pop()
			vm.Stack.Push(
				asm.ValueBool(
					StringValue(val1) != StringValue(val2),
				),
			)
		case asm.OP_ListLoad:
			val := instr.Args[0]
			vm.Stack.Push(
				val,
			)
		case asm.OP_ListConstruct:
			argc := I64Value(instr.Args[0])
			items := make([]asm.Value, argc)
			for i := argc - 1; i >= 0; i-- {
				items[i] = vm.Stack.Pop()
			}
			vm.Stack.Push(
				asm.ValueList(items...),
			)
		case asm.OP_ListGet:
			index := I64Value(vm.Stack.Pop())
			list := ListValue(vm.Stack.Pop())
			vm.Stack.Push(
				list[index],
			)
		case asm.OP_ListSet:
			val := vm.Stack.Pop()
			index := I64Value(vm.Stack.Pop())
			list := ListValue(vm.Stack.Pop())
			list[index] = val
			vm.Stack.Push(
				asm.ValueList(list...),
			)
		case asm.OP_ListInsert:
			val := vm.Stack.Pop()
			index := I64Value(vm.Stack.Pop())
			list := ListValue(vm.Stack.Pop())
			list = append(list[:index], append([]asm.Value{val}, list[index:]...)...)
			vm.Stack.Push(
				asm.ValueList(list...),
			)
		default:
			return fmt.Errorf("unexpected instruction: %s", instr.Inspect())
		}
	}
}
