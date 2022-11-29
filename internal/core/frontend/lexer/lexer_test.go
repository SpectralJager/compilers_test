package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	"testing"
)

func Test(t *testing.T) {
	code := `
	()
	{}
	[]

	add
	div
	sub
	mul

	and
	or
	not

	nil
	
	def 
	fn
	
	false
	true

	symb

	123
	"123"
	123.123
	`

	tests := []struct {
		expectedType  tokens.TokenType
		expectedValue string
	}{
		{tokens.LParen, "("},
		{tokens.RParen, ")"},
		{tokens.LBrace, "{"},
		{tokens.RBrace, "}"},
		{tokens.LBracket, "["},
		{tokens.RBracket, "]"},

		{tokens.Add, "add"},
		{tokens.Div, "div"},
		{tokens.Sub, "sub"},
		{tokens.Mul, "mul"},

		{tokens.And, "and"},
		{tokens.Or, "or"},
		{tokens.Not, "not"},

		{tokens.Nil, "nil"},

		{tokens.Def, "def"},
		{tokens.Fn, "fn"},

		{tokens.False, "false"},
		{tokens.True, "true"},

		{tokens.Symbol, "symb"},

		{tokens.Number, "123"},
		{tokens.String, "\"123\""},
		{tokens.Float, "123.123"},

		{tokens.EOF, "EOF"},
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
