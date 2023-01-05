package vm

import (
	"fmt"
	"grimlang/internal/core/backend/bytecode"
	"grimlang/internal/core/backend/chunk"
	"grimlang/internal/core/backend/object"
)

type VM struct {
	Stack        []object.Object
	MainChunk    chunk.Chunk
	Chunks       []chunk.Chunk
	IP           uint16
	SP           uint16
	ExecuteChunk chunk.Chunk
}

func NewVM(main chunk.Chunk) *VM {
	return &VM{MainChunk: main}
}

func (v *VM) WriteChunk(ch chunk.Chunk) {
	v.Chunks = append(v.Chunks, ch)
}

func (v *VM) Run() {
	v.MainChunk.StartAddress = v.IP
	v.ExecuteChunk = v.MainChunk
	retIP := v.IP
	v.runChunk()
	v.IP = retIP
}

func (v *VM) runChunk() {
	for v.IP = 0; v.IP < uint16(len(v.ExecuteChunk.Code)); v.IP++ {
		op := v.ExecuteChunk.Code[v.IP]
		switch op {
		case bytecode.OP_LOAD_CONST:
			v.IP += 1
			const_obj := v.ExecuteChunk.Data[v.IP]
			v.PushStack(const_obj)
		case bytecode.OP_RET:
			return
		default:
			panic("unexpected bytecode: " + fmt.Sprint(op))
		}
	}
}

func (v *VM) PushStack(obj object.Object) {
	v.Stack = append(v.Stack, obj)
	v.SP += 1
}
func (v *VM) PopStack() object.Object {
	if v.SP == 0 {
		panic("cant pop value from empty stack")
	}
	obj := v.Stack[v.SP]
	v.Stack[v.SP] = object.Object{}
	v.SP -= 1
	return obj
}

func (v *VM) TraceStack() {

}
