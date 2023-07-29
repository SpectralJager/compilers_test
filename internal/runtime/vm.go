package runtime

import (
	"gl/internal/ir"
	"log"
)

type VM struct {
	Program     ir.Program
	Constants   []Object
	GlobalFrame *Frame
}

func (vm *VM) MustExecute(program ir.Program) {
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
			}
		}
		log.Print("DONE\n")
	} else {
		log.Print("SKIPED\n")
	}
	log.Print("execute code...")
	if len(vm.Program.InitCode) > 0 {
		vm.GlobalFrame = NewFrame("global", vm.Program.Globals, vm.Program.InitCode)
		vm.run(vm.GlobalFrame)
		log.Print("DONE\n")
	} else {
		log.Print("SKIPED\n")
	}
	log.Println("execution finished")
}

func (vm *VM) run(f *Frame) Object {
	for ip := 0; ip < len(f.Code); {
		i := f.Code[ip]
		switch i := i.(type) {
		case *ir.Return:
			if i.Count == 0 {
				return nil
			}
			return f.Stack.Pop()
		case *ir.Load:
			ind := i.ConstIndex
			f.Stack.Push(vm.Constants[ind])
		case *ir.GlobalSet:
			ind := i.ConstIndex
			obj := vm.Constants[ind]
			symbol := i.Symbol
			vm.GlobalFrame.Varibles[symbol] = obj
		case *ir.GlobalLoad:
			symbol := i.Symbol
			obj := vm.GlobalFrame.Varibles[symbol]
			f.Stack.Push(obj)
		case *ir.GlobalSave:
			symbol := i.Symbol
			obj := f.Stack.Pop()
			vm.GlobalFrame.Varibles[symbol] = obj
		case *ir.LocalSet:
			ind := i.ConstIndex
			obj := vm.Constants[ind]
			symbol := i.Symbol
			f.Varibles[symbol] = obj
		case *ir.LocalLoad:
			symbol := i.Symbol
			obj := f.Varibles[symbol]
			f.Stack.Push(obj)
		case *ir.LocalSave:
			symbol := i.Symbol
			obj := f.Stack.Pop()
			f.Varibles[symbol] = obj
		case *ir.Jump:
			ip = i.Address
			continue
		case *ir.RelativeJump:
			ip += i.Count
			continue
		case *ir.ConditionalJump:
		case *ir.Call:
			symbol := i.FuncName
			fn := vm.Program.Functions[symbol]
			argCount := len(vm.Program.Globals[symbol].(*ir.FunctionDef).Arguments)
			objs := make([]Object, 0)
			for i := 0; i < argCount; i++ {
				objs = append(objs, f.Stack.Pop())
			}
			fr := NewFrame(f.Name, fn.Locals, fn.BodyCode)
			copy(fr.Stack, objs)
			res := vm.run(fr)
			if res != nil {
				f.Stack.Push(res)
			}
		case *ir.CallBuildin:
			symbol := i.FuncName
			switch symbol {
			case "iadd":
				obj2 := f.Stack.Pop()
				obj1 := f.Stack.Pop()
				res := iadd(obj1.(*Integer), obj2.(*Integer))
				f.Stack.Push(res)
			}
		}
		ip += 1
	}
	return nil
}
