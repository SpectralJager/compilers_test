package ir

import (
	"fmt"
	"grimlang/internal/object"
	tp "grimlang/internal/type"
	"testing"
)

func TestIr(t *testing.T) {
	m := NewModule("test")
	glob := NewFunction("global")
	glob.PushInstructions(
		VarNew("alpha", tp.NewInt()),
		ConstLoad(object.NewInt(0)),
		StackType(tp.NewInt()),
		VarSave("alpha"),
		Call("main", object.NewInt(0)),
	)
	m.SetGlobal(glob)
	main := NewFunction("main")
	main.PushInstructions(
		ConstLoad(object.NewInt(10)),
		ConstLoad(object.NewInt(20)),
		Call("int/add", object.NewInt(2)),
		StackType(tp.NewInt()),
		VarNew("res", tp.NewInt()),
		VarSave("res"),
	)
	m.PushFunctions(main)

	fmt.Println(m)
}
