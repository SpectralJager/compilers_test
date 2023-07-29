package runtime

import (
	"bytes"
	"fmt"
	"gl/internal/ir"
	"log"
)

type VM struct {
	Stack       Stack
	Program     ir.Program
	Constants   []Object
	currentCode []ir.IInstruction
}

func (vm *VM) MustExecute(program ir.Program) {
	vm.Program = program
	log.Println("Start executing")
	log.Print("read constants...")
	if len(program.Constants) > 0 {
		for _, c := range program.Constants {
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
	log.Print("execute initalization...")
	if len(program.InitCode) > 0 {
		vm.currentCode = program.InitCode
		vm.run()
		log.Print("DONE\n")
	} else {
		log.Print("SKIPED\n")
	}
	log.Println("execution finished")
}

func (vm *VM) StackTrace() string {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "=== Stack trace ===")
	for i, o := range vm.Stack {
		fmt.Fprintf(&buf, "%d: %s\n", i, o.String())
	}
	fmt.Fprintln(&buf, "=== =========== ===")
	return buf.String()
}

func (vm *VM) run() {
	for _, instr := range vm.currentCode {
		switch instr := instr.(type) {
		case *ir.Load:
			index := instr.ConstIndex
			vm.Stack.Push(vm.Constants[index])
		}
	}
}
