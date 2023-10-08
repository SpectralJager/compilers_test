package main

import (
	"fmt"
	"gl/internal/compiler"
	"gl/internal/ir"
	"gl/internal/runtime"
	"log"
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
	data, err := os.ReadFile("test.grim")
	if err != nil {
		panic(err)
	}
	programm, err := compiler.Parser.ParseBytes("test.grim", data)
	if err != nil {
		log.Fatalf("can't parse program: %s", err)
	}
	var g compiler.Generator
	res, err := g.GenerateProgram(*programm)
	if err != nil {
		log.Fatalf("can't generate program: %s", err)
	}
	return &res
}
