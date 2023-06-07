package frontend

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
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

type SymbolTable map[string]SymbolDefinition

/*
scope levels:
# contains buildin types, functions, ...
buildin |>

	# contains user defined global constants, varibles, functions, ...
	...global |>

		# contains user defined local varibles, constants, ...
		...local.
*/
func NewSymbolTable(scopeLevel string) SymbolTable {
	table := make(SymbolTable)
	if scopeLevel == "buildin" {
		// primitive data types
		table.AddSymbol("int", NewDataTypeDefinition("buildin"))
		table.AddSymbol("float", NewDataTypeDefinition("buildin"))
		table.AddSymbol("bool", NewDataTypeDefinition("buildin"))
		// integer functions
		table.AddSymbol("iadd", NewFuncDefinition("int", "...int"))
		table.AddSymbol("isub", NewFuncDefinition("int", "...int"))
		table.AddSymbol("imul", NewFuncDefinition("int", "...int"))
		table.AddSymbol("idiv", NewFuncDefinition("int", "...int"))
		table.AddSymbol("floatToInt", NewFuncDefinition("int", "float"))
		table.AddSymbol("iaddTwo", NewFuncDefinition("int", "int", "int"))
		//float functions
		table.AddSymbol("fadd", NewFuncDefinition("float", "...float"))
		table.AddSymbol("fsub", NewFuncDefinition("float", "...float"))
		table.AddSymbol("fmul", NewFuncDefinition("float", "...float"))
		table.AddSymbol("fdiv", NewFuncDefinition("float", "...float"))
		table.AddSymbol("intToFloat", NewFuncDefinition("float", "int"))
		// boolean functions
		table.AddSymbol("ilt", NewFuncDefinition("bool", "...int"))
	}
	return table
}

func (table SymbolTable) AddSymbol(symbol string, definition SymbolDefinition) error {
	if def, ok := table[symbol]; ok {
		return fmt.Errorf("symbol %s already exists with definition %s", symbol, def)
	}
	table[symbol] = definition
	return nil
}

func (table SymbolTable) ResolveSymbol(symbol string) (SymbolDefinition, error) {
	if def, ok := table[symbol]; ok {
		return def, nil
	}
	return nil, fmt.Errorf("symbol %s undefined", symbol)
}

type SymbolDefinition map[string]any

func (def SymbolDefinition) getType() string {
	res := def["type"].(string)
	return res
}

func NewFuncDefinition(returnType string, args ...string) SymbolDefinition {
	return SymbolDefinition{
		"type": "function",
		"args": args,
		"ret":  returnType,
	}
}

func NewVarDefinition(dataType string) SymbolDefinition {
	return SymbolDefinition{
		"type":     "varible",
		"dataType": dataType,
	}
}

func NewConstDefinition(dataType string, value string) SymbolDefinition {
	return SymbolDefinition{
		"type":     "constant",
		"dataType": dataType,
		"value":    value,
	}
}

func NewDataTypeDefinition(typeKind string) SymbolDefinition {
	return SymbolDefinition{
		"type": "dataType",
		"kind": typeKind,
	}
}

type SemanticContext map[string]any

func NewSemanticContext() SemanticContext {
	return SemanticContext{}
}

type SemanticAnalyser struct {
	scopeStack []string
	ctx        SemanticContext
}

func NewSemanticAnalyser(ctx SemanticContext) SemanticAnalyser {
	return SemanticAnalyser{
		scopeStack: []string{"buildin"},
		ctx:        ctx,
	}
}
func (s *SemanticAnalyser) pushScope(scope string) {
	s.scopeStack = append(s.scopeStack, scope)
}

func (s *SemanticAnalyser) top() string {
	return s.scopeStack[len(s.scopeStack)-1]
}

func (s *SemanticAnalyser) popScope() error {
	if s.top() == "buildin" && len(s.scopeStack) == 1 {
		return fmt.Errorf("can't remove builin scope")
	}
	s.scopeStack = s.scopeStack[:len(s.scopeStack)-1]
	return nil
}

func (s *SemanticAnalyser) currentSymbolTable() SymbolTable {
	return s.ctx[s.top()].(SymbolTable)
}

func (s *SemanticAnalyser) searchSymbol(symbol string) (SymbolDefinition, string, error) {
	for i := len(s.scopeStack) - 1; i >= 0; i-- {
		if def, ok := s.ctx[s.scopeStack[i]].(SymbolTable)[symbol]; ok {
			return def, s.scopeStack[i], nil
		}
	}
	return nil, "undefined", fmt.Errorf("undefined symbol %s", symbol)
}

