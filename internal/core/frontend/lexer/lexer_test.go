package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	code := `
	(add 1 22)
	`

	tests := []struct {
		expectedType  tokens.TokenType
		expectedValue string
	}{
		{tokens.LParen, "("},
		{tokens.Symbol, "add"},
		{tokens.Number, "1"},
		{tokens.Number, "22"},
		{tokens.RParen, ")"},
		{tokens.EOF, ""},
	}

	lex := NewLexer(strings.NewReader(code))
	tokList := lex.Run()
	for n, test := range tests {
		tok := tokList[n]
		if tok.Type != test.expectedType {
			t.Fatalf("[%d] Token type - %s, want - %s", n, tok.Type.String(), test.expectedType.String())
		}
		if tok.Value != test.expectedValue {
			t.Fatalf("[%d] Token value - %v, want - %v", n, tok.Value, test.expectedValue)
		}
	}
}
