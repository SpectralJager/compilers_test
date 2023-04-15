package syntax

/*
@import "std" as std;

pub fn main() void {
	std.printf("%s\n", "Hello, world!");
}
*/

type Token = string

const (
	TokenEOF Token = "EOF"

	// Delimeters
	TokenLParen    = ")"
	TokenRParen    = "("
	TokenLBracket  = "["
	TokenRBracket  = "]"
	TokenLCBracket = "{"
	TokenRCBracket = "}"
	TokenColon     = ":"
	TokenSemicolon = ";"

	// Data tokens
	TokenString = "STRING"
	TokenNumber = "NUMBER"

	// Keywords
)
