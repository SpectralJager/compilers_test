package lexer

import "testing"

func TestLexer(t *testing.T) {
	code := `
@import "std" as std;

@const pi:float = 3.14;

@fn main:void() {
	@var x:int = 2;
	@set x = (std/floatToInt (add (std/intToFloat x) pi))
	@if (eq x 10) {
		(std/println (std/intToString x))
	}
}
`
	lexer := NewLexer(code)
	tokens := lexer.Lex()
	if len(lexer.errors) != 0 {
		for _, e := range lexer.errors {
			t.Logf("Error: %v", e)
		}
		t.FailNow()
	}
	for _, tok := range *tokens {
		t.Logf("%s\n", tok.String())
	}
	// t.FailNow()
}
