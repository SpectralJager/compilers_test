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
	DocString string              `parser:"@String"`
	Body      []BlockContext      `parser:"@@+ 'end' ';'"`
}

func (pkgCtx *FunctionCommand) pkg() {}

// Block Context
type BlockContext interface {
	blk()
}

type ReturnCommand struct {
	ReturnValue string `parser:"'ret' @Int ';'"`
}

func (blkCtx *ReturnCommand) blk() {}

type LetCommand struct {
	Symbol SymbolDeclaration   `parser:"'let' @@ "`
	Value  ExpressionArguments `parser:" @@ ';'"`
}

func (blkCtx *LetCommand) blk() {}

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
}

type Int struct {
	Value string `parser:" @Int "`
}

func (at *Int) atom() {}
func (e *Int) expr()  {}

type Symbol struct {
	Value string `parser:" @Ident "`
}

func (at *Symbol) atom() {}
func (e *Symbol) expr()  {}

type Float struct {
	Value string `parser:" @Float "`
}

func (at *Float) atom() {}
func (e *Float) expr()  {}

type String struct {
	Value string `parser:" @String "`
}

func (at *String) atom() {}
func (e *String) expr()  {}

// Utils

type SymbolDeclaration struct {
	Name string `parser:"@Ident"`
	Type string `parser:"':' @Ident"`
}
