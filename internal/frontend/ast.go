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
type ExprArgs interface {
	exprArgs()
}
type Atom interface {
	atom()
}
type Command interface {
	command()
}

// Programm
type Programm struct {
	Body []GlobalCom `@@*`
}

// Utils
type VarSymb struct {
	Symb Symbol `@@`
	Type Symbol `":" @@`
}

// Commands
type FnCom struct {
	Name      Symbol    `"(" "fn" @@`
	RetType   Symbol    `":" @@`
	Params    []VarSymb `"[" @@* "]"`
	DocString String    `@@?`
	Body      FnBody    `@@ ")"`
}

func (glCom *FnCom) globalCom() {}

type Begin struct {
	BegBody []BeginBody `"(" "begin" @@+ `
	Atm     Atom        `@@? ")"`
}

func (fnBd *Begin) fnBody() {}

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

// Expression
type Expression struct {
	Op   Symbol     `"(" @@`
	Args []ExprArgs `@@* ")"`
}

func (exAr *Expression) exprArgs()  {}
func (bgBd *Expression) beginBody() {}

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

// Nil
type Nil struct {
	Value string `"nil"`
}

func (fnBd *Nil) fnBody() {}
func (atm *Nil) atom()    {}
