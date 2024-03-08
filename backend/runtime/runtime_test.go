package runtime

import (
	"fmt"
	"grimlang/backend/asm"
	"testing"
)

var prog = Must(asm.NewProgram(
	asm.NewFunction("main_main",
		asm.NewBlock(
			asm.InstructionI64Load(12),
			asm.InstructionI64Load(18),
			asm.InstructionI64Add(),
			asm.InstructionI64Load(18),
			asm.InstructionHalt(),
		),
	),
))

func TestRunBlock(t *testing.T) {
	vm := NewVM()
	err := vm.LoadProgram(prog)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.RunBlock()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\n%s\n", vm.Stack.TraceMemory())
}
