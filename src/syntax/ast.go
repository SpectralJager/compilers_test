package syntax

/*
programm 	= globalUnion* ;
globalUnion	= const | func ;
const		= '@const' symbolDef '=' LITTERAL ';' ;
func		= '@fn' SYMBOL '(' symbolDef* ')' [SYMBOL] '{' localUnion* '}' ;

symbolDef	= SYMBOL ':' SYMBOL ;
*/

// Programm main node
type Program struct {
	Filename string        `json:"filename"`
	Body     []GlobalUnion `json:"body"`
}

// Unions
type GlobalUnion interface {
	glob()
}

type LocalUnion interface {
	local()
}

// ConstNode
type ConstNode struct {
	Symbol SymbolDef `json:"symbol_def"`
	Value  Token     `json:"value"`
}

func (c *ConstNode) glob() {}

// Function node
type FunctionNode struct {
	Symbol     Token        `json:"symbol"`
	ReturnType Token        `json:"return_type"`
	Args       []SymbolDef  `json:"args"`
	Body       []LocalUnion `json:"body"`
}

func (f *FunctionNode) glob() {}

// Utils
type SymbolDef struct {
	Symbol Token `json:"symbol"`
	Type   Token `json:"type"`
}
