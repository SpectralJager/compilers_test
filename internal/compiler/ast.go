package compiler

type AST interface {
	// ast()
}

type Global interface {
	AST
	global()
}

func (g *FunctionAST) global() {}
func (g *VaribleAST) global()  {}

type Local interface {
	AST
	local()
}

func (l *IfAST) local()               {}
func (l *ReturnAST) local()           {}
func (l *SymbolExpressionAST) local() {}
func (l *VaribleAST) local()          {}
func (l *SetAST) local()              {}

type Expression interface {
	expr()
}

func (e *SymbolExpressionAST) expr() {}
func (e *SymbolAST) expr()           {}
func (e *IntegerAST) expr()          {}
func (e *FloatAST) expr()            {}

type DataType interface {
	AST
	dat()
}

func (d *SimpleDataTypeAST) dat() {}

type ProgramAST struct {
	Name string
	Body []Global `parser:"@@+"`
}

type FunctionAST struct {
	Ident     string          `parser:"'@fn' @Symbol"`
	Arguments []SymbolDeclAST `parser:"'['@@*']'"`
	RetTypes  []DataType      `parser:"('<'@@+'>')?"`
	Body      []Local         `parser:"'{'@@+'}'"`
}

type IfAST struct {
	Expr     Expression `parser:"'@if' @@"`
	ThenBody []Local    `parser:"'{'@@+"`
	ElseBody []Local    `parser:"('else''=>' @@+)?'}'"`
}

type ReturnAST struct {
	Expression Expression `parser:"'@return' @@?';'"`
}

type VaribleAST struct {
	Symbol SymbolDeclAST `parser:"'@var' @@ '='"`
	Value  Expression    `parser:"@@ ';'"`
}

type SetAST struct {
	Ident string     `parser:"'@set' @Symbol '='"`
	Value Expression `parser:"@@ ';'"`
}

type SymbolDeclAST struct {
	Ident    string   `parser:"@Symbol"`
	DataType DataType `parser:"':' @@"`
}

type SimpleDataTypeAST struct {
	Value string `parser:"@Symbol"`
}

type SymbolExpressionAST struct {
	Symbol    string       `parser:"'(' @Symbol"`
	Arguments []Expression `parser:"@@* ')'"`
}

type SymbolAST struct {
	Symbol string `parser:"@Symbol"`
}

type IntegerAST struct {
	Integer int `parser:"@Integer"`
}

type FloatAST struct {
	Integer float64 `parser:"@Float"`
}
