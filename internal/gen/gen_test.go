package gen

import (
	"fmt"
	"grimlang/internal/parser"
	"os"
	"testing"
)

func TestGenIr(t *testing.T) {
	data, err := os.ReadFile("../../src/func.grim")
	if err != nil {
		t.Fatalf("can't read file: %v", err)
	}
	prog, err := parser.Parser.ParseBytes("test", data)
	if err != nil {
		t.Fatalf("can't parse file: %v", err)
	}
	prog.Name = "test"
	mod, err := IRGenerator{}.GenerateModule(prog)
	if err != nil {
		t.Fatalf("can't generate module: %v", err)
	}
	fmt.Println(mod.String())
}
