package lexer

import (
	"grimlang/internal/core/frontend/tokens"
	symboltable "grimlang/internal/core/symbol_table"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	(5 6 1231)
	(+ (567 123)
		(123 567))
	(def a 12)

	(def add (fn [a b c] 
		(* a b c)))

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
		{tokens.Def, "def"},
		{tokens.Identifier, "a"},
		{tokens.Int, "12"},
		{tokens.RightParen, ")"},

		{tokens.LeftParen, "("},
		{tokens.Def, "def"},
		{tokens.Identifier, "add"},
		{tokens.LeftParen, "("},
		{tokens.Fn, "fn"},
		{tokens.LeftSBracet, "["},
		{tokens.Identifier, "a"},
		{tokens.Identifier, "b"},
		{tokens.Identifier, "c"},
		{tokens.RightSBracet, "]"},
		{tokens.LeftParen, "("},
		{tokens.Multimply, "*"},
		{tokens.Identifier, "a"},
		{tokens.Identifier, "b"},
		{tokens.Identifier, "c"},
		{tokens.RightParen, ")"},
		{tokens.RightParen, ")"},
		{tokens.RightParen, ")"},
	}
	st := symboltable.GetSymbolTableEntity()
	reader := strings.NewReader(input)
	tokenChan := make(chan tokens.Token)
	_, toks := NewLexer(reader, tokenChan)
	for i, tt := range tests {
		tok := <-toks
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Type == tokens.Identifier {
			_, err := st.Get(tok.Literal)
			if err != nil {
				t.Fatalf("tests[%d] - %q", i, err)
			}
		}
		if _, ok := tokens.Keywords[tok.Literal]; ok {
			_, err := st.Get(tok.Literal)
			if err == nil {
				t.Fatalf("tests[%d] - keyword %q shouldn't be inserted into symtable", i, &tok)
			}
		}
	}
}
