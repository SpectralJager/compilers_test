package frontend

import "github.com/alecthomas/participle/v2"

var Parser = participle.MustBuild[Programm](
	participle.Lexer(lex),
	participle.Unquote("String"),
	participle.Union[GlobalBody](
		&FunctionCom{},
		&ConstantCom{},
		&GlobalVaribleCom{},
	),
	participle.Union[LocalBody](
		&ReturnCom{},
		&ConstantCom{},
		&LocalVaribleCom{},
		&Expression{},
	),
	participle.Union[Atom](
		&Integer{},
		&Double{},
		&String{},
		&Bool{},
		&Symbol{},
		&List{},
		&Map{},
	),
	participle.Union[ExpressionArguments](
		&Expression{},
		&Integer{},
		&Double{},
		&String{},
		&Bool{},
		&Symbol{},
		&List{},
		&Map{},
	),
)
