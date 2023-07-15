package internal

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	Def = lexer.MustStateful(lexer.Rules{
		"Root": {
			{Name: `whitespace`, Pattern: `\s+`, Action: nil},
			{Name: `Import`, Pattern: "@import", Action: nil},
			{Name: `Varible`, Pattern: "@var", Action: nil},
			{Name: `Constant`, Pattern: "@const", Action: nil},
			{Name: `Function`, Pattern: "@fn", Action: nil},
			{Name: `Struct`, Pattern: "@struct", Action: nil},
			{Name: `Enum`, Pattern: "@enum", Action: nil},
			{Name: `Interface`, Pattern: "@interface", Action: nil},
			{Name: `Return`, Pattern: "@return", Action: nil},
			{Name: `If`, Pattern: "@if", Action: nil},
			{Name: `Else`, Pattern: "else", Action: nil},
			{Name: `ElseIf`, Pattern: "elif", Action: nil},
			{Name: `While`, Pattern: "@while", Action: nil},
			{Name: `For`, Pattern: "@for", Action: nil},
			{Name: `Foreach`, Pattern: "@foreach", Action: nil},
			{Name: `As`, Pattern: "as", Action: nil},
			{Name: `In`, Pattern: "in", Action: nil},
			{Name: `NoError`, Pattern: "ner", Action: nil},
			{Name: `Float`, Pattern: `\d+\.\d+`, Action: nil},
			{Name: `Integer`, Pattern: `\d+`, Action: nil},
			{Name: `Symbol`, Pattern: `\w+`, Action: nil},
			{Name: `stringStart`, Pattern: `"`, Action: lexer.Push("String")},
			{Name: `LeftCBracket`, Pattern: `{`, Action: nil},
			{Name: `RightCBracket`, Pattern: `}`, Action: nil},
			{Name: `LeftParen`, Pattern: `\(`, Action: nil},
			{Name: `RightParen`, Pattern: `\)`, Action: nil},
			{Name: `LeftABracket`, Pattern: `<`, Action: nil},
			{Name: `RightABracket`, Pattern: `>`, Action: nil},
			{Name: `DoubleColon`, Pattern: `::`, Action: nil},
			{Name: `Colon`, Pattern: `:`, Action: nil},
			{Name: `Comma`, Pattern: `,`, Action: nil},
			{Name: `UpperComma`, Pattern: `'`, Action: nil},
			{Name: `Dot`, Pattern: `\.`, Action: nil},
			{Name: `Assign`, Pattern: `=`, Action: nil},
			{Name: `RightArrow`, Pattern: `->`, Action: nil},
			{Name: `Brace`, Pattern: `#`, Action: nil},
		},
		"String": {
			{Name: "stringEnd", Pattern: `"`, Action: lexer.Pop()},
			{Name: "String", Pattern: `(\\"|[^"])*`, Action: nil},
		},
	})
	Parser = participle.MustBuild[Program](
		participle.Lexer(Def),
		participle.UseLookahead(1),
		participle.Union[Global](
			&SingleConst{},
			&MultiConst{},
			&SingleVar{},
			&MultiVar{},
			&SingleImport{},
			&MultiImport{},
			&Enum{},
		),
		participle.Union[DataType](
			&ComplexType{},
			&PrimitiveType{},
		),
		participle.Union[Atom](
			&Integer{},
			&Float{},
			&String{},
			&Bool{},
			&List{},
			&Map{},
		),
		participle.Union[ExprArg](
			&Integer{},
			&Float{},
			&String{},
			&Bool{},
			&Expression{},
			&List{},
			&Map{},
		),
	)
)

type Global interface{ gl() }
type Local interface{ lc() }
type DataType interface{ dt() }
type ExprArg interface{ ex() }
type Atom interface{ at() }

type Program struct {
	Body []Global `parser:"@@+"`
}

type Enum struct {
	Ident string        `parser:"'@enum' @Symbol"`
	Type  PrimitiveType `parser:"(':' @@)?"`
	Items []EnumItem    `parser:"'{' @@+ '}'"`
}

func (gl *Enum) gl() {}

type EnumItem struct {
	Name  string `parser:"@Symbol"`
	Value Atom   `parser:"('->' @@)?"`
}

type SingleImport struct {
	Import ImportBody `parser:"'@import' @@"`
}

func (gl *SingleImport) gl() {}

type MultiImport struct {
	Imports []ImportBody `parser:"'@import' '{' @@+ '}'"`
}

func (gl *MultiImport) gl() {}

type ImportBody struct {
	Path  String `parser:"@@ 'as' "`
	Ident string `parser:"@Symbol"`
}

type SingleConst struct {
	Const ConstBody `parser:"'@const' @@"`
}

func (gl *SingleConst) gl() {}

type MultiConst struct {
	Consts []ConstBody `parser:"'@const' '{' @@+ '}'"`
}

func (gl *MultiConst) gl() {}

type ConstBody struct {
	Ident string   `parser:"@Symbol ':'"`
	Type  DataType `parser:"@@ '='"`
	Value Atom     `parser:"@@"`
}

type SingleVar struct {
	Var VarBody `parser:"'@var' @@"`
}

func (gl *SingleVar) gl() {}

type MultiVar struct {
	Vars []VarBody `parser:"'@var' '{' @@+ '}'"`
}

func (gl *MultiVar) gl() {}

type VarBody struct {
	Ident string   `parser:"@Symbol ':'"`
	Type  DataType `parser:"@@ '='"`
	Value ExprArg  `parser:"@@"`
}

type PrimitiveType struct {
	Type string `parser:"@Symbol"`
}

func (dt *PrimitiveType) dt() {}

type ComplexType struct {
	Type     string     `parser:"@Symbol '<'"`
	Subtypes []DataType `parser:"(@@ (',' @@)* )+ '>'"`
}

func (dt *ComplexType) dt() {}

type Expression struct {
	Function string    `parser:"'(' @Symbol"`
	Args     []ExprArg `parser:"@@* ')'"`
}

func (ex *Expression) ex() {}

type Integer struct {
	Value string `parser:"@Integer"`
}

func (at *Integer) at() {}
func (ea *Integer) ex() {}

type Float struct {
	Value string `parser:"@Float"`
}

func (at *Float) at() {}
func (ea *Float) ex() {}

type String struct {
	Value string `parser:"@String"`
}

func (at *String) at() {}
func (ea *String) ex() {}

type Bool struct {
	Value string `parser:"@('true' | 'false')"`
}

func (at *Bool) at() {}
func (ea *Bool) ex() {}

type List struct {
	Items []Atom `parser:"UpperComma '(' @@* ')'"`
}

func (at *List) at() {}
func (ea *List) ex() {}

type Map struct {
	Items []MapItem `parser:"'#' '(' @@* ')'"`
}

type MapItem struct {
	Key   Atom `parser:"@@ '::'"`
	Value Atom `parser:"@@"`
}

func (at *Map) at() {}
func (ea *Map) ex() {}
