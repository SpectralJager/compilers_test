package frontend

import (
	"encoding/json"
	"testing"
)

func TestCollectMeta(t *testing.T) {
	code := `
@const alpha:int = 12;
@var beta:float = 12.2;
@fn main:void(a:int b:float) {
	@const alpha:string = "alpha";
	@var gamma:string = "gamma";
	(iadd a 12)
}
	`
	programm := NewParser(*NewLexer(code).Lex()).Parse()
	if err := CollectMeta(programm.(*ProgramNode)); err != nil {
		t.Fatal(err)
	}
	data, _ := json.MarshalIndent(programm, "", "  ")
	t.Fatalf("%s\n", data)
}
