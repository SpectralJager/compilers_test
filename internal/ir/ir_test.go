package ir

import (
	"fmt"
	"testing"
)

func TestIr(t *testing.T) {
	module := NewModule("test", ".")
	module.AddChunk(
		NewChunk("GLOBAL").
			AddBlock(
				NewBlock("start").
					AddInstructions(
						VarNew(NewSymbol("n", nil), NewSymbol("int", nil)),
						VarSet(NewSymbol("n", nil), NewInteger(40)),
						Call(NewSymbol("main", nil), NewInteger(0)),
					)))
	module.AddChunk(
		NewChunk("main").
			AddBlock(
				NewBlock("start").
					AddInstructions(
						VarNew(NewSymbol("res", nil), NewSymbol("int", nil)),
						VarPush(NewSymbol("n", nil)),
						Call(NewSymbol("fib", nil), NewInteger(1)),
						VarPop(NewSymbol("res", nil)),
						VarPush(NewSymbol("res", nil)),
						CallBuiltin(NewSymbol("io", NewSymbol("println", nil)), NewInteger(1)),
					),
			),
	)
	module.AddChunk(
		NewChunk("fib").
			AddBlocks(
				NewBlock("start").
					AddInstructions(
						VarNew(NewSymbol("n", nil), NewSymbol("int", nil)),
						VarPop(NewSymbol("n", nil)),
						VarPush(NewSymbol("n", nil)),
						ConstPush(NewInteger(2)),
						CallBuiltin(NewSymbol("int", NewSymbol("lt", nil)), NewInteger(2)),
						BrTrue(NewSymbol("if_00000001", nil), NewSymbol("endif_00000001", nil)),
					),
				NewBlock("if_00000001").
					AddInstructions(
						VarPush(NewSymbol("n", nil)),
						StackType(NewSymbol("int", nil)),
						VarFree(NewSymbol("n", nil)),
						Return(NewInteger(1)),
					),
				NewBlock("endif_00000001").
					AddInstructions(
						VarPush(NewSymbol("n", nil)),
						ConstPush(NewInteger(1)),
						CallBuiltin(NewSymbol("int", NewSymbol("sub", nil)), NewInteger(2)),
						VarPush(NewSymbol("n", nil)),
						ConstPush(NewInteger(2)),
						CallBuiltin(NewSymbol("int", NewSymbol("sub", nil)), NewInteger(2)),
						CallBuiltin(NewSymbol("int", NewSymbol("add", nil)), NewInteger(2)),
						Call(NewSymbol("fib", nil), NewInteger(1)),
						StackType(NewSymbol("int", nil)),
						VarFree(NewSymbol("n", nil)),
						Return(NewInteger(1)),
					),
			),
	)
	fmt.Println(module.String())
}
