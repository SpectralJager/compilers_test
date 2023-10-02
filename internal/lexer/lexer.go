package lexer

import "github.com/alecthomas/participle/v2/lexer"

var Lexer = lexer.MustStateful(lexer.Rules{
	"Root": {
		{Name: "comment", Pattern: `//[^\n]*\n`},
		{Name: "whitespace", Pattern: `[ \t\r\n]+`},

		{Name: "Function", Pattern: `@fn`},
		{Name: "Varible", Pattern: `@var`},
		{Name: "While", Pattern: `@while`},
		{Name: "If", Pattern: `@if`},
		{Name: "Return", Pattern: `@ret`},
		{Name: "Else", Pattern: `else`},

		{Name: "Symbol", Pattern: `[a-zA-Z_]+[a-zA-Z0-9_]*`},
		{Name: "Integer", Pattern: `[-+]?[0-9]+`},

		{Name: "ArrowR", Pattern: `=>`},
		{Name: "Brace[", Pattern: `\[`},
		{Name: "Brace]", Pattern: `\]`},
		{Name: "Brace(", Pattern: `\(`},
		{Name: "Brace)", Pattern: `\)`},
		{Name: "Brace<", Pattern: `<`},
		{Name: "Brace>", Pattern: `>`},
		{Name: "Brace{", Pattern: `{`},
		{Name: "Brace}", Pattern: `}`},
		{Name: "Colon", Pattern: `:`},
		{Name: "SemiColon", Pattern: `;`},
		{Name: "Assign", Pattern: `=`},
		{Name: "Slash", Pattern: `/`},
	},
})
