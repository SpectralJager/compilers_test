package internal

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	Def = lexer.MustStateful(lexer.Rules{
		"Root": {
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
	Parser = participle.MustBuild[Program](
		participle.Lexer(Def),
		participle.UseLookahead(1),
		participle.Union[Global](
			&Function{},
			&Varible{},
		),
		participle.Union[Local](
			&Varible{},
			&Set{},
			&FunctionCall{},
			&WhileLoop{},
			&EachLoop{},
			&IfStatement{},
		),
		participle.Union[DataType](
			&ComplexType{},
			&SimpleType{},
		),
		participle.Union[Expression](
			&FunctionCall{},
			&Symbol{},
			&Integer{},
			&List{},
		),
		participle.Union[Atom](
			&Integer{},
			&List{},
		),
	)
)

type Global interface{ glob() }

func (g *Function) glob() {}
func (g *Varible) glob()  {}

type Local interface{ local() }

func (l *Varible) local()      {}
func (l *Set) local()          {}
func (l *FunctionCall) local() {}
func (l *WhileLoop) local()    {}
func (l *EachLoop) local()     {}
func (l *IfStatement) local()  {}

type DataType interface{ dt() }

func (dt *SimpleType) dt()  {}
func (dt *ComplexType) dt() {}

type Expression interface{ expr() }

func (ex *FunctionCall) expr() {}
func (ex *Symbol) expr()       {}
func (ex *Integer) expr()      {}
func (ex *List) expr()         {}

type Atom interface{ atom() }

func (at *Integer) atom() {}
func (at *List) atom()    {}

type Program struct {
	Body []Global `parser:"@@+"`
}

type Function struct {
	Ident       string     `parser:"'@fn' @Symbol"`
	Args        []FuncArg  `parser:"'(' @@* ')'"`
	ReturnTypes []DataType `parser:"'<' @@+ '>'"`
	Body        []Local    `parser:"'{' @@* '}'"`
}

type Varible struct {
	Ident string     `parser:"'@var' @Symbol ':'"`
	Type  DataType   `parser:"@@ "`
	Expr  Expression `parser:"'=' @@ ';'"`
}

type Set struct {
	Ident string     `parser:"'@set' @Symbol"`
	Expr  Expression `parser:"'=' @@ ';'"`
}

type WhileLoop struct {
	Condition Expression `parser:"'@while' @@"`
	Body      []Local    `parser:"'{' @@+ '}'"`
}

type EachLoop struct {
	IteratorIdent  string     `parser:"'@each' @Symbol"`
	IteratorType   DataType   `parser:"':' @@"`
	IterableObject Expression `parser:"'<-' @@"`
	Body           []Local    `parser:"'{' @@+ '}'"`
}

type IfStatement struct {
	Condition Expression `parser:"'@if' @@"`
	Body      []Local    `parser:"'{' @@+ '}'"`
	ElseIf    []ElseIf   `parser:"@@*"`
	Else      *Else      `parser:"@@?"`
}

type ElseIf struct {
	Condition Expression `parser:"@@"`
	Body      []Local    `parser:"'{' @@+ '}'"`
}

type Else struct {
	Body []Local `parser:"'else' '{' @@+ '}'"`
}

type FunctionCall struct {
	Func string       `parser:"'(' @Symbol"`
	Args []Expression `parser:"@@* ')'"`
}

type Symbol struct {
	Value string `parser:"@Symbol"`
}

type Integer struct {
	Value int `parser:"@Integer"`
}

type List struct {
	Values []Atom `parser:"ListStart @@* ')'"`
}

type FuncArg struct {
	Ident string   `parser:"@Symbol"`
	Type  DataType `parser:"':' @@"`
}

type SimpleType struct {
	Name string `parser:"@Symbol"`
}

type ComplexType struct {
	Name     string     `parser:"@Symbol "`
	Subtypes []DataType `parser:"'<' @@+ '>'"`
}