func (s *SemanticAnalyser) CollectSymbols(node Node) error {
	switch node := node.(type) {
	case *ProgramNode:
		programmScope := fmt.Sprintf("global_%s", node.Package)
		s.pushScope(programmScope)
		node.ScopeName = programmScope
		s.ctx[programmScope] = NewSymbolTable("global")
		for _, item := range node.Body {
			err := s.CollectSymbols(item)
			if err != nil {
				return err
			}
		}
		s.popScope()
		return nil
	case *ConstNode:
		symbol := resolveSymbol(node.Name.(*SymbolNode))
		dataType := resolveDataType(node.DataType)
		s.currentSymbolTable().AddSymbol(symbol, NewConstDefinition(dataType, ""))
		return nil
	case *VarNode:
		symbol := resolveSymbol(node.Name.(*SymbolNode))
		dataType := resolveDataType(node.DataType)
		s.currentSymbolTable().AddSymbol(symbol, NewVarDefinition(dataType))
		return nil
	case *FunctionNode:
		symbol := resolveSymbol(node.Name.(*SymbolNode))
		returnType := resolveDataType(node.ReturnType)
		funcScope := fmt.Sprintf("function_%s", symbol)
		s.pushScope(funcScope)
		node.ScopeName = funcScope
		s.ctx[funcScope] = NewSymbolTable("local")
		argsDataType := []string{}
		for _, item := range node.Params {
			paramSymbol := resolveSymbol(item.(*ParamNode).Name.(*SymbolNode))
			dataType := resolveDataType(item.(*ParamNode).DataType)
			argsDataType = append(argsDataType, dataType)
			s.currentSymbolTable().AddSymbol(paramSymbol, NewVarDefinition(dataType))
		}
		s.popScope()
		s.currentSymbolTable().AddSymbol(symbol, NewFuncDefinition(returnType, argsDataType...))
		s.pushScope(funcScope)
		for _, item := range node.Body {
			err := s.CollectSymbols(item)
			if err != nil {
				return err
			}
		}
		s.popScope()
		return nil
	default:
		return nil
	}

}

