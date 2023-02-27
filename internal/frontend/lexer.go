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
	{Name: `SemiColon`, Pattern: `\;`},
	// keywords
	{Name: `Package`, Pattern: `package`},
	{Name: `Fn`, Pattern: `fn`},
	{Name: `Var`, Pattern: `var`},
	{Name: `If`, Pattern: `if`},
	{Name: `Elif`, Pattern: `elif`},
	{Name: `Else`, Pattern: `else`},
	{Name: `While`, Pattern: `while`},
	// atoms
	{Name: `Bool`, Pattern: `true|false`},
	{Name: `Ident`, Pattern: `[a-zA-Z][a-zA-Z_]*`},
	{Name: `String`, Pattern: `"(\\"|[^"])*"`},
	{Name: `Float`, Pattern: `[-]?[0-9]+[.][0-9]*`},
	{Name: `Int`, Pattern: `[-]?[0-9]+`},
	// whitespaces
	{Name: "whitespace", Pattern: `[\t\n\r ]+`},
	{Name: "comment", Pattern: `(?is);[^\n]*`},
})
