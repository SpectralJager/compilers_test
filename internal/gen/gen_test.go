package gen

import (
	"fmt"
	"grimlang/internal/parser"
	"os"
	"testing"
)

func TestGenIr(t *testing.T) {
	data, err := os.ReadFile("../../src/fib.grim")
	if err != nil {
		t.Fatalf("can't read file: %v", err)
	}
	prog, err := parser.Parser.ParseBytes("test", data)
	if err != nil {
		t.Fatalf("can't parse file: %v", err)
	}
	prog.Name = "test"
	mod := _GenModule(prog)
	fmt.Println(mod.String())
}
