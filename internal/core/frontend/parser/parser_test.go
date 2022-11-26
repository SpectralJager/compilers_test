package parser

import (
	"grimlang/internal/core/frontend/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	code := `
	(add 1 2)
	(add 1 (add 1 1))
	`
	tests := []struct {
		ExpectedLiteral string
	}{
		{"(add)"},
		{"(add)"},
	}
	lex := lexer.NewLexer(code)
	toks := lex.Run()
	pars := NewParser(toks)
	prog := pars.Run()
	for i, test := range tests {
		if test.ExpectedLiteral != prog.Expresions[i].TokenLiteral() {
			t.Fatalf("[#%d] Want %v, got %v", i, test.ExpectedLiteral, prog.Expresions[i].TokenLiteral())
		}
	}
}
