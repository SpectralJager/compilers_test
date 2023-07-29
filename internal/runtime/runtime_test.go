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
			&ir.Integer{Value: 60},
		},
		Globals: map[string]ir.ISymbolDef{
			"res": &ir.VaribleDef{Name: "a", Type: &ir.Primitive{Name: "int"}},
			"sum": &ir.FunctionDef{
				Name: "sum",
				// Arguments: []ir.IDataType{
				// 	&ir.Primitive{Name: "int"},
				// 	&ir.Primitive{Name: "int"},
				// },
				Returns: []ir.IDataType{
					&ir.Primitive{Name: "int"},
				},
			},
		},
		InitCode: []ir.IInstruction{
			&ir.Call{FuncName: "sum"},
			&ir.GlobalSave{Symbol: "res"},
		},
		Functions: map[string]ir.Function{
			"sum": {
				Name: "sum",
				BodyCode: []ir.IInstruction{
					&ir.Load{ConstIndex: 0},
					&ir.Load{ConstIndex: 1},
					&ir.CallBuildin{FuncName: "iadd"},
					&ir.Return{Count: 1},
				},
			},
		},
	}
	vm := VM{}
	vm.MustExecute(program)
	fmt.Println(vm.GlobalFrame.String())
}
