package syntax

import "fmt"

type TokenType = string

const (
	TokenEOF     TokenType = "EOF"
	TokenIllegal           = "ILLEGAL"

	// Delimeters
	TokenLeftParen   = "("
	TokenRightParen  = ")"
	TokenLeftBrace   = "{"
	TokenRightBrace  = "}"
	TokenLeftSquare  = "["
	TokenRightSquare = "]"

	// Symbols
	TokenColon     = "COLON"
	TokenSemicolon = "SEMiCOLON"
	TokenAssign    = "ASSIGN"

	// Data tokens
	TokenInteger = "INT"
	TokenFloat   = "FLOAT"
	TokenSymbol  = "SYMBOL"

	// Keywords
	TokenConst = "CONST"
	TokenFunc  = "FUNC"
)

var keywords = map[string]TokenType{
	"@const": TokenConst,
	"@fn":    TokenFunc,
}

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
