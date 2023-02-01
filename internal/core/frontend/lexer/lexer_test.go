package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	"testing"
)

func Test(t *testing.T) {
	code := `
	(fn main:void [] 
		(begin 
			(let a:i32 12)
			nil)	
	)
	`

	tests := []struct {
		expectedType  tokens.TokenType
		expectedValue string
	}{
		{tokens.LParen, "("},
		{tokens.Fn, "fn"},
		{tokens.Symbol, "main"},
		{tokens.Colon, ":"},
		{tokens.Symbol, "void"},
		{tokens.LBracket, "["},
		{tokens.RBracket, "]"},
		{tokens.LParen, "("},
		{tokens.Begin, "begin"},
		{tokens.LParen, "("},
		{tokens.Let, "let"},
		{tokens.Symbol, "a"},
		{tokens.Colon, ":"},
		{tokens.Symbol, "i32"},
		{tokens.Int, "12"},
		{tokens.RParen, ")"},
		{tokens.Nil, "nil"},
		{tokens.RParen, ")"},
		{tokens.RParen, ")"},
		{tokens.EOF, "EOF"},
	}

	lex := NewLexer(code)
	tokList := lex.Run()
	for n, test := range tests {
		tok := tokList[n]
		if tok.Type != test.expectedType {
			t.Fatalf("[%d] Token type - %v, want - %v, at %d row", n, tok.Type.String(), test.expectedType.String(), tok.Row)
		}
		if tok.Value != test.expectedValue {
			t.Fatalf("[%d] Token value - %v, want - %v at %d row", n, tok.Value, test.expectedValue, tok.Row)
		}
	}
}
