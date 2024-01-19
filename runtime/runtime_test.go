package runtime

import (
	"fmt"
	"testing"
)

func TestRuntime(t *testing.T) {
	context := NewContext("text", nil)
	context.Insert(
		NewConstantSymbol("alpha", NewIntLit(12)),
	)
	context.Insert(
		NewVariableSymbol("test", NewIntType(), NewIntLit(0)),
	)
	test_variable := context.Search("test")
	test_variable.Set(NewIntLit(10))
	fmt.Println(test_variable)
	context.Update(test_variable)
	fmt.Println(context)
}
