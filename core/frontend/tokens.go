package frontend

import "fmt"

type TokenType uint8

const (
	TokenIllegal TokenType = iota
	TokenEOF

	// literals
	TokenNumber
	TokenString
	TokenSymbol

	// delimeters
	TokenLeftParen
	TokenRightParen
	TokenLeftBracket
	TokenRightBracket
	TokenLeftBrace
	TokenRightBrace

	// special symbols
	TokenColon
	TokenDoubleColon
	TokenSemicolon
	TokenAssign
	TokenSlash

	// keywords
	TokenConst
	TokenVar
	TokenSet
	TokenFn
	TokenImport
	Tokenlambda
	TokenIf
	TokenFor
	TokenWhile

	// reserved symbols
	TokenElif
	TokenElse
	TokenFalse
	TokenTrue
	TokenAs
	TokenTo
	TokenFrom
	TokenList
	TokenMap
)

var tokenTypeMap = map[TokenType]string{
	TokenEOF:     "EOF",
	TokenIllegal: "ILLEGAL",

	TokenLeftParen:    "(",
	TokenRightParen:   ")",
	TokenLeftBracket:  "[",
	TokenRightBracket: "]",
	TokenLeftBrace:    "{",
	TokenRightBrace:   "}",

	TokenColon:       ":",
	TokenDoubleColon: "::",
	TokenSemicolon:   ";",
	TokenAssign:      "=",
	TokenSlash:       "/",

	TokenConst:  "@const",
	TokenVar:    "@var",
	TokenSet:    "@set",
	TokenFn:     "@fn",
	TokenImport: "@import",
	Tokenlambda: "@lambda",
	TokenIf:     "@if",
	TokenFor:    "@for",
	TokenWhile:  "@while",

	TokenElif:  "elif",
	TokenElse:  "else",
	TokenFalse: "false",
	TokenTrue:  "true",
	TokenAs:    "as",
	TokenTo:    "to",
	TokenFrom:  "from",
	TokenList:  "list",
	TokenMap:   "map",

	TokenNumber: "NUMBER",
	TokenString: "STRING",
	TokenSymbol: "SYMBOL",
}

// keywords
var keywordMap = map[string]TokenType{
	"@const":  TokenConst,
	"@var":    TokenVar,
	"@set":    TokenSet,
	"@fn":     TokenFn,
	"@import": TokenImport,
	"@lambda": Tokenlambda,
	"@if":     TokenIf,
	"@for":    TokenFor,
	"@while":  TokenWhile,
}

func IsKeyword(val string) (TokenType, bool) {
	tt, ok := keywordMap[val]
	return tt, ok
}

// reserved symbols
var reservedMap = map[string]TokenType{
	"elif":  TokenElif,
	"else":  TokenElse,
	"false": TokenFalse,
	"true":  TokenTrue,
	"as":    TokenAs,
	"to":    TokenTo,
	"from":  TokenFrom,
	"list":  TokenList,
	"map":   TokenMap,
}

func IsReserved(val string) (TokenType, bool) {
	tt, ok := reservedMap[val]
	return tt, ok
}

// Token
type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

func (t Token) String() string {
	return fmt.Sprintf("%d:%d(Type %s, Value %s)", t.Line, t.Column, tokenTypeMap[t.Type], t.Value)
}

func NewToken(tt TokenType, v string, lp, cp int) *Token {
	return &Token{
		Type:   tt,
		Value:  v,
		Line:   lp,
		Column: cp,
	}
}
