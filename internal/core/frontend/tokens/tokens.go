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
	Illegal = "Illegal"
	EOF     = "EOF"

	// Identifiers + Literals
	Identifier = "Identifier" // print, x ,y, z, ...
	Int        = "Int"        // 1239164198273
	String     = "String"     // "x", "y"

	// Operators
	Apostrophe = "'"
	Assign     = "="
	Plus       = "+"
	Minus      = "-"
	Bang       = "!"
	Asterisk   = "*"
	Slash      = "/"

	// Separators
	LeftParen  = "("
	RightParen = ")"
)

// func (tt TokenType) String() string {
// 	return fmt.Sprintf("%s", tt)
// }

func (t *Token) String() string {
	return fmt.Sprintf("Type: %v, Value: %s, Position: %q", &t.Type, t.Literal)
}
