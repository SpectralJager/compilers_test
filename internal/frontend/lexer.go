package frontend

import "github.com/alecthomas/participle/v2/lexer"

var lex = lexer.MustSimple([]lexer.SimpleRule{
	// delimeters
	{Name: `LParen`, Pattern: `\(`},
	{Name: `RParen`, Pattern: `\)`},
	{Name: `LBracket`, Pattern: `\[`},
	{Name: `RBracket`, Pattern: `\]`},
	{Name: `LCurB`, Pattern: `\{`},
	{Name: `RCurB`, Pattern: `\}`},
	{Name: `Colon`, Pattern: `\:`},
	{Name: `,`, Pattern: `\,`},
	// keywords
	{Name: `Fn`, Pattern: `fn`},
	{Name: `Begin`, Pattern: `begin`},
	{Name: `Let`, Pattern: `let`},
	{Name: `Set`, Pattern: `set`},
	{Name: `Glob`, Pattern: `glob`},
	{Name: `Const`, Pattern: `const`},
	{Name: `If`, Pattern: `if`},
	{Name: `Cond`, Pattern: `cond`},
	{Name: `Dotimes`, Pattern: `dotimes`},
	{Name: `While`, Pattern: `while`},
	{Name: `Iterate`, Pattern: `iterate`},
	{Name: `Bool`, Pattern: `true|false`},
	{Name: `Nil`, Pattern: `nil`},
	// atoms
	{Name: `Ident`, Pattern: `[a-zA-Z][a-zA-Z0-9_]*`},
	{Name: `String`, Pattern: `"(\\"|[^"])*"`},
	{Name: `Float`, Pattern: `[-]?[0-9]+[.][0-9]+`},
	{Name: `Int`, Pattern: `[-]?[0-9]+`},
	// whitespace
	{Name: "whitespace", Pattern: `[\t\n\r ]+`},
	{Name: "comment", Pattern: `(?is);[^\n]*`},
})
