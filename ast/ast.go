package ast

type Ast interface {
	ast()
}

type Global interface {
	glb()
}

type Local interface {
	loc()
}

type Expression interface {
	exp()
}

type Atom interface {
	atm()
}

type Type interface {
	typ()
}

// ======================================

type _ast struct{}

func (*_ast) ast() {}

// --------------------------------------
type _global struct{ _ast }

func (*_global) glb() {}

// --------------------------------------
type _local struct{ _ast }

func (*_local) loc() {}

// --------------------------------------
type _expression struct{ _ast }

func (*_expression) exp() {}

// --------------------------------------
type _atom struct{ _ast }

func (*_atom) atm() {}

// --------------------------------------
type _type struct{ _ast }

func (*_type) typ() {}

// ======================================

type Module struct {
	_ast
	Kind string   `parser:"'(' '@module' @(':main') "`
	Body []Global `parser:" @@+ ')'"`
}

type ConstantDecl struct {
	_global
	_local
	Identifier string `parser:"'(' '@const' @Symbol "`
	Value      Atom   `parser:" @@ ')'"`
}

type VariableDecl struct {
	_local
	Identifier string     `parser:"'(' '@var' @Symbol "`
	Type       Type       `parser:" ':of' @@ "`
	Value      Expression `parser:" @@ ')'"`
}

type FunctionDecl struct {
	_global
	Identifier string          `parser:"'(' '@fn' @Symbol "`
	Arguments  []*VariableDefn `parser:"('[' @@+ ']')?"`
	Return     Type            `parser:"('<' @@ '>')?"`
	Body       []Local         `parser:"':do' @@+ ')'"`
}

type VariableDefn struct {
	_ast
	Identifier string `parser:" @Symbol "`
	Type       Type   `parser:" '::' @@ "`
}

type RecordDefn struct {
	_global
	Identifier string         `parser:"'(' '@record' @Symbol"`
	Fields     []VariableDefn `parser:"':fields' @@+ ')'"`
}

type SymbolCall struct {
	_local
	_expression
	Call      *SymbolExpr  `parser:"'(' @@"`
	Arguments []Expression `parser:" @@* ')'"`
}

type ReturnStmt struct {
	_local
	Value Expression `parser:" '(' '@return' @@? ')'"`
}

type SetStmt struct {
	_local
	Symbol *SymbolExpr `parser:"'(' '@set' @@"`
	Value  Expression  `parser:"@@ ')'"`
}

type IfStmt struct {
	_local
	Condition Expression `parser:"'(' '@if' @@"`
	ThenBody  []Local    `parser:"':then' @@+ "`
	Elif      []Elif     `parser:"@@*"`
	ElseBody  []Local    `parser:"(':else' @@+)? ')'"`
}

type WhileStmt struct {
	_local
	Condition Expression `parser:"'(' '@while' @@"`
	ThenBody  []Local    `parser:"':do' @@+ "`
	ElseBody  []Local    `parser:"(':else' @@+)? ')'"`
}

type SymbolExpr struct {
	_expression
	Identifier string      `parser:"@Symbol"`
	Next       *SymbolExpr `parser:"('/' @@)?"`
}

type NewExpr struct {
	_expression
	Symbol *SymbolExpr  `parser:"'(' '@new' @@"`
	Fields []Expression `parser:"'{' @@+ '}'')'"`
}

type IntAtom struct {
	_expression
	_atom
	Value int `parser:"@Integer"`
}

type BoolAtom struct {
	_expression
	_atom
	Value bool `parser:"@Boolean"`
}

type FloatAtom struct {
	_expression
	_atom
	Value float64 `parser:"@Float"`
}

type StringAtom struct {
	_expression
	_atom
	Value string `parser:"@String"`
}

type ListAtom struct {
	_expression
	_atom
	Type  *ListType    `parser:"@@"`
	Items []Expression `parser:"'{' @@* '}'"`
}

type IntType struct {
	_type
	Identifier string `parser:"@'int'"`
}

type BoolType struct {
	_type
	Identifier string `parser:"@'bool'"`
}

type FloatType struct {
	_type
	Identifier string `parser:"@'float'"`
}
type StringType struct {
	_type
	Identifier string `parser:"@'string'"`
}

type ListType struct {
	_type
	Child Type `parser:"'list' '<' @@ '>'"`
}

type RecordType struct {
	_type
	Symbol *SymbolExpr `parser:"@@"`
}

// ======================================

type Elif struct {
	Condition Expression `parser:"':elif' @@"`
	Body      []Local    `parser:"'=>' @@+"`
}
