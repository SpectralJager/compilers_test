package runtime

import (
	"fmt"
	"grimlang/backend/asm"
	"testing"
)

var prog = fib

var gcd = Must(asm.NewProgram(
	asm.NewFunction("main_main",
		map[string]asm.Value{
			"a": asm.ValueI64(0),
			"b": asm.ValueI64(0),
		},
		asm.NewBlock(
			asm.InstructionI64Load(126),
			asm.InstructionLocalSave("a"),
			asm.InstructionI64Load(120),
			asm.InstructionLocalSave("b"),
			asm.InstructionBr(1),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad("a"),
			asm.InstructionI64Load(0),
			asm.InstructionI64Neq(),
			asm.InstructionLocalLoad("b"),
			asm.InstructionI64Load(0),
			asm.InstructionI64Neq(),
			asm.InstructionBoolAnd(),
			asm.InstructionBrTrue(2, 3),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad("a"),
			asm.InstructionLocalLoad("b"),
			asm.InstructionI64Gt(),
			asm.InstructionBrTrue(4, 5),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad("a"),
			asm.InstructionLocalLoad("b"),
			asm.InstructionI64Add(),
			asm.InstructionHalt(),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad("a"),
			asm.InstructionLocalLoad("b"),
			asm.InstructionI64Mod(),
			asm.InstructionLocalSave("a"),
			asm.InstructionBr(1),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad("b"),
			asm.InstructionLocalLoad("a"),
			asm.InstructionI64Mod(),
			asm.InstructionLocalSave("b"),
			asm.InstructionBr(1),
		),
	),
))

var fib = Must(asm.NewProgram(
	asm.NewFunction("main_main",
		asm.Vars{
			"n":       asm.ValueI64(50),
			"i":       asm.ValueI64(0),
			"fib1":    asm.ValueI64(1),
			"fib2":    asm.ValueI64(1),
			"fib_sum": asm.ValueI64(0),
		},
		asm.NewBlock(
			asm.InstructionLocalLoad("i"),
			asm.InstructionLocalLoad("n"),
			asm.InstructionI64Load(2),
			asm.InstructionI64Sub(),
			asm.InstructionI64Lt(),
			asm.InstructionBrTrue(1, 2),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad("fib1"),
			asm.InstructionLocalLoad("fib2"),
			asm.InstructionI64Add(),
			asm.InstructionLocalSave("fib_sum"),
			asm.InstructionLocalLoad("fib2"),
			asm.InstructionLocalSave("fib1"),
			asm.InstructionLocalLoad("fib_sum"),
			asm.InstructionLocalSave("fib2"),
			asm.InstructionLocalLoad("i"),
			asm.InstructionI64Load(1),
			asm.InstructionI64Add(),
			asm.InstructionLocalSave("i"),
			asm.InstructionBr(0),
		),
		asm.NewBlock(
			asm.InstructionLocalLoad("fib2"),
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