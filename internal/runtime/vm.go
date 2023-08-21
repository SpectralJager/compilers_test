package runtime

import (
	"encoding/binary"
	"fmt"
	"gl/internal/ir"
	"log"
)

type VM struct {
	Program    *ir.Program
	Constants  []Object
	Stack      Stack
	Frames     [512]Frame
	frameIndex int
}

func (vm *VM) MustExecute(program *ir.Program) {
	vm.Stack = make(Stack, 0, 2048)
	vm.Program = program
	log.Println("Start executing")
	log.Print("read constants...")
	if len(vm.Program.Constants) > 0 {
		for _, c := range vm.Program.Constants {
			switch c := c.(type) {
			case *ir.Integer:
				vm.Constants = append(vm.Constants, &Integer{c.Value})
			case *ir.Float:
				vm.Constants = append(vm.Constants, &Float{c.Value})
			case *ir.String:
				vm.Constants = append(vm.Constants, &String{c.Value})
			case *ir.True:
				vm.Constants = append(vm.Constants, &Boolean{true})
			case *ir.False:
				vm.Constants = append(vm.Constants, &Boolean{false})
			}
		}
		log.Print("DONE\n")
	} else {
		log.Print("SKIPED\n")
	}
	log.Print("execute code...")
	if vm.Program.InitCode.Len() > 0 {
		fr := NewFrame("global", vm.Program.Globals, vm.Program.InitCode, 0)
		vm.PushFrame(fr)
		vm.run()
		log.Print("DONE\n")
	} else {
		log.Print("SKIPED\n")
	}
	log.Println("execution finished")
}

func (vm *VM) GlobalFrame() *Frame {
	return &vm.Frames[0]
}

func (vm *VM) CurrentFrame() *Frame {
	return &vm.Frames[vm.frameIndex-1]
}

func (vm *VM) PushFrame(frame Frame) {
	vm.Frames[vm.frameIndex] = frame
	vm.frameIndex += 1
}

func (vm *VM) PopFrame() {
	if len(vm.Frames) == 1 {
		return
	}
	vm.frameIndex -= 1
}

