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
	Nil    // null

	// delimiters
	LParen   // (
	RParen   // )
	LBracket // [
	RBracket // ]

	Colon // :

	// keywords
	Let      // binding symbol to something
	Set      // change symbol value
	Fn       // define function
	Begin    // begin block
	Const    // define constant value
	Glob     // define global varible
	If       // if command
	Cond     // multiple condition expresions
	While    // while loop
	Import   // import
	Dotimes  // dotimes loop
	Ret      // return command
	Break    // break current loop command
	Continue // next iteration of current loop command

)

// Representation of TokenType as String
var tokenTypeString = map[TokenType]string{
	EOF:      "EOF",
	Illegal:  "Illegal",
	Int:      "Int",
	Float:    "Float",
	String:   "String",
	Bool:     "Bool",
	Nil:      "Nil",
	Symbol:   "Symbol",
	Colon:    "Colon",
	LParen:   "LParen",
	RParen:   "RParen",
	LBracket: "LBracket",
	RBracket: "RBracket",
	Let:      "Let",
	Set:      "Set",
	Fn:       "Fn",
	Begin:    "Begin",
	Break:    "Break",
	Continue: "Continue",
	Ret:      "Ret",
	Dotimes:  "Dotimes",
	While:    "While",
	Cond:     "Cond",
	If:       "If",
	Const:    "Const",
	Glob:     "Glob",
	Import:   "Import",
}

var keywords = map[string]TokenType{
	"let":      Let,
	"set":      Set,
	"fn":       Fn,
	"begin":    Begin,
	"break":    Break,
	"continue": Continue,
	"ret":      Ret,
	"dotimes":  Dotimes,
	"while":    While,
	"cond":     Cond,
	"if":       If,
	"const":    Const,
	"glob":     Glob,
	"import":   Import,
	// bool values
	"true":  Bool,
	"false": Bool,
	"nil":   Nil,
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
	panic("Undefined TokenType ")
}

// -----------------------------------------

// Token -----------------------------------
type Token struct {
	Type  TokenType
	Value string
	Row   int
}

func NewToken(tt TokenType, val string, row int) *Token {
	return &Token{Type: tt, Value: val, Row: row}
}

func (t *Token) String() string {
	return fmt.Sprintf("{Type: %v, Value: %v, Row: %d}", t.Type.String(), t.Value, t.Row)
}

// -----------------------------------------
