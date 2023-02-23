package frontend

import "github.com/alecthomas/participle/v2"

var Parser = participle.MustBuild[Package](
	participle.Lexer(lex),
	participle.Unquote("String"),
	participle.Union[PackageContext](
		&FunctionCommand{},
	),
	participle.Union[BlockContext](
		&ReturnCommand{},
		&LetCommand{},
		&Expression{},
	),
	participle.Union[ExpressionArguments](
		&Expression{},
		&Int{},
		&Symbol{},
		&String{},
		&Float{},
	),
	participle.Union[Atom](
		&Int{},
		&Symbol{},
		&String{},
		&Float{},
	),
)
