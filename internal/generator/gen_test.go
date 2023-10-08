package gen

import (
	"fmt"
	"grimlang/internal/parser"
	"os"
	"testing"
)

func TestIrGen(t *testing.T) {
	data, err := os.ReadFile("../../src/fib.grim")
	if err != nil {
		t.Fatalf("cant read file: %v", err)
	}
	prog, err := parser.Parser.ParseBytes("test", data)
	if err != nil {
		t.Fatalf("cant parse file: %v", err)
	}
	m := GenIrModule(prog)
	fmt.Println(m.String())
}
