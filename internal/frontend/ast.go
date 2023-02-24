package frontend

// Package Context
type PackageContext interface {
	pkg()
}

type Package struct {
	Filename  string
	Name      string           `parser:"'package' @Ident ';'"`
	Variables PackageVariables `parser:"@@?"`
	Body      []PackageContext `parser:"@@*"`
}

type PackageVariables struct {
	Variables []PackageVariable `parser:"'var' @@+ 'end' ';'"`
}

type PackageVariable struct {
	Symbol SymbolDeclaration `parser:"@@"`
	Value  Atom              `parser:"@@ ';'"`
}

type FunctionCommand struct {
	Symbol    SymbolDeclaration   `parser:"'fn' @@"`
	Args      []SymbolDeclaration `parser:"'(' @@* ')'"`
	DocString string              `parser:"@String?"`
	Body      []BlockContext      `parser:"@@+ 'end' ';'"`
}

func (pkgCtx *FunctionCommand) pkg() {}

// Block Context
type BlockContext interface {
	blk()
}

type ReturnCommand struct {
	ReturnValue ExpressionArguments `parser:"'ret' @@ ';'"`
}

func (blkCtx *ReturnCommand) blk() {}

type LetCommand struct {
	Symbol SymbolDeclaration   `parser:"'let' @@ "`
	Value  ExpressionArguments `parser:" @@ ';'"`
}

func (blkCtx *LetCommand) blk() {}

type SetCommand struct {
	Symbol Symbol              `parser:"'set' @@ "`
	Value  ExpressionArguments `parser:" @@ ';'"`
}

func (blkCtx *SetCommand) blk() {}

type IfCommand struct {
	Condition    Expression     `parser:"'if' @@ "`
	ThenBody     []BlockContext `parser:" @@+ "`
	ElseIfBodies []ElseIf       `parser:" @@* "`
	ElseBodies   []BlockContext `parser:"('else' @@+)? 'end'';'"`
}

func (blkCtx *IfCommand) blk() {}

// Expression Arguments
type ExpressionArguments interface {
	expr()
}

type Expression struct {
	Symbol string                `parser:"'(' @Ident"`
	Args   []ExpressionArguments `parser:" @@* ')'"`
}

func (blkCtx *Expression) blk() {}
func (e *Expression) expr()     {}

// Atom
type Atom interface {
	atom()
	AtomType() string
	AtomValue() string
}

type Int struct {
	Value string `parser:" @Int "`
}

func (at *Int) atom() {}
func (at *Int) AtomType() string {
	return "int"
}
func (at *Int) AtomValue() string {
	return at.Value
}
func (e *Int) expr() {}

type Symbol struct {
	Value string `parser:" @Ident "`
}

func (at *Symbol) atom() {}
func (at *Symbol) AtomType() string {
	return "symbol"
}
func (at *Symbol) AtomValue() string {
	return at.Value
}
func (e *Symbol) expr() {}

type Float struct {
	Value string `parser:" @Float "`
}

func (at *Float) atom() {}
func (at *Float) AtomType() string {
	return "float"
}
func (at *Float) AtomValue() string {
	return at.Value
}
func (e *Float) expr() {}

type String struct {
	Value string `parser:" @String "`
}

func (at *String) atom() {}
func (at *String) AtomType() string {
	return "string"
}
func (at *String) AtomValue() string {
	return at.Value
}
func (e *String) expr() {}

// Utils

type SymbolDeclaration struct {
	Name string `parser:"@Ident"`
	Type string `parser:"':' @Ident"`
}

type ElseIf struct {
	Condition ExpressionArguments `parser:"'elif' @@ "`
	ThenBody  []BlockContext      `parser:" @@+ "`
}
