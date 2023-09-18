package ast

import "github.com/alecthomas/participle/v2/lexer"

var Lexer = lexer.MustStateful(lexer.Rules{
	"Root": []lexer.Rule{
		{Name: "whitespace", Pattern: `[ \t\r\n]+`},

		{Name: "Function", Pattern: `@fn`},
		{Name: "If", Pattern: `@if`},
		{Name: "Return", Pattern: `@return`},
		{Name: "Varible", Pattern: `@var`},
		{Name: "Set", Pattern: `@set`},
		{Name: "Else", Pattern: `else`},

		{Name: "Symbol", Pattern: `[a-zA-Z_]+[a-zA-Z0-9_]*`},
		{Name: "Float", Pattern: `[0-9]+\.[0-9]+`},
		{Name: "Integer", Pattern: `[0-9]+`},

		{Name: "Arrow", Pattern: `=>`},
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
	},
})
