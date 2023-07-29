package runtime

import (
	"fmt"
	"gl/internal/ir"
	"testing"
)

func TestRuntime(t *testing.T) {
	program := ir.Program{
		Name: "test",
		Constants: []ir.IConstant{
			&ir.Integer{Value: 40},
			&ir.Integer{Value: 2},
			&ir.Integer{Value: 1},
		},
		InitCode: []ir.IInstruction{
			&ir.Load{ConstIndex: 0},
			&ir.Load{ConstIndex: 2},
			&ir.Load{ConstIndex: 1},
		},
	}
	vm := VM{
		Stack: make(Stack, 0),
	}
	vm.MustExecute(program)
	defer fmt.Println(vm.StackTrace())
}
