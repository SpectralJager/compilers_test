package parser

import (
	"grimlang/internal/ast"
	"grimlang/internal/lexer"

	"github.com/alecthomas/participle/v2"
)

var Parser = participle.MustBuild[ast.ProgramAST](
	participle.Lexer(lexer.Lexer),
	participle.Union[ast.GLOBAL](
		&ast.VarAST{},
		&ast.FunctionAST{},
	),
	participle.Union[ast.LOCAL](
		&ast.SCallAST{},
		&ast.VarAST{},
		&ast.ReturnAST{},
	),
	participle.Union[ast.EXPR](
		&ast.SymbolAST{},
		&ast.IntAST{},
		&ast.SCallAST{},
	),
)
