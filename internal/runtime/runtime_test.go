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
			&ir.Integer{Value: 12},
		},
		Globals: map[string]ir.ISymbolDef{
			"sum": &ir.FunctionDef{
				Name: "sum",
				Arguments: []ir.IDataType{
					&ir.Primitive{Name: "int"},
				},
			},
			"res": &ir.VaribleDef{Name: "res", Type: &ir.Primitive{Name: "int"}},
		},
		InitCode: []ir.IInstruction{
			&ir.Load{ConstIndex: 0},
			&ir.Call{FuncName: "sum"},
			&ir.GlobalSave{Symbol: "res"},
			&ir.Halt{},
		},
		Functions: map[string]*ir.Function{
			"sum": {
				Name: "sum",
				Locals: map[string]ir.ISymbolDef{
					"n": &ir.VaribleDef{Name: "n", Type: &ir.Primitive{Name: "n"}},
				},
				BodyCode: []ir.IInstruction{
					&ir.LocalSave{Symbol: "n"},
					&ir.Load{ConstIndex: 0},
					&ir.LocalLoad{Symbol: "n"},
					&ir.CallBuildin{FuncName: "iadd"},
					&ir.Return{Count: 1},
				},
			},
		},
	}
	vm := VM{}
	vm.MustExecute(&program)
	fmt.Println(vm.GlobalFrame().String())
	fmt.Println(vm.Stack.StackTrace())
}
