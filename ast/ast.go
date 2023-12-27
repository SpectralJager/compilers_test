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
	Kind string
	Body []Global
}

type ConstantDecl struct {
	_global
	_local
	Identifier string
	Value      Atom
}

type VariableDecl struct {
	_local
	Identifier string
	Type       Type
	Value      Expression
}

type FunctionDecl struct {
	_global
	Identifier string
	Arguments  []*VariableDefn
	Return     Type
	Body       []Local
}

type VariableDefn struct {
	_ast
	Identifier string
	Type       Type
}

type SymbolCall struct {
	_local
	_expression
	Call      *SymbolExpr
	Arguments []Expression
}

type ReturnStmt struct {
	_local
	Value Expression
}

type SymbolExpr struct {
	_expression
	Identifier string
	Next       *SymbolExpr
}

type IntAtom struct {
	_expression
	_atom
	Value int
}

type IntType struct {
	_type
	Identifier string
}