func (vm *VM) run() {
	for vm.CurrentFrame().Ip < vm.CurrentFrame().Code.Len() {
		f := vm.CurrentFrame()
		i := f.Instruction()
		switch i {
		case ir.OP_NULL:
			f.Ip += 1
		case ir.OP_HALT:
			return
		case ir.OP_RETURN:
			f.Ip += 1
			count := vm.CurrentFrame().Code.ReadBytes(f.Ip, 1)[0]
			f.Ip += 1
			if count == 0 {
				vm.Stack = vm.Stack[:f.Bp]
				vm.PopFrame()
			} else {
				obj := vm.Stack.Pop()
				vm.Stack = vm.Stack[:f.Bp]
				vm.PopFrame()
				vm.Stack.Push(obj)
			}
		case ir.OP_LOAD:
			f.Ip += 1
			ind := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			obj := vm.Constants[ind]
			vm.Stack.Push(obj)
		case ir.OP_GLOBAL_SET:
			f.Ip += 1
			indV := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			indC := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			obj := vm.Constants[indC]
			vm.GlobalFrame().Varibles[indV] = obj
		case ir.OP_GLOBAL_LOAD:
			f.Ip += 1
			ind := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			vm.Stack.Push(vm.GlobalFrame().Varibles[ind])
		case ir.OP_GLOBAL_SAVE:
			f.Ip += 1
			ind := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			obj := vm.Stack.Pop()
			vm.GlobalFrame().Varibles[ind] = obj
		case ir.OP_LOCAL_SET:
			f.Ip += 1
			indV := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			indC := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			obj := vm.Constants[indC]
			f.Varibles[indV] = obj
		case ir.OP_LOCAL_LOAD:
			f.Ip += 1
			ind := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			vm.Stack.Push(f.Varibles[ind])
		case ir.OP_LOCAL_SAVE:
			f.Ip += 1
			ind := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			obj := vm.Stack.Pop()
			f.Varibles[ind] = obj
		case ir.OP_JUMP:
			f.Ip += 1
			address := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip = int(address)
		case ir.OP_JUMP_CONDITION:
			f.Ip += 1
			address := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			if vm.Stack.Pop().(*Boolean).Value {
				f.Ip = int(address)
				continue
			}
			f.Ip += 4
		case ir.OP_JUMP_NOT_CONDITION:
			f.Ip += 1
			address := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			if !vm.Stack.Pop().(*Boolean).Value {
				f.Ip = int(address)
				continue
			}
			f.Ip += 4
		case ir.OP_CALL:
			f.Ip += 1
			ind := binary.LittleEndian.Uint32(f.Code.ReadBytes(f.Ip, 4))
			f.Ip += 4
			def := vm.Program.Globals[ind].(*ir.FunctionDef)
			fn := vm.Program.Functions[def.Name]
			fr := NewFrame(def.Name, fn.Locals, fn.BodyCode, vm.Stack.Len()-len(def.Arguments))
			vm.PushFrame(fr)
		case ir.OP_INT_FUNC:
			f.Ip += 1
			id := f.Code.ReadBytes(f.Ip, 1)[0]
			f.Ip += 1
			switch id {
			case 0:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := iadd(obj1, obj2)
				vm.Stack.Push(res)
			case 1:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := isub(obj1, obj2)
				vm.Stack.Push(res)
			case 2:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := imul(obj1, obj2)
				vm.Stack.Push(res)
			case 3:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := idiv(obj1, obj2)
				vm.Stack.Push(res)
			case 4:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := ilt(obj1, obj2)
				vm.Stack.Push(res)
			case 5:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := igt(obj1, obj2)
				vm.Stack.Push(res)
			case 6:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := ileq(obj1, obj2)
				vm.Stack.Push(res)
			case 7:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := igeq(obj1, obj2)
				vm.Stack.Push(res)
			case 8:
				obj1 := vm.Stack.Pop().(*Integer)
				obj2 := vm.Stack.Pop().(*Integer)
				res := ieq(obj1, obj2)
				vm.Stack.Push(res)
			case 9:
				obj1 := vm.Stack.Pop().(*Integer)
				res := itof(obj1)
				vm.Stack.Push(res)
			case 10:
				obj1 := vm.Stack.Pop().(*Integer)
				res := itos(obj1)
				vm.Stack.Push(res)
			case 11:
				obj1 := vm.Stack.Pop().(*Integer)
				res := itob(obj1)
				vm.Stack.Push(res)
			}
		case ir.OP_FLOAT_FUNC:
			f.Ip += 1
			id := f.Code.ReadBytes(f.Ip, 1)[0]
			f.Ip += 1
			switch id {
			case 0:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := fadd(obj1, obj2)
				vm.Stack.Push(res)
			case 1:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := fsub(obj1, obj2)
				vm.Stack.Push(res)
			case 2:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := fmul(obj1, obj2)
				vm.Stack.Push(res)
			case 3:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := fdiv(obj1, obj2)
				vm.Stack.Push(res)
			case 4:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := flt(obj1, obj2)
				vm.Stack.Push(res)
			case 5:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := fgt(obj1, obj2)
				vm.Stack.Push(res)
			case 6:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := fleq(obj1, obj2)
				vm.Stack.Push(res)
			case 7:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := fgeq(obj1, obj2)
				vm.Stack.Push(res)
			case 8:
				obj1 := vm.Stack.Pop().(*Float)
				obj2 := vm.Stack.Pop().(*Float)
				res := feq(obj1, obj2)
				vm.Stack.Push(res)
			case 9:
				obj1 := vm.Stack.Pop().(*Float)
				res := ftoi(obj1)
				vm.Stack.Push(res)
			case 10:
				obj1 := vm.Stack.Pop().(*Float)
				res := ftos(obj1)
				vm.Stack.Push(res)
			case 11:
				obj1 := vm.Stack.Pop().(*Float)
				res := ftob(obj1)
				vm.Stack.Push(res)
			}
		case ir.OP_STRING_FUNC:
			f.Ip += 1
			id := f.Code.ReadBytes(f.Ip, 1)[0]
			f.Ip += 1
			switch id {
			case 0:
				obj1 := vm.Stack.Pop().(*String)
				obj2 := vm.Stack.Pop().(*String)
				res := sconcat(obj1, obj2)
				vm.Stack.Push(res)
			case 1:
				obj1 := vm.Stack.Pop().(*String)
				res := slen(obj1)
				vm.Stack.Push(res)
			case 2:
				obj1 := vm.Stack.Pop().(*String)
				res := stoi(obj1)
				vm.Stack.Push(res)
			case 3:
				obj1 := vm.Stack.Pop().(*String)
				res := stof(obj1)
				vm.Stack.Push(res)
			case 4:
				obj1 := vm.Stack.Pop().(*String)
				res := stob(obj1)
				vm.Stack.Push(res)
			}
		default:
			panic(fmt.Sprintf("invalid instruction %d", i))
		}
	}
}
