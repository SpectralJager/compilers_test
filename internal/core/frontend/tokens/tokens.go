package tokens

import "fmt"

// TokenType -----------------------------
type TokenType int

const (
	EOF TokenType = iota
	Illegal

	// data
	Number // 123
	Symbol // abc ab_c
	String // "string"
	Float  // 123.123

	// delimiters
	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket

	// keywords
	Def // defenition
	Fn  // function definition
	Not
	Neg
	And
	Or
	Nil
	True
	False
	Add
	Sub
	Mul
	Div
	Lt
	Gt
	Leq
	Geq
)

// Representation of TokenType as String
var tokenTypeString = map[TokenType]string{
	EOF:      "EOF",
	Illegal:  "Illegal",
	Number:   "Number",
	Float:    "Flaot",
	String:   "String",
	LParen:   "(",
	RParen:   ")",
	LBrace:   "{",
	RBrace:   "}",
	LBracket: "[",
	RBracket: "]",
	Def:      "Def",
	Fn:       "Fn",
	Not:      "Not",
	Neg:      "Neg",
	And:      "And",
	Or:       "Or",
	Nil:      "Nil",
	True:     "True",
	False:    "False",
	Add:      "Add",
	Sub:      "Sub",
	Mul:      "Mul",
	Div:      "Div",
	Lt:       "Lt",
	Gt:       "Gt",
	Leq:      "Leq",
	Geq:      "Geq",
}

var keywords = map[string]TokenType{
	"def":   Def,
	"fn":    Fn,
	"not":   Not,
	"neg":   Neg,
	"and":   And,
	"or":    Or,
	"nil":   Nil,
	"true":  True,
	"false": False,
	"add":   Add,
	"sub":   Sub,
	"mul":   Mul,
	"div":   Div,
	"lt":    Lt,
	"gt":    Gt,
	"leq":   Leq,
	"geq":   Geq,
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
