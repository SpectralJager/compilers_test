package ir

import (
	"fmt"
	"testing"
)

func TestIr(t *testing.T) {
	m := NewModule("test")
	m.WriteInstrs(
		Call(NewSymbol("main", nil), NewInt(0)),
	)
	main := NewFunction(NewSymbol("main", nil))
	main.WriteInstrs()
	m.WriteFunctions(
		main,
	)
	fmt.Println(m.String())
}
