package parser

import (
	"grimlang/internal/core/frontend/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	code := `
	12
	224
	123123
	`
	tests := []struct {
		ExpectedLiteral string
	}{
		{"12"},
		{"224"},
		{"123123"},
	}
	lex := lexer.NewLexer(code)
	toks := lex.Run()
	pars := NewParser(toks)
	prog := pars.Run()
	for i, test := range tests {
		if test.ExpectedLiteral != prog.Body[i].TokenLiteral() {
			t.Fatalf("[#%d] Want %v, got %v", i, test.ExpectedLiteral, prog.Body[i].TokenLiteral())
		}
	}
}
