package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	Lexer = lexer.MustStateful(lexer.Rules{
		"Root": {
			{Name: "comment", Pattern: `;[^\n]*\n`},
			{Name: "whitespace", Pattern: `[ \t\r\n]+`},

			{Name: "Module", Pattern: `@module`},
			{Name: "Function", Pattern: `@fn`},
			{Name: "Constant", Pattern: `@const`},
			{Name: "Varible", Pattern: `@var`},
			{Name: "Return", Pattern: `@return`},
			{Name: "Set", Pattern: `@set`},
			{Name: "If", Pattern: `@if`},
			{Name: "While", Pattern: `@while`},
			{Name: "Record", Pattern: `@record`},
			{Name: "New", Pattern: `@new`},

			{Name: "MainKW", Pattern: `:main`},
			{Name: "DoKW", Pattern: `:do`},
			{Name: "OfKW", Pattern: `:of`},
			{Name: "ThenKW", Pattern: `:then`},
			{Name: "ElifKW", Pattern: `:elif`},
			{Name: "ElseKW", Pattern: `:else`},
			{Name: "FieldsKW", Pattern: `:fields`},

			{Name: "String", Pattern: `"(\\"|[^"])*"`},
			{Name: "Float", Pattern: `[-+]?[0-9]+\.[0-9]+`},
			{Name: "Integer", Pattern: `[-+]?[0-9]+`},
			{Name: "Boolean", Pattern: `(true|false)`},
			{Name: "Symbol", Pattern: `[a-zA-Z_]+[a-zA-Z0-9_]*`},

			{Name: "Arrow", Pattern: `=>`},
			{Name: "Brace[", Pattern: `\[`},
			{Name: "Brace]", Pattern: `\]`},
			{Name: "Brace(", Pattern: `\(`},
			{Name: "Brace)", Pattern: `\)`},
			{Name: "Brace<", Pattern: `<`},
			{Name: "Brace>", Pattern: `>`},
			{Name: "Brace{", Pattern: `{`},
			{Name: "Brace}", Pattern: `}`},
			{Name: "DColon", Pattern: `::`},
			{Name: "Colon", Pattern: `:`},
			{Name: "Slash", Pattern: `/`},
		},
	})
	Parser = participle.MustBuild[Module](
		participle.Lexer(Lexer),
		participle.UseLookahead(4),
		participle.Union[Global](
			&ConstantDecl{},
			&FunctionDecl{},
			&RecordDefn{},
		),
		participle.Union[Local](
			&ConstantDecl{},
			&VariableDecl{},
			&ReturnStmt{},
			&SetStmt{},
			&SymbolCall{},
			&IfStmt{},
			&WhileStmt{},
		),
		participle.Union[Expression](
			&IntAtom{},
			&BoolAtom{},
			&FloatAtom{},
			&StringAtom{},
			&ListAtom{},
			&SymbolCall{},
			&NewExpr{},
			&SymbolExpr{},
		),
		participle.Union[Atom](
			&IntAtom{},
			&BoolAtom{},
			&FloatAtom{},
			&StringAtom{},
			&ListAtom{},
		),
		participle.Union[Type](
			&IntType{},
			&BoolType{},
			&FloatType{},
			&StringType{},
			&ListType{},
			&RecordType{},
		),
	)
)
