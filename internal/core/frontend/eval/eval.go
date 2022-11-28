package eval

import (
	"fmt"
	"grimlang/internal/core/frontend/ast"
	"grimlang/internal/core/frontend/tokens"
	"strconv"
)

func Eval(node ast.Node, env *Env) (interface{}, error) {
	switch node := node.(type) {
	case *ast.PrefixExpr:
		return evalPrefix(*node, env)
	case *ast.Number:
		return evalAtom(node)
	default:
		return nil, fmt.Errorf("unknow Node type: %+v", node)
	}
}

func evalPrefix(node ast.PrefixExpr, env *Env) (interface{}, error) {
	values := make([]interface{}, len(node.Args))
	for i, argument := range node.Args {
		res, err := Eval(argument, env)
		if err != nil {
			return nil, err
		}
		values[i] = res
	}

	switch node.Operator.Type {
	case tokens.Add:
		res := 0
		for _, val := range values {
			res += val.(int)
		}
		return res, nil
	default:
		return nil, fmt.Errorf("wrong prefix expression symbol")
	}
}

func evalAtom(node ast.Atom) (interface{}, error) {
	switch node := node.(type) {
	case *ast.Number:
		return strconv.Atoi(node.Token.Value)
	default:
		return nil, fmt.Errorf("unsupported atom type")
	}
}
