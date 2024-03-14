package runtime

import (
	"fmt"
	"grimlang/backend/asm"
	"testing"
)

var prog = tmp

var gcd = Must(asm.NewProgram(
	asm.NewFunction("main_main",
		false,
		asm.NewBlock(
			asm.InstructionI64Load(126),
			asm.InstructionLocalSave(0),
			asm.InstructionI64Load(120),
			asm.InstructionLocalSave(1),
			asm.InstructionBr(1),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad(0),
			asm.InstructionI64Load(0),
			asm.InstructionI64Neq(),
			asm.InstructionLocalLoad(1),
			asm.InstructionI64Load(0),
			asm.InstructionI64Neq(),
			asm.InstructionBoolAnd(),
			asm.InstructionBrTrue(2, 3),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad(0),
			asm.InstructionLocalLoad(1),
			asm.InstructionI64Gt(),
			asm.InstructionBrTrue(4, 5),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad(0),
			asm.InstructionLocalLoad(1),
			asm.InstructionI64Add(),
			asm.InstructionHalt(),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad(0),
			asm.InstructionLocalLoad(1),
			asm.InstructionI64Mod(),
			asm.InstructionLocalSave(0),
			asm.InstructionBr(1),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad(1),
			asm.InstructionLocalLoad(0),
			asm.InstructionI64Mod(),
			asm.InstructionLocalSave(1),
			asm.InstructionBr(1),
		),
	),
))

var fib_rec = Must(asm.NewProgram(
	asm.NewFunction("main_main",
		false,
		asm.NewBlock(
			asm.InstructionI64Load(35),
			asm.InstructionCall("main_fib", 1),
			asm.InstructionHalt(),
		),
	),
	asm.NewFunction("main_fib",
		true,
		asm.NewBlock(
			asm.InstructionLocalSave(0),
			asm.InstructionLocalLoad(0),
			asm.InstructionI64Load(2),
			asm.InstructionI64Lt(),
			asm.InstructionBrTrue(1, 2),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad(0),
			asm.InstructionReturn(1),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad(0),
			asm.InstructionI64Load(1),
			asm.InstructionI64Sub(),
			asm.InstructionCall("main_fib", 1),
			asm.InstructionLocalLoad(0),
			asm.InstructionI64Load(2),
			asm.InstructionI64Sub(),
			asm.InstructionCall("main_fib", 1),
			asm.InstructionI64Add(),
			asm.InstructionReturn(1),
		),
	),
))

var tmp = Must(asm.NewProgram(
	asm.NewFunction("main_main", false,
		asm.NewBlock(
			asm.InstructionF64Load(12.),
			asm.InstructionF64Load(30.),
			asm.InstructionF64Div(),
			asm.InstructionHalt(),
		),
	),
))

func TestRunBlock(t *testing.T) {
	fmt.Printf("\nProgram:\n%s\n", prog.InspectIndent(2))
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
