package vm

import (
	"fmt"
	"grimlang/internal/core/backend/bytecode"
)

type VM struct {
	Chunk *bytecode.Chunk
	Stack []bytecode.Value
}

func (vm *VM) InitVM() {
	vm.ClearVM()
}

func (vm *VM) ClearVM() {
	vm.Stack = make([]bytecode.Value, 0)
}

func (vm *VM) RunChunk(chunk *bytecode.Chunk) error {
	vm.Chunk = chunk
	for ip := 0; ip < len(vm.Chunk.Code); ip++ {
		switch vm.Chunk.Code[ip] {
		case bytecode.OP_RETURN:
			return nil
		case bytecode.OP_CONSTANT:
			ip += 1
			val_pt := vm.Chunk.Code[ip]
			val := vm.Chunk.Data[val_pt]
			vm.Push(val)
		case bytecode.OP_NEG:
			val := vm.Pop()
			vm.Push(-val)
		case bytecode.OP_ADD:
			a := vm.Pop()
			b := vm.Pop()
			vm.Push(a + b)
		case bytecode.OP_SUB:
			a := vm.Pop()
			b := vm.Pop()
			vm.Push(a - b)
		case bytecode.OP_MUL:
			a := vm.Pop()
			b := vm.Pop()
			vm.Push(a * b)
		case bytecode.OP_DIV:
			a := vm.Pop()
			b := vm.Pop()
			vm.Push(a / b)
		}
	}
	return nil
}

func (vm *VM) Push(val bytecode.Value) {
	vm.Stack = append(vm.Stack, val)
}

func (vm *VM) Pop() bytecode.Value {
	if len(vm.Stack) != 0 {
		val := vm.Stack[len(vm.Stack)-1]
		vm.Stack = vm.Stack[:len(vm.Stack)-1]
		return val
	}
	return 0
}

func (vm *VM) PrintTopStack() {
	fmt.Println(vm.Stack[len(vm.Stack)-1])
}
