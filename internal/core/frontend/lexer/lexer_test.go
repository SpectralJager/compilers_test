package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	"testing"
)

func Test(t *testing.T) {
	code := `
	123
	123.123
	(add 123 123.123)
	(add (add 123 123) sym)
	(def sym 123)
	(def sym (add 123 123))
	(set sym 123.123)
	(fn sym [sl al] (
		(ret (add sl al))
	))
	`

	tests := []struct {
		expectedType  tokens.TokenType
		expectedValue string
	}{
		{tokens.Int, "123"},
		{tokens.Float, "123.123"},
		{tokens.LParen, "("},
		{tokens.Symbol, "add"},
		{tokens.Int, "123"},
		{tokens.Float, "123.123"},
		{tokens.RParen, ")"},
		{tokens.LParen, "("},
		{tokens.Symbol, "add"},
		{tokens.LParen, "("},
		{tokens.Symbol, "add"},
		{tokens.Int, "123"},
		{tokens.Int, "123"},
		{tokens.RParen, ")"},
		{tokens.Symbol, "sym"},
		{tokens.RParen, ")"},
		{tokens.LParen, "("},
		{tokens.Def, "def"},
		{tokens.Symbol, "sym"},
		{tokens.Int, "123"},
		{tokens.RParen, ")"},
		{tokens.LParen, "("},
		{tokens.Def, "def"},
		{tokens.Symbol, "sym"},
		{tokens.LParen, "("},
		{tokens.Symbol, "add"},
		{tokens.Int, "123"},
		{tokens.Int, "123"},
		{tokens.RParen, ")"},
		{tokens.RParen, ")"},
		{tokens.LParen, "("},
		{tokens.Set, "set"},
		{tokens.Symbol, "sym"},
		{tokens.Float, "123.123"},
		{tokens.RParen, ")"},
		{tokens.LParen, "("},
		{tokens.Fn, "fn"},
		{tokens.Symbol, "sym"},
		{tokens.LBracket, "["},
		{tokens.Symbol, "sl"},
		{tokens.Symbol, "al"},
		{tokens.RBracket, "]"},
		{tokens.LParen, "("},
		{tokens.LParen, "("},
		{tokens.Ret, "ret"},
		{tokens.LParen, "("},
		{tokens.Symbol, "add"},
		{tokens.Symbol, "sl"},
		{tokens.Symbol, "al"},
		{tokens.RParen, ")"},
		{tokens.RParen, ")"},
		{tokens.RParen, ")"},
		{tokens.RParen, ")"},
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
