package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	(5 6 1231)
	(+ (567 123)
		(123 567))
	(def a 12)

	}
	`

	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.LeftParen, "("},
		{tokens.Int, "5"},
		{tokens.Int, "6"},
		{tokens.Int, "1231"},
		{tokens.RightParen, ")"},

		{tokens.LeftParen, "("},
		{tokens.Plus, "+"},
		{tokens.LeftParen, "("},
		{tokens.Int, "567"},
		{tokens.Int, "123"},
		{tokens.RightParen, ")"},
		{tokens.LeftParen, "("},
		{tokens.Int, "123"},
		{tokens.Int, "567"},
		{tokens.RightParen, ")"},
		{tokens.RightParen, ")"},

		{tokens.LeftParen, "("},
		{tokens.Identifier, "def"},
		{tokens.Identifier, "a"},
		{tokens.Int, "12"},
		{tokens.RightParen, ")"},

		{tokens.Illegal, "}"},
	}
	reader := strings.NewReader(input)
	tokenChan := make(chan tokens.Token)
	_, tokens := NewLexer(reader, tokenChan)
	for i, tt := range tests {
		tok := <-tokens
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
