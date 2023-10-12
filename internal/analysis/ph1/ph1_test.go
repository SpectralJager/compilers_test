package ph1

import (
	"fmt"
	"grimlang/internal/gen"
	"grimlang/internal/parser"
	"os"
	"testing"
)

func TestPh1(t *testing.T) {
	data, err := os.ReadFile("../../../src/func.grim")
	if err != nil {
		t.Fatalf("can't read file: %v", err)
	}
	prog, err := parser.Parser.ParseBytes("test", data)
	if err != nil {
		t.Fatalf("can't parse file: %v", err)
	}
	prog.Name = "test"
	mod := gen.GenModule(prog)
	fmt.Println(mod.String())
	AnalyseModule(mod)
	fmt.Println(mod.String())
}
