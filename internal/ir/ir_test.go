package ir

import (
	"fmt"
	"testing"
)

func TestIr(t *testing.T) {
	mod := NewModule(
		"test",
		"",
		nil,
		NewChunk(
			"GLOBAL",
			OPBlockBegin(NewSymbol("global", nil)),
			OPVarNew(NewSymbol("n", nil), NewSymbol("int", nil)),
			OPVarSet(NewSymbol("n", nil), NewInteger(40)),
			OPCall(NewSymbol("main", nil), NewInteger(1)),
			OPBlockEnd(),
		),
		NewChunk(
			"main",
			OPBlockBegin(NewSymbol("function", nil)),
			OPVarNew(NewSymbol("res", nil), NewSymbol("int", nil)),
			OPVarPush(NewSymbol("n", nil)),
			OPCall(NewSymbol("fib", nil), NewInteger(1)),
			OPVarPop(NewSymbol("res", nil)),
			OPVarPush(NewSymbol("res", nil)),
			OPCallb(NewSymbol("io", NewSymbol("println", nil)), NewInteger(1)),
			OPBlockEnd(),
		),
		NewChunk(
			"fib",
			OPBlockBegin(NewSymbol("function", nil)),
			OPVarNew(NewSymbol("n", nil), NewSymbol("int", nil)),
			OPVarPop(NewSymbol("n", nil)),
			OPVarPush(NewSymbol("n", nil)),
			OPConstPush(NewInteger(2)),
			OPCallb(NewSymbol("int", NewSymbol("lt", nil)), NewInteger(2)),
			OPJumpFalse(NewSymbol("if_end_token", nil)),
			OPBlockBegin(NewSymbol("local", nil)),
			OPVarPush(NewSymbol("n", nil)),
			OPStackType(NewSymbol("int", nil)),
			OPReturn(NewInteger(1)),
			OPBlockEnd(),
			OPLabel(NewSymbol("if_end_token", nil)),
			OPVarPush(NewSymbol("n", nil)),
			OPConstPush(NewInteger(1)),
			OPCallb(NewSymbol("int", NewSymbol("sub", nil)), NewInteger(2)),
			OPVarPush(NewSymbol("n", nil)),
			OPConstPush(NewInteger(2)),
			OPCallb(NewSymbol("int", NewSymbol("sub", nil)), NewInteger(2)),
			OPCallb(NewSymbol("int", NewSymbol("add", nil)), NewInteger(2)),
			OPCall(NewSymbol("fib", nil), NewInteger(1)),
			OPStackType(NewSymbol("int", nil)),
			OPReturn(NewInteger(1)),
			OPBlockEnd(),
		),
	)
	fmt.Println(mod.String())
}
