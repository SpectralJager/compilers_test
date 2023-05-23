package parser

import (
	"encoding/json"
	"gl/core/frontend/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	code := `
@const a:int = 1;
@var b:int = (add 2 a);
@fn main:void() {
	@const a:int = 1;
	@set b = a;
}
	`
	lex := lexer.NewLexer(code)
	tokens := lex.Lex()
	if len(lex.Error()) != 0 {
		for _, e := range lex.Error() {
			t.Logf("Error: %v", e)
		}
		t.FailNow()
	}
	prs := NewParser(*tokens)
	programm := prs.Parse()
	if len(prs.Errors()) != 0 {
		for _, e := range prs.Errors() {
			t.Logf("Error: %v", e)
		}
		t.FailNow()
	}
	data, _ := json.MarshalIndent(programm, "", "\t")
	t.Fatalf("%s\n", data)
}
