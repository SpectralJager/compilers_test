package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	"testing"
)

func Test(t *testing.T) {
	code := `
	5
	6
	12
	`

	tests := []struct {
		expectedType  tokens.TokenType
		expectedValue string
	}{
		{tokens.Int, "5"},
		{tokens.Int, "6"},
		{tokens.Int, "12"},
	}

	lex := NewLexer(code)
	tokList := lex.Run()
	for n, test := range tests {
		tok := tokList[n]
		if tok.Type != test.expectedType {
			t.Fatalf("[%d] Token type - %v, want - %v", n, tok.Type.String(), test.expectedType.String())
		}
		if tok.Value != test.expectedValue {
			t.Fatalf("[%d] Token value - %v, want - %v", n, tok.Value, test.expectedValue)
		}
	}
}
