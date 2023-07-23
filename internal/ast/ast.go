package ast

// Unions for parse tree

// interface for global scope stmts
type Global interface{ glob() }

func (g *Function) glob() {}
func (g *Varible) glob()  {}

// interface for local scope stmts
type Local interface{ local() }

func (l *Varible) local()      {}
func (l *Set) local()          {}
func (l *FunctionCall) local() {}
func (l *WhileLoop) local()    {}
func (l *EachLoop) local()     {}
func (l *IfStatement) local()  {}

// interface for simple and complex datatypes
type DataType interface{ dt() }

func (dt *SimpleType) dt()  {}
func (dt *ComplexType) dt() {}

// interface for expressions
type Expression interface{ expr() }

func (ex *FunctionCall) expr() {}
func (ex *Symbol) expr()       {}
func (ex *Integer) expr()      {}
func (ex *List) expr()         {}

// interface for atoms
type Atom interface{ atom() }

func (at *Integer) atom() {}
func (at *List) atom()    {}

// Parse tree

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
