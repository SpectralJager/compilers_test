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
		},
		Globals: []ir.ISymbolDef{
			&ir.FunctionDef{
				Name:      "main",
				Arguments: []ir.IDataType{},
				Returns:   []ir.IDataType{},
			},
		},
		InitCode: ir.NewCode().
			WriteBytes(ir.Call(0)...).
			WriteByte(ir.Halt()),
		Functions: map[string]*ir.Function{
			"main": {
				Name:   "main",
				Locals: []ir.ISymbolDef{},
				BodyCode: ir.NewCode().
					WriteBytes(ir.Load(0)...).
					WriteBytes(ir.Return(1)...),
			},
		},
	}
	fmt.Println(program.String())
	vm := VM{}
	vm.MustExecute(&program)
	fmt.Println(vm.GlobalFrame().String())
	fmt.Println(vm.Stack.StackTrace())
}
