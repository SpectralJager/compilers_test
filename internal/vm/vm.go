package vm

import (
	"bytes"
	"fmt"
	"gl/internal/repr"
)

type VM struct {
	Stack []repr.Object
	Chunk *repr.Chunk
}

func NewVM() *VM {
	return &VM{}
}

func (vm *VM) ExecuteChunk(chunk *repr.Chunk) error {
	vm.Chunk = chunk
	return vm.Eval()
}

func (vm *VM) Eval() error {
	for offset := 0; offset < len(vm.Chunk.Code); {
		instr := vm.GetCode(offset)
		switch instr {
		case repr.OP_NULL:
			offset += 1
		case repr.OP_HALT:
			return nil
		case repr.OP_CONST:
			offset += 1
			ind := vm.GetCode(offset)
			vm.PushObject(vm.GetConstant(int(ind)))
		case repr.OP_INEG:
			offset += 1
			obj1 := vm.PopObject()
			newObj := repr.Object{Value: Neg(obj1.Value.(int)), Kind: repr.ObjectInt}
			vm.PushObject(newObj)
		case repr.OP_IINC:
			offset += 1
			obj1 := vm.PopObject()
			newObj := repr.Object{Value: Add(obj1.Value.(int), 1), Kind: repr.ObjectInt}
			vm.PushObject(newObj)
		case repr.OP_IDEC:
			offset += 1
			obj1 := vm.PopObject()
			newObj := repr.Object{Value: Sub(obj1.Value.(int), 1), Kind: repr.ObjectInt}
			vm.PushObject(newObj)
		case repr.OP_IADD:
			offset += 1
			obj2 := vm.PopObject()
			obj1 := vm.PopObject()
			newObj := repr.Object{Value: Add(obj1.Value.(int), obj2.Value.(int)), Kind: repr.ObjectInt}
			vm.PushObject(newObj)
		case repr.OP_ISUB:
			offset += 1
			obj2 := vm.PopObject()
			obj1 := vm.PopObject()
			newObj := repr.Object{Value: Sub(obj1.Value.(int), obj2.Value.(int)), Kind: repr.ObjectInt}
			vm.PushObject(newObj)
		case repr.OP_IMUL:
			offset += 1
			obj2 := vm.PopObject()
			obj1 := vm.PopObject()
			newObj := repr.Object{Value: Mul(obj1.Value.(int), obj2.Value.(int)), Kind: repr.ObjectInt}
			vm.PushObject(newObj)
		case repr.OP_IDIV:
			offset += 1
			obj2 := vm.PopObject()
			obj1 := vm.PopObject()
			newObj := repr.Object{Value: Div(obj1.Value.(int), obj2.Value.(int)), Kind: repr.ObjectInt}
			vm.PushObject(newObj)
		default:
			return fmt.Errorf("unexpected opcode %04xh", instr)
		}
	}
	return fmt.Errorf("program finished without halt instruction")
}

func (vm *VM) StackTrace() string {
	var buf bytes.Buffer
	fmt.Fprint(&buf, "=== Stack ===\n")
	for i, obj := range vm.Stack {
		fmt.Fprintf(&buf, "%d: %s\n", i, obj.String())
	}
	return buf.String()
}

func (vm *VM) GetCode(offset int) byte {
	return vm.Chunk.Code[offset]
}

func (vm *VM) GetConstant(index int) repr.Object {
	return vm.Chunk.ConstPool[index]
}

func (vm *VM) PushObject(obj repr.Object) {
	vm.Stack = append(vm.Stack, obj)
}

func (vm *VM) PopObject() repr.Object {
	if len(vm.Stack) == 0 {
		panic("empty stack")
	}
	obj := vm.Stack[len(vm.Stack)-1]
	vm.Stack = vm.Stack[:len(vm.Stack)-1]
	return obj
}
