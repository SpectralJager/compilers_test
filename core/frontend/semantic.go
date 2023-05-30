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

func CollectMeta(node *ProgramNode) error {
	node.Meta = make(MetaData)
	node.Meta["symbolTable"] = make(SymbolTable)
	node.Meta.appendSymbol("int", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("float", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("string", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("bool", DataTypeDefinition{"buildin"})
	node.Meta.appendSymbol("iadd", FunctionDefinition{"int", []string{"...int"}})
	for _, item := range node.Body {
		switch item := item.(type) {
		case *ConstNode:
			err := _constMeta(item)
			if err != nil {
				return err
			}
			err = node.Meta.appendSymbol(item.Meta["name"].(string), ConstantDefinition{item.Meta["type"].(string)})
			if err != nil {
				return err
			}
		case *VarNode:
			err := _varMeta(item)
			if err != nil {
				return err
			}
			err = node.Meta.appendSymbol(item.Meta["name"].(string), VaribleDefinition{item.Meta["type"].(string)})
			if err != nil {
				return err
			}
		case *FunctionNode:
			err := _funcMeta(item)
			if err != nil {
				return err
			}
			err = node.Meta.appendSymbol(item.Meta["name"].(string), FunctionDefinition{item.Meta["returnType"].(string), item.Meta["paramsType"].([]string)})
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unexpected node %T", item)
		}
	}
	return nil
}

func _constMeta(node *ConstNode) error {
	node.Meta = make(MetaData)
	node.Meta["name"] = resolveSymbol(node.Name.(*SymbolNode))
	dataType, err := resolveDataTypeFromNode(node.DataType)
	if err != nil {
		return err
	}
	node.Meta["type"] = dataType
	value, err := resolveValueFromNode(node.Value)
	if err != nil {
		return err
	}
	node.Meta["value"] = value
	return nil
}

func _varMeta(node *VarNode) error {
	node.Meta = make(MetaData)
	node.Meta["name"] = resolveSymbol(node.Name.(*SymbolNode))
	dataType, err := resolveDataTypeFromNode(node.DataType)
	if err != nil {
		return err
	}
	node.Meta["type"] = dataType
	return nil
}

func _funcMeta(node *FunctionNode) error {
	node.Meta = make(MetaData)
	node.Meta["symbolTable"] = make(SymbolTable)
	node.Meta["name"] = resolveSymbol(node.Name.(*SymbolNode))
	returnType, err := resolveDataTypeFromNode(node.ReturnType)
	if err != nil {
		return err
	}
	node.Meta["returnType"] = returnType
	node.Meta["paramsType"] = []string{}
	for _, item := range node.Params {
		err := _paramMeta(item.(*ParamNode))
		if err != nil {
			return err
		}
		node.Meta["paramsType"] = append(node.Meta["paramsType"].([]string), item.(*ParamNode).Meta["dataType"].(string))
		err = node.Meta.appendSymbol(item.(*ParamNode).Meta["name"].(string), VaribleDefinition{item.(*ParamNode).Meta["dataType"].(string)})
		if err != nil {
			return err
		}
	}
	for _, item := range node.Body {
		switch item := item.(type) {
		case *ConstNode:
			err := _constMeta(item)
			if err != nil {
				return err
			}
			err = node.Meta.appendSymbol(item.Meta["name"].(string), ConstantDefinition{item.Meta["type"].(string)})
			if err != nil {
				return err
			}
		case *VarNode:
			err := _varMeta(item)
			if err != nil {
				return err
			}
			err = node.Meta.appendSymbol(item.Meta["name"].(string), VaribleDefinition{item.Meta["type"].(string)})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _paramMeta(node *ParamNode) error {
	node.Meta = make(MetaData)
	node.Meta["name"] = resolveSymbol(node.Name.(*SymbolNode))
	dataType, err := resolveDataTypeFromNode(node.DataType)
	if err != nil {
		return err
	}
	node.Meta["dataType"] = dataType
	return nil
}

func resolveSymbol(node *SymbolNode) string {
	return node.Value.Value
}

func resolveValueFromNode(node Node) (string, error) {
	switch value := node.(type) {
	case *IntegerNode:
		return value.Value.Value, nil
	case *FloatNode:
		return value.Value.Value, nil
	case *BooleanNode:
		return value.Value.Value, nil
	case *StringNode:
		return value.Value.Value, nil
	default:
		return "", fmt.Errorf("can't use %T as const value", value)
	}
}

func resolveDataTypeFromNode(node Node) (string, error) {
	switch dataType := node.(type) {
	case *SymbolNode:
		return dataType.Value.Value, nil
	default:
		return "", fmt.Errorf("unsupported data type %T", node)
	}
}
