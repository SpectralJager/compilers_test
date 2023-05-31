package frontend

import (
	"encoding/json"
	"testing"
)

func TestCollectMeta(t *testing.T) {
	code := `
@const alpha:int = 12.12;
@var beta:float = 12.2;
@fn main:void(a:int b:float) {
	@const alpha:string = "alpha";
	@var gamma:string = "gamma";
	(iadd a 12)
}
	`
	programm := NewParser(*NewLexer(code).Lex()).Parse()
	data, _ := json.MarshalIndent(programm, "", "  ")
	t.Fatalf("%s\n", data)
}
