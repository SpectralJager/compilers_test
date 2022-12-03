package tokens

import "fmt"

// TokenType -----------------------------
type TokenType int

const (
	EOF TokenType = iota
	Illegal

	// data
	Int    // 123
	Float  // 123.123
	Symbol // abc ab_c
	String // "string"
	Bool   // true false

	// delimiters
	LParen   // (
	RParen   // )
	LBrace   // {
	RBrace   // }
	LBracket // [
	RBracket // ]

	// keywords
	Def // binding symbol to something
)

// Representation of TokenType as String
var tokenTypeString = map[TokenType]string{
	EOF:      "EOF",
	Illegal:  "Illegal",
	Int:      "Int",
	Float:    "Float",
	String:   "String",
	Bool:     "Bool",
	LParen:   "(",
	RParen:   ")",
	LBrace:   "{",
	RBrace:   "}",
	LBracket: "[",
	RBracket: "]",
	Def:      "Def",
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
	return fmt.Sprintf("{Type: %v, Value: %v}", t.Type.String(), t.Value)
}

// -----------------------------------------
