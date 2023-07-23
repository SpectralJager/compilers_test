package parser

import (
	"gl/internal/ast"

	"github.com/alecthomas/participle/v2"
)

var (
	Parser = participle.MustBuild[ast.Program](
		participle.Lexer(Lexer),
		participle.UseLookahead(1),
		participle.Union[ast.Global](
			&ast.Function{},
			&ast.Varible{},
		),
		participle.Union[ast.Local](
			&ast.Varible{},
			&ast.Set{},
			&ast.FunctionCall{},
			&ast.WhileLoop{},
			&ast.EachLoop{},
			&ast.IfStatement{},
		),
		participle.Union[ast.DataType](
			&ast.ComplexType{},
			&ast.SimpleType{},
		),
		participle.Union[ast.Expression](
			&ast.FunctionCall{},
			&ast.Symbol{},
			&ast.Integer{},
			&ast.List{},
		),
		participle.Union[ast.Atom](
			&ast.Integer{},
			&ast.List{},
		),
	)
)
