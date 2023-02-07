package frontend

// ------------Unions-----------------

// body for global context
type GlobalBody interface {
	global()
}

// body for local context (functions, loops, condition statements, etc.)
type LocalBody interface {
	local()
}

// atom union for strings, numbers, symbols, lists, structs, maps.
type Atom interface {
	atom()
}

// expression arguments
type ExpressionArguments interface {
	exprArg()
}

// Programm
type Programm struct {
	Package string       `parser:"" json:"package"`
	Body    []GlobalBody `parser:"@@*" json:"body"`
}

// ------------Commands-----------------

// Function command
type FunctionCom struct {
	Symbol    SymbolDecl   `parser:"'(' 'fn' @@" json:"symbol"`
	Args      []SymbolDecl `parser:"'[' @@* ']'" jons:"args"`
	DocString string       `parser:"@String?" json:"docString"`
	Body      []LocalBody  `parser:"@@+ ')'" json:"body"`
}

func (gb *FunctionCom) global() {}

// Return command
type ReturnCom struct {
	Ret Atom `parser:"'(' 'ret' @@ ')'" json:"ret"`
}

func (lb *ReturnCom) local() {}

// Constant command
type ConstantCom struct {
	Symbol SymbolDecl `parser:"'(' 'const' @@" json:"symbol"`
	Value  Atom       `parser:"@@ ')'" json:"value"`
}

func (gb *ConstantCom) global() {}
func (lb *ConstantCom) local()  {}

// Global varible command
type GlobalVaribleCom struct {
	Symbol SymbolDecl `parser:"'(' 'var' @@" json:"symbol"`
	Value  Atom       `parser:"@@ ')'" json:"value"`
}

func (gb *GlobalVaribleCom) global() {}

// local varible command
type LocalVaribleCom struct {
	Symbol SymbolDecl          `parser:"'(' 'let' @@" json:"symbol"`
	Value  ExpressionArguments `parser:"@@ ')'" json:"value"`
}

func (lb *LocalVaribleCom) local() {}

// ------------Expression-----------------
type Expression struct {
	Function Symbol              `parser:"'(' @@" json:"function"`
	Args     ExpressionArguments `parser:"@@* ')'" json:"args"`
}

func (lb *Expression) local()   {}
func (ea *Expression) exprArg() {}

// ------------Atoms-----------------

// Integer
type Integer struct {
	Value string `parser:"@Int" json:"value"`
}

func (at *Integer) atom()    {}
func (ea *Integer) exprArg() {}

// Double
type Double struct {
	Value string `parser:"@Double" json:"value"`
}

func (at *Double) atom()    {}
func (ea *Double) exprArg() {}

// String
type String struct {
	Value string `parser:"@String" json:"value"`
}

func (at *String) atom()    {}
func (ea *String) exprArg() {}

// Bool
type Bool struct {
	Value string `parser:"@Bool" json:"value"`
}

func (at *Bool) atom()    {}
func (ea *Bool) exprArg() {}

// Symbol
type Symbol struct {
	Value string `parser:"@Ident" json:"value"`
}

func (at *Symbol) atom()    {}
func (ea *Symbol) exprArg() {}

// List
type List struct {
	Items []Atom `parser:"'[' @@* ']'" json:"items"`
}

func (at *List) atom()    {}
func (ea *List) exprArg() {}

// List
type Map struct {
	Items []MapPair `parser:"'{' @@* '}'" json:"items"`
}

func (at *Map) atom()    {}
func (ea *Map) exprArg() {}

// ------------Utils-----------------

// symbol declaration
type SymbolDecl struct {
	Name            string                 `parser:"@Ident" json:"name"`
	CompositionType []CompositionTypeIdent `parser:"':' @@*" json:"compositionType"`
	PrimitiveType   string                 `parser:"@Ident" json:"primitiveType"`
}

type CompositionTypeIdent struct {
	Identifier string `parser:"@('['']')|@('{''}')" json:"identifier"`
}

type MapPair struct {
	Key   string `parser:"@String" json:"key"`
	Value Atom   `parser:"':'':' @@" json:"value"`
}
