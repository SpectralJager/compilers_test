package models_test

import (
	"grimlang/internal/core/models"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	(5 6 1231)
	(+ (567 123)
		(123 567))
	`

	tests := []struct {
		expectedType    models.TokenType
		expectedLiteral string
	}{
		{models.LeftParen, "("},
		{models.Int, "5"},
		{models.Int, "6"},
		{models.Int, "1231"},
		{models.RightParen, ")"},

		{models.LeftParen, "("},
		{models.Plus, "+"},
		{models.LeftParen, "("},
		{models.Int, "567"},
		{models.Int, "123"},
		{models.RightParen, ")"},
		{models.LeftParen, "("},
		{models.Int, "123"},
		{models.Int, "567"},
		{models.RightParen, ")"},
		{models.RightParen, ")"},
	}
	reader := strings.NewReader(input)
	_, tokens := models.NewLexer(reader)
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
