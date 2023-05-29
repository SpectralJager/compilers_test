package frontend

import (
	"encoding/json"
)

type Node interface {
	json.Marshaler
	node()
}

func (n ProgramNode) node()  {}
func (n ConstNode) node()    {}
func (n VarNode) node()      {}
func (n SetNode) node()      {}
func (n FunctionNode) node() {}
func (n WhileNode) node()    {}
func (n IfNode) node()       {}

func (n ExpressionNode) node() {}
func (n ParamNode) node()      {}

func (n IntegerNode) node() {}
func (n FloatNode) node()   {}
func (n StringNode) node()  {}
func (n BooleanNode) node() {}
func (n SymbolNode) node()  {}

type ProgramNode struct {
	Package string   `json:"package"`
	Body    []Node   `json:"body"`
	Meta    MetaData `json:"meta"`
}

func (progNode ProgramNode) MarshalJSON() ([]byte, error) {
	type tmp ProgramNode
	return toJson(tmp(progNode), "programNode")
}

type ConstNode struct {
	Name     Node     `json:"name"`
	DataType Node     `json:"dataType"`
	Value    Node     `json:"value"`
	Meta     MetaData `json:"meta"`
}

func (constNode ConstNode) MarshalJSON() ([]byte, error) {
	type tmp ConstNode
	return toJson(tmp(constNode), "constNode")
}

type VarNode struct {
	Name     Node     `json:"name"`
	DataType Node     `json:"dataType"`
	Value    Node     `json:"value"`
	Meta     MetaData `json:"meta"`
}

func (varNode VarNode) MarshalJSON() ([]byte, error) {
	type tmp VarNode
	return toJson(tmp(varNode), "varNode")
}

type SetNode struct {
	Name  Node     `json:"name"`
	Value Node     `json:"value"`
	Meta  MetaData `json:"meta"`
}

func (setNode SetNode) MarshalJSON() ([]byte, error) {
	type tmp SetNode
	return toJson(tmp(setNode), "setNode")
}

type FunctionNode struct {
	Name       Node   `json:"name"`
	ReturnType Node   `json:"returnType"`
	Params     []Node `json:"params"`
	Body       []Node `json:"body"`
}

func (funcNode FunctionNode) MarshalJSON() ([]byte, error) {
	type tmp FunctionNode
	return toJson(tmp(funcNode), "functionNode")
}

type WhileNode struct {
	ConditionExpressions Node     `json:"conditionExpressions"`
	ThenBody             []Node   `json:"thenBody"`
	Meta                 MetaData `json:"meta"`
}

func (whileNode WhileNode) MarshalJSON() ([]byte, error) {
	type tmp WhileNode
	return toJson(tmp(whileNode), "whileNode")
}

type IfNode struct {
	ConditionExpressions Node     `json:"conditionExpressions"`
	ThenBody             []Node   `json:"thenBody"`
	ElseBody             []Node   `json:"elseBody"`
	Meta                 MetaData `json:"meta"`
}

func (ifNode IfNode) MarshalJSON() ([]byte, error) {
	type tmp IfNode
	return toJson(tmp(ifNode), "ifNode")
}

type ParamNode struct {
	Name     Node     `json:"name"`
	DataType Node     `json:"dataType"`
	Meta     MetaData `json:"meta"`
}

func (paramNode ParamNode) MarshalJSON() ([]byte, error) {
	type tmp ParamNode
	return toJson(tmp(paramNode), "paramNode")
}

type ExpressionNode struct {
	Function Node     `json:"function"`
	Args     []Node   `json:"args"`
	Meta     MetaData `json:"meta"`
}

func (exprNode ExpressionNode) MarshalJSON() ([]byte, error) {
	type tmp ExpressionNode
	return toJson(tmp(exprNode), "expressionNode")
}

type IntegerNode struct {
	Value Token    `json:"value"`
	Meta  MetaData `json:"meta"`
}

func (intNode IntegerNode) MarshalJSON() ([]byte, error) {
	type tmp IntegerNode
	return toJson(tmp(intNode), "integerNode")
}

type FloatNode struct {
	Value Token    `json:"value"`
	Meta  MetaData `json:"meta"`
}

func (floatNode FloatNode) MarshalJSON() ([]byte, error) {
	type tmp FloatNode
	return toJson(tmp(floatNode), "floatNode")
}

type StringNode struct {
	Value Token    `json:"value"`
	Meta  MetaData `json:"meta"`
}

func (stringNode StringNode) MarshalJSON() ([]byte, error) {
	type tmp StringNode
	return toJson(tmp(stringNode), "stringNode")
}

type BooleanNode struct {
	Value Token    `json:"value"`
	Meta  MetaData `json:"meta"`
}

func (booleanNode BooleanNode) MarshalJSON() ([]byte, error) {
	type tmp BooleanNode
	return toJson(tmp(booleanNode), "booleanNode")
}

type SymbolNode struct {
	Value Token    `json:"value"`
	Meta  MetaData `json:"meta"`
}

func (symbolNode SymbolNode) MarshalJSON() ([]byte, error) {
	type tmp SymbolNode
	return toJson(tmp(symbolNode), "symbolNode")
}

// Helpers

func toJson(v any, tp string) ([]byte, error) {
	res, _ := json.Marshal(v)
	var temp map[string]interface{}
	json.Unmarshal(res, &temp)
	temp["type"] = tp
	return json.Marshal(temp)
}
