package main

import (
	"bytes"
	"fmt"
	"grimlang/ast"
	"grimlang/context"
	"grimlang/eval"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
)

func main() {
	file, err := os.ReadFile("examples/test.grim")
	if err != nil {
		log.Fatalf("something goes wrong -> %s", err.Error())
	}
	var errBuf bytes.Buffer
	module, err := ast.Parser.ParseBytes("", file, participle.Trace(&errBuf))
	if err != nil {
		fmt.Printf("%s\n", errBuf.String())
		log.Fatalf("something goes wrong -> %s", err.Error())
	}
	builtinContext := context.NewBuiltinContext()
	if err := new(eval.EvalState).EvalModule(builtinContext, module); err != nil {
		log.Fatalf("something goes wrong -> %s", err.Error())
	}
	os.Exit(0)
}
