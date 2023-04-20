package syntax

// ------------- Interfaces for all syntax nodes
// General interface for all syntax nodes.
type Node interface {
	NodeType() string
}

// Interface for global expressions.
type GlobalExpression interface {
	global()
}

// Interface for local expressions.
type LocalExpression interface {
	local()
}

// Interface for litterals
type Litteral interface {
	litteral()
}

// ------------- Nodes
// Program root node
type Programm struct {
	FileName    string             `json:"file_name"`
	GlobalNodes []GlobalExpression `json:"globals`
}

// ------------- Keyword Nodes
// @import expression
type ImportExpression struct {
	Sybol      SymbolLitteral `json:"import_sybol"`
	ModuleName StringLitteral `json:"module_name"`
}

func (e *ImportExpression) global()          {}
func (n *ImportExpression) NodeType() string { return "import_expression" }

// @const expression
type ConstExpression struct {
	Sybol    SymbolLitteral `json:"const_sybol"`
	DataType SequenceSymbol `json:"const_type"`
	Value    Litteral       `json:"value"`
}

func (e *ConstExpression) global()          {}
func (n *ConstExpression) NodeType() string { return "const_expression" }

// ------------- Litteral Nodes
// Symbol litteral
type SymbolLitteral struct {
	Value string `json:"value"`
}

func (l *SymbolLitteral) litteral()        {}
func (n *SymbolLitteral) NodeType() string { return "symbol_litteral" }

// symbol/symbol/.../symbol litteral
type SequenceSymbol struct {
	Symbols []SymbolLitteral `json:"symbols"`
}

func (l *SequenceSymbol) litteral()        {}
func (n *SequenceSymbol) NodeType() string { return "sequance_litteral" }

// String literal
type StringLitteral struct {
	Value string `json:"value"`
}

func (l *StringLitteral) litteral()        {}
func (n *StringLitteral) NodeType() string { return "string_litteral" }

// ------------- Utiles Nodes
