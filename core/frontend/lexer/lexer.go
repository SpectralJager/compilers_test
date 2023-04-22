package lexer

import "gl/core/frontend/tokens"

type Lexer struct {
	input  string
	tokens []tokens.Token
	errors []error

	pos    int
	line   int
	column int
}
