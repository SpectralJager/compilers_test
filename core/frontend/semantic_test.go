package frontend

import (
	"encoding/json"
	"testing"
)

func TestCollectMeta(t *testing.T) {
	code := `
@const alpha:int = 12;

@fn main:int() {
	@var a:int = 12;
	@if (ilt a 10) {
		@var l:int = 2;
		(sum beta a)
	} else {
		(sum a alpha)
	}
	@while (ilt a 10) {
		@var i:int = 0;
		(iadd i 1)
		(iadd a 1)
	}
}

@fn sum:int(a:int b:int) {
	@var result:int = (iadd a b);
}

@var beta:int = (sum alpha 12);
	`
	programm := NewParser(*NewLexer(code).Lex()).Parse()
	programm.(*ProgramNode).Package = "test"
	ctx := NewSemanticContext()
	ctx["buildin"] = NewSymbolTable("buildin")
	semanticAnalyser := NewSemanticAnalyser(ctx)
	err := semanticAnalyser.CollectSymbols(programm)
	if err != nil {
		t.Fatal(err)
	}
	err = semanticAnalyser.Semantic(programm)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.MarshalIndent(ctx, "", "  ")
	t.Fatalf("%s\n", data)
}
