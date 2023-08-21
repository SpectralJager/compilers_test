package compiler

import (
	"github.com/alecthomas/participle/v2"
)

var Parser = participle.MustBuild[ProgramAST](
	participle.Lexer(Lexer),
	participle.Union[Global](
		&FunctionAST{},
		&VaribleAST{},
	),
	participle.Union[Local](
		&IfAST{},
		&ReturnAST{},
		&SymbolExpressionAST{},
		&VaribleAST{},
		&SetAST{},
	),
	participle.Union[Expression](
		&SymbolExpressionAST{},
		&SymbolAST{},
		&IntegerAST{},
		&FloatAST{},
	),
	participle.Union[DataType](
		&SimpleDataTypeAST{},
	),
)
