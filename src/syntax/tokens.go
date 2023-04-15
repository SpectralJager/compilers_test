package syntax

import "fmt"

type TokenType = string

const (
	TokenEOF TokenType = "EOF"

	// Delimeters
	TokenLParen    = ")"
	TokenRParen    = "("
	TokenLBracket  = "["
	TokenRBracket  = "]"
	TokenLCBracket = "{"
	TokenRCBracket = "}"
	TokenColon     = ":"
	TokenSemicolon = ";"

	// Data tokens
	TokenString = "STRING"
	TokenNumber = "NUMBER"
	TokenSymbol = "SYMBOL"

	// Keywords
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

func NewToken(typ TokenType, value string, line int, column int) *Token {
	return &Token{
		Type:   typ,
		Value:  value,
		Line:   line,
		Column: column,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%d:%d{Type: %s, Value: %s}", t.Line, t.Column, t.Type, t.Value)
}
