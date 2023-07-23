package parser

import "github.com/alecthomas/participle/v2/lexer"

var (
	Lexer = lexer.MustStateful(lexer.Rules{
		"Root": {
			{Name: "commentary", Pattern: `\/\/[\S \t]*\n`, Action: nil},
			{Name: "Function", Pattern: `@fn`, Action: nil},
			{Name: "Varible", Pattern: `@var`, Action: nil},
			{Name: "Set", Pattern: `@set`, Action: nil},
			{Name: "While", Pattern: `@while`, Action: nil},
			{Name: "Each", Pattern: `@each`, Action: nil},
			{Name: "If", Pattern: `@if`, Action: nil},
			{Name: "Else", Pattern: `else`, Action: nil},
			{Name: "ListStart", Pattern: `'\(`, Action: nil},
			{Name: "Range", Pattern: `<-`, Action: nil},
			{Name: "LParen", Pattern: `\(`, Action: nil},
			{Name: "RParen", Pattern: `\)`, Action: nil},
			{Name: "LTBracet", Pattern: `<`, Action: nil},
			{Name: "RTBracet", Pattern: `>`, Action: nil},
			{Name: "LCBracet", Pattern: `{`, Action: nil},
			{Name: "RCBracet", Pattern: `}`, Action: nil},
			{Name: "Colon", Pattern: `:`, Action: nil},
			{Name: "Semicolon", Pattern: `;`, Action: nil},
			{Name: "Assign", Pattern: `=`, Action: nil},
			{Name: "Symbol", Pattern: `[a-zA-Z_]+`, Action: nil},
			{Name: "Integer", Pattern: `[0-9]+`, Action: nil},
			{Name: "whitespace", Pattern: `[\n\t\r ]+`, Action: nil},
		},
	})
)