func (s *SemanticAnalyser) Semantic(node Node) error {
	switch node := node.(type) {
	case *ProgramNode:
		programmScope := fmt.Sprintf("global_%s", node.Package)
		s.pushScope(programmScope)
		for _, item := range node.Body {
			err := s.Semantic(item)
			if err != nil {
				return err
			}
		}
		s.popScope()
		return nil
	case *ConstNode:
		symbol := resolveSymbol(node.Name.(*SymbolNode))
		dataType := resolveDataType(node.DataType)
		def, _, err := s.searchSymbol(dataType)
		if err != nil {
			return err
		}
		if def.getType() != "dataType" {
			return fmt.Errorf("%s is not dataType", dataType)
		}
		valueType, value, err := s.resolveValue(node.Value)
		if err != nil {
			return err
		}
		if dataType != valueType {
			return fmt.Errorf("invalid dataType, want %s, got %s", dataType, valueType)
		}
		if strings.Contains(s.top(), "global") {
			_, scope, err := s.searchSymbol(symbol)
			if scope == "buildin" {
				return fmt.Errorf("symbol %s is reserved", symbol)
			} else if scope != "undefined" {
				return err
			}
		}
		s.currentSymbolTable()[symbol] = NewConstDefinition(dataType, value)
		return nil
	case *VarNode:
		symbol := resolveSymbol(node.Name.(*SymbolNode))
		dataType := resolveDataType(node.DataType)
		def, _, err := s.searchSymbol(dataType)
		if err != nil {
			return err
		}
		if def.getType() != "dataType" {
			return fmt.Errorf("%s is not dataType", dataType)
		}
		valueType, _, err := s.resolveValue(node.Value)
		if err != nil {
			return err
		}
		if dataType != valueType {
			return fmt.Errorf("invalid dataType, want %s, got %s", dataType, valueType)
		}
		if strings.Contains(s.top(), "global") {
			_, scope, err := s.searchSymbol(symbol)
			if scope == "buildin" {
				return fmt.Errorf("symbol %s is reserved", symbol)
			} else if scope != "undefined" {
				return err
			}
		}
		s.currentSymbolTable()[symbol] = NewVarDefinition(dataType)
		return nil
	case *FunctionNode:
		symbol := resolveSymbol(node.Name.(*SymbolNode))
		returnType := resolveDataType(node.ReturnType)
		def, _, err := s.searchSymbol(returnType)
		if err != nil {
			return err
		}
		if def.getType() != "dataType" {
			return fmt.Errorf("%s is not dataType", returnType)
		}
		funcScope := fmt.Sprintf("function_%s", symbol)
		s.pushScope(funcScope)
		for _, item := range node.Params {
			dataType := resolveDataType(item.(*ParamNode).DataType)
			def, _, err := s.searchSymbol(dataType)
			if err != nil {
				return err
			}
			if def.getType() != "dataType" {
				return fmt.Errorf("%s is not dataType", dataType)
			}
		}
		for _, item := range node.Body {
			err := s.Semantic(item)
			if err != nil {
				return err
			}
		}
		s.popScope()
		return nil
	case *ExpressionNode:
		symbol := resolveSymbol(node.Function.(*SymbolNode))
		def, _, err := s.searchSymbol(symbol)
		if err != nil {
			return err
		}
		if def.getType() != "function" {
			return fmt.Errorf("symbol %s is not a function", symbol)
		}
		args := def["args"].([]string)
		for i, arg := range args {
			if strings.Contains(arg, "...") {
				if i != len(args)-1 {
					return fmt.Errorf("...type can be used only in last position")
				}
				tempType := strings.Split(arg, "...")[1]
				for j := i; j < len(node.Args); j++ {
					argType, _, err := s.resolveValue(node.Args[j])
					if err != nil {
						return err
					}
					if tempType != argType {
						return fmt.Errorf("argument #%d of wront valueType, want %s, get %s", j, tempType, argType)
					}
				}
			} else {
				argType, _, err := s.resolveValue(node.Args[i])
				if err != nil {
					return err
				}
				if arg != argType {
					return fmt.Errorf("argument #%d of wront valueType, want %s, get %s", i, arg, argType)
				}

			}
		}
		return nil
	case *IfNode:
		condValueType, _, err := s.resolveValue(node.ConditionExpressions)
		if err != nil {
			return err
		}
		if condValueType != "bool" {
			return fmt.Errorf("cant use %s as condition dataType, expected bool", condValueType)
		}
		token := GenerateToken(12)
		if token == "undefined" {
			return fmt.Errorf("cant generate token")
		}
		fnName := strings.Split(s.top(), "_")[1]
		node.ScopeName = fmt.Sprintf("%s_%s", fnName, token)
		scope := fmt.Sprintf("%s_ifBody_%s", fnName, token)
		s.pushScope(scope)
		s.ctx[scope] = NewSymbolTable("local")
		for _, item := range node.ThenBody {
			err := s.Semantic(item)
			if err != nil {
				return err
			}
		}
		s.popScope()
		scope = fmt.Sprintf("%s_elseBody_%s", fnName, token)
		s.pushScope(scope)
		s.ctx[scope] = NewSymbolTable("local")
		for _, item := range node.ElseBody {
			err := s.Semantic(item)
			if err != nil {
				return err
			}
		}
		s.popScope()
		return nil
	case *WhileNode:
		condValueType, _, err := s.resolveValue(node.ConditionExpressions)
		if err != nil {
			return err
		}
		if condValueType != "bool" {
			return fmt.Errorf("cant use %s as condition dataType, expected bool", condValueType)
		}
		token := GenerateToken(12)
		if token == "undefined" {
			return fmt.Errorf("cant generate token")
		}
		fnName := strings.Split(s.top(), "_")[1]
		node.ScopeName = fmt.Sprintf("%s_%s", fnName, token)
		scope := fmt.Sprintf("%s_while_%s", fnName, token)
		s.pushScope(scope)
		s.ctx[scope] = NewSymbolTable("local")
		for _, item := range node.ThenBody {
			err := s.Semantic(item)
			if err != nil {
				return err
			}
		}
		s.popScope()
		return nil
	default:
		return fmt.Errorf("unexpected type %T during semantic analysis", node)
	}
}

func resolveSymbol(node *SymbolNode) string {
	return node.Value.Value
}

func resolveDataType(node Node) string {
	switch node := node.(type) {
	case *SymbolNode:
		return resolveSymbol(node)
	default:
		return ""
	}
}

func (s *SemanticAnalyser) resolveValue(node Node) (string, string, error) {
	switch val := node.(type) {
	case *IntegerNode:
		return "int", val.Value.Value, nil
	case *FloatNode:
		return "float", val.Value.Value, nil
	case *BooleanNode:
		return "bool", val.Value.Value, nil
	case *SymbolNode:
		symbol := resolveSymbol(val)
		def, _, err := s.searchSymbol(symbol)
		if err != nil {
			return "undefined", "", err
		}
		switch def.getType() {
		case "constant", "varible":
			return def["dataType"].(string), symbol, nil
		default:
			return "undefined", "", fmt.Errorf("can't use %s of %s as value", symbol, def.getType())
		}
	case *ExpressionNode:
		def, _, err := s.searchSymbol(resolveSymbol(val.Function.(*SymbolNode)))
		if err != nil {
			return "undefined", "", err
		}
		if def.getType() != "function" {
			return "undefined", "", fmt.Errorf("expected function, got %s", def.getType())
		}
		err = s.Semantic(val)
		if err != nil {
			return "undefined", "", err
		}
		return def["ret"].(string), "", nil
	default:
		return "undefined", "", fmt.Errorf("unsupported valuet type %T", val)
	}
}

func GenerateToken(length int) string {
	data := make([]byte, length)
	if _, err := rand.Read(data); err != nil {
		return "undefined"
	}
	return hex.EncodeToString(data)
}
