package parser

import (
	"grimlang/internal/core/frontend/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	code := `
	`
	tests := []struct {
		ExpectedLiteral string
	}{}
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
