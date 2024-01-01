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
	if len(os.Args) != 2 {
		fmt.Printf("usage: grim <filename>")
		os.Exit(1)
	}
	file, err := os.ReadFile(os.Args[1])
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
