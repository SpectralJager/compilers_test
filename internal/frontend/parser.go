package frontend

import "github.com/alecthomas/participle/v2"

var parser = participle.MustBuild[Programm](
	participle.Lexer(lex),
	participle.Unquote("String"),
	participle.Union[GlobalCom](&FnCom{}, &GlobVar{}, &Constant{}),
	participle.Union[FnBody](&Begin{}, &Nil{}, &Int{}, &Float{}, &String{}, &Bool{}),
	participle.Union[BeginBody](&Expression{}, &Let{}, &Set{}, &Constant{}, &IfCom{}, &CondCom{}),
	participle.Union[ExprArgs](&Expression{}, &Symbol{}, &Int{}, &Float{}, &String{}, &Bool{}),
	participle.Union[Atom](&Symbol{}, &Int{}, &Float{}, &String{}, &Bool{}, &Nil{}),
	participle.Union[IfBody](&Begin{}, &IfCom{}, &CondCom{}, &Expression{}, &Set{}),
)
