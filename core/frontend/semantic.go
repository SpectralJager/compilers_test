package frontend

import (
	"fmt"
)

/*
global pkgName |>

	#symbolTable:
		fn add: ...int -> int
		fn sub: ...int -> int
		fn mul: ...int -> int
		fn div: ...int -> int
		fn printf: format ...strings -> void
		fn main: void -> void
		const a:int
		var b:int
	function main |>
		#symbolTable:
			const a:int
			const b:int
		local ifBody_{sometoken1}
		local elifBody_{sometoken1}_1
		local elifBody_{sometoken1}_2
		local elseBody_{sometoken1}
		local whileBody_{sometoken2}
		local elseBody_{sometoken2}
		local forBody_{sometoken3} |>
			#symbolTable:
				var i:int
*/
type MetaData map[string]any

func (md MetaData) appendSymbol(symbol string, definition SymbolDefinition) error {
	symbolTable, ok := md["symbolTable"]
	if !ok {
		return fmt.Errorf("symbol table not found")
	}
	return symbolTable.(SymbolTable).appendSymbol(symbol, definition)
}

// scopes
type ScopeKind string

const (
	ScopeGlobal ScopeKind = "global"
)

// symbol table
type SymbolTable map[string]SymbolDefinition

func (st SymbolTable) appendSymbol(symbol string, definition SymbolDefinition) error {
	if def, ok := st[symbol]; ok {
		return fmt.Errorf("symbol %s already exists with definition: %s", symbol, def)
	}
	st[symbol] = definition
	return nil
}

type SymbolDefinition interface {
	fmt.Stringer
	symbdef()
}

func (sd FunctionDefinition) symbdef() {}
func (sd VaribleDefinition) symbdef()  {}
func (sd ConstantDefinition) symbdef() {}
func (sd DataTypeDefinition) symbdef() {}

type FunctionDefinition struct {
	Return string   `json:"return"`
	Args   []string `json:"args"`
}

func (df FunctionDefinition) String() string {
	return fmt.Sprintf("fn %s -> %s", df.Return, df.Args)
}

type VaribleDefinition struct {
	Type string `json:"type"`
}

func (df VaribleDefinition) String() string {
	return fmt.Sprintf("var %s", df.Type)
}

type ConstantDefinition struct {
	Type string `json:"type"`
}

func (df ConstantDefinition) String() string {
	return fmt.Sprintf("const %s", df.Type)
}

type DataTypeDefinition struct {
	Kind string `json:"type"`
}

func (df DataTypeDefinition) String() string {
	return fmt.Sprintf("%s type", df.Kind)
}

func CollectMeta(node Node) error {
	switch node := node.(type) {
	case *ProgramNode:
		return _programmMeta(node)
	default:
		return fmt.Errorf("unexpected node type %T", node)
	}
}

func _programmMeta(node *ProgramNode) error {
	node.Meta = make(MetaData)
	node.Meta["scope"] = string(ScopeGlobal)
	node.Meta["symbolTable"] = make(SymbolTable)
	node.Meta.appendSymbol("int", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("float", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("string", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("bool", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("iadd", FunctionDefinition{"int", []string{"...int"}})
	for _, item := range node.Body {
		switch item := item.(type) {
		case *ConstNode:
			symbol := resolveSymbol(item.Name.(*SymbolNode))
			dataType, err := resolveDataTypeFromNode(item.DataType)
			if err != nil {
				return err
			}
			if err := node.Meta.appendSymbol(symbol, ConstantDefinition{dataType}); err != nil {
				return err
			}
		case *VarNode:
			symbol := resolveSymbol(item.Name.(*SymbolNode))
			dataType, err := resolveDataTypeFromNode(item.DataType)
			if err != nil {
				return err
			}
			if err := node.Meta.appendSymbol(symbol, VaribleDefinition{dataType}); err != nil {
				return err
			}
		case *FunctionNode:
			symbol := resolveSymbol(item.Name.(*SymbolNode))
			dataType, err := resolveDataTypeFromNode(item.ReturnType)
			if err != nil {
				return err
			}
			definition := FunctionDefinition{
				Return: dataType,
			}
			for _, arg := range item.Params {
				dataType, err = resolveDataTypeFromNode(arg.(*ParamNode).DataType)
				if err != nil {
					return err
				}
				definition.Args = append(definition.Args, dataType)
			}
			if err := node.Meta.appendSymbol(symbol, definition); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected node %T", item)
		}
	}
	return nil
}

func resolveSymbol(node *SymbolNode) string {
	return node.Value.Value
}

func resolveDataTypeFromNode(node Node) (string, error) {
	switch dataType := node.(type) {
	case *SymbolNode:
		return dataType.Value.Value, nil
	default:
		return "", fmt.Errorf("unsupported data type %T", node)
	}
}
