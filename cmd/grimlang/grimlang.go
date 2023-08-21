package main

import (
	"fmt"
	"gl/internal/compiler"
	"gl/internal/ir"
	"gl/internal/runtime"
	"os"
)

func main() {
	program := test()
	fmt.Println(program.String())
	vm := runtime.VM{}
	// fl, err := os.Create("fib.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// pprof.StartCPUProfile(fl)
	vm.MustExecute(program)
	// pprof.StopCPUProfile()
	fmt.Println(vm.GlobalFrame().String())
	fmt.Println(vm.Stack.StackTrace())
}

func test() *ir.Program {
	// 	code := `
	// @var result:int = 0;
	// @fn main[] {
	// 	@var temp:int = 1;
	// 	@if (ieq temp 0) {
	// 		@set result = 10;
	// 	} else {
	// 		@set result = 20;
	// 	}
	// }
	// 	`
	data, err := os.ReadFile("test.grim")
	if err != nil {
		panic(err)
	}
	programm, err := compiler.Parser.ParseString("test.grim", string(data))
	if err != nil {
		panic(err)
	}
	var g compiler.Generator
	res, err := g.GenerateProgram(*programm)
	if err != nil {
		panic(err)
	}
	return &res
}

func fibProg() *ir.Program {
	return &ir.Program{
		Name: "test",
		Constants: []ir.IConstant{
			&ir.Integer{Value: 35},
			&ir.Integer{Value: 2},
			&ir.Integer{Value: 1},
		},
		Globals: []ir.ISymbolDef{
			&ir.FunctionDef{
				Name: "fib",
				Arguments: []ir.IDataType{
					&ir.Primitive{Name: "int"},
				},
				Returns: []ir.IDataType{
					&ir.Primitive{Name: "int"},
				},
			},
			&ir.FunctionDef{
				Name: "main",
			},
			&ir.VaribleDef{Name: "res", Type: &ir.Primitive{Name: "int"}},
		},
		InitCode: ir.NewCode().
			WriteBytes(ir.Call(1)...).
			WriteByte(ir.Halt()),
		Functions: map[string]*ir.Function{
			"main": {
				Name:   "main",
				Locals: []ir.ISymbolDef{},
				BodyCode: ir.NewCode().
					WriteBytes(ir.Load(0)...).
					WriteBytes(ir.Call(0)...).
					WriteBytes(ir.GlobalSave(2)...).
					WriteBytes(ir.Return(0)...),
			},
			"fib": {
				Name: "fib",
				Locals: []ir.ISymbolDef{
					&ir.VaribleDef{
						Name: "n",
						Type: &ir.Primitive{Name: "int"},
					},
				},
				BodyCode: ir.NewCode().
					WriteBytes(ir.LocalSave(0)...).
					WriteBytes(ir.Load(1)...).
					WriteBytes(ir.LocalLoad(0)...).
					WriteBytes(ir.IntFunc(4)...).
					WriteBytes(ir.JumpCondition(0x3c)...).
					WriteBytes(ir.Load(1)...).
					WriteBytes(ir.LocalLoad(0)...).
					WriteBytes(ir.IntFunc(1)...).
					WriteBytes(ir.Call(0)...).
					WriteBytes(ir.Load(2)...).
					WriteBytes(ir.LocalLoad(0)...).
					WriteBytes(ir.IntFunc(1)...).
					WriteBytes(ir.Call(0)...).
					WriteBytes(ir.IntFunc(0)...).
					WriteBytes(ir.Return(1)...).
					WriteBytes(ir.LocalLoad(0)...).
					WriteBytes(ir.Return(1)...),
			},
		},
	}
}
