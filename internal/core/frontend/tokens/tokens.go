package tokens

import "fmt"

// TokenType -----------------------------
type TokenType int

const (
	EOF TokenType = iota
	Illegal

	// data
	Number // 123 1.23
	Symbol // abc ab_c

	// delimiters
	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket

	// keywords
	Def // defenition
)

// Representation of TokenType as String
var tokenTypeString = map[TokenType]string{
	EOF:      "EOF",
	Illegal:  "Illegal",
	Number:   "Number",
	Def:      "Def",
	LParen:   "(",
	RParen:   ")",
	LBrace:   "{",
	RBrace:   "}",
	LBracket: "[",
	RBracket: "]",
}

var keywords = map[string]TokenType{
	"def": Def,
}

func LookupSymbolType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Symbol
}

func (tt *TokenType) String() string {
	if val, ok := tokenTypeString[*tt]; ok {
		return val
	}
	panic("Undefined TokenType")
}

// -----------------------------------------

// Token -----------------------------------
type Token struct {
	Type  TokenType
	Value string
}

func NewToken(tt TokenType, val string) *Token {
	return &Token{Type: tt, Value: val}
}

func (t *Token) String() string {
	return fmt.Sprintf("{Type: %s, Value: %s", t.Type.String(), t.Value)
}

// -----------------------------------------
