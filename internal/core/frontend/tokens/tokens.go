package tokens

import (
	"fmt"
	"grimlang/internal/core/frontend/utils"
)

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Position utils.Position
}

const (
	Illegal = "illegal"
	EOF     = "EOF"

	// Identifiers + Literals
	Identifier = "identifier" // print, x ,y, z, ...
	Int        = "int"        // 1239164198273
	Float      = "float"      // 12.2
	String     = "string"     // "x", "y"

	// Operators
	Quote  = "'"
	Assign = "="
	//// Math
	Plus      = "+"
	Minus     = "-"
	Multimply = "*"
	Divide    = "/"

	// Separators
	LeftParen    = "("
	RightParen   = ")"
	LeftSBracet  = "["
	RightSBracet = "]"
	LeftCBracet  = "{"
	RightCBracet = "}"

	// keywords
	Fn  = "fn"  // lambda function
	Def = "def" // bind value to symbol
	Do  = "do"  // multiple expressions
	Let = "let" // local varibles
)

// func (tt TokenType) String() string {
// 	return fmt.Sprintf("%s", tt)
// }

func (t *Token) String() string {
	return fmt.Sprintf("Type: %s, Value: %s, Position: %q", t.Type, t.Literal, &t.Position)
}
