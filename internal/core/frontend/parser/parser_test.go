package parser

import (
	"grimlang/internal/core/frontend/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	code := `
	(add 12 (mul 2 2))
	(def a 12)
	(set a (add a 12))
	(fn test [] (
		(ret 12)
	))
	(test)
	(def b (add 12 (mul a 2)))
	`
	tests := []struct {
		ExpectedLiteral string
	}{
		{"s-expr"},
		{"def"},
		{"set"},
		{"fn"},
		{"s-expr"},
		{"def"},
	}
	lex := lexer.NewLexer(code)
	toks := lex.Run()
	pars := NewParser(toks)
	prog := pars.Run()
	for i, test := range tests {
		if test.ExpectedLiteral != prog.Body[i].Type() {
			t.Fatalf("[#%d] Want %v, got %v", i, test.ExpectedLiteral, prog.Body[i].Type())
		}
	}
}
