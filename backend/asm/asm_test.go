package asm

import (
	"fmt"
	"testing"
)

func TestBlock(t *testing.T) {
	block := NewBlock(
		InstructionI64Load(12),
		InstructionI64Load(18),
		InstructionI64Add(),
		InstructionHalt(),
	)
	fmt.Printf("\n%s\n", block.Inspect(0))
}

func TestFunction(t *testing.T) {
	fn := NewFunction("main_main", nil)
	fn.Blocks = append(fn.Blocks, NewBlock(
		InstructionI64Load(12),
		InstructionI64Load(18),
		InstructionI64Add(),
		InstructionHalt(),
	))
	fmt.Printf("\n%s\n", fn.InspectIndent(2))
}

func TestProgram(t *testing.T) {
	prog, err := NewProgram(
		NewFunction("main", nil,
			NewBlock(
				InstructionI64Load(12),
				InstructionI64Load(18),
				InstructionI64Add(),
				InstructionHalt(),
			),
		),
		NewFunction("fib", nil,
			NewBlock(
				InstructionI64Load(12),
				InstructionI64Load(18),
				InstructionI64Add(),
				InstructionHalt(),
			),
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\n%s\n", prog.InspectIndent(2))
}
