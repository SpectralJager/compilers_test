package syntax

import "fmt"

type TokenType = string

const (
	TokenEOF     TokenType = "EOF"
	TokenIllegal           = "ILLEGAL"

	// Delimeters
	TokenLParen    = "LPAREN"
	TokenRParen    = "RPAREN"
	TokenLBracket  = "LBRACKET"
	TokenRBracket  = "RBRACKET"
	TokenLCBracket = "LCBRACKET"
	TokenRCBracket = "RCBRACKET"
	TokenLess      = "LESS"
	TokenMore      = "MORE"

	// Symbols
	TokenColon     = "COLON"
	TokenDColon    = "DOUBLECOLON"
	TokenSemicolon = "SEMiCOLON"
	TokenAssign    = "ASSIGN"
	TokenSlash     = "SLASH"

	// Data tokens
	TokenString = "STRING"
	TokenNumber = "NUMBER"
	TokenSymbol = "SYMBOL"

	// Keywords
	TokenImport = "IMPORT"
	TokenFn     = "FUNCTION"
	TokenIf     = "IF"
	TokenEach   = "EACH"
	TokenVar    = "VAR"
	TokenSet    = "SET"
	TokenConst  = "CONST"
	TokenReturn = "RETURN"
)

var keywords = map[string]TokenType{
	"@import": TokenImport,
	"@fn":     TokenFn,
	"@if":     TokenIf,
	"@each":   TokenEach,
	"@var":    TokenVar,
	"@set":    TokenSet,
	"@const":  TokenConst,
	"@return": TokenReturn,
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
