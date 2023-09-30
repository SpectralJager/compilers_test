package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var Lexer = lexer.MustStateful(lexer.Rules{
	"Root": []lexer.Rule{
		{Name: "whitespace", Pattern: `[ \t\r\n]+`},
		{Name: "comment", Pattern: `//[^\n]*\n`},

		{Name: "Varible", Pattern: `@var`},

		{Name: "Symbol", Pattern: `[a-zA-Z_]+[a-zA-Z0-9_]*`},
		{Name: "Integer", Pattern: `[-+]?[0-9]+`},

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

var Parser = participle.MustBuild[ProgramAST](
	participle.Lexer(Lexer),
	participle.Union[Global](
		&VaribleAST{},
	),
	participle.Union[Expression](
		&SymbolAST{},
		&IntegerAST{},
	),
)

type Global interface {
	global()
}

func (g *VaribleAST) global() {}

type Expression interface {
	expr()
}

func (e *SymbolAST) expr()  {}
func (e *IntegerAST) expr() {}

type ProgramAST struct {
	Name string
	Body []Global `parser:"@@+"`
}

type VaribleAST struct {
	Symbol SymbolAST  `parser:"'@var' @@ "`
	Type   TypeAST    `parser:"':' @@ "`
	Value  Expression `parser:" '=' @@ ';'"`
}

type TypeAST struct {
	Primary   SymbolAST `parser:"@@"`
	Secondary []TypeAST `parser:"('<' @@+ '>')?"`
}

type SymbolAST struct {
	Value string `parser:"@Symbol"`
}

type IntegerAST struct {
	Value int `parser:"@Integer"`
}
