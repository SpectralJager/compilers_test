package main

import (
	"fmt"
	"gl/internal/ir"
	"gl/internal/runtime"
	"os"
	"runtime/pprof"
)

func main() {
	program := ir.Program{
		Name: "test",
		Constants: []ir.IConstant{
			&ir.Integer{Value: 35},
			&ir.Integer{Value: 2},
			&ir.Integer{Value: 1},
		},
		Globals: map[string]ir.ISymbolDef{
			"fib": &ir.FunctionDef{
				Name: "fib",
				Arguments: []ir.IDataType{
					&ir.Primitive{Name: "int"},
				},
				Returns: []ir.IDataType{
					&ir.Primitive{Name: "int"},
				},
			},
			"main": &ir.FunctionDef{
				Name: "main",
			},
			"res": &ir.VaribleDef{Name: "res", Type: &ir.Primitive{Name: "int"}},
		},
		InitCode: []ir.IInstruction{
			&ir.Call{FuncName: "main"},
			&ir.Halt{},
		},
		Functions: map[string]*ir.Function{
			"main": {
				Name:   "main",
				Locals: map[string]ir.ISymbolDef{},
				BodyCode: []ir.IInstruction{
					&ir.Load{ConstIndex: 0},
					&ir.Call{FuncName: "fib"},
					&ir.GlobalSave{Symbol: "res"},
					&ir.Return{},
				},
			},
			"fib": {
				Name: "fib",
				Locals: map[string]ir.ISymbolDef{
					"n": &ir.VaribleDef{
						Name: "n",
						Type: &ir.Primitive{Name: "int"},
					},
				},
				BodyCode: []ir.IInstruction{
					&ir.LocalSave{Symbol: "n"},
					&ir.Load{ConstIndex: 1},
					&ir.LocalLoad{Symbol: "n"},
					&ir.CallBuildin{FuncName: "ilt"},
					&ir.ConditionalJump{Address: 15},
					&ir.Load{ConstIndex: 1},
					&ir.LocalLoad{Symbol: "n"},
					&ir.CallBuildin{FuncName: "isub"},
					&ir.Call{FuncName: "fib"},
					&ir.Load{ConstIndex: 2},
					&ir.LocalLoad{Symbol: "n"},
					&ir.CallBuildin{FuncName: "isub"},
					&ir.Call{FuncName: "fib"},
					&ir.CallBuildin{FuncName: "iadd"},
					&ir.Return{Count: 1},
					&ir.LocalLoad{Symbol: "n"},
					&ir.Return{Count: 1},
				},
			},
		},
	}
	vm := runtime.VM{}
	fl, err := os.Create("fib.prof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(fl)
	vm.MustExecute(&program)
	pprof.StopCPUProfile()
	fmt.Println(vm.GlobalFrame().String())
	fmt.Println(vm.Stack.StackTrace())
}
