package frontend

import (
	"encoding/json"
	"testing"
)

func TestCollectMeta(t *testing.T) {
	code := `
@const alpha:int = 12;
@var beta:int = (iaddTwo alpha 1.1);
	`
	programm := NewParser(*NewLexer(code).Lex()).Parse()
	programm.(*ProgramNode).Package = "test"
	ctx := NewSemanticContext()
	ctx["buildin"] = NewSymbolTable("buildin")
	semanticAnalyser := NewSemanticAnalyser(ctx)
	err := semanticAnalyser.Semantic(programm)
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.MarshalIndent(ctx, "", "  ")
	t.Fatalf("%s\n", data)
}
