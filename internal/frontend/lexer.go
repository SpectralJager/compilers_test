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
	{Name: `Comma`, Pattern: `\,`},
	// keywords
	{Name: `Package`, Pattern: `package`},
	{Name: `Fn`, Pattern: `fn`},
	{Name: `Let`, Pattern: `let`},
	{Name: `Set`, Pattern: `set`},
	{Name: `Var`, Pattern: `var`},
	{Name: `Const`, Pattern: `const`},
	{Name: `If`, Pattern: `if`},
	{Name: `Cond`, Pattern: `cond`},
	{Name: `Dotimes`, Pattern: `dotimes`},
	{Name: `While`, Pattern: `while`},
	{Name: `Nil`, Pattern: `nil`},
	{Name: `Bool`, Pattern: `true|false`},
	// atoms
	{Name: `Ident`, Pattern: `[a-zA-Z][a-zA-Z_]*`},
	{Name: `String`, Pattern: `"(\\"|[^"])*"`},
	{Name: `Double`, Pattern: `[-]?[0-9]+[.][0-9]+`},
	{Name: `Int`, Pattern: `[-]?[0-9]+`},
	// whitespaces
	{Name: "whitespace", Pattern: `[\t\n\r ]+`},
	{Name: "comment", Pattern: `(?is);[^\n]*`},
})
