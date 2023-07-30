package runtime

import (
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
	if len(vm.Program.InitCode) > 0 {
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
	for vm.CurrentFrame().Ip < len(vm.CurrentFrame().Code) {
		f := vm.CurrentFrame()
		i := f.Instruction()
		switch i := i.(type) {
		case *ir.Halt:
			return
		case *ir.Return:
			if i.Count == 1 {
				retObj := vm.Stack.Pop()
				vm.Stack = vm.Stack[:f.Bp]
				vm.PopFrame()
				vm.Stack.Push(retObj)
			} else {
				vm.Stack = vm.Stack[:f.Bp]
				vm.PopFrame()
			}
			continue
		case *ir.Load:
			vm.Stack.Push(vm.Constants[i.ConstIndex])
		case *ir.GlobalSet:
			vm.GlobalFrame().Varibles[i.Symbol] = vm.Constants[i.ConstIndex]
		case *ir.GlobalLoad:
			vm.Stack.Push(vm.GlobalFrame().Varibles[i.Symbol])
		case *ir.GlobalSave:
			vm.GlobalFrame().Varibles[i.Symbol] = vm.Stack.Pop()
		case *ir.LocalSet:
			f.Varibles[i.Symbol] = vm.Constants[i.ConstIndex]
		case *ir.LocalLoad:
			vm.Stack.Push(f.Varibles[i.Symbol])
		case *ir.LocalSave:
			f.Varibles[i.Symbol] = vm.Stack.Pop()
		case *ir.Jump:
			vm.CurrentFrame().Ip = i.Address
			continue
		case *ir.RelativeJump:
			vm.CurrentFrame().Ip += i.Count
			continue
		case *ir.ConditionalJump:
			obj := vm.Stack.Pop().(*Boolean)
			if obj.Value {
				vm.CurrentFrame().Ip = i.Address
				continue
			}
		case *ir.Call:
			fn := vm.Program.Functions[i.FuncName]
			args := len(vm.Program.Globals[i.FuncName].(*ir.FunctionDef).Arguments)
			bp := vm.Stack.Len() - args
			// if bp < 0 {
			// 	panic("base pointer is out of bounds")
			// }
			fr := NewFrame(fn.Name, fn.Locals, fn.BodyCode, bp)
			vm.PushFrame(fr)
		case *ir.CallBuildin:
			symbol := i.FuncName
			switch symbol {
			case "iadd":
				obj1 := vm.Stack.Pop()
				obj2 := vm.Stack.Pop()
				res := iadd(obj1.(*Integer), obj2.(*Integer))
				vm.Stack.Push(res)
			case "isub":
				obj1 := vm.Stack.Pop()
				obj2 := vm.Stack.Pop()
				res := isub(obj1.(*Integer), obj2.(*Integer))
				vm.Stack.Push(res)
			case "ilt":
				obj1 := vm.Stack.Pop()
				obj2 := vm.Stack.Pop()
				res := ilt(obj1.(*Integer), obj2.(*Integer))
				vm.Stack.Push(res)
			}
		}
		f.Ip += 1
	}
}
