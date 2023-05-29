package frontend

import (
	"encoding/json"
	"testing"
)

func TestSemantic(t *testing.T) {
	code := `
@const alpha:int = 12;
@var x:float = 12.2;
@fn main:void(a:int b:float) {}
	`
	programm := NewParser(*NewLexer(code).Lex()).Parse()
	if err := CollectMeta(programm); err != nil {
		t.Fatal(err)
	}
	data, _ := json.MarshalIndent(programm, "", "  ")
	t.Fatalf("%s\n", data)
}
