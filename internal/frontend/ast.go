package frontend

// unions
type GlobalCom interface {
	globalCom()
}
type LocalCom interface {
	localCom()
}
type FnBody interface {
	fnBody()
}
type BeginBody interface {
	beginBody()
}
type IfBody interface {
	ifBody()
}
type ExprArgs interface {
	exprArgs()
}
type Atom interface {
	atom()
}
type Type interface {
	typE()
}

// Programm
type Programm struct {
	Body []GlobalCom `@@*`
}

// Utils
type VarSymb struct {
	Symb Symbol `@@`
	Tp   Type   `":" @@`
}
type SimpleType struct {
	Symb Symbol `@@`
}

func (tp *SimpleType) typE() {}

type SequenceType struct {
	Tp Symbol `(("[" "]") | ("{" "}")) @@ `
}

func (tp *SequenceType) typE() {}

type Case struct {
	BoolExpr Expression `"(" @@`
	Body     IfBody     `@@ ")"`
}

type MapItem struct {
	Key   String `@@`
	Value Atom   `":"":" @@`
}

// Commands
// Constant value
type Constant struct {
	Var   VarSymb `"(" "const" @@ `
	Value Atom    `@@ ")"`
}

func (glCom *Constant) globalCom() {}
func (bgBd *Constant) beginBody()  {}

// Global commands
// Function declaration
type FnCom struct {
	Name      Symbol    `"(" "fn" @@`
	RetType   Symbol    `":" @@`
	Params    []VarSymb `"[" @@* "]"`
	DocString String    `@@?`
	Body      FnBody    `@@ ")"`
}

func (glCom *FnCom) globalCom() {}

// global varible
type GlobVar struct {
	Var   VarSymb `"(" "glob" @@ `
	Value Atom    `@@ ")"`
}

func (glCom *GlobVar) globalCom() {}

// Local commands
// begin block
type Begin struct {
	BegBody []BeginBody `"(" "begin" @@+ `
	Atm     Atom        `@@? ")"`
}

func (fnBd *Begin) fnBody() {}
func (ifBd *Begin) ifBody() {}

type Let struct {
	Varible VarSymb  `"(" "let" @@`
	Value   ExprArgs `@@ ")"`
}

func (bgBd *Let) beginBody() {}

type Set struct {
	Varible Symbol   `"(" "set" @@`
	Value   ExprArgs `@@ ")"`
}

func (bgBd *Set) beginBody() {}
func (ifBd *Set) ifBody()    {}

type IfCom struct {
	BoolExpr Expression `"(" "if" @@`
	ThenBody IfBody     `@@`
	ElseBody IfBody     `@@ ")"`
}

func (bgBd *IfCom) beginBody() {}
func (ifBd *IfCom) ifBody()    {}

type CondCom struct {
	Cases   []Case `"(" "cond" @@+`
	Default IfBody `@@ ")"`
}

func (bgBd *CondCom) beginBody() {}
func (ifBd *CondCom) ifBody()    {}

// Expression
type Expression struct {
	Op   Symbol     `"(" @@`
	Args []ExprArgs `@@* ")"`
}

func (exAr *Expression) exprArgs()  {}
func (bgBd *Expression) beginBody() {}
func (ifBd *Expression) ifBody()    {}

// Atoms---------------------
// Symbol
type Symbol struct {
	Value string `@Ident`
}

func (exAr *Symbol) exprArgs() {}
func (atm *Symbol) atom()      {}

// Int
type Int struct {
	Value string `@Int`
}

func (fnBd *Int) fnBody()   {}
func (exAr *Int) exprArgs() {}
func (atm *Int) atom()      {}

// Float
type Float struct {
	Value string `@Float`
}

func (fnBd *Float) fnBody()   {}
func (exAr *Float) exprArgs() {}
func (atm *Float) atom()      {}

// String
type String struct {
	Value string `@String`
}

func (fnBd *String) fnBody()   {}
func (exAr *String) exprArgs() {}
func (atm *String) atom()      {}

// Bool
type Bool struct {
	Value string `@Bool`
}

func (fnBd *Bool) fnBody()   {}
func (exAr *Bool) exprArgs() {}
func (atm *Bool) atom()      {}

// List
type List struct {
	Value []Atom `"[" @@* "]"`
}

func (fnBd *List) fnBody()   {}
func (exAr *List) exprArgs() {}
func (atm *List) atom()      {}

// Map
type Map struct {
	Value []MapItem `"{" @@* "}"`
}

func (fnBd *Map) fnBody()   {}
func (exAr *Map) exprArgs() {}
func (atm *Map) atom()      {}

// Nil
type Nil struct {
	Value string `"nil"`
}

func (fnBd *Nil) fnBody() {}
func (atm *Nil) atom()    {}
